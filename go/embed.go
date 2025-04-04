package examples

import (
	"context"
	"fmt"
	"os"
	"log"
	"encoding/json"

	"google.golang.org/genai"
)

func EmbedContent() error {
	// [START embed_content]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	text := "Hello World!"
	outputDim := int32(10)
	contents := []*genai.Content{
		genai.NewContentFromText(text, "user"),
	}
	result, err := client.Models.EmbedContent(ctx, "text-embedding-004", 
		contents, &genai.EmbedContentConfig{
			OutputDimensionality: &outputDim,
	})
	if err != nil {
		log.Fatal(err)
	}

	embeddings, err := json.MarshalIndent(result.Embeddings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(embeddings))
	// [END embed_content]
	return err
}

func BatchEmbedContents() error {
	// [START batch_embed_contents]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	contents := []*genai.Content{
		genai.NewContentFromText("What is the meaning of life?", "user"),
		genai.NewContentFromText("How much wood would a woodchuck chuck?", "user"),
		genai.NewContentFromText("How does the brain work?", "user"),
	}

	outputDim := int32(10)
	result, err := client.Models.EmbedContent(ctx, "text-embedding-004", contents, &genai.EmbedContentConfig{
		OutputDimensionality: &outputDim,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	embeddings, err := json.MarshalIndent(result.Embeddings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(embeddings))
	// [END batch_embed_contents]
	return err
}
