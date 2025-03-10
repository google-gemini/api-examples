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
from absl.testing import absltest
import pathlib

media = pathlib.Path(__file__).parents[1] / "third_party"


class UnitTests(absltest.TestCase):

    def test_chat(self):
        # [START chat]
        from google import genai
        from google.genai import types

        client = genai.Client()
        # Pass initial history using the "history" argument
        chat = client.chats.create(
            model="gemini-2.0-flash",
            history=[
                types.Content(role="user", parts=[types.Part(text="Hello")]),
                types.Content(
                    role="model",
                    parts=[
                        types.Part(
                            text="Great to meet you. What would you like to know?"
                        )
                    ],
                ),
            ],
        )
        response = chat.send_message(message="I have 2 dogs in my house.")
        print(response.text)
        response = chat.send_message(message="How many paws are in my house?")
        print(response.text)
        # [END chat]

    def test_chat_streaming(self):
        # [START chat_streaming]
        from google import genai
        from google.genai import types

        client = genai.Client()
        chat = client.chats.create(
            model="gemini-2.0-flash",
            history=[
                types.Content(role="user", parts=[types.Part(text="Hello")]),
                types.Content(
                    role="model",
                    parts=[
                        types.Part(
                            text="Great to meet you. What would you like to know?"
                        )
                    ],
                ),
            ],
        )
        response = chat.send_message_stream(message="I have 2 dogs in my house.")
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        response = chat.send_message_stream(message="How many paws are in my house?")
        for chunk in response:
            print(chunk.text)
            print("_" * 80)

        print(chat.get_history())
        # [END chat_streaming]

    def test_chat_streaming_with_images(self):
        # [START chat_streaming_with_images]
        from google import genai

        client = genai.Client()
        chat = client.chats.create(model="gemini-2.0-flash")

        response = chat.send_message_stream(
            message="Hello, I'm interested in learning about musical instruments. Can I show you one?"
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)

        # Upload image file locally
        image_file = client.files.upload(file=media / "organ.jpg")
        response = chat.send_message_stream(
            message=[
                "What family of instruments does this instrument belong to?",
                image_file,
            ]
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END chat_streaming_with_images]


if __name__ == "__main__":
    absltest.main()
