package models

import "time"

type Book struct {
    Timestamp      time.Time `bson:"timestamp"`
    BookID         string    `bson:"book_id"`
    BookName       string    `bson:"book_name"`
    BookStatus     string    `bson:"book_status"`
    Department     string    `bson:"book_department"`
    Remarks        string    `bson:"remarks"`
}
