package f3client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type RequestBody struct {
	Data requestData
}

type requestData struct {
	Type           string      `json:"type,omitempty"`
	ID             uuid.UUID   `json:"id,omitempty"`
	Version        int         `json:"version,omitempty"`
	OrganisationID uuid.UUID   `json:"organisation_id,omitempty"`
	CreatedOn      string      `json:"created_on,omitempty"`
	ModifiedOn     string      `json:"modified_on,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

// NewRequest creates http.Request object and returns a pointer to it
// if the request creation is not suceessful then it returns error and
// the request object is returned as nil. The body of the request needs
// to be of type of pointer to RequestBody
func (c *Client) NewRequest(method, url string, body *RequestBody) (*http.Request, error) {

	u, err := c.BaseURL.Parse(url)

	if err != nil {
		return nil, err
	}

	var httpReq *http.Request

	switch method {
	case http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodDelete:
		httpReq, err = http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil {
			return nil, err
		}

	default:
		encodedBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		httpReq, err = http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(encodedBody))
		if err != nil {
			return nil, err
		}
	}

	return httpReq, nil
}
