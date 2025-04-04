package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

const (
	Blue   = "\033[34m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

type LibraryController struct {
	service *services.Library
}

func NewLibraryController(service *services.Library) *LibraryController {
	return &LibraryController{service: service}
}
func (s *LibraryController) Landing() {
	fmt.Println("=== Library Management System ===")
	fmt.Println(Blue + "\t1. Sign In")
	fmt.Println("\t2. Exit")
	fmt.Print(Yellow + "Enter your choice: " + Reset)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("Invalid input. Please try again.")
		s.Landing()
		return
	}
	switch input {
	case "1":
		s.UserLogin()

	case "2":
		fmt.Println("Manager Login")
	default:
		fmt.Println("existed")

	}
}
func (c *LibraryController) MainOption() {
	reader := bufio.NewReader(os.Stdin)

	currentMember := c.service.GetCurrentUser()
	fmt.Printf(Green+"\n---- Welcome back, %s ----"+Blue+"\n1. Add Book\n2. Remove Book\n3. Available Books\n4. Borrowed Books\n5. Borrow Book\n6. Return Book\n7. Exit"+Yellow+"\n\nEnter your choice: ", currentMember.Name)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		c.AddBook()
	case "2":
		c.RemoveBook()

	case "3":
		c.AvailableBooks()

	case "4":
		c.BorrowedBooks()

	case "5":
		c.BorrowBook()

	case "6":
		c.ReturnBook()

	case "7":
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println(Red + "Invalid choice. Please try again.")
		c.MainOption()
	}
}

func (c *LibraryController) UserLogin() {
	var name string
	fmt.Println(Green + "\n=== User Login ===")
	fmt.Print(Yellow + "Enter name : " + Reset)
	fmt.Scan(&name)
	c.AddMember(name)
	c.MainOption()

}

// Member
func (c *LibraryController) AddMember(name string) {
	totalMember := c.service.GetMembers()
	member := models.Member{
		ID:            len(totalMember),
		Name:          name,
		BorrowedBooks: []*models.Book{},
	}

	c.service.AddMember(member)
}

// Book
func (c *LibraryController) AddBook() {
	reader := bufio.NewReader(os.Stdin)
	var title, author string
	fmt.Println(Green + "\n=== Add Book ===")
	fmt.Print(Yellow + "Enter title : " + Reset)
	title, _ = reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print(Yellow + "Enter author : " + Reset)
	author, _ = reader.ReadString('\n')
	author = strings.TrimSpace(author)
	if title == "" || author == "" {
		fmt.Println(Red + "Title and author are required." + Reset)
		c.FinalOption()
		return
	}
	err := c.service.AddBook(title, author)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf(Green + "Book added successfully!\n" + Reset)
	c.FinalOption()
}

func (c *LibraryController) RemoveBook() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(Green + "\n=== Remove Book ===")
	fmt.Print(Yellow + "Enter book ID to remove: " + Reset)
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println(Red + "Invalid book ID." + Reset)
		c.FinalOption()
		return
	}
	err = c.service.RemoveBook(bookID)
	if err != nil {
		fmt.Println(Red + err.Error() + Reset)
		c.FinalOption()
		return
	}
	fmt.Println(Green + "Book removed successfully!" + Reset)
	c.FinalOption()
}
func (c *LibraryController) ReturnBook() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(Green + "\n=== Return Book ===")
	fmt.Print(Yellow + "Enter book ID to return: " + Reset)
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println(Red + "Invalid book ID." + Reset)
		c.FinalOption()
		return
	}
	user := c.service.GetCurrentUser()
	err = c.service.ReturnBook(bookID, user.ID)
	if err != nil {
		fmt.Println(Red + err.Error() + Reset)
		c.FinalOption()
		return
	}
	fmt.Println(Green + "Book returned successfully!")
	c.FinalOption()

}
func (c *LibraryController) AvailableBooks() {
	books := c.service.ListAvailableBooks()
	fmt.Println(Green + "\n=== Available Books ===" + Reset)

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)

	fmt.Fprintln(writer, "ID\tTitle\tAuthor")
	fmt.Fprintln(writer, Blue+"----\t-----------------------------\t----------------------")

	for _, book := range books {
		fmt.Fprintf(writer, "%d\t%s\t%s\n", book.ID, book.Title, book.Author)
	}

	writer.Flush()
	c.FinalOption()
}
func (c *LibraryController) BorrowedBooks() {
	books := c.service.ListBorrowedBooks(c.service.GetCurrentUser().ID)
	fmt.Println(Green + "\n=== Borrowed Books ===" + Reset)

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', 0)

	fmt.Fprintln(writer, "ID\tTitle\tAuthor")
	fmt.Fprintln(writer, Blue+"----\t-----------------------------\t----------------------")

	for _, book := range books {
		fmt.Fprintf(writer, "%d\t%s\t%s\n", book.ID, book.Title, book.Author)
	}

	writer.Flush()
	c.FinalOption()
}
func (c *LibraryController) BorrowBook() {
	fmt.Println(Green + "\n=== Borrow Book ===")
	fmt.Print(Yellow + "Enter book ID to borrow: " + Reset)
	reader := bufio.NewReader(os.Stdin)
	bookIDStr, _ := reader.ReadString('\n')
	bookIDStr = strings.TrimSpace(bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println(Red + "Invalid book ID." + Reset)
		c.FinalOption()
		return
	}
	err = c.service.BorrowBook(bookID, c.service.GetCurrentUser().ID)
	if err != nil {
		fmt.Println(Red + "Error:" + err.Error() + Reset)
		c.FinalOption()
		return
	}
	fmt.Println(Green + "Book borrowed successfully!" + Reset)
	c.FinalOption()
}

// Final Option
func (c *LibraryController) FinalOption() {
	fmt.Print(Reset + "Do you want to continue? (y/n): ")
	var input string
	fmt.Scan(&input)
	switch input {
	case "y":
		c.MainOption()
	case "n":
		fmt.Println("Exiting...")
	default:
		fmt.Println(Red + "Invalid choice. Please try again." + Reset)
		c.FinalOption()
	}
}
