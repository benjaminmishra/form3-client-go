// The basic form3 client object that needs to be instanciated for every request being sent to form3 apis

package f3client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	Get    string = "GET"
	Put    string = "PUT"
	Post   string = "POST"
	Delete string = "DELETE"
)

type Client struct {
	BaseURL    url.URL
	common     service
	HttpClient *http.Client
	UserAgent  string
	Version    string

	// Services for interacting with different parts of the API
	Accounts *AccountService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse("http://localhost:8080")
	c := &Client{
		BaseURL:    *baseURL,
		HttpClient: httpClient,
		UserAgent:  "form3-go-client",
		Version:    "v1",
	}

	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)

	return c
}

// NewRequest creates http.Request object and returns a pointer to it
// if the request creation is not suceessful then it returns error and
// the request object is returned as nil.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, objectType string, body interface{}) (*http.Request, error) {

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var httpReq *http.Request
	var encodedBody io.Reader
	var contentLen int64

	switch method {

	case http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodDelete:
		encodedBody = nil
		contentLen = 0
	default:
		requestBody, err := MarshalToRequestBody(body, objectType)
		if err != nil {
			return nil, err
		}

		encodedBody = bytes.NewReader((*requestBody))
		contentLen = int64(len((*requestBody)))
		httpReq, err = http.NewRequest(method, u.String(), encodedBody)
		if err != nil {
			return nil, err
		}
	}

	httpReq.ContentLength = contentLen

	return httpReq, nil
}

func (c *Client) SendRequest(ctx context.Context, request *http.Request) (*Response, error) {
	response := new(Response)
	errorResponse := new(struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	})

	httpResp, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode >= 400 && httpResp.StatusCode <= 504 {
		err = json.Unmarshal(bodyBytes, errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(errorResponse.Message)
	}

	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
