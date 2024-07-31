# Library Management System Documentation

## Overview
This is a simple console-based library management system implemented in Go. It demonstrates the use of structs, interfaces, and other Go functionalities such as methods, slices, and maps.

## Features
- Add a new book
- Remove an existing book
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Folder Structure
library_management/
- ├── main.go
- ├── controllers/
- │ └── library_controller.go
- ├── models/
- │ └── book.go
- │ └── member.go
- ├── services/
- │ └── library_service.go
- ├── docs/
- │ └── documentation.md
- └── go.mod

## Models
### Book
Fields:
- ID (int)
- Title (string)
- Author (string)
- Status (string) // can be "Available" or "Borrowed"

### Member
Fields:
- ID (int)
- Name (string)
- BorrowedBooks ([]Book) // a slice to hold borrowed books

## Interfaces
### LibraryManager
Methods:
- AddBook(book Book)
- RemoveBook(bookID int)
- BorrowBook(bookID int, memberID int) error
- ReturnBook(bookID int, memberID int) error
- ListAvailableBooks() []Book
- ListBorrowedBooks(memberID int) []Book

## Implementation
The `LibraryManager` interface is implemented in the `Library` struct. The `Library` struct has fields to store all books (using a map with book ID as the key) and members (using a map with member ID as the key).

## Running the Application
1. Navigate to the `library_management` directory.
2. Run `go run main.go` to start the console-based application.
3. Follow the on-screen instructions to interact with the library management system.

