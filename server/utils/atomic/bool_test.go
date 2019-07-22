package atomic_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/utils/atomic"
	"github.com/stretchr/testify/assert"
)

func TestAtomicBoolNew(t *testing.T) {
	b := atomic.NewAtomicBool(true)
	assert.True(t, b.Get())
}

func TestAtomicBoolSet(t *testing.T) {
	b := atomic.NewAtomicBool(false)
	assert.False(t, b.Get())
	b.Set(true)
	assert.True(t, b.Get())
}

func TestAtomicBoolGetAndSet(t *testing.T) {
	b := atomic.NewAtomicBool(false)
	assert.False(t, b.GetAndSet(true))
	assert.True(t, b.Get())
}
