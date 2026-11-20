package examples

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

func ConfigureModelParameters() (*genai.GenerateContentResponse, error) {
	// [START configure_model_parameters]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create local variables for parameters.
	candidateCount := int32(1)
	maxOutputTokens := int32(4096)

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-4-hydrogen",
		genai.Text("Tell me a story about a magic backpack."),
		&genai.GenerateContentConfig{
			CandidateCount:  candidateCount,
			StopSequences:   []string{"x"},
			MaxOutputTokens: maxOutputTokens,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(response)
	// [END configure_model_parameters]
	return response, err
}
