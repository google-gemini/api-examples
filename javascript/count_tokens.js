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

import {
  GoogleGenAI,
  createUserContent,
  createPartFromUri,
  createPartFromBase64,
} from '@google/genai';
import path from 'path';
import fs from 'fs';
import {fileURLToPath} from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const media = path.join(__dirname, '..', 'third_party');

// A simple sleep helper.
const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

export async function tokensTextOnly() {
  // [START tokens_text_only]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt = 'The quick brown fox jumps over the lazy dog.';
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log(generateResponse.usageMetadata);
  // [END tokens_text_only]
  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensChat() {
  // [START tokens_chat]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  // Initial chat history.
  const history = [
    {role: 'user', parts: [{text: 'Hi my name is Bob'}]},
    {role: 'model', parts: [{text: 'Hi Bob!'}]},
  ];
  const chat = ai.chats.create({
    model: 'gemini-2.0-flash',
    history: history,
  });

  // Count tokens for the current chat history.
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: chat.getHistory(),
  });
  console.log(countTokensResponse.totalTokens);

  const chatResponse = await chat.sendMessage({
    message: 'In one sentence, explain how a computer works to a young child.',
  });
  console.log(chatResponse.usageMetadata);

  // Add an extra user message to the history.
  const extraMessage = {
    role: 'user',
    parts: [{text: 'What is the meaning of life?'}],
  };
  const combinedHistory = chat.getHistory();
  combinedHistory.push(extraMessage);
  const combinedCountTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: combinedHistory,
  });
  console.log(
    'Combined history token count:',
    combinedCountTokensResponse.totalTokens,
  );
  // [END tokens_chat]
  return {
    historyTokenCount: countTokensResponse.totalTokens,
    usage: chatResponse.usageMetadata,
    combinedTokenCount: combinedCountTokensResponse.totalTokens,
  };
}

export async function tokensMultimodalImageInline() {
  // [START tokens_multimodal_image_inline]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt = 'Tell me about this image';
  const imageBuffer = fs.readFileSync(path.join(media, 'organ.jpg'));

  // Convert buffer to base64 string.
  const imageBase64 = imageBuffer.toString('base64');

  // Build contents using createUserContent and createPartFromBase64.
  const contents = createUserContent([
    prompt,
    createPartFromBase64(imageBase64, 'image/jpeg'),
  ]);

  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: contents,
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: contents,
  });
  console.log(generateResponse.usageMetadata);
  // [END tokens_multimodal_image_inline]

  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensMultimodalImageFileApi() {
  // [START tokens_multimodal_image_file_api]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt = 'Tell me about this image';
  const organ = await ai.files.upload({
    file: path.join(media, 'organ.jpg'),
    config: {mimeType: 'image/jpeg'},
  });

  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(organ.uri, organ.mimeType),
    ]),
  });
  console.log(generateResponse.usageMetadata);
  // [END tokens_multimodal_image_file_api]
  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensMultimodalVideoAudioFileApi() {
  // [START tokens_multimodal_video_audio_file_api]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt = 'Tell me about this video';
  let videoFile = await ai.files.upload({
    file: path.join(media, 'Big_Buck_Bunny.mp4'),
    config: {mimeType: 'video/mp4'},
  });

  // Poll until the video file is completely processed (state becomes ACTIVE).
  while (!videoFile.state || videoFile.state.toString() !== 'ACTIVE') {
    console.log('Processing video...');
    console.log('File state: ', videoFile.state);
    await sleep(5000);
    videoFile = await ai.files.get({name: videoFile.name});
  }

  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(videoFile.uri, videoFile.mimeType),
    ]),
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(videoFile.uri, videoFile.mimeType),
    ]),
  });
  console.log(generateResponse.usageMetadata);
  // [END tokens_multimodal_video_audio_file_api]

  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensMultimodalPdfFileApi() {
  // [START tokens_multimodal_pdf_file_api]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const samplePdf = await ai.files.upload({
    file: path.join(media, 'test.pdf'),
    config: {mimeType: 'application/pdf'},
  });
  const prompt = 'Give me a summary of this document.';
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(samplePdf.uri, samplePdf.mimeType),
    ]),
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents: createUserContent([
      prompt,
      createPartFromUri(samplePdf.uri, samplePdf.mimeType),
    ]),
  });
  console.log(generateResponse.usageMetadata);
  // [START tokens_multimodal_pdf_file_api]
  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensCachedContent() {
  // [START tokens_cached_content]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const textFile = await ai.files.upload({
    file: path.join(media, 'a11.txt'),
    config: {mimeType: 'text/plain'},
  });

  const cache = await ai.caches.create({
    model: 'gemini-1.5-flash-001',
    config: {
      contents: createUserContent([
        'Here the Apollo 11 transcript:',
        createPartFromUri(textFile.uri, textFile.mimeType),
      ]),
      system_instruction: null,
      tools: null,
    },
  });

  const prompt = 'Please give a short summary of this file.';
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log(countTokensResponse.totalTokens);

  const generateResponse = await ai.models.generateContent({
    model: 'gemini-1.5-flash-001',
    contents: prompt,
    config: {cachedContent: cache.name},
  });
  console.log(generateResponse.usageMetadata);

  await ai.caches.delete({name: cache.name});
  // [START tokens_cached_content]
  return {
    totalTokens: countTokensResponse.totalTokens,
    usage: generateResponse.usageMetadata,
  };
}

