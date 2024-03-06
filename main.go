package main

import (
	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controller"
	"github.com/SakthiMahendran/SmartLibrary/server"
)

func main() {
	dbdriver.Nothing()
	controller.Nothing()

	smartServer := server.NewServer()
	smartServer.Start("8080")
}
