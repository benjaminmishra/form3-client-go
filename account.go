package f3client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type AccountService struct {
	service
	ObjectType string
}

// Account represents an account in the form3 org section.
//
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type Account struct {
	ID             uuid.UUID         `json:"id,omitempty"`
	Version        int               `json:"version,omitempty"`
	OrganisationID uuid.UUID         `json:"organisation_id,omitempty"`
	CreatedOn      string            `json:"created_on,omitempty"`
	ModifiedOn     string            `json:"modified_on,omitempty"`
	Attributes     AccountAttributes `json:"attributes,omitempty"`
}

type AccountAttributes struct {
	Country                 string   `json:"country,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	CustomerID              string   `json:"customer_id,omitempty"`
	Name                    []string `json:"name,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	AccountClassification   string   `json:"account_classification,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
	ProcessingService       string   `json:"processing_service,omitempty"`
	UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
	ValidationType          string   `json:"validation_type,omitempty"`
	ReferenceMask           string   `json:"reference_mask,omitempty"`
	AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	Status                  string   `json:"status,omitempty"`
}

func (as *AccountService) Create(ctx context.Context, account *Account) error {
	var err error

	path := "/v1/organisation/accounts"

	// validate for mandatory account fields before creating new request
	if account.Attributes.Country == "" {
		return NewArgError("country", "country is mandatory for account create request")
	} else if account.Attributes.Name == nil {
		return NewArgError("name", "names are mandatory for account create request")
	}

	// create new request
	req, err := as.client.NewRequest(ctx, Post, path, as.ObjectType, account)
	if err != nil {
		return err
	}

	// send the request and catch the response
	resp, err := as.client.SendRequest(ctx, req)
	if err != nil {
		return err
	}

	err = resp.ConvertTo(account)
	if err != nil {
		return err
	}

	return nil
}

func (as *AccountService) Fetch(ctx context.Context, accountId uuid.UUID) (*Account, error) {

	acc := new(Account)

	path := "/v1/organisation/accounts/" + accountId.String()

	req, err := as.client.NewRequest(ctx, Get, path, as.ObjectType, nil)
	if err != nil {
		return acc, err
	}

	resp, err := as.client.SendRequest(ctx, req)
	if err != nil {
		return acc, err
	}

	err = resp.ConvertTo(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (as *AccountService) Delete(ctx context.Context, accountId uuid.UUID, accountVersion int) (bool, error) {
	path := "/v1/organisation/accounts/" + accountId.String() + "?version=" + fmt.Sprint(accountVersion)

	req, err := as.client.NewRequest(ctx, Delete, path, as.ObjectType, nil)
	if err != nil {
		return false, err
	}
	// no return expected , hence ignore the response object
	resp, err := as.client.SendRequest(ctx, req)
	if err != nil {
		return false, err
	}

	if resp == nil {
		return true, nil
	}

	return false, nil
}
