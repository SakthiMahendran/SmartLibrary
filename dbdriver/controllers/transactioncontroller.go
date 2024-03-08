package controllers

import (
	"context"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct {
	Client *mongo.Client
}

func (tc *TransactionController) Borrow(bookID primitive.ObjectID, student *models.Student) error {

}

func (tc *TransactionController) Return(bookID primitive.ObjectID, student *models.Student) error {
	
}

func (tc *TransactionController) FindStudentTransactions(student *models.Student) *[]models.Transaction {

}

func (tc *TransactionController) FindBookTransactions(student *models.Book) *[]models.Transaction {

}

func (tc *TransactionController) FindDuedTransactions(dueDate time.Time) *[]models.Transaction {

}
