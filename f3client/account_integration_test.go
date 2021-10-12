package f3client_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	f3client "github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var (
	host     string = os.Getenv("PSQL_HOST")
	port     string = os.Getenv("PSQL_PORT")
	user     string = os.Getenv("PSQL_PORT")
	password string = os.Getenv("PSQL_PASSWORD")
	dbname   string = "interview_accountapi"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	// setup test data in db before running integration tests
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insertStatement := `insert into public."Account" 
	(id,organisation_id,version, is_deleted, is_locked,created_on, modified_on,record, pagination_id)
	values ($1,$2,$3, $4,$5, $6, $7, $8,$9)`

	_, err = db.Exec(insertStatement,
		"bc8fb900-d6fd-41d0-b187-dc23ba928712",
		"2e5087aa-2b22-11ec-bfa9-acde48001122",
		0,
		false,
		false,
		time.Now(),
		time.Now(),
		"{\"bic\": \"NWBKGB22\", \"name\": [\"hello world\"], \"bank_id\": \"400300\", \"country\": \"GB\", \"bank_id_code\": \"GBDSC\", \"base_currency\": \"GBP\", \"alternative_bank_account_names\": [\"abc\"]}",
		1)
	if err != nil {
		panic(err)
	}

	// run tests
	exitVal := m.Run()

	// clean up the table after test
	_, err = db.Exec("delete from public.\"Account\"")
	if err != nil {
		panic(err)
	}

	return exitVal

}

// --------------------- Create account method --------------------------- //

func TestAccountService_Create(t *testing.T) {

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient(f3client.WithHostUrl("http://localhost:8080"))
	if err != nil {
		panic(err)
	}

	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
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
	assert.Equal(t, expected.Attributes.Country, newAccount.Attributes.Country)
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
	actual := client.Accounts.Create(ctx, &newAccount)
	var target *f3client.ArgumentError

	if assert.Error(t, actual) {
		assert.ErrorAs(t, actual, &target)
	}
}

// ------------------- Delete account tests -------------------- //

func TestAccountService_Delete(t *testing.T) {

	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}

	success, err := client.Accounts.Delete(ctx, accountId, 0)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, true, success)
}

// ------------------- Fetch Account tests ---------------------- //

func TestAccountService_Fetch(t *testing.T) {

	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	client, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		panic(err)
	}

	fmt.Print(account)
}
