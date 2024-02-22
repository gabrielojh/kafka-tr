package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gabrielojh/kafka-tr/collections"
	"github.com/gabrielojh/kafka-tr/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config.InitEnvironment()

	client, err := config.GetDBInstance()
	if err != nil {
		log.Fatalf("Failed to get MongoDB client: %v", err)
	}

	// Close db connection
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	consumer, err := getKafkaConsumer()
	if err != nil {
		log.Fatal("Failed to create Consumer:", err)
	}

	topic := "test-topic"

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal("Failed to subscribe to topic:", err)
	}

	processTransactions(consumer)
}

func processTransactions(consumer *kafka.Consumer) {

	const BATCH_SIZE int = 10000
	var operations []mongo.WriteModel

	for {

		// Process transactions in batches
		for i := 0; i < BATCH_SIZE; i++ {

			msg, err := consumer.ReadMessage(time.Second)
			if err != nil {
				break
			}

			// log.Println("Received message", string(msg.Value))

			temp := strings.Split(string(msg.Value), ",")
			if temp[1] == "Name" {
				consumer.Commit()
				continue
			}

			name := temp[1]
			credit, err := strconv.Atoi(temp[2])
			if err != nil {
				log.Fatalln("Error converting credit to int: ", err)
			}
			category := temp[3]

			filter := bson.M{"name": name, "category": category}
			update := bson.M{
				"$setOnInsert": bson.M{"name": name, "category": category},
				"$inc":         bson.M{"credit": credit},
			}
			operations = append(operations, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
			// transaction, err := collections.RetrieveSpecificTransactionByNameAndCategory(name, category)
			// if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
			// 	log.Fatalf("Error retrieving transaction: %v", err)
			// }

			// if transaction == nil {
			// 	transaction = &models.Transaction{
			// 		Name:     name,
			// 		Credit:   credit,
			// 		Category: category,
			// 	}
			// 	log.Println("Creating transaction", *transaction)
			// 	_, err := collections.CreateTransaction(*transaction)
			// 	if err != nil {
			// 		log.Fatalf("Error creating transaction: %v", err)
			// 	}
			// } else {
			// 	log.Println("Updating transaction", *transaction)
			// 	transaction.Credit += credit
			// 	_, err := collections.UpdateTransaction(*transaction)
			// 	if err != nil {
			// 		log.Fatalf("Error updating transaction: %v", err)
			// 	}
			// }
		}

		if len(operations) == 0 {
			continue
		}

		log.Println("Committing transactions: ", len(operations))
		_, err := collections.CreateBulkTransactions(operations)
		if err != nil {
			log.Fatalf("Error creating bulk transactions: %v", err)
		}
		log.Println("Committed transactions successfully")
		consumer.Commit()
		operations = nil
	}
}

func getKafkaConsumer() (consumer *kafka.Consumer, err error) {
	server := os.Getenv("KAFKA_BOOTSTRAP_SERVER")

	return kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        server,
		"group.id":                 "CsvWorkerGroup",
		"client.id":                "CsvProcessor",
		"enable.auto.commit":       false,
		"enable.auto.offset.store": true,
		"auto.offset.reset":        "earliest",
		"isolation.level":          "read_committed",
	})
}
