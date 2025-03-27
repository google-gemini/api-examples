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
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.function.BiFunction;

public class FunctionCalling {
    public static void main(String[] args) throws IOException {
        // [START function_calling]
        Client client = new Client();

        FunctionDeclaration addFunction =
                FunctionDeclaration.builder()
                        .name("addNumbers")
                        .parameters(
                                Schema.builder()
                                        .type("object")
                                        .properties(Map.of(
                                                "firstParam", Schema.builder().type("number").description("First number").build(),
                                                "secondParam", Schema.builder().type("number").description("Second number").build()))
                                        .required(Arrays.asList("firstParam", "secondParam"))
                                        .build())
                        .build();

        FunctionDeclaration subtractFunction =
                FunctionDeclaration.builder()
                        .name("subtractNumbers")
                        .parameters(
                                Schema.builder()
                                        .type("object")
                                        .properties(Map.of(
                                                "firstParam", Schema.builder().type("number").description("First number").build(),
                                                "secondParam", Schema.builder().type("number").description("Second number").build()))
                                        .required(Arrays.asList("firstParam", "secondParam"))
                                        .build())
                        .build();

        FunctionDeclaration multiplyFunction =
                FunctionDeclaration.builder()
                        .name("multiplyNumbers")
                        .parameters(
                                Schema.builder()
                                        .type("object")
                                        .properties(Map.of(
                                                "firstParam", Schema.builder().type("number").description("First number").build(),
                                                "secondParam", Schema.builder().type("number").description("Second number").build()))
                                        .required(Arrays.asList("firstParam", "secondParam"))
                                        .build())
                        .build();

        FunctionDeclaration divideFunction =
                FunctionDeclaration.builder()
                        .name("divideNumbers")
                        .parameters(
                                Schema.builder()
                                        .type("object")
                                        .properties(Map.of(
                                                "firstParam", Schema.builder().type("number").description("First number").build(),
                                                "secondParam", Schema.builder().type("number").description("Second number").build()))
                                        .required(Arrays.asList("firstParam", "secondParam"))
                                        .build())
                        .build();

        GenerateContentConfig config = GenerateContentConfig.builder()
                .toolConfig(ToolConfig.builder().functionCallingConfig(
                        FunctionCallingConfig.builder().mode("ANY").build()
                ).build())
                .tools(
                        Collections.singletonList(
                                Tool.builder().functionDeclarations(
                                        Arrays.asList(
                                                addFunction,
                                                subtractFunction,
                                                divideFunction,
                                                multiplyFunction
                                        )
                                ).build()

                        )
                )
                .build();

        GenerateContentResponse response;
        try {
            response = client.models.generateContent("gemini-2.0-flash", "I have 57 cats, each owns 44 mittens, how many mittens is that in total?", config);
        } catch (HttpException e) {
            throw new RuntimeException(e);
        }

        if (response.functionCalls() == null || response.functionCalls().isEmpty()) {
            System.err.println("No function call received.");
            return;
        }

        var functionCall = response.functionCalls().getFirst();
        String functionName = functionCall.name().get();
        var arguments = functionCall.args();

        Map<String, BiFunction<Double, Double, Double>> functionMapping = new HashMap<>();
        functionMapping.put("addNumbers", (a, b) -> a + b);
        functionMapping.put("subtractNumbers", (a, b) -> a - b);
        functionMapping.put("multiplyNumbers", (a, b) -> a * b);
        functionMapping.put("divideNumbers", (a, b) -> b != 0 ? a / b : Double.NaN);

        BiFunction<Double, Double, Double> function = functionMapping.get(functionName);

        Number firstParam = (Number) arguments.get().get("firstParam");
        Number secondParam = (Number) arguments.get().get("secondParam");
        Double result = function.apply(firstParam.doubleValue(), secondParam.doubleValue());

        System.out.println(result);
        // [END function_calling]
    }
}
