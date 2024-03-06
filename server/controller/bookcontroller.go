// bookcontroller.go
package controller

import (
	"fmt"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
)

type BookController struct {
	dbDriver dbdriver.DBDriver
}

func NewBookController(dbDriver dbdriver.DBDriver) *BookController {
	return &BookController{
		dbDriver: dbDriver,
	}
}

func (bc *BookController) UpdateBookStatus(id, status string) error {
	book, err := bc.dbDriver.FindBookByID(id)
	if err != nil {
		return err
	}
	if book == nil {
		return fmt.Errorf("book not found with ID: %s", id)
	}
	book.BookStatus = status
	if err := bc.dbDriver.InsertBook(book); err != nil {
		return err
	}
	return nil
}

func (bc *BookController) AddBook(book *models.Book) error {
	return bc.dbDriver.InsertBook(book)
}

// func (bc *BookController) RemoveBook(id string) error {
// 	if err := bc.dbDriver.RemoveBook(id); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (bc *BookController) AddReturnedBook(bookID, studentID string) error {
	// Get the book
	book, err := bc.dbDriver.FindBookByID(bookID)
	if err != nil {
		return err
	}
	if book == nil {
		return fmt.Errorf("book not found with ID: %s", bookID)
	}

	// Update book status to available
	book.BookStatus = "Available"
	if err := bc.dbDriver.InsertBook(book); err != nil {
		return err
	}

	// Create a new transaction for book return
	transaction := &models.Transaction{
		BookID:             bookID,
		BookName:           book.BookName,
		BookDepartment:     book.Department,
		StudentName:        "", // Since the book is returned, no student name
		RegNo:              studentID,
		BorrowTimestamp:    time.Time{}, // Zero value, as it's returned
		DueDate:            time.Time{}, // Zero value, as it's returned
		ReturnTimestamp:    time.Now(),
		NumOfNotifications: 0, // Reset notifications for returned book
		BookStatus:         "Available",
	}

	if err := bc.dbDriver.InsertTransaction(transaction); err != nil {
		return err
	}

	return nil
}

// func (bc *BookController) RemoveBorrowedBook(bookID, studentID string) error {
// 	if err := bc.dbDriver.RemoveBorrowedBook(bookID, studentID); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (bc *BookController) UpdateBookRemarksBorrowed(bookID, remarks string) error {
// 	if err := bc.dbDriver.UpdateBookRemarksBorrowed(bookID, remarks); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (bc *BookController) UpdateBookRemarksReturned(bookID, remarks string) error {
// 	if err := bc.dbDriver.UpdateBookRemarksReturned(bookID, remarks); err != nil {
// 		return err
// 	}
// 	return nil
// }
