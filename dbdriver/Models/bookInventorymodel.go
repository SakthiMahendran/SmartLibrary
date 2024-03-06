package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookInventory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	BookName  string             `bson:"book_name"`
	Author    string             `bson:"author"`
	BookDept  string             `bson:"book_dept"`
	AddedDate time.Time          `bson:"added_date"`
	Count     int                `bson:"count"`
}
