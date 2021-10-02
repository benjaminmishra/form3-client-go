package f3client

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Request struct {
}

type RequestBody struct {
	Data struct {
		ObjType    string      `json:"type"`
		Id         uuid.UUID   `json:"id"`
		Version    int         `json:"version"`
		OrgId      uuid.UUID   `json:"organisation_id"`
		Attributes interface{} `json:"attributes"`
	} `json:"data"`
}

func ConvertToRequestBody(req interface{}, requestType string) (*RequestBody, error) {

	mandatoryFields := [2]string{"id", "organisation_id"}

	var requestBody RequestBody
	inInteface := make(map[string]interface{})
	inReqBody := make(map[string]interface{})

	byteReq, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteReq, &inInteface)
	if err != nil {
		return nil, err
	}

	// validate the request body
	// check for missing mandatory fields
	for _, field := range mandatoryFields {
		if value, ok := inInteface[field]; ok {
			if value == "00000000-0000-0000-0000-000000000000" {
				return nil, fmt.Errorf("%s is mandatory in the request body", field)
			}
		} else {
			return nil, fmt.Errorf("%s is mandatory in the request body", field)
		}
	}

	// add generic RequestBody attributes to the inInterface
	inInteface["type"] = requestType
	inInteface["version"] = 0

	// wrap the whole thing in data , same as ReqBody type
	inReqBody["data"] = inInteface

	// convert to RequestBody type
	jsonStr, err := json.Marshal(&inReqBody)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonStr, &requestBody)
	if err != nil {
		return nil, err
	}

	return &requestBody, nil
}
