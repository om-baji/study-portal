package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/om-baji/models"
	"google.golang.org/api/option"
)

func getResponse(limit int64, tag string) (*genai.GenerateContentResponse, error) {
	api_key := os.Getenv("GEMINI_API_KEY")

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(api_key))

	if err != nil {
		log.Fatal("Error connecting to gemini!")
	}

	prompt := fmt.Sprintf(
		"Give %d MCQ question(s) on advanced, high-level, and core %s. "+
			"Questions should be incredibly hard testing one's core knowledge and fundamentals at its finest. "+
			"It should return an array of objects with fields for options, question, correctAnswer, and tag with the subject name matching my input. "+
			"Avoid repeated questions. It should only have a json response not anything other than that",
		limit, tag)

	model := client.GenerativeModel("gemini-1.5-flash")

	defer client.Close()

	response, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	return response, nil

}

func getQuestion(limit int64, tag string) []models.Question {

	var questions []models.Question

	response, err := getResponse(limit, tag)

	if err != nil {
		log.Fatal("Something went wrong!")
	}

	err = json.Unmarshal([]byte(response.PromptFeedback.BlockReason.String()), &questions)

	if err != nil {
		fmt.Printf("Something went wrong!\n %v\n", err.Error())
	}

	return questions
}
