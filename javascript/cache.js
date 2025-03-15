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

const media = path.join(__dirname, "..", "third_party");

export async function cacheCreate() {
  // [START cache_create]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });
  console.log("Cache created:", cache);

  const response = await ai.models.generateContent({
    model: modelName,
    contents: "Please summarize this transcript",
    config: { cached_content: cache.name },
  });
  console.log("Response text:", response.text);
  // [END cache_create]

  await ai.caches.delete({ name: cache.name });
  return response.text;
}

export async function cacheCreateFromName() {
  // [START cache_create_from_name]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });
  const cacheName = cache.name; // Save the name for later

  // Later retrieve the cache
  const retrievedCache = await ai.caches.get({ name: cacheName });
  const response = await ai.models.generateContent({
    model: modelName,
    contents: "Find a lighthearted moment from this transcript",
    config: { cached_content: retrievedCache.name },
  });
  console.log("Response text:", response.text);
  // [END cache_create_from_name]

  await ai.caches.delete({ name: retrievedCache.name });
  return response.text;
}

export async function cacheCreateFromChat() {
  // [START cache_create_from_chat]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const modelName = "gemini-1.5-flash-001";
  const systemInstruction = "You are an expert analyzing transcripts.";

  // Create a chat session with the system instruction.
  const chat = ai.chats.create({
    model: modelName,
    config: { system_instruction: systemInstruction },
  });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);

  let response = await chat.sendMessage({
    message: createUserContent([
      "Hi, could you summarize this transcript?",
      createPartFromUri(document.uri, document.mimeType),
    ]),
  });
  console.log("\n\nmodel:", response.text);

  response = await chat.sendMessage({
    message: "Okay, could you tell me more about the trans-lunar injection",
  });
  console.log("\n\nmodel:", response.text);

  // To cache the conversation so far, pass the chat history as the list of contents.
  const chatHistory = chat.getHistory();
  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: chatHistory,
      system_instruction: systemInstruction,
    },
  });

  // Continue the conversation using the cached content.
  const chatWithCache = ai.chats.create({
    model: modelName,
    config: { cached_content: cache.name },
  });
  response = await chatWithCache.sendMessage({
    message:
      "I didn't understand that last part, could you explain it in simpler language?",
  });
  console.log("\n\nmodel:", response.text);
  // [END cache_create_from_chat]

  await ai.caches.delete({ name: cache.name });
  return response.text;
}

export async function cacheDelete() {
  // [START cache_delete]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });
  await ai.caches.delete({ name: cache.name });
  console.log("Cache deleted:", cache.name);
  // [END cache_delete]
}

export async function cacheGet() {
  // [START cache_get]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });
  const retrievedCache = await ai.caches.get({ name: cache.name });
  console.log("Retrieved Cache:", retrievedCache);
  // [END cache_get]

  await ai.caches.delete({ name: cache.name });
  return retrievedCache;
}

export async function cacheList() {
  // [START cache_list]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  // Create a cache for demonstration.
  const cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });

  console.log("My caches:");
  const pager = await ai.caches.list({ config: { pageSize: 10 } });
  let page = pager.page;
  while (true) {
    for (const c of page) {
      console.log("    ", c.name);
    }
    if (!pager.hasNextPage()) break;
    page = await pager.nextPage();
  }
  // [END cache_list]

  await ai.caches.delete({ name: cache.name });
}

export async function cacheUpdate() {
  // [START cache_update]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const filePath = path.join(media, "a11.txt");
  const document = await ai.files.upload({
    file: filePath,
    config: { mimeType: "text/plain" },
  });
  console.log("Uploaded file name:", document.name);
  const modelName = "gemini-1.5-flash-001";

  const contents = [
    createUserContent(createPartFromUri(document.uri, document.mimeType)),
  ];

  let cache = await ai.caches.create({
    model: modelName,
    config: {
      contents: contents,
      system_instruction: "You are an expert analyzing transcripts.",
    },
  });

  // Update the cache's time-to-live (ttl)
  const ttl = `${2 * 3600}s`; // 2 hours in seconds
  cache = await ai.caches.update({
    name: cache.name,
    config: { ttl },
  });
  console.log("After update (TTL):", cache);

  // Alternatively, update the expire_time directly (in RFC 3339 format with a "Z" suffix)
  const expireTime = new Date(Date.now() + 15 * 60000)
    .toISOString()
    .replace(/\.\d{3}Z$/, "Z");
  cache = await ai.caches.update({
    name: cache.name,
    config: { expire_time: expireTime },
  });
  console.log("After update (expire_time):", cache);
  // [END cache_update]

  await ai.caches.delete({ name: cache.name });
  return cache;
}
