package f3client

import (
	"encoding/json"

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

func (r *Response) ConvertTo(targetType interface{}) error {
	encoded, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(encoded, targetType)
	if err != nil {
		return err
	}

	return nil
}
