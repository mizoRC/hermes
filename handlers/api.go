package handlers

import (
	"encoding/json"
	"github.com/councilbox/hermes/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type config struct {
	port    string
	version string
}

type apiInfo struct {
	mongo  string
	config config
	alive  bool
}

func health(c *gin.Context) {
	ready := db.Status()
	var status string
	var httpStatus int
	if ready {
		status = "Ready"
		httpStatus = http.StatusOK
	} else {
		status = "Failed"
		httpStatus = http.StatusServiceUnavailable
	}

	info, err := json.Marshal(apiInfo{
		mongo: status,
		config: config{
			port:    "PORT",
			version: "version",
		},
		alive: ready,
	})

	if err != nil {
		httpStatus = http.StatusBadRequest
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(httpStatus, info)
	}
	/*c.JSON(httpStatus, gin.H{
		"db": status,
		"config": gin.H{
			"port":    "PORT",
			"version": "1",
		},
		"alive": strconv.FormatBool(ready),
	})*/
}

func ApiRoutes(rg *gin.RouterGroup) {
	rg.GET("/health", health)
}
