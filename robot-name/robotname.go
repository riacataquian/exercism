// Package robotname manage robot factory settings.
package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// LetterBytes is the allowed letters for robot name.
const LetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Robot describes a robot.
// It encapsulates a name.
type Robot struct {
	name string
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano() * rand.Int63())
}

// Name returns the name of the r robot.
//
// Generates a new name for robots without names yet.
func (r *Robot) Name() string {
	if r.name == "" {
		r.name = RandStringBytes() + strconv.Itoa(RandInt())
	}
	return r.name
}

// Reset clears out a robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

// RandStringBytes generate random string bytes from LetterBytes.
//
// Returns a 2 character string.
func RandStringBytes() string {
	b := make([]byte, 2)
	max := len(LetterBytes)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(max)]
	}
	return string(b)
}

// RandInt generate a random int from 100...999.
//
// Returns a 3 int.
func RandInt() int {
	r := rand.Intn(999)
	// Prevent values from 0 to 99.
	for r < 100 {
		r = rand.Intn(999)
	}
	return r
}
