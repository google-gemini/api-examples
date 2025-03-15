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
import { safetySettings, safetySettingsMulti } from "./safety_settings.js";

describe("safety_settings", () => {
  test("safetySettings", async () => {
    const response = await safetySettings();
    assert.ok(response.text.length > 0);
    assert.ok(
      response.candidates[0].safetyRatings,
      "Safety ratings should be present",
    );
  });

  test("safetySettingsMulti", async () => {
    const response = await safetySettingsMulti();
    assert.ok(response.text.length > 0);
    // Generated text might be empty if blocked, so we test for safety ratings.
    assert.ok(
      response.candidates[0].safetyRatings,
      "Safety ratings should be present",
    );
  });
});
