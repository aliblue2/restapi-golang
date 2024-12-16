package router

import (
	"azno-space.com/azno/middleware"
	"github.com/gin-gonic/gin"
)

func RouterHandler(server *gin.Engine) {

	server.GET("/events", GetAllEventsHandler)
	server.GET("/events/:eventId", GetEventByIdHandler)

	/// protected routes
	authRequiresPath := server.Group("/")
	authRequiresPath.Use(middleware.Authenticate)
	authRequiresPath.DELETE("/events/:eventId", DeleteEventHandler)
	authRequiresPath.PUT("/events/:eventId", EditEventHandler)
	authRequiresPath.POST("/events", AddNewEventHandler)
	authRequiresPath.POST("/events/:eventId/register", RegisterEventHandler)
	authRequiresPath.DELETE("/events/:eventId/register", CancelEventHandler)

	server.POST("/signup", SignupUserHandler)
	server.POST("/login", LoginUserHandler)

}
