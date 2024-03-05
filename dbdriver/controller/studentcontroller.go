// studentcontroller.go
package controller

import (
	"errors"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
)

type StudentController struct {
	dbDriver dbdriver.DBDriver
}

func NewStudentController(dbDriver dbdriver.DBDriver) *StudentController {
	return &StudentController{
		dbDriver: dbDriver,
	}
}

func (sc *StudentController) AuthenticateStudent(regNo, dob string) (*models.Student, error) {
	student, err := sc.dbDriver.FindStudentByID(regNo)
	if err != nil {
		return nil, err
	}

	if student == nil {
		return nil, errors.New("student not found")
	}

	if student.DOB != dob {
		return nil, errors.New("invalid date of birth")
	}

	return student, nil
}

func (sc *StudentController) BorrowBook(regNo, bookID string, dueDate time.Time) error {
	book, err := sc.dbDriver.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	if book.BookStatus != "available" {
		return errors.New("book is not available for borrowing")
	}

	// Update book status and remarks
	book.BookStatus = "borrowed"
	book.Remarks = "Borrowed by student with RegNo: " + regNo

	if err := sc.dbDriver.InsertTransaction(&models.Transaction{
		BookID:             bookID,
		BookName:           book.BookName,
		BookDepartment:     book.Department,
		StudentName:        regNo,
		RegNo:              regNo,
		BorrowTimestamp:    time.Now(),
		DueDate:            dueDate,
		MailID:             "",
		StudentDepartment:  book.Department,
		ReturnTimestamp:    time.Time{},
		NumOfNotifications: 0,
		BookStatus:         "borrowed",
	}); err != nil {
		return err
	}

	return sc.dbDriver.InsertBook(book)
}

func (sc *StudentController) ReturnBook(regNo, bookID string) error {
	book, err := sc.dbDriver.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if book == nil {
		return errors.New("book not found")
	}

	// Update book status and remarks
	book.BookStatus = "available"
	book.Remarks = "Returned by student with RegNo: " + regNo

	if err := sc.dbDriver.InsertTransaction(&models.Transaction{
		BookID:            bookID,
		BookName:          book.BookName,
		BookDepartment:    book.Department,
		StudentName:       regNo, // Assuming student name is same as regNo
		RegNo:             regNo,
		ReturnTimestamp:   time.Now(),
		StudentDepartment: book.Department,
		BookStatus:        "available",
	}); err != nil {
		return err
	}

	return sc.dbDriver.InsertBook(book)
}
