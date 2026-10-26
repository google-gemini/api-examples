# -*- coding: utf-8 -*-
# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


"""
Demonstrates text and audio generation capabilities.

Setup:
- Requires GEMINI_API_KEY environment variable
- Generates text responses and audio files

Key features:
- Text generation with streaming output
- Audio generation saved as WAV file
"""

import asyncio
import wave
from google import genai
from google.genai import types
from absl.testing import absltest
import os
import pathlib


class UnitTests(absltest.TestCase):
    def setUp(self):
        self.client = genai.Client(
            api_key=os.getenv("GEMINI_API_KEY"), http_options={"api_version": "v1alpha"}
        )

    async def generate_text_live(self):
        try:
            full_response = ""

            async with self.client.aio.live.connect(
                model="gemini-2.0-flash-exp",
                config=types.LiveConnectConfig(
                    response_modalities=[types.Modality.TEXT]
                ),
            ) as session:
                # Send initial message
                await session.send(
                    input="Tell me a short story about vibe coding", end_of_turn=True
                )

                # Process responses
                async for response in session.receive():
                    if response.text:
                        full_response += response.text
                        print("Streaming:", response.text)

                    if (
                        response.server_content
                        and response.server_content.turn_complete
                    ):
                        break

            print("\nFull response:", full_response)
            return full_response

        except Exception as e:
            print("Error:", str(e))
            raise

    async def generate_audio_live(self):
        try:
            config = types.LiveConnectConfig(response_modalities=[types.Modality.AUDIO])

            async with self.client.aio.live.connect(
                model="gemini-2.0-flash-exp", config=config
            ) as session:

                # Create WAV file
                wf = wave.open("response.wav", "wb")
                wf.setnchannels(1)  # Mono
                wf.setsampwidth(2)  # 2 bytes per sample (16-bit)
                wf.setframerate(24000)  # 24kHz sample rate

                # Send text message to get audio response
                message = (
                    "Hello! Can you tell me a short story about a peaceful garden?"
                )
                await session.send(input=message, end_of_turn=True)

                # Process audio responses
                async for response in session.receive():
                    if response.data is not None:
                        wf.writeframes(response.data)
                        print("Received audio data chunk")

                    if (
                        response.server_content
                        and response.server_content.turn_complete
                    ):
                        break

                wf.close()
                print("Audio response saved to response.wav")

        except Exception as e:
            print("Error:", str(e))
            raise

    def test_text_generation(self):
        # [START text_generation]
        asyncio.run(self.generate_text_live())
        # [END text_generation]

    def test_audio_generation(self):
        # [START audio_generation]
        asyncio.run(self.generate_audio_live())
        # [END audio_generation]


if __name__ == "__main__":
    absltest.main()
