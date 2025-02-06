package scb_inkvartering_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	scb_inkvartering "github.com/omniboost/go-scb-inkvartering"
)

func TestPostSurveyUploadCSVRequest(t *testing.T) {
	f, err := os.ReadFile("survey-upload-csv-test.txt")
	if err != nil {
		t.Fatal(err)
	}

	req := client.NewPostSurveyUploadCSVRequest()
	req.PathParams().Period = scb_inkvartering.DateTime{time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}
	req.RequestBody().RawCSV = string(f)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
