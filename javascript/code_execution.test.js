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

import assert from 'node:assert';
import {test, describe} from 'node:test';
import {
  codeExecutionBasic,
  codeExecutionRequestOverride,
  codeExecutionChat,
} from './code_execution.js';

describe('code_execution', () => {
  test('codeExecutionBasic', async () => {
    const result = await codeExecutionBasic();
    // Check that the response contains parts and non-empty results.
    assert.ok(Array.isArray(result.parts) && result.parts.length > 0);
    assert.ok(result.text.length > 0);
  });

  test('codeExecutionRequestOverride', async () => {
    const result = await codeExecutionRequestOverride();
    // Check that the response contains non-empty results.
    assert.ok(result.executableCode.length > 0);
    assert.ok(result.codeExecutionResult.length > 0);
  });

  test('codeExecutionChat', async () => {
    const result = await codeExecutionChat();
    // Check that the response contains non-empty results.
    assert.ok(result.executableCode.length > 0);
    assert.ok(result.codeExecutionResult.length > 0);
  });
});
