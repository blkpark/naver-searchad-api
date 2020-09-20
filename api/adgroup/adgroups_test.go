package adgroup

import (
	"encoding/json"
	"testing"

	"github.com/blkpark/naver-searchad-api/api/businesschannel"
)

func TestAdGroup(t *testing.T) {

}

func get() []byte {
	r := GetAdgroups(nil)
	if len(r) < 1 {
		panic(r)
	}
	return r
}

func post() []byte {

	// crete businesschannel
	b := businesschannel.GetBusinessChannel()
	bi :=
		json.Unmarshal(b, i)

	adg := &AdGroups{}
	adg.MobileChannelId = ""
	adg.Name = ""
	adg.NccCampaignId = ""
	adg.PcChannelId = ""
	r := PostAdgroups(nil, adg)
	if len(r) < 1 {
		panic(r)
	}
	return r
}
