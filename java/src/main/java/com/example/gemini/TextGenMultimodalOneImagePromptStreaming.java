package com.example.gemini;

import com.google.common.collect.ImmutableList;
import com.google.genai.Client;
import com.google.genai.ResponseStream;
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

public class TextGenMultimodalOneImagePromptStreaming {
    public static void main(String[] args) throws IOException, HttpException {
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

        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
        }

        responseStream.close();
        // [END text_gen_multimodal_one_image_prompt_streaming]
    }
}
