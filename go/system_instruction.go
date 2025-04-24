package examples

import (
	"context"
	"os"
	"log"

	"google.golang.org/genai"
)

func SystemInstruction() error {
	// [START system_instruction]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Construct the user message contents.
	contents := []*genai.Content{
		genai.NewContentFromText("Good morning! How are you?", "user"),
	}

	// Set the system instruction as a *genai.Content.
	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a cat. Your name is Neko.", "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, config)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END system_instruction]
	return err
}
