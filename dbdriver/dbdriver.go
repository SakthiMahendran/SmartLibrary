package models

type DBDriver interface {
    Connect() error
    Disconnect() error
    InsertAdmin(admin *Admin) error
    FindAdminByID(id string) (*Admin, error)
    InsertBook(book *Book) error
    FindBookByID(id string) (*Book, error)
    InsertStudent(student *Student) error
    FindStudentByID(id string) (*Student, error)
    InsertTransaction(transaction *Transaction) error
    FindTransactionsByStudentID(id string) ([]*Transaction, error)
}
