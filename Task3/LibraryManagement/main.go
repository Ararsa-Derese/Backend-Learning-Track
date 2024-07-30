package main

import (
	"fmt"
	"library/controllers"
	"library/models"
	"library/services"
)

func validateword(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return len(name) > 0
	
}
func display(bookcontroller controllers.BookController) {
	var choice string
	for {
	fmt.Println("############################################")
	fmt.Println("Hello, welcome to Library management system")
	fmt.Println("1. Add a new book")
	fmt.Println("2. Remove an existing book.")
	fmt.Println("3. Borrow a book.")
	fmt.Println("4. Return a book.")
	fmt.Println("5. List all available books.")
	fmt.Println("6. List all borrowed books by a member.")
	fmt.Println("7. Exit")
	fmt.Println("Enter your choice")
	fmt.Scanln(&choice)
	switch choice {
	case "1":
		addBook(bookcontroller)
	case "2":
		removeBook(bookcontroller)
	// case "3":
	// 	borrowBook()
	// case "4":
	// 	returnBook()
	case "5":
		listAvailableBooks(bookcontroller)
	// case "6":
	// 	listBorrowedBooks()
	// case "7":
	// 	return 
	default:
		fmt.Println("Invalid choice")
	}
	if choice == "7"{
		break
	}
	}
	
}
func addBook(bookcontroller controllers.BookController) {
	var title,author,status string
	fmt.Println("Please Enter the necessary informatin: ")
	for {
	fmt.Println("Please Enter title of the book")
	_,err := fmt.Scanln(&title)
	if err!= nil{
		fmt.Println("Error reading data")
		continue

	}
	if !validateword(title){
		fmt.Println("Please enter a valid Title")
		continue
	}
	break	
}

for {
	fmt.Println("Please Enter author of the book")
	_,err := fmt.Scanln(&author)
	if err!= nil{
		fmt.Println("Error reading data")
		continue

	}
	if !validateword(author){
		fmt.Println("Please enter a valid author")
		continue
	}
	break
}
for {
	fmt.Println("Please Enter status of the book(Available/Borrowed)")
	_,err := fmt.Scanln(&status)
	if err!= nil{
		fmt.Println("Error reading data")
		continue

	}
	if !validateword(status){
		fmt.Println("Please enter a valid status")
		continue
	}
	if status != "Available" && status != "Borrowed"{
		fmt.Println("Please enter a valid status")
		continue
	}
	break
}
book := models.Book{
	Title: title,
	Author: author,
	Status: status,
}
 bookcontroller.AddBook(book)
}
func listAvailableBooks(bookcontroller controllers.BookController) {
	fmt.Println("List of available books")
	fmt.Println("|Title|\t|Author|\t|Status|")
	books := bookcontroller.GetBooks()
	if len(books) == 0 {
		fmt.Println("No books available")
		return
	}
	for _, book := range books {
		if book.Status == "Available" {
			fmt.Println("|",book.Title,"|", "\t", book.Author, "\t", book.Status)
		}
	}

}
func removeBook(bookcontroller controllers.BookController) {
	var title string
	fmt.Println("Please Enter the title of the book to be removed")
	_,err := fmt.Scanln(&title)
	if err!= nil{
		fmt.Println("Error reading data")
		return

	}
	if !validateword(title){
		fmt.Println("Please enter a valid Title")
		return
	}
	bookcontroller.RemoveBook(title)
}
func main() {
	bookservice := services.NewBookService()
	bookcontroller := controllers.NewBookController(bookservice)

    display(*bookcontroller)
	
}
