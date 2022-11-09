//basic CRUD of mongoDB using go programming language

package commit

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	//Connect
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongoDB!")
	collection := client.Database("test").Collection("user")

	//Insert
	qwq := User{1, "qwq", 1}
	qwqwq := User{2, "qwqwq", 2}
	users := []interface{}{qwq, qwqwq}
	insertManyResult, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	//Update
	filter := bson.D{{Key: "id", Value: 1}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "age", Value: 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	//Search
	var result User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %v\n", result)

	//Delete
	delteeResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the users collection\n", delteeResult.DeletedCount)

	//Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Disconnected from mongoDB!")
}
