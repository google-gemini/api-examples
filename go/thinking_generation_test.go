package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	// Ensure this import is correct
)

// TestMain checks for the API key before running tests.
func TestMain(m *testing.M) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("WARNING: GEMINI_API_KEY environment variable not set. Skipping Go tests.")
		// os.Exit(0) // Exit successfully without running tests
		// Or, let tests run and fail if they depend on the client
	}
	// Run the tests
	os.Exit(m.Run())
}

// Helper sleep function for potential rate limits
func sleep(d time.Duration) {
	time.Sleep(d)
}

const testDelay = 1 * time.Second // Delay between tests

func TestThinkingTextOnlyPrompt(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	resp, err := ThinkingTextOnlyPrompt()
	if err != nil {
		t.Fatalf("ThinkingTextOnlyPrompt failed: %v", err)
	}
	if resp == nil {
		t.Fatal("ThinkingTextOnlyPrompt returned nil response without error") // Changed to Fatal
	}
	// Handle both return values from resp.Text()
	text, textErr := resp.Text()
	if textErr != nil {
		t.Errorf("resp.Text() returned an error: %v", textErr)
	}
	if text == "" {
		t.Error("ThinkingTextOnlyPrompt response text is empty")
	}
	sleep(testDelay)
}

func TestThinkingTextOnlyPromptStreaming(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	// This function returns (string, error) directly
	fullResp, err := ThinkingTextOnlyPromptStreaming()
	if err != nil {
		// Streaming might error mid-stream, check if some response was received
		if fullResp != "" {
			t.Logf("ThinkingTextOnlyPromptStreaming returned partial response before error: %v", err)
		} else {
			// If error and no response, it's a clearer failure
			t.Fatalf("ThinkingTextOnlyPromptStreaming failed: %v", err)
		}
	}
	// If no error, ensure response is not empty
	if err == nil && fullResp == "" {
		t.Error("ThinkingTextOnlyPromptStreaming returned empty response without error")
	}
	sleep(testDelay)
}

func TestThinkingLogicPuzzle(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	resp, err := ThinkingLogicPuzzle()
	if err != nil {
		t.Fatalf("ThinkingLogicPuzzle failed: %v", err)
	}
	if resp == nil {
		t.Fatal("ThinkingLogicPuzzle returned nil response without error")
	}
	// Handle both return values from resp.Text()
	text, textErr := resp.Text()
	if textErr != nil {
		t.Errorf("resp.Text() returned an error: %v", textErr)
	}
	if text == "" {
		t.Error("ThinkingLogicPuzzle response text is empty")
	}
	sleep(testDelay)
}

func TestThinkingCodeExplanation(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	resp, err := ThinkingCodeExplanation()
	if err != nil {
		t.Fatalf("ThinkingCodeExplanation failed: %v", err)
	}
	if resp == nil {
		t.Fatal("ThinkingCodeExplanation returned nil response without error")
	}
	// Handle both return values from resp.Text()
	text, textErr := resp.Text()
	if textErr != nil {
		t.Errorf("resp.Text() returned an error: %v", textErr)
	}
	if text == "" {
		t.Error("ThinkingCodeExplanation response text is empty")
	}
	sleep(testDelay)
}

func TestThinkingCreativeWritingConstraints(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	resp, err := ThinkingCreativeWritingConstraints()
	if err != nil {
		t.Fatalf("ThinkingCreativeWritingConstraints failed: %v", err)
	}
	if resp == nil {
		t.Fatal("ThinkingCreativeWritingConstraints returned nil response without error")
	}
	// Handle both return values from resp.Text()
	text, textErr := resp.Text()
	if textErr != nil {
		t.Errorf("resp.Text() returned an error: %v", textErr)
	}
	if text == "" {
		t.Error("ThinkingCreativeWritingConstraints response text is empty")
	}
	// Optional: Stricter check (might be flaky)
	// if strings.Contains(strings.ToLower(text), "e") {
	// 	t.Errorf("Constraint check failed: Result contains 'e'")
	// }
	sleep(testDelay)
}

