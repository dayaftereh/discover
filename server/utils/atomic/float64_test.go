package atomic_test

import (
	"testing"

	"github.com/dayaftereh/discover/server/mathf"
	"github.com/dayaftereh/discover/server/utils/atomic"
	"github.com/stretchr/testify/assert"
)

func TestAtomicFloat64New(t *testing.T) {
	f := atomic.NewAtomicFloat64(42.31)
	assert.InDelta(t, 42.31, f.Get(), mathf.Epsilon)
}

func TestAtomicFloat64Set(t *testing.T) {
	f := atomic.NewAtomicFloat64(0.0)
	f.Set(42.31)
	assert.InDelta(t, 42.31, f.Get(), mathf.Epsilon)
}
