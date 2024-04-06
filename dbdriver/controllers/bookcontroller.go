package controllers

import (
	"context"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookController struct {
	collection *mongo.Collection
}

func NewBookController(client *mongo.Client) *BookController {
	return &BookController{
		collection: client.Database(dbName).Collection("Books"),
	}
}

func (bc *BookController) IsBorrowed(bookID string) (bool, error) {
	filter := bson.M{"book_id": bookID}

	ctx := context.TODO()

	var book models.Book
	err := bc.collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Book ID not found
			return false, nil
		}
		// Other errors
		return false, err
	}

	// Book ID found, check BookStatus
	return !book.BookStatus, nil
}
