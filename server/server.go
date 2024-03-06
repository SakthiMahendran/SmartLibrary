package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app   *fiber.App
	books []Book
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func NewServer() *Server {
	books := []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1"},
		{ID: 2, Title: "Book 2", Author: "Author 2"},
		{ID: 3, Title: "Book 3", Author: "Author 3"},
		{ID: 4, Title: "Book 4", Author: "Author 4"},
		{ID: 5, Title: "Book 5", Author: "Author 5"},
	}

	return &Server{
		app:   fiber.New(),
		books: books,
	}
}

func (s *Server) Start(port string) error {
	s.setupRoutes()
	return s.app.Listen(":" + port)
}

func (s *Server) setupRoutes() {
	// Authentication middleware
	s.app.Use(func(c *fiber.Ctx) error {
		// Perform authentication logic here
		// For simplicity, assume all requests are authenticated
		return c.Next()
	})

	// Get all books
	s.app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(s.books)
	})

	// Get a single book
	s.app.Get("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		bookID, _ := strconv.Atoi(id)

		for _, book := range s.books {
			if book.ID == bookID {
				return c.JSON(book)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "book not found"})
	})

	// Create a new book
	s.app.Post("/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return err
		}

		book.ID = len(s.books) + 1
		s.books = append(s.books, *book)

		return c.Status(fiber.StatusCreated).JSON(book)
	})

	// Update a book
	s.app.Put("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		bookID, _ := strconv.Atoi(id)

		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return err
		}

		for i, b := range s.books {
			if b.ID == bookID {
				s.books[i] = *book
				return c.Status(fiber.StatusOK).JSON(s.books[i])
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "book not found"})
	})

	// Delete a book
	s.app.Delete("/books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		bookID, _ := strconv.Atoi(id)

		for i, book := range s.books {
			if book.ID == bookID {
				s.books = append(s.books[:i], s.books[i+1:]...)
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "book deleted successfully"})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "book not found"})
	})
}
