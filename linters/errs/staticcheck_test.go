package errs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_staticcheck(t *testing.T) {
	err := staticcheck()
	require.Error(t, err)
}
