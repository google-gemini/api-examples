package examples

import (
	"context"
	"fmt"
	"strings"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

func Chat() error {
	// [START chat]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Pass initial history using the History field.
	history := []*genai.Content{
		genai.NewContentFromText("Hello", genai.RoleUser),
		genai.NewContentFromText("Great to meet you. What would you like to know?", genai.RoleModel),
	}

	chat, err := client.Chats.Create(ctx, "gemini-2.0-flash", nil, history)
	if err != nil {
		log.Fatal(err)
	}

	firstResp, err := chat.SendMessage(ctx, genai.Part{Text: "I have 2 dogs in my house."})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(firstResp.Text())

	secondResp, err := chat.SendMessage(ctx, genai.Part{Text: "How many paws are in my house?"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(secondResp.Text())
	// [END chat]

	return nil
}

func ChatStreaming() error {
	// [START chat_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	history := []*genai.Content{
		genai.NewContentFromText("Hello", genai.RoleUser),
		genai.NewContentFromText("Great to meet you. What would you like to know?", genai.RoleModel),
	}
	chat, err := client.Chats.Create(ctx, "gemini-2.0-flash", nil, history)
	if err != nil {
		log.Fatal(err)
	}

	for chunk, err := range chat.SendMessageStream(ctx, genai.Part{Text: "I have 2 dogs in my house."}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(chunk.Text())
		fmt.Println(strings.Repeat("_", 64))
	}

	for chunk, err := range chat.SendMessageStream(ctx, genai.Part{Text: "How many paws are in my house?"}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(chunk.Text())
		fmt.Println(strings.Repeat("_", 64))
	}

	fmt.Println(chat.History(false))
	// [END chat_streaming]

	return nil
}

func ChatStreamingWithImages() error {
	// [START chat_streaming_with_images]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	chat, err := client.Chats.Create(ctx, "gemini-2.0-flash", nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for chunk, err := range chat.SendMessageStream(ctx, genai.Part{
		Text: "Hello, I'm interested in learning about musical instruments. Can I show you one?"}) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(chunk.Text())
		fmt.Println(strings.Repeat("_", 64))
	}

	image, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "organ.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Upload image file
	parts := make([]genai.Part, 2)
	parts[0] = genai.Part{Text: "What family of instruments does this instrument belong to?"}
	parts[1] = genai.Part{
		FileData: &genai.FileData{
			FileURI :      image.URI,
			MIMEType: image.MIMEType,
		},
	}

	for chunk, err := range chat.SendMessageStream(ctx, parts...) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(chunk.Text())
		fmt.Println(strings.Repeat("_", 64))
	}
	// [END chat_streaming_with_images]

	return nil
}
