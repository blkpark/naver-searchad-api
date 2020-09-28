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

// List returns the list of business channel.
// Send a channelTp to the parameter, it returns the business channel list of channel type.
// Send a ids to the parameter, it returns the business channel list of input business channel ids.

// GET /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels
//
func List(params url.Values) ([]byte, error) {
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

	res, err := searchad.GetAPI(api, params)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Get(nccBusinessChannelID string) ([]byte, error) {
	api := BASE + "/" + nccBusinessChannelID
	res, err := searchad.GetAPI(api, nil)

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

func Update(params url.Values, payload interface{}, businessChannelID string) ([]byte, error) {
	api := BASE + "/" + businessChannelID
	res, err := searchad.PutAPI(api, params, payload)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// DELETE /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/DELETE/~2Fncc~2Fchannels%7B%3Fids%7D
func Delete(businessChannelID string) ([]byte, error) {
	api := BASE + "/" + businessChannelID
	res, err := searchad.DeleteAPI(api, nil)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteIds Remove business channels in id list.
//
// Request: DELETE /ncc/channels/{businessChannelId}
//
// Ref: https://naver.github.io/searchad-apidoc/#/operations/DELETE/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D
func DeleteIds(params url.Values) ([]byte, error) {
	api := BASE

	// check parameters.
	if params.Get("ids") == "" {
		msg := "have to send ids"
		err := errors.New(msg)
		return nil, err
	}

	res, err := searchad.DeleteAPI(api, params)
	if err != nil {
		return nil, err
	}

	return res, nil
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
