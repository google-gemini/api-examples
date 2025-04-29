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

import pathlib
import os
from absl.testing import absltest

media = pathlib.Path(__file__).parents[1] / "third_party"
output_dir = pathlib.Path(__file__).parent / "generated_images"
if not output_dir.exists():
    output_dir.mkdir(parents=True)


class UnitTests(absltest.TestCase):

    def test_image_generation_basic(self):
        # [START image_generation_basic]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO

        client = genai.Client()

        contents = (
            "Create a 3D rendered image of a penguin wearing a top hat "
            "and monocle, standing in front of a futuristic cityscape."
        )

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=contents,
            config=types.GenerateContentConfig(response_modalities=["Text", "Image"]),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                image = Image.open(BytesIO(part.inline_data.data))
                image.save("generated-penguin.png")
                print("Image saved as 'generated-penguin.png'")
        # [END image_generation_basic]

    def test_image_generation_with_parameters(self):
        # [START image_generation_with_parameters]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO

        client = genai.Client()

        contents = (
            "Create a watercolor painting of a mountain landscape "
            "with a flowing river, tall pine trees, and a cabin."
        )

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=contents,
            config=types.GenerateContentConfig(
                response_modalities=["Text", "Image"],
                temperature=0.9,
                top_p=0.8,
            ),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                image = Image.open(BytesIO(part.inline_data.data))
                image.save("generated-landscape.png")
                print("Image saved as 'generated-landscape.png'")
        # [END image_generation_with_parameters]

    def test_image_generation_with_style_directions(self):
        # [START image_generation_with_style_directions]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO

        client = genai.Client()

        contents = (
            "Create a digital art image in the style of cyberpunk anime "
            "showing a futuristic city at night with neon lights and flying vehicles."
        )

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=contents,
            config=types.GenerateContentConfig(response_modalities=["Text", "Image"]),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                image = Image.open(BytesIO(part.inline_data.data))
                image.save("generated-cyberpunk.png")
                print("Image saved as 'generated-cyberpunk.png'")
        # [END image_generation_with_style_directions]

    def test_image_generation_with_save_to_path(self):
        # [START image_generation_with_save_to_path]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO
        import os
        import pathlib

        client = genai.Client()

        contents = (
            "Create a photorealistic image of a hummingbird "
            "hovering next to a bright red flower."
        )

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=contents,
            config=types.GenerateContentConfig(response_modalities=["Text", "Image"]),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                # Create output directory if it doesn't exist
                output_dir = pathlib.Path(__file__).parent / "generated_images"
                if not output_dir.exists():
                    output_dir.mkdir(parents=True)

                # Save the image to the specified path
                output_path = output_dir / "hummingbird.png"
                image = Image.open(BytesIO(part.inline_data.data))
                image.save(output_path)
                print(f"Image saved to {output_path}")
        # [END image_generation_with_save_to_path]


if __name__ == "__main__":
    absltest.main()
