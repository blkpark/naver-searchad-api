package campaign

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

type Campaign struct {
	CampaignTp     string `json:"campaignTp"`
	Name           string `json:"name"`
	CustomerId     string `json:"customerId"`
	DailyBudget    int64  `json:"dailyBudget,omitempty"`
	DeliveryMethod string `json:"deliveryMethod,omitempty"`
	PeriodEndDt    string `json:"periodEndDt,omitempty"`
	PeriodStartDt  string `json:"periodStartDt,omitempty"`
	TrackingMode   string `json:"yrackingMode,omitempty"`
	UseDailyBudget bool   `json:"useDailyBudget,omitempty"`
	UsePeriod      bool   `json:"usePeriod,omitempty"`
	UserLock       bool   `json:"userLock,omitempty"`
}

// BASE api path
const BASE string = "/ncc/campaigns"

// GetCampaignList /ncc/campaigns?campaignType,baseSearchId,recordSize,selector
func GetCampaignList(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

// not working
func PostCampaign(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}
