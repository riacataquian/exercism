// Package letter counts the frequency of letters in texts using parallel computation.
package letter

// FreqMap maps runes and their frequency counts.
type FreqMap map[rune]int

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
	return FreqMap{}
}
