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
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	account, err := fetchAccount(ctx, accountId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(account)
}

func fetchAccount(ctx context.Context, accountId uuid.UUID) (*f3client.Account, error) {
	client := f3client.NewClient(nil)
	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}
