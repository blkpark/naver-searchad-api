package relkwdstat

import (
	"net/url"
	"testing"
)

func TestGetKeywordsTool(t *testing.T) {
	params := url.Values{}
	params.Add("hintKeywords", "마스크")
	params.Add("showDetail", "0")

	r := GetKeywordsTool(params)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}
