// Package clock implements a clock that handles times without dates.
package clock

import (
	"fmt"
)

// Clock describes a clock object.
type Clock struct {
	hour int
	min  int
}

// New is the Clock constructor.
func New(hour, min int) Clock {
	const (
		maxMin  = 60
		maxHour = 24
	)

	// rmin is the remainder of a given minute after dividing it to hours.
	rmin := min % maxMin
	if rmin < 0 {
		rmin += maxMin
		hour--
	}

	hoursPerMins := min / maxMin
	totalHours := (hour + hoursPerMins) % maxHour
	if totalHours < 0 {
		totalHours += maxHour
	}

	return Clock{
		hour: totalHours,
		min:  rmin,
	}
}

// String is the Stringer implementation of Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Subtract returns a new clock subtracted the minute provided.
//
// It also handles subtraction by accepting negative values.
func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.min-m)
}

// Add returns a new clock added the minute provided.
func (c Clock) Add(m int) Clock {
	return New(c.hour, c.min+m)
}
