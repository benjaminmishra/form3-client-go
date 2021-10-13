package f3client_test

import (
	"encoding/json"
	"testing"

	f3client "github.com/benjaminmishra/form3-client-go/v1.0.0-alpha/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Unit_MarshalToRequestBody(t *testing.T) {
	accUUId := uuid.New()
	orgUUId := uuid.New()

	accCreateReq := f3client.Account{
		ID:             accUUId,
		OrganisationID: orgUUId,
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	reqBody := map[string]interface{}{
		"data": map[string]interface{}{
			"type":            "accounts",
			"id":              accUUId.String(),
			"version":         0,
			"organisation_id": orgUUId.String(),
			"attributes": map[string]interface{}{
				"country":            "GB",
				"base_currency":      "GBP",
				"bank_id":            "400300",
				"bank_id_code":       "GBDSC",
				"bic":                "NWBKGB22",
				"processing_service": "ABC Bank",
			},
		},
	}

	// simulate json conversion
	expected, err := json.Marshal(reqBody)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// execute the function
	actual, err := f3client.MarshalToRequestBody(accCreateReq, "accounts")
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, *actual, expected)
}

func Test_Unit_MarshalToRequestBody_NilRequest(t *testing.T) {

	_, err := f3client.MarshalToRequestBody(nil, "accounts")
	if err != nil {
		assert.IsType(t, &f3client.ArgumentError{}, err)
		assert.EqualError(t, err, "request : request cannot be nil")
	} else {
		assert.FailNow(t, "Coversion doesn't check for nil req")
	}
}

func Test_Unit_MarshalToRequestBody_NilRequestType(t *testing.T) {
	accUUId := uuid.New()
	orgUUId := uuid.New()

	accCreateReq := f3client.Account{
		ID:             accUUId,
		OrganisationID: orgUUId,
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.MarshalToRequestBody(accCreateReq, "")
	if err != nil {
		assert.IsType(t, &f3client.ArgumentError{}, err)
		assert.EqualError(t, err, "requestType : requestType cannot be empty")
	} else {
		assert.FailNow(t, "Coversion doesn't check for nil requestType")
	}
}

func Test_Unit_MarshalToRequestBody_OrganisationIDMissing(t *testing.T) {
	accUUId := uuid.New()
	accCreateReq := f3client.Account{
		ID: accUUId,
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.MarshalToRequestBody(accCreateReq, "accounts")

	assert.EqualError(t, err, "organisation_id is mandatory in the request body")
}

func Test_Unit_MarshalToRequestBody_IDMissing(t *testing.T) {
	orgUUId := uuid.New()
	accCreateReq := f3client.Account{
		OrganisationID: orgUUId,
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.MarshalToRequestBody(accCreateReq, "accounts")

	assert.EqualError(t, err, "id is mandatory in the request body")
}
