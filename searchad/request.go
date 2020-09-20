package searchad

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetAPI(path string, params url.Values) []byte {

	// get headers
	s := initialize(http.MethodGet, path, params, nil)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func PostAPI(path string, params url.Values, payload interface{}) []byte {

	// get headers
	s := initialize(http.MethodPost, path, params, payload)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func DeleteAPI() {

}

func PutAPI(path string, params url.Values, payload interface{}) []byte {

	// get headers
	s := initialize(http.MethodPut, path, params, payload)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}

	return data
}
