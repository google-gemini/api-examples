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

import {GoogleGenAI} from '@google/genai';

export async function configureModelParameters() {
  // [START configure_model_parameters]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: 'Tell me a story about a magic backpack.',
    config: {
      candidateCount: 1,
      stopSequences: ['x'],
      maxOutputTokens: 20,
      temperature: 1.0,
    },
  });

  console.log(response.text);
  // [END configure_model_parameters]

  return response.text;
}
