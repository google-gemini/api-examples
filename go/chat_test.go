package examples

import (
	"testing"
)

func TestChat(t *testing.T) {
	err := Chat()
	if err != nil {
		t.Errorf("Chat returned an error: %v", err)
	}
}

func TestChatStreaming(t *testing.T) {
	err := ChatStreaming()
	if err != nil {
		t.Errorf("ChatStreaming returned an error: %v", err)
	}
}

func TestChatStreamingWithImages(t *testing.T) {
	err := ChatStreamingWithImages()
	if err != nil {
		t.Errorf("ChatStreamingWithImages returned an error: %v", err)
	}
}
