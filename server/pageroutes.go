package server

import (
	"github.com/gofiber/fiber/v2"
)

var basePath = "Frontend/Login-form/Sachin-boot/hope-ui-html-2.0%202/html/dashboard/"

func (s *Server) setupPageRoutes() {
	s.app.Static("/assets", "Frontend/Login-form/Sachin-boot/hope-ui-html-2.0 2/html/assets")

	pageRoutes := s.app.Group("/")

	pageRoutes.Get("/", s.lockScreenHandler)
	pageRoutes.Get("/adminlogin", s.adminLoginHandler)
	pageRoutes.Get("/studentlogin", s.studentLoginHandler)
	pageRoutes.Get("/dashboard", s.dashboardHandler)
}

func (s *Server) lockScreenHandler(c *fiber.Ctx) error {
	return c.SendFile(basePath + "auth/lock-screen.html")
}

func (s *Server) adminLoginHandler(c *fiber.Ctx) error {
	return c.SendFile(basePath + "auth/admin-login.html")
}

func (s *Server) studentLoginHandler(c *fiber.Ctx) error {
	return c.SendFile(basePath + "auth/student-login.html")
}

func (s *Server) dashboardHandler(c *fiber.Ctx) error {
	return c.SendFile(basePath + "index.html")
}
