package examples

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"

	"google.golang.org/genai"
)

// Helper for getting the path to the media directory
func getMedia() string {
	// runtime.Caller returns information about the caller.
	// 0 identifies the getMedia function itself.
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Unable to determine caller information")
	}
	// file is the full path to this source file.
	dir := filepath.Dir(file)
	// Adjust the relative path as needed.
	return filepath.Join(dir, "..", "third_party")
}

// Helping for printing the response.
func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part.Text)
			}
		}
	}
}
