package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type transactionsRouteParams struct {
	Operation string `json:"operation"`
	RegNo     string `json:"regNo"`
	BookID    string `json:"bookID"`
}

func (s *Server) setupTransactionsRoutes() {
	transactionsRouter := s.app.Group("/api/transactions")

	transactionsRouter.Get("/student", s.findStudentTransactionsHandler)
	transactionsRouter.Get("/book", s.findBookTransactionsHandler)
	transactionsRouter.Get("/due", s.findDueTransactionsHandler)
}

func (s *Server) findStudentTransactionsHandler(c *fiber.Ctx) error {
	var params transactionsRouteParams

	if err := c.QueryParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if params.Operation == "findStudentTransactions" {
		transactions, err := s.transactionsController.FindStudentTransactions(params.RegNo)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(transactions)
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}

func (s *Server) findBookTransactionsHandler(c *fiber.Ctx) error {
	var params transactionsRouteParams

	if err := c.QueryParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if params.Operation == "findBookTransactions" {
		transactions, err := s.transactionsController.FindBookTransactions(params.BookID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(transactions)
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}

func (s *Server) findDueTransactionsHandler(c *fiber.Ctx) error {
	var params transactionsRouteParams

	if err := c.QueryParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if params.Operation == "findDueTransactions" {
		transactions, err := s.transactionsController.FindDuedTransactions()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(transactions)
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}
