package businesschannel

import (
	"net/url"
	"testing"
)

func TestGetBusinessChannel(t *testing.T) {
	params := url.Values{}
	r := GetBusinessChannel(params)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}
