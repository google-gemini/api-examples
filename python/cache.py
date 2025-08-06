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

    def test_cache_create(self):
        # [START cache_create]
        from google import genai
        from google.genai import types

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config=types.CreateCachedContentConfig(
                contents=[document],
                system_instruction="You are an expert analyzing transcripts.",
            ),
        )
        print(cache)

        response = client.models.generate_content(
            model=model_name,
            contents="Please summarize this transcript",
            config=types.GenerateContentConfig(cached_content=cache.name),
        )
        print(response.text)
        # [END cache_create]
        client.caches.delete(name=cache.name)

    def test_cache_create_from_name(self):
        # [START cache_create_from_name]
        from google import genai
        from google.genai import types

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config=types.CreateCachedContentConfig(
                contents=[document],
                system_instruction="You are an expert analyzing transcripts.",
            ),
        )
        cache_name = cache.name  # Save the name for later

        # Later retrieve the cache
        cache = client.caches.get(name=cache_name)
        response = client.models.generate_content(
            model=model_name,
            contents="Find a lighthearted moment from this transcript",
            config=types.GenerateContentConfig(cached_content=cache.name),
        )
        print(response.text)
        # [END cache_create_from_name]
        client.caches.delete(name=cache.name)

    def test_cache_create_from_chat(self):
        # [START cache_create_from_chat]
        from google import genai
        from google.genai import types

        client = genai.Client()
        model_name = "gemini-2.5-flash"
        system_instruction = "You are an expert analyzing transcripts."

        # Create a chat session with the given system instruction.
        chat = client.chats.create(
            model=model_name,
            config=types.GenerateContentConfig(system_instruction=system_instruction),
        )
        document = client.files.upload(file=media / "a11.txt")

        response = chat.send_message(
            message=["Hi, could you summarize this transcript?", document]
        )
        print("\n\nmodel:  ", response.text)
        response = chat.send_message(
            message=["Okay, could you tell me more about the trans-lunar injection"]
        )
        print("\n\nmodel:  ", response.text)

        # To cache the conversation so far, pass the chat history as the list of contents.
        cache = client.caches.create(
            model=model_name,
            config={
                "contents": chat.get_history(),
                "system_instruction": system_instruction,
            },
        )
        # Continue the conversation using the cached content.
        chat = client.chats.create(
            model=model_name,
            config=types.GenerateContentConfig(cached_content=cache.name),
        )
        response = chat.send_message(
            message="I didn't understand that last part, could you explain it in simpler language?"
        )
        print("\n\nmodel:  ", response.text)
        # [END cache_create_from_chat]
        client.caches.delete(name=cache.name)

    def test_cache_delete(self):
        # [START cache_delete]
        from google import genai

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config={
                "contents": [document],
                "system_instruction": "You are an expert analyzing transcripts.",
            },
        )
        client.caches.delete(name=cache.name)
        # [END cache_delete]

    def test_cache_get(self):
        # [START cache_get]
        from google import genai

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config={
                "contents": [document],
                "system_instruction": "You are an expert analyzing transcripts.",
            },
        )
        print(client.caches.get(name=cache.name))
        # [END cache_get]
        client.caches.delete(name=cache.name)

    def test_cache_list(self):
        # [START cache_list]
        from google import genai

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config={
                "contents": [document],
                "system_instruction": "You are an expert analyzing transcripts.",
            },
        )
        print("My caches:")
        for c in client.caches.list():
            print("    ", c.name)
        # [END cache_list]
        client.caches.delete(name=cache.name)

    def test_cache_update(self):
        # [START cache_update]
        from google import genai
        from google.genai import types
        import datetime

        client = genai.Client()
        document = client.files.upload(file=media / "a11.txt")
        model_name = "gemini-2.5-flash"

        cache = client.caches.create(
            model=model_name,
            config={
                "contents": [document],
                "system_instruction": "You are an expert analyzing transcripts.",
            },
        )

        # Update the cache's time-to-live (ttl)
        ttl = f"{int(datetime.timedelta(hours=2).total_seconds())}s"
        client.caches.update(
            name=cache.name, config=types.UpdateCachedContentConfig(ttl=ttl)
        )
        print(f"After update:\n {cache}")

        # Alternatively, update the expire_time directly
        # Update the expire_time directly in valid RFC 3339 format (UTC with a "Z" suffix)
        expire_time = (
            (
                datetime.datetime.now(datetime.timezone.utc)
                + datetime.timedelta(minutes=15)
            )
            .isoformat()
            .replace("+00:00", "Z")
        )
        client.caches.update(
            name=cache.name,
            config=types.UpdateCachedContentConfig(expire_time=expire_time),
        )
        # [END cache_update]
        client.caches.delete(name=cache.name)


if __name__ == "__main__":
    absltest.main()
