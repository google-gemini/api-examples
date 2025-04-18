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
import com.google.genai.types.GenerateContentResponse;
import org.apache.http.HttpException;

import java.io.IOException;

public class JsonNoSchema {
    public static void main(String[] args) throws HttpException, IOException {
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
    }
}
