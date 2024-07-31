package services

import (
    "errors"
    "library/models"
	"strconv"
)

type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
    books   map[int]models.Book
    members map[int]models.Member
}

func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

func (l *Library) AddBook(book models.Book) {
    l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
    delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }
    if book.Status != "Available" {
        return errors.New("book is not available")
    }

    member, exists := l.members[memberID]
    if !exists {
        member = models.Member{ID: memberID, Name: "Member" + strconv.Itoa(memberID)}
    }

    book.Status = "Borrowed"
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.books[bookID] = book
    l.members[memberID] = member

    return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }

    member, exists := l.members[memberID]
    if !exists {
        return errors.New("member not found")
    }

    borrowedBooks := member.BorrowedBooks
    for i, b := range borrowedBooks {
        if b.ID == bookID {
            member.BorrowedBooks = append(borrowedBooks[:i], borrowedBooks[i+1:]...)
            break
        }
    }

    book.Status = "Available"
    l.books[bookID] = book
    l.members[memberID] = member

    return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
    availableBooks := []models.Book{}
    for _, book := range l.books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, exists := l.members[memberID]
    if !exists {
        return []models.Book{}
    }
    return member.BorrowedBooks
}
