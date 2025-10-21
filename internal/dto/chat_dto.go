package dto

type ChatRequest struct {
	Message string `json:"message" binding:"required,min=3"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}
