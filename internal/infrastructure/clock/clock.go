package clock

import (
	"time"

	"github.com/hogartr/go-hexagonal-template/internal/domain"
)

type Clock interface {
	Now() domain.Date
}

type realClock struct{}

func (rc *realClock) Now() domain.Date { return domain.Date(time.Now()) }

type fakeClock struct{ now domain.Date }

func (fc *fakeClock) Now() domain.Date { return fc.now }

// compile-time interface assertions
var _ Clock = (*realClock)(nil)
var _ Clock = (*fakeClock)(nil)

func NewRealClock() Clock {
	return &realClock{}
}

func NewFakeClock(t domain.Date) Clock {
	return &fakeClock{now: t}
}
