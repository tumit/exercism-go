package letter

// FreqMap is map of rune frequency
type FreqMap map[rune]int

// Frequency it return frequency of string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency it return frequency of string in concurrent process
func ConcurrentFrequency(ss []string) FreqMap {

	mc := make(chan FreqMap, len(ss))
	for _, s := range ss {
		go func(s string) {
			mc <- Frequency(s)
		}(s)
	}

	fm := FreqMap{}

	for range ss {
		m := <-mc
		for k, v := range m {
			fm[k] += v
		}
	}

	return fm
}
