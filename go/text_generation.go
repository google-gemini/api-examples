package examples

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

func TextGenTextOnlyPrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_text_only_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Write a story about a magic backpack.", "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_text_only_prompt]
	return response, err
}

func TextGenTextOnlyPromptStreaming() error {
	// [START text_gen_text_only_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Write a story about a magic backpack.", "user"),
	}
	for response, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(response.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_text_only_prompt_streaming]
	return err
}

func TextGenMultimodalOneImagePrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_one_image_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	imagePath := filepath.Join(getMedia(), "organ.jpg")
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	imagePart := &genai.Part{
		InlineData: &genai.Blob{
			Data:     data,
			MIMEType: "image/jpeg",
		},
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Tell me about this instrument", "user"),
		genai.NewContentFromParts([]*genai.Part{imagePart}, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_multimodal_one_image_prompt]
	return response, err
}

func TextGenMultimodalOneImagePromptStreaming() error {
	// [START text_gen_multimodal_one_image_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	imagePath := filepath.Join(getMedia(), "organ.jpg")
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	imagePart := &genai.Part{
		InlineData: &genai.Blob{
			Data:     data,
			MIMEType: "image/jpeg",
		},
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Tell me about this instrument", "user"),
		genai.NewContentFromParts([]*genai.Part{imagePart}, "user"),
	}
	for response, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(response.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_one_image_prompt_streaming]
	return err
}

func TextGenMultimodalMultiImagePrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_multi_image_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	organPath := filepath.Join(getMedia(), "organ.jpg")
	cajunPath := filepath.Join(getMedia(), "Cajun_instruments.jpg")
	organFile, err := os.Open(organPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer organFile.Close()
	organData, err := io.ReadAll(organFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	cajunFile, err := os.Open(cajunPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer cajunFile.Close()
	cajunData, err := io.ReadAll(cajunFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	parts := []*genai.Part{
		{InlineData: &genai.Blob{Data: organData, MIMEType: "image/jpeg"}},
		{InlineData: &genai.Blob{Data: cajunData, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(
			"What is the difference between both of these instruments?", "user",
		),
		genai.NewContentFromParts(parts, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_multimodal_multi_image_prompt]
	return response, err
}

func TextGenMultimodalMultiImagePromptStreaming() error {
	// [START text_gen_multimodal_multi_image_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	organPath := filepath.Join(getMedia(), "organ.jpg")
	cajunPath := filepath.Join(getMedia(), "Cajun_instruments.jpg")
	organFile, err := os.Open(organPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer organFile.Close()
	organData, err := io.ReadAll(organFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	cajunFile, err := os.Open(cajunPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer cajunFile.Close()
	cajunData, err := io.ReadAll(cajunFile)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	parts := []*genai.Part{
		{InlineData: &genai.Blob{Data: organData, MIMEType: "image/jpeg"}},
		{InlineData: &genai.Blob{Data: cajunData, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(
			"What is the difference between both of these instruments?", "user",
		),
		genai.NewContentFromParts(parts, "user"),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_multi_image_prompt_streaming]
	return err
}

func TextGenMultimodalAudio() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_audio]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "Give me a summary of this audio file."
	audioPath := filepath.Join(getMedia(), "sample.mp3")
	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	audioPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "audio/mpeg"},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, "user"),
		genai.NewContentFromParts([]*genai.Part{audioPart}, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_multimodal_audio]
	return response, err
}

func TextGenMultimodalAudioStreaming() error {
	// [START text_gen_multimodal_audio_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "Give me a summary of this audio file."
	audioPath := filepath.Join(getMedia(), "sample.mp3")
	file, err := os.Open(audioPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	audioPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "audio/mpeg"},
	}
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, "user"),
		genai.NewContentFromParts([]*genai.Part{audioPart}, "user"),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_audio_streaming]
	return err
}

func TextGenMultimodalVideoPrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_video_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	videoPath := filepath.Join(getMedia(), "Big_Buck_Bunny.mp4")
	file, err := os.Open(videoPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	parts := []*genai.Part{
		{Text: "Describe this video clip"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "video/mp4"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_multimodal_video_prompt]
	return response, err
}

func TextGenMultimodalVideoPromptStreaming() error {
	// [START text_gen_multimodal_video_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	videoPath := filepath.Join(getMedia(), "Big_Buck_Bunny.mp4")
	file, err := os.Open(videoPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	parts := []*genai.Part{
		{Text: "Describe this video clip"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "video/mp4"}},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_video_prompt_streaming]
	return err
}

func TextGenMultimodalPdf() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_pdf]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	pdfPath := filepath.Join(getMedia(), "test.pdf")
	file, err := os.Open(pdfPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	pdfPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"},
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Give me a summary of this document:", "user"),
		genai.NewContentFromParts([]*genai.Part{pdfPart}, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_multimodal_pdf]
	return response, err
}

func TextGenMultimodalPdfStreaming() error {
	// [START text_gen_multimodal_pdf_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	pdfPath := filepath.Join(getMedia(), "test.pdf")
	file, err := os.Open(pdfPath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	pdfPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"},
	}
	contents := []*genai.Content{
		genai.NewContentFromText("Give me a summary of this document:", "user"),
		genai.NewContentFromParts([]*genai.Part{pdfPart}, "user"),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_pdf_streaming]
	return err
}
