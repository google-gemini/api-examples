/**
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package examples

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/genai"
)

// TokensTextOnly demonstrates counting tokens for text-only content
func TokensTextOnly() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_text_only]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	prompt := "The quick brown fox jumps over the lazy dog."

	// Count tokens for the prompt
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(countTokensResponse.TotalTokens)

	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_text_only]

	return countTokensResponse, generateResponse, nil
}

// TokensChat demonstrates counting tokens for chat content
func TokensChat() (*genai.CountTokensResponse, *genai.GenerateContentResponse, *genai.CountTokensResponse, error) {
	// [START tokens_chat]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Initial chat history
	history := []*genai.Content{
		{
			Role:  "user",
			Parts: []genai.Part{genai.Text("Hi my name is Bob")},
		},
		{
			Role:  "model",
			Parts: []genai.Part{genai.Text("Hi Bob!")},
		},
	}

	chat := client.Chats.Create(ctx, "gemini-2.0-flash", history)

	// Count tokens for the current chat history
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		chat.GetHistory(),
		nil,
	)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println(countTokensResponse.TotalTokens)

	// Send a new message and get usage metadata
	chatResponse, err := chat.SendMessage(
		ctx,
		genai.Text("In one sentence, explain how a computer works to a young child."),
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", chatResponse.UsageMetadata)

	// Add an extra user message to the history
	extraMessage := &genai.Content{
		Role:  "user",
		Parts: []genai.Part{genai.Text("What is the meaning of life?")},
	}
	combinedHistory := chat.GetHistory()
	combinedHistory = append(combinedHistory, extraMessage)

	// Count tokens for the combined history
	combinedCountTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		combinedHistory,
		nil,
	)
	if err != nil {
		return countTokensResponse, chatResponse, nil, err
	}
	fmt.Println("Combined history token count:", combinedCountTokensResponse.TotalTokens)
	// [END tokens_chat]

	return countTokensResponse, chatResponse, combinedCountTokensResponse, nil
}

// TokensMultimodalImageInline demonstrates counting tokens for content with inline images
func TokensMultimodalImageInline() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_multimodal_image_inline]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	prompt := "Tell me about this image"
	
	// Get the absolute path to the media directory
	mediaDir := filepath.Join("..", "third_party")
	imagePath := filepath.Join(mediaDir, "organ.jpg")
	
	// Read the image file
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, nil, err
	}
	
	// Create image part
	imagePart := genai.ImageData("image/jpeg", imageBytes)
	
	// Count tokens for the combined text and image
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), imagePart},
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(countTokensResponse.TotalTokens)
	
	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), imagePart},
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_multimodal_image_inline]
	
	return countTokensResponse, generateResponse, nil
}

// TokensMultimodalImageFileAPI demonstrates counting tokens for content with images using the File API
func TokensMultimodalImageFileAPI() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_multimodal_image_file_api]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	prompt := "Tell me about this image"
	
	// Get the absolute path to the media directory
	mediaDir := filepath.Join("..", "third_party")
	imagePath := filepath.Join(mediaDir, "organ.jpg")
	
	// Upload the file
	file, err := client.Files.Upload(ctx, imagePath, &genai.FileConfig{
		MimeType: "image/jpeg",
	})
	if err != nil {
		return nil, nil, err
	}
	
	// Count tokens for the combined text and file
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), genai.FileData(file.URI, file.MimeType)},
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(countTokensResponse.TotalTokens)
	
	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), genai.FileData(file.URI, file.MimeType)},
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_multimodal_image_file_api]
	
	return countTokensResponse, generateResponse, nil
}

// TokensMultimodalVideoAudioFileAPI demonstrates counting tokens for content with video/audio using the File API
func TokensMultimodalVideoAudioFileAPI() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_multimodal_video_audio_file_api]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	prompt := "Tell me about this video"
	
	// Get the absolute path to the media directory
	mediaDir := filepath.Join("..", "third_party")
	videoPath := filepath.Join(mediaDir, "Big_Buck_Bunny.mp4")
	
	// Upload the file
	videoFile, err := client.Files.Upload(ctx, videoPath, &genai.FileConfig{
		MimeType: "video/mp4",
	})
	if err != nil {
		return nil, nil, err
	}
	
	// Poll until the video file is completely processed (state becomes ACTIVE)
	for videoFile.State != "ACTIVE" {
		fmt.Println("Processing video...")
		fmt.Println("File state:", videoFile.State)
		time.Sleep(5 * time.Second)
		videoFile, err = client.Files.Get(ctx, videoFile.Name)
		if err != nil {
			return nil, nil, err
		}
	}
	
	// Count tokens for the combined text and video
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), genai.FileData(videoFile.URI, videoFile.MimeType)},
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(countTokensResponse.TotalTokens)
	
	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text(prompt), genai.FileData(videoFile.URI, videoFile.MimeType)},
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_multimodal_video_audio_file_api]
	
	return countTokensResponse, generateResponse, nil
}

// TokensMultimodalPdfFileAPI demonstrates counting tokens for content with PDF files using the File API
func TokensMultimodalPdfFileAPI() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_multimodal_pdf_file_api]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Get the absolute path to the media directory
	mediaDir := filepath.Join("..", "third_party")
	pdfPath := filepath.Join(mediaDir, "test.pdf")
	
	// Upload the PDF file
	pdfFile, err := client.Files.Upload(ctx, pdfPath, &genai.FileConfig{
		MimeType: "application/pdf",
	})
	if err != nil {
		return nil, nil, err
	}
	
	// Count tokens for the combined text and PDF
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text("Give me a summary of this document."), genai.FileData(pdfFile.URI, pdfFile.MimeType)},
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Token count: %d\n", countTokensResponse.TotalTokens)
	
	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		[]genai.Part{genai.Text("Give me a summary of this document."), genai.FileData(pdfFile.URI, pdfFile.MimeType)},
		nil,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_multimodal_pdf_file_api]
	
	return countTokensResponse, generateResponse, nil
}

// TokensSystemInstruction demonstrates counting tokens for content with system instructions
func TokensSystemInstruction() (*genai.CountTokensResponse, *genai.GenerateContentResponse, error) {
	// [START tokens_system_instruction]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	systemInstruction := "You are a helpful assistant that speaks like a pirate."
	prompt := "Tell me about the weather today."
	
	// Create generation config with system instruction
	genConfig := &genai.GenerationConfig{
		SystemInstruction: systemInstruction,
	}
	
	// Count tokens for the prompt with system instruction
	countTokensResponse, err := client.Models.CountTokens(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		genConfig,
	)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("Total tokens with system instruction: %d\n", countTokensResponse.TotalTokens)
	
	// Generate content and get usage metadata
	generateResponse, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(prompt),
		genConfig,
	)
	if err != nil {
		return countTokensResponse, nil, err
	}
	fmt.Printf("Usage Metadata: %+v\n", generateResponse.UsageMetadata)
	// [END tokens_system_instruction]
	
	return countTokensResponse, generateResponse, nil
}
