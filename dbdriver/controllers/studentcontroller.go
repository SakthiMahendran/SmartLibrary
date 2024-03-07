package controllers

import (
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentController struct {
	db                    *mongo.Database
	studentCollectionName string
}

func NewStudentController(client *mongo.Client, dbName, studentCollectionName string) *StudentController {
	return &StudentController{
		db:                    client.Database(dbName),
		studentCollectionName: studentCollectionName,
	}
}

func (sc *StudentController) Authstu(RegNo, dob string) bool {

	return true
}

func (sc *StudentController) CreateStudent(Admin *models.Admin, RegNo, student_name, mail, phone_no, year_of_joining, year_of_passing, department string) error {

	return nil
}
