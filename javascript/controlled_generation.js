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

import {GoogleGenAI, createUserContent, createPartFromUri} from '@google/genai';
import path from 'path';
import {fileURLToPath} from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const media = path.join(__dirname, '..', 'third_party');

export async function jsonControlledGeneration() {
  // [START json_controlled_generation]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: 'List a few popular cookie recipes.',
    config: {
      responseMimeType: 'application/json',
      responseSchema: {
        type: 'array',
        items: {
          type: 'object',
          properties: {
            recipeName: {type: 'string'},
            ingredients: {type: 'array', items: {type: 'string'}},
          },
          required: ['recipeName', 'ingredients'],
        },
      },
    },
  });
  console.log(response.text);
  // [END json_controlled_generation]
  return response;
}

export async function jsonNoSchema() {
  // [START json_no_schema]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt =
    'List a few popular cookie recipes in JSON format.\n\n' +
    'Use this JSON schema:\n\n' +
    "Recipe = {'recipeName': str, 'ingredients': list[str]}\n" +
    'Return: list[Recipe]';
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log(response.text);
  // [END json_no_schema]
  return response;
}

export async function jsonEnum() {
  // [START json_enum]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const imagePath = path.join(media, 'organ.jpg');
  const organ = await ai.files.upload({
    file: imagePath,
    config: {mimeType: 'image/jpeg'},
  });
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      'What kind of instrument is this?',
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
    config: {
      responseMimeType: 'application/json',
      responseSchema: {
        type: 'string',
        enum: ['Percussion', 'String', 'Woodwind', 'Brass', 'Keyboard'],
      },
    },
  });
  console.log(response.text);
  // [END json_enum]
  return response;
}

export async function enumInJson() {
  // [START enum_in_json]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: 'List about 10 cookie recipes, grade them based on popularity',
    config: {
      responseMimeType: 'application/json',
      responseSchema: {
        type: 'array',
        items: {
          type: 'object',
          properties: {
            recipeName: {type: 'string'},
            grade: {type: 'string', enum: ['a+', 'a', 'b', 'c', 'd', 'f']},
          },
          required: ['recipeName', 'grade'],
        },
      },
    },
  });
  console.log(response.text);
  // [END enum_in_json]
  return response;
}

export async function jsonEnumRaw() {
  // [START json_enum_raw]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const imagePath = path.join(media, 'organ.jpg');
  const organ = await ai.files.upload({
    file: imagePath,
    config: {mimeType: 'image/jpeg'},
  });
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      'What kind of instrument is this?',
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
    config: {
      responseMimeType: 'application/json',
      responseSchema: {
        type: 'string',
        enum: ['Percussion', 'String', 'Woodwind', 'Brass', 'Keyboard'],
      },
    },
  });
  console.log(response.text);
  // [END json_enum_raw]
  return response;
}

export async function xEnum() {
  // [START x_enum]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const imagePath = path.join(media, 'organ.jpg');
  const organ = await ai.files.upload({
    file: imagePath,
    config: {mimeType: 'image/jpeg'},
  });
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      'What kind of instrument is this?',
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
    config: {
      responseMimeType: 'text/x.enum',
      responseSchema: {
        type: 'string',
        enum: ['Percussion', 'String', 'Woodwind', 'Brass', 'Keyboard'],
      },
    },
  });
  console.log(response.text);
  // [END x_enum]
  return response;
}

export async function xEnumRaw() {
  // [START x_enum_raw]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const imagePath = path.join(media, 'organ.jpg');
  const organ = await ai.files.upload({
    file: imagePath,
    config: {mimeType: 'image/jpeg'},
  });
  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      'What kind of instrument is this?',
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
    config: {
      responseMimeType: 'text/x.enum',
      responseSchema: {
        type: 'string',
        enum: ['Percussion', 'String', 'Woodwind', 'Brass', 'Keyboard'],
      },
    },
  });
  console.log(response.text);
  // [END x_enum_raw]
  return response;
}
