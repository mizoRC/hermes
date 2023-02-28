package handlers

import (
	"fmt"
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func participantAuth(c *gin.Context) {
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

func AuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/participant", participantAuth)
}
