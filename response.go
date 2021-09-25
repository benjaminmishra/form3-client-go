package f3client

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
)

type Response struct {
	*http.Response

	Data  []Data
	Links Links
}

type Data struct {
	Type           string      `json:"type,omitempty"`
	ID             uuid.UUID   `json:"id,omitempty"`
	Version        int         `json:"version,omitempty"`
	OrganisationID uuid.UUID   `json:"organisation_id,omitempty"`
	CreatedOn      string      `json:"created_on,omitempty"`
	ModifiedOn     string      `json:"modified_on,omitempty"`
	Attributes     interface{} `json:"attributes,omitempty"`
}

type Links struct {
	Self  url.URL `json:"self,omitempty"`
	First url.URL `json:"first,omitempty"`
	Last  url.URL `json:"last,omitempty"`
	Next  url.URL `json:"next,omitempty"`
	Prev  url.URL `json:"prev,omitempty"`
}
