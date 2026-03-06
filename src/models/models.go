package models

type ChatRequest struct {
	Prompt string `json:"prompt" example:"Explique o que é Go" binding:"required"`
}

type ChatResponse struct {
	Response string `json:"response" example:"Go é uma linguagem de programação..."`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Failed to chat with Ollama"`
}
