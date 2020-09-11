package managedcustomerlink

import (
	"net/url"

	"github.com/blkpark/naver-searchad-api/searchad"
)

// GetCustomerLinks
func GetCustomerLinks(params url.Values) []byte {
	api := "/customer-links"
	return searchad.GetAPI(api, params)
}
