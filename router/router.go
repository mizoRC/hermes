package router

import (
	"github.com/councilbox/hermes/handlers"
	"github.com/gin-gonic/gin"
)

func New(server *gin.Engine) *gin.Engine {
	// Setting up API routes
	api := server.Group("/api")
	auth := server.Group("/auth")
	users := server.Group("/users")
	spaces := server.Group("/spaces")
	messages := server.Group("/messages")

	handlers.ApiRoutes(api)
	handlers.AuthRoutes(auth)
	handlers.UserRoutes(users)
	handlers.SpaceRoutes(spaces)
	handlers.MessageRoutes(messages)

	return server
}
