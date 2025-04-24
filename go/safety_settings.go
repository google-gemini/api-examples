package examples

import (
	"context"
	"fmt"
	"os"
	"log"
	"encoding/json"

	"google.golang.org/genai"
)

func SafetySettings() error {
	// [START safety_settings]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	unsafePrompt := "I support Martians Soccer Club and I think Jupiterians Football Club sucks! " +
		"Write a ironic phrase about them including expletives."

	config := &genai.GenerateContentConfig{
		SafetySettings: []*genai.SafetySetting{
			{
				Category:  "HARM_CATEGORY_HARASSMENT",
				Threshold: "BLOCK_ONLY_HIGH",
			},
		},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(unsafePrompt, genai.RoleUser),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, config)
	if err != nil {
		log.Fatal(err)
	}

	// Print the finish reason and safety ratings from the first candidate.
	if len(response.Candidates) > 0 {
		fmt.Println("Finish reason:", response.Candidates[0].FinishReason)
		safetyRatings, err := json.MarshalIndent(response.Candidates[0].SafetyRatings, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println("Safety ratings:", string(safetyRatings))
	} else {
		fmt.Println("No candidate returned.")
	}
	// [END safety_settings]
	return err
}

func SafetySettingsMulti() error {
	// [START safety_settings_multi]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	unsafePrompt := "I support Martians Soccer Club and I think Jupiterians Football Club sucks! " +
		"Write a ironic phrase about them including expletives."

	config := &genai.GenerateContentConfig{
		SafetySettings: []*genai.SafetySetting{
			{
				Category:  "HARM_CATEGORY_HATE_SPEECH",
				Threshold: "BLOCK_MEDIUM_AND_ABOVE",
			},
			{
				Category:  "HARM_CATEGORY_HARASSMENT",
				Threshold: "BLOCK_ONLY_HIGH",
			},
		},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(unsafePrompt, genai.RoleUser),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, config)
	if err != nil {
		log.Fatal(err)
	}

	// Print the generated text.
	text := response.Text()
	fmt.Println("Generated text:", text)

	// Print the and safety ratings from the first candidate.
	if len(response.Candidates) > 0 {
		fmt.Println("Finish reason:", response.Candidates[0].FinishReason)
		safetyRatings, err := json.MarshalIndent(response.Candidates[0].SafetyRatings, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println("Safety ratings:", string(safetyRatings))
	} else {
		fmt.Println("No candidate returned.")
	}
	// [END safety_settings_multi]
	return err
}
