package main 

import (
    "testing"
)

func TestPalindromeCheck(t *testing.T) {
	word := "madam"
	result := PalindromeCheck(word)
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}