package clock

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (rc *realClock) Now() time.Time { return time.Now() }

type fakeClock struct{ now time.Time }

func (fc *fakeClock) Now() time.Time { return fc.now }

// compile-time interface assertions
var _ Clock = (*realClock)(nil)
var _ Clock = (*fakeClock)(nil)

func NewRealClock() Clock {
	return &realClock{}
}

func NewFakeClock(t time.Time) Clock {
	return &fakeClock{now: t}
}
