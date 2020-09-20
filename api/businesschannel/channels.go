package businesschannel

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

type Site struct {
	Site string `json:"site"`
}

const BASE string = "/ncc/channels"

type BusinessChannel struct {
	ChannelTp         string `json:"channelTp,omitempty"`         // post
	Name              string `json:"name,omitempty"`              // post
	InspectRequestMsg string `json:"inspectRequestMsg,omitempty"` //post
	BusinessInfo      Site   `json:"businessInfo,omitempty"`      // post
	// nidAut
	// nidSes
	// passNaTokenToSa

	AdultStatus          string `json:"adultStatus,omitempty"`
	BlackStatus          string `json:"blackStatus,omitempty"`
	ChannelKey           string `json:"channelKey,omitempty"`
	CustomerID           int64  `json:"customerId,omitempty"`
	EditTm               string `json:"editTm,omitempty"`
	FirstChargeTm        string `json:"firstChargeTm,omitempty"`
	InspectRequestTm     string `json:"inspectRequestTm,omitempty"`
	InspectTm            string `json:"inspectTm,omitempty"`
	MobileInspectStatus  string `json:"mobileInspectStatus,omitempty"`
	NccBusinessChannelID string `json:"nccBusinessChannelId,omitempty"`
	PcInspectStatus      string `json:"pcInspectStatus,omitempty"`
	RegTm                string `json:"regTm,omitempty"`
	Status               string `json:"status,omitempty"`
	StatusReason         string `json:"statusReason,omitempty"`
}

// GetBusinessChannel
func GetBusinessChannel(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

func PostBusinessChannel(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}
