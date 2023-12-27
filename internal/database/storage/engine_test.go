package storage

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEngine(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	storage := NewStorage()
	key := "test"
	value := "lol"

	v, ok := storage.Get(ctx, key)
	require.False(t, ok)
	require.Equal(t, "", v)

	err := storage.Set(ctx, key, value)
	require.Nil(t, err)

	v, ok = storage.Get(ctx, key)
	require.True(t, ok)
	require.Equal(t, value, v)

	err = storage.Del(ctx, key)
	require.Nil(t, err)

	v, ok = storage.Get(ctx, key)
	require.False(t, ok)
	require.Equal(t, "", v)
}
