package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connect only one time
func init() {
	// client option
	clientoption := options.Client().ApplyURI(connectionString)

	// connect with mongodb
	client, err := mongo.Connect(context.TODO(), clientoption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Established")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is Ready")
}
