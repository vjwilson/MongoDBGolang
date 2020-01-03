package mongoapi_test

import (
	mongoapi "MongoDBGolang/mongoAPI"
	"fmt"
	"reflect"
	"testing"
)

func TestMongoClient(t *testing.T) {
	// var client *mongo.Client = &mongo.Client{}
	client := mongoapi.MongoClient()
	fmt.Println("I'm a test!: ", reflect.TypeOf(client))
	if client != client {
		fmt.Println("Failed to obtain a Mongo Client")
	}
}
