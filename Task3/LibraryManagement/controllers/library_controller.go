package controllers

import (
	"fmt"
	"library/models"
	"library/services"
)

type BookController struct {
	Service *services.BookService
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{Service: service}
}

func (c *BookController) AddBook(book models.Book) {
	if c.Service.GetBook(book.Title).Title != "" {
		fmt.Println("Book already exists")
		return
	}	
	c.Service.AddBook(book)
	fmt.Println("Book added")
}
func (c *BookController) GetBooks() []models.Book {
	return c.Service.GetBooks()
}
func (c *BookController) RemoveBook(title string) {
	c.Service.RemoveBook(title)
	fmt.Println("Book removed")
}
func (c *BookController) BorrowBook(title,name string) {
	books := c.Service.GetBooks()
	if len(books) == 0 {
		fmt.Println("No books available")
		return
	}
	if c.Service.GetBook(title).Title == "" {
		fmt.Println("Book not found")
		return
	}
	c.Service.BorrowBook(title,name)
	fmt.Println("Book borrowed")
}
func (c *BookController) GetBorrowedBooks() []models.Member {
	return c.Service.BorrowedBooks()
}
func (c *BookController) ReturnBook(title,name string) {
	books := c.Service.GetBooks()
	if len(books) == 0 {
		fmt.Println("No books available")
		return
	}
	if c.Service.GetBook(title).Title == "" {
		fmt.Println("Book not found")
		return
	}
	c.Service.ReturnBook(title,name)
	fmt.Println("Book returned")
}