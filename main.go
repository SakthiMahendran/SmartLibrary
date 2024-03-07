package main

import "github.com/SakthiMahendran/SmartLibrary/server"

func main() {
	// dbdriver.Nothing()
	// controller.Nothing()

	smartServer := server.NewServer()
	smartServer.Start("8080")

	//client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	//adminController := controllers.NewAdminController(client, "SMLS", "Admin")

}
