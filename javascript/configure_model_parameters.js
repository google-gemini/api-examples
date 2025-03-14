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

// [START gemini_javascript_configure_model_parameters]
import {GoogleGenAI} from '@google/genai';

/**
 * Configure model parameters example with the Gemini API
 * 
 * This example demonstrates how to configure generation parameters:
 * - candidate_count: Number of candidates to generate
 * - stop_sequences: Sequences that stop generation
 * - max_output_tokens: Maximum number of tokens to generate
 * - temperature: Controls randomness of generation
 */
export async function configureModelParameters() {
  try {
    // Initialize the client with your API key
    const GEMINI_API_KEY = process.env.GEMINI_API_KEY;
    const ai = new GoogleGenAI({apiKey: GEMINI_API_KEY});

    // Get the Gemini model
    const model = ai.getGenerativeModel({model: 'gemini-pro'});

    // Configure generation parameters
    const generationConfig = {
      candidateCount: 1,
      stopSequences: ['x'],
      maxOutputTokens: 20,
      temperature: 1.0,
    };

    // Generate content with configured parameters
    const result = await model.generateContent({
      contents: [{text: 'Tell me a story about a magic backpack.'}],
      generationConfig,
    });

    console.log('Generated text:', result.response.text());
  } catch (error) {
    console.error('Error:', error);
  }
}

// Run the example if this file is executed directly
if (process.argv[1] === new URL(import.meta.url).pathname) {
  if (!process.env.GEMINI_API_KEY) {
    console.error('Please set the GEMINI_API_KEY environment variable');
    process.exit(1);
  }

  configureModelParameters().catch(console.error);
}
// [END gemini_javascript_configure_model_parameters] 