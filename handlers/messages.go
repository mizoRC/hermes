package handlers

import (
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Message struct {
	ID             uint `gorm:"primary_key"`
	Space_id       uint
	Participant_id uint
	Text           string
	Created_at     time.Time `gorm:"default:current_timestamp"`
}

func getAllMessages(c *gin.Context) {
	var messages []Message
	result := db.Client.Find(&messages)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

func getSpaceMessages(c *gin.Context) {
	var messages []Message
	spaceID := c.Query("spaceID")
	result := db.Client.First(&messages, spaceID)
	db.Client.Where("space_id = ?", spaceID).Find(&messages)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"messages": messages,
		})
	} else {
		c.String(http.StatusBadRequest, result.Error.Error())
	}
}

/*func addMessage(c *gin.Context) {
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
}*/

func MessageRoutes(rg *gin.RouterGroup) {
	rg.GET("/all", getAllMessages)
	rg.GET("/space", getSpaceMessages)
	//rg.POST("/add", addMessage)
}
