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
  cacheCreate,
  cacheCreateFromName,
  cacheCreateFromChat,
  cacheDelete,
  cacheGet,
  cacheList,
  cacheUpdate,
} from "./cache.js";

describe("cache", () => {
  test("cacheCreate", async () => {
    const text = await cacheCreate();
    assert.ok(text && text.length > 0);
  });

  test("cacheCreateFromName", async () => {
    const text = await cacheCreateFromName();
    assert.ok(text && text.length > 0);
  });

  test("cacheCreateFromChat", async () => {
    const text = await cacheCreateFromChat();
    assert.ok(text && text.length > 0);
  });

  test("cacheDelete", async () => {
    // cacheDelete does not return anything; ensure it runs without error.
    await cacheDelete();
  });

  test("cacheGet", async () => {
    const cache = await cacheGet();
    assert.ok(cache && cache.name);
  });

  test("cacheList", async () => {
    // cacheList prints the list; ensure it runs without error.
    await cacheList();
  });

  test("cacheUpdate", async () => {
    const cache = await cacheUpdate();
    assert.ok(cache && cache.name);
  });
});
