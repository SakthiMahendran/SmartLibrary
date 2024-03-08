package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SakthiMahendran/SmartLibrary/dbdriver/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongoDB() (*mongo.Client, error) {
	// Set MongoDB connection string
	connectionString := "mongodb://localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return client, nil
}

func main() {
	// dbdriver.Nothing()
	// controller.Nothing()

	//smartServer := server.NewServer()
	//smartServer.Start("8080")

	//client, _ := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))

	//adminController := controllers.NewAdminController(client, "SMLS", "Admin")

	//book controller checking
	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	database := client.Database("SMLS") // Replace "your_database_name_here" with your actual database name

	bookController := controllers.NewBookController(database)

	bookID := "7hhh" // Replace "your_book_id_here" with an actual book ID
	borrowed, err := bookController.IsBorrowed(bookID)
	fmt.Println(borrowed)
	fmt.Println(err)*/
	/*client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)

		// Instantiate BookInventoryController
		bookController := controllers.NewBookInventoryController(client)

		//AddBook(bookID, bookName, author, bookDept string)

		// Add book
		var x string
		x, err = bookController.AddBook("7hhh", "abc book", "gjhj author", "hh department")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(x)

		fmt.Println("Book added successfully!")

	}*/
	/*client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create a BookInventoryController instance
	bc := controllers.BookInventoryController{Client: client}

	// Prepare present and to be updated book inventory data
	presentData := &models.BookInventory{
		BookName: "abc book",
	}

	toBeUpdatedData := &models.BookInventory{
		Author: "pranav",
	}

	// Call the UpdateBook function
	err = bc.UpdateBook(presentData, toBeUpdatedData)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("UpdateBook function executed successfully")*/

	client, err := connectMongoDB()
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB client:", err)
		}
	}()

	// Initialize controller
	bookInventoryController := controllers.NewBookInventoryController(client)

	// Sample book ID to delete
	bookID := "id_sample"

	// Test DeleteBook function
	err = bookInventoryController.DeleteBook(bookID)
	if err != nil {
		log.Fatal("Error deleting book:", err)
	}

	fmt.Println("Book deleted successfully!")
}
