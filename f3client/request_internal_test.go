package f3client

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestValidateRequest_ValidationIdMissing(t *testing.T) {

	// prepare
	requestMap := make(map[string]interface{})

	requestMap["x"] = "x"
	requestMap["y"] = "y"
	requestMap["organisation_id"] = uuid.New().String()

	// execute
	err := validateReq(&requestMap)

	// assert
	if err != nil {
		assert.EqualError(t, err, "id is mandatory in the request body")
	} else {
		assert.FailNow(t, "Validation did not work")
	}
}

func TestValidateRequest_ValidationOrgIdMissing(t *testing.T) {

	// prepare
	requestMap := make(map[string]interface{})

	requestMap["x"] = "x"
	requestMap["y"] = "y"
	requestMap["id"] = uuid.New().String()

	// execute
	err := validateReq(&requestMap)

	// assert
	if err != nil {
		assert.EqualError(t, err, "organisation_id is mandatory in the request body")
	} else {
		assert.FailNow(t, "Validation did not work")
	}
}

func TestValidateRequest_ValidationForIdZero(t *testing.T) {

	// prepare
	requestMap := make(map[string]interface{})

	requestMap["x"] = "x"
	requestMap["y"] = "y"
	requestMap["id"] = "00000000-0000-0000-0000-000000000000"
	requestMap["organisation_id"] = uuid.New().String()

	// execute
	err := validateReq(&requestMap)

	// assert
	if err != nil {
		assert.EqualError(t, err, "id is mandatory in the request body")
	} else {
		assert.FailNow(t, "Validation did not work")
	}

}

func TestValidateRequest_ValidationForOrgIdZero(t *testing.T) {

	// prepare
	requestMap := make(map[string]interface{})

	requestMap["x"] = "x"
	requestMap["y"] = "y"
	requestMap["id"] = uuid.New().String()
	requestMap["organisation_id"] = "00000000-0000-0000-0000-000000000000"

	// execute
	err := validateReq(&requestMap)

	// assert
	if err != nil {
		assert.EqualError(t, err, "organisation_id is mandatory in the request body")
	} else {
		assert.FailNow(t, "Validation did not work")
	}

}
