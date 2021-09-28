package f3client

import (
	"github.com/google/uuid"
)

type requestBody interface {
	validate() error
	convertToJson() ([]byte, int64, error)
}

type Data struct {
	RequestData RequestData `json:"data"`
}

type RequestData struct {
	ObjType    string      `json:"type,omitempty"`
	Id         uuid.UUID   `json:"id,omitempty"`
	Version    int         `json:"version,omitempty"`
	OrgId      uuid.UUID   `json:"organisation_id,omitempty"`
	CreatedOn  string      `json:"created_on,omitempty"`
	ModifiedOn string      `json:"modified_on,omitempty"`
	Attributes interface{} `json:"attributes,omitempty"`
}
