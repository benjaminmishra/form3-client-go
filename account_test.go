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
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

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

	err = client.Accounts.Create(context.Background(), actual)
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

func TestAccountService_Create_RequestValidationFailed_ID(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dummy response"))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	organisationUUID, _ := uuid.Parse("ee2fb143-6dfe-4787-b183-de8ddd4164d1")

	// create request that needs to be passed
	actual := &f3client.Account{
		OrganisationID: organisationUUID,
		Attributes: f3client.AccountAttributes{
			Country: "US",
			Name:    []string{"Jon Doe", "Jane Doe"},
		},
	}

	expected := &f3client.Account{
		OrganisationID: organisationUUID,
		Attributes: f3client.AccountAttributes{
			Country: "US",
			Name:    []string{"Jon Doe", "Jane Doe"},
		},
	}

	err = client.Accounts.Create(context.Background(), actual)
	if err != nil {
		assert.EqualError(t, err, "id is mandatory in the request body")
	}

	assert.Equal(t, expected, actual)
}

func TestAccountService_Create_RequestValidationFailed_OrganisationID(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dummy response"))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	accountid, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		assert.FailNow(t, "failed with error : "+err.Error())
	}

	// create request that needs to be passed
	actual := &f3client.Account{
		ID: accountid,
		Attributes: f3client.AccountAttributes{
			Country: "US",
			Name:    []string{"Jon Doe", "Jane Doe"},
		},
	}

	expected := &f3client.Account{
		ID: accountid,
		Attributes: f3client.AccountAttributes{
			Country: "US",
			Name:    []string{"Jon Doe", "Jane Doe"},
		},
	}

	err = client.Accounts.Create(context.Background(), actual)
	if err != nil {
		assert.EqualError(t, err, "organisation_id is mandatory in the request body")
	}

	assert.Equal(t, expected, actual)
}

func TestAccountService_Create_RequestValidationFailed_Country(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dummy response"))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	accountid, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		panic(err)
	}

	organisationid, err := uuid.Parse("ee2fb143-6dfe-4787-b183-de8ddd4164d1")
	if err != nil {
		panic(err)
	}

	// create request that needs to be passed
	actual := &f3client.Account{
		ID:             accountid,
		OrganisationID: organisationid,
		Attributes: f3client.AccountAttributes{
			Name: []string{"Jon Doe", "Jane Doe"},
		},
	}

	expected := &f3client.Account{
		ID:             accountid,
		OrganisationID: organisationid,
		Attributes: f3client.AccountAttributes{
			Name: []string{"Jon Doe", "Jane Doe"},
		},
	}

	err = client.Accounts.Create(context.Background(), actual)
	if err != nil {
		assert.IsType(t, &f3client.ArgumentError{}, err)
		assert.EqualError(t, err, "country : country is mandatory for account create request")
	}

	assert.Equal(t, expected, actual)
}

func TestAccountService_Create_RequestValidationFailed_Names(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Dummy response"))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	accountid, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		panic(err)
	}

	organisationid, err := uuid.Parse("ee2fb143-6dfe-4787-b183-de8ddd4164d1")
	if err != nil {
		panic(err)
	}

	// create request that needs to be passed
	actual := &f3client.Account{
		ID:             accountid,
		OrganisationID: organisationid,
		Attributes: f3client.AccountAttributes{
			Country: "IN",
		},
	}

	expected := &f3client.Account{
		ID:             accountid,
		OrganisationID: organisationid,
		Attributes: f3client.AccountAttributes{
			Country: "IN",
		},
	}

	err = client.Accounts.Create(context.Background(), actual)
	if err != nil {
		assert.IsType(t, &f3client.ArgumentError{}, err)
		assert.EqualError(t, err, "name : names are mandatory for account create request")
	}

	assert.Equal(t, expected, actual)
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
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

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
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

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
		w.Write([]byte(`{
			"error_code": "Not Found",
			"error_message": "Account Not Found"
			}`))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")

	actual, err := client.Accounts.Fetch(context.Background(), resourceUUID)
	if err != nil {
		assert.EqualError(t, err, "Account Not Found")
	}

	expected := new(f3client.Account)

	assert.Equal(t, expected, actual)
}

func TestAccountService_Delete_NotFound(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{
			"error_code": "Not Found",
			"error_message": "Account Not Found"
			}`))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")

	actual, err := client.Accounts.Delete(context.Background(), resourceUUID, 0)
	if err != nil {
		assert.EqualError(t, err, "Account Not Found")
	}

	expected := false

	assert.Equal(t, expected, actual)
}

func TestAccountService_Version_Conflict(t *testing.T) {
	// mock the server and json response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{
			"error_code": "Conflict",
			"error_message": "Specified version incorrect"
			}`))
	}))

	// close the server once this test is done executing
	defer server.Close()

	// client is the Form 3 API client being tested and is
	// taking the mock server url.
	client, err := f3client.NewClient(f3client.WithHostUrl(server.URL))
	if err != nil {
		panic(err)
	}

	// test code
	resourceUUID, _ := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")

	actual, err := client.Accounts.Delete(context.Background(), resourceUUID, 1)
	if err != nil {
		assert.EqualError(t, err, "Specified version incorrect")
	}

	expected := false

	assert.Equal(t, expected, actual)
}
