package searchad

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SearchAd struct {
	URL         string
	Method      string
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
		l.Base + path, method, "application/json; charset=UTF-8", now, l.APIKey, l.CustomerID, sign, nil, nil,
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
func (s SearchAd) request() *http.Response {

	// add payload
	buf := new(bytes.Buffer)
	if s.Payload != nil {
		json.NewEncoder(buf).Encode(s.Payload)
	}

	// make new http request
	req, _ := http.NewRequest(s.Method, s.URL, buf)

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
	res, err := client.Do(req)

	if err != nil {
		// request
		fmt.Println("----------- request ----------")
		PrintInterface(s.Payload)
		panic(err)
	}

	// error check
	if res.StatusCode/100 != 2 {
		// response
		r, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		fmt.Println("----------- response ----------")
		PrintJSON(r)

		// request
		fmt.Println("----------- request ----------")
		fmt.Println(s.Method, s.URL, s.Params)
		PrintInterface(s.Payload)

		panic(nil)
	}
	return res
}
