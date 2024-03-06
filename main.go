package main

import (
	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
	"github.com/SakthiMahendran/SmartLibrary/server"
	"github.com/SakthiMahendran/SmartLibrary/server/controller"
)

func main() {
	dbdriver.Nothing()
	controller.Nothing()

	smartServer := server.NewServer()
	smartServer.Start("8080")
}
