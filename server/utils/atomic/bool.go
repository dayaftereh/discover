package atomic

import (
	"sync/atomic"
)

type AtomicBool int32

func NewAtomicBool(ok bool) *AtomicBool {
	atomicBool := new(AtomicBool)
	atomicBool.Set(ok)
	return atomicBool
}

func (atomicBool *AtomicBool) Set(b bool) {
	if b {
		atomic.StoreInt32((*int32)(atomicBool), 1)
	} else {
		atomic.StoreInt32((*int32)(atomicBool), 0)
	}
}

func (atomicBool *AtomicBool) GetAndSet(b bool) bool {
	if b {
		return atomic.SwapInt32((*int32)(atomicBool), 1) == 1
	}
	return atomic.SwapInt32((*int32)(atomicBool), 0) == 1
}

func (atomicBool *AtomicBool) Get() bool {
	return atomic.LoadInt32((*int32)(atomicBool)) == 1
}
