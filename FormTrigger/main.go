// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const (
// 	dbName         = "GolangDB"
// 	collectionName = "GolangName"
// )

// func main() {
// 	err := godotenv.Load("credential.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	MongoURI := os.Getenv("MONGODB_URI")

// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoURI))
// 	if err != nil {
// 		log.Fatalf("Error connecting to MongoDB: %v", err)
// 	}

// 	defer func() {
// 		if err := client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	coll := client.Database(dbName).Collection(collectionName)

// 	insertName(coll, "Doug")
// 	queryName(coll, "Doug")

// 	// deleteDocuments(coll)
// }

// func insertName(collection *mongo.Collection, name string) {
// 	type Person struct {
// 		Name string
// 	}

// 	doc := Person{Name: name}
// 	InsertResult, err := collection.InsertOne(context.TODO(), doc)
// 	fmt.Printf("Inserted document with _id: %v\n", InsertResult.InsertedID)

// 	basicErrorChecker(err)
// }

// func queryName(collection *mongo.Collection, name string) {
// 	filter := bson.M{"name": name}
// 	cursor, err := collection.Find(context.TODO(), filter)

// 	basicErrorChecker(err)

// 	defer cursor.Close(context.TODO())
// 	var QueryResult []bson.M

// 	// Iterate through the cursor and decode the results
// 	for cursor.Next(context.TODO()) {
// 		var result bson.M
// 		if err := cursor.Decode(&result); err != nil {
// 			panic(err)
// 		}
// 		QueryResult = append(QueryResult, result)
// 	}

// 	// Marshal and print the results as JSON
// 	jsonData, err := json.MarshalIndent(QueryResult, "", "    ")

// 	basicErrorChecker(err)
// 	fmt.Printf("%s\n", jsonData)
// }

// func deleteDocuments(collection *mongo.Collection) {
// 	filter := bson.M{"name": bson.M{"$ne": ""}}
// 	count, Err := collection.CountDocuments(context.Background(), filter)
// 	result, err := collection.DeleteMany(context.Background(), filter)

// 	if err != nil {
// 		log.Fatalf("Error deleting documents: %v", err)
// 	}
// 	fmt.Printf("Deleted %v document(s)\n", count)

// 	_, _ = result, Err
// }

// func basicErrorChecker(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
