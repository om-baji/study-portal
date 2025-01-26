package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func getQuestion(limit int64, tag string) (*genai.GenerateContentResponse, error) {
	api_key := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))

	if err != nil {
		log.Fatal("Error connecting to gemini!")
	}

	prompt := fmt.Sprintf(
		"Give %d MCQ question(s) on advanced, high-level, and core %s. It should return an array of objects with fields for options, question, correctAnswer, and tag with the subject name matching my input. Avoid repeated questions.",
		limit, tag)

	model := client.GenerativeModel("gemini-1.5-flash")

	defer client.Close()

	response, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	return response, nil

}
