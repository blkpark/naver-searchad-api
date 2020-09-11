package searchad

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SearchAd struct {
	URL         string
	ContentType string
	Timestamp   string
	APIKey      string
	CustomerID  string
	Signature   string
	Params      url.Values
	Payload     interface{}
}

// get headers
func initialize(method string, path string, params url.Values, payload interface{}) *SearchAd {

	// get naver api licence
	l, err := GetLicense()
	if err != nil {
		panic(err)
	}

	// get signature
	now := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	sign := l.GetSignature(method, path, now)

	m := &SearchAd{
		l.Base + path, "application/json; charset=UTF-8", now, l.APIKey, l.CustomerID, sign, nil, nil,
	}

	if params != nil {
		m.Params = params
	}

	if payload != nil {
		m.Payload = payload
	}

	return m
}

// send http request
func (s SearchAd) request(method string) (*http.Response, error) {

	// add payload
	buf := new(bytes.Buffer)
	if s.Payload != nil {
		json.NewEncoder(buf).Encode(s.Payload)
	}

	// make new http request
	req, _ := http.NewRequest(method, s.URL, buf)

	// add query string
	if s.Params != nil {
		req.URL.RawQuery = s.Params.Encode()
	}

	// add headers
	req.Header.Add("Content-Type", s.ContentType)
	req.Header.Add("X-Timestamp", s.Timestamp)
	req.Header.Add("X-API-KEY", s.APIKey)
	req.Header.Add("X-Customer", s.CustomerID)
	req.Header.Add("X-Signature", s.Signature)

	// send
	client := &http.Client{}
	res, _ := client.Do(req)

	// error check
	if res.StatusCode/100 != 2 {
		statusCode := strconv.Itoa(res.StatusCode)

		// response
		r, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		PrintJSON(r)

		// request
		b, _ := ioutil.ReadAll(req.Body)
		PrintJSON(b)

		return nil, errors.New(statusCode)
	}
	return res, nil
}
