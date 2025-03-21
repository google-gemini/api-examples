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
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import org.apache.http.HttpException;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Base64;

import static com.example.gemini.BuildConfig.media_path;

public class TextGenMultimodalMultiImagePrompt {
    public static void main(String[] args) throws IOException, HttpException {
        // [START text_gen_multimodal_multi_image_prompt]
        Client client = new Client();

        String organPath = media_path + "organ.jpg";
        byte[] organImageData = Files.readAllBytes(Paths.get(organPath));
        String organImageBase64 = Base64.getEncoder().encodeToString(organImageData);
        Part organImagePart = Part.builder()
                .inlineData(Blob.builder().data(organImageBase64)
                        .mimeType("image/jpeg").build()).build();

        String cajunPath = media_path + "Cajun_instruments.jpg";
        byte[] cajunImageData = Files.readAllBytes(Paths.get(cajunPath));
        String cajunImageBase64 = Base64.getEncoder().encodeToString(cajunImageData);
        Part cajunImagePart = Part.builder()
                .inlineData(Blob.builder().data(cajunImageBase64)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("What is the difference between both of these instruments?").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, organImagePart, cajunImagePart)).build();

        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash", content, null);

        System.out.println(response.text());
        // [END text_gen_multimodal_multi_image_prompt]
    }
}
