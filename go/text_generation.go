package examples

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func GenerateContentTextOnly() (*genai.GenerateContentResponse, error) {
	// [START text_gen_text_only_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("Write a story about a magic backpack."),
		&genai.GenerateContentConfig{
			Temperature: genai.Ptr[float32](0.5),
		},
	)
	printResponse(result)
	// [END text_gen_text_only_prompt]
	return result, err
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part.Text)
			}
		}
	}
}
