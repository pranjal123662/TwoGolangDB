package controller

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName1 = "DuplicateLoginBucket"
const dbName2 = "UserDataBucket"
const colName1 = "LoginData"
const colName2 = "UserData"
const connectionString = "mongodb+srv://Pranjal:Pranjal%40123@cluster0.sc7ucqz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

var collection1 *mongo.Collection
var collection2 *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection1 = client.Database(dbName1).Collection(colName1)
	collection2 = client.Database(dbName2).Collection(colName2)
}
func InsertIntoLoginBucket(number string) bool {
	filter := bson.M{"number": number}
	insertedId, err := collection1.InsertOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return insertedId != nil
}
func InsertIntoUserDataBucket(name string) bool {
	filter := bson.M{"name": name}
	insertedId, err := collection2.InsertOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return insertedId != nil
}
func FetchFromLoginDB(number string) bool {
	filter := bson.M{"number": number}
	var result bson.M
	err := collection1.FindOne(context.Background(), filter).Decode(&result)
	return err != mongo.ErrNoDocuments
}
func FetchFromUserDataDB(name string) bool {
	filter := bson.M{"name": name}
	var result bson.M
	err := collection2.FindOne(context.Background(), filter).Decode(&result)
	return err != mongo.ErrNoDocuments
}
