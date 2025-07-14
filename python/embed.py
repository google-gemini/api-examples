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


class UnitTests(absltest.TestCase):

    def test_embed_content(self):
        # [START embed_content]
        from google import genai
        from google.genai import types

        client = genai.Client()
        text = "Hello World!"
        result = client.models.embed_content(
            model="gemini-embedding-001",
            contents=text,
            config=types.EmbedContentConfig(output_dimensionality=10),
        )
        print(result.embeddings)
        # [END embed_content]

    def test_batch_embed_contents(self):
        # [START batch_embed_contents]
        from google import genai
        from google.genai import types

        client = genai.Client()
        texts = [
            "What is the meaning of life?",
            "How much wood would a woodchuck chuck?",
            "How does the brain work?",
        ]
        result = client.models.embed_content(
            model="gemini-embedding-001",
            contents=texts,
            config=types.EmbedContentConfig(output_dimensionality=10),
        )
        print(result.embeddings)
        # [END batch_embed_contents]


if __name__ == "__main__":
    absltest.main()
