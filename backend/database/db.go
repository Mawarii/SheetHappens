package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func Init() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Couldn't connect to the database:", err)
	}

	db = client.Database("sheethappens")
	if err := createUniqueIndex("users", "username"); err != nil {
		log.Fatal("Could not create index:", err)
	} else {
		log.Println("Unique index on 'username' created successfully")
	}
}

func createUniqueIndex(coll string, index string) error {
	collection := db.Collection(coll)

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: index, Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	return err
}

func CloseConnection() {
	err := db.Client().Disconnect(context.Background())
	if err != nil {
		log.Fatal("Couldn't close the database connection:", err)
	}
}
