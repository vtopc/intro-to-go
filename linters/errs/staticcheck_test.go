package errs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_staticcheck(t *testing.T) {
	err := staticcheck1()
	require.Error(t, err)
}
