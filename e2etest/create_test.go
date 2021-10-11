package e2etest

import (
	"context"
	"log"
	"testing"
	"time"

	f3client "github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_Create(t *testing.T) {

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient(f3client.WithHostUrl("http://localhost:8080"))
	if err != nil {
		panic(err)
	}

	accountId, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	orgId, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	// create new account object
	newAccount := f3client.Account{
		ID:             accountId,
		OrganisationID: orgId,
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			ProcessingService: "ABC Bank",
			Name:              []string{"hello world"},
			AlternativeNames:  []string{"abc"},
		},
	}

	// call the create account api
	err = client.Accounts.Create(ctx, &newAccount)
	if err != nil {
		log.Fatal(err)
	}

	expected := f3client.Account{
		ID:             accountId,
		Version:        0,
		OrganisationID: orgId,
		CreatedOn:      time.Now().String(),
		ModifiedOn:     time.Now().String(),
		Attributes: f3client.AccountAttributes{
			Country:           "GB",
			BaseCurrency:      "GBP",
			BankID:            "400300",
			BankIDCode:        "GBDSC",
			Bic:               "NWBKGB22",
			Name:              []string{"hello world"},
			AlternativeNames:  []string{"abc"},
			Switched:          false,
			ProcessingService: "ABC Bank",
		},
	}

	assert.Equal(t, expected.ID, newAccount.ID)
	assert.Equal(t, expected.OrganisationID, newAccount.OrganisationID)
	assert.Equal(t, expected.Version, newAccount.Version)
	assert.Equal(t, expected.Attributes.Country, newAccount.Attributes.Name)
	assert.Equal(t, expected.Attributes.Status, newAccount.Attributes.Status)

}

func TestAccountService_Create_EmptyBody(t *testing.T) {

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient(f3client.WithHostUrl("http://localhost:8080"))
	if err != nil {
		panic(err)
	}

	_, err = uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	_, err = uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}

	// create new account object
	newAccount := f3client.Account{}

	// call the create account api
	err = client.Accounts.Create(ctx, &newAccount)

	if assert.Error(t, err) {
		assert.ErrorIs(t, err, &f3client.ArgumentError{})
	}
}
