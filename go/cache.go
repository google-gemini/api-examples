package examples

import (
	"context"
	"fmt"
	"io"
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
		{Text: "Please summarize this transcript"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
	fmt.Println("Generated content:")
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
		{Text: "Please summarize this transcript"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
		{InlineData: &genai.Blob{Data: data, MIMEType: "text/plain"}},
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
	fmt.Println("After TTL update:")
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
