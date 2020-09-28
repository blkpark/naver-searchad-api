package searchad

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAPI return http GET reponse.
func GetAPI(path string, params url.Values) ([]byte, error) {

	// get headers
	s := initialize(http.MethodGet, path, params, nil)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// PostAPI return http post response.
func PostAPI(path string, params url.Values, payload interface{}) ([]byte, error) {

	// get headers
	s := initialize(http.MethodPost, path, params, payload)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return data, nil
}

// DeleteAPI return http delete response.
func DeleteAPI(path string, params url.Values) ([]byte, error) {

	// get headers
	s := initialize(http.MethodDelete, path, params, nil)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	return data, nil
}

// PutAPI return http put response.
func PutAPI(path string, params url.Values, payload interface{}) ([]byte, error) {

	// get headers
	s := initialize(http.MethodPut, path, params, payload)

	// send
	res := s.request()

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}
