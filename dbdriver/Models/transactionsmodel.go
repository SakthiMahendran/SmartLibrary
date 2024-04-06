package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	StudentPtr *primitive.ObjectID `bson:"student_ptr"`
	BookPtr    *primitive.ObjectID `bson:"book_ptr"`
	DueDate    time.Time           `bson:"due_date"`
	BorrowDate time.Time           `bson:"borrow_date"`
	ReturnDate time.Time           `bson:"return_date"`
}
