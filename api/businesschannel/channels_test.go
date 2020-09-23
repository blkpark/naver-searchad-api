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
	created := create()
	finded := get(created.NccBusinessChannelID)

	if !reflect.DeepEqual(created.NccBusinessChannelID, finded.NccBusinessChannelID) {
		msg := "\nbusiness Channel is not equel" + "\ncreated:" + created.NccBusinessChannelID + "\nfinded" + finded.NccBusinessChannelID
		err := errors.New(msg)
		log.Fatal(msg)
		panic(err)
	}

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
}

// GET /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels
func list() {
	params := url.Values{}
	r := List(params)
	if len(r) < 1 {
		panic(r)
	}
	searchad.PrintJSON(r)
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

	r := Create(nil, newCh)
	searchad.PrintJSON(r)

	json.Unmarshal(r, &newCh)

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

	r := Update(params, bc, bc.NccBusinessChannelID)
	searchad.PrintJSON(r)

	rbc := BusinessChannel{}
	json.Unmarshal(r, &rbc)

	return rbc
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
