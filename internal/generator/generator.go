package generator

import (
	"encoding/base64"
	"log"
	"net/url"
	"strconv"
	"time"
)

func IsUrlValid(rawUrl string) bool {
	parsed, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		log.Printf("Cannot parse url: %s - Error: %v\n", rawUrl, err)
		return false
	}

	log.Printf("Parsed url: %s\n", parsed)
	return true
}

func GenerateKey() string {
	millis := time.Now().UnixNano() % 100000
	stringMillis := strconv.FormatInt(millis, 10)
	return base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString([]byte(stringMillis))
}
