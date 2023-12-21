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

	err = release1()
	require.NoError(t, err)

	locked3, release3, err := GetLock("lock1", expiredAt)
	require.NoError(t, err)
	assert.True(t, locked3, "lock succeeded because lock1 was released")

	err = release3()
	require.NoError(t, err)
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

func TestReleaseByOther(t *testing.T) {
	locked1, release1, err := GetLock("lock1", time.Now().UTC().Add(2*time.Second))
	require.NoError(t, err)
	assert.True(t, locked1)

	time.Sleep(2 * time.Second)

	locked2, release2, err := GetLock("lock1", time.Now().UTC().Add(2*time.Second))
	require.NoError(t, err)
	assert.True(t, locked2, "lock succeeded because lock1 was expired")

	err = release1()
	require.NoError(t, err, "release1 should be successful but ")

	locked3, _, err := GetLock("lock1", time.Now().UTC().Add(2*time.Second))
	require.NoError(t, err)
	assert.False(t, locked3, "lock failed because release1 does not release lock1")

	err = release2()
	require.NoError(t, err, "release2 should be successful")

	locked4, release4, err := GetLock("lock1", time.Now().UTC().Add(2*time.Second))
	require.NoError(t, err)
	assert.True(t, locked4, "lock succeeded because lock1 was released by release2")

	release4()
}
