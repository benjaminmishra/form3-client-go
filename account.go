// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
package f3client

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type AccountService service

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
	SecondaryIdentification bool     `json:"secondary_identification,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
	ProcessingService       string   `json:"processing_service,omitempty"`
	UserDefinedInformation  string   `json:"user_defined_information,omitempty"`
	ValidationType          string   `json:"validation_type,omitempty"`
	ReferenceMask           string   `json:"reference_mask,omitempty"`
	AcceptanceQualifier     string   `json:"acceptance_qualifier,omitempty"`
}

func (account *AccountService) Create(ctx context.Context, acc *Account) error {
	path := account.client.Version + "/organisation/accounts"

	req, err := account.client.NewRequest(ctx, Post, path, "accounts", acc)
	if err != nil {
		return err
	}

	resp, err := account.client.SendRequest(ctx, req)
	if err != nil {
		return err
	}

	resp.ConvertTo(acc)

	return nil
}

func (account *AccountService) Fetch(ctx context.Context, accountId uuid.UUID) (*Account, error) {

	acc := new(Account)

	path := account.client.Version + "/organisation/accounts/" + accountId.String()

	req, err := account.client.NewRequest(ctx, Get, path, "accounts", nil)
	if err != nil {
		return nil, err
	}

	resp, err := account.client.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	encoded, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(encoded, acc)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (account *AccountService) Delete(accountId uuid.UUID) error {
	return nil
}
