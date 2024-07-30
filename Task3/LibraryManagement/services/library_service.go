package services

import (
	"library/models"
	"sync"

)
type BookService struct {
	mu    sync.Mutex
	books []models.Book
	nextID uint
}
func NewBookService() *BookService {
	return &BookService{
		books:  []models.Book{},
		nextID: 1,
	}
}
func (s *BookService) AddBook(book models.Book) {
	s.mu.Lock()
	defer s.mu.Unlock()
	book.ID = s.nextID
	s.nextID++
	s.books = append(s.books, book)
}
func (s *BookService) GetBooks() []models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.books
}
func (s *BookService) GetBook(title string) models.Book {
	s.mu.Lock()
	defer s.mu.Unlock()
	var thebook models.Book
	for _, book := range s.books {
		if book.Title == title {
			thebook= book
			break
		}
	}
	return thebook
} 
func (s *BookService) RemoveBook(title string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, book := range s.books {
		if book.Title == title {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return
		}
	}
}
