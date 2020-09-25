package searchad

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func PrintJSON(b []byte) {
	var o interface{}
	json.Unmarshal(b, &o)

	js, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(js))
}

func PrintInterface(i interface{}) {
	b, _ := json.Marshal(i)
	PrintJSON(b)
}

func Env(key string, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	} else {
		return v
	}
}

func RandomString(len int) string {

	// if length 0
	if len == 0 {
		return ""
	}

	// re-seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// make string
	c := int(math.Round(float64(len) / 2))
	b := make([]byte, c)
	r.Read(b)
	s := hex.EncodeToString(b)
	return s[:len]
}
