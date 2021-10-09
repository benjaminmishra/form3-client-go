package main

import (
	"context"
	"fmt"
	"log"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
)

func main() {

	ctx := context.Background()
	// instantiate the client
	client := f3client.NewClient(nil, "http://localhost:8080")

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
