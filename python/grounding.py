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

MODEL_ID = "gemini-2.5-flash"


class GroundingUnitTests(absltest.TestCase):

    def test_grounding_maps(self):
        # [START grounding_maps]
        """Generates text using Google Maps as a grounding tool."""
        from google import genai
        from google.genai.types import GenerateContentConfig, GoogleMaps, LatLng, RetrievalConfig, Tool, ToolConfig

        client = genai.Client()

        maps_tool = Tool(
            google_maps=GoogleMaps()
        )

        location_context = RetrievalConfig(
            lat_lng=LatLng(latitude=34.050481, longitude=-118.248526)
        )

        prompt = "What are the best Italian restaurants within a 15-minute walk from here?"
        
        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt,
            config=GenerateContentConfig(
                tools=[maps_tool],
                tool_config=ToolConfig(retrieval_config=location_context),
            )
        )

        # Display the text output.
        print(response.text)

        # Display the grounding sources.
        if grounding := response.candidates[0].grounding_metadata:
          if grounding.grounding_chunks:
            print('-' * 40)
            print("Sources:")
            for chunk in grounding.grounding_chunks:
              print(f'- [{chunk.maps.title}]({chunk.maps.uri})')

          if widget_token := grounding.google_maps_widget_context_token:
            print('-' * 40)
            print(f'Maps token: {widget_token[:14]}...')
        # [END grounding_maps]

if __name__ == "__main__":
    absltest.main()

