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

// Helper for sleeping (used in video polling)
const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

export async function filesCreateText() {
  // [START files_create_text]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const myfile = await ai.files.upload({
    file: path.join(media, "poem.txt"),
  });
  console.log("Uploaded file:", myfile);

  const result = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents: createUserContent([
      createPartFromUri(myfile.uri, myfile.mimeType),
      "\n\n",
      "Can you add a few more lines to this poem?",
    ]),
  });
  console.log("result.text=", result.text);
  // [END files_create_text]
  return result.text;
}

export async function filesCreateImage() {
  // [START files_create_image]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const myfile = await ai.files.upload({
    file: path.join(media, "Cajun_instruments.jpg"),
    config: { mimeType: "image/jpeg" },
  });
  console.log("Uploaded file:", myfile);

  const result = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents: createUserContent([
      createPartFromUri(myfile.uri, myfile.mimeType),
      "\n\n",
      "Can you tell me about the instruments in this photo?",
    ]),
  });
  console.log("result.text=", result.text);
  // [END files_create_image]
  return result.text;
}

export async function filesCreateAudio() {
  // [START files_create_audio]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const myfile = await ai.files.upload({
    file: path.join(media, "sample.mp3"),
    config: { mimeType: "audio/mpeg" },
  });
  console.log("Uploaded file:", myfile);

  const result = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents: createUserContent([
      createPartFromUri(myfile.uri, myfile.mimeType),
      "Describe this audio clip",
    ]),
  });
  console.log("result.text=", result.text);
  // [END files_create_audio]
  return result.text;
}

export async function filesCreateVideo() {
  // [START files_create_video]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  let myfile = await ai.files.upload({
    file: path.join(media, "Big_Buck_Bunny.mp4"),
    config: { mimeType: "video/mp4" },
  });
  console.log("Uploaded video file:", myfile);

  // Poll until the video file is completely processed (state becomes ACTIVE).
  while (!myfile.state || myfile.state.toString() !== "ACTIVE") {
    console.log("Processing video...");
    console.log("File state: ", myfile.state);
    await sleep(5000);
    myfile = await ai.files.get({ name: myfile.name });
  }

  const result = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents: createUserContent([
      createPartFromUri(myfile.uri, myfile.mimeType),
      "Describe this video clip",
    ]),
  });
  console.log("result.text=", result.text);
  // [END files_create_video]
  return result.text;
}

export async function filesCreatePdf() {
  // [START files_create_pdf]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const samplePdf = await ai.files.upload({
    file: path.join(media, "test.pdf"),
    config: { mimeType: "application/pdf" },
  });
  const response = await ai.models.generateContent({
    model: "gemini-2.0-flash",
    contents: createUserContent([
      "Give me a summary of this pdf file.",
      createPartFromUri(samplePdf.uri, samplePdf.mimeType),
    ]),
  });
  console.log("Result text:", response.text);
  // [END files_create_pdf]
  return response.text;
}

export async function filesList() {
  // [START files_list]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  console.log("My files:");
  // Using the pager style to list files
  const pager = await ai.files.list({ config: { pageSize: 10 } });
  let page = pager.page;
  const names = [];
  while (true) {
    for (const f of page) {
      console.log("  ", f.name);
      names.push(f.name);
    }
    if (!pager.hasNextPage()) break;
    page = await pager.nextPage();
  }
  // [END files_list]
  return names;
}

export async function filesGet() {
  // [START files_get]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });
  const myfile = await ai.files.upload({
    file: path.join(media, "poem.txt"),
  });
  const fileName = myfile.name;
  console.log(fileName);

  const fetchedFile = await ai.files.get({ name: fileName });
  console.log(fetchedFile);
  // [END files_get]
  return fetchedFile;
}
