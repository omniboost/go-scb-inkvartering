package scb_inkvartering

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/omniboost/go-scb-inkvartering/utils"
)

func (c *Client) NewPostSurveyUploadCSVRequest() PostSurveyUploadCSVRequest {
	return PostSurveyUploadCSVRequest{
		client:      c,
		queryParams: c.NewPostSurveyUploadCSVQueryParams(),
		pathParams:  c.NewPostSurveyUploadCSVPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewPostSurveyUploadCSVRequestBody(),
	}
}

type PostSurveyUploadCSVRequest struct {
	client      *Client
	queryParams *PostSurveyUploadCSVQueryParams
	pathParams  *PostSurveyUploadCSVPathParams
	method      string
	headers     http.Header
	requestBody PostSurveyUploadCSVRequestBody
}

func (c *Client) NewPostSurveyUploadCSVQueryParams() *PostSurveyUploadCSVQueryParams {
	return &PostSurveyUploadCSVQueryParams{}
}

type PostSurveyUploadCSVQueryParams struct{}

func (p PostSurveyUploadCSVQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *PostSurveyUploadCSVRequest) QueryParams() *PostSurveyUploadCSVQueryParams {
	return r.queryParams
}

func (c *Client) NewPostSurveyUploadCSVPathParams() *PostSurveyUploadCSVPathParams {
	return &PostSurveyUploadCSVPathParams{}
}

type PostSurveyUploadCSVPathParams struct {
	Period DateTime `schema:"period"`
}

func (p *PostSurveyUploadCSVPathParams) Params() map[string]string {
	return map[string]string{
		"period": p.Period.Format("200601"),
	}
}

func (r *PostSurveyUploadCSVRequest) PathParams() *PostSurveyUploadCSVPathParams {
	return r.pathParams
}

func (r *PostSurveyUploadCSVRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PostSurveyUploadCSVRequest) SetMethod(method string) {
	r.method = method
}

func (r *PostSurveyUploadCSVRequest) Method() string {
	return r.method
}

func (s *Client) NewPostSurveyUploadCSVRequestBody() PostSurveyUploadCSVRequestBody {
	return PostSurveyUploadCSVRequestBody{}
}

type PostSurveyUploadCSVRequestBody struct {
	RawCSV string `schema:"-"`
}

func (r *PostSurveyUploadCSVRequest) RequestBody() *PostSurveyUploadCSVRequestBody {
	return &r.requestBody
}

func (r *PostSurveyUploadCSVRequest) RequestBodyInterface() interface{} {
	return strings.NewReader(r.requestBody.RawCSV)
}

func (r *PostSurveyUploadCSVRequest) SetRequestBody(body PostSurveyUploadCSVRequestBody) {
	r.requestBody = body
}

func (r *PostSurveyUploadCSVRequest) NewResponseBody() *PostSurveyUploadCSVResponseBody {
	return &PostSurveyUploadCSVResponseBody{}
}

type PostSurveyUploadCSVResponseBody SurveyUploadCVSResponse

func (r *PostSurveyUploadCSVRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("survey-upload-csv/{{.period}}", r.PathParams())
	return &u
}

func (r *PostSurveyUploadCSVRequest) Do() (PostSurveyUploadCSVResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Set("Content-Type", "text/plain")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Reponse is string, so we unmarshal it to json
	var rawJSON string
	_, err = r.client.Do(req, &rawJSON)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	if err := json.Unmarshal([]byte(rawJSON), responseBody); err != nil {
		return *r.NewResponseBody(), err
	}

	return *responseBody, nil
}
