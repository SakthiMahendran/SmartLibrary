package server

import (
	"net/http"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Define a struct to represent the parameters for each student operation
type studentRouteParams struct {
	Operation     string    `json:"operation"`
	RegNo         string    `json:"regNo"`
	Dob           time.Time `json:"dob"`
	StudentName   string    `json:"studentName"`
	Mail          string    `json:"mail"`
	PhoneNo       string    `json:"phoneNo"`
	YearOfJoining time.Time `json:"yearOfJoining"`
	YearOfPassing time.Time `json:"yearOfPassing"`
	Department    string    `json:"department"`
}

func (s *Server) setupStudentRoutes() {
	studentRouter := s.app.Group("api/student")

	studentRouter.Post("/", s.studentHandler)
	// Add other routes as needed
}

func (s *Server) studentHandler(c *fiber.Ctx) error {
	var params studentRouteParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	switch params.Operation {
	case "auth":
		return s.authStudentHandler(c, params)
	case "create":
		return s.createStudentHandler(c, params)
	case "delete":
		return s.deleteStudentHandler(c, params)
	case "update":
		return s.updateStudentHandler(c, params)
	case "find":
		return s.findStudentsHandler(c, params)
	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid operation"})
	}
}

func (s *Server) authStudentHandler(c *fiber.Ctx, params studentRouteParams) error {
	auth := s.studentController.Authstu(params.RegNo, params.Dob)
	if auth {
		return c.JSON(fiber.Map{"message": "Authentication successful"})
	}
	return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authentication failed"})
}

func (s *Server) createStudentHandler(c *fiber.Ctx, params studentRouteParams) error {
	student := models.Student{
		RegNo:         params.RegNo,
		Dob:           params.Dob,
		StudentName:   params.StudentName,
		Mail:          params.Mail,
		PhoneNo:       params.PhoneNo,
		YearOfJoining: params.YearOfJoining,
		YearOfPassing: params.YearOfPassing,
		Department:    params.Department,
	}

	err := s.studentController.CreateStudent(&student)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Student created successfully"})
}

func (s *Server) deleteStudentHandler(c *fiber.Ctx, params studentRouteParams) error {
	student := models.Student{
		RegNo: params.RegNo,
	}

	err := s.studentController.DeleteStudent(&student)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Student deleted successfully"})
}

func (s *Server) updateStudentHandler(c *fiber.Ctx, params studentRouteParams) error {
	student := models.Student{
		RegNo:         params.RegNo,
		Dob:           params.Dob,
		StudentName:   params.StudentName,
		Mail:          params.Mail,
		PhoneNo:       params.PhoneNo,
		YearOfJoining: params.YearOfJoining,
		YearOfPassing: params.YearOfPassing,
		Department:    params.Department,
	}

	// Assuming updateData is a map[string]interface{} containing update fields

	err := s.studentController.UpdateStudent(student.RegNo, &student)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Student updated successfully"})
}

func (s *Server) findStudentsHandler(c *fiber.Ctx, params studentRouteParams) error {

	// Construct filter based on parameters received in the request
	filter := bson.M{}
	if params.RegNo != "" {
		filter["reg_no"] = params.RegNo
	}
	// Add other filters as needed

	// Call the FindStudents method to retrieve students
	students, err := s.studentController.FindStudents(filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the list of students as JSON response
	return c.JSON(students)
}
