package main

import (
	"fmt"
	"github.com/councilbox/hermes/db"
	"github.com/councilbox/hermes/logger"
	"github.com/councilbox/hermes/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const VERSION = "0.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	logger.Init()

	// DISPLAY LOGO
	DisplayLogo()

	// Connect to MongoDb
	dbErr := db.Connect()

	if dbErr == nil {
		logger.Logger.Info("MongoDB connected")
		runningMsg := fmt.Sprintf("Hermes v%s running on  http://0.0.0.0%s", VERSION, port)
		logger.Logger.Info(runningMsg)

		// Setting up Gin
		gin.SetMode(gin.ReleaseMode)
		server := gin.Default()
		server = router.New(server)
		err := server.Run(port)
		if err != nil {
			db.Disconnect()
			log.Fatalf("Error running server: %v", err)
		}
	} else {
		log.Fatalf("Error connecting mongoDB: %v", dbErr)
	}
}
