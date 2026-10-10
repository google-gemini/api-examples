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
import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export async function imageGenerationBasic() {
    // [START image_generation_basic]
    // Make sure to include the following import:
    // import { GoogleGenAI } from "@google/genai";
    // import fs from "fs";
    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    const contents = "Create a 3D rendered image of a penguin wearing a top hat " +
        "and monocle, standing in front of a futuristic cityscape.";

    try {
        const response = await ai.models.generateContent({
            model: "gemini-2.0-flash-exp-image-generation",
            contents: contents,
            config: {
                responseModalities: ["Text", "Image"]
            },
        });

        // Process text responses
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                // Process and save image data
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                fs.writeFileSync('generated-image.png', buffer);
                console.log('Image saved as generated-image.png');
            }
        }
    } catch (error) {
        console.error("Error generating image:", error);
    }
    // [END image_generation_basic]
}

export async function imageGenerationWithParameters() {
    // [START image_generation_with_parameters]
    // Make sure to include the following import:
    // import { GoogleGenAI } from "@google/genai";
    // import fs from "fs";
    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    const contents = "Create a watercolor painting of a mountain landscape " +
        "with a flowing river, tall pine trees, and a cabin.";

    try {
        const response = await ai.models.generateContent({
            model: "gemini-2.0-flash-exp-image-generation",
            contents: contents,
            config: {
                responseModalities: ["Text", "Image"],
                // Additional parameters can be added as needed
                temperature: 0.9,
                topP: 0.8,
            },
        });

        // Process text responses
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                // Process and save image data
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                fs.writeFileSync('generated-landscape.png', buffer);
                console.log('Image saved as generated-landscape.png');
            }
        }
    } catch (error) {
        console.error("Error generating image:", error);
    }
    // [END image_generation_with_parameters]
}

export async function imageGenerationWithStyleDirections() {
    // [START image_generation_with_style_directions]
    // Make sure to include the following import:
    // import { GoogleGenAI } from "@google/genai";
    // import fs from "fs";
    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    const contents = "Create a digital art image in the style of cyberpunk anime " +
        "showing a futuristic city at night with neon lights and flying vehicles.";

    try {
        const response = await ai.models.generateContent({
            model: "gemini-2.0-flash-exp-image-generation",
            contents: contents,
            config: {
                responseModalities: ["Text", "Image"]
            },
        });

        // Process text responses
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                // Process and save image data
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');
                fs.writeFileSync('generated-cyberpunk.png', buffer);
                console.log('Image saved as generated-cyberpunk.png');
            }
        }
    } catch (error) {
        console.error("Error generating image:", error);
    }
    // [END image_generation_with_style_directions]
}

export async function imageGenerationWithSaveToPath() {
    // [START image_generation_with_save_to_path]
    // Make sure to include the following imports:
    // import { GoogleGenAI } from "@google/genai";
    // import fs from "fs";
    // import path from "path";
    const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

    const contents = "Create a photorealistic image of a hummingbird " +
        "hovering next to a bright red flower.";

    try {
        const response = await ai.models.generateContent({
            model: "gemini-2.0-flash-exp-image-generation",
            contents: contents,
            config: {
                responseModalities: ["Text", "Image"]
            },
        });

        // Process text responses
        for (const part of response.candidates[0].content.parts) {
            if (part.text) {
                console.log(part.text);
            } else if (part.inlineData) {
                // Process and save image data
                const imageData = part.inlineData.data;
                const buffer = Buffer.from(imageData, 'base64');

                // Create a directory if it doesn't exist
                const outputDir = path.join(__dirname, 'generated_images');
                if (!fs.existsSync(outputDir)) {
                    fs.mkdirSync(outputDir, { recursive: true });
                }

                const outputPath = path.join(outputDir, 'hummingbird.png');
                fs.writeFileSync(outputPath, buffer);
                console.log(`Image saved to ${outputPath}`);
            }
        }
    } catch (error) {
        console.error("Error generating image:", error);
    }
    // [END image_generation_with_save_to_path]
} 