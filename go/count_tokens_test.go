package examples

import (
	"testing"
)

func TestTokensContextWindow(t *testing.T) {
	err := TokensContextWindow()
	if err != nil {
		t.Errorf("TokensContextWindow returned an error.")
	}
}

func TestTokensTextOnly(t *testing.T) {
	err := TokensTextOnly()
	if err != nil {
		t.Errorf("TokensTextOnly returned an error.")
	}
}

func TestTokensChat(t *testing.T) {
	if err := TokensChat(); err != nil {
		t.Errorf("TokensChat returned an error: %v", err)
	}
}

func TestTokensMultimodalImageFileApi(t *testing.T) {
	err := TokensMultimodalImageFileApi()
	if err != nil {
		t.Errorf("TokensMultimodalImageFileApi returned an error.")
	}
}

func TestTokensMultimodalVideoAudioFileApi(t *testing.T) {
	err := TokensMultimodalVideoAudioFileApi()
	if err != nil {
		t.Errorf("TokensMultimodalVideoAudioFileApi returned an error.")
	}
}

func TestTokensMultimodalPdfFileApi(t *testing.T) {
	err := TokensMultimodalPdfFileApi()
	if err != nil {
		t.Errorf("TokensMultimodalPdfFileApi returned an error.")
	}
}

func TestTokensCachedContent(t *testing.T) {
	err := TokensCachedContent()
	if err != nil {
		t.Errorf("TokensCachedContent returned an error.")
	}
}
