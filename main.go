package main

import (
	"context"
	"fmt"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	/*dbdriver.Nothing()
	// controller.Nothing()

	// smartServer := server.NewServer()
	// smartServer.Start("8080")*/

	client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	adminController := controllers.NewAdminController(client, "SMLS", "Admin")
	fmt.Printf("adminController.Auth(\"Abba\", \"uK1'Ra<5U'7W1n,\"): %v\n", adminController.Auth("Abba", "uK1'Ra<5U'7W1n,"))
}
