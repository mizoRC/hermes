package handlers

import (
	"fmt"
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	ID         uint `gorm:"primary_key"`
	Name       string
	Created_at time.Time `gorm:"default:current_timestamp"`
}

func getAllUsers(c *gin.Context) {
	var users []User
	result := db.Client.Find(&users)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func getUser(c *gin.Context) {
	var user User
	userID := c.Query("userID")
	result := db.Client.First(&user, userID)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func createUser(c *gin.Context) {
	name := c.PostForm("name")
	user := User{
		Name: name,
	}

	result := db.Client.Create(&user)

	// check for errors in the insertion
	if result.Error == nil {
		// display the id of the newly inserted object
		fmt.Println(user.ID)

		c.JSON(http.StatusOK, gin.H{
			"userID": user.ID,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func UserRoutes(rg *gin.RouterGroup) {
	rg.GET("/all", getAllUsers)
	rg.GET("/get", getUser)
	rg.POST("/create", createUser)
}
