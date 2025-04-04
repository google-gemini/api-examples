package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

func TokensContextWindow() error {
	// [START tokens_context_window]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelInfo, err := client.Models.Get(ctx, "gemini-2.0-flash", &genai.GetModelConfig{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("input_token_limit=%d\n", modelInfo.InputTokenLimit)
	fmt.Printf("output_token_limit=%d\n", modelInfo.OutputTokenLimit)
	// [END tokens_context_window]
	return err
}

// TokensTextOnly counts tokens for a text prompt and prints usage metadata.
func TokensTextOnly() error {
	// [START tokens_text_only]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "The quick brown fox jumps over the lazy dog."

	// Convert prompt to a slice of *genai.Content using the helper.
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, "user"),
	}
	countResp, err := client.Models.CountTokens(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		return err
	}
	fmt.Println("total_tokens:", countResp.TotalTokens)

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usageMetadata))
	// [END tokens_text_only]
	return err
}

func TokensMultimodalImageInline() error {
	// [START tokens_multimodal_image_inline]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file.
	file, err := os.Open(filepath.Join(getMedia(), "organ.jpg"))
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Read the file.
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	parts := []*genai.Part{
		{Text: "Tell me about this image"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	tokenResp, err := client.Models.CountTokens(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Multimodal inline token count:", tokenResp.TotalTokens)

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usageMetadata))
	// [END tokens_multimodal_image_inline]
	return err
}

func TokensMultimodalVideoAudioInline() error {
	// [START tokens_multimodal_video_audio_inline]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file.
	file, err := os.Open(filepath.Join(getMedia(), "organ.jpg"))
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Read the file.
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	parts := []*genai.Part{
		{Text: "Tell me about this video"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "video/mp4"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	tokenResp, err := client.Models.CountTokens(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Multimodal video/audio token count:", tokenResp.TotalTokens)
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usageMetadata))
	// [END tokens_multimodal_video_audio_inline]
	return err
}

func TokensMultimodalPdfInline() error {
	// [START tokens_multimodal_pdf_inline]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file.
	file, err := os.Open(filepath.Join(getMedia(), "test.pdf"))
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Read the file.
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	parts := []*genai.Part{
		{Text: "Give me a summary of this document."},
		{InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	tokenResp, err := client.Models.CountTokens(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Multimodal PDF token count: %d\n", tokenResp.TotalTokens)
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(usageMetadata))
	// [END tokens_multimodal_pdf_inline]
	return err
}

func TokensCachedContent() error {
	// [START tokens_cached_content]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Open the file.
	file, err := os.Open(filepath.Join(getMedia(), "a11.txt"))
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Read the file.
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	parts := []*genai.Part{
		{Text: "Here the Apollo 11 transcript:"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	// Create cached content using a simple slice with text and a file.
	cache, err := client.Caches.Create(ctx, "gemini-1.5-flash-001", &genai.CreateCachedContentConfig{
		Contents: contents,
	})
	if err != nil {
		log.Fatal(err)
	}

	prompt := "Please give a short summary of this file."
	countResp, err := client.Models.CountTokens(ctx, "gemini-2.0-flash", []*genai.Content{
		genai.NewContentFromText(prompt, "user"),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", countResp.TotalTokens)
	response, err := client.Models.GenerateContent(ctx, "gemini-1.5-flash-001", []*genai.Content{
		genai.NewContentFromText(prompt, "user"),
	}, &genai.GenerateContentConfig{
		CachedContent: cache.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	usageMetadata, err := json.MarshalIndent(response.UsageMetadata, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	// Returns `nil` for some reason
	fmt.Println(string(usageMetadata))
	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	// [END tokens_cached_content]
	return err
}
