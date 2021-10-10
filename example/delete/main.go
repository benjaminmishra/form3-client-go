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

	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}

	success, err := client.Accounts.Delete(ctx, accountId, 1)
	if err != nil {
		log.Fatal(err)
	}

	if success {
		fmt.Print("Account has been deleted")
	} else {
		fmt.Print("Account was not deleted")
	}
}
