# Library Management System

## Project Overview

This project is a console-based Library Management System implemented in Go. The system provides functionality to manage books, members, and their borrowing status. It allows adding/removing books, borrowing/returning books, and listing available or borrowed books.

## Features:

- Add a new book to the library
- Remove a book from the library
- Borrow a book
- Return a borrowed book
- List available books
- List borrowed books by a specific member

---

## System Requirements

### Minimum Requirements:

- Go 1.18 or higher
- A terminal/console for running the application
- Operating System: Cross-platform (Windows, macOS, Linux)

---

## Folder Structure

library_management/
├── main.go # Entry point of the application
├── controllers/ # Handles user input and services
│ └── library_controller.go # Handles user interactions (e.g., sign-in, book borrowing)
├── models/ # Contains data models like Book and Member
│ └── book.go # Definition of Book struct
│ └── member.go # Definition of Member struct
├── services/ # Contains business logic
│ └── library_service.go # Implements LibraryManager interface and its methods
├── docs/ # Documentation files
│ └── documentation.md # System documentation
| └── go.mod # Go module and dependencies

---

## Setup Instructions

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/library-management.git
   ```

2. Navigate to the project directory:

   ```sh
   cd library-management
   ```

3. Ensure Go is installed by checking the version:

   ```sh
   go version
   ```

4. Run the application:
   ```sh
   go run main.go
   ```

This will launch the console-based library management system.

---

## Features and Functionality

### Add a New Book

- **Command**: `1. Add Book`
- The user is prompted to enter a book title and author, which are then added to the library.

### Remove a Book

- **Command**: `2. Remove Book`
- The user is asked for the book ID. If the book exists, it will be removed from the library.

### Borrow a Book

- **Command**: `3. Borrow Book`
- The user can borrow a book if it's available. The book's status will change to "borrowed" and will be added to the user's borrowed list.

### Return a Book

- **Command**: `4. Return Book`
- The user can return a borrowed book, which will change its status back to "available."

### List Available Books

- **Command**: `5. List Available Books`
- Displays a list of books that are available to borrow.

### List Borrowed Books

- **Command**: `6. List Borrowed Books`
- Displays a list of books that have been borrowed by the current user.

---

## LibraryManager Interface Methods

### AddBook

```go
func (s *Library) AddBook(title, author string) error
```

    - Adds a new book to the library with the given title and author.

---

### RemoveBook

```go
func (s *Library) RemoveBook(bookID int) error
```

    - Removes a book from the library by its ID.

### BorrowBook

```go
func (s *Library) BorrowBook(bookID int, memberID int) error
```

    - Allows a member to borrow a book by its ID.

### ReturnBook

```go
func (s *Library) ReturnBook(bookID int, memberID int) error
```

    - Allows a member to return a borrowed book by its ID.

### ListAvailableBooks

```go
func (s *Library) ListAvailableBooks() []Book
```

    - Lists all available books in the library.

### ListBorrowedBooks

```go
func (s *Library) ListBorrowedBooks(memberID int) []Book
```

    - Lists all books borrowed by a specific member.

## Error Handling

The system performs error handling for the following scenarios:

### Book Not Found:

    - If a book is not found during borrow or return operations, an error is returned.

### Invalid Input:

    - If the user provides invalid input (e.g., empty title or author), the system prompts the user again.

### Member Not Found:

    - If a member is not found, the system returns an error.

# Contributors

    - Yoseph Alemu - https://github.com/Yosef64
