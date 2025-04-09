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

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class SafetySettingsTest {
    @Test
    public void test_safetySettings() {
        GenerateContentResponse response = assertDoesNotThrow(SafetySettings::safetySettings,
                "safetySettings returned an error");

        assertNotNull(response, "Response should not be null");

        assertTrue(
                response.candidates().get().getFirst().finishReason().isPresent(),
                "Safety finish reason should be present"
        );
        assertTrue(
                response.candidates().get().getFirst().safetyRatings().isPresent(),
                "Safety ratings should be present"
        );
    }

    @Test
    public void test_safetySettingsMulti() {
        GenerateContentResponse response = assertDoesNotThrow(SafetySettings::safetySettingsMulti,
                "safetySettingsMulti returned an error");

        assertNotNull(response, "Response should not be null");

        assertTrue(
                response.candidates().get().getFirst().safetyRatings().isPresent(),
                "Safety ratings should be present"
        );
    }
}
