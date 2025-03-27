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
import com.google.genai.types.SafetySetting;
import org.apache.http.HttpException;

import java.io.IOException;
import java.util.Collections;

public class SafetySettings {
    public static void main(String[] args) throws IOException, HttpException {
        // [START safety_settings]
        Client client = new Client();

        String unsafePrompt = """
                 I support Martians Soccer Club and I think Jupiterians Football Club sucks!
                 Write a ironic phrase about them including expletives.
                """;

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .safetySettings(Collections.singletonList(
                                SafetySetting.builder()
                                        .category("HARM_CATEGORY_HARASSMENT")
                                        .threshold("BLOCK_ONLY_HIGH")
                                        .build()
                        )).build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        unsafePrompt,
                        config);

        System.out.println(response.candidates().get().getFirst().finishReason());
        System.out.println(response.candidates().get().getFirst().safetyRatings());
        // [END safety_settings]
    }
}
