package adgroup

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

// AdGroups struct
type AdGroups struct {
	MobileChannelID string `json:"mobileChannelId"`
	Name            string `json:"name"`
	NccCampaignID   string `json:"nccCampaignId"`
	PcChannelID     string `json:"pcChannelId"`
}

// BASE path
const BASE string = "/ncc/adgroups"

// GetAdgroups
func GetAdgroups(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

func PostAdgroups(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}
