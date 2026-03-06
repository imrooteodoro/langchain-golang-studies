package routes

import (
	"langchain-chat/src/models"
	"langchain-chat/src/services"

	"github.com/gin-gonic/gin"
)

func ChatRoute(r *gin.Engine) {
	r.POST("/chat", chatHandler)
}

func chatHandler(ctx *gin.Context) {
	var request models.ChatRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, models.ErrorResponse{Error: "Invalid request"})
		return
	}

	client, err := services.CreateOllamaClient("qwen2.5vl:latest")
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{Error: "Failed to create Ollama client"})
		return
	}

	response, err := services.ChatWithOllama(ctx.Request.Context(), client, request.Prompt)
	if err != nil {
		ctx.JSON(500, models.ErrorResponse{Error: "Failed to chat with Ollama"})
		return
	}

	ctx.JSON(200, models.ChatResponse{Response: response})
}