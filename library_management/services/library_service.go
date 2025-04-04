package services

import (
	"errors"
	"library_management/models"
	"slices"
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
	books         []models.Book
	members       []models.Member
	currentMember *models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   []models.Book{},
		members: []models.Member{},
	}
}

// Member
func (s *Library) AddMember(member models.Member) {
	s.members = append(s.members, member)
	s.currentMember = &member
}
func (s *Library) GetMembers() []models.Member {
	return s.members
}
func (s *Library) GetCurrentUser() *models.Member {
	return s.currentMember
}

// Book
func (s *Library) AddBook(title, author string) error {
	if title == "" || author == "" {
		return errors.New("title and author are required")
	}
	newBook := (models.Book{Title: title, Author: author, Status: "available", ID: len(s.books)})
	s.books = append(s.books, newBook)
	return nil
}
func (s *Library) RemoveBook(bookId int) error {

	n := len(s.books)
	for i, book := range s.books {
		if book.ID == bookId {

			s.books[n-1], s.books[i] = s.books[i], s.books[n-1]
			s.books = s.books[:n-1]

			return nil
		}
	}

	return errors.New("book not found")

}
func (s *Library) ReturnBook(bookID int, memberID int) error {
	var member *models.Member
	memberFound := false
	for _, m := range s.members {
		if m.ID == memberID {
			member = &m
			memberFound = true
			break
		}
	}
	if !memberFound {
		return errors.New("member not found")
	}

	for i, book := range member.BorrowedBooks {
		if book.ID == bookID {
			book.Status = "available"

			member.BorrowedBooks = slices.Delete(member.BorrowedBooks, i, i+1)
			s.currentMember = member
			return nil
		}
	}

	return errors.New("book not found in borrowed list")
}

func (s *Library) BorrowBook(bookID int, memberID int) error {
	var member *models.Member
	for i := range s.members {
		if s.members[i].ID == memberID {
			member = &s.members[i]
			break
		}
	}
	if member == nil {
		return errors.New("member not found")
	}

	for i := range s.books {
		if s.books[i].ID == bookID {
			if s.books[i].Status == "borrowed" {
				return errors.New("book is already borrowed")
			}

			s.books[i].Status = "borrowed"

			member.BorrowedBooks = append(member.BorrowedBooks, &s.books[i])

			s.currentMember = member

			return nil
		}
	}
	return errors.New("book not found")
}

func (s *Library) ListAvailableBooks() []models.Book {
	availableBooks := []models.Book{}
	for _, book := range s.books {
		if book.Status == "available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}
func (s *Library) ListBorrowedBooks(memberID int) []models.Book {
	borrowedBooks := []models.Book{}
	for _, book := range s.books {
		if book.Status == "borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	return borrowedBooks
}
