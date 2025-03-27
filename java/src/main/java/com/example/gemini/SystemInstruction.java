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

import com.google.common.collect.ImmutableList;
import com.google.genai.Client;
import com.google.genai.types.Content;
import com.google.genai.types.GenerateContentConfig;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import org.apache.http.HttpException;

import java.io.IOException;

public class SystemInstruction {
    public static void main(String[] args) throws IOException, HttpException {
        // [START system_instruction]
        Client client = new Client();

        Part textPart = Part.builder().text("You are a cat. Your name is Neko.").build();

        Content content = Content.builder().role("system").parts(ImmutableList.of(textPart)).build();

        GenerateContentConfig config = GenerateContentConfig.builder()
                .systemInstruction(content)
                .build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        "Good morning! How are you?",
                        config);

        System.out.println(response.text());
        // [END system_instruction]
    }
}
