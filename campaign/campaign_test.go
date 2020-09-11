package campaign

import (
	"math/rand"
	"net/url"
	"strconv"
	"testing"

	"github.com/blkpark/naver-searchad-api-golang/searchad"
)

func TestPostCampaign(t *testing.T) {
	c := &Campaign{}
	c.CampaignTp = "SHOPPING"
	c.Name = "api_test_" + strconv.Itoa(rand.Int())

	customerID := searchad.Env("CUSTOMER_ID", "")
	c.CustomerId = customerID

	r := PostCampaign(nil, c)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}

func TestGetCampaignList(t *testing.T) {
	p := url.Values{}
	p.Add("campaignType", "SHOPPING")

	r := GetCampaignList(p)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}
