package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func GetDBInstance() (clientInstance *mongo.Client, clientInstanceError error) {

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
		client, err := mongo.Connect(ctx, opts)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		// Send a ping to confirm a successful connection
		log.Println("Pinging your deployment...")
		if err := client.Ping(context.Background(), nil); err != nil{
			log.Println("Can't connect")
			panic(err)
		}
		log.Println("Pinged your deployment. You successfully connected to MongoDB!")
		// // initialise indexes
		// InitIndexes(client)
		// log.Println("Success!")
	})
	return clientInstance, clientInstanceError
}

// func InitIndexes(client *mongo.Client) {

// 	// transactions_transactions_-1 index
// 	transactionCollection := OpenCollection(client, "transactions")

// 	transactionIndexModel := mongo.IndexModel{
// 		Keys:    bson.D{{Key: "transaction_id", Value: -1}},
// 		Options: options.Index().SetUnique(true),
// 	}
// 	transactionIndexCreated, err := transactionCollection.Indexes().CreateOne(context.Background(), transactionIndexModel)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// unprocessed_unprocessed-1 index
// 	unprocessedCollection := OpenCollection(client, "unprocessed")

// 	unprocessedIndexModel := mongo.IndexModel{
// 		Keys:    bson.D{{Key: "transaction_id", Value: -1}},
// 		Options: options.Index().SetUnique(true),
// 	}
// 	unprocessedIndexCreated, err := unprocessedCollection.Indexes().CreateOne(context.Background(), unprocessedIndexModel)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// campaigns_campaigns_-1 index
// 	campaignCollection := OpenCollection(client, "campaigns")

// 	campaignIndexModel := mongo.IndexModel{
// 		Keys:    bson.D{{Key: "campaign_id", Value: -1}},
// 		Options: options.Index().SetUnique(true),
// 	}

// 	campaignIndexCreated, err := campaignCollection.Indexes().CreateOne(context.Background(), campaignIndexModel)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// cards_cards_-1 index
// 	cardCollection := OpenCollection(client, "cards")

// 	cardIndexModel := mongo.IndexModel{
// 		Keys:    bson.D{{Key: "card_id", Value: -1}},
// 		Options: options.Index().SetUnique(true),
// 	}

// 	cardIndexCreated, err := cardCollection.Indexes().CreateOne(context.Background(), cardIndexModel)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// user_users_-1 index
// 	userCollection := OpenCollection(client, "users")

// 	userIndexModel := mongo.IndexModel{
// 		Keys: bson.D{{Key: "user_id", Value: -1}},
// 		Options: options.Index().SetUnique(true),
// 	}

// 	userIndexCreated, err := userCollection.Indexes().CreateOne(context.Background(), userIndexModel)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Printf("Created Transaction Index %s\n", transactionIndexCreated)
// 	log.Printf("Created Unprocessed Index %s\n", unprocessedIndexCreated)
// 	log.Printf("Created Campaign Index %s\n", campaignIndexCreated)
// 	log.Printf("Created Card Index %s\n", cardIndexCreated)
// 	log.Printf("Created User Index %s\n", userIndexCreated)
// }

// func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

// 	var collection *mongo.Collection = client.Database("loyalty").Collection(collectionName)

// 	return collection
// }
