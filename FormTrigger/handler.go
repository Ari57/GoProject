package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "GolangDB"
	collectionName = "GolangName"
)

type NameResult struct {
	Name string `json:"name"`
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // gets name value e.g. url?name=...
	delete := r.URL.Query().Get("delete")

	if delete == "yes" {
		message := deleteDocuments() // deletes all name records
		w.Write([]byte(message))
	}

	if name == "" && delete != "yes" {
		w.Write([]byte("Hi"))
	}

	if name != "" {
		w.Write([]byte("Hey " + name + "\n"))
		insertName(name)

		queryData := queryName()
		w.Write([]byte("\n"))
		w.Write([]byte("Current list of names stored in the database: \n"))
		w.Write([]byte(queryData + "\n"))
	}
}

func main() {
	listenAddr := ":7071"                                            // default address
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok { // will use this address if env variable is set
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/FormTrigger", nameHandler) // registers the nameHandler function to handle incoming requests for the endpoint

	go func() {
		log.Fatal(http.ListenAndServe(listenAddr, nil))
	}()

	select {}
}

func getCollection() *mongo.Collection {
	err := godotenv.Load("credential.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	MongoURI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	coll := client.Database(dbName).Collection(collectionName)
	return coll

}

func insertName(name string) {
	collection := getCollection()

	type Person struct {
		Name string
	}

	doc := Person{Name: name}
	InsertResult, err := collection.InsertOne(context.Background(), doc)
	fmt.Printf("Inserted document with _id: %v\n", InsertResult.InsertedID)

	basicErrorChecker(err)
}

func queryName() string {
	collection := getCollection()
	filter := bson.M{"name": bson.M{"$ne": ""}} // where name is not empty

	cursor, err := collection.Find(context.Background(), filter)
	basicErrorChecker(err)

	// defer cursor.Close(context.TODO())
	var names []string

	// Iterate through the cursor and decode the results
	for cursor.Next(context.Background()) {
		var result NameResult
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		names = append(names, result.Name)
	}

	resultString := strings.Join(names, "\n")
	return resultString
}

func basicErrorChecker(err error) {
	if err != nil {
		panic(err)
	}
}

func deleteDocuments() string {
	collection := getCollection()
	filter := bson.M{"name": bson.M{"$ne": ""}} // where name not empty

	result, err := collection.DeleteMany(context.Background(), filter)
	_ = result

	if err != nil {
		log.Fatalf("Error deleting documents: %v", err)
	}

	message := "Deleted all name records"
	return message
}
