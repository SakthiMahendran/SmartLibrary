package models

import "time"

type Admin struct {
    Timestamp   time.Time `bson:"timestamp"`
    Name        string    `bson:"admin_name"`
    ID          string    `bson:"admin_id"`
    Password    string    `bson:"admin_password"`
}
