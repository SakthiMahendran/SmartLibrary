package controllers

import (
	"context"
	"time"

	models "github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct {
	Client *mongo.Client
}

func (tc *TransactionController) FindStudentTransactions(studentID primitive.ObjectID) ([]models.Transaction, error) {
	// Find transactions for a given student
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transactionsCollection := tc.Client.Database("library").Collection("transactions")

	filter := bson.M{"student_id": studentID}
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

func (tc *TransactionController) FindBookTransactions(bookID primitive.ObjectID) ([]models.Transaction, error) {
	// Find transactions for a given book
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transactionsCollection := tc.Client.Database("library").Collection("transactions")

	filter := bson.M{"book_id": bookID}
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


func (tc *TransactionController) FindDuedTransactions(dueDate time.Time) ([]models.Transaction, error) {
	// Find transactions with due date before the given date
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	transactionsCollection := tc.Client.Database("library").Collection("transactions")

	filter := bson.M{"due_date": bson.M{"$lt": dueDate}}
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
