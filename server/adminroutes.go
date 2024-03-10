package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Define a struct to represent the parameters for each operation
type adminRouteParams struct {
	Operation           string `json:"operation"`
	AuthAdminUserName   string `json:"authAdminUserName"`
	AuthAdminPassword   string `json:"authAdminPassword"`
	TargetAdminUserName string `json:"targetAdminUserName"`
	TargetAdminPassword string `json:"targetAdminPassword"`
	OldUserName         string `json:"oldUserName"`
	OldPassword         string `json:"oldPassword"`
	NewUserName         string `json:"newUserName"`
	NewPassword         string `json:"newPassword"`
	UserName            string `json:"userName"`
	Password            string `json:"password"`
}

func (s *Server) setupAdminRoutes() {
	adminRouter := s.app.Group("api/admin")

	adminRouter.Delete("/", s.deleteAdminHandler)
	adminRouter.Put("/", s.createAdminHandler)
	adminRouter.Post("/", s.authHandler)
}

func (s *Server) deleteAdminHandler(c *fiber.Ctx) error {
	var params adminRouteParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if params.Operation == "deleteadmin" {
		err := s.adminController.DeleteAdmin(params.AuthAdminUserName, params.AuthAdminPassword, params.TargetAdminUserName, params.TargetAdminPassword)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "Success"})
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}

func (s *Server) createAdminHandler(c *fiber.Ctx) error {
	var params adminRouteParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("Operation: ", params.Operation)

	if params.Operation == "createadmin" {
		err := s.adminController.CreateAdmin(params.OldUserName, params.OldPassword, params.NewUserName, params.NewPassword)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"message": "Success"})
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}

func (s *Server) authHandler(c *fiber.Ctx) error {
	var params adminRouteParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println(params)

	if params.Operation == "auth" {
		if s.adminController.Auth(params.UserName, params.Password) {
			return c.JSON(fiber.Map{"message": "Success"})
		}
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed"})
	}

	return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
}
