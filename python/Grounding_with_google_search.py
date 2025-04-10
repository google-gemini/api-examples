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

    def test_grounding_with_google_search(self):
        # [START grounding_with_google_search]
        
        from google import genai
        from google.genai.types import Tool, GenerateContentConfig, GoogleSearch
        
        client = genai.Client()
        model_id = "gemini-2.0-flash"

        google_search_tool = Tool(
            google_search = GoogleSearch()
        )

        response = client.models.generate_content(
            model=model_id,
            contents="When is the next total solar eclipse in the United States?",
            config=GenerateContentConfig(
                tools=[google_search_tool],
                response_modalities=["TEXT"],
                )
            )

        for each in response.candidates[0].content.parts:
            print(each.text)
        # Example response:
        # The next total solar eclipse visible in the contiguous United States will be on ...
        from google import genai
        from google.genai.types import Tool, GenerateContentConfig, GoogleSearch
        client = genai.Client()
        model_id = "gemini-2.0-flash"

        google_search_tool = Tool(
            google_search = GoogleSearch()
        )

        response = client.models.generate_content(
        model=model_id,
        contents="When is the next total solar eclipse in the United States?",
        config=GenerateContentConfig(
            tools=[google_search_tool],
            response_modalities=["TEXT"],
            )
        )

        for each in response.candidates[0].content.parts:
            print(each.text)
        # Example response:
        # The next total solar eclipse visible in the contiguous United States will be on ...

        # To get grounding metadata as web content.
        print(response.candidates[0].grounding_metadata.search_entry_point.rendered_content)
        # [END grounding_with_google_search]

if __name__ == "__main__":
    absltest.main()


