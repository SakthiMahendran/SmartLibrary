package controllers

import (
	"context"
	"time"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentController struct {
	db             *mongo.Database
	collectionName string
}

func NewStudentController(client *mongo.Client) *StudentController {
	return &StudentController{
		db:             client.Database(dbName),
		collectionName: studentCollectionName,
	}
}

func (sc *StudentController) Authstu(RegNo string, dob time.Time) bool {
	filter := bson.M{
		"reg_no": RegNo,
		"dob":    dob,
	}
	var student models.Student
	err := sc.db.Collection(sc.collectionName).FindOne(context.Background(), filter).Decode(&student)
	if err == nil {
		return true
	} else {
		return false
	}
}

func (sc *StudentController) CreateStudent(student *models.Student) error {
	_, err := sc.db.Collection(sc.collectionName).InsertOne(context.Background(), student)
	if err != nil {
		return err
	}
	return nil
}

func (sc *StudentController) DeleteStudent(student *models.Student) error {
	filter := bson.M{"reg_no": student.RegNo}
	_, err := sc.db.Collection(sc.collectionName).DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (sc *StudentController) UpdateStudent(regNo string, updateData *models.Student) error {
	filter := bson.M{"reg_no": regNo}
	update := bson.M{"$set": updateData}
	_, err := sc.db.Collection(sc.collectionName).UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (sc *StudentController) FindStudents(filter bson.M) ([]*models.Student, error) {
	var students []*models.Student
	cur, err := sc.db.Collection(sc.collectionName).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var student models.Student
		err := cur.Decode(&student)
		if err != nil {
			return nil, err
		}
		students = append(students, &student)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
