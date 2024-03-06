package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Username string             `bson:"username"`
    Password string             `bson:"password"`
}
