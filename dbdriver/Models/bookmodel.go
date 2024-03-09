package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	BookID           string             `bson:"book_id"`
	BookStatus       bool                `bson:"book_status"`
	BookInventoryPtr *BookInventory     `bson:"book_inventory_ptr"`
	
}
