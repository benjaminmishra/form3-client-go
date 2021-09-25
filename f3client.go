// The basic form3 client object that needs to be instanciated for every request being sent to form3 apis

package f3client

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    url.URL
	common     service
	HttpClient *http.Client
	UserAgent  string

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

	baseURL, _ := url.Parse("http://api.form3.tech")
	c := &Client{
		BaseURL:    *baseURL,
		HttpClient: httpClient,
		UserAgent:  "form3-go-client",
	}

	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)

	return c
}
