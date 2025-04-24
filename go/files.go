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

func FilesCreateText() (*genai.GenerateContentResponse, error) {
	// [START files_create_text]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	myfile, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "poem.txt"), 
		&genai.UploadFileConfig{
			MIMEType : "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("myfile=%+v\n", myfile)

	parts := []*genai.Part{
		genai.NewPartFromURI(myfile.URI, myfile.MIMEType),
		genai.NewPartFromText("\n\n"),
		genai.NewPartFromText("Can you add a few more lines to this poem?"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Printf("result.text=%s\n", text)
	// [END files_create_text]
	return response, err
}

func FilesCreateImage() (*genai.GenerateContentResponse, error) {
	// [START files_create_image]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	myfile, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Cajun_instruments.jpg"), 
		&genai.UploadFileConfig{
			MIMEType : "image/jpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("myfile=%+v\n", myfile)

	parts := []*genai.Part{
		genai.NewPartFromURI(myfile.URI, myfile.MIMEType),
		genai.NewPartFromText("\n\n"),
		genai.NewPartFromText("Can you tell me about the instruments in this photo?"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Printf("result.text=%s\n", text)
	// [END files_create_image]
	return response, err
}

func FilesCreateAudio() (*genai.GenerateContentResponse, error) {
	// [START files_create_audio]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	myfile, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "sample.mp3"), 
		&genai.UploadFileConfig{
			MIMEType : "audio/mpeg",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("myfile=%+v\n", myfile)

	parts := []*genai.Part{
		genai.NewPartFromURI(myfile.URI, myfile.MIMEType),
		genai.NewPartFromText("Describe this audio clip"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Printf("result.text=%s\n", text)
	// [END files_create_audio]
	return response, err
}

func FilesCreateVideo() (*genai.GenerateContentResponse, error) {
	// [START files_create_video]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	myfile, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "Big_Buck_Bunny.mp4"), 
		&genai.UploadFileConfig{
			MIMEType : "video/mp4",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("myfile=%+v\n", myfile)

	// Poll until the video file is completely processed (state becomes ACTIVE).
	for myfile.State == genai.FileStateUnspecified || myfile.State != genai.FileStateActive {
		fmt.Println("Processing video...")
		fmt.Println("File state:", myfile.State)
		time.Sleep(5 * time.Second)

		myfile, err = client.Files.Get(ctx, myfile.Name, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	parts := []*genai.Part{
		genai.NewPartFromURI(myfile.URI, myfile.MIMEType),
		genai.NewPartFromText("Describe this video clip"),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Printf("result.text=%s\n", text)
	// [END files_create_video]
	return response, err
}

func FilesCreatePdf() (*genai.GenerateContentResponse, error) {
	// [START files_create_pdf]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	samplePdf, err := client.Files.UploadFromPath(
		ctx,
		filepath.Join(getMedia(), "test.pdf"),
		&genai.UploadFileConfig{
			MIMEType: "application/pdf",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this pdf file."),
		genai.NewPartFromURI(samplePdf.URI, samplePdf.MIMEType),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Println(text)
	// [END files_create_pdf]
	return response, err
}

func FilesCreateFromIO() (*genai.GenerateContentResponse, error) {
	// [START files_create_io]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(filepath.Join(getMedia(), "test.pdf"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	samplePdf, err := client.Files.Upload(ctx, f, &genai.UploadFileConfig{
		MIMEType : "application/pdf",
	})
	if err != nil {
		log.Fatal(err)
	}

	parts := []*genai.Part{
		genai.NewPartFromText("Give me a summary of this pdf file."),
		genai.NewPartFromURI(samplePdf.URI, samplePdf.MIMEType),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text := response.Text()
	fmt.Println(text)
	// [END files_create_io]
	return response, err
}

func FilesList() error {
	// [START files_list]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("My files:")
	page, err := client.Files.List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range page.Items {
		fmt.Println("  ", f.Name)
	}
	// [END files_list]
	return nil
}

func FilesGet() (*genai.File, error) {
	// [START files_get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	myfile, err := client.Files.UploadFromPath(
		ctx,
		filepath.Join(getMedia(), "poem.txt"), 
		&genai.UploadFileConfig{
			MIMEType: "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fileName := myfile.Name
	fmt.Println(fileName)
	file, err := client.Files.Get(ctx, fileName, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
	// [END files_get]
	return file, err
}

func FilesDelete() error {
	// [START files_delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	myfile, err := client.Files.UploadFromPath(
		ctx, 
		filepath.Join(getMedia(), "poem.txt"), 
		&genai.UploadFileConfig{
			MIMEType: "text/plain",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	// Delete the file.
	_, err = client.Files.Delete(ctx, myfile.Name, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Attempt to use the deleted file.
	parts := []*genai.Part{
		genai.NewPartFromURI(myfile.URI, myfile.MIMEType,),
		genai.NewPartFromText("Describe this file."),
	}

	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	
	_, err = client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	// Expect an error when using a deleted file.
	if err != nil {
		return nil
	}
	return fmt.Errorf("expected an error when using deleted file")
	// [END files_delete]
}
