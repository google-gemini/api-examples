package examples

import (
	"context"
	"path/filepath"
	"os"
	"log"

	"google.golang.org/genai"
)

func JsonControlledGeneration() (*genai.GenerateContentResponse, error) {
	// [START json_controlled_generation]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	schema := &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"recipe_name": {Type: genai.TypeString},
				"ingredients": {
					Type:  genai.TypeArray,
					Items: &genai.Schema{Type: genai.TypeString},
				},
			},
			Required: []string{"recipe_name"},
		},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   schema,
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("List a few popular cookie recipes."),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END json_controlled_generation]
	return response, err
}

func JsonNoSchema() (*genai.GenerateContentResponse, error) {
	// [START json_no_schema]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "List a few popular cookie recipes in JSON format.\n\n" +
			  "Use this JSON schema:\n\n" +
			  "Recipe = {'recipe_name': str, 'ingredients': list[str]}\n" +
		      "Return: list[Recipe]"
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END json_no_schema]
	return response, err
}

func JsonEnum() (*genai.GenerateContentResponse, error) {
	// [START json_enum]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Choice is a custom type representing a musical instrument category.
	type Choice string

	const (
		Percussion Choice = "Percussion"
		String     Choice = "String"
		Woodwind   Choice = "Woodwind"
		Brass      Choice = "Brass"
		Keyboard   Choice = "Keyboard"
	)

	// Define a schema restricting the response to the allowed Choice enum values.
	schema := &genai.Schema{
		Type: genai.TypeString,
		Enum: []string{
			string(Percussion),
			string(String),
			string(Woodwind),
			string(Brass),
			string(Keyboard),
		},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   schema,
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
		genai.NewPartFromText("What kind of instrument is this:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END json_enum]
	return response, err
}

func EnumInJson() (*genai.GenerateContentResponse, error) {
	// [START enum_in_json]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// We use a schema representing an array of objects.
	schema := &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"recipe_name": {Type: genai.TypeString},
				"grade":       {Type: genai.TypeString, Enum: []string{"a+", "a", "b", "c", "d", "f"}},
			},
			Required: []string{"recipe_name", "grade"},
		},
	}
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   schema,
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		genai.Text("List about 10 cookie recipes, grade them based on popularity"),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	// Expected output: a JSON-parsed list with recipe names and grades (e.g., "a+")
	printResponse(response)
	// [END enum_in_json]
	return response, err
}

func JsonEnumRaw() (*genai.GenerateContentResponse, error) {
	// [START json_enum_raw]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"), 
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeString,
			Enum: []string{"Percussion", "String", "Woodwind", "Brass", "Keyboard"},
		},
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
		genai.NewPartFromText("What kind of instrument is this:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}

	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	printResponse(response)
	// [END json_enum_raw]
	return response, err
}

func XEnum() (*genai.GenerateContentResponse, error) {
	// [START x_enum]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Choice is a custom type representing a musical instrument category.
	type Choice string

	const (
		Percussion Choice = "Percussion"
		String     Choice = "String"
		Woodwind   Choice = "Woodwind"
		Brass      Choice = "Brass"
		Keyboard   Choice = "Keyboard"
	)

	// Define a schema restricting the response to the allowed Choice enum values.
	schema := &genai.Schema{
		Type: genai.TypeString,
		Enum: []string{
			string(Percussion),
			string(String),
			string(Woodwind),
			string(Brass),
			string(Keyboard),
		},
	}

	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "text/x.enum",
		ResponseSchema:   schema,
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
		genai.NewPartFromText("What kind of instrument is this:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END x_enum]
	return response, err
}

func XEnumRaw() (*genai.GenerateContentResponse, error) {
	// [START x_enum_raw]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	rawSchema := &genai.Schema{
		Type: genai.TypeString,
		Enum: []string{"Percussion", "String", "Woodwind", "Brass", "Keyboard"},
	}
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "text/x.enum",
		ResponseSchema:   rawSchema,
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
		genai.NewPartFromText("What kind of instrument is this:"),
		genai.NewPartFromURI(file.URI, file.MIMEType),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, "user"),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash",
		contents,
		config,
	)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// Expected output: "Keyboard"
	// [END x_enum_raw]
	return response, err
}
