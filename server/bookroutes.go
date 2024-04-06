package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BookData struct {
	BookID string `json:"bookID"`
}

func (s *Server) setupBookRoutes() {
	bookRouter := s.app.Group("/api/books")

	bookRouter.Get("/borrowed", s.isBookBorrowedHandler)
}

func (s *Server) isBookBorrowedHandler(c *fiber.Ctx) error {
	bookID := c.Query("bookID")

	isBorrowed, err := s.bookController.IsBorrowed(bookID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"isBorrowed": isBorrowed})
}
