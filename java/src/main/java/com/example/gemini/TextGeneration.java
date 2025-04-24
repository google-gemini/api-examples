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
import com.google.genai.ResponseStream;
import com.google.genai.types.Content;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import org.jspecify.annotations.Nullable;

import java.nio.file.Files;
import java.nio.file.Paths;

import static com.example.gemini.BuildConfig.media_path;

public class TextGeneration {
    public static @Nullable String textGenTextOnlyPrompt() {
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

    public static String textGenTextOnlyPromptStreaming() {
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

        Content content =
                Content.fromParts(
                        Part.fromText("Tell me about this instrument."),
                        Part.fromBytes(imageData, "image/jpeg"));

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

        Content content =
                Content.fromParts(
                        Part.fromText("Tell me about this instrument."),
                        Part.fromBytes(imageData, "image/jpeg"));


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

        String cajunPath = media_path + "Cajun_instruments.jpg";
        byte[] cajunImageData = Files.readAllBytes(Paths.get(cajunPath));

        Content content =
                Content.fromParts(
                        Part.fromText("What is the difference between both of these instruments?"),
                        Part.fromBytes(organImageData, "image/jpeg"),
                        Part.fromBytes(cajunImageData, "image/jpeg"));


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

        String cajunPath = media_path + "Cajun_instruments.jpg";
        byte[] cajunImageData = Files.readAllBytes(Paths.get(cajunPath));

        Content content =
                Content.fromParts(
                        Part.fromText("What is the difference between both of these instruments?"),
                        Part.fromBytes(organImageData, "image/jpeg"),
                        Part.fromBytes(cajunImageData, "image/jpeg"));

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

        Content content =
                Content.fromParts(Part.fromText("Give me a summary of this audio file."),
                        Part.fromBytes(audioData, "audio/mpeg"));

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

        Content content =
                Content.fromParts(Part.fromText("Give me a summary of this audio file."),
                        Part.fromBytes(audioData, "audio/mpeg"));

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

        Content content =
                Content.fromParts(Part.fromText("Describe this video clip."),
                        Part.fromBytes(videoData, "video/mp4"));

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

        Content content =
                Content.fromParts(Part.fromText("Describe this video clip."),
                        Part.fromBytes(videoData, "video/mp4"));

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

        Content content =
                Content.fromParts(Part.fromText("Give me a summary of this document."),
                        Part.fromBytes(pdfData, "application/pdf"));

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

        Content content =
                Content.fromParts(Part.fromText("Give me a summary of this document."),
                        Part.fromBytes(pdfData, "application/pdf"));

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