export async function tokensSystemInstruction() {
  // [START tokens_system_instruction]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt = 'The quick brown fox jumps over the lazy dog.';
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log(
    'base_count (no system instruction):',
    countTokensResponse.totalTokens,
  );

  const countTokensWithSysInstResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: prompt,
    config: {systemInstruction: 'You are a cat. Your name is Neko.'},
  });
  console.log(
    'total_tokens (with system instruction):',
    countTokensWithSysInstResponse.totalTokens,
  );
  // [START tokens_system_instruction]
  return {
    responseTokenCount: countTokensResponse.totalTokens,
    sysInstResponseTokenCount: countTokensWithSysInstResponse.totalTokens,
  };
}

// Error: tools parameter is not supported in Gemini API.
export async function tokensTools() {
  // [START tokens_tools]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});
  const prompt =
    'I have 57 cats, each owns 44 mittens, how many mittens is that in total?';
  const countTokensResponse = await ai.models.countTokens({
    model: 'gemini-2.0-flash',
    contents: prompt,
  });
  console.log('total_tokens (no tools):', countTokensResponse.totalTokens);

  // Define function declarations for arithmetic operations.
  const addDeclaration = {
    name: 'addNumbers',
    parameters: {
      type: 'object',
      description: 'Return the result of adding two numbers.',
      properties: {
        firstParam: {
          type: 'number',
          description:
            'The first parameter which can be an integer or a floating point number.',
        },
        secondParam: {
          type: 'number',
          description:
            'The second parameter which can be an integer or a floating point number.',
        },
      },
      required: ['firstParam', 'secondParam'],
    },
  };

  const subtractDeclaration = {
    name: 'subtractNumbers',
    parameters: {
      type: 'object',
      description:
        'Return the result of subtracting the second number from the first.',
      properties: {
        firstParam: {
          type: 'number',
          description: 'The first parameter.',
        },
        secondParam: {
          type: 'number',
          description: 'The second parameter.',
        },
      },
      required: ['firstParam', 'secondParam'],
    },
  };

  const multiplyDeclaration = {
    name: 'multiplyNumbers',
    parameters: {
      type: 'object',
      description: 'Return the product of two numbers.',
      properties: {
        firstParam: {
          type: 'number',
          description: 'The first parameter.',
        },
        secondParam: {
          type: 'number',
          description: 'The second parameter.',
        },
      },
      required: ['firstParam', 'secondParam'],
    },
  };

  const divideDeclaration = {
    name: 'divideNumbers',
    parameters: {
      type: 'object',
      description:
        'Return the quotient of dividing the first number by the second.',
      properties: {
        firstParam: {
          type: 'number',
          description: 'The first parameter.',
        },
        secondParam: {
          type: 'number',
          description: 'The second parameter.',
        },
      },
      required: ['firstParam', 'secondParam'],
    },
  };

  const tools = {
    functionDeclarations: [
      addDeclaration,
      subtractDeclaration,
      multiplyDeclaration,
      divideDeclaration,
    ],
  };

  // Uncomment if the API gets support for tools in count_tokens.
  /*
  const countTokensWithToolsResponse = await ai.models.countTokens({
    model: "gemini-2.0-flash",
    contents: prompt,
    config: { tools: tools }
  });
  console.log("total_tokens (with tools):", countTokensWithToolsResponse.totalTokens);
  */
  // [END tokens_tools]
  return {
    responseTokenCount: countTokensResponse.totalTokens,
    // Uncomment if the API gets support for tools in count_tokens.
    // toolsResponseTokenCount : countTokensWithToolsResponse.totalTokens
  };
}