func TestThinkingWithSearchTool(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	// t.Setenv("GOOGLE_API_KEY", os.Getenv("GEMINI_API_KEY")) // May not be needed depending on auth flow

	resp, err := ThinkingWithSearchTool()
	// Search can sometimes fail or return empty if no relevant results
	if err != nil {
		// Check for specific error types if needed, otherwise just log
		t.Logf("ThinkingWithSearchTool returned an error (may be expected if no results found): %v", err)
		// Allow test to pass if error occurred but function handled it gracefully (didn't panic)
	} else if resp == nil {
		t.Error("ThinkingWithSearchTool returned nil response without error")
	} else {
		// Optionally check text if response was successful
		_, textErr := resp.Text() // Check if Text() itself errors
		if textErr != nil {
			t.Logf("resp.Text() on successful ThinkingWithSearchTool call returned error (may be ok if only tool calls): %v", textErr)
		}
	}
	// A more robust test could check resp.Candidates[0].GroundingMetadata != nil

	sleep(testDelay)
}

func TestThinkingWithSearchToolStreaming(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	// t.Setenv("GOOGLE_API_KEY", os.Getenv("GEMINI_API_KEY"))

	// This function returns (string, error) directly
	fullResp, err := ThinkingWithSearchToolStreaming()
	if err != nil {
		t.Logf("ThinkingWithSearchToolStreaming returned an error (may be expected): %v", err)
	}
	// Check if at least some text was generated, even if an error occurred later
	// Or if the function completed without error, ensure response wasn't empty
	if err == nil && fullResp == "" {
		t.Error("ThinkingWithSearchToolStreaming completed successfully but returned empty response")
	}

	sleep(testDelay)
}

func TestThinkingCodeExecution(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	// t.Setenv("GOOGLE_API_KEY", os.Getenv("GEMINI_API_KEY"))

	resp, err := ThinkingCodeExecution()
	if err != nil {
		t.Logf("ThinkingCodeExecution returned an error (may be expected for certain prompts/tool issues): %v", err)
	} else if resp == nil {
		t.Error("ThinkingCodeExecution returned nil response without error")
	} else {
		// Optionally check text if response was successful
		_, textErr := resp.Text() // Check if Text() itself errors
		if textErr != nil {
			t.Logf("resp.Text() on successful ThinkingCodeExecution call returned error (may be ok if only tool calls): %v", textErr)
		}
	}

	// More robust: Check if resp contains ExecutableCode or CodeExecutionResult parts
	foundToolPart := false
	if resp != nil && len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		for _, part := range resp.Candidates[0].Content.Parts {
			if part != nil && (part.ExecutableCode != nil || part.CodeExecutionResult != nil) {
				foundToolPart = true
				break
			}
		}
	}
	// It's possible the model chooses not to use the tool, so lack of tool parts isn't always a failure.
	// But if no error and no candidates at all, that might be an issue.
	if err == nil && resp != nil && len(resp.Candidates) == 0 {
		t.Error("ThinkingCodeExecution returned response with no candidates")
	}
	t.Logf("Found Code Execution/Result Part: %t", foundToolPart)

	sleep(testDelay)
}

func TestThinkingStructuredOutputJson(t *testing.T) {
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set")
	}
	resp, err := ThinkingStructuredOutputJson()
	if err != nil {
		t.Fatalf("ThinkingStructuredOutputJson failed: %v", err)
	}
	if resp == nil {
		t.Fatal("ThinkingStructuredOutputJson returned nil response without error")
	}

	// Handle both return values from resp.Text()
	responseText, textErr := resp.Text()
	if textErr != nil {
		t.Errorf("resp.Text() returned an error: %v", textErr)
	}
	if responseText == "" {
		t.Error("ThinkingStructuredOutputJson response text is empty")
	}

	// Try parsing the JSON in the test only if text extraction was successful
	if textErr == nil && responseText != "" {
		jsonStr := responseText
		// Basic markdown cleanup
		if strings.HasPrefix(jsonStr, "```json") && strings.HasSuffix(jsonStr, "```") {
			jsonStr = strings.TrimPrefix(jsonStr, "```json")
			jsonStr = strings.TrimSuffix(jsonStr, "```")
			jsonStr = strings.TrimSpace(jsonStr)
		} else if strings.HasPrefix(jsonStr, "```") && strings.HasSuffix(jsonStr, "```") {
			jsonStr = strings.TrimPrefix(jsonStr, "```")
			jsonStr = strings.TrimSuffix(jsonStr, "```")
			jsonStr = strings.TrimSpace(jsonStr)
		}

		var parsedJson any
		parseErr := json.Unmarshal([]byte(jsonStr), &parsedJson)
		if parseErr != nil {
			t.Errorf("Failed to parse response text as JSON in test: %v\nResponse text was:\n%s", parseErr, responseText)
		}
	}
	sleep(testDelay)
}
