package dbdriver

import models "github.com/SakthiMahendran/SmartLibrary/dbdriver/Models"

type DBDriver interface {
	Connect() error
	Disconnect() error
	InsertAdmin(admin *models.Admin) error
	FindAdminByID(id string) (*models.Admin, error)
	InsertBook(book *models.Book) error
	FindBookByID(id string) (*models.Book, error)
	InsertStudent(student *models.Book) error
	FindStudentByID(id string) (*models.Student, error)
	InsertTransaction(transaction *models.Transaction) error
	FindTransactionsByStudentID(id string) ([]*models.Transaction, error)
}

func Nothing() {

}
