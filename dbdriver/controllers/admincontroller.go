package controllers

import (
	"context"
	"errors"
	"fmt"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminController struct {
	db                  *mongo.Database
	adminCollectionName string
}

func NewAdminController(client *mongo.Client, dbName, adminCollectionName string) *AdminController {
	return &AdminController{
		db:                  client.Database(dbName),
		adminCollectionName: adminCollectionName,
	}
}
func (ac *AdminController) DeleteAdmin(authAdmin *models.Admin, targetUserName, targetPassword string) error {
	if !ac.Auth(authAdmin.Username, authAdmin.Password) {
		return errors.New("authentication failed")
	}
	if !ac.Auth(targetUserName, targetPassword) {
		return errors.New("authentication failed for target admin")
	}
	var targetAdmin models.Admin
	err := ac.db.Collection(ac.adminCollectionName).FindOne(context.TODO(), bson.M{
		"username": targetUserName,
		"password": targetPassword,
	}).Decode(&targetAdmin)
	if err != nil {
		return err
	}
	_, err = ac.db.Collection(ac.adminCollectionName).DeleteOne(context.TODO(), bson.M{"_id": targetAdmin.ID})
	if err != nil {
		return err
	}
	return nil
}

/*
	func (ac *AdminController) DeleteAdmin(oldAdmin *models.Admin, userName, passWord string) error {
		if !ac.Auth(userName, passWord) {
			return errors.New("authentication failed")
		}
		_, err := ac.db.Collection(ac.adminCollectionName).DeleteOne(context.TODO(), bson.M{"_id": oldAdmin.ID})
		if err != nil {
			return err
		}
		return nil
	}
*/
func (ac *AdminController) CreateAdmin(oldUserName, oldPassword, newUsername, newPassword string) error {
	if !ac.Auth(oldUserName, oldPassword) {
		return errors.New("authentication failed for old admin")
	}

	newAdmin := &models.Admin{
		Username: newUsername,
		Password: newPassword,
	}

	_, err := ac.db.Collection(ac.adminCollectionName).InsertOne(context.TODO(), newAdmin)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AdminController) Auth(userName, passWord string) bool {
	admin := &models.Admin{}
	err := ac.db.Collection(ac.adminCollectionName).FindOne(context.TODO(), bson.M{
		"username": userName,
	}).Decode(admin)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Admin not found")
			return false
		}
		fmt.Println("Error while finding admin:", err)
		return false
	}

	if admin.Password != passWord {
		fmt.Println("Incorrect password")
		return false
	}

	fmt.Println("Authentication successful")
	return true
}
