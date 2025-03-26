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

import { GoogleGenAI, Modality } from "@google/genai";

export async function liveClientBasic() {
    // [START live_client_basic]
    // Make sure to include the following import:
    // import {GoogleGenAI,  Modality} from '@google/genai';

    const responseQueue = [];
    let session = null;

    try {
        // This should use an async queue.
        async function waitMessage() {
            let done = false;
            let message = undefined;
            while (!done) {
                message = responseQueue.shift();
                if (message) {
                    if (message.serverContent && message.serverContent.modelTurn) {
                        console.log('streaming response', message.serverContent.modelTurn.parts);
                    }
                    done = true;
                } else {
                    await new Promise((resolve) => setTimeout(resolve, 100));
                }
            }
            // [streaming response console.log statements start]
            //  streaming response[{ text: 'The' }]
            //  streaming response[
            //                 {
            //                     text: ' neon signs of Neo-Kyoto buzzed with a restless energy, mirroring the'
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: ' frantic pulse of the city itself. Kai, a young coder with perpetually tired eyes'
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: ", hunched over his console, fingers flying across the holographic interface. He wasn't coding in the traditional sense, not with algorithms and data structures. Kai"
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: " was a vibe coder, and his canvas was the city's emotional landscape.\n" +
            //                         '\n' +
            //                         'He was hired by the Neo-Kyoto Harmony Collective, a group striving'
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: " to mitigate the city's simmering tension. They believed, and Kai had come to believe, that emotions weren't just subjective feelings; they were waves, frequencies, that could be manipulated with the right code.\n" +
            //                         '\n' +
            //                         'Kai worked with "'
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: 'emotes," snippets of code that translated abstract feelings into tangible sensory experiences. He had programmed the scent of cherry blossoms into the air of a particularly stressed-out district, hoping to coax a sense of calm. He had also crafted a low'
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: ', comforting hum, an auditory lullaby, to counter the screeching traffic and blaring advertisements.\n' +
            //                         '\n' +
            //                         "Tonight, however, was different. The city was a boiling pot. A recent political scandal had inflamed passions, and Kai could feel the city's anxiety surging around him, a discordant symphony of fear and anger."
            //                 }
            //             ]
            //  streaming response[
            //                 {
            //                     text: ' He needed to craft something powerful, something that could cut through the negativity.\n' +
            //                         '\n' +
            //                         `He delved deep into the code, pulling up files he hadn't touched in months: the "Hope" library. He hesitated. Hope was a tricky emotion; too subtle and it would be ignored, too potent and it could back`
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: 'fire. He needed to be precise.\n' +
            //                         '\n' +
            //                         'He decided to build an emote based on the feeling of a shared sunrise, the quiet dawn that always follows the darkest night. He layered in the sounds of birdsong, the faint shimmer of sunlight, and the cool scent of dew-kissed grass. He amplified the feeling'
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: ' of anticipation, of a fresh start, not in a forceful way, but with a gentle, almost hesitant, touch.\n' +
            //                         '\n' +
            //                         "He compiled the code and uploaded it to the city's network. He held his breath, watching the holographic monitor that displayed the city’s emotional energy. For a moment, nothing seemed to"
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: ' change. The red spikes of anxiety continued to pierce the screen. Then, a faint blue shimmer began to spread, a soft ripple against the backdrop of turmoil.\n' +
            //                         '\n' +
            //                         "It wasn't instantaneous. It wasn't a magic spell that erased all the city's problems. But it was there, a subtle current of"
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: " quiet optimism weaving its way through the city's chaotic energy. People paused in their arguments, their faces softening as they looked up at the sky. The blare of advertisements seemed to fade into the background.\n" +
            //                         '\n' +
            //                         'Kai slumped back in his chair, exhausted but strangely exhilarated. He had managed to nudge the city towards'
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: " a glimmer of hope. He knew his work was far from over. The city's emotions were volatile, constantly shifting and changing. But he was ready. He was a vibe coder, and he had the city’s emotional pulse at his fingertips. He would continue to weave his subtle threads of positive energy, hoping"
            //                 }
            //             ]
            // streaming response[
            //                 {
            //                     text: ' that, little by little, he could help the city find its rhythm again. The code, after all, was just the beginning. The real magic, he knew, was in the feeling.\n'
            //                 }
            //             ]
            // [streaming response console.log statements end]

            return message;
        }

        async function handleTurn() {
            const turn = [];
            let done = false;
            while (!done) {
                const message = await waitMessage();
                turn.push(message);
                if (message.serverContent && message.serverContent.turnComplete) {
                    done = true;
                }
            }
            return turn;
        }

        // Initialize the client with API key
        const client = new GoogleGenAI({
            apiKey: process.env.GEMINI_API_KEY,
            httpOptions: {
                apiVersion: 'v1alpha',
            },
        });

        session = await client.live.connect({
            model: 'gemini-2.0-flash-exp',
            callbacks: {
                onopen: function () {
                    console.log('Live session opened successfully');
                },
                onmessage: function (message) {
                    responseQueue.push(message);
                },
                onerror: function (e) {
                    console.log('Error:', e.message);
                },
                onclose: function (e) {
                    console.log('Close:', e.reason);
                },
            },
            config: { responseModalities: [Modality.TEXT] },
        });

        // Send a simple text message
        console.log('-'.repeat(80));
        console.log('Sending: Tell me a short story about vibe coding');
        session.sendClientContent({ turns: 'Tell me a short story about vibe coding', turnComplete: true });

        // Wait for and process the response
        const turnMessages = await handleTurn();

        let fullResponse = '';
        turnMessages.forEach(msg => {
            if (msg.serverContent && msg.serverContent.modelTurn && msg.serverContent.modelTurn.parts) {
                msg.serverContent.modelTurn.parts.forEach(part => {
                    if (part.text) {
                        fullResponse += part.text;
                    }
                });
            }
        });
        console.log('fullResponse', fullResponse);
        // [full response console.log statements start]

        // fullResponse The neon signs of Neo - Kyoto buzzed with a restless energy, mirroring the frantic pulse of the city itself.Kai, a young coder with perpetually tired eyes, hunched over his console, fingers flying across the holographic interface.He wasn't coding in the traditional sense, not with algorithms and data structures. Kai was a vibe coder, and his canvas was the city's emotional landscape.

        // He was hired by the Neo - Kyoto Harmony Collective, a group striving to mitigate the city's simmering tension. They believed, and Kai had come to believe, that emotions weren't just subjective feelings; they were waves, frequencies, that could be manipulated with the right code.

        // Kai worked with "emotes," snippets of code that translated abstract feelings into tangible sensory experiences.He had programmed the scent of cherry blossoms into the air of a particularly stressed - out district, hoping to coax a sense of calm.He had also crafted a low, comforting hum, an auditory lullaby, to counter the screeching traffic and blaring advertisements.

        //             Tonight, however, was different.The city was a boiling pot.A recent political scandal had inflamed passions, and Kai could feel the city's anxiety surging around him, a discordant symphony of fear and anger. He needed to craft something powerful, something that could cut through the negativity.

        // He delved deep into the code, pulling up files he hadn't touched in months: the "Hope" library. He hesitated. Hope was a tricky emotion; too subtle and it would be ignored, too potent and it could backfire. He needed to be precise.

        // He decided to build an emote based on the feeling of a shared sunrise, the quiet dawn that always follows the darkest night.He layered in the sounds of birdsong, the faint shimmer of sunlight, and the cool scent of dew - kissed grass.He amplified the feeling of anticipation, of a fresh start, not in a forceful way, but with a gentle, almost hesitant, touch.

        // He compiled the code and uploaded it to the city's network. He held his breath, watching the holographic monitor that displayed the city’s emotional energy. For a moment, nothing seemed to change. The red spikes of anxiety continued to pierce the screen. Then, a faint blue shimmer began to spread, a soft ripple against the backdrop of turmoil.

        // It wasn't instantaneous. It wasn't a magic spell that erased all the city's problems. But it was there, a subtle current of quiet optimism weaving its way through the city's chaotic energy.People paused in their arguments, their faces softening as they looked up at the sky.The blare of advertisements seemed to fade into the background.

        // Kai slumped back in his chair, exhausted but strangely exhilarated.He had managed to nudge the city towards a glimmer of hope.He knew his work was far from over.The city's emotions were volatile, constantly shifting and changing. But he was ready. He was a vibe coder, and he had the city’s emotional pulse at his fingertips. He would continue to weave his subtle threads of positive energy, hoping that, little by little, he could help the city find its rhythm again. The code, after all, was just the beginning. The real magic, he knew, was in the feeling.

        // [full response console.log statements end]

        // Close the session
        session.close();

        console.log("Session closed successfully");
        return turnMessages;
    } catch (error) {
        console.error('Error in liveClientBasic:', error);
        // Ensure session is closed even if an error occurs
        if (session) {
            try {
                session.close();
            } catch (closeError) {
                console.error('Error closing session:', closeError);
            }
        }
        throw error;
    }
    // [END live_client_basic]
}

