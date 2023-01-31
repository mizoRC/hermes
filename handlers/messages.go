package handlers

import (
	"context"
	"fmt"
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Space     string             `bson:"space" json:"space"`
	UserID    string             `bson:"userID" json:"userID"`
	Text      string             `bson:"text" json:"text"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

func getAll(c *gin.Context) {
	db := db.Db.Database("hermes")
	collection := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(
		ctx,
		bson.D{},
	)
	if err == nil {
		var messages []Message
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var result Message
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Result: %v", result)
			messages = append(messages, result)
		}
		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	} else {
		c.Error(err)
	}
}

func getSpaceMessages(c *gin.Context) {
	space := c.Query("space")
	db := db.Db.Database("hermes")
	collection := db.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(
		ctx,
		bson.D{
			{Key: "space", Value: space},
		},
	)
	if err == nil {
		var messages []Message
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var result Message
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Result: %v", result)
			messages = append(messages, result)
		}
		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	} else {
		c.Error(err)
	}
}

func add(c *gin.Context) {
	text := c.PostForm("text")
	userID := uuid.New().String()
	collection := db.Db.Database("hermes").Collection("messages")
	// insert a single document into a collection
	// create a bson.D object
	/*message := bson.D{
		{"space", "prueba"},
		{"userID", userID,
		{"text", "Probando, probando!"},
	}*/
	message := Message{
		ID:        primitive.NewObjectID(),
		Space:     "prueba",
		UserID:    userID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	// insert the bson object using InsertOne()
	result, err := collection.InsertOne(context.TODO(), message)
	// check for errors in the insertion
	if err != nil {
		c.Error(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	c.JSON(http.StatusOK, gin.H{
		"messageID": result.InsertedID,
	})
}

func MessageRoutes(rg *gin.RouterGroup) {
	rg.GET("/all", getAll)
	rg.GET("/space", getSpaceMessages)
	rg.POST("/add", add)
}
