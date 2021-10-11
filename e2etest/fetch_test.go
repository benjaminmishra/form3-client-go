package e2etest

import (
	"context"
	"fmt"
	"log"
	"testing"

	f3client "github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/google/uuid"
)

func TestAccountService_Fetch(t *testing.T) {

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
	client, err := f3client.NewClient()
	if err != nil {
		return nil, err
	}
	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		return nil, err
	}

	return account, nil
}