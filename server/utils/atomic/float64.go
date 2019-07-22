package atomic

import (
	"math"
	"sync/atomic"
)

type AtomicFloat64 uint64

func NewAtomicFloat64(value float64) *AtomicFloat64 {
	atomicFloat64 := new(AtomicFloat64)
	// set the given value
	atomicFloat64.Set(value)
	return atomicFloat64

}

func (atomicFloat64 *AtomicFloat64) Set(f float64) {
	value := math.Float64bits(f)
	atomic.StoreUint64((*uint64)(atomicFloat64), value)
}

func (atomicFloat64 *AtomicFloat64) Get() float64 {
	return math.Float64frombits(atomic.LoadUint64((*uint64)(atomicFloat64)))
}
