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

import com.google.genai.Chat;
import com.google.genai.Client;
import com.google.genai.types.Content;
import com.google.genai.types.GenerateContentConfig;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;

import java.util.Collections;
import java.util.List;

public class ChatSession {
    public static List<GenerateContentResponse> chat() {
        // [START chat]
        Client client = new Client();

        Content userContent = Content.fromParts(Part.fromText("Hello"));
        Content modelContent =
                Content.builder()
                        .role("model")
                        .parts(
                                Collections.singletonList(
                                        Part.fromText("Great to meet you. What would you like to know?")
                                )
                        ).build();

        Chat chat = client.chats.create(
                "gemini-2.0-flash",
                GenerateContentConfig.builder()
                        .systemInstruction(userContent)
                        .systemInstruction(modelContent)
                        .build()
        );

        GenerateContentResponse response1 = chat.sendMessage("I have 2 dogs in my house.");
        System.out.println(response1.text());

        GenerateContentResponse response2 = chat.sendMessage("How many paws are in my house?");
        System.out.println(response2.text());

        // [END chat]
        return List.of(response1, response2);
    }
}
