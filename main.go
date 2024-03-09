package main

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

	/*client, err := connectMongoDB()
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

	fmt.Println("Book deleted successfully!")*/

	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Initialize controller
	controller := controllers.BookInventoryController{Client: client}

	// Test GetBookCount function
	bookName := "updated sample book" // Example book name
	count, err := controller.GetBookCount(bookName)
	if err != nil {
		fmt.Println("Error fetching book count:", err)
		return
	}
	fmt.Printf("Count of '%s' in inventory: %d\n", bookName, count)*/
	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Initialize controller
	controller := controllers.BookInventoryController{Client: client}

	// Test GetCategoryCount function
	department := "updated fake department" // Example department
	count, err := controller.GetCategoryCount(department)
	if err != nil {
		fmt.Println("Error fetching category count:", err)
		return
	}
	fmt.Printf("Count of books in '%s' department: %d\n", department, count)*/
	/*
		// Set up MongoDB client
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			fmt.Println("Error connecting to MongoDB:", err)
			return
		}
		defer func() {
			if err := client.Disconnect(context.Background()); err != nil {
				fmt.Println("Error disconnecting from MongoDB:", err)
			}
		}()

		// Initialize controller
		controller := controllers.BookInventoryController{Client: client}

		// Test FindCategory function
		department := "updated fake department" // Example department
		books, err := controller.FindCategory(department)
		if err != nil {
			fmt.Println("Error fetching books by category:", err)
			return
		}
		fmt.Println(books)

		// Print found books
		fmt.Printf("Books in '%s' department:\n", department)
		for _, book := range books {
			fmt.Printf("Book Name: %s, Author: %s\n", book.BookName, book.Author)
		}*/

	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()

	// Initialize controller
	controller := controllers.BookInventoryController{Client: client}

	// Test IsAvailable function
	bookName := "updated sample book" // Example book name
	available, err := controller.IsAvailable(bookName)
	if err != nil {
		fmt.Println("Error checking book availability:", err)
		return
	}

	if available {
		fmt.Printf("Book '%s' is available!\n", bookName)
	} else {
		fmt.Printf("Book '%s' is not available!\n", bookName)
	}*/

	/*client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Error creating MongoDB client:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// Initialize BookInventoryController
	bc := &controllers.BookInventoryController{
		Client: client,
	}

	// Simulate a book borrowing transaction
	bookID := "78hh"
	studentRegNo := "GHI789"
	err = bc.Borrow(bookID, studentRegNo)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
		return
	}

	fmt.Println("Book borrowed successfully!")*/
	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Create BookInventoryController instance
	bc := controllers.BookInventoryController{
		Client: client,
	}

	// Sample book and student data
	bookID := "78hh"         // Sample book ID
	studentRegNo := "GHI789" // Sample student registration number

	// Test Return function
	err = bc.Return(bookID, studentRegNo)
	if err != nil {
		fmt.Println("Error returning book:", err)
		return
	}

	fmt.Println("Book returned successfully!")*/
	/*client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	// Create a TransactionController instance
	transController := controllers.TransactionController{Client: client}

	// Example student registration number
	studentRegNo := "GHI789"

	// Find transactions for the student
	transactions, err := transController.FindStudentTransactions(studentRegNo)
	if err != nil {
		fmt.Println("Error finding student transactions:", err)
		return
	}

	// Print the transactions
	fmt.Println(transactions)*/
	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	// Create a new instance of the TransactionController
	tc := &controllers.TransactionController{Client: client}

	// Provide a book ID to find transactions for
	bookID := "78hh" // Replace with a valid book ID

	// Call the FindBookTransactions function
	transactions, err := tc.FindBookTransactions(bookID)
	if err != nil {
		fmt.Println("Error finding book transactions:", err)
		return
	}

	// Print the found transactions
	fmt.Println(transactions)*/
	/*clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize TransactionController
	tc := controllers.TransactionController{
		Client: client,
	}

	// Test FindDuedTransactions function
	transactions, err := tc.FindDuedTransactions()
	if err != nil {
		log.Fatal(err)
	}

	// Print the found transactions
	fmt.Println(transactions)*/

}
