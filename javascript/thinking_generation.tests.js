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

import assert from "node:assert";
import { test, describe, before } from "node:test";

// Import the functions to test
import {
  thinkingTextOnlyPrompt,
  thinkingTextOnlyPromptStreaming,
  thinkingLogicPuzzle,
  thinkingCodeExplanation,
  thinkingCreativeWritingConstraints,
  thinkingWithSearchTool,
  thinkingWithSearchToolStreaming,
  thinkingCodeExecution,
  thinkingStructuredOutputJson,
} from "./thinking_generation.js"; 

// Simple delay function for potential rate limiting issues between tests
const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));
const TEST_DELAY = 1000; // 1 second delay between tests

describe("thinking_generation", { timeout: 300000 }, () => {
  // Check for API key before running tests
  before(() => {
    if (!process.env.GEMINI_API_KEY) {
      throw new Error(
        "GEMINI_API_KEY environment variable not set. Tests cannot run."
      );
    }
  });

  test("thinkingTextOnlyPrompt", async () => {
    const result = await thinkingTextOnlyPrompt();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    await sleep(TEST_DELAY);
  });

  test("thinkingTextOnlyPromptStreaming", async () => {
    const result = await thinkingTextOnlyPromptStreaming();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    await sleep(TEST_DELAY);
  });

  test("thinkingLogicPuzzle", async () => {
    const result = await thinkingLogicPuzzle();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    await sleep(TEST_DELAY);
  });

  test("thinkingCodeExplanation", async () => {
    const result = await thinkingCodeExplanation();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    await sleep(TEST_DELAY);
  });

  test("thinkingCreativeWritingConstraints", async () => {
    const result = await thinkingCreativeWritingConstraints();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    await sleep(TEST_DELAY);
  });

  test("thinkingWithSearchTool", async () => {
    const result = await thinkingWithSearchTool();
    // Search results can sometimes be empty if nothing relevant is found *right now*
    // A basic check that the function ran and returned *something* (even if just explanatory text)
    assert.ok(result !== undefined, "Test failed: Function did not return");
    // A more robust test might check if grounding metadata was accessed or if specific keywords appeared.
    await sleep(TEST_DELAY);
  });

   test("thinkingWithSearchToolStreaming", async () => {
    const result = await thinkingWithSearchToolStreaming();
    assert.ok(result !== undefined, "Test failed: Function did not return");
    await sleep(TEST_DELAY);
  });

  test("thinkingCodeExecution", async () => {
    const result = await thinkingCodeExecution();
    assert.ok(result !== undefined, "Test failed: Function did not return");
    await sleep(TEST_DELAY);
  });

  test("thinkingStructuredOutputJson", async () => {
    const result = await thinkingStructuredOutputJson();
    assert.ok(
      result && result.length > 0,
      "Test failed: No result or empty result"
    );
    assert.doesNotThrow(() => {
        const jsonMatch = result.match(/```json\s*([\s\S]*?)\s*```/);
        const jsonToParse = jsonMatch ? jsonMatch[1] : result;
        JSON.parse(jsonToParse);
     }, "Test failed: Result could not be parsed as JSON");
    await sleep(TEST_DELAY);
  });
});
