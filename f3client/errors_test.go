package f3client_test

import (
	"testing"

	"github.com/benjaminmishra/form3-client-go/f3client"
	"github.com/stretchr/testify/assert"
)

func TestNewArgError(t *testing.T) {

	err := f3client.NewArgError("foo", "foo is missing")
	var target *f3client.ArgumentError

	if assert.Error(t, err) {
		assert.ErrorAs(t, err, &target)
	}

}