export async function liveClientWithImage() {
    // [START live_client_with_image]
    // Make sure to include the following import:
    // import {GoogleGenAI,  Modality} from '@google/genai';

    const responseQueue = [];
    let session = null;

    try {
        // This should use an async queue.
        async function waitMessage() {
            let done = false;
            let message = undefined;
            while (!done) {
                message = responseQueue.shift();
                if (message) {
                    if (message.serverContent && message.serverContent.modelTurn) {
                        console.log('streaming response', message.serverContent.modelTurn.parts);
                    }
                    done = true;
                } else {
                    await new Promise((resolve) => setTimeout(resolve, 100));
                }
            }
            return message;
        }

        async function handleTurn() {
            const turn = [];
            let done = false;
            let timeout = 0;
            const MAX_TIMEOUT = 100; // 10 seconds timeout (100 * 100ms)

            while (!done && timeout < MAX_TIMEOUT) {
                const message = await waitMessage();
                turn.push(message);
                if (message.serverContent && message.serverContent.turnComplete) {
                    done = true;
                }
                // Reset timeout on message receipt
                timeout = 0;
            }

            if (timeout >= MAX_TIMEOUT) {
                console.warn('Warning: Timed out waiting for complete turn');
            }

            return turn;
        }

        // Initialize the client with API key
        const client = new GoogleGenAI({
            apiKey: process.env.GEMINI_API_KEY,
            httpOptions: {
                apiVersion: 'v1alpha',
            },
        });

        session = await client.live.connect({
            model: 'gemini-2.0-flash-exp',
            callbacks: {
                onopen: function () {
                    console.debug('Live session opened');
                },
                onmessage: function (message) {
                    responseQueue.push(message);
                },
                onerror: function (e) {
                    console.error('Error in live session:', e.message);
                },
                onclose: function (e) {
                    console.debug('Close:', e.reason);
                },
            },
            config: { responseModalities: [Modality.TEXT] },
        });

        // Send a message with an inline image
        const turns = [
            'This image is just black, can you see it?',
            {
                inlineData: {
                    // 2x2 black PNG, base64 encoded.
                    data: 'iVBORw0KGgoAAAANSUhEUgAAAAIAAAACCAIAAAD91JpzAAAAC0lEQVR4nGNgQAYAAA4AAamRc7EAAAAASUVORK5CYII=',
                    mimeType: 'image/png',
                },
            },
        ];
        session.sendClientContent({ turns: turns, turnComplete: true });

        // Wait for and process the response
        const imageTurn = await handleTurn();
        console.log('Image turn response:', imageTurn);


        let fullResponse = '';
        imageTurn.forEach(msg => {
            if (msg.serverContent && msg.serverContent.modelTurn && msg.serverContent.modelTurn.parts) {
                msg.serverContent.modelTurn.parts.forEach(part => {
                    if (part.text) {
                        fullResponse += part.text;
                    }
                });
            }
        });
        console.log('fullResponse', fullResponse);

        // Close the session
        session.close();
        return imageTurn;
    } catch (error) {
        console.error('Error in liveClientWithImage:', error);
        // Ensure session is closed even if an error occurs
        if (session) {
            try {
                session.close();
            } catch (closeError) {
                console.error('Error closing session:', closeError);
            }
        }
        throw error;
    }
    // [END live_client_with_image]
}
