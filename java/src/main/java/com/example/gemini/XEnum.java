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
import com.google.genai.types.*;
import org.apache.http.HttpException;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.Base64;

import static com.example.gemini.BuildConfig.media_path;

public class XEnum {
    public static void main(String[] args) throws HttpException, IOException {
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
    }
}
