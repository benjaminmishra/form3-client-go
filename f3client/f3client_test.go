package f3client_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	f3client "github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/stretchr/testify/assert"
)

// ------- f3client NewClient function unit tests --------- //

func TestNewClient_NoOptions(t *testing.T) {

	actual, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	assert.IsType(t, &f3client.Client{}, actual)
}

func TestNewClient_CustomHostUrl(t *testing.T) {
	actual, err := f3client.NewClient(f3client.WithHostUrl("http://foo.bar.com"))
	if err != nil {
		panic(err)
	}

	expected, err := url.Parse("http://foo.bar.com")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, *expected, actual.BaseURL)
}

func TestNewClient_CustomHttpClient(t *testing.T) {
	actual, err := f3client.NewClient(f3client.WithHttpClient(&http.Client{}))
	if err != nil {
		panic(err)
	}

	expected := &http.Client{}
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expected, actual.HttpClient)
}

// ------------- f3client NewRequest function unit tests -------------//

func TestNewRequest_WithoutUrlStr(t *testing.T) {
	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	_, err = c.NewRequest(context.Background(), f3client.Get, "", c.Accounts.ObjectType, nil)

	var targetErr *f3client.ArgumentError

	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &targetErr)
	}

}

func TestNewRequest_WithoutObjectType(t *testing.T) {
	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	_, err = c.NewRequest(context.Background(), f3client.Post, "/v1/org/acc", "", &f3client.Account{})

	var targetErr *f3client.ArgumentError

	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &targetErr)
	}

}

func TestNewRequest_Post_NilRequestBody(t *testing.T) {
	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	_, err = c.NewRequest(context.Background(), f3client.Post, "/v1/org/acc", c.Accounts.ObjectType, nil)

	var targetErr *f3client.ArgumentError

	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &targetErr)
	}

}

func TestNewRequest_Put_EmptyRequestBody(t *testing.T) {
	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	_, err = c.NewRequest(context.Background(), f3client.Put, "/v1/org/acc", c.Accounts.ObjectType, "")

	var targetErr *f3client.ArgumentError

	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &targetErr)
	}

}

// -------- f3client SendRequest function unit tests ---------------- //

func TestSendRequest_EmptyRequest(t *testing.T) {

	c, err := f3client.NewClient()
	if err != nil {
		panic(err)
	}
	_, err = c.SendRequest(context.Background(), &http.Request{})

	var targetErr *url.Error
	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &targetErr)
	}

}
