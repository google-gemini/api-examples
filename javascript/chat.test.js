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
import { test, describe } from "node:test";
import { chat, chatStreaming, chatStreamingWithImages } from "./chat.js";

describe("chat", () => {
  test("chat", async () => {
    const { response1, response2 } = await chat();
    assert.ok(response1 && response1.length > 0);
    assert.ok(response2 && response2.length > 0);
  });

  test("chatStreaming", async () => {
    const result = await chatStreaming();
    assert.strictEqual(result, true);
  });

  test("chatStreamingWithImages", async () => {
    const result = await chatStreamingWithImages();
    assert.strictEqual(result, true);
  });
});
