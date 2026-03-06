package main

import (
	_ "langchain-chat/docs"
	"langchain-chat/src/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	server := gin.Default()
	routes.ChatRoute(server)

	server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "server is healthy",
		})
	})

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
