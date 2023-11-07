package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "GolangDB"
	collectionName = "GolangName"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // gets name value e.g. url?name=...
	if name == "" {
		w.Write([]byte("Hi"))
	} else {
		w.Write([]byte("Hey " + name))

		insertName(name)
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

func queryName(name string) {
	collection := getCollection()
	filter := bson.M{"name": name}
	cursor, err := collection.Find(context.TODO(), filter)

	basicErrorChecker(err)

	// defer cursor.Close(context.TODO())
	var QueryResult []bson.M

	// Iterate through the cursor and decode the results
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			panic(err)
		}
		QueryResult = append(QueryResult, result)
	}

	// Marshal and print the results as JSON
	jsonData, err := json.MarshalIndent(QueryResult, "", "    ")

	basicErrorChecker(err)
	fmt.Printf("%s\n", jsonData)
}

func basicErrorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
