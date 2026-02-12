package main

import (
	"testing"
	"time"

	"github.com/go-pkgz/expirable-cache/v3"
	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	// make cache with short TTL and 3 max keys
	c := cache.NewCache[string, string]().WithMaxKeys(3).WithTTL(24 * time.Hour)

	// with 0 ttl (last parameter) will use cache-wide setting instead (10ms).
	c.Set("key1", "val1", 0)
	c.Set("key2", "val2", 0)
	c.Set("key3", "val3", 0)
	c.Set("key4", "val4", 0)

	// get value under key1
	r, ok := c.Get("key1")
	require.False(t, ok)

	t.Logf("%+v\n", r)
}
