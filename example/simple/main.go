package main

import (
	"context"
	"fmt"
	"log"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
)

func main() {
	client := f3client.NewClient(nil)
	ctx := context.Background()

	accountId, err := uuid.Parse("7826c3cb-d6fd-41d0-b187-dc23ba928772")
	if err != nil {
		log.Fatal(err)
	}

	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(account)
}
