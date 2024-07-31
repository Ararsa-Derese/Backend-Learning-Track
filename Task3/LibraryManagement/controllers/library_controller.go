package controllers

import (
    "bufio"
    "fmt"
    "strconv"
    "strings"
    "library/models"
    "library/services"
)

type LibraryController struct {
    library services.LibraryManager
}

func NewLibraryController() *LibraryController {
    return &LibraryController{library: services.NewLibrary()}
}

func (c *LibraryController) AddBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID: ")
    idStr, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idStr))

    fmt.Print("Enter book title: ")
    title, _ := reader.ReadString('\n')

    fmt.Print("Enter book author: ")
    author, _ := reader.ReadString('\n')

    book := models.Book{
        ID:     id,
        Title:  strings.TrimSpace(title),
        Author: strings.TrimSpace(author),
        Status: "Available",
    }

    c.library.AddBook(book)
    fmt.Println("Book added successfully.")
}

func (c *LibraryController) RemoveBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID to remove: ")
    idStr, _ := reader.ReadString('\n')
    id, _ := strconv.Atoi(strings.TrimSpace(idStr))

    c.library.RemoveBook(id)
    fmt.Println("Book removed successfully.")
}

func (c *LibraryController) BorrowBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID to borrow: ")
    bookIDStr, _ := reader.ReadString('\n')
    bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

    fmt.Print("Enter member ID: ")
    memberIDStr, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

    err := c.library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully.")
    }
}

func (c *LibraryController) ReturnBook(reader *bufio.Reader) {
    fmt.Print("Enter book ID to return: ")
    bookIDStr, _ := reader.ReadString('\n')
    bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

    fmt.Print("Enter member ID: ")
    memberIDStr, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

    err := c.library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully.")
    }
}

func (c *LibraryController) ListAvailableBooks() {
    books := c.library.ListAvailableBooks()
    if len(books) == 0 {
        fmt.Println("No available books.")
    } else {
        fmt.Println("Available Books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}

func (c *LibraryController) ListBorrowedBooks(reader *bufio.Reader) {
    fmt.Print("Enter member ID: ")
    memberIDStr, _ := reader.ReadString('\n')
    memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

    books := c.library.ListBorrowedBooks(memberID)
    if len(books) == 0 {
        fmt.Println("No borrowed books.")
    } else {
        fmt.Println("Borrowed Books:")
        for _, book := range books {
            fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
        }
    }
}
