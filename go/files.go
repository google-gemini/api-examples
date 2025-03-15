/**
 * @license
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
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
	"os"
	"path/filepath"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Helper for getting the path to the media directory
func getMediaDir() string {
	// Get the current file's directory
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// Navigate to the third_party directory
	return filepath.Join(currentDir, "..", "third_party")
}

// Helper for sleeping (used in video polling)
func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// FilesCreateText demonstrates uploading a text file and using it with Gemini
func FilesCreateText() (string, error) {
	// [START files_create_text]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "poem.txt"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "text/plain")
	if err != nil {
		return "", err
	}
	fmt.Printf("Uploaded file: %v\n", myfile)

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("\n\nCan you add a few more lines to this poem?"), genai.FileData{
		MimeType: "text/plain",
		FileURI:  myfile.URI,
	})
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	fmt.Printf("result.text= %s\n", result)
	// [END files_create_text]
	return result, nil
}

// FilesCreateImage demonstrates uploading an image file and using it with Gemini
func FilesCreateImage() (string, error) {
	// [START files_create_image]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "Cajun_instruments.jpg"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "image/jpeg")
	if err != nil {
		return "", err
	}
	fmt.Printf("Uploaded file: %v\n", myfile)

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.FileData{
		MimeType: "image/jpeg",
		FileURI:  myfile.URI,
	}, genai.Text("\n\nCan you tell me about the instruments in this photo?"))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	fmt.Printf("result.text= %s\n", result)
	// [END files_create_image]
	return result, nil
}

// FilesCreateAudio demonstrates uploading an audio file and using it with Gemini
func FilesCreateAudio() (string, error) {
	// [START files_create_audio]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "sample.mp3"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "audio/mpeg")
	if err != nil {
		return "", err
	}
	fmt.Printf("Uploaded file: %v\n", myfile)

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.FileData{
		MimeType: "audio/mpeg",
		FileURI:  myfile.URI,
	}, genai.Text("Describe this audio clip"))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	fmt.Printf("result.text= %s\n", result)
	// [END files_create_audio]
	return result, nil
}

// FilesCreateVideo demonstrates uploading a video file and using it with Gemini
func FilesCreateVideo() (string, error) {
	// [START files_create_video]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "Big_Buck_Bunny.mp4"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "video/mp4")
	if err != nil {
		return "", err
	}
	fmt.Printf("Uploaded video file: %v\n", myfile)

	// Poll until the video file is completely processed (state becomes ACTIVE)
	for myfile.State != "ACTIVE" {
		fmt.Println("Processing video...")
		fmt.Printf("File state: %s\n", myfile.State)
		sleep(5000)
		myfile, err = client.Files.Get(ctx, myfile.Name)
		if err != nil {
			return "", err
		}
	}

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.FileData{
		MimeType: "video/mp4",
		FileURI:  myfile.URI,
	}, genai.Text("Describe this video clip"))
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	fmt.Printf("result.text= %s\n", result)
	// [END files_create_video]
	return result, nil
}

// FilesCreatePdf demonstrates uploading a PDF file and using it with Gemini
func FilesCreatePdf() (string, error) {
	// [START files_create_pdf]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "test.pdf"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	samplePdf, err := client.Files.Upload(ctx, file, "application/pdf")
	if err != nil {
		return "", err
	}

	model := client.GenerativeModel("gemini-2.0-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("Give me a summary of this pdf file."), genai.FileData{
		MimeType: "application/pdf",
		FileURI:  samplePdf.URI,
	})
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])
	fmt.Printf("Result text: %s\n", result)
	// [END files_create_pdf]
	return result, nil
}

// FilesList demonstrates listing all files
func FilesList() ([]string, error) {
	// [START files_list]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	fmt.Println("My files:")
	files, err := client.Files.List(ctx)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, f := range files {
		fmt.Printf("  %s\n", f.Name)
		names = append(names, f.Name)
	}
	// [END files_list]
	return names, nil
}

// FilesGet demonstrates retrieving a specific file by name
func FilesGet() (string, error) {
	// [START files_get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "poem.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "text/plain")
	if err != nil {
		panic(err)
	}
	fileName := myfile.Name
	fmt.Println(fileName)

	fetchedFile, err := client.Files.Get(ctx, fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(fetchedFile)
	// [END files_get]
}

// FilesDelete demonstrates deleting a file
func FilesDelete() {
	// [START files_delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	mediaDir := getMediaDir()
	file, err := os.Open(filepath.Join(mediaDir, "poem.txt"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	myfile, err := client.Files.Upload(ctx, file, "text/plain")
	if err != nil {
		panic(err)
	}

	err = client.Files.Delete(ctx, myfile.Name)
	if err != nil {
		panic(err)
	}

	// Try to use the deleted file (should fail)
	model := client.GenerativeModel("gemini-2.0-flash")
	_, err = model.GenerateContent(ctx, genai.FileData{
		MimeType: "text/plain",
		FileURI:  myfile.URI,
	}, genai.Text("Describe this file."))
	
	if err != nil {
		fmt.Println("Error using deleted file (expected):", err)
	}
	// [END files_delete]
}

// Each function in this file can be called independently
// and returns appropriate responses for testing
