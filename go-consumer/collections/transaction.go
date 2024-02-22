package collections

import (
	"context"
	"time"

	"github.com/gabrielojh/kafka-tr/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TransactionCollection *mongo.Collection;

func RetrieveAllTransactions() (transactions []models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	defer cancel()

	options := options.Find().SetSort(bson.M{"Name": 1})
	cursor, err := TransactionCollection.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &transactions)

	return transactions, err
}

func RetrieveSpecificTransactionByNameAndCategory(name string, category string) (result *models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "name", Value: name}, {Key: "category", Value: category}}

	result = &models.Transaction{}
	err = TransactionCollection.FindOne(ctx, filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return result, err
}

func CreateTransaction(transaction models.Transaction) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err = TransactionCollection.InsertOne(ctx, transaction)

	return result, err
}

func UpdateTransaction(transaction models.Transaction) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": transaction.Name, "category": transaction.Category}
	update := bson.M{"$set": transaction}
	result, err = TransactionCollection.UpdateOne(ctx, filter, update)

	return result, err
}

func CreateBulkTransactions(operations []mongo.WriteModel) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

    result, err = TransactionCollection.BulkWrite(ctx, operations)

	return result, err
}

// func RetrieveCardValuesFromTransaction(cardId string) (result float64, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	pipeline := []bson.M{
// 		{"$match": bson.M{"card_id": cardId}},
// 		{"$group": bson.M{
// 			"_id": nil,
// 			"totalPoints": bson.M{"$sum": "$points"},
// 			"totalMiles": bson.M{"$sum": "$miles"},
// 			"totalCashback": bson.M{"$sum": "$cash_back"},
// 		}},
// 	}
	
// 	cursor, err := TransactionCollection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return result, err
// 	}
	
// 	var temp struct {
// 		TotalPoints   float64 `bson:"totalPoints"`
// 		TotalMiles    float64 `bson:"totalMiles"`
// 		TotalCashback float64 `bson:"totalCashback"`
// 	}

// 	if cursor.Next(context.Background()) {
        
// 		if err = cursor.Decode(&temp); err != nil {
// 			log.Println(err.Error())
// 			return result, err
// 		}
// 	}

// 	result += temp.TotalCashback + temp.TotalMiles + temp.TotalPoints

// 	return result, err
// }

// func DeleteUnprocessedByTransactionId(transactionIdList []string) (result *mongo.DeleteResult, err error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	filter := bson.M{"transaction_id": bson.M{"$in": transactionIdList}}

// 	result, err = unprocessedCollection.DeleteMany(ctx, filter)

// 	return result, err
// }