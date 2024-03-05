package main

import (
	"context"

	"fmt"
	"time"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBDriver struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func Connect(connectionString, dbName, collectionName string) (*MongoDBDriver, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	} else {
		fmt.Println("connected")
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &MongoDBDriver{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func main() {
	Connect("mongodb://localhost:27017", "SMLS", "Student")

}

/*func (m *MongoDBDriver) Disconnect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err := m.client.Disconnect(ctx)
    if err != nil {
        return fmt.Errorf("error disconnecting from MongoDB: %v", err)
    }

    return nil
}

func (m *MongoDBDriver) InsertAdmin(admin *Admin) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := m.collection.InsertOne(ctx, admin)
    if err != nil {
        return fmt.Errorf("error inserting admin into MongoDB: %v", err)
    }

    return nil
}

func (m *MongoDBDriver) FindAdminByID(id string) (*Admin, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var admin Admin
    err := m.collection.FindOne(ctx, bson.M{"admin_id": id}).Decode(&admin)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, nil // Admin not found
        }
        return nil, fmt.Errorf("error finding admin by ID from MongoDB: %v", err)
    }

    return &admin, nil
}

func (m *MongoDBDriver) InsertBook(book *Book) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := m.collection.InsertOne(ctx, book)
    if err != nil {
        return fmt.Errorf("error inserting book into MongoDB: %v", err)
    }

    return nil
}


func (m *MongoDBDriver) FindBookByID(id string) (*Book, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var book Book
    err := m.collection.FindOne(ctx, bson.M{"book_id": id}).Decode(&book)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, nil // Book not found
        }
        return nil, fmt.Errorf("error finding book by ID from MongoDB: %v", err)
    }

    return &book, nil
}

func (m *MongoDBDriver) InsertStudent(student *Student) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := m.collection.InsertOne(ctx, student)
    if err != nil {
        return fmt.Errorf("error inserting student into MongoDB: %v", err)
    }

    return nil
}

func (m *MongoDBDriver) FindStudentByID(id string) (*Student, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var student Student
    err := m.collection.FindOne(ctx, bson.M{"reg_no": id}).Decode(&student)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, nil // Student not found
        }
        return nil, fmt.Errorf("error finding student by ID from MongoDB: %v", err)
    }

    return &student, nil
}

func (m *MongoDBDriver) InsertTransaction(transaction *Transaction) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := m.collection.InsertOne(ctx, transaction)
    if err != nil {
        return fmt.Errorf("error inserting transaction into MongoDB: %v", err)
    }

    return nil
}

func (m *MongoDBDriver) FindTransactionsByStudentID(id string) ([]*Transaction, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := m.collection.Find(ctx, bson.M{"reg_no": id})
    if err != nil {
        return nil, fmt.Errorf("error finding transactions by student ID from MongoDB: %v", err)
    }
    defer cursor.Close(ctx)

    var transactions []*Transaction
    for cursor.Next(ctx) {
        var transaction Transaction
        if err := cursor.Decode(&transaction); err != nil {
            return nil, fmt.Errorf("error decoding transaction from MongoDB: %v", err)
        }
        transactions = append(transactions, &transaction)
    }

    return transactions, nil
}*/
