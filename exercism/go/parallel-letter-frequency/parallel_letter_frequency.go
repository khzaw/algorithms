package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.

func ConcurrentFrequency(l []string) FreqMap {
	m := make(FreqMap)
	c := make(chan byte, 1)

	for _, s := range l {
		s := s
		go func() {
			for _, r := range s {
				m[r]++
			}
			c <- 1
		}()
		<-c
	}
	return m
}
