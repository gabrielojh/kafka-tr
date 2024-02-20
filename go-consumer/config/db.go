package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gabrielojh/kafka-tr/collections"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetDBInstance() (*mongo.Client, error) {

	mongoOnce.Do(func() {
		// Load environment variables
		user := os.Getenv("MONGO_USERNAME")
		pass := os.Getenv("MONGO_PASSWORD")
		host := os.Getenv("MONGO_HOST")
		port := os.Getenv("MONGO_PORT")

		conn := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		log.Println("Connection String: ", conn)

		// Use the SetServerAPIOptions() method to set the Stable API version to 1
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		opts := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPI)
		// Create a new client and connect to the server
		clientInstance, clientInstanceError = mongo.Connect(ctx, opts)
		if clientInstanceError != nil {
			panic(clientInstanceError)
		}

		// Send a ping to confirm a successful connection
		log.Println("Pinging your deployment...")
		if clientInstanceError = clientInstance.Ping(context.Background(), nil); clientInstanceError != nil {
			log.Println("Can't connect")
			panic(clientInstanceError)
		}
		log.Println("Pinged your deployment. You successfully connected to MongoDB!")

		initIndexes(clientInstance)

		collections.TransactionCollection = OpenCollection("transactions")
	})

	return clientInstance, clientInstanceError
}

func initIndexes(client *mongo.Client) {

	// transactions_transactions_-1 index
	transactionCollection := OpenCollection("transactions")

	transactionIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: -1}, {Key: "category", Value: -1}},
		Options: options.Index().SetUnique(true),
	}
	transactionIndexCreated, err := transactionCollection.Indexes().CreateOne(context.Background(), transactionIndexModel)
	if err != nil {
		log.Println("Error creating index")
		log.Fatal(err)
	}

	log.Printf("Created Transaction Index %s\n", transactionIndexCreated)
}

func OpenCollection(collectionName string) *mongo.Collection {

	var collection *mongo.Collection = clientInstance.Database("is459").Collection(collectionName)

	return collection
}
