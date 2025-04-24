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

func CacheCreate() (*genai.GenerateContentResponse, error) {
	// [START cache_create]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents: contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cache created:")
	fmt.Println(cache)

	// Use the cache for generating content.
	response, err := client.Models.GenerateContent(
		ctx,
		modelName,
		genai.Text("Please summarize this transcript"),
		&genai.GenerateContentConfig{
			CachedContent: cache.Name,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END cache_create]

	// Delete the cache.
	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	return response, err
}

func CacheCreateFromName() (*genai.GenerateContentResponse, error) {
	// [START cache_create_from_name]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	cacheName := cache.Name

	// Later retrieve the cache.
	cache, err = client.Caches.Get(ctx, cacheName, &genai.GetCachedContentConfig{})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Models.GenerateContent(
		ctx,
		modelName,
		genai.Text("Find a lighthearted moment from this transcript"),
		&genai.GenerateContentConfig{
			CachedContent: cache.Name,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response from cache (create from name):")
	printResponse(response)
	// [END cache_create_from_name]

	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	return response, err
}

func CacheCreateFromChat() (*genai.GenerateContentResponse, error) {
	// [START cache_create_from_chat]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	systemInstruction := "You are an expert analyzing transcripts."

	// Create initial chat with a system instruction.
	chat, err := client.Chats.Create(ctx, modelName, &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemInstruction, "user"),
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Send first message with the transcript.
	parts := make([]genai.Part, 2)
	parts[0] = genai.Part{Text: "Hi, could you summarize this transcript?"}
	parts[1] = genai.Part{
		FileData: &genai.FileData{
			FileURI :      document.URI,
			MIMEType: document.MIMEType,
		},
	}

	// Send chat message.
	resp, err := chat.SendMessage(ctx, parts...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel: ", resp.Text())

	resp, err = chat.SendMessage(
		ctx, 
		genai.Part{
			Text: "Okay, could you tell me more about the trans-lunar injection",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel: ", resp.Text())

	// To cache the conversation so far, pass the chat history as the list of contents.
	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          chat.History(false),
		SystemInstruction: genai.NewContentFromText(systemInstruction, "user"),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Continue the conversation using the cached history.
	chat, err = client.Chats.Create(ctx, modelName, &genai.GenerateContentConfig{
		CachedContent: cache.Name,
	}, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = chat.SendMessage(
		ctx, 
		genai.Part{
			Text: "I didn't understand that last part, could you explain it in simpler language?",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel: ", resp.Text())
	// [END cache_create_from_chat]

	// Clean up the cache.
	if _, err := client.Caches.Delete(ctx, cache.Name, nil); err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func CacheDelete() error {
	// [START cache_delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cache deleted:", cache.Name)
	// [END cache_delete]
	return err
}

func CacheGet() error {
	// [START cache_get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	cache, err = client.Caches.Get(ctx, cache.Name, &genai.GetCachedContentConfig{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved cache:")
	fmt.Println(cache)
	// [END cache_get]

	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	return err
}

func CacheList() error {
	// [START cache_list]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// For demonstration, create a cache first.
	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	// List caches using the List method with a page size of 2.
	page, err := client.Caches.List(ctx, &genai.ListCachedContentsConfig{PageSize: 2})
	if err != nil {
		log.Fatal(err)
	}

	pageIndex := 1
	for {
		fmt.Printf("Listing caches (page %d):\n", pageIndex)
		for _, item := range page.Items {
			fmt.Println("   ", item.Name)
		}
		if page.NextPageToken == "" {
			break
		}
		page, err = page.Next(ctx)
		if err == genai.ErrPageDone {
			break
		} else if err != nil {
			return err
		}
		pageIndex++
	}
	// [END cache_list]

	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	return err
}

func CacheUpdate() error {
	// [START cache_update]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelName := "gemini-1.5-flash-001"
	document, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "a11.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		genai.NewPartFromURI(document.URI, document.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	cache, err := client.Caches.Create(ctx, modelName, &genai.CreateCachedContentConfig{
		Contents:          contents,
		SystemInstruction: genai.NewContentFromText(
			"You are an expert analyzing transcripts.", "user",
		),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Update the TTL (2 hours).
	cache, err = client.Caches.Update(ctx, cache.Name, &genai.UpdateCachedContentConfig{
		TTL: 7200 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After update:")
	fmt.Println(cache)

	// Alternatively, update expire_time directly.
	expire := time.Now().Add(15 * time.Minute).UTC()
	cache, err = client.Caches.Update(ctx, cache.Name, &genai.UpdateCachedContentConfig{
		ExpireTime: expire,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("After expire_time update:")
	fmt.Println(cache)
	// [END cache_update]

	_, err = client.Caches.Delete(ctx, cache.Name, &genai.DeleteCachedContentConfig{})
	return err
}
