package errs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ineffassign(t *testing.T) {
	err := ineffassign()
	require.Error(t, err)
}
