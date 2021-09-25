// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
package f3client

import (
	"github.com/google/uuid"
)

type AccountService service

type Account struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type CreateAccountRequest struct {
	Country                 *string  `json:"country,omitempty"`
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

func (account *AccountService) Create(accountId uuid.UUID, orgnisationId uuid.UUID, createAccReq *CreateAccountRequest) (*Account, error) {
	return &Account{}, nil
}

func (account *AccountService) Fetch(accountId uuid.UUID) (*Account, error) {

	return &Account{}, nil
}

func (account *AccountService) Delete(accountId uuid.UUID) error {
	return nil
}
