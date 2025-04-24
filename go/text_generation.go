package examples

import (
	"context"
	"fmt"
	"log"
	"time"
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "organ.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromText("Tell me about this instrument"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
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
	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "organ.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromText("Tell me about this instrument"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
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

	organ, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "organ.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	cajun, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Cajun_instruments.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("What is the difference between both of these instruments?"),
		genai.NewPartFromURI(organ.URI, organ.MIMEType),
		genai.NewPartFromURI(cajun.URI, cajun.MIMEType),
	}

	contents := []*genai.Content{
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

	organ, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "organ.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	cajun, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Cajun_instruments.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("What is the difference between both of these instruments?"),
		genai.NewPartFromURI(organ.URI, organ.MIMEType),
		genai.NewPartFromURI(cajun.URI, cajun.MIMEType),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "sample.mp3"), 
		&genai.UploadFileConfig{
			MIMEType : "audio/mpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this audio file."),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "sample.mp3"), 
		&genai.UploadFileConfig{
			MIMEType : "audio/mpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this audio file."),
		genai.NewPartFromURI(file.URI, file.MIMEType),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Big_Buck_Bunny.mp4"), 
		&genai.UploadFileConfig{
			MIMEType : "video/mp4",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Poll until the video file is completely processed (state becomes ACTIVE).
	for file.State == genai.FileStateUnspecified || file.State != genai.FileStateActive {
		fmt.Println("Processing video...")
		fmt.Println("File state:", file.State)
		time.Sleep(5 * time.Second)

		file, err = client.Files.Get(ctx, file.Name, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Describe this video clip"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Big_Buck_Bunny.mp4"), 
		&genai.UploadFileConfig{
			MIMEType : "video/mp4",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Poll until the video file is completely processed (state becomes ACTIVE).
	for file.State == genai.FileStateUnspecified || file.State != genai.FileStateActive {
		fmt.Println("Processing video...")
		fmt.Println("File state:", file.State)
		time.Sleep(5 * time.Second)

		file, err = client.Files.Get(ctx, file.Name, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Describe this video clip"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "test.pdf"), 
		&genai.UploadFileConfig{
			MIMEType : "application/pdf",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this document:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
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

	file, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "test.pdf"), 
		&genai.UploadFileConfig{
			MIMEType : "application/pdf",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this document:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
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
	// [END text_gen_multimodal_pdf_streaming]
	return err
}
