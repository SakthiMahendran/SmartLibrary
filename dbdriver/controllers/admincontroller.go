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
	db             *mongo.Database
	collectionName string
}

func NewAdminController(client *mongo.Client) *AdminController {
	return &AdminController{
		db:             client.Database(dbName),
		collectionName: adminCollectionName,
	}
}

func (ac *AdminController) DeleteAdmin(authAdminUsername, authAdminPassword, targetUserName, targetPassword string) error {
	if !ac.Auth(authAdminUsername, authAdminPassword) {
		return errors.New("authentication failed")
	}
	if !ac.Auth(targetUserName, targetPassword) {
		return errors.New("authentication failed for target admin")
	}
	var targetAdmin models.Admin
	err := ac.db.Collection(ac.collectionName).FindOne(context.TODO(), bson.M{
		"username": targetUserName,
		"password": targetPassword,
	}).Decode(&targetAdmin)
	if err != nil {
		return err
	}
	_, err = ac.db.Collection(ac.collectionName).DeleteOne(context.TODO(), bson.M{"_id": targetAdmin.ID})
	if err != nil {
		return err
	}
	return nil
}

func (ac *AdminController) CreateAdmin(oldUserName, oldPassword, newUsername, newPassword string) error {
	if !ac.Auth(oldUserName, oldPassword) {
		return errors.New("authentication failed for old admin")
	}

	newAdmin := &models.Admin{
		Username: newUsername,
		Password: newPassword,
	}

	_, err := ac.db.Collection(ac.collectionName).InsertOne(context.TODO(), newAdmin)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AdminController) Auth(userName, passWord string) bool {
	admin := &models.Admin{}
	err := ac.db.Collection(ac.collectionName).FindOne(context.TODO(), bson.M{
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
