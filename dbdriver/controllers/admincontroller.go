package controllers

import (
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminController struct {
	db                  *mongo.Database
	adminCollectionName string
}

func NewAdminController(client *mongo.Client, dbName, adminCollectionName string) *AdminController {
	return nil
}

func (ac *AdminController) CreateAdmin(newAdmin *models.Admin) error {
	return nil
}

func (ac *AdminController) Auth(userName, passWord string) bool {
	return false
}
