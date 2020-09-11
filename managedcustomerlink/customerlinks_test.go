package managedcustomerlink

import "testing"

func TestGetCustomerLinks(t *testing.T) {
	params := make(map[string][]string)
	//params["type"] = []string{"myclients"}
	r := GetCustomerLinks(params)
	if len(r) < 1 {
		panic(r)
	}
	//	searchad.PrintJSON(r)
}
