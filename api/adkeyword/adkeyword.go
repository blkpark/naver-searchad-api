package adkeyword

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

type AdKeyword struct {
	BidAmt  int64  // 70 ~ 100,000 default 70
	Keyword string // require
	Links   AdKeywordLink
}

type AdKeywordLink struct {
	Pc     string
	Mobile string
}

const BASE string = "/ncc/keywords"

func CreateAdKeywords(params url.Values, payload url.Values) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}

func GetAdKeywords(params url.Values) []byte {
	api := "/ncc/keywords"
	return searchad.GetAPI(api, params)
}
