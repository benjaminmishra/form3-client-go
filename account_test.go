package f3client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_CreateAccount(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
				{
					"data": {
						"attributes": {
							"account_classification": "Personal",
							"account_matching_opt_out": false,
							"account_number": "41426819",
							"alternative_names": [
								"mollit elit",
								"enim mollit"
							],
							"bank_id": "400300",
							"bank_id_code": "GBDSC",
							"base_currency": "GBP",
							"bic": "NWBKGB22",
							"country": "GB",
							"iban": "GB11NWBK40030041426819",
							"joint_account": false,
							"name": [
								"Jon Doe",
								"Jane Doe"
							],
							"secondary_identification": "Duis consectetur proident anim",
							"status": "pending",
							"switched": false
						},
						"created_on": "2021-10-03T13:44:27.809Z",
						"id": "bc8fb900-d6fd-41d0-b187-dc23ba928712",
						"modified_on": "2021-10-03T13:44:27.809Z",
						"organisation_id": "ee2fb143-6dfe-4787-b183-de8ddd4164d1",
						"type": "accounts",
						"version": 0
					},
					"links": {
						"self": "/v1/organisation/accounts/bc8fb900-d6fd-41d0-b187-dc23ba928712"
					}
				}
			  `))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client := f3client.NewClient(nil, server.URL)

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	organisationUUID, _ := uuid.Parse("ee2fb143-6dfe-4787-b183-de8ddd4164d1")

	// create request that needs to be passed
	actual := &f3client.Account{
		ID:             resourceUUID,
		OrganisationID: organisationUUID,
		Attributes: f3client.AccountAttributes{
			Country: "US",
			Name:    []string{"Jon Doe", "Jane Doe"},
		},
	}

	err := client.Accounts.Create(context.Background(), actual)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	expected := f3client.Account{
		ID:             resourceUUID,
		OrganisationID: organisationUUID,
		Version:        0,
		CreatedOn:      "2021-10-03T13:44:27.809Z",
		ModifiedOn:     "2021-10-03T13:44:27.809Z",
		Attributes: f3client.AccountAttributes{
			Country:                 "GB",
			BaseCurrency:            "GBP",
			AccountNumber:           "41426819",
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			Bic:                     "NWBKGB22",
			Iban:                    "GB11NWBK40030041426819",
			Status:                  "pending",
			AccountMatchingOptOut:   false,
			AccountClassification:   "Personal",
			AlternativeNames:        []string{"mollit elit", "enim mollit"},
			JointAccount:            false,
			Name:                    []string{"Jon Doe", "Jane Doe"},
			SecondaryIdentification: "Duis consectetur proident anim",
			Switched:                false,
		},
	}

	assert.Equal(t, expected, *actual)
}

func TestAccountService_FetchAccount(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
				{
					"data": {
						"attributes": {
							"account_classification": "Personal",
							"account_matching_opt_out": false,
							"account_number": "41426819",
							"alternative_names": [
								"mollit elit",
								"enim mollit"
							],
							"bank_id": "400300",
							"bank_id_code": "GBDSC",
							"base_currency": "GBP",
							"bic": "NWBKGB22",
							"country": "IN",
							"iban": "GB11NWBK40030041426819",
							"joint_account": false,
							"name": [
								"esse",
								"exercitation"
							],
							"secondary_identification": "Duis consectetur proident anim",
							"status": "pending",
							"switched": false
						},
						"created_on": "2021-10-03T13:44:27.809Z",
						"id": "bc8fb900-d6fd-41d0-b187-dc23ba928712",
						"modified_on": "2021-10-03T13:44:27.809Z",
						"organisation_id": "ee2fb143-6dfe-4787-b183-de8ddd4164d1",
						"type": "accounts",
						"version": 0
					},
					"links": {
						"self": "/v1/organisation/accounts/bc8fb900-d6fd-41d0-b187-dc23ba928712"
					}
				}
			  `))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client := f3client.NewClient(nil, server.URL)

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	organisationUUID, _ := uuid.Parse("ee2fb143-6dfe-4787-b183-de8ddd4164d1")

	actual, err := client.Accounts.Fetch(context.Background(), resourceUUID)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	expected := f3client.Account{
		ID:             resourceUUID,
		OrganisationID: organisationUUID,
		Version:        0,
		CreatedOn:      "2021-10-03T13:44:27.809Z",
		ModifiedOn:     "2021-10-03T13:44:27.809Z",
		Attributes: f3client.AccountAttributes{
			Country:                 "IN",
			BaseCurrency:            "GBP",
			AccountNumber:           "41426819",
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			Bic:                     "NWBKGB22",
			Iban:                    "GB11NWBK40030041426819",
			Status:                  "pending",
			AccountMatchingOptOut:   false,
			AccountClassification:   "Personal",
			AlternativeNames:        []string{"mollit elit", "enim mollit"},
			JointAccount:            false,
			Name:                    []string{"esse", "exercitation"},
			SecondaryIdentification: "Duis consectetur proident anim",
			Switched:                false,
		},
	}

	assert.Equal(t, expected, *actual)
}

func TestAccountService_DeleteAccount_NoContent(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client := f3client.NewClient(nil, server.URL)

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")

	actual, err := client.Accounts.Delete(context.Background(), resourceUUID, 0)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	expected := true

	assert.Equal(t, expected, actual)
}

func TestAccountService_Fetch_NotFound(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client := f3client.NewClient(nil, server.URL)

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")

	actual, err := client.Accounts.Fetch(context.Background(), resourceUUID)
	if err != nil {
		assert.EqualError(t, err, "Not found")
	}

	assert.Equal(t, *new(f3client.Account), actual)
}
