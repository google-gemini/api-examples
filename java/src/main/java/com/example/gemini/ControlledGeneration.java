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
import com.google.genai.types.Blob;
import com.google.genai.types.Content;
import com.google.genai.types.GenerateContentConfig;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import com.google.genai.types.Schema;
import org.jspecify.annotations.Nullable;

import java.lang.reflect.Array;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.Base64;
import java.util.List;
import java.util.Map;

import static com.example.gemini.BuildConfig.media_path;

public class ControlledGeneration {
    public static @Nullable String jsonControlledGeneration() throws Exception {
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
        return response.text();
    }

    public static @Nullable String jsonNoSchema() throws Exception {
        // [START json_no_schema]
        Client client = new Client();

        String prompt = """
                List a few popular cookie recipes in JSON format.
                
                
                Use this JSON schema:
                
                
                Recipe = {'recipe_name': string, 'ingredients': list[string]}
                
                Return: list[Recipe].
                """;

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        prompt,
                        null);

        System.out.println(response.text());
        // [END json_no_schema]
        return response.text();
    }

    public static @Nullable String jsonEnum() throws Exception {
        // [START json_enum]
        Client client = new Client();

        Schema schema =
                Schema.builder()
                        .type(String.class.getSimpleName())
                        .enum_(Arrays.asList("Percussion", "String", "Woodwind", "Brass", "Keyboard"))
                        .build();

        String path = media_path + "organ.jpg";
        byte[] imageData = Files.readAllBytes(Paths.get(path));
        String base64Image = Base64.getEncoder().encodeToString(imageData);
        Part imagePart = Part.builder()
                .inlineData(Blob.builder().data(base64Image)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("What kind of instrument is this:").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, imagePart)).build();

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .responseMimeType("application/json")
                        .candidateCount(1)
                        .responseSchema(schema)
                        .build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        content,
                        config);

        System.out.println(response.text());
        // [END json_enum]
        return response.text();
    }

    public static @Nullable String enumInJson() throws Exception {
        // [START enum_in_json]
        Client client = new Client();

        Schema recipeSchema = Schema.builder()
                .type(Array.class.getSimpleName())
                .items(Schema.builder()
                        .type(Object.class.getSimpleName())
                        .properties(
                                Map.of("recipe_name", Schema.builder()
                                                .type(String.class.getSimpleName())
                                                .build(),
                                        "grade", Schema.builder()
                                                .type(String.class.getSimpleName())
                                                .enum_(Arrays.asList("a+", "a", "b", "c", "d", "f"))
                                                .build())
                        )
                        .required(List.of("recipe_name", "grade"))
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
                        "List about 10 cookie recipes, grade them based on popularity",
                        config);

        System.out.println(response.text());
        // [END enum_in_json]
        return response.text();
    }

    public static @Nullable String xEnum() throws Exception {
        // [START x_enum]
        Client client = new Client();

        Schema schema =
                Schema.builder()
                        .type(String.class.getSimpleName())
                        .enum_(Arrays.asList("Percussion", "String", "Woodwind", "Brass", "Keyboard"))
                        .build();

        String path = media_path + "organ.jpg";
        byte[] imageData = Files.readAllBytes(Paths.get(path));
        String base64Image = Base64.getEncoder().encodeToString(imageData);
        Part imagePart = Part.builder()
                .inlineData(Blob.builder().data(base64Image)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("What kind of instrument is this:").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, imagePart)).build();

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .responseMimeType("text/x.enum")
                        .candidateCount(1)
                        .responseSchema(schema)
                        .build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        content,
                        config);

        System.out.println(response.text());
        // [END x_enum]
        return response.text();
    }
}
