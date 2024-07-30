package main

import (
	"testing"

)

func TestWordFrequencyCount(t *testing.T) {
	word := "hello"
	result := wordFrequencyCount(word)
	if result["h"] != 1 || result["e"] != 1 || result["l"] != 2 || result["o"] != 1 {
		t.Errorf("Expected map[h:1 e:1 l:2 o:1], got %v", result)
	}
}