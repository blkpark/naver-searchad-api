package campaign

import (
	"encoding/json"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/blkpark/naver-searchad-api/searchad"
)

func TestCampaign(t *testing.T) {
	get()
	postCampaign := post()
	put(postCampaign)
}

func post() []byte {
	c := &Campaign{}
	c.CampaignTp = "SHOPPING"
	now := strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	c.Name = "api_test_" + now
	c.CustomerID, _ = strconv.ParseInt(searchad.Env("CUSTOMER_ID", ""), 10, 64)

	r := PostCampaign(nil, c)
	if len(r) < 1 {
		panic(r)
	}
	searchad.PrintJSON(r)

	return r
}

func get() []byte {
	p := url.Values{}
	p.Add("campaignType", "SHOPPING")

	r := GetCampaignList(p)
	if len(r) < 1 {
		panic(r)
	}
	searchad.PrintJSON(r)

	return r
}

func put(b []byte) []byte {
	c := Campaign{}
	json.Unmarshal(b, &c)
	// c.UserLock = false

	params := url.Values{}
	params.Add("field", "userLock,budget,period")

	r := UpdateCampaign(params, c, c.NccCampaignID)
	if len(r) < 1 {
		panic(r)
	}

	searchad.PrintJSON(r)
	return r
}
