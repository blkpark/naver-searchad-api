package businesschannel

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

const BASE string = "/ncc/channels"

type Site struct {
	Site string `json:"site"`
}

// BusinessChannel
type BusinessChannel struct {
	ChannelTp         string `json:"channelTp,omitempty"`         // post
	Name              string `json:"name,omitempty"`              // post
	InspectRequestMsg string `json:"inspectRequestMsg,omitempty"` // post
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

// List
func List(params url.Values) []byte {
	api := BASE
	return searchad.GetAPI(api, params)
}

func Get(nccBusinessChannelID string) []byte {
	api := BASE + "/" + nccBusinessChannelID
	return searchad.GetAPI(api, nil)
}

func Create(params url.Values, payload interface{}) []byte {
	api := BASE
	return searchad.PostAPI(api, params, payload)
}

func Update(params url.Values, payload interface{}, businessChannelID string) []byte {
	api := BASE + "/" + businessChannelID
	return searchad.PutAPI(api, params, payload)

}
