package f3client

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

// Response is Form3 standard resposne object which warps the service specific attributes and
// other common fields sent back in every api response.
type Response struct {
	*http.Response

	Data  responseData `json:"data,omitempty"`
	Links links        `json:"links,omitempty"`
}

type responseData struct {
	Type           string      `json:"type,omitempty"`
	ID             uuid.UUID   `json:"id,omitempty"`
	Version        int         `json:"version,omitempty"`
	OrganisationID uuid.UUID   `json:"organisation_id,omitempty"`
	CreatedOn      string      `json:"created_on,omitempty"`
	ModifiedOn     string      `json:"modified_on,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type links struct {
	Self  url.URL `json:"self,omitempty"`
	First url.URL `json:"first,omitempty"`
	Last  url.URL `json:"last,omitempty"`
	Next  url.URL `json:"next,omitempty"`
	Prev  url.URL `json:"prev,omitempty"`
}

type ErrorResponse struct {
	*http.Response
	Code    string `json:"error_code,omitempty"`
	Message string `json:"error_message,omitempty"`
}

func convert(httpResponse *http.Response, *ErrorResponse, error) {
	
	var e ErrorResponse
	var i map[string]interface{}

	err := json.NewDecoder(httpResponse.Body).Decode(i)
	if err != nil {
		return nil, nil, err
	}

	if httpResponse.StatusCode >= 299 {
		for key, value := range i {
			switch key {
			case "data":
				e.Code = value.(string)
			case "error_message":
				e.Message = value.(string)
			default:
				err = errors.New("error response conversion error")
			}
		}
	} else {
		for key, value := range i {
			switch key {
			case "data":
				s.Data = value.(responseData)
			case "links":
				s.Links = value.(links)
			default:
				err = errors.New("sucess response conversion error")
			}
		}
	}
	return &s, &e, err
}
