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

    def test_configure_model_parameters(self):
        # [START configure_model_parameters]
        from google import genai
        from google.genai import types

        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.5-flash",
            contents="Tell me a story about a magic backpack.",
            config=types.GenerateContentConfig(
                candidate_count=1,
                stop_sequences=["x"],
                max_output_tokens=20,
                temperature=1.0,
            ),
        )
        print(response.text)
        # [END configure_model_parameters]


if __name__ == "__main__":
    absltest.main()
