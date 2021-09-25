package f3client

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

// Response is Form3 standard resposne object which warps the service specific attributes and
// other common fields sent back in every api response.
type Response struct {
	*http.Response

	Data  []responseData `json:"data"`
	Links links          `json:"links,omitempty"`
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
