package examples

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

// Define the thinking model centrally
const modelID = "gemini-2.5-pro-exp-03-25"

// Helper function to initialize the client
func newGenAIClient(ctx context.Context) (*genai.Client, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set")
	}
	return genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI, // Assuming Gemini API backend
	})
}

func ThinkingTextOnlyPrompt() (*genai.GenerateContentResponse, error) {
	// [START thinking_text_only_prompt]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	prompt := "Explain the concept of Occam's Razor and provide a simple, everyday example."
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, nil)
	if err != nil {
		log.Printf("GenerateContent failed: %v", err)
		return nil, err
	}

	fmt.Println(resp.Text())
	// [END thinking_text_only_prompt]
	return resp, nil
}

func ThinkingTextOnlyPromptStreaming() (string, error) {
	// [START thinking_text_only_prompt_streaming]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return "", err
	}

	prompt := "Explain the concept of Occam's Razor and provide a simple, everyday example."
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	var fullResponse strings.Builder
	stream := client.Models.GenerateContentStream(ctx, modelID, contents, nil)
	for resp, err := range stream {
		if err != nil {
			log.Printf("Stream error: %v", err)
			return fullResponse.String(), err
		}
		// Check if there are candidates and parts before accessing
		if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
			textPart := resp.Candidates[0].Content.Parts[0].Text
			fmt.Print(textPart) // Print chunk directly
			fullResponse.WriteString(textPart)
		}
	}
	fmt.Println("\n" + strings.Repeat("_", 80))
	// [END thinking_text_only_prompt_streaming]
	return fullResponse.String(), nil
}

func ThinkingLogicPuzzle() (*genai.GenerateContentResponse, error) {
	// [START thinking_logic_puzzle]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	prompt := `
        Solve this logic puzzle and explain your reasoning step-by-step:
        There are three boxes. One contains only apples, one contains only oranges,
        and one contains both apples and oranges. The boxes have been incorrectly
        labeled such that no label is correct. You are allowed to draw one fruit
        from only one box of your choosing (without looking inside). Which box
        would you draw from to correctly label all boxes, and why?
        `
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, nil)
	if err != nil {
		log.Printf("GenerateContent failed: %v", err)
		return nil, err
	}

	fmt.Println(resp.Text())
	// [END thinking_logic_puzzle]
	return resp, nil
}

func ThinkingCodeExplanation() (*genai.GenerateContentResponse, error) {
	// [START thinking_code_explanation]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	prompt := `
        Explain this Python code snippet step-by-step, including what it does
        and why recursion is used here:

        def factorial(n):
            '''Calculates the factorial of a non-negative integer.'''
            if not isinstance(n, int) or n < 0:
                raise ValueError("Input must be a non-negative integer")
            if n == 0:
                return 1
            else:
                return n * factorial(n-1)

        result = factorial(5)
        print(f"The factorial of 5 is: {result}")
        `
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, nil)
	if err != nil {
		log.Printf("GenerateContent failed: %v", err)
		return nil, err
	}

	fmt.Println(resp.Text())
	// [END thinking_code_explanation]
	return resp, nil
}

func ThinkingCreativeWritingConstraints() (*genai.GenerateContentResponse, error) {
	// [START thinking_creative_writing_constraints]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	prompt := `
        Write a short story (max 150 words) about a detective investigating a
        mystery in a library, but the story must not contain the letter 'e'.
        `
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, nil)
	if err != nil {
		log.Printf("GenerateContent failed: %v", err)
		return nil, err
	}

	fmt.Println(resp.Text())
	// [END thinking_creative_writing_constraints]
	return resp, nil
}

func ThinkingWithSearchTool() (*genai.GenerateContentResponse, error) {
	// [START thinking_with_search_tool]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	googleSearchTool := &genai.Tool{
		GoogleSearch: &genai.GoogleSearch{}, // Empty struct for default search
	}

	prompt := "What were the major scientific breakthroughs announced last week?"
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{googleSearchTool},
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, config)
	if err != nil {
		log.Printf("GenerateContent with search tool failed: %v", err)
		return nil, err
	}

	fmt.Println(resp)

	// [END thinking_with_search_tool]
	return resp, nil
}

func ThinkingWithSearchToolStreaming() (string, error) {
	// [START thinking_with_search_tool_streaming]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return "", err
	}

	googleSearchTool := &genai.Tool{
		GoogleSearch: &genai.GoogleSearch{},
	}

	prompt := "When is the next total solar eclipse visible from mainland Europe?"
	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{googleSearchTool},
	}

	var fullResponseText strings.Builder
	// var finalResponse *genai.GenerateContentResponse // Store the last response chunk

	stream := client.Models.GenerateContentStream(ctx, modelID, contents, config)
	for resp, err := range stream {
		if err != nil {
			log.Printf("Stream error: %v", err)
			fmt.Println("\nCould not access grounding metadata from stream response likely due to error.")
			return fullResponseText.String(), err
		}
		// Process text chunks
		if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil && len(resp.Candidates[0].Content.Parts) > 0 {
			textPart := resp.Candidates[0].Content.Parts[0].Text
			if textPart != "" {
				fmt.Print(textPart)
				fullResponseText.WriteString(textPart)
			}
		}
		// finalResponse = resp // Keep track of the latest response which might contain aggregated data
	}

	fmt.Println("\n" + strings.Repeat("_", 80)) // Separator

	// [END thinking_with_search_tool_streaming]
	return fullResponseText.String(), nil
}

func ThinkingCodeExecution() (*genai.GenerateContentResponse, error) {
	// [START thinking_code_execution]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	prompt := "What is the sum of the first 50 prime numbers? " +
		"Generate and run Python code for the calculation, and make sure you get all 50. " +
		"Provide the final sum clearly."

	codeExecutionTool := &genai.Tool{
		CodeExecution: &genai.ToolCodeExecution{}, // Enable code execution
	}

	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}
	config := &genai.GenerateContentConfig{
		Tools: []*genai.Tool{codeExecutionTool}, // Provide the tool
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, config)
	if err != nil {
		log.Printf("GenerateContent with code execution failed: %v", err)
		return nil, err
	}

	fmt.Println(resp)

	// [END thinking_code_execution]
	return resp, nil
}

func ThinkingStructuredOutputJson() (*genai.GenerateContentResponse, error) {
	// [START thinking_structured_output_json]
	ctx := context.Background()
	client, err := newGenAIClient(ctx)
	if err != nil {
		log.Printf("Failed to create client: %v", err)
		return nil, err
	}

	// Prompt clearly asks for JSON and provides the schema inline
	prompt := `
        Provide a list of 3 famous physicists and their key contributions
        in JSON format.

        Use this JSON schema:

        Physicist = {'name': str, 'contribution': str, 'era': str}
        Return: list[Physicist]
        `

	contents := []*genai.Content{
		genai.NewContentFromText(prompt, genai.RoleUser),
	}

	resp, err := client.Models.GenerateContent(ctx, modelID, contents, nil)

	// For stricter JSON mode (if structured output is supported):
	// config := &genai.GenerateContentConfig{
	// 	ResponseMIMEType: "application/json",
	//  // ResponseSchema: &genai.Schema{...} // Define schema if needed
	// }
	// resp, err := client.Models.GenerateContent(ctx, modelID, contents, config)

	if err != nil {
		log.Printf("GenerateContent failed: %v", err)
		return nil, err
	}

	fmt.Println(resp.Text())
	// [END thinking_structured_output_json]
	return resp, nil
}
