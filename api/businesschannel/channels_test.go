package businesschannel

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"reflect"
	"testing"

	"github.com/blkpark/naver-searchad-api/searchad"
)

func TestGetBusinessChannel(t *testing.T) {

	// test create
	created := create()

	// test get
	finded := get(created.NccBusinessChannelID)

	if !reflect.DeepEqual(created.NccBusinessChannelID, finded.NccBusinessChannelID) {
		msg := "\nbusiness Channel is not equel" + "\ncreated:" + created.NccBusinessChannelID + "\nfinded" + finded.NccBusinessChannelID
		err := errors.New(msg)
		log.Fatal(msg)
		panic(err)
	}

	// test update
	marshaled, _ := json.Marshal(finded)
	updated := update(marshaled)
	findAgain := get(created.NccBusinessChannelID)

	// check update
	if !reflect.DeepEqual(updated.Name, findAgain.Name) {
		msg := "\nbusiness channel name not changed.\nupdated:" + updated.Name + "\nfind again:" + findAgain.Name
		err := errors.New(msg)
		log.Fatal(msg)
		panic(err)
	}

	// check delete

	// check list
	list()
	listByChannelTp()
}

// GET /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels
func list() {
	r := List(nil)
	if len(r) < 1 {
		panic(r)
	}
	//searchad.PrintJSON(r)
}

// POST /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/POST/~2Fncc~2Fchannels
func create() BusinessChannel {
	newCh := BusinessChannel{}
	newCh.ChannelTp = "SITE"
	str := searchad.RandomString(8)
	newCh.Name = "bc_" + str
	newCh.InspectRequestMsg = "test msg"

	site := Site{}
	site.Site = "http://www." + str + ".com"
	newCh.BusinessInfo = site

	created := Create(nil, newCh)
	//searchad.PrintJSON(created)

	json.Unmarshal(created, &newCh)

	return newCh
}

// PUT /ncc/channels/{businessChannelId}{?fields}
// https://naver.github.io/searchad-apidoc/#/operations/PUT/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D%7B%3Ffields%7D
func update(b []byte) BusinessChannel {

	// unmarshal
	bc := BusinessChannel{}
	json.Unmarshal(b, &bc)

	// change name
	bc.Name = "changed_" + bc.Name

	// set params
	params := url.Values{}
	params.Add("fields", "name")

	updated := Update(params, bc, bc.NccBusinessChannelID)
	//searchad.PrintJSON(updated)

	r := BusinessChannel{}
	json.Unmarshal(updated, &r)

	return r
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
func listByChannelTp() []BusinessChannel {

	// set channel type
	params := url.Values{}
	params.Add("channelTp", "SITE")

	// get list
	r := List(params)

	// set return value
	var chs []BusinessChannel
	json.Unmarshal(r, &chs)

	// for _, v := range chs {
	// 	marshaled, _ := json.Marshal(v)
	// 	searchad.PrintJSON(marshaled)
	// }

	return chs
}

// GET /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels%7B%3Fids%7D
func listByIds() []BusinessChannel {

	// set ids
	params := url.Values{}
	params.Add("ids", "")

	// get list
	r := List(params)

	// set return value
	var chs []BusinessChannel
	json.Unmarshal(r, &chs)

	return chs
}

// GET /ncc/channels/}{businessChannelId
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D
func get(businessChannelID string) BusinessChannel {
	r := Get(businessChannelID)
	b := BusinessChannel{}
	json.Unmarshal(r, &b)
	return b
}

// GET /ncc/purchasable-place-channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fpurchasable-place-channels
func getPurchasablePlaceChannels() {
	// TODO
}
