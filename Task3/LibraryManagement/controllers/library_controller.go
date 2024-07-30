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