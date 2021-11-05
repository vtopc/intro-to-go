package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatJSON(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want Cat
	}{
		{
			b:    []byte(`{"name":"Cheshire"}`),
			want: Cat{Animal: &Animal{Name: "Cheshire"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Cat
			err := json.Unmarshal(tt.b, &got)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
