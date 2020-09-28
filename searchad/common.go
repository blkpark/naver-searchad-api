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

// PrintJSON is print pretty JSON.
func PrintJSON(b []byte) {
	var o interface{}
	json.Unmarshal(b, &o)

	js, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(js))
}

// PrintInterface print interface.
func PrintInterface(i interface{}) {
	b, _ := json.Marshal(i)
	PrintJSON(b)
}

// Env return os environment.
func Env(key string, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	} else {
		return v
	}
}

// RandomString returns random length string.
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

// CheckValidParameter returns whether param is included.
func CheckValidParameter(f func() []string, param string) bool {

	if param == "" {
		return true
	}

	types := f()

	for _, val := range types {
		if val == param {
			return true
		}
	}
	return false
}
