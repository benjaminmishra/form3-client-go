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

// f3client struct constants
const (
	UserAgent string = "form3-client-go/1.0.0"
	Accepts   string = "application/vnd.api+json"
)

// Http methods
const (
	Get    string = "GET"
	Put    string = "PUT"
	Post   string = "POST"
	Delete string = "DELETE"
)

// The basic form3 client object that
// needs to be instansiated for every request being sent to form3 apis
type Client struct {
	BaseURL    url.URL
	common     service
	HttpClient *http.Client
	UserAgent  string
	Accepts    string

	// Services for interacting with different parts of the API
	Accounts *AccountService
}

type service struct {
	client *Client
}

type Option func(*Client) error

// NewClient creates an f3client.Client instance
//
// If there is an error, non-nil error is returned and
// the f3Cleint.Client is set to nil
//
// Accepts a variable number of f3client.Options functions
// that configure the f3Client.Client instance as per their inputs
// example f3Cleint.WithHostUrl , f3client.WithHttpClient
func NewClient(options ...Option) (*Client, error) {

	defaultBaseUrl, err := url.Parse("http://localhost:8080")
	if err != nil {
		return nil, err
	}
	defaultHttpClient := http.DefaultClient

	c := &Client{
		BaseURL:    *defaultBaseUrl,
		HttpClient: defaultHttpClient,
		UserAgent:  UserAgent,
		Accepts:    Accepts,
	}

	for _, option := range options {
		err := option(c)
		if err != nil {
			return c, err
		}
	}

	c.common.client = c

	c.Accounts = &AccountService{
		service:    c.common,
		ObjectType: "accounts",
	}

	return c, nil
}

// WithHostUrl configures f3client.Client to add the constom url being passed
func WithHostUrl(hostAddr string) Option {
	f := func(c *Client) error {
		baseUrl, err := url.Parse(hostAddr)
		c.BaseURL = *baseUrl
		return err
	}
	return f
}

// WithHttpClient confiures f3client.Client to override the default http.Clinet
// and assigns the constom http.Client object being passed
func WithHttpClient(customClient *http.Client) Option {
	f := func(c *Client) error {
		c.HttpClient = customClient
		return nil
	}
	return f
}

// NewRequest creates http.Request object and returns a pointer to it
//
// if the request creation is not suceessful then it returns error and
// the request object is returned as nil. It is a wrapper on http.NewRequest
func (c *Client) NewRequest(ctx context.Context, method, urlStr, objectType string, body interface{}) (*http.Request, error) {
	var u *url.URL
	var err error

	if urlStr != "" {
		u, err = c.BaseURL.Parse(urlStr)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, NewArgError("urlStr", "urlStr cannot be empty")
	}

	var httpReq *http.Request
	var encodedBody io.Reader = nil
	var contentLen int64 = 0

	if method != http.MethodGet && method != http.MethodOptions && method != http.MethodHead && method != http.MethodDelete {
		// only aplicable for requests that require a body i.e. put,post , patch
		var requestBody *[]byte
		if objectType != "" && body != nil && body != "" {
			requestBody, err = MarshalToRequestBody(body, objectType)

			if err != nil {
				return nil, err
			}
			encodedBody = bytes.NewReader((*requestBody))
			contentLen = int64(len((*requestBody)))

		}

		if objectType == "" {
			return nil, NewArgError("objectType", "objectType cannot be empty")
		}

		if body == "" || body == nil {
			return nil, NewArgError("body", "body cannot nil or empty for Post requests")
		}
	}

	httpReq, err = http.NewRequest(method, u.String(), encodedBody)
	if err != nil {
		return nil, err
	}

	httpReq.ContentLength = contentLen
	httpReq.Header.Add("Accepts", c.Accepts)
	httpReq.Header.Add("User-Agent", c.UserAgent)

	return httpReq, nil
}

// SendRequest executes the http request to the apis and returns their response
//
// An error is returned if there is any error in executing the request
//
// Arguments context object , pointer to http.Request object
//
// Returns f3cleint.Response struct, which contains the repose body
func (c *Client) SendRequest(ctx context.Context, request *http.Request) (*Response, error) {
	response := new(Response)
	errorResponse := new(struct {
		Code    string `json:"error_code,omitempty"`
		Message string `json:"error_message,omitempty"`
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
	} else if httpResp.StatusCode == 204 {
		return nil, nil
	}

	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
