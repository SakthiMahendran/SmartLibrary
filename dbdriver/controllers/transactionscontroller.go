package controllers

import (
	"context"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionsController struct {
	Client *mongo.Client
}

func NewTransactionsController(client *mongo.Client) *TransactionsController {
	return &TransactionsController{
		Client: client,
	}
}

func (tc *TransactionsController) FindStudentTransactions(studentRegNo string) ([]models.Transactions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	studentsCollection := tc.Client.Database(dbName).Collection(studentCollectionName)
	transactionsCollection := tc.Client.Database(dbName).Collection(transactionsCollectionName)

	var student models.Student
	err := studentsCollection.FindOne(ctx, bson.M{"reg_no": studentRegNo}).Decode(&student)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"student_ptr": student.ID}
	cursor, err := transactionsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []models.Transactions
	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tc *TransactionsController) FindBookTransactions(bookID string) ([]models.Transactions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	booksCollection := tc.Client.Database(dbName).Collection(bookCollectionName)
	transactionsCollection := tc.Client.Database(dbName).Collection(transactionsCollectionName)

	var book models.Book
	err := booksCollection.FindOne(ctx, bson.M{"book_id": bookID}).Decode(&book)
	if err != nil {
		return nil, err
	}

	cursor, err := transactionsCollection.Find(ctx, bson.M{"book_ptr": book.ID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []models.Transactions
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tc *TransactionsController) FindDuedTransactions() ([]models.Transactions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transactionsCollection := tc.Client.Database(dbName).Collection(transactionsCollectionName)

	filter := bson.M{
		"due_date":    bson.M{"$lt": time.Now()},
		"return_date": time.Time{},
	}

	cursor, err := transactionsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []models.Transactions
	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
