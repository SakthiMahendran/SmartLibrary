package controllers

import "github.com/SakthiMahendran/SmartLibrary/dbdriver/models"

type BookInventoryController struct {
}

func (bc *BookInventoryController) AddBook(newBook *models.Book) error {

}

func (bc *BookInventoryController) UpdateBook(newBook *models.Book) error {

}

func (bc *BookInventoryController) DeleteBook(book *models.Book) error {

}

func (bc *BookInventoryController) GetBookCount(bookName string) int {

}

func (bc *BookInventoryController) GetCategoryCount(category *models.Book) int {

}

func (bc *BookInventoryController) FindCategory(category *models.Book) *[]models.BookInventory {

}

func (bc *BookInventoryController) IsAvailable(bookName string) bool {

}

func (bc *BookInventoryController) Borrow(book *models.Book, student *models.Student) error {

}

func (bc *BookInventoryController) Return(book *models.Book, student *models.Student) error {

}
