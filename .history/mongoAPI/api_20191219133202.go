package mongoapi

import (
	"MongoDBGolang/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient returns a connection to an instance of Mongo
func MongoClient() *mongo.Client {
	// TODO: Make the URI here configurable.
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// err = client.Connect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

/*

	CREATE

*/

// AddToGamerCollection adds a new gamer to the database. You may add an arbitrary number of gamers.
func AddToGamerCollection(collection *mongo.Collection, gamer ...interface{}) error {
	// TODO: Make the clientDB into an interface so these methods could be used with any DB.
	if len(gamer) > 1 {
		_, err := collection.InsertMany(context.TODO(), gamer)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := collection.InsertOne(context.TODO(), gamer[0])
	if err != nil {
		return err
	}
	return nil
}

/*

	READ

*/

// FindOneInCollection returns a gamer containing the fields
func FindOneInCollection(databaseName string, collectionName string, gamerName interface{}, projections []interface{}) (models.Gamer, error) {
	collection := MongoClient().Database(databaseName).Collection(collectionName)

	var result models.Gamer
	filter := bson.M{
		"name": gamerName,
	}
	projectionOpts := bson.M{}
	for _, key := range projections {
		projectionOpts[key.(string)] = 1
	}

	collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectionOpts)).Decode(&result)
	fmt.Println("\n\nI'm the Result!\n\n", result)
	return result, nil
}

/*

	UPDATE

*/

// UpdateOneGamerByName allows the gamers name and age to be changed.
func UpdateOneGamerByName(databaseName string, collectionName string, gamerName string, updateInfo interface{}) {
	collection := MongoClient().Database(databaseName).Collection(collectionName)
	filter := bson.M{
		"name": gamerName,
	}
	update := bson.M{
		"$set": bson.M{
			"age": updateInfo,
		},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Failed to update: ", err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// AddGameToGamerGamelist adds a game to a gamer's game list.
func AddGameToGamerGamelist() {
	// 	client := mongoapi.MongoClient()
	// 	gamerCollection := client.Database("BR").Collection("Gamers")
	// 	filter := bson.M{
	// 		"name": "Sean",
	// 	}
	// game := models.Game{
	// 		Name:     "Tobal #1",
	// 		Genre:    "1 v 1 Fighting",
	// 		Platform: "Playstation",
	// 	}
	// updateJSON, err := json.Marshal(game)
	// 	var updateJSONMap map[string]interface{}
	// 	err = json.Unmarshal(updateJSON, &updateJSONMap)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// update := bson.M{
	// 		"$push": bson.M{
	// 			"gamelist": updateJSONMap,
	// 		},
	// 	}

	// 	collection, err := gamerCollection.UpdateOne(context.TODO(), filter, update)
	// 	if err != nil {
	// 		fmt.Println("Failed to update: ", err)
	// 	}

	// 	fmt.Println("Collection: ", collection.UpsertedID)
}

// DeleteOneGamerFromCollectionByName deletes any gamers passed as arguments.
func DeleteOneGamerFromCollectionByName(collection *mongo.Collection, gamerName interface{}) error {
	filter := bson.M{
		"name": gamerName,
	}
	results, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println(results.DeletedCount)
	return nil
}

// DropCollection drops the passed collection from the database.
func DropCollection(collection *mongo.Collection) error {
	err := collection.Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
