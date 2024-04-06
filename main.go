package main

import (
	"fmt"

	"github.com/SakthiMahendran/SmartLibrary/server"
)

func main() {
	fmt.Println("Starting server")
	smartServer := server.NewServer()

	fmt.Printf("smartServer.Start(\"8080\"): %v\n", smartServer.Start("8080").Error())
}
