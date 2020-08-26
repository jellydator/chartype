package chartype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// equalError uses testify's assert package to check if errors
// are equal or, if assert.AnError is expected, whether an error exists
// or not.
func equalError(t *testing.T, exp, err error) {
	t.Helper()

	if exp != nil {
		if exp == assert.AnError { //nolint:goerr113 // direct check is needed
			assert.Error(t, err)
			return
		}

		assert.Equal(t, exp, err)

		return
	}

	assert.NoError(t, err)
}
