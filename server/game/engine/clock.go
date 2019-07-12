package engine

import (
	"github.com/dayaftereh/discover/server/utils"
)

type Clock struct {
	StartTime   float64
	ElapsedTime float64
	// private
	oldTime float64
	running bool
}

func NewClock() *Clock {
	return &Clock{
		running: false,
	}
}

func (clock *Clock) Start() {
	clock.running = true
	clock.ElapsedTime = 0.0
	clock.StartTime = utils.SystemMillis()
	clock.oldTime = clock.StartTime
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

	clock.ElapsedTime += diff

	return diff
}
