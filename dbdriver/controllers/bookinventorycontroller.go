package controllers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookInventoryController struct {
	client                 *mongo.Client
	inventoryDatabase      *mongo.Database
	booksCollection        *mongo.Collection
	inventoryCollection    *mongo.Collection
	transactionsCollection *mongo.Collection
	studentsCollection     *mongo.Collection
}

func NewBookInventoryController(client *mongo.Client) *BookInventoryController {
	db := client.Database(dbName)
	return &BookInventoryController{
		client:                 client,
		inventoryDatabase:      db,
		booksCollection:        db.Collection(bookCollectionName),
		inventoryCollection:    db.Collection(bookInventoryCollectionName),
		transactionsCollection: db.Collection(transactionsCollectionName),
		studentsCollection:     db.Collection(studentCollectionName),
	}
}

func (bc *BookInventoryController) AddBook(bookID, bookName, author, bookDept string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"book_id": bookID}
	var existingBook models.Book
	err := bc.booksCollection.FindOne(ctx, filter).Decode(&existingBook)
	if err == nil {
		return fmt.Sprintf("Book with ID %s already exists", bookID), nil
	}

	filter = bson.M{"book_name": bookName}
	var existingInventory models.BookInventory
	err = bc.inventoryCollection.FindOne(ctx, filter).Decode(&existingInventory)
	if err == nil {
		update := bson.M{"$inc": bson.M{"count": 1}}
		_, err := bc.inventoryCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return "", err
		}
	} else if err == mongo.ErrNoDocuments {
		newInventory := models.BookInventory{
			BookName:  bookName,
			Author:    author,
			BookDept:  bookDept,
			AddedDate: time.Now(),
			Count:     1,
		}
		_, err := bc.inventoryCollection.InsertOne(ctx, newInventory)
		if err != nil {
			return "", err
		}
	} else {
		return "", err
	}

	newBook := models.Book{
		BookID:     bookID,
		BookStatus: true,
		BookInventoryPtr: &models.BookInventory{
			BookName:  bookName,
			Author:    author,
			BookDept:  bookDept,
			AddedDate: time.Now(),
			Count:     1,
		},
	}
	_, err = bc.booksCollection.InsertOne(ctx, newBook)
	if err != nil {
		return "", err
	}

	_, err = bc.booksCollection.UpdateMany(ctx, bson.M{}, bson.M{"$set": bson.M{"book_inventory_ptr": newBook.BookInventoryPtr}})
	if err != nil {
		return "", err
	}

	return "Book added successfully", nil
}

func (bc *BookInventoryController) UpdateBook(presentData, toBeUpdatedData *models.BookInventory) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"book_name": presentData.BookName}
	var existingInventory models.BookInventory
	err := bc.inventoryCollection.FindOne(ctx, filter).Decode(&existingInventory)
	if err != nil {
		return err
	}
	updateFields := bson.M{}
	if toBeUpdatedData.BookName != "" && toBeUpdatedData.BookName != existingInventory.BookName {
		updateFields["book_name"] = toBeUpdatedData.BookName
	}
	if toBeUpdatedData.Author != "" && toBeUpdatedData.Author != existingInventory.Author {
		updateFields["author"] = toBeUpdatedData.Author
	}
	if toBeUpdatedData.BookDept != "" && toBeUpdatedData.BookDept != existingInventory.BookDept {
		updateFields["book_dept"] = toBeUpdatedData.BookDept
	}
	if !toBeUpdatedData.AddedDate.IsZero() {
		updateFields["added_date"] = toBeUpdatedData.AddedDate
	}
	if toBeUpdatedData.Count != existingInventory.Count {
		updateFields["count"] = toBeUpdatedData.Count
	}
	update := bson.M{"$set": updateFields}
	_, err = bc.inventoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	var updatedInventory models.BookInventory
	err = bc.inventoryCollection.FindOne(ctx, filter).Decode(&updatedInventory)
	if err != nil {
		return err
	}

	filter = bson.M{"book_inventory_ptr.book_name": existingInventory.BookName}
	_, err = bc.booksCollection.UpdateMany(ctx, filter, bson.M{"$set": bson.M{"book_inventory_ptr": updatedInventory}})
	if err != nil {
		return err
	}

	return nil
}

func (bc *BookInventoryController) DeleteBook(bookID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"book_id": bookID}
	var book models.Book
	err := bc.booksCollection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		fmt.Println("Book not present")
		return err
	}

	bookName := book.BookInventoryPtr.BookName

	_, err = bc.booksCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	filter = bson.M{"book_name": bookName}
	update := bson.M{"$inc": bson.M{"count": -1}}
	_, err = bc.inventoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	var updatedInventory models.BookInventory
	err = bc.inventoryCollection.FindOne(ctx, filter).Decode(&updatedInventory)
	if err != nil {
		return err
	}

	filter = bson.M{"book_inventory_ptr.book_name": bookName}
	_, err = bc.booksCollection.UpdateMany(ctx, filter, bson.M{"$set": bson.M{"book_inventory_ptr": updatedInventory}})
	if err != nil {
		return err
	}

	return nil
}

