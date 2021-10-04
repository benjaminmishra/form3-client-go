package f3client_test

import (
	"testing"
	"time"

	"github.com/benjaminmishra/f3client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestConvertTo(t *testing.T) {

	accId := uuid.New()
	orgId := uuid.New()

	response := f3client.Response{
		Data: f3client.ResponseData{
			Type:           "account",
			ID:             accId,
			Version:        1,
			OrganisationID: orgId,
			CreatedOn:      time.Now().String(),
			ModifiedOn:     time.Now().String(),
			Attributes:     nil,
		},
		Links: f3client.Links{
			Self: "http://localhost:8080/v1/accounts/" + accId.String(),
		},
	}

	expected := f3client.Account{
		ID:             accId,
		Version:        1,
		OrganisationID: orgId,
		CreatedOn:      time.Now().String(),
		ModifiedOn:     time.Now().String(),
		Attributes:     f3client.AccountAttributes{},
	}
	acc := new(f3client.Account)
	err := response.ConvertTo(acc)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	assert.EqualValues(t, acc, expected)

}
