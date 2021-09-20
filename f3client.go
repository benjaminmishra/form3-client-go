// The basic form3 client object that needs to be instanciated for every request being sent to form3 apis

package f3client

import (
	"net/http"
)

type Client struct {
	Host       string
	common     service
	HttpClient *http.Client
	UserAgent  string

	// Services for interacting with different parts of the API
	Account    *AccountService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	c := &Client{
		Host:       "api.form3.tech",
		HttpClient: httpClient,
		UserAgent:  "form3-go-client",
	}

	c.common.client = c
	c.Account = (*AccountService)(&c.common)

	return c
}
