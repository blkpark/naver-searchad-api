package searchad

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

type Licence struct {
	Base       string
	APIKey     string
	SecretKey  string
	CustomerID string
}

// GetSignature return sha256-hmac (encoded base64)
func (l *Licence) GetSignature(method string, path string, timestamp string) string {
	// create hmac with secret key
	hash := hmac.New(sha256.New, []byte(l.SecretKey))

	// set message
	message := timestamp + "." + method + "." + path

	// Write message to hash
	hash.Write([]byte(message))
	sha := hash.Sum(nil)

	// Encode base64
	b64 := base64.StdEncoding.EncodeToString([]byte(sha))

	return b64
}

// GetLicense return environment variables for naver search ad.
// http://searchad.naver.com/
func GetLicense() (*Licence, error) {
	base := "https://api.naver.com"
	apiKey := Env("API_KEY", "")
	secretKey := Env("SECRET_KEY", "")
	customerID := Env("CUSTOMER_ID", "")

	var err error

	for {
		if apiKey == "" {
			err = errors.New("apiKey not found.")
			break
		}

		if secretKey == "" {
			err = errors.New("secretKey not found")
			break
		}

		if customerID == "" {
			err = errors.New("customerID not found")
		}

		break
	}

	if err != nil {
		return nil, err
	}

	l := &Licence{
		Base:       base,
		APIKey:     apiKey,
		SecretKey:  secretKey,
		CustomerID: customerID,
	}

	return l, nil
}
