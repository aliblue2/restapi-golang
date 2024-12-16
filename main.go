package main

import (
	"azno-space.com/azno/db"
	"azno-space.com/azno/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db.IntiDatabase()
	server := gin.Default()
	router.RouterHandler(server)
	server.Run(":8080")

}
