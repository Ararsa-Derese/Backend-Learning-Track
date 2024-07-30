package main

import (
	"fmt"
)

func calculateAverage(grades map[string]int) float64 {
	var sum float64
	for _, grade := range grades {
		sum += float64(grade)
	}
	return sum / float64(len(grades))
}

func main() {

	var name string
	var subject int
	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error in reading input")
		return
	}
	fmt.Print("Enter How many subjects you take: ")
	_, err = fmt.Scan(&subject)
	if err != nil {
		fmt.Println("Error in reading input")
		return
	}

	subjectswithgrades := make(map[string]int)

	for i := 0; i < subject; i++ {
		var subjectname string
		var grade int
		fmt.Print("Enter subject name ", i+1, ": ")
		_, err = fmt.Scan(&subjectname)
		if err != nil {
			fmt.Println("Error in reading input")
			return
		}
	reEnter:
		fmt.Print("Enter grade: ")
		_, err = fmt.Scan(&grade)
		if err != nil {
			fmt.Println("Error in reading input")
			return
		}
		if grade < 0 || grade > 100 {
			fmt.Println("Grade should be between 0 and 100")
			goto reEnter
		}
		subjectswithgrades[subjectname] = grade
	}

	fmt.Println("Your average is: ", calculateAverage(subjectswithgrades))

}
