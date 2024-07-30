package main

import (
	"fmt"
	"strconv"
)

func calculateAverage(grades map[string]int) float64 {
	var sum float64
	for _, grade := range grades {
		sum += float64(grade)
	}
	return sum / float64(len(grades))
}

func validatename(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return len(name) > 0
}


func main() {

	var name string
	var subject string
entername:
	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if !validatename(name) {
		fmt.Println("Name should only contain letters")
		goto entername
	}
	if err != nil {
		fmt.Println("error reading data")
		return
	}
entersubject:
	fmt.Print("Enter How many subjects you take: ")
	_, err = fmt.Scan(&subject)
	if err != nil {
		fmt.Println("Error in reading input")
		goto entersubject
	}
	
	 num, err := strconv.Atoi(subject)
	if err != nil {
		fmt.Println("Please enter a number not a letter")
		goto entersubject
	}

	subjectswithgrades := make(map[string]int)
	for i := 0; i < num; i++ {
		var subjectname string
		var grade string
	entersubjectname:
		fmt.Print("Enter subject name ", i+1, ": ")
		_, err = fmt.Scan(&subjectname)
		if !validatename(subjectname) {
			fmt.Println("Subject name should only contain letters")
			goto entersubjectname
		}

		if err != nil {
			fmt.Println("Error in reading input")
			return
		}
	reEnter:
		fmt.Print("Enter grade: ")
		_, err = fmt.Scan(&grade)
		if err != nil {
			fmt.Println("Error in reading input")
			goto entersubject
		}
		
		g, err := strconv.Atoi(grade)
		if err != nil {
			fmt.Println("Please enter a number not a letter")
			goto reEnter
		}
		if g < 0 || g > 100 {
			fmt.Println("Grade should be between 0 and 100")
			goto reEnter
		}
		subjectswithgrades[subjectname] = g
	}

	fmt.Println("Your average is: ", calculateAverage(subjectswithgrades))

}
