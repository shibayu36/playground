package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLock(t *testing.T) {
	expiredAt := time.Now().UTC().Add(10 * time.Second)

	locked1, release1, err := GetLock("lock1", expiredAt)
	require.NoError(t, err)
	assert.True(t, locked1)

	locked2, _, err := GetLock("lock1", expiredAt)
	require.NoError(t, err)
	assert.False(t, locked2, "lock failed because of lock1")

	release1()

	locked3, _, err := GetLock("lock1", expiredAt)
	require.NoError(t, err)
	assert.True(t, locked3, "lock succeeded because lock1 was released")
}

func TestLockExpired(t *testing.T) {
	expiredAt := time.Now().UTC().Add(3 * time.Second)

	locked1, _, err := GetLock("lock1_with_expired", expiredAt)
	require.NoError(t, err)
	assert.True(t, locked1)

	locked2, _, err := GetLock("lock1_with_expired", expiredAt)
	require.NoError(t, err)
	assert.False(t, locked2, "lock failed because of lock1_with_expired")

	time.Sleep(3 * time.Second)

	locked3, release, err := GetLock("lock1_with_expired", expiredAt)
	require.NoError(t, err)
	assert.True(t, locked3, "lock succeeded because lock1_with_expired was expired")

	release()
}
