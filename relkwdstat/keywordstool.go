package relkwdstat

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

const BASE = "/keywordstool"

// GetKeywordsTool https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fkeywordstool
func GetKeywordsTool(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}
