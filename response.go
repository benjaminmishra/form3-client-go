package f3client

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// Response is Form3 standard response object which warps the service specific attributes and
// other common fields sent back in every api response.
type Response struct {
	Data  ResponseData `json:"data,omitempty"`
	Links Links        `json:"links,omitempty"`
}

type ResponseData struct {
	Type           string      `json:"type,omitempty"`
	ID             uuid.UUID   `json:"id,omitempty"`
	Version        int         `json:"version,omitempty"`
	OrganisationID uuid.UUID   `json:"organisation_id,omitempty"`
	CreatedOn      string      `json:"created_on,omitempty"`
	ModifiedOn     string      `json:"modified_on,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type Links struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}

// ConvertsTo converts the response object's data field to the type being passed in the argument.
// pre requisite is the response object data field would have to be of same fields as the target type
func (r *Response) ConvertTo(targetType interface{}) error {
	// if the targettype is nil, then do nothing
	if r.Data.Attributes != nil && targetType != nil {
		encoded, err := json.Marshal(r.Data)
		if err != nil {
			return err
		}

		err = json.Unmarshal(encoded, targetType)
		if err != nil {
			return err
		}

		return nil
	} else {
		return errors.New("targetType cannot be nil")
	}
}
