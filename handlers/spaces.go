package handlers

import (
	"fmt"
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Space struct {
	ID         uint `gorm:"primary_key"`
	Name       string
	Created_at time.Time `gorm:"default:current_timestamp"`
}

func getAllSpaces(c *gin.Context) {
	var spaces []Space
	result := db.Client.Find(&spaces)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"spaces": spaces,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func getSpace(c *gin.Context) {
	var space Space
	spaceID := c.Query("spaceID")
	result := db.Client.First(&space, spaceID)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"space": space,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func createSpace(c *gin.Context) {
	name := c.PostForm("name")
	space := Space{
		Name: name,
	}

	result := db.Client.Create(&space)

	// check for errors in the insertion
	if result.Error == nil {
		// display the id of the newly inserted object
		fmt.Println(space.ID)

		c.JSON(http.StatusOK, gin.H{
			"spaceID": space.ID,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func SpaceRoutes(rg *gin.RouterGroup) {
	rg.GET("/all", getAllSpaces)
	rg.GET("/get", getSpace)
	rg.POST("/create", createSpace)
}
