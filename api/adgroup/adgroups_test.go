package adgroup

import (
	"testing"

	"github.com/blkpark/naver-searchad-api/searchad"
)

func TestAdGroup(t *testing.T) {
	get()
}

func get() []byte {
	res, err := List(nil)

	if err != nil {
		panic(err)
	}

	if len(res) < 1 {
		panic(res)
	}

	searchad.PrintJSON(res)
	return res
}

// func post() []byte {

// 	// crete businesschannel
// 	var i interface{}

// 	b, err := businesschannel.List(nil)
// 	bi := json.Unmarshal(b, &i)

// adg := &AdGroups{}
// adg.MobileChannelId = ""
// adg.Name = ""
// adg.NccCampaignId = ""
// adg.PcChannelId = ""
// r := PostAdgroups(nil, adg)
// if len(r) < 1 {
// 	panic(r)
// }
// return r
// }
