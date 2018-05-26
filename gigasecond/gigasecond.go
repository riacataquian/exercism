// Package gigasecond calculates the moment when someone has lived for 10^9 seconds.
package gigasecond

// import path for the time package from the standard library
import "time"

// Gigasecond is 10^9 (1,000,000,000) seconds.
const Gigasecond = time.Duration(1e9) * time.Second

// AddGigasecond adds Gigasecond to current time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
