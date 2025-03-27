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
import com.google.genai.types.GenerateContentResponse;
import org.apache.http.HttpException;

import java.io.IOException;

public class TextGenTextOnlyPromptStreaming {
    public static void main(String[] args) throws IOException, HttpException {
        // [START text_gen_text_only_prompt_streaming]
        Client client = new Client();

        ResponseStream<GenerateContentResponse> responseStream =
                client.models.generateContentStream(
                        "gemini-2.0-flash",
                        "Write a story about a magic backpack.",
                        null);

        for (GenerateContentResponse res : responseStream) {
            System.out.print(res.text());
        }

        responseStream.close();
        // [END text_gen_text_only_prompt_streaming]
    }

}
