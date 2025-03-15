/**
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package examples

import (
	"testing"
)

func TestTokensTextOnly(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensTextOnly()
	if err != nil {
		t.Errorf("TokensTextOnly returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensChat(t *testing.T) {
	historyTokenCount, chatResponse, combinedTokenCount, err := TokensChat()
	if err != nil {
		t.Errorf("TokensChat returned an error: %v", err)
	}
	if historyTokenCount.TotalTokens <= 0 {
		t.Errorf("Expected history token count to be greater than 0, got %d", historyTokenCount.TotalTokens)
	}
	if combinedTokenCount.TotalTokens <= historyTokenCount.TotalTokens {
		t.Errorf("Expected combined token count to be greater than history token count, got %d vs %d", 
			combinedTokenCount.TotalTokens, historyTokenCount.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if chatResponse != nil && chatResponse.UsageMetadata != nil {
		if chatResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", chatResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensMultimodalImageInline(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensMultimodalImageInline()
	if err != nil {
		t.Errorf("TokensMultimodalImageInline returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensMultimodalImageFileAPI(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensMultimodalImageFileAPI()
	if err != nil {
		t.Errorf("TokensMultimodalImageFileAPI returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensMultimodalVideoAudioFileAPI(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensMultimodalVideoAudioFileAPI()
	if err != nil {
		t.Errorf("TokensMultimodalVideoAudioFileAPI returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensMultimodalPdfFileAPI(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensMultimodalPdfFileAPI()
	if err != nil {
		t.Errorf("TokensMultimodalPdfFileAPI returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}

func TestTokensSystemInstruction(t *testing.T) {
	countTokensResponse, generateResponse, err := TokensSystemInstruction()
	if err != nil {
		t.Errorf("TokensSystemInstruction returned an error: %v", err)
	}
	if countTokensResponse.TotalTokens <= 0 {
		t.Errorf("Expected total tokens to be greater than 0, got %d", countTokensResponse.TotalTokens)
	}
	// Check that usage metadata has numeric values
	if generateResponse != nil && generateResponse.UsageMetadata != nil {
		if generateResponse.UsageMetadata.PromptTokenCount < 0 {
			t.Errorf("Expected prompt token count to be greater than or equal to 0, got %d", generateResponse.UsageMetadata.PromptTokenCount)
		}
	} else {
		t.Error("Expected usage metadata to be present")
	}
}
