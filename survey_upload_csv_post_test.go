package scb_inkvartering_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestPostSurveyUploadCSVRequest(t *testing.T) {
	f, err := os.ReadFile("survey-upload-csv-tes.txt")
	if err != nil {
		t.Fatal(err)
	}

	req := client.NewPostSurveyUploadCSVRequest()
	req.PathParams().Period = "YYYYMM"
	req.RequestBody().RawCSV = string(f)
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
