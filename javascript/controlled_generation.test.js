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
  jsonControlledGeneration,
  jsonNoSchema,
  jsonEnum,
  enumInJson,
  jsonEnumRaw,
  xEnum,
  xEnumRaw,
} from "./controlled_generation.js";

describe("controlled_generation", () => {
  test("jsonControlledGeneration", async () => {
    const result = await jsonControlledGeneration();
    assert.ok(result);
  });

  test("jsonNoSchema", async () => {
    const result = await jsonNoSchema();
    assert.ok(result);
  });

  test("jsonEnum", async () => {
    const result = await jsonEnum();
    assert.ok(typeof result.text === "string");
    assert.ok(result.text.includes("Keyboard"));
  });

  test("enumInJson", async () => {
    const result = await enumInJson();
    const parsed = result.parsed || [];
    assert.ok(Array.isArray(parsed));
  });

  test("jsonEnumRaw", async () => {
    const result = await jsonEnumRaw();
    assert.ok(typeof result.text === "string");
    assert.ok(result.text.includes("Keyboard"));
  });

  test("xEnum", async () => {
    const result = await xEnum();
    assert.ok(typeof result.text === "string");
    assert.ok(result.text.includes("Keyboard"));
  });

  test("xEnumRaw", async () => {
    const result = await xEnumRaw();
    assert.ok(typeof result.text === "string");
    assert.ok(result.text.includes("Keyboard"));
  });
});
