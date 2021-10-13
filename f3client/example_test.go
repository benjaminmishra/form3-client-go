package f3client_test

import (
	"context"
	"fmt"
	"log"

	"github.com/benjaminmishra/form3-client-go/v1/f3client"
	"github.com/google/uuid"
)

func ExampleAccountService_Create() {

	// create new f3client, with default options
	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	// create background context
	bctx := context.Background()

	// create an account object with f3client.AccountAttributes
	// ID , OrganisationID , Country, Name are mandatory while
	// doing a create request
	accountId, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	orgId, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	account := f3client.Account{
		ID:             accountId,
		OrganisationID: orgId,
		Attributes: f3client.AccountAttributes{
			Country: "GB",
			Name:    []string{"jane doe", "john doe"},
		},
	}

	err = c.Accounts.Create(bctx, &account)
	if err != nil {
		panic(err)
	}

	// use the account object for further operations
}

func ExampleAccountService_Delete() {
	// create context
	ctx := context.Background()

	// get or create account uuid
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	// create form3 client object
	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}

	// call the Accounts delete request with the account uuid and version no of the account
	// the delete function returns a bool indicating whether the request was successful or not
	success, err := client.Accounts.Delete(ctx, accountId, 1)
	if err != nil {
		log.Fatal(err)
	}

	// Use the returned bool to perform further operations
	if success {
		fmt.Print("Account has been deleted")
	} else {
		fmt.Print("Account was not deleted")
	}
}

func ExampleAccountService_Fetch() {
	// create context
	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	// create new f3client object
	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}

	// call the fetch function on the client.Account service
	// accountId is mandatory
	// this returns a account object and nil error
	// in case of error account is nil and err is non-nil
	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		panic(err)
	}

	// do next operations with the new account object
	fmt.Print(account)
}
