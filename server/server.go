package server

import (
	"context"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controllers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app               *fiber.App
	adminController   *controllers.AdminController
	studentController *controllers.StudentController
}

func NewServer() *Server {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	return &Server{

		app:               fiber.New(),
		adminController:   controllers.NewAdminController(client),
		studentController: controllers.NewStudentController(client),
	}
}

func (s *Server) Start(port string) error {
	s.setupAdminRoutes()
	s.setupStudentRoutes()
	s.setupPageRoutes()

	return s.app.Listen(":" + port)
}
