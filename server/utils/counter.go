package utils

type IDCounter struct {
	n int64
}

func NewIDCounter() *IDCounter {
	return &IDCounter{
		n: 0,
	}
}

func (counter *IDCounter) Next() int64 {
	counter.n = counter.n + 1
	return counter.n
}
