package f3client_test

import (
	"encoding/json"
	"testing"
	"time"

	f3client "github.com/benjaminmishra/form3-client-go/v1.0.0-beta.2/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Unit_ConvertTo(t *testing.T) {

	accId := uuid.New()
	orgId := uuid.New()
	date := time.Now().String()

	response := f3client.Response{
		Data: f3client.ResponseData{
			Type:           "account",
			ID:             accId,
			Version:        1,
			OrganisationID: orgId,
			CreatedOn:      date,
			ModifiedOn:     date,
			Attributes: f3client.AccountAttributes{
				BankID:       "XYZ1234",
				BaseCurrency: "INR",
				BankIDCode:   "1234ASD",
				Bic:          "ASD",
			},
		},
		Links: f3client.Links{
			Self: "http://localhost:8080/v1/accounts/" + accId.String(),
		},
	}

	expected := f3client.Account{
		ID:             accId,
		Version:        1,
		OrganisationID: orgId,
		CreatedOn:      date,
		ModifiedOn:     date,
		Attributes: f3client.AccountAttributes{
			BankID:       "XYZ1234",
			BaseCurrency: "INR",
			BankIDCode:   "1234ASD",
			Bic:          "ASD",
		},
	}
	acc := new(f3client.Account)
	err := response.ConvertTo(acc)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.EqualValues(t, acc, &expected)

}

func Test_Unit_ConvertTo_NilTargetType(t *testing.T) {

	accId := uuid.New()
	orgId := uuid.New()
	date := time.Now().String()

	response := f3client.Response{
		Data: f3client.ResponseData{
			Type:           "account",
			ID:             accId,
			Version:        1,
			OrganisationID: orgId,
			CreatedOn:      date,
			ModifiedOn:     date,
			Attributes: f3client.AccountAttributes{
				BankID:       "XYZ1234",
				BaseCurrency: "INR",
				BankIDCode:   "1234ASD",
				Bic:          "ASD",
			},
		},
		Links: f3client.Links{
			Self: "http://localhost:8080/v1/accounts/" + accId.String(),
		},
	}

	err := response.ConvertTo(nil)

	assert.EqualError(t, err, "targetType cannot be nil")

}

func Test_Unit_ConvertTo_IncompatableTargetType(t *testing.T) {

	accId := uuid.New()
	orgId := uuid.New()
	date := time.Now().String()

	response := f3client.Response{
		Data: f3client.ResponseData{
			Type:           "account",
			ID:             accId,
			Version:        1,
			OrganisationID: orgId,
			CreatedOn:      date,
			ModifiedOn:     date,
			Attributes: f3client.AccountAttributes{
				BankID:       "XYZ1234",
				BaseCurrency: "INR",
				BankIDCode:   "1234ASD",
				Bic:          "ASD",
			},
		},
		Links: f3client.Links{
			Self: "http://localhost:8080/v1/accounts/" + accId.String(),
		},
	}

	targetType := new(string)
	actualerr := response.ConvertTo(targetType)
	var expectederr *json.UnmarshalTypeError

	assert.ErrorAs(t, actualerr, &expectederr)

}
