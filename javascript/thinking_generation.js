/**
 * @license
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { GoogleGenAI } from "@google/genai";

// Ensure the API key is set in your environment variables
if (!process.env.GEMINI_API_KEY) {
  throw new Error("GEMINI_API_KEY environment variable not set.");
}

// Define the thinking model centrally
const MODEL_ID = "gemini-2.5-pro-exp-03-25";

const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

export async function thinkingTextOnlyPrompt() {
  // [START thinking_text_only_prompt]
  /**
   * Generates text based on a simple reasoning prompt.
   */
  const prompt =
    "Explain the concept of Occam's Razor and provide a simple, everyday example.";

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
  });

  console.log(response.text); // Direct text access for simple responses

  return response.text;
  // [END thinking_text_only_prompt]
}

export async function thinkingTextOnlyPromptStreaming() {
  // [START thinking_text_only_prompt_streaming]
  /**
   * Generates text based on a simple reasoning prompt using streaming.
   */
  const prompt =
    "Explain the concept of Occam's Razor and provide a simple, everyday example.";

  const response = await ai.models.generateContentStream({
    model: MODEL_ID,
    contents: prompt,
  });

  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END thinking_text_only_prompt_streaming]
  return text;
}

export async function thinkingLogicPuzzle() {
  // [START thinking_logic_puzzle]
  /**
   * Tests the model's ability to solve a classic logic puzzle.
   */
  const prompt = `
        Solve this logic puzzle and explain your reasoning step-by-step:
        There are three boxes. One contains only apples, one contains only oranges,
        and one contains both apples and oranges. The boxes have been incorrectly
        labeled such that no label is correct. You are allowed to draw one fruit
        from only one box of your choosing (without looking inside). Which box
        would you draw from to correctly label all boxes, and why?
        `;

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
  });

  console.log(response.text);

  return response.text;
  // [END thinking_logic_puzzle]
}

export async function thinkingCodeExplanation() {
  // [START thinking_code_explanation]
  /**
   * Tests the model's ability to understand and explain code.
   */
  const prompt = `
        Explain this Python code snippet step-by-step, including what it does
        and why recursion is used here:

        def factorial(n):
            '''Calculates the factorial of a non-negative integer.'''
            if not isinstance(n, int) or n < 0:
                raise ValueError("Input must be a non-negative integer")
            if n == 0:
                return 1
            else:
                return n * factorial(n-1)

        result = factorial(5)
        print(f"The factorial of 5 is: {result}")
        `;

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
  });

  console.log(response.text);

  return response.text;
  // [END thinking_code_explanation]
}

export async function thinkingCreativeWritingConstraints() {
  // [START thinking_creative_writing_constraints]
  /**
   * Tests creative writing with specific constraints.
   */
  const prompt = `
        Write a short story (max 150 words) about a detective investigating a
        mystery in a library, but the story must not contain the letter 'e'.
        `;
  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
  });

  console.log(response.text);

  return response.text;
  // [END thinking_creative_writing_constraints]
}

export async function thinkingWithSearchTool() {
  // [START thinking_with_search_tool]
  /**
   * Uses the search tool to answer a question requiring current information.
   */
  // Define the search tool. The empty object {} signifies default web search.
  const googleSearchTool = {
    googleSearch: {},
  };

  const prompt = "What were the major scientific breakthroughs announced last week?";

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
    config: {
      tools: [googleSearchTool], // Pass the tool in the config
    },
  });

  console.log(response);

  return response.text;
  // [END thinking_with_search_tool]
}


export async function thinkingWithSearchToolStreaming() {
  // [START thinking_with_search_tool_streaming]
  /**
   * Uses the search tool with streaming output and retrieves grounding.
   */
  const googleSearchTool = {
    googleSearch: {},
  };

  const prompt = "When is the next total solar eclipse visible from mainland Europe?";

  const response = await ai.models.generateContentStream({
    model: MODEL_ID,
    contents: prompt,
    config: {
      tools: [googleSearchTool],
    },
  });

  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }

  return text;
  // [END thinking_with_search_tool_streaming]
}


export async function thinkingCodeExecution() {
  // [START thinking_code_execution]
  /**
   * Tests the model's ability to generate and execute code.
   */
  const prompt =
    "What is the sum of the first 50 prime numbers? " +
    "Generate and run Python code for the calculation, and make sure you get all 50. " +
    "Provide the final sum clearly.";

  // Define the code execution tool
  const codeExecutionTool = {
    codeExecution: {},
  };

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
    config: {
      tools: [codeExecutionTool],
    },
  });

  console.log(response);

  return response.text;
  // [END thinking_code_execution]
}

export async function thinkingStructuredOutputJson() {
  // [START thinking_structured_output_json]
  /**
   * Tests the model's ability to generate structured JSON output via prompting.
   */
  // Prompt clearly asks for JSON and provides the schema inline
  const prompt = `
        Provide a list of 3 famous physicists and their key contributions
        in JSON format.

        Use this JSON schema:

        Physicist = {'name': str, 'contribution': str, 'era': str}
        Return: list[Physicist]
        `;

  const response = await ai.models.generateContent({
    model: MODEL_ID,
    contents: prompt,
    // For stricter JSON mode (if structured output is supported):
    // config: {
    //   responseMimeType: "application/json",
    //   // jsonSchema: { type: "array", items: { type: "object", properties: { ... } } } // Define schema if needed
    // }
  });

  console.log(response);

  return response.text;
  // [END thinking_structured_output_json]
}
