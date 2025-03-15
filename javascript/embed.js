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

import { GoogleGenAI } from "@google/genai";

export async function embedContent() {
  // [START embed_content]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const text = "Hello World!";
  const result = await ai.models.embedContent({
    model: "text-embedding-004",
    contents: text,
    config: { outputDimensionality: 10 },
  });
  console.log(result.embeddings);
  // [END embed_content]
  return result.embeddings;
}

export async function batchEmbedContents() {
  // [START batch_embed_content]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const texts = [
    "What is the meaning of life?",
    "How much wood would a woodchuck chuck?",
    "How does the brain work?",
  ];
  const result = await ai.models.embedContent({
    model: "text-embedding-004",
    contents: texts,
    config: { outputDimensionality: 10 },
  });
  console.log(result.embeddings);
  // [END batch_embed_content]
  return result.embeddings;
}
