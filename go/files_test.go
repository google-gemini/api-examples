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
	"os"
	"testing"
)

// TestFilesCreateText tests the text file upload and generation functionality
func TestFilesCreateText(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesCreateText()
}

// TestFilesCreateImage tests the image file upload and generation functionality
func TestFilesCreateImage(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesCreateImage()
}

// TestFilesCreateAudio tests the audio file upload and generation functionality
func TestFilesCreateAudio(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesCreateAudio()
}

// TestFilesCreateVideo tests the video file upload and generation functionality
func TestFilesCreateVideo(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesCreateVideo()
}

// TestFilesCreatePdf tests the PDF file upload and generation functionality
func TestFilesCreatePdf(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesCreatePdf()
}

// TestFilesList tests the file listing functionality
func TestFilesList(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesList()
}

// TestFilesGet tests the file retrieval functionality
func TestFilesGet(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesGet()
}

// TestFilesDelete tests the file deletion functionality
func TestFilesDelete(t *testing.T) {
	// Skip test if API key is not set
	if os.Getenv("GEMINI_API_KEY") == "" {
		t.Skip("GEMINI_API_KEY not set, skipping test")
	}
	
	// Run the function and check for errors
	FilesDelete()
}
