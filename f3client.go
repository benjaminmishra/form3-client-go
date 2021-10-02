// The basic form3 client object that needs to be instanciated for every request being sent to form3 apis

package f3client

import (
	"net/http"
	"net/url"
)

const (
	Get    string = "GET"
	Put           = "PUT"
	Post          = "POST"
	Delete        = "DELETE"
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
	}

	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)

	return c
}

// NewRequest creates http.Request object and returns a pointer to it
// if the request creation is not suceessful then it returns error and
// the request object is returned as nil. The body of the request needs
// to be of type of pointer to RequestBody
/*func (c *Client) NewRequest(ctx context.Context, method, relativeUrl string, requestBody interface{}) (*RequestBody, error) {

	u, err := c.BaseURL.Parse(relativeUrl)
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
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil, err
		}
		encodedBody = bytes.NewBuffer(jsonBody)
		contentLen = int64(len(jsonBody))
	}

	httpReq, newReqErr := http.NewRequest(method, u.String(), encodedBody)
	if newReqErr != nil {
		return nil, newReqErr
	}

	httpReq.ContentLength = contentLen

	return httpReq, nil
}

func (c *Client) SendRequest(ctx context.Context, request *http.Request) (*SuccessResponse, error) {

	httpResp, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	sucessResp, errResp, err := convert(httpResp)
	if err != nil {
		return nil, err
	}
	if errResp != nil {
		return nil, errors.New(errResp.Message)
	}

	return sucessResp, nil
}
*/
