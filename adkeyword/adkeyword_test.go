package adkeyword

import (
	"net/url"
	"testing"
)

// get add group id

func TestCreateAdKeywords(t *testing.T) {

}

func TestGetAdKeywords(t *testing.T) {
	params := url.Values{}
	params.Add("nccAdggroupId", "")
	r := GetAdKeywords(params)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}
