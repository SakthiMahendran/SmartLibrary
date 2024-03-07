package controllers

import (
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
)

type TransctionController struct {
}

func (tc *TransctionController) FindStudentTransactions(student *models.Student) *[]models.Transaction {

}

func (tc *TransctionController) FindBookTransactions(student *models.Book) *[]models.Transaction {

}

func (tc *TransctionController) FindDuedTransactions(dueDate time.Time) *[]models.Transaction {

}
