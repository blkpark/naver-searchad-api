package adgroup

import (
	"testing"
)

func TestGetAdgroups(t *testing.T) {
	GetAdgroups(nil)
	//searchad.PrintJSON(GetAdgroups(nil))
}

func TestPostAdgroups(t *testing.T) {
	adg := &AdGroups{}
	adg.MobileChannelId = ""
	adg.Name = ""
	adg.NccCampaignId = ""
	adg.PcChannelId = ""
	r := PostAdgroups(nil, adg)
	if len(r) < 1 {
		panic(r)
	}
	// searchad.PrintJSON(r)
}
