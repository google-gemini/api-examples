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
import { embedContent, batchEmbedContents } from "./embed.js";

describe("embed", () => {
  test("embedContent returns a valid embedding", async () => {
    const embeddings = await embedContent();
    // Ensure the result is an array and non-empty.
    assert.ok(Array.isArray(embeddings));
    assert.ok(embeddings.length > 0);
  });

  test("batchEmbedContents returns embeddings for all inputs", async () => {
    const embeddings = await batchEmbedContents();
    // Expect an array with three embeddings, one per input text.
    assert.ok(Array.isArray(embeddings));
    assert.strictEqual(embeddings.length, 3);
  });
});
