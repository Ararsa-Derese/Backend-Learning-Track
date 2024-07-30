package main

import (
	"fmt"

)

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
		fmt.Print("Enter a word: ")
		_, err = fmt.Scan(&word)

		if PalindromeCheck(word) {
			fmt.Println("The word is a palindrome")
		} else {
			fmt.Println("The word is not a palindrome")
		}
	} else if choice == 2 {
		var word string
		fmt.Print("Enter a word: ")
		_, err = fmt.Scan(&word)
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