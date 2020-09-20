package searchad

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
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

func RandomString(l int) string {
	if l == 0 {
		return ""
	}

	c := int(math.Round(float64(l) / 2))
	b := make([]byte, c)
	rand.Read(b)
	s := hex.EncodeToString(b)
	return s[:l]
}
