/*
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.example.gemini;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertFalse;


public class TextGenerationTest {
    @Test
    public void test_textGenTextOnlyPrompt() {
        String result = assertDoesNotThrow(TextGeneration::textGenTextOnlyPrompt,
                "textGenTextOnlyPrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenTextOnlyPromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenTextOnlyPromptStreaming,
                "textGenTextOnlyPromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalOneImagePrompt() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalOneImagePrompt,
                "textGenMultimodalOneImagePrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalOneImagePromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalOneImagePromptStreaming,
                "textGenMultimodalOneImagePromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalMultiImagePrompt() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalMultiImagePrompt,
                "textGenMultimodalMultiImagePrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalMultiImagePromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalMultiImagePromptStreaming,
                "textGenMultimodalMultiImagePromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalAudio() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalAudio,
                "textGenMultimodalAudio returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalAudioStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalAudioStreaming,
                "textGenMultimodalAudioStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalVideoPrompt() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalVideoPrompt,
                "textGenMultimodalVideoPrompt returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalVideoPromptStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalVideoPromptStreaming,
                "textGenMultimodalVideoPromptStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalPdf() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalPdf,
                "textGenMultimodalPdf returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }

    @Test
    public void test_textGenMultimodalPdfStreaming() {
        String result = assertDoesNotThrow(TextGeneration::textGenMultimodalPdfStreaming,
                "textGenMultimodalPdfStreaming returned an error");

        assertNotNull(result, "Response should not be null");
        assertFalse(result.trim().isEmpty(), "Response should not be empty");
    }
}
