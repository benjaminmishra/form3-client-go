package f3client_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_FetchAccount(t *testing.T) {
	client, server, _, teardown := setup()
	defer teardown()

	var resourceId = "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	u := "organisation/accounts/" + resourceId

	server.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			r := strings.Split(r.RequestURI, "/")

			if r[len(r)-1] == resourceId {
				w.Write([]byte(`{
				"data": {
				  "type": "accounts",
				  "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
				  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
				  "version": 0,
				  "attributes": {
					"country": "GB",
					"base_currency": "GBP",
					"account_number": "41426819",
					"bank_id": "400300",
					"bank_id_code": "GBDSC",
					"bic": "NWBKGB22",
					"iban": "GB11NWBK40030041426819",
					"status": "confirmed"
				  },
				  "relationships": {
					"account_events": {
					  "data": [
						{
						  "type": "account_events",
						  "id": "c1023677-70ee-417a-9a6a-e211241f1e9c"
						},
						{
						  "type": "account_events",
						  "id": "aca32528-d4cf-4d54-93fe-5d80d27ab773"
						}
					  ]
					}
				  }
				}
			  }
			  `))
			} else {
				w.Write([]byte("Not Found!!"))
			}
		}
	})

	resourceUUID, _ := uuid.Parse(resourceId)

	artifact := client.Accounts.Fetch(resourceUUID)
	var ver int64 = 0
	var country string = "GB"
	var status string = "confirmed"

	account := f3client.Account{
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Version:        &ver,
		Attributes: &f3client.AccountAttributes{
			Country:       &country,
			BaseCurrency:  "GBP",
			AccountNumber: "41426819",
			BankID:        "400300",
			BankIDCode:    "GBDSC",
			Bic:           "NWBKGB22",
			Iban:          "GB11NWBK40030041426819",
			Status:        &status,
		},
	}

	expected, _ := json.Marshal(account)

	assert.Equal(t, expected, artifact)

}