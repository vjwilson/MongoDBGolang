package mongoapi

import (
	"MongoDBGolang/models"
	collectionhelper "MongoDBGolang/mongoAPI/collectionHelper"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Move these to MAIN and create a layer called app, which then references the various packages.
// Consider using Viper to get these things.
const DATABASE = "BR"
const COLLECTION = "Gamers"

// type DataAccess interface {
// 	Connect(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error)
// }

// type RealMongoClient struct {
// 	cl *mongo.Client
// }

// func (rm RealMongoClient) Connect(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {

// 	return mongo.Connect(ctx, opts...)
// }

// MongoClient returns a connection to an instance of Mongo
// func MongoClient() *mongo.Client {
// 	// TODO : Make the URI here configurable.
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")

// 	return client
// }

/*

	CREATE

*/

// AddToGamerCollection adds a new gamer to the database. You may add an arbitrary number of gamers.
// func AddToGamerCollection(collection *mongo.Collection, gamer ...interface{}) error {
func AddToGamerCollection(collection collectionhelper.ICollectionHelper, gamer ...interface{}) error {
	// TODO: Make the clientDB into an interface so these methods could be used with any DB.
	if gamer[0] == nil {
		return errors.New("Nothing to add to DB")
	}
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
func FindOneInCollection(collection collectionhelper.ICollectionHelper, gamerName interface{}, projections []interface{}) (models.Gamer, error) {
	// var mc RealMongoClient
	// collection := mc.MongoClient().Database(DATABASE).Collection(COLLECTION)
	if gamerName == nil {
		return models.Gamer{}, errors.New("Must enter a gamer name")
	}
	var result models.Gamer
	filter := bson.M{
		"name": gamerName,
	}
	projectionOpts := bson.M{}
	for _, key := range projections {
		projectionOpts[key.(string)] = 1
	}

	sr := collection.FindOne(context.TODO(), filter, options.FindOne().SetProjection(projectionOpts)).Decode(&result)
	fmt.Println("\n\nI'm the Result!\n\n", result)
	fmt.Println("I'm a single result:", sr)
	return result, nil
}

// /*

// 	UPDATE

// */

// UpdateOneGamerByName allows the gamers name and age to be changed.
func UpdateOneGamerByName(collection collectionhelper.ICollectionHelper, gamerName string, updateInfo interface{}) {
	// var mc RealMongoClient
	// collection := mc.MongoClient().Database(databaseName).Collection(collectionName)

	if gamerName == nil {
		return models.Gamer{}, errors.New("Must enter a gamer name")
	}
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

	fmt.Printf("Matched %v document and updated %v document.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// // AddGameToGamerGamelist adds a game to a gamer's game list.
// func AddGameToGamerGamelist(databaseName string, collectionName string, listUpdate models.GamelistUpdate) {
// 	var mc RealMongoClient

// 	collection := mc.MongoClient().Database(databaseName).Collection(collectionName)
// 	filter := bson.M{
// 		"name": listUpdate.Name,
// 	}
// 	fmt.Println("List Update: ", listUpdate)
// 	update := bson.M{
// 		"$push": bson.M{
// 			"gamelist": listUpdate.Game,
// 		},
// 	}

// 	_, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		fmt.Println("Failed to update: ", err)
// 	}
// }

/*

	DELETE

*/

// DeleteOneGamerFromCollectionByName deletes any gamers passed as arguments.
func DeleteOneGamerFromCollectionByName(collection collectionhelper.ICollectionHelper, gamerName interface{}) error {
	var filter bson.M
	if gamerName != nil {
		filter = bson.M{
			"name": gamerName,
		}
	} else {
		return errors.New("Name argument cannot be blank or <nil>")
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

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
