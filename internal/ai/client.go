package ai

import (
	"context"
	"errors"
	"log"

	"github.com/RodrigoGuerraCortes/ai-backend/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	Client *genai.Client
}

func NewGeminiClient() *GeminiClient {
	config.LoadEnv()
	apiKey := config.GetEnv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ Missing GEMINI_API_KEY in environment variables")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("❌ Failed to create Gemini client: %v", err)
	}

	return &GeminiClient{Client: client}
}

func (g *GeminiClient) TestConnection() error {
	ctx := context.Background()
	model := g.Client.GenerativeModel("gemini-2.0-flash-exp")

	_, err := model.GenerateContent(ctx, genai.Text("ping"))
	return err
}

func (g *GeminiClient) Chat(ctx context.Context, message string) (string, error) {
	model := g.Client.GenerativeModel("gemini-2.0-flash-exp")

	resp, err := model.GenerateContent(ctx, genai.Text(message))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 ||
		resp.Candidates[0].Content == nil ||
		len(resp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("no response from AI model")
	}

	reply := resp.Candidates[0].Content.Parts[0].(genai.Text)
	return string(reply), nil
}
