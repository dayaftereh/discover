package engine

import (
	"github.com/dayaftereh/discover/server/utils"
)

type Clock struct {
	startTime   float64
	elapsedTime float64
	oldTime     float64
	running     bool
}

func NewClock() *Clock {
	return &Clock{
		running: false,
	}
}

func (clock *Clock) StarTime() float64 {
	return clock.startTime / 1000.0 // to seconds
}

func (clock *Clock) ElapsedTime() float64 {
	return clock.elapsedTime / 1000.0 // to seconds
}

func (clock *Clock) Start() {
	clock.running = true
	clock.elapsedTime = 0.0
	clock.startTime = utils.SystemMillis()
	clock.oldTime = clock.startTime
}

func (clock *Clock) Stop() {
	clock.Delta()
	clock.running = false
}

func (clock *Clock) Delta() float64 {
	if !clock.running {
		clock.Start()
		return 0.0
	}

	newTime := utils.SystemMillis()
	diff := newTime - clock.oldTime
	clock.oldTime = newTime

	clock.elapsedTime += diff

	return diff / 1000.0 // to seconds
}
