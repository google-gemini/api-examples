package examples

import (
	"testing"

	"google.golang.org/genai"
)

func TestPrintResponseNil(t *testing.T) {
	// Must not panic on nil response
	printResponse(nil)
}

func TestPrintResponseEmpty(t *testing.T) {
	resp := &genai.GenerateContentResponse{}
	printResponse(resp)
}

func TestPrintResponseWithCandidates(t *testing.T) {
	resp := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []*genai.Part{
						{Text: "test output"},
					},
				},
			},
		},
	}
	printResponse(resp)
}
