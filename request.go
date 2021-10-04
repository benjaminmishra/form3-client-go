package f3client

import (
	"encoding/json"
	"fmt"
)

func ConvertToRequestBody(req interface{}, requestType string) (*[]byte, error) {

	inInteface := make(map[string]interface{})
	inReqBody := make(map[string]interface{})

	// validate arguments to the function
	if req == nil {
		return nil, fmt.Errorf("req cannot be nil")
	}

	if requestType == "" {
		return nil, fmt.Errorf("requestType cannot be empty")
	}

	byteReq, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteReq, &inInteface)
	if err != nil {
		return nil, err
	}

	// validate the incoming body
	err = validateReq(&inInteface)
	if err != nil {
		return nil, err
	}

	// add generic RequestBody attributes to the inInterface
	inInteface["type"] = requestType
	inInteface["version"] = 0

	// wrap the whole thing in data
	inReqBody["data"] = inInteface

	// convert to RequestBody type
	encodedRequestBody, err := json.Marshal(&inReqBody)
	if err != nil {
		return nil, err
	}

	return &encodedRequestBody, nil
}

func validateReq(inInteface *map[string]interface{}) error {
	// TODO : move to a more generic place, to make this dependecy more explicit
	mandatoryFields := [2]string{"id", "organisation_id"}

	// check for missing mandatory fields
	for _, field := range mandatoryFields {
		// yay found! check if the field values are 0
		if value, ok := (*inInteface)[field]; ok {
			if value == "00000000-0000-0000-0000-000000000000" {
				return fmt.Errorf("%s is mandatory in the request body", field)
			}
		} else {
			// not found! throw error
			return fmt.Errorf("%s is mandatory in the request body", field)
		}
	}

	return nil
}
