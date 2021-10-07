package main

import (
	"context"
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

	client := f3client.NewClient(nil)

	err = client.Accounts.Delete(ctx, accountId, 1)
	if err != nil {
		log.Fatal(err)
	}
}
