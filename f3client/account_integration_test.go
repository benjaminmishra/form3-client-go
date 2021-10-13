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
	user     string = os.Getenv("PSQL_USER")
	password string = os.Getenv("PSQL_PASSWORD")
	dbname   string = "interview_accountapi"
	apihost  string = os.Getenv("API_HOST")
)

// setup func inserts records into the database before running the test case
// it returns an function that can then be used to cleanup the database after the tests are done
// if there is an error setting up the database then this function returns an err and a teardown func
// that only warps db.close() method
func setup() (func(), error) {
	// setup test data in db before running integration tests
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return func() { db.Close() }, err
	}

	// define teardown func , to be used later to cleanup db
	teardown := func() {
		// clean up the table after test
		_, err = db.Exec("delete from public.\"Account\"")
		if err != nil {
			db.Close()
		}
		db.Close()
	}

	/// start with a clean table
	_, err = db.Exec("delete from public.\"Account\"")
	if err != nil {
		return func() { db.Close() }, err
	}

	// insert your record
	insertStatement := `insert into public."Account" 
	(id,organisation_id,version, is_deleted, is_locked,created_on, modified_on,record)
	values ($1,$2,$3, $4,$5, $6, $7, $8)`

	_, err = db.Exec(insertStatement,
		"bc8fb900-d6fd-41d0-b187-dc23ba928712",
		"2e5087aa-2b22-11ec-bfa9-acde48001122",
		0,
		false,
		false,
		time.Now(),
		time.Now(),
		"{\"bic\": \"NWBKGB22\", \"name\": [\"hello world\"], \"bank_id\": \"400300\", \"country\": \"GB\", \"bank_id_code\": \"GBDSC\", \"base_currency\": \"GBP\", \"alternative_bank_account_names\": [\"abc\"]}")
	if err != nil {
		return teardown, err
	}

	// return a teardown function, that is later used to cleanup the db
	return teardown, nil
}

// --------------------- Create account method --------------------------- //

func Test_Integration_AccountService_Create(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		log.Println(err.Error())
		return
	}

	accountId, err := uuid.NewRandom()
	if err != nil {
		log.Println(err.Error())
		return
	}
	orgId, err := uuid.NewUUID()
	if err != nil {
		log.Println(err.Error())
		return
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
			Name:              []string{"hello world1"},
			AlternativeNames:  []string{"abc"},
		},
	}

	// call the create account api
	err = client.Accounts.Create(ctx, &newAccount)
	if err != nil {
		log.Println(err.Error())
		return
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

func Test_Integration_AccountService_Create_EmptyBody(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	// instantiate the client
	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	_, err = uuid.NewUUID()
	if err != nil {
		log.Println(err.Error())
		return
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

// ------------------- Fetch Account tests ---------------------- //

func Test_Integration_AccountService_Fetch(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Fatal(err)
	}

	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		log.Println(err.Error())
		return
	}
	account, err := client.Accounts.Fetch(ctx, accountId)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.NoError(t, err, fmt.Sprint("Found Account with Id : "+account.ID.String()))
}

func Test_Integration_AccountService_Fetch_NotFound(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	accountId, err := uuid.NewRandom()
	if err != nil {
		log.Println(err.Error())
		return
	}

	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		panic(err)
	}
	_, err = client.Accounts.Fetch(ctx, accountId)

	if !assert.Error(t, err) {
		assert.Fail(t, "Expected error but there was no error")
	} else {
		assert.Positive(t, "Error generated as expected")
	}
}

// ------------------- Delete account tests -------------------- //

func Test_Integration_AccountService_Delete(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Println(err.Error())
		return
	}

	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		log.Println(err.Error())
		return
	}

	success, err := client.Accounts.Delete(ctx, accountId, 0)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.Equal(t, true, success)
}

func Test_Integration_AccountService_Delete_WrongVersion(t *testing.T) {

	teardown, err := setup()
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer teardown()

	ctx := context.Background()
	accountId, err := uuid.Parse("bc8fb900-d6fd-41d0-b187-dc23ba928712")
	if err != nil {
		log.Println(err.Error())
		return
	}

	client, err := f3client.NewClient(f3client.WithHostUrl(apihost))
	if err != nil {
		log.Println(err.Error())
		return
	}

	_, err = client.Accounts.Delete(ctx, accountId, 1)
	if err != nil {
		assert.Error(t, err, err.Error())
	} else {
		assert.Fail(t, "Expected error but no error was found")
	}

}
