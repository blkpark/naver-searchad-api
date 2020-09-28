package adgroup

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/blkpark/naver-searchad-api/searchad"
)

// AdGroups struct
type AdGroup struct {
	MobileChannelID string `json:"mobileChannelId"`
	Name            string `json:"name"`
	NccCampaignID   string `json:"nccCampaignId"`
	PcChannelID     string `json:"pcChannelId"`
}

// AdGroupResponse response
type AdGroupResponse struct {
	AdRollingType            string      `json:"adRollingType"`
	AdgroupAttrJSON          interface{} `json:"adgroupAttrJson"`
	AdgroupType              string      `json:"adgroupType"`
	BidAmt                   int64       `json:"bidAmt"`
	BudgetLock               bool        `json:"budgetLock"`
	ContentsNetworkBidAmt    int         `json:"contentsNetworkBidAmt"`
	ContentsNetworkBidWeight int64       `json:"contentsNetworkBidWeight"`
	CustomerID               int64       `json:"customerId"`
	DailyBudget              int64       `json:"dailyBudget"`
	DelFlag                  bool        `json:"delFlag"`
	EditTm                   string      `json:"editTm"`
	ExpectCost               int         `json:"expectCost"`
	KeywordPlusWeight        int         `json:"keywordPlusWeight"`
	MigType                  int         `json:"migType"`
	MobileChannelID          string      `json:"mobileChannelId"`
	MobileChannelKey         string      `json:"mobileChannelKey"`
	MobileNetworkBidWeight   int64       `json:"mobileNetworkBidWeight"`
	Name                     string      `json:"name"`
	NccAdgroupID             string      `json:"nccAdgroupId"`
	NccCampaignID            string      `json:"nccCampaignId"`
	PcChannelID              string      `json:"pcChannelId"`
	PcChannelKey             string      `json:"pcChannelKey"`
	PcNetworkBidWeight       int64       `json:"pcNetworkBidWeight"`
	RegTm                    string      `json:"regTm"`
	Status                   string      `json:"status"`
	StatusReason             string      `json:"statusReason"`
	SystemBiddingType        string      `json:"systemBiddingType"`
	TargetSummary            interface{} `json:"targetSummary"`
	UseCntsNetworkBidAmt     bool        `json:"useCntsNetworkBidAmt"`
	UseCntsNetworkBidWeight  bool        `json:"useCntsNetworkBidWeight"`
	UseDailyBudget           bool        `json:"useDailyBudget"`
	UseKeywordPlus           bool        `json:"useKeywordPlus"`
	UserLock                 bool        `json:"userLock"`
}

// BASE path
const BASE string = "/ncc/adgroups"

func getSelectors() []string {
	return []string{"PREVIOUS", "NEXT"}
}

// GetAdgroups
func List(params url.Values) ([]byte, error) {
	api := BASE

	var (
		nccLabelID = params.Get("nccLabelId")
		ids        = params.Get("ids")
	)

	if nccLabelID != "" && ids != "" {
		msg := "have to send either nccLabelID or ids"
		err := errors.New(msg)
		panic(err)
	}

	res, err := searchad.GetAPI(api, params)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// ListByCampaignID returns the list of partially AdGroups of input campaign id
//
// GET /ncc/adgroups{?nccCampaignId,baseSearchId,recordSize,selector}
//
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fadgroups%7B%3FnccCampaignId,baseSearchId,recordSize,selector%7D
func ListByCampaignID(params url.Values) ([]byte, error) {

	var (
		api           = BASE
		selector      = params.Get("selector")
		recordSize, _ = strconv.Atoi(params.Get("recordSize"))
	)

	if !(recordSize > 1 && recordSize < 1000) {
		msg := "recordSize should be enter between 1 to 1000."
		err := errors.New(msg)
		return nil, err
	}

	if searchad.CheckValidParameters(getSelectors, selector) {
		msg := "selector should be PREVIOUS or NEXT."
		err := errors.New(msg)
		return nil, err
	}

	res, err := searchad.GetAPI(api, params)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func Create(params url.Values, payload interface{}) ([]byte, error) {
	api := BASE
	res, err := searchad.PostAPI(api, params, payload)

	if err != nil {
		return nil, err
	}

	return res, nil
}
