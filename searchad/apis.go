package searchad

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GetAPI(path string, params url.Values) []byte {

	// get headers
	s := initialize(http.MethodGet, path, params, nil)

	// send
	res, err := s.request(http.MethodGet)
	if err != nil {
		panic(err)
	}

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return data
}

func PostAPI(path string, params url.Values, payload interface{}) []byte {
	// get headers

	s := initialize(http.MethodPost, path, params, payload)

	// send
	res, err := s.request(http.MethodPost)
	if err != nil {
		log.Fatal(res)
		panic(err)
	}

	// read body
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return data
}

func DeleteAPI() {

}

func PutAPI() {

}
