package e2etest

import (
	"context"
	"fmt"
	"log"
	"testing"

	f3client "github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/google/uuid"
)

func TestAccountService_Create(t *testing.T) {

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}

	accountId, err := uuid.NewUUID()
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
		},
	}

	// call the create account api
	err = client.Accounts.Create(ctx, &newAccount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(newAccount)
}
