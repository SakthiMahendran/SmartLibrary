package server

import (
	"context"
	"net/http"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controllers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app             *fiber.App
	adminController *controllers.AdminController
}

func NewServer() *Server {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	return &Server{
		app:             fiber.New(),
		adminController: controllers.NewAdminController(client, "SmartLibrary", "Admin"),
	}
}

func (s *Server) Start(port string) error {
	s.setupAdminRoutes()
	return s.app.Listen(":" + port)
}

func (s *Server) setupAdminRoutes() {
	adminRouter := s.app.Group("admin/")

	adminRouter.Delete("/", func(c *fiber.Ctx) error {
		queries := c.Queries()

		if queries["operation"] == "deleteadmin" {
			err := s.adminController.DeleteAdmin(queries["authAdminUserName"], queries["authAdminPassWord"], queries["targetAdminUserName"], queries["targetAdminPassWord"])
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(fiber.Map{"message": "Success"})
		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
	})

	adminRouter.Put("/", func(c *fiber.Ctx) error {
		queries := c.Queries()

		if queries["operation"] == "createadmin" {
			err := s.adminController.CreateAdmin(queries["oldUserName"], queries["oldPassWord"], queries["newUserName"], queries["newPassWord"])
			if err != nil {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(fiber.Map{"message": "Success"})
		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
	})

	adminRouter.Post("/", func(c *fiber.Ctx) error {
		queries := c.Queries()

		if queries["operation"] == "auth" {
			if s.adminController.Auth(queries["userName"], queries["passWord"]) {
				return c.JSON(fiber.Map{"message": "Success"})
			}
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed"})
		}

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
	})
}
