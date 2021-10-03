package f3client_test

import (
	"encoding/json"
	"testing"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConvertToRequestBody(t *testing.T) {
	accUUId := uuid.New()
	orgUUId := uuid.New()

	expected := new(f3client.RequestBody)

	accCreateReq := f3client.AccountCreateRequest{
		ID:             accUUId,
		OrganisationID: orgUUId,
		Attributes: struct {
			Country                 string   `json:"country,omitempty"`
			BaseCurrency            string   `json:"base_currency,omitempty"`
			BankID                  string   `json:"bank_id,omitempty"`
			BankIDCode              string   `json:"bank_id_code,omitempty"`
			Bic                     string   `json:"bic,omitempty"`
			Iban                    string   `json:"iban,omitempty"`
			CustomerID              string   `json:"customer_id,omitempty"`
			Name                    []string `json:"name,omitempty"`
			AlternativeNames        []string `json:"alternative_names,omitempty"`
			AccountClassification   string   `json:"account_classification,omitempty"`
			JointAccount            bool     `json:"joint_account,omitempty"`
			AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
			SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
			Switched                bool     `json:"switched,omitempty"`
			ProcessingService       string   `json:"processing_service,omitempty"`
			UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
			ValidationType          string   `json:"validation_type,omitempty"`
			ReferenceMask           string   `json:"reference_mask,omitempty"`
			AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
		}{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	reqBody := &f3client.RequestBody{
		Data: struct {
			ObjType    string      `json:"type"`
			Id         uuid.UUID   `json:"id"`
			Version    int         `json:"version"`
			OrgId      uuid.UUID   `json:"organisation_id"`
			Attributes interface{} `json:"attributes"`
		}{
			ObjType: "accounts",
			Id:      accUUId,
			Version: 0,
			OrgId:   orgUUId,
			Attributes: struct {
				Country                 string   `json:"country,omitempty"`
				BaseCurrency            string   `json:"base_currency,omitempty"`
				BankID                  string   `json:"bank_id,omitempty"`
				BankIDCode              string   `json:"bank_id_code,omitempty"`
				Bic                     string   `json:"bic,omitempty"`
				Iban                    string   `json:"iban,omitempty"`
				CustomerID              string   `json:"customer_id,omitempty"`
				Name                    []string `json:"name,omitempty"`
				AlternativeNames        []string `json:"alternative_names,omitempty"`
				AccountClassification   string   `json:"account_classification,omitempty"`
				JointAccount            bool     `json:"joint_account,omitempty"`
				AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
				SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
				Switched                bool     `json:"switched,omitempty"`
				ProcessingService       string   `json:"processing_service,omitempty"`
				UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
				ValidationType          string   `json:"validation_type,omitempty"`
				ReferenceMask           string   `json:"reference_mask,omitempty"`
				AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
			}{
				Country:           "GB",
				BaseCurrency:      "GBP",
				BankID:            "400300",
				BankIDCode:        "GBDSC",
				Bic:               "NWBKGB22",
				ProcessingService: "ABC Bank",
			},
		},
	}

	// simulate json conversion
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	err = json.Unmarshal(jsonReqBody, expected)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	actual, err := f3client.ConvertToRequestBody(accCreateReq, "accounts")
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, actual, expected)
}

func TestConvertToRequestBody_NilRequest(t *testing.T) {

	_, err := f3client.ConvertToRequestBody(nil, "accounts")
	if err != nil {
		assert.EqualError(t, err, "req cannot be nil")
	} else {
		assert.FailNow(t, "Coversion doesn't check for nil req")
	}
}

func TestConvertToRequestBody_NilRequestType(t *testing.T) {
	accUUId := uuid.New()
	orgUUId := uuid.New()

	accCreateReq := f3client.AccountCreateRequest{
		ID:             accUUId,
		OrganisationID: orgUUId,
		Attributes: struct {
			Country                 string   `json:"country,omitempty"`
			BaseCurrency            string   `json:"base_currency,omitempty"`
			BankID                  string   `json:"bank_id,omitempty"`
			BankIDCode              string   `json:"bank_id_code,omitempty"`
			Bic                     string   `json:"bic,omitempty"`
			Iban                    string   `json:"iban,omitempty"`
			CustomerID              string   `json:"customer_id,omitempty"`
			Name                    []string `json:"name,omitempty"`
			AlternativeNames        []string `json:"alternative_names,omitempty"`
			AccountClassification   string   `json:"account_classification,omitempty"`
			JointAccount            bool     `json:"joint_account,omitempty"`
			AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
			SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
			Switched                bool     `json:"switched,omitempty"`
			ProcessingService       string   `json:"processing_service,omitempty"`
			UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
			ValidationType          string   `json:"validation_type,omitempty"`
			ReferenceMask           string   `json:"reference_mask,omitempty"`
			AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
		}{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.ConvertToRequestBody(accCreateReq, "")
	if err != nil {
		assert.EqualError(t, err, "requestType cannot be empty")
	} else {
		assert.FailNow(t, "Coversion doesn't check for nil requestType")
	}
}

func TestConvertToRequestBody_OrganisationIDMissing(t *testing.T) {
	accUUId := uuid.New()
	accCreateReq := f3client.AccountCreateRequest{
		ID: accUUId,
		Attributes: struct {
			Country                 string   `json:"country,omitempty"`
			BaseCurrency            string   `json:"base_currency,omitempty"`
			BankID                  string   `json:"bank_id,omitempty"`
			BankIDCode              string   `json:"bank_id_code,omitempty"`
			Bic                     string   `json:"bic,omitempty"`
			Iban                    string   `json:"iban,omitempty"`
			CustomerID              string   `json:"customer_id,omitempty"`
			Name                    []string `json:"name,omitempty"`
			AlternativeNames        []string `json:"alternative_names,omitempty"`
			AccountClassification   string   `json:"account_classification,omitempty"`
			JointAccount            bool     `json:"joint_account,omitempty"`
			AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
			SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
			Switched                bool     `json:"switched,omitempty"`
			ProcessingService       string   `json:"processing_service,omitempty"`
			UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
			ValidationType          string   `json:"validation_type,omitempty"`
			ReferenceMask           string   `json:"reference_mask,omitempty"`
			AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
		}{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.ConvertToRequestBody(accCreateReq, "accounts")

	assert.EqualError(t, err, "organisation_id is mandatory in the request body")
}

func TestConvertToRequestBody_IDMissing(t *testing.T) {
	orgUUId := uuid.New()
	accCreateReq := f3client.AccountCreateRequest{
		OrganisationID: orgUUId,
		Attributes: struct {
			Country                 string   `json:"country,omitempty"`
			BaseCurrency            string   `json:"base_currency,omitempty"`
			BankID                  string   `json:"bank_id,omitempty"`
			BankIDCode              string   `json:"bank_id_code,omitempty"`
			Bic                     string   `json:"bic,omitempty"`
			Iban                    string   `json:"iban,omitempty"`
			CustomerID              string   `json:"customer_id,omitempty"`
			Name                    []string `json:"name,omitempty"`
			AlternativeNames        []string `json:"alternative_names,omitempty"`
			AccountClassification   string   `json:"account_classification,omitempty"`
			JointAccount            bool     `json:"joint_account,omitempty"`
			AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
			SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
			Switched                bool     `json:"switched,omitempty"`
			ProcessingService       string   `json:"processing_service,omitempty"`
			UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
			ValidationType          string   `json:"validation_type,omitempty"`
			ReferenceMask           string   `json:"reference_mask,omitempty"`
			AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
		}{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
		},
	}

	_, err := f3client.ConvertToRequestBody(accCreateReq, "accounts")

	assert.EqualError(t, err, "id is mandatory in the request body")
}
