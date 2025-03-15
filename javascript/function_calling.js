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

import { GoogleGenAI, FunctionCallingConfigMode } from "@google/genai";

export async function functionCallingTest() {
  // [START function_calling]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  /**
   * The add function returns the sum of two numbers.
   * @param {number} a
   * @param {number} b
   * @returns {number}
   */
  function add(a, b) {
    return a + b;
  }

  /**
   * The subtract function returns the difference (a - b).
   * @param {number} a
   * @param {number} b
   * @returns {number}
   */
  function subtract(a, b) {
    return a - b;
  }

  /**
   * The multiply function returns the product of two numbers.
   * @param {number} a
   * @param {number} b
   * @returns {number}
   */
  function multiply(a, b) {
    return a * b;
  }

  /**
   * The divide function returns the quotient of a divided by b.
   * @param {number} a
   * @param {number} b
   * @returns {number}
   */
  function divide(a, b) {
    return a / b;
  }

  const addDeclaration = {
    name: "addNumbers",
    parameters: {
      type: "object",
      description: "Return the result of adding two numbers.",
      properties: {
        firstParam: {
          type: "number",
          description:
            "The first parameter which can be an integer or a floating point number.",
        },
        secondParam: {
          type: "number",
          description:
            "The second parameter which can be an integer or a floating point number.",
        },
      },
      required: ["firstParam", "secondParam"],
    },
  };

  const subtractDeclaration = {
    name: "subtractNumbers",
    parameters: {
      type: "object",
      description:
        "Return the result of subtracting the second number from the first.",
      properties: {
        firstParam: {
          type: "number",
          description: "The first parameter.",
        },
        secondParam: {
          type: "number",
          description: "The second parameter.",
        },
      },
      required: ["firstParam", "secondParam"],
    },
  };

  const multiplyDeclaration = {
    name: "multiplyNumbers",
    parameters: {
      type: "object",
      description: "Return the product of two numbers.",
      properties: {
        firstParam: {
          type: "number",
          description: "The first parameter.",
        },
        secondParam: {
          type: "number",
          description: "The second parameter.",
        },
      },
      required: ["firstParam", "secondParam"],
    },
  };

  const divideDeclaration = {
    name: "divideNumbers",
    parameters: {
      type: "object",
      description:
        "Return the quotient of dividing the first number by the second.",
      properties: {
        firstParam: {
          type: "number",
          description: "The first parameter.",
        },
        secondParam: {
          type: "number",
          description: "The second parameter.",
        },
      },
      required: ["firstParam", "secondParam"],
    },
  };

  // Step 1: Call generateContent with function calling enabled.
  const generateContentResponse = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents:
      "I have 57 cats, each owns 44 mittens, how many mittens is that in total?",
    config: {
      toolConfig: {
        functionCallingConfig: {
          mode: FunctionCallingConfigMode.ANY,
        },
      },
      tools: [
        {
          functionDeclarations: [
            addDeclaration,
            subtractDeclaration,
            multiplyDeclaration,
            divideDeclaration,
          ],
        },
      ],
    },
  });

  // Step 2: Extract the function call.(
  // Assuming the response contains a 'functionCalls' array.
  const functionCall =
    generateContentResponse.functionCalls &&
    generateContentResponse.functionCalls[0];
  console.log(functionCall);

  // Parse the arguments.
  const args = functionCall.args;
  // Expected args format: { firstParam: number, secondParam: number }

  // Step 3: Invoke the actual function based on the function name.
  const functionMapping = {
    addNumbers: add,
    subtractNumbers: subtract,
    multiplyNumbers: multiply,
    divideNumbers: divide,
  };
  const func = functionMapping[functionCall.name];
  if (!func) {
    console.error("Unimplemented error:", functionCall.name);
    return generateContentResponse;
  }
  const resultValue = func(args.firstParam, args.secondParam);
  console.log("Function result:", resultValue);

  // Step 4: Use the chat API to send the result as the final answer.
  const chat = ai.chats.create({ model: "gemini-2.0-flash" });
  const chatResponse = await chat.sendMessage({
    message: "The final result is " + resultValue,
  });
  console.log(chatResponse.text);
  // [START function_calling]
  return chatResponse;
}
