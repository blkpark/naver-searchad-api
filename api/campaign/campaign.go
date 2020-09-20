package campaign

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

type Campaign struct {
	CampaignTp     string `json:"campaignTp,omitempty"`
	Name           string `json:"name,omitempty"`
	CustomerID     int64  `json:"customerId,omitempty"`
	DailyBudget    int64  `json:"dailyBudget,omitempty"`
	DeliveryMethod string `json:"deliveryMethod,omitempty"`
	PeriodEndDt    string `json:"periodEndDt,omitempty"`
	PeriodStartDt  string `json:"periodStartDt,omitempty"`
	TrackingMode   string `json:"trackingMode,omitempty"`
	UseDailyBudget bool   `json:"useDailyBudget"`
	UsePeriod      bool   `json:"usePeriod"`
	UserLock       bool   `json:"userLock"`
	TrackingURL    string `json:"trackingUrl,omitempty"`

	NccCampaignID string `json:"nccCampaignId,omitempty"` // put

	EditTm       string `json:"editTm,omitempty"`
	RegTm        string `json:"regTm,omitempty"`
	Status       string `json:"status,omitempty"`
	StatusReason string `json:"statusReason,omitempty"`
}

// BASE api path
const BASE string = "/ncc/campaigns"

// GetCampaignList /ncc/campaigns?campaignType,baseSearchId,recordSize,selector
func GetCampaignList(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

func PostCampaign(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}

func UpdateCampaign(params url.Values, payload interface{}, campaignId string) []byte {
	api := BASE + "/" + campaignId
	return searchad.PutAPI(api, params, payload)
}

func GetCampaignType() []string {
	s := make([]string)
}
