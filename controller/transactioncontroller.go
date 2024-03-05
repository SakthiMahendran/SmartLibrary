// transactioncontroller.go
package controllers

import (
    "errors"
    "github.com/yourusername/yourproject/models"
    "time"
)

type TransactionController struct {
    dbDriver models.DBDriver
}

func NewTransactionController(dbDriver models.DBDriver) *TransactionController {
    return &TransactionController{
        dbDriver: dbDriver,
    }
}

func (tc *TransactionController) UpdateTransactionDetails(bookID, studentRegNo string, notifications int) error {
    book, err := tc.dbDriver.FindBookByID(bookID)
    if err != nil {
        return err
    }

    if book == nil {
        return errors.New("book not found")
    }

    student, err := tc.dbDriver.FindStudentByID(studentRegNo)
    if err != nil {
        return err
    }

    if student == nil {
        return errors.New("student not found")
    }

    // Update transaction details
    transaction := &models.Transaction{
        BookID:              bookID,
        BookName:            book.BookName,
        BookDepartment:      book.Department,
        StudentName:         student.Name,
        RegNo:               student.RegNo,
        BorrowTimestamp:     book.DateOfBorrow,
        DueDate:             book.DueDate,
        MailID:              student.MailID,
        StudentDepartment:   student.Department,
        ReturnTimestamp:     time.Now(),
        NumOfNotifications: notifications,
        BookStatus:          book.BookStatus,
    }

    return tc.dbDriver.InsertTransaction(transaction)
}
