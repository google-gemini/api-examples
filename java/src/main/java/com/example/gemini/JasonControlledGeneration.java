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
import com.google.genai.types.*;
import org.apache.http.HttpException;

import java.io.IOException;
import java.lang.reflect.Array;
import java.util.List;
import java.util.Map;

public class JasonControlledGeneration {
    public static void main(String[] args) throws HttpException, IOException {
        // [START json_controlled_generation]
        Client client = new Client();

        Schema recipeSchema = Schema.builder()
                .type(Array.class.getSimpleName())
                .items(Schema.builder()
                        .type(Object.class.getSimpleName())
                        .properties(
                                Map.of("recipe_name", Schema.builder()
                                                .type(String.class.getSimpleName())
                                                .build(),
                                        "ingredients", Schema.builder()
                                                .type(Array.class.getSimpleName())
                                                .items(Schema.builder()
                                                        .type(String.class.getSimpleName())
                                                        .build())
                                                .build())
                        )
                        .required(List.of("recipe_name", "ingredients"))
                        .build())
                .build();

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .responseMimeType("application/json")
                        .responseSchema(recipeSchema)
                        .build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        "List a few popular cookie recipes.",
                        config);

        System.out.println(response.text());
        // [END json_controlled_generation]
    }
}
