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

import com.google.genai.Client;
import com.google.genai.types.GenerateContentConfig;
import com.google.genai.types.GenerateContentResponse;
import org.jspecify.annotations.Nullable;

import java.util.List;

public class ConfigureModelParameters {
    public static @Nullable String configureModelParameters() throws Exception {
        // [START configure_model_parameters]
        Client client = new Client();

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .candidateCount(1)
                        .stopSequences(List.of("x"))
                        .maxOutputTokens(20)
                        .temperature(1.0F)
                        .build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        "Tell me a story about a magic backpack.",
                        config);

        System.out.println(response.text());
        // [END configure_model_parameters]
        return response.text();
    }
}
