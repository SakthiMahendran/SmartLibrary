package server

import (
	"context"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controllers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	app                     *fiber.App
	adminController         *controllers.AdminController
	studentController       *controllers.StudentController
	transactionsController  *controllers.TransactionsController
	bookController          *controllers.BookController
	bookInventoryController *controllers.BookInventoryController
}

func NewServer() *Server {
	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	return &Server{
		app:                     fiber.New(),
		adminController:         controllers.NewAdminController(client),
		studentController:       controllers.NewStudentController(client),
		transactionsController:  controllers.NewTransactionsController(client),
		bookController:          controllers.NewBookController(client),
		bookInventoryController: controllers.NewBookInventoryController(client),
	}
}

func (s *Server) Start(port string) error {
	s.setupAdminRoutes()
	s.setupStudentRoutes()
	s.setupPageRoutes()
	s.setupTransactionsRoutes()
	s.setupBookRoutes()
	s.setupBookInventoryRoutes()

	return s.app.Listen(":" + port)
}
