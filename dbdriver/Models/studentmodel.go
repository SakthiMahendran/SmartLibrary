package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
    ID            primitive.ObjectID `bson:"_id,omitempty"`
    RegNo         string             `bson:"reg_no"`
    Dob           time.Time          `bson:"dob"`
    StudentName   string             `bson:"student_name"`
    Mail          string             `bson:"mail"`
    PhoneNo       string             `bson:"phone_no"`
    YearOfJoining time.Time          `bson:"year_of_joining"`
    YearOfPassing time.Time          `bson:"year_of_passing"`
    Department    string             `bson:"department"`
}

