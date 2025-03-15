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
} from "@google/genai";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Define common media directory path (was third_party)
const media = path.join(__dirname, "..", "third_party");

export async function chat() {
  // [START chat]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const chat = ai.chats.create({
    model: "gemini-2.0-flash",
    history: [
      {
        role: "user",
        parts: [{ text: "Hello" }],
      },
      {
        role: "model",
        parts: [{ text: "Great to meet you. What would you like to know?" }],
      },
    ],
  });

  const response1 = await chat.sendMessage({
    message: "I have 2 dogs in my house.",
  });
  console.log("Chat response 1:", response1.text);

  const response2 = await chat.sendMessage({
    message: "How many paws are in my house?",
  });
  console.log("Chat response 2:", response2.text);
  // [END chat]

  return { response1: response1.text, response2: response2.text };
}

export async function chatStreaming() {
  // [START chat_streaming]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const chat = ai.chats.create({
    model: "gemini-2.0-flash",
    history: [
      {
        role: "user",
        parts: [{ text: "Hello" }],
      },
      {
        role: "model",
        parts: [{ text: "Great to meet you. What would you like to know?" }],
      },
    ],
  });

  console.log("Streaming response for first message:");
  const stream1 = await chat.sendMessageStream({
    message: "I have 2 dogs in my house.",
  });
  for await (const chunk of stream1) {
    console.log(chunk.text);
    console.log("_".repeat(80));
  }

  console.log("Streaming response for second message:");
  const stream2 = await chat.sendMessageStream({
    message: "How many paws are in my house?",
  });
  for await (const chunk of stream2) {
    console.log(chunk.text);
    console.log("_".repeat(80));
  }

  console.log(chat.getHistory());
  // [END chat_streaming]
  return true;
}

export async function chatStreamingWithImages() {
  // [START chat_streaming_with_images]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const chat = ai.chats.create({ model: "gemini-2.0-flash" });

  console.log("Streaming response for initial text message:");
  const stream1 = await chat.sendMessageStream({
    message:
      "Hello, I'm interested in learning about musical instruments. Can I show you one?",
  });
  for await (const chunk of stream1) {
    console.log(chunk.text);
    console.log("_".repeat(80));
  }

  // Upload an image file (organ.jpg from the media folder)
  const imagePath = path.join(media, "organ.jpg");
  const image = await ai.files.upload({
    file: imagePath,
    config: { mimeType: "image/jpeg" },
  });
  console.log("Uploaded image file name:", image.name);

  console.log("Streaming response for message with image:");
  const stream2 = await chat.sendMessageStream({
    message: createUserContent([
      "What family of instruments does this instrument belong to?",
      createPartFromUri(image.uri, image.mimeType),
    ]),
  });
  for await (const chunk of stream2) {
    console.log(chunk.text);
    console.log("_".repeat(80));
  }
  // [END chat_streaming_with_images]
  return true;
}
