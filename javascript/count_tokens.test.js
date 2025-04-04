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
import {
  tokensTextOnly,
  tokensChat,
  tokensMultimodalImageInline,
  tokensMultimodalImageFileApi,
  tokensMultimodalVideoAudioFileApi,
  tokensMultimodalPdfFileApi,
  tokensCachedContent,
  tokensSystemInstruction,
  tokensTools,
} from "./count_tokens.js";

describe("count_tokens", () => {
  test("tokensTextOnly", async () => {
    const result = await tokensTextOnly();
    assert.ok(result.totalTokens > 0);
    // Check that usage metadata has numeric values.
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensChat", async () => {
    const result = await tokensChat();
    assert.ok(result.historyTokenCount > 0);
    assert.ok(result.combinedTokenCount > result.historyTokenCount);
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensMultimodalImageInline", async () => {
    const result = await tokensMultimodalImageInline();
    assert.ok(result.totalTokens > 0);
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensMultimodalImageFileApi", async () => {
    const result = await tokensMultimodalImageFileApi();
    assert.ok(result.totalTokens > 0);
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensMultimodalVideoAudioFileApi", async () => {
    const result = await tokensMultimodalVideoAudioFileApi();
    assert.ok(result.totalTokens > 0);
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensMultimodalPdfFileApi", async () => {
    const result = await tokensMultimodalPdfFileApi();
    assert.ok(result.totalTokens > 0);
    assert.ok(result.usage.promptTokenCount >= 0);
  });

  test("tokensCachedContent", async () => {
    const result = await tokensCachedContent();
    assert.ok(result.totalTokens > 0);
    assert.ok(result.usage.promptTokenCount >= 0);
  });
});
