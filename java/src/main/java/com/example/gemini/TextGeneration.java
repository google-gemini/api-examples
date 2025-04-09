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
import com.google.genai.ResponseStream;
import com.google.genai.types.Blob;
import com.google.genai.types.Content;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import org.jspecify.annotations.Nullable;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Base64;

import static com.example.gemini.BuildConfig.media_path;

@SuppressWarnings("resource")
public class TextGeneration {
    public static @Nullable String textGenTextOnlyPrompt() throws Exception {
        // [START text_gen_text_only_prompt]
        Client client = new Client();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        "Write a story about a magic backpack.",
                        null);

        System.out.println(response.text());
        // [END text_gen_text_only_prompt]
        return response.text();
    }

    public static String textGenTextOnlyPromptStreaming() throws Exception {
        // [START text_gen_text_only_prompt_streaming]
        Client client = new Client();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream(
                        "gemini-2.0-flash",
                        "Write a story about a magic backpack.",
                        null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_text_only_prompt_streaming]
        return response.toString();
    }

    public static @Nullable String textGenMultimodalOneImagePrompt() throws Exception {
        // [START text_gen_multimodal_one_image_prompt]
        Client client = new Client();

        String path = media_path + "organ.jpg";
        byte[] imageData = Files.readAllBytes(Paths.get(path));
        String base64Image = Base64.getEncoder().encodeToString(imageData);
        Part imagePart = Part.builder()
                .inlineData(Blob.builder().data(base64Image)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("Tell me about this instrument").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, imagePart)).build();

        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash", content, null);

        System.out.println(response.text());
        // [END text_gen_multimodal_one_image_prompt]
        return response.text();
    }

    public static String textGenMultimodalOneImagePromptStreaming() throws Exception {
        // [START text_gen_multimodal_one_image_prompt_streaming]
        Client client = new Client();

        String path = media_path + "organ.jpg";
        byte[] imageData = Files.readAllBytes(Paths.get(path));
        String base64Image = Base64.getEncoder().encodeToString(imageData);
        Part imagePart = Part.builder()
                .inlineData(Blob.builder().data(base64Image)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("Tell me about this instrument").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, imagePart)).build();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream(
                        "gemini-2.0-flash",
                        content,
                        null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_one_image_prompt_streaming]
        return response.toString();
    }

    public static @Nullable String textGenMultimodalMultiImagePrompt() throws Exception {
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
        return response.text();
    }

    public static String textGenMultimodalMultiImagePromptStreaming() throws Exception {
        // [START text_gen_multimodal_multi_image_prompt_streaming]
        Client client = new Client();

        String organPath = media_path + "organ.jpg";
        byte[] organImageData = Files.readAllBytes(Paths.get(organPath));
        String organBase64Image = Base64.getEncoder().encodeToString(organImageData);
        Part organImagePart = Part.builder()
                .inlineData(Blob.builder().data(organBase64Image)
                        .mimeType("image/jpeg").build()).build();

        String cajunPath = media_path + "Cajun_instruments.jpg";
        byte[] cajunImageData = Files.readAllBytes(Paths.get(cajunPath));
        String cajunImageBase64 = Base64.getEncoder().encodeToString(cajunImageData);
        Part cajunImagePart = Part.builder()
                .inlineData(Blob.builder().data(cajunImageBase64)
                        .mimeType("image/jpeg").build()).build();

        Part textPart = Part.builder().text("What is the difference between both of these instruments?").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, organImagePart, cajunImagePart)).build();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream("gemini-2.0-flash", content, null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_multi_image_prompt_streaming]
        return response.toString();
    }

    public static @Nullable String textGenMultimodalAudio() throws Exception {
        // [START text_gen_multimodal_audio]
        Client client = new Client();

        String path = media_path + "sample.mp3";
        byte[] audioData = Files.readAllBytes(Paths.get(path));
        String audioBase64 = Base64.getEncoder().encodeToString(audioData);
        Part audioPart = Part.builder()
                .inlineData(Blob.builder().data(audioBase64)
                        .mimeType("audio/mpeg").build()).build();

        Part textPart = Part.builder().text("Give me a summary of this audio file.").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, audioPart)).build();

        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash", content, null);

        System.out.println(response.text());
        // [END text_gen_multimodal_audio]
        return response.text();
    }

    public static String textGenMultimodalAudioStreaming() throws Exception {
        // [START text_gen_multimodal_audio_streaming]
        Client client = new Client();

        String path = media_path + "sample.mp3";
        byte[] audioData = Files.readAllBytes(Paths.get(path));
        String audioBase64 = Base64.getEncoder().encodeToString(audioData);
        Part audioPart = Part.builder()
                .inlineData(Blob.builder().data(audioBase64)
                        .mimeType("audio/mpeg").build()).build();

        Part textPart = Part.builder().text("Give me a summary of this audio file.").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, audioPart)).build();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream("gemini-2.0-flash", content, null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_audio_streaming]
        return response.toString();
    }

    public static @Nullable String textGenMultimodalVideoPrompt() throws Exception {
        // [START text_gen_multimodal_video_prompt]
        Client client = new Client();

        String path = media_path + "Big_Buck_Bunny.mp4";
        byte[] videoData = Files.readAllBytes(Paths.get(path));
        String videoBase64 = Base64.getEncoder().encodeToString(videoData);
        Part videoPart = Part.builder()
                .inlineData(Blob.builder().data(videoBase64)
                        .mimeType("video/mp4").build()).build();

        Part textPart = Part.builder().text("Describe this video clip").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, videoPart)).build();

        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash", content, null);

        System.out.println(response.text());
        // [END text_gen_multimodal_video_prompt]
        return response.text();
    }

    public static String textGenMultimodalVideoPromptStreaming() throws Exception {
        // [START text_gen_multimodal_video_prompt_streaming]
        Client client = new Client();

        String path = media_path + "Big_Buck_Bunny.mp4";
        byte[] videoData = Files.readAllBytes(Paths.get(path));
        String videoBase64 = Base64.getEncoder().encodeToString(videoData);
        Part videoPart = Part.builder()
                .inlineData(Blob.builder().data(videoBase64)
                        .mimeType("video/mp4").build()).build();

        Part textPart = Part.builder().text("Describe this video clip").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, videoPart)).build();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream("gemini-2.0-flash", content, null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_video_prompt_streaming]
        return response.toString();
    }

    public static @Nullable String textGenMultimodalPdf() throws Exception {
        // [START text_gen_multimodal_pdf]
        Client client = new Client();

        String path = media_path + "test.pdf";
        byte[] pdfData = Files.readAllBytes(Paths.get(path));
        String pdfBase64 = Base64.getEncoder().encodeToString(pdfData);
        Part pdfPart = Part.builder()
                .inlineData(Blob.builder().data(pdfBase64)
                        .mimeType("application/pdf").build()).build();

        Part textPart = Part.builder().text("Give me a summary of this document.").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, pdfPart)).build();

        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash", content, null);

        System.out.println(response.text());
        // [END text_gen_multimodal_pdf]
        return response.text();
    }

    public static String textGenMultimodalPdfStreaming() throws Exception {
        // [START text_gen_multimodal_pdf_streaming]
        Client client = new Client();

        String path = media_path + "test.pdf";
        byte[] pdfData = Files.readAllBytes(Paths.get(path));
        String pdfBase64 = Base64.getEncoder().encodeToString(pdfData);
        Part pdfPart = Part.builder()
                .inlineData(Blob.builder().data(pdfBase64)
                        .mimeType("application/pdf").build()).build();

        Part textPart = Part.builder().text("Give me a summary of this document.").build();

        Content content = Content.builder().role("user").parts(ImmutableList.of(textPart, pdfPart)).build();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream("gemini-2.0-flash", content, null);

        StringBuilder response = new StringBuilder();
        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
            response.append(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_pdf_streaming]
        return response.toString();
    }
}
