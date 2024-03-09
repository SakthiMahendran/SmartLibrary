package controllers

import (
	"context"
	"time"

	models "github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct {
	Client *mongo.Client
}

func (tc *TransactionController) FindStudentTransactions(studentRegNo string) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	studentsCollection := tc.Client.Database("SMLS").Collection("Students")
	transactionsCollection := tc.Client.Database("SMLS").Collection("Transaction")

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

	var transactions []models.Transaction
	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tc *TransactionController) FindBookTransactions(bookID string) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	booksCollection := tc.Client.Database("SMLS").Collection("Books")
	transactionsCollection := tc.Client.Database("SMLS").Collection("Transaction")

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

	var transactions []models.Transaction
	if err := cursor.All(ctx, &transactions); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tc *TransactionController) FindDuedTransactions() ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transactionsCollection := tc.Client.Database("SMLS").Collection("Transaction")

	filter := bson.M{
		"due_date":    bson.M{"$lt": time.Now()},
		"return_date": time.Time{},
	}

	cursor, err := transactionsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var transactions []models.Transaction
	err = cursor.All(ctx, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
