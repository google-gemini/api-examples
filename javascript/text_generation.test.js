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
import {describe, test} from 'node:test';

import {textGenMultimodalAudio, textGenMultimodalAudioStreaming, textGenMultimodalMultiImagePrompt, textGenMultimodalMultiImagePromptStreaming, textGenMultimodalOneImagePrompt, textGenMultimodalOneImagePromptStreaming, textGenMultimodalPdf, textGenMultimodalPdfStreaming, textGenMultimodalVideoPrompt, textGenMultimodalVideoPromptStreaming, textGenTextOnlyPrompt, textGenTextOnlyPromptStreaming,} from './text_generation.js';

describe('text_generation', () => {
  test('textGenTextOnlyPrompt', async () => {
    const text = await textGenTextOnlyPrompt();
    assert.ok(text.length > 0);
  });

  test('textGenTextOnlyPromptStreaming', async () => {
    const text = await textGenTextOnlyPromptStreaming();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalOneImagePrompt', async () => {
    const text = await textGenMultimodalOneImagePrompt();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalOneImagePromptStreaming', async () => {
    const text = await textGenMultimodalOneImagePromptStreaming();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalMultiImagePrompt', async () => {
    const text = await textGenMultimodalMultiImagePrompt();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalMultiImagePromptStreaming', async () => {
    const text = await textGenMultimodalMultiImagePromptStreaming();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalAudio', async () => {
    const text = await textGenMultimodalAudio();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalAudioStreaming', async () => {
    const text = await textGenMultimodalAudioStreaming();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalVideoPrompt', async () => {
    const text = await textGenMultimodalVideoPrompt();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalVideoPromptStreaming', async () => {
    const text = await textGenMultimodalVideoPromptStreaming();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalPdf', async () => {
    const text = await textGenMultimodalPdf();
    assert.ok(text.length > 0);
  });

  test('textGenMultimodalPdfStreaming', async () => {
    const text = await textGenMultimodalPdfStreaming();
    assert.ok(text.length > 0);
  });
});
