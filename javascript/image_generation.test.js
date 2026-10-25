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
    imageGenerationBasic,
    imageGenerationWithParameters,
    imageGenerationWithStyleDirections,
    imageGenerationWithSaveToPath
} from './image_generation.js';

import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const outputDir = path.join(__dirname, 'generated_images');

// Clean up image files before and after tests
beforeAll(() => {
    cleanupImageFiles();
});

afterAll(() => {
    cleanupImageFiles();
});

function cleanupImageFiles() {
    // Delete generated image files
    const filesToDelete = [
        'generated-image.png',
        'generated-landscape.png',
        'generated-cyberpunk.png'
    ];

    filesToDelete.forEach(file => {
        if (fs.existsSync(file)) {
            fs.unlinkSync(file);
        }
    });

    // Clean up directory images if it exists
    if (fs.existsSync(outputDir)) {
        const dirFiles = fs.readdirSync(outputDir);
        dirFiles.forEach(file => {
            fs.unlinkSync(path.join(outputDir, file));
        });

        // Remove directory if empty
        if (fs.readdirSync(outputDir).length === 0) {
            fs.rmdirSync(outputDir);
        }
    }
}

describe('Image Generation Examples', () => {
    // Set a longer timeout for image generation
    jest.setTimeout(30000);

    test('Basic image generation', async () => {
        await imageGenerationBasic();
        expect(fs.existsSync('generated-image.png')).toBeTruthy();
    });

    test('Image generation with parameters', async () => {
        await imageGenerationWithParameters();
        expect(fs.existsSync('generated-landscape.png')).toBeTruthy();
    });

    test('Image generation with style directions', async () => {
        await imageGenerationWithStyleDirections();
        expect(fs.existsSync('generated-cyberpunk.png')).toBeTruthy();
    });

    test('Image generation with save to path', async () => {
        await imageGenerationWithSaveToPath();

        // Check if directory and file exist
        expect(fs.existsSync(outputDir)).toBeTruthy();
        expect(fs.existsSync(path.join(outputDir, 'hummingbird.png'))).toBeTruthy();
    });
});