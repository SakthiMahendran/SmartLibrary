package server

import (
	"net/http"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BookID           string             `json:"book_id" bson:"book_id"`
	BookStatus       bool               `json:"book_status" bson:"book_status"`
	BookInventoryPtr *BookInventory     `json:"book_inventory_ptr" bson:"book_inventory_ptr"`
}

type BookInventory struct {
	BookName  string    `json:"book_name" bson:"book_name"`
	Author    string    `json:"author" bson:"author"`
	BookDept  string    `json:"book_dept" bson:"book_dept"`
	AddedDate time.Time `json:"added_date" bson:"added_date"`
	Count     int       `json:"count" bson:"count"`
}

type BorrowData struct {
	BookID       string `json:"book_id"`
	StudentRegNo string `json:"student_reg_no"`
}

type ReturnData struct {
	BookID       string `json:"book_id"`
	StudentRegNo string `json:"student_reg_no"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (s *Server) setupBookInventoryRoutes() {
	bookInventoryRouter := s.app.Group("/api/bookinventory")

	// Endpoint to add a book
	bookInventoryRouter.Post("/addbook", s.addBookHandler)

	// Endpoint to update a book
	bookInventoryRouter.Put("/updatebook", s.updateBookHandler)

	// Endpoint to delete a book
	bookInventoryRouter.Delete("/deletebook", s.deleteBookHandler)

	// Endpoint to check availability of a book
	bookInventoryRouter.Get("/checkavailability", s.checkAvailabilityHandler)

	// Endpoint to borrow a book
	bookInventoryRouter.Post("/borrowbook", s.borrowBookHandler)

	// Endpoint to return a book
	bookInventoryRouter.Post("/returnbook", s.returnBookHandler)
}

func (s *Server) addBookHandler(c *fiber.Ctx) error {
	var bookData Book
	if err := c.BodyParser(&bookData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	// Call the AddBook method from the BookInventoryController with the book data
	message, err := s.bookInventoryController.AddBook(bookData.BookID, bookData.BookInventoryPtr.BookName, bookData.BookInventoryPtr.Author, bookData.BookInventoryPtr.BookDept)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with success message
	return c.JSON(SuccessResponse{Message: message})
}

func (s *Server) updateBookHandler(c *fiber.Ctx) error {
	var presentData, toBeUpdatedData models.BookInventory
	if err := c.BodyParser(&presentData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}
	if err := c.BodyParser(&toBeUpdatedData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	// Call the UpdateBook method from the BookInventoryController with the present and updated book data
	err := s.bookInventoryController.UpdateBook(&presentData, &toBeUpdatedData)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with success message
	return c.JSON(SuccessResponse{Message: "Book updated successfully"})
}

func (s *Server) deleteBookHandler(c *fiber.Ctx) error {
	var bookData Book
	if err := c.BodyParser(&bookData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	// Call the DeleteBook method from the BookInventoryController with the book ID
	err := s.bookInventoryController.DeleteBook(bookData.BookID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with success message
	return c.JSON(SuccessResponse{Message: "Book deleted successfully"})
}

func (s *Server) checkAvailabilityHandler(c *fiber.Ctx) error {
	bookName := c.Query("bookName")

	// Call the IsAvailable method from the BookInventoryController with the book name
	isAvailable, err := s.bookInventoryController.IsAvailable(bookName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with availability status
	return c.JSON(fiber.Map{"isAvailable": isAvailable})
}

func (s *Server) borrowBookHandler(c *fiber.Ctx) error {
	var borrowData BorrowData
	if err := c.BodyParser(&borrowData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	// Call the Borrow method from the BookInventoryController with the book ID and student registration number
	err := s.bookInventoryController.Borrow(borrowData.BookID, borrowData.StudentRegNo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with success message
	return c.JSON(SuccessResponse{Message: "Book borrowed successfully"})
}

func (s *Server) returnBookHandler(c *fiber.Ctx) error {
	var returnData ReturnData
	if err := c.BodyParser(&returnData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
	}

	// Call the Return method from the BookInventoryController with the book ID and student registration number
	err := s.bookInventoryController.Return(returnData.BookID, returnData.StudentRegNo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{Error: err.Error()})
	}

	// Respond with success message
	return c.JSON(SuccessResponse{Message: "Book returned successfully"})
}
