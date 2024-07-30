package main

import (
	"fmt"

)
func validateword(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return len(name) > 0
}

func main(){

	var choice int
	choice:
	fmt.Println("Enter your Intention: ")
	fmt.Print("1. Palindrome Check\n2. Word Frequency Count\n")
	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Error in reading input")
		return
	}

	if choice == 1 {
		var word string
		Enterword:
		fmt.Print("Enter a word: ")
		_, err = fmt.Scan(&word)
		if !validateword(word) {
			fmt.Println("Word should only contain letters")
			goto Enterword
		}
		if PalindromeCheck(word) {
			fmt.Println("The word is a palindrome")
		} else {
			fmt.Println("The word is not a palindrome")
		}
	} else if choice == 2 {
		var word string
		Enter1word:
		fmt.Print("Enter a word: ")
		_, err = fmt.Scan(&word)
		if !validateword(word) {
			fmt.Println("Word should only contain letters")
			goto Enter1word
		}
		wordFrequency := wordFrequencyCount(word)
		fmt.Println("Word Frequency: ")
		for key, value := range wordFrequency {
			fmt.Println(key, ":", value)
		}
	} else {
		fmt.Println("Invalid choice")
		goto choice
	}

	if err != nil {
		fmt.Println("Error in reading input")
		return
	}
}