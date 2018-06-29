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
	c := make(chan FreqMap)
	for _, s := range strs {
		go func(s string) {
			c <- Frequency(s) // send to channel.
		}(s)
	}

	m := FreqMap{}
	for range strs {
		for r, f := range <-c { // receive from channel.
			m[r] += f
		}
	}
	return m
}
