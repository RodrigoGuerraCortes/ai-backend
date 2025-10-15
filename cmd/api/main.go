package main

import (
	"fmt"
	"log"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/http/router"
)

func main() {
	fmt.Println("🚀 Starting AI Backend with Gemini + Gin...")

	// Init AI Client
	geminiClient := ai.NewGeminiClient()
	if err := geminiClient.TestConnection(); err != nil {
		log.Fatalf("❌ Gemini connection failed: %v", err)
	}

	// Create router and inject dependencies
	r := router.NewRouter(geminiClient)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
