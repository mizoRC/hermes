package db

import (
	. "github.com/councilbox/hermes/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var Client *gorm.DB

func Status() bool {
	if client, err := Client.DB(); err != nil {
		log.Fatalf("Error connecting to db database: %v", err)
		return false
	} else {
		if pingErr := client.Ping(); pingErr != nil {
			log.Fatalf("Error connecting to db database: %v", err)
			return false
		} else {
			return true
		}
	}
}

func Connect() error {
	var postgresURI string = os.Getenv("POSTGRES_URL")
	Logger.Info("Trying to connect DB")
	client, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  postgresURI,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return err
	} else {
		Client = client
		return nil
	}
}

func Disconnect() {
	if client, err := Client.DB(); err != nil {
		panic(err)
	} else {
		client.Close()
	}
}
