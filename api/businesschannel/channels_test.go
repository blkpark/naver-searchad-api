package businesschannel

import (
	"net/url"
	"testing"

	"github.com/blkpark/naver-searchad-api/searchad"
)

func TestGetBusinessChannel(t *testing.T) {
	list()
	create()
}

// GET /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels
func list() {
	params := url.Values{}
	r := GetBusinessChannel(params)
	if len(r) < 1 {
		panic(r)
	}
	searchad.PrintJSON(r)
}

// POST /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/POST/~2Fncc~2Fchannels
func create() {
	b := BusinessChannel{}
	b.ChannelTp = "SITE"
	b.Name = "bc_" + searchad.RandomString(8)
	b.InspectRequestMsg = "test msg"

	m := Site{}
	m.Site = "http://www.my-site.com"
	b.BusinessInfo = m

	r := PostBusinessChannel(nil, b)
	searchad.PrintJSON(r)
}

// PUT /ncc/channels/{businessChannelId}{?fields}
// https://naver.github.io/searchad-apidoc/#/operations/PUT/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D%7B%3Ffields%7D
func update() {
	// TODO
}

// PUT /ncc/channels/{businessChannelId}/inspect
// https://naver.github.io/searchad-apidoc/#/operations/PUT/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D~2Finspect
func requestInspect() {
	// TODO
}

// DELETE /ncc/channels/{businessChannelId}
// https://naver.github.io/searchad-apidoc/#/operations/DELETE/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D
func delete() {
	// TODO
}

// DELETE /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/DELETE/~2Fncc~2Fchannels%7B%3Fids%7D
func deleteByIDs() {
	// TODO
}

// GET /ncc/channels{?channelTp}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels%7B%3FchannelTp%7D
func listByChannelTp() {
	// TODO
}

// GET /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels%7B%3Fids%7D
func listByIds() {
	// TODO
}

// GET /ncc/channels/{businessChannelId}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D
func get() {
	// TODO
}

// GET /ncc/purchasable-place-channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fpurchasable-place-channels
func getPurchasablePlaceChannels() {
	// TODO
}
