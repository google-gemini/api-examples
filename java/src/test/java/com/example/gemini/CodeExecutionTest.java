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
import com.google.genai.types.Part;
import org.junit.jupiter.api.Test;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertNotNull;


public class CodeExecutionTest {
    @Test
    public void test_codeExecutionBasic() {
        GenerateContentResponse response = assertDoesNotThrow(CodeExecution::codeExecutionBasic,
                "codeExecutionBasic returned an error");

        assertNotNull(response, "Response should not be null");

        List<Part> parts = response.candidates().get().getFirst().content().get().parts().get();
        assertNotNull(parts, "Response parts should not be null");
        assert !parts.isEmpty();

        assertNotNull(response.text(), "Response text should not be null");
        assertFalse(response.text().trim().isEmpty(), "Response text should not be empty");
    }

    @Test
    public void test_codeExecutionRequestOverride() {
        GenerateContentResponse response = assertDoesNotThrow(CodeExecution::codeExecutionRequestOverride,
                "codeExecutionRequestOverride returned an error");

        assertNotNull(response, "Response should not be null");

        assertNotNull(response.executableCode(), "ExecutableCode response should not be null");
        assertNotNull(response.codeExecutionResult(), "CodeExecutionResult response should not be null");
    }
}
