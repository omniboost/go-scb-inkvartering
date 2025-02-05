package inkvartering_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	inkvartering "github.com/omniboost/go-inkvartering"
)

var (
	client *inkvartering.Client
)

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	apiKey := os.Getenv("INKVARTERING_API_KEY")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	client = inkvartering.NewClient(nil)
	client.SetApiKey(apiKey)

	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	m.Run()
}
