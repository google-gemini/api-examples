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

import com.google.genai.types.GenerateContentResponse;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;

public class ChatSessionTest {
    @Test
    public void test_chat() {
        List<GenerateContentResponse> chatResponses = assertDoesNotThrow(ChatSession::chat,
                "chat returned an error");

        for (GenerateContentResponse response : chatResponses) {
            assertNotNull(response, "Response should not be null");
            assertNotNull(response.text(), "Response text should not be null");
            assertFalse(response.text().trim().isEmpty(), "Response text should not be empty");
        }
    }
}
