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
		hour-- // offset
	}

	// hoursPerMins is the count of hours given a minute.
	hoursPerMins := min / maxMin

	// totalHours is the remainder of sum of hours and hoursPerMins after dividing it to maxHour (or a day),
	// we only care about the remainder.
	totalHours := (hour + hoursPerMins) % maxHour
	if totalHours < 0 {
		// If total hours is less than 0, it is a negative number.
		// Add to maxHour (24) to compute for the diff.
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
