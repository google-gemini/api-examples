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
  filesCreateText,
  filesCreateImage,
  filesCreateAudio,
  filesCreateVideo,
  filesCreatePdf,
  filesList,
  filesGet,
} from "./files.js";

describe("Files API Tests", () => {
  test("filesCreateText", async () => {
    const text = await filesCreateText();
    assert.ok(text && text.length > 0);
  });

  test("filesCreateImage", async () => {
    const text = await filesCreateImage();
    assert.ok(text && text.length > 0);
  });

  test("filesCreateAudio", async () => {
    const text = await filesCreateAudio();
    assert.ok(text && text.length > 0);
  });

  test("filesCreateVideo", async () => {
    const text = await filesCreateVideo();
    assert.ok(text && text.length > 0);
  });

  test("filesCreatePdf", async () => {
    const text = await filesCreatePdf();
    assert.ok(text && text.length > 0);
  });

  test("filesList", async () => {
    const names = await filesList();
    assert.ok(Array.isArray(names) && names.length > 0);
  });

  test("filesGet", async () => {
    const file = await filesGet();
    assert.ok(file && file.name);
  });
});
