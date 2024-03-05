package models

import "time"

type Student struct {
    Timestamp      time.Time `bson:"timestamp"`
    Name           string    `bson:"name"`
    RegNo          string    `bson:"reg_no"`
    RollNo         int       `bson:"roll_no"`
    DOB            time.Time `bson:"dob"`
    Department     string    `bson:"stu_department"`
    MailID         string    `bson:"mail_id"`
    PhoneNumber    string    `bson:"phone_no"`
    YearOfJoining  int       `bson:"year_of_joining"`
    YearOfPassing  int       `bson:"year_of_passing"`
    DateOfBorrow   time.Time `bson:"date_of_borrow"`
    DueDate        time.Time `bson:"due_date"`
}
