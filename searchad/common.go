package searchad

import (
	"encoding/json"
	"fmt"
	"log"
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
