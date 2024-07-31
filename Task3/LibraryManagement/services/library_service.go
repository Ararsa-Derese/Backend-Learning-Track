package services

import (
	"library/models"
	"sync"

)
type BookService struct {
	mu    sync.Mutex
	books []models.Book
	members []models.Member
	nextID uint
	memeberID uint
}
func NewBookService() *BookService {
	return &BookService{
		books:  []models.Book{},
		members: []models.Member{},
		nextID: 1,
		memeberID: 1,
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
func (s *BookService) BorrowBook(title,name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, book := range s.books {
		if book.Title == title {
			s.books[i].Status = "Borrowed"
			for j, member := range s.members {
				if member.Name == name {
					s.members[j].BorrowedBooks = append(s.members[j].BorrowedBooks, book)
					return
				}
			}
			s.members = append(s.members, models.Member{ID: s.memeberID, Name: name, BorrowedBooks: []models.Book{book}})
			s.memeberID++
			return
		}
	}
}
func (s *BookService) BorrowedBooks() []models.Member {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.members
}
func (s *BookService) ReturnBook(title string, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, member := range s.members {
		if member.Name == name {
			for j, book := range member.BorrowedBooks {
				if book.Title == title {
					s.members[i].BorrowedBooks = append(s.members[i].BorrowedBooks[:j], s.members[i].BorrowedBooks[j+1:]...)
					for k, book := range s.books {
						if book.Title == title {
							s.books[k].Status = "Available"
							break
						}
					}
					return
				}
			}
		}
	}
}
func (s *BookService) GetMember  (name string) []models.Member {
	s.mu.Lock()
	defer s.mu.Unlock()
	var themember []models.Member
	for _, member := range s.members {
		if member.Name == name {
			themember= append(themember, member)
		}
	}
	return themember
}