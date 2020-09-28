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

func TestBusinessChannel(t *testing.T) {

	// test create
	created := create()

	// test get
	finded := get(created.NccBusinessChannelID)

	// check made id.
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
	delete(created.NccBusinessChannelID)

	// check list
	list()
	listByChannelTp()

	// delete ids
	// created channels
	count := 3
	var ids []string

	for i := 0; i < count; i++ {
		c := create()
		ids = append(ids, c.NccBusinessChannelID)
	}

	// delete ids
	listByIds(ids)
	deleteByIDs(ids)
}

// GET /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels
func list() []BusinessChannel {

	// request
	r, err := List(nil)

	// check error
	if err != nil {
		panic(err)
	}

	// set return value
	var chs []BusinessChannel
	json.Unmarshal(r, &chs)

	return chs
}

// POST /ncc/channels
// https://naver.github.io/searchad-apidoc/#/operations/POST/~2Fncc~2Fchannels
func create() BusinessChannel {

	// make payload
	newCh := BusinessChannel{}
	newCh.ChannelTp = "SITE"
	str := searchad.RandomString(8)
	newCh.Name = "bc_" + str
	newCh.InspectRequestMsg = "test msg"

	site := Site{}
	site.Site = "http://www." + str + ".com"
	newCh.BusinessInfo = site

	// request
	created, err := Create(nil, newCh)

	// check err
	if err != nil {
		panic(err)
	}

	// set return value
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

	// request
	updated, err := Update(params, bc, bc.NccBusinessChannelID)

	// check err
	if err != nil {
		panic(err)
	}

	// set return value
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
func delete(businessChannelID string) {
	Delete(businessChannelID)
}

// DELETE /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/DELETE/~2Fncc~2Fchannels%7B%3Fids%7D
func deleteByIDs(ids []string) {

	// set params
	params := url.Values{}

	// add params
	for _, v := range ids {
		params.Add("ids", v)
	}

	// request
	r, err := DeleteIds(params)

	// check err
	if err != nil {
		panic(err)
	}

	searchad.PrintJSON(r)
}

// GET /ncc/channels{?channelTp}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels%7B%3FchannelTp%7D
func listByChannelTp() []BusinessChannel {

	// set channel type
	params := url.Values{}
	params.Add("channelTp", "SITE")

	// get list
	r, err := List(params)

	// check err
	if err != nil {
		panic(err)
	}

	// set return value
	var chs []BusinessChannel
	json.Unmarshal(r, &chs)

	return chs
}

// GET /ncc/channels{?ids}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels%7B%3Fids%7D
func listByIds(ids []string) []BusinessChannel {

	// set ids
	params := url.Values{}

	for _, v := range ids {
		params.Add("ids", v)
	}

	// get list
	r, err := List(params)
	// searchad.PrintJSON(r)

	// check err
	if err != nil {
		panic(err)
	}

	// set return value
	var chs []BusinessChannel
	json.Unmarshal(r, &chs)

	return chs
}

// GET /ncc/channels/{businessChannelId}
// https://naver.github.io/searchad-apidoc/#/operations/GET/~2Fncc~2Fchannels~2F%7BbusinessChannelId%7D
func get(businessChannelID string) BusinessChannel {

	// request
	r, err := Get(businessChannelID)

	// check err
	if err != nil {
		panic(err)
	}

	// set return value
	b := BusinessChannel{}
	json.Unmarshal(r, &b)
	return b
}
