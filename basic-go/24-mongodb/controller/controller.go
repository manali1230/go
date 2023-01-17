package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/manali1230/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// insert one record - helper file

func insertOneRecord(movie model.Netflix) {
	fmt.Println("Insert One Record")
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data inserted successfully with id: ", inserted.InsertedID)
}
