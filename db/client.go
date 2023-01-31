package db

import (
	"context"
	. "github.com/councilbox/hermes/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var Db *mongo.Client

func Status() bool {
	err := Db.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatalf("Error connecting to db database: %v", err)
		return false
	} else {
		return true
	}
}

func Connect() error {
	var mongodbURI string = os.Getenv("MONGO_URL")
	Logger.Info("Trying to connect MongoDB")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	pingErr := client.Ping(context.TODO(), readpref.Primary())
	if err != nil || pingErr != nil {
		return err
	} else {
		Db = client
		return nil
	}
}

func Disconnect() {
	if err := Db.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
