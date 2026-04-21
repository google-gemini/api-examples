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
  createPartFromUri 
} from "@google/genai";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const media = path.join(__dirname, '..', 'third_party');

// Sleep helper for video polling.
const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

export async function textGenTextOnlyPrompt() {
  // [START text_gen_text_only_prompt]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: "Write a story about a magic backpack.",
  });
  console.log(response.text);
  // [END text_gen_text_only_prompt]
  return response.text;
}

export async function textGenTextOnlyPromptStreaming() {
  // [START text_gen_text_only_prompt_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: "Write a story about a magic backpack.",
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_text_only_prompt_streaming]
  return text;
}

export async function textGenMultimodalOneImagePrompt() {
  // [START text_gen_multimodal_one_image_prompt]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const organ = await ai.files.upload({
    file: path.join(media, "organ.jpg"),
  });

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Tell me about this instrument", 
        createPartFromUri(organ.uri, organ.mimeType)
      ]),
    ],
  });
  console.log(response.text);
  // [END text_gen_multimodal_one_image_prompt]
  return response.text;
}

export async function textGenMultimodalOneImagePromptStreaming() {
  // [START text_gen_multimodal_one_image_prompt_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const organ = await ai.files.upload({
    file: path.join(media, "organ.jpg"),
  });

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Tell me about this instrument", 
        createPartFromUri(organ.uri, organ.mimeType)
      ]),
    ],
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_multimodal_one_image_prompt_streaming]
  return text;
}

export async function textGenMultimodalMultiImagePrompt() {
  // [START text_gen_multimodal_multi_image_prompt]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  
  const organ = await ai.files.upload({
    file: path.join(media, "organ.jpg"),
  });

  const cajun = await ai.files.upload({
    file: path.join(media, "Cajun_instruments.jpg"),
    config: { mimeType: "image/jpeg" },
  });

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "What is the difference between both of these instruments?",
        createPartFromUri(organ.uri, organ.mimeType),
        createPartFromUri(cajun.uri, cajun.mimeType),
      ]),
    ],
  });
  console.log(response.text);
  // [END text_gen_multimodal_multi_image_prompt]
  return response.text;
}

export async function textGenMultimodalMultiImagePromptStreaming() {
  // [START text_gen_multimodal_multi_image_prompt_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const organ = await ai.files.upload({
    file: path.join(media, "organ.jpg"),
  });

  const cajun = await ai.files.upload({
    file: path.join(media, "Cajun_instruments.jpg"),
  });

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "What is the difference between both of these instruments?",
        createPartFromUri(organ.uri, organ.mimeType),
        createPartFromUri(cajun.uri, cajun.mimeType),
      ]),
    ],
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_multimodal_multi_image_prompt_streaming]
  return text;
}

export async function textGenMultimodalAudio() {
  // [START text_gen_multimodal_audio]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const audio = await ai.files.upload({
    file: path.join(media, "sample.mp3"),
  });
  
  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Give me a summary of this audio file.",
        createPartFromUri(audio.uri, audio.mimeType),
      ]),
    ],
  });
  console.log(response.text);
  // [END text_gen_multimodal_audio]
  return response.text;
}

export async function textGenMultimodalAudioStreaming() {
  // [START text_gen_multimodal_audio_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const audio = await ai.files.upload({
    file: path.join(media, "sample.mp3"),
  });

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Give me a summary of this audio file.",
        createPartFromUri(audio.uri, audio.mimeType),
      ]),
    ],
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_multimodal_audio_streaming]
  return text;
}

export async function textGenMultimodalVideoPrompt() {
  // [START text_gen_multimodal_video_prompt]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  let video = await ai.files.upload({
    file: path.join(media, 'Big_Buck_Bunny.mp4'),
  });

  // Poll until the video file is completely processed (state becomes ACTIVE).
  while (!video.state || video.state.toString() !== 'ACTIVE') {
    console.log('Processing video...');
    console.log('File state: ', video.state);
    await sleep(5000);
    video = await ai.files.get({name: video.name});
  }

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Describe this video clip",
        createPartFromUri(video.uri, video.mimeType),
      ]),
    ],
  });
  console.log(response.text);
  // [END text_gen_multimodal_video_prompt]
  return response.text;
}

export async function textGenMultimodalVideoPromptStreaming() {
  // [START text_gen_multimodal_video_prompt_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  let video = await ai.files.upload({
    file: path.join(media, 'Big_Buck_Bunny.mp4'),
  });

  // Poll until the video file is completely processed (state becomes ACTIVE).
  while (!video.state || video.state.toString() !== 'ACTIVE') {
    console.log('Processing video...');
    console.log('File state: ', video.state);
    await sleep(5000);
    video = await ai.files.get({name: video.name});
  }

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Describe this video clip",
        createPartFromUri(video.uri, video.mimeType),
      ]),
    ],
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_multimodal_video_prompt_streaming]
  return text;
}

export async function textGenMultimodalPdf() {
  // [START text_gen_multimodal_pdf]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const pdf = await ai.files.upload({
    file: path.join(media, "test.pdf"),
  });

  const response = await ai.models.generateContent({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Give me a summary of this document:",
        createPartFromUri(pdf.uri, pdf.mimeType),
      ]),
    ],
  });
  console.log(response.text);
  // [END text_gen_multimodal_pdf]
  return response.text;
}

export async function textGenMultimodalPdfStreaming() {
  // [START text_gen_multimodal_pdf_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

  const pdf = await ai.files.upload({
    file: path.join(media, "test.pdf"),
  });

  const response = await ai.models.generateContentStream({
    model: "gemini-2.5-flash",
    contents: [
      createUserContent([
        "Give me a summary of this document:",
        createPartFromUri(pdf.uri, pdf.mimeType),
      ]),
    ],
  });
  let text = "";
  for await (const chunk of response) {
    console.log(chunk.text);
    text += chunk.text;
  }
  // [END text_gen_multimodal_pdf_streaming]
  return text;
}
