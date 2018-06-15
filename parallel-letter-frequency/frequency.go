// Package letter counts the frequency of letters in strs using parallel computation.
package letter

import (
	"sync"
)

// FreqMap maps runes and their frequency counts.
type FreqMap map[rune]int

// SafeCounter holds the frequency map and its locker.
type SafeCounter struct {
	v   FreqMap
	mux sync.Mutex
}

// Frequency counts the runes and their frequency synchronously.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the runes and their frequency concurrently.
func ConcurrentFrequency(strs []string) FreqMap {
	c := SafeCounter{v: make(FreqMap)}

	for _, str := range strs {
		for _, s := range str {
			go c.Inc(s)
		}
	}

	return c.v
}

// Inc increments a key's count.
func (c *SafeCounter) Inc(r rune) {
	c.mux.Lock()
	c.v[r]++ // Lock so only one goroutine at a time can access the map c.v.
	c.mux.Unlock()
}
