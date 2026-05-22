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

import pkg from '@google/genai';
const { GoogleGenAI } = pkg;
import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';

// Get current file directory in ES modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

// Set up the API key - replace with your actual API key
const API_KEY = process.env.GOOGLE_API_KEY || 'YOUR_API_KEY_HERE';

// Initialize the Generative AI client
const genAI = new GoogleGenAI(API_KEY);

// Input and output directories
const INPUT_DIR = path.join(__dirname, '../input');
const OUTPUT_DIR = path.join(__dirname, '../output');

// Helper function to validate PNG format
function validatePngFormat(imagePath) {
    if (!imagePath) {
        throw new Error('Image path is required');
    }
    const ext = path.extname(imagePath).toLowerCase();
    if (ext !== '.png') {
        throw new Error('Only PNG images are supported for this example. Please provide a PNG image.');
    }
    return imagePath;
}

// Helper function to ensure input image exists
async function ensureInputImage() {
    await fs.mkdir(INPUT_DIR, { recursive: true });

    const files = await fs.readdir(INPUT_DIR);
    const pngFiles = files.filter(file => {
        const ext = path.extname(file).toLowerCase();
        return ext === '.png';
    });

    if (pngFiles.length === 0) {
        console.error(`No PNG images found in ${INPUT_DIR}. Please add a PNG image before running this script.`);
        process.exit(1);
    }

    return path.join(INPUT_DIR, pngFiles[0]);
}

export async function imageEditingBasic() {
    // [START image_editing_basic]
    // Make sure to include the following import:
    // import { GoogleGenAI } from '@google/genai';

    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    // Ensure output directory exists
    await fs.mkdir(OUTPUT_DIR, { recursive: true });
    const imagePath = validatePngFormat(await ensureInputImage());
    try {
        // Get the first image from the input directory and validate format

        const imageBuffer = await fs.readFile(imagePath);

        // Convert the image to base64
        const base64Image = imageBuffer.toString('base64');
        const mimeType = 'image/png';

        // Prepare the content parts
        const textInput = "Add another cat near the cat in the image";

        const response = await ai.models.generateContent({
            model: 'gemini-2.0-flash-exp-image-generation',
            contents: [
                { text: textInput },
                {
                    inlineData: {
                        mimeType: mimeType,
                        data: base64Image
                    }
                }
            ],
            config: {
                responseModalities: ['Text', 'Image']
            },
        });

        // Process the response
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                const outputPath = path.join(OUTPUT_DIR, 'edited-image-dragon.png');
                await fs.writeFile(outputPath, buffer);
                console.log(`Edited image saved as ${outputPath}`);
            }
        }
    } catch (error) {
        console.error("Error processing image:", error);
    }

    return `Processed image: ${imagePath}`;
}

export async function imageEditingWithStyle() {
    // [START image_editing_with_style]
    // Make sure to include the following import:
    // import { GoogleGenAI } from '@google/genai';

    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    // Ensure output directory exists
    await fs.mkdir(OUTPUT_DIR, { recursive: true });
    const imagePath = validatePngFormat(await ensureInputImage());
    try {
        // Get the first image from the input directory and validate format

        const imageBuffer = await fs.readFile(imagePath);

        // Convert the image to base64
        const base64Image = imageBuffer.toString('base64');
        const mimeType = 'image/png';

        // Prepare the content parts with style directions
        const textInput = "Transform this image into a watercolor painting style";

        const response = await ai.models.generateContent({
            model: 'gemini-2.0-flash-exp-image-generation',
            contents: [
                { text: textInput },
                {
                    inlineData: {
                        mimeType: mimeType,
                        data: base64Image
                    }
                }
            ],
            config: {
                responseModalities: ['Text', 'Image']
            },
        });

        // Process the response
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                const outputPath = path.join(OUTPUT_DIR, 'edited-image-watercolor.png');
                await fs.writeFile(outputPath, buffer);
                console.log(`Edited image saved as ${outputPath}`);
            }
        }
    } catch (error) {
        console.error("Error processing image:", error);
    }

    return `Processed image: ${imagePath}`;
}

export async function imageEditingWithParameters() {
    // [START image_editing_with_parameters]
    // Make sure to include the following import:
    // import { GoogleGenAI } from '@google/genai';

    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    // Ensure output directory exists
    await fs.mkdir(OUTPUT_DIR, { recursive: true });
    const imagePath = validatePngFormat(await ensureInputImage());
    try {
        // Get the first image from the input directory and validate format

        const imageBuffer = await fs.readFile(imagePath);

        // Convert the image to base64
        const base64Image = imageBuffer.toString('base64');
        const mimeType = 'image/png';

        // Prepare the content parts
        const textInput = "Remove the background and replace it with a beach scene";

        const response = await ai.models.generateContent({
            model: 'gemini-2.0-flash-exp-image-generation',
            contents: [
                { text: textInput },
                {
                    inlineData: {
                        mimeType: mimeType,
                        data: base64Image
                    }
                }
            ],
            config: {
                responseModalities: ['Text', 'Image'],
                temperature: 0.9,
                topP: 0.8,
            },
        });

        // Process the response
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                const outputPath = path.join(OUTPUT_DIR, 'edited-image-beach.png');
                await fs.writeFile(outputPath, buffer);
                console.log(`Edited image saved as ${outputPath}`);
            }
        }
    } catch (error) {
        console.error("Error processing image:", error);
    }

    return `Processed image: ${imagePath}`;
}

// Run all examples function
async function runAll() {
    console.log("Running image editing examples:");

    try {
        console.log("\n1. Basic Image Editing:");
        await imageEditingBasic();

        console.log("\n2. Image Editing with Style:");
        await imageEditingWithStyle();

        console.log("\n3. Image Editing with Parameters:");
        await imageEditingWithParameters();

        console.log("\nAll examples completed successfully!");
    } catch (error) {
        console.error("Error running examples:", error);
    }
}

// If this file is run directly, execute all examples
if (process.argv[1] === fileURLToPath(import.meta.url)) {
    if (!process.env.GEMINI_API_KEY) {
        console.error("Please set the GEMINI_API_KEY environment variable");
        process.exit(1);
    }

    runAll().catch(console.error);
}