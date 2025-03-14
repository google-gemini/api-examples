/**
 * Copyright 2024 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START gemini_javascript_configure_model_parameters_test]
import {jest} from '@jest/globals';
import {configureModelParameters} from './configure_model_parameters.js';

describe('Configure Model Parameters Example', () => {
  const originalEnv = process.env;

  beforeEach(() => {
    jest.resetModules();
    process.env = {...originalEnv};
    process.env.GEMINI_API_KEY = 'test-api-key';
  });

  afterEach(() => {
    process.env = originalEnv;
  });

  test('should throw error if API key is not set', async () => {
    delete process.env.GEMINI_API_KEY;
    
    await expect(async () => {
      await configureModelParameters();
    }).rejects.toThrow();
  });

  // Note: Add more tests as needed for specific functionality
  // Currently keeping tests minimal as they would require mocking the Gemini API
});
// [END gemini_javascript_configure_model_parameters_test] 