package adgroup

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

type AdGroups struct {
	MobileChannelId string `json:"mobileChannelId"`
	Name            string `json:"name"`
	NccCampaignId   string `json:"nccCampaignId"`
	PcChannelId     string `json:"pcChannelId"`
}

const BASE string = "/ncc/adgroups"

// // GetAdgroups
func GetAdgroups(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

func PostAdgroups(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}
