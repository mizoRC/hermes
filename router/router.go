package router

import (
	"github.com/councilbox/hermes/handlers"
	"github.com/gin-gonic/gin"
)

func New(server *gin.Engine) *gin.Engine {
	// Setting up API routes
	api := server.Group("/api")
	messages := server.Group("/messages")
	handlers.ApiRoutes(api)
	handlers.MessageRoutes(messages)

	return server
}
