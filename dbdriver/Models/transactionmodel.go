package models

import "time"

type Transaction struct {
    BookID              string    `bson:"book_id"`
    BookName            string    `bson:"book_name"`
    BookDepartment      string    `bson:"book_department"`
    StudentName         string    `bson:"stud_name"`
    RegNo               string    `bson:"reg_no"`
    BorrowTimestamp     time.Time `bson:"timestamp_of_borrow"`
    DueDate             time.Time `bson:"due_date"`
    MailID              string    `bson:"mail_id"`
    StudentDepartment   string    `bson:"stu_department"`
    ReturnTimestamp     time.Time `bson:"timestamp_of_return"`
    NumOfNotifications int       `bson:"no_of_notifications"`
    BookStatus          string    `bson:"book_status"`
}
