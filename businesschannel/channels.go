package businesschannel

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

// GetBusinessChannel
func GetBusinessChannel(params url.Values) []byte {
	api := "/ncc/channels"
	return searchad.GetAPI(api, params)
}
