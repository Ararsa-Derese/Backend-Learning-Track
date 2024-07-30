package main
func wordFrequencyCount(word string) map[string]int {
	wordFrequency := make(map[string]int)
	for _, char := range word {
		wordFrequency[string(char)]++
	}
	return wordFrequency
}

