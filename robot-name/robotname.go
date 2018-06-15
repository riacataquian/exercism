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
		r.name = RandStringBytes(2) + strconv.Itoa(RandInt(3))
	}
	return r.name
}

// Reset clears out a robot's name.
func (r *Robot) Reset() {
	r.name = ""
}

// RandStringBytes generate random string bytes from LetterBytes.
func RandStringBytes(n int) string {
	b := make([]byte, n)
	max := len(LetterBytes)
	for i := range b {
		b[i] = LetterBytes[rand.Intn(max)]
	}
	return string(b)
}

// RandInt generate a random int from 0...999.
func RandInt(n int) int {
	r := rand.Intn(999)
	for r < 100 {
		r = rand.Intn(999)
	}
	return r
}