func (bc *BookInventoryController) GetBookCount(bookName string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var bookInventory models.BookInventory
	filter := bson.M{"book_name": bookName}
	err := bc.inventoryCollection.FindOne(ctx, filter).Decode(&bookInventory)
	if err != nil {
		return 0, err
	}

	return bookInventory.Count, nil
}

func (bc *BookInventoryController) GetCategoryCount(department string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := bc.inventoryCollection.Find(ctx, bson.M{"book_dept": department})
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var totalCount int
	for cursor.Next(ctx) {
		var bookInventory models.BookInventory
		if err := cursor.Decode(&bookInventory); err != nil {
			return 0, err
		}
		totalCount += bookInventory.Count
	}
	if err := cursor.Err(); err != nil {
		return 0, err
	}

	return totalCount, nil
}

func (bc *BookInventoryController) FindCategory(department string) ([]models.BookInventory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := bc.inventoryCollection.Find(ctx, bson.M{"book_dept": department})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var books []models.BookInventory
	for cursor.Next(ctx) {
		var bookInventory models.BookInventory
		if err := cursor.Decode(&bookInventory); err != nil {
			return nil, err
		}
		books = append(books, bookInventory)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, fmt.Errorf("no books found for department %s", department)
	}

	return books, nil
}

func (bc *BookInventoryController) IsAvailable(bookName string) (bool, error) {
	count, err := bc.GetBookCount(bookName)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (bc *BookInventoryController) Borrow(bookID, studentRegNo string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	booksCollection := bc.booksCollection
	inventoryCollection := bc.inventoryCollection
	transactionsCollection := bc.transactionsCollection
	studentCollection := bc.studentsCollection

	// Check if the book is available for borrowing
	filter := bson.M{"book_id": bookID, "book_status": true}
	var book models.Book
	err := booksCollection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		return errors.New("book is not available for borrowing")
	}

	// Update book status to unavailable
	bookUpdate := bson.M{"$set": bson.M{"book_status": false}}
	_, err = booksCollection.UpdateOne(ctx, filter, bookUpdate)
	if err != nil {
		return err
	}

	// Decrement book count in inventory
	bookName := book.BookInventoryPtr.BookName
	inventoryFilter := bson.M{"book_name": bookName}
	inventoryUpdate := bson.M{"$inc": bson.M{"count": -1}}
	_, err = inventoryCollection.UpdateOne(ctx, inventoryFilter, inventoryUpdate)
	if err != nil {
		return err
	}

	// Retrieve student information
	var student models.Student
	err = studentCollection.FindOne(ctx, bson.M{"reg_no": studentRegNo}).Decode(&student)
	if err != nil {
		return err
	}

	// Calculate due date (15 days from now)
	dueDate := time.Now().AddDate(0, 0, 15)

	// Create transaction record
	transaction := models.Transactions{
		StudentPtr: &student.ID,
		BookPtr:    &book.ID,
		BorrowDate: time.Now(),
		DueDate:    dueDate,
	}

	_, err = transactionsCollection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (bc *BookInventoryController) Return(bookID, studentRegNo string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	booksCollection := bc.booksCollection
	inventoryCollection := bc.inventoryCollection
	transactionsCollection := bc.transactionsCollection
	studentCollection := bc.studentsCollection

	// Find the book by ID
	var book models.Book
	err := booksCollection.FindOne(ctx, bson.M{"book_id": bookID}).Decode(&book)
	if err != nil {
		return err
	}

	// Find the student by registration number
	var student models.Student
	err = studentCollection.FindOne(ctx, bson.M{"reg_no": studentRegNo}).Decode(&student)
	if err != nil {
		return err
	}

	// Find the transaction record
	var transaction models.Transactions
	err = transactionsCollection.FindOne(ctx, bson.M{
		"book_ptr":    book.ID,
		"student_ptr": student.ID,
	}).Decode(&transaction)
	if err != nil {
		return err
	}

	// Update the transaction record with return date
	_, err = transactionsCollection.UpdateOne(ctx,
		bson.M{"_id": transaction.ID},
		bson.M{"$set": bson.M{"return_date": time.Now()}},
	)
	if err != nil {
		return err
	}

	// Update book status to available
	bookUpdate := bson.M{"$set": bson.M{"book_status": true}}
	_, err = booksCollection.UpdateOne(ctx,
		bson.M{"_id": book.ID},
		bookUpdate,
	)
	if err != nil {
		return err
	}

	// Increment book count in inventory
	bookName := book.BookInventoryPtr.BookName
	inventoryFilter := bson.M{"book_name": bookName}
	inventoryUpdate := bson.M{"$inc": bson.M{"count": 1}}
	_, err = inventoryCollection.UpdateOne(ctx,
		inventoryFilter,
		inventoryUpdate,
	)
	if err != nil {
		return err
	}

	return nil
}
