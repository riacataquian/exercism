// Package letter ...
package letter

// FreqMap ...
type FreqMap map[rune]int

// Frequency ...
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency ...
func ConcurrentFrequency(strs []string) FreqMap {
	return FreqMap{}
}
