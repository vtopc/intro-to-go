package ai

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode_Act(t *testing.T) {
	tests := []struct {
		name string
		v    uint8
		node *Node
		want uint8
	}{
		{
			name: "not",
			v:    0,
			node: &Node{
				action:   Not,
				constant: 0,
			},
			want: 0xff,
		},
		{
			name: "xor-0",
			v:    0,
			node: &Node{
				action:   Xor,
				constant: 1,
			},
			want: 1,
		},
		{
			name: "xor-1",
			v:    1,
			node: &Node{
				action:   Xor,
				constant: 1,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.node.Act(tt.v)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
