package businesschannel

import (
	"errors"
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
	channelTp := params.Get("channelTp")
	ids := params.Get("ids")

	if channelTp != "" && ids != "" {
		msg := "have to send either channelTp or ids"
		err := errors.New(msg)
		panic(err)
	}

	if !checkSupportedChannelTp(channelTp) {
		msg := "not supported channel type."
		err := errors.New(msg)
		panic(err)
	}

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

func Delete(businessChannelID string) []byte {
	api := BASE + "/" + businessChannelID
	return searchad.DeleteAPI(api, nil)
}

func DeleteIds(params url.Values) []byte {
	api := BASE
	return searchad.DeleteAPI(api, params)
}

func checkSupportedChannelTp(name string) bool {

	if name == "" {
		return true
	}

	types := []string{"SITE", "PHONE", "ADDRESS", "BOOKING", "TALK", "MALL", "CONTENTS", "PLACE", "CATALOG"}

	for _, val := range types {
		if val == name {
			return true
		}
	}
	return false
}
