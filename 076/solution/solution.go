package main

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main (){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Create a variable named "uri" containing your MONGODB_URI string connection.
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Create a collection called "users" into your Database.
	usersCollection := client.Database("TestCluster").Collection("users")

	log.Println("You got connected!")
}