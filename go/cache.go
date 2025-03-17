// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package examples contains code samples that demonstrate the usage of the Gemini API.
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

// CacheCreate demonstrates how to create a cache with a file and use it for content generation.
// It uploads a document, creates a cache with the document, and then uses the cache
// for generating content with the model.
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

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cache created:", cache.Name)

	// Use the cache for content generation
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

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: cache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return response, nil
}

// CacheCreateFromName demonstrates how to retrieve a previously created cache by name
// and use it for content generation. This is useful when you want to reuse a cache
// across different sessions or components of your application.
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

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	cacheName := cache.Name // Save the name for later

	// Later retrieve the cache
	retrievedCache, err := client.Caches.Get(ctx, &genai.GetCacheRequest{Name: cacheName})
	if err != nil {
		log.Fatal(err)
	}

	// Use the retrieved cache for content generation
	response, err := client.Models.GenerateContent(
		ctx,
		modelName,
		genai.Text("Find a lighthearted moment from this transcript"),
		&genai.GenerateContentConfig{
			CachedContent: retrievedCache.Name,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END cache_create_from_name]

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: retrievedCache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return response, nil
}

// CacheCreateFromChat demonstrates how to create a cache from chat history and use it
// to continue a conversation. This allows you to preserve context from previous
// interactions and maintain continuity in multi-turn conversations.
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

	// Create a chat session with the system instruction
	chat, err := client.Chats.Create(ctx, &genai.CreateChatRequest{
		Model: modelName,
		Config: &genai.GenerateContentConfig{
			SystemInstruction: systemInstruction,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)

	// Send messages to the chat
	response, err := chat.SendMessage(ctx, &genai.SendMessageRequest{
		Message: []genai.Part{
			genai.Text("Hi, could you summarize this transcript?"),
			document,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel:", response.Text())

	response, err = chat.SendMessage(ctx, &genai.SendMessageRequest{
		Message: genai.Text("Okay, could you tell me more about the trans-lunar injection"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel:", response.Text())

	// To cache the conversation so far, pass the chat history as the list of contents
	chatHistory, err := chat.GetHistory(ctx)
	if err != nil {
		log.Fatal(err)
	}

	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          chatHistory,
			SystemInstruction: systemInstruction,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Continue the conversation using the cached content
	chatWithCache, err := client.Chats.Create(ctx, &genai.CreateChatRequest{
		Model: modelName,
		Config: &genai.GenerateContentConfig{
			CachedContent: cache.Name,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err = chatWithCache.SendMessage(ctx, &genai.SendMessageRequest{
		Message: genai.Text("I didn't understand that last part, could you explain it in simpler language?"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nmodel:", response.Text())
	// [END cache_create_from_chat]

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: cache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return response, nil
}

// CacheDelete demonstrates how to delete a cache when it's no longer needed.
// This is important for managing resources and maintaining privacy by removing
// cached content that is no longer required.
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

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Delete the cache
	err = client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: cache.Name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cache deleted:", cache.Name)
	// [END cache_delete]

	return nil
}

// CacheGet demonstrates how to retrieve a cache by its name. This allows you
// to access a cache's details and metadata after it has been created.
func CacheGet() (*genai.Cache, error) {
	// [START cache_get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get the cache by name
	retrievedCache, err := client.Caches.Get(ctx, &genai.GetCacheRequest{Name: cache.Name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved cache:", retrievedCache)
	// [END cache_get]

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: cache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return retrievedCache, nil
}

// CacheList demonstrates how to list all available caches. This is useful for
// inventory management and discovering what caches are available for use.
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

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// List all caches
	fmt.Println("My caches:")
	caches, err := client.Caches.List(ctx, &genai.ListCachesRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range caches {
		fmt.Println("    ", c.Name)
	}
	// [END cache_list]

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: cache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return nil
}

// CacheUpdate demonstrates how to update a cache's properties, specifically
// its time-to-live (TTL) value. This allows you to extend or shorten the
// lifespan of a cache based on your application's needs.
func CacheUpdate() (*genai.Cache, error) {
	// [START cache_update]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Upload a file
	mediaPath, err := filepath.Abs("../third_party/a11.txt")
	if err != nil {
		log.Fatal(err)
	}
	document, err := client.Files.Upload(ctx, &genai.FileUploadConfig{
		File:     mediaPath,
		MimeType: "text/plain",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Uploaded file name:", document.Name)
	modelName := "gemini-1.5-flash-001"

	// Create a cache with the document
	cache, err := client.Caches.Create(ctx, &genai.CreateCacheRequest{
		Model: modelName,
		Config: &genai.CreateCachedContentConfig{
			Contents:          []genai.Part{document},
			SystemInstruction: "You are an expert analyzing transcripts.",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Update the cache's time-to-live (TTL)
	ttlDuration := 2 * time.Hour
	ttl := fmt.Sprintf("%ds", int(ttlDuration.Seconds()))

	updatedCache, err := client.Caches.Update(ctx, &genai.UpdateCacheRequest{
		Name: cache.Name,
		Cache: &genai.Cache{
			Ttl: ttl,
		},
		UpdateMask: "ttl",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated cache TTL to:", updatedCache.Ttl)
	// [END cache_update]

	// Clean up
	if err := client.Caches.Delete(ctx, &genai.DeleteCacheRequest{Name: updatedCache.Name}); err != nil {
		log.Printf("Failed to delete cache: %v", err)
	}

	return updatedCache, nil
}
