package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) setupPageRoutes() {
	pageRoutes := s.app.Group("/")

	pageRoutes.Get("/adminpage", s.adminPageHandler)
	pageRoutes.Get("/studentpage", s.studentPageHandler)
}

func (s *Server) adminPageHandler(c *fiber.Ctx) error {
	return c.SendFile("server/testpages/adminpage.html")
}
func (s *Server) studentPageHandler(c *fiber.Ctx) error {
	return c.SendFile("server/testpages/studentpage.html")
}
