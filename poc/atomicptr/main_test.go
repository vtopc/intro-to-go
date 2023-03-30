package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAtomicPointer(t *testing.T) {
	v := ptr[int](1)

	var s1 S
	s1.i.Store(v)
	require.Equal(t, 1, *s1.i.Load())

	var s2 S
	s2.i.Store(v)
	require.Equal(t, 1, *s2.i.Load())

	s1.i.Store(ptr[int](2))
	assert.Equal(t, 2, *s1.i.Load())
	assert.Equal(t, 2, *s2.i.Load())
}

func ptr[T any](v T) *T {
	return &v
}
