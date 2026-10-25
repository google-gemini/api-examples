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

    def test_image_editing_basic(self):
        # [START image_editing_basic]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO
        import pathlib

        client = genai.Client()

        # Define paths
        test_dir = pathlib.Path(__file__).parent
        input_image_path = test_dir / "test_image.jpg"  #  provide The image

        # Check if test image exists
        if not input_image_path.exists():
            print(f"Test image not found at {input_image_path}")
            print("Creating a simple test image instead")
            # Create a simple test image
            image = Image.new("RGB", (300, 200), color="gray")
            image.save(input_image_path)

        # Open the image
        image = Image.open(input_image_path)

        # Prepare input content
        text_input = "Hi, This is a picture of a cyberpunk character. Can you add a dragon above the character?"

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=[text_input, image],
            config=types.GenerateContentConfig(response_modalities=["Text", "Image"]),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                edited_image = Image.open(BytesIO(part.inline_data.data))
                edited_image.save("edited-image-llama.png")
                print("Edited image saved as 'edited-image-llama.png'")
        # [END image_editing_basic]

    def test_image_editing_with_style(self):
        # [START image_editing_with_style]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO
        import pathlib

        client = genai.Client()

        # Define paths
        test_dir = pathlib.Path(__file__).parent
        input_image_path = test_dir / "test_image.jpg"  #  provide The image

        # Check if test image exists
        if not input_image_path.exists():
            print(f"Test image not found at {input_image_path}")
            return

        # Open the image
        image = Image.open(input_image_path)

        # Prepare input content with style directions
        text_input = "Transform this image into a watercolor painting style"

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=[text_input, image],
            config=types.GenerateContentConfig(response_modalities=["Text", "Image"]),
        )

        # Process the response
        for part in response.candidates[0].content.parts:
            if part.text is not None:
                print(part.text)
            elif part.inline_data is not None:
                edited_image = Image.open(BytesIO(part.inline_data.data))
                edited_image.save("edited-image-watercolor.png")
                print("Edited image saved as 'edited-image-watercolor.png'")
        # [END image_editing_with_style]

    def test_image_editing_with_parameters(self):
        # [START image_editing_with_parameters]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO
        import pathlib

        client = genai.Client()

        # Define paths
        test_dir = pathlib.Path(__file__).parent
        input_image_path = test_dir / "test_image.jpg"

        # Check if test image exists
        if not input_image_path.exists():
            print(f"Test image not found at {input_image_path}")
            return

        # Open the image
        image = Image.open(input_image_path)

        # Prepare input content
        text_input = "This is a picture. Can you add a dragon in the sky?"

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=[text_input, image],
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
                edited_image = Image.open(BytesIO(part.inline_data.data))
                edited_image.save("edited-image-dragon.png")
                print("Edited image saved as 'edited-image-dragon.png'")
        # [END image_editing_with_parameters]

    def test_image_editing_with_save_to_path(self):
        # [START image_editing_with_save_to_path]
        from google import genai
        from google.genai import types
        from PIL import Image
        from io import BytesIO
        import pathlib

        client = genai.Client()

        # Define paths
        test_dir = pathlib.Path(__file__).parent
        input_image_path = test_dir / "test_image.jpg"

        # Check if test image exists
        if not input_image_path.exists():
            print(f"Test image not found at {input_image_path}")
            return

        # Open the image
        image = Image.open(input_image_path)

        # Prepare input content
        text_input = "Can you remove the background and replace it with a beach scene?"

        response = client.models.generate_content(
            model="gemini-2.0-flash-exp-image-generation",
            contents=[text_input, image],
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
                output_path = output_dir / "edited-beach-background.png"
                edited_image = Image.open(BytesIO(part.inline_data.data))
                edited_image.save(output_path)
                print(f"Edited image saved to {output_path}")
        # [END image_editing_with_save_to_path]


if __name__ == "__main__":
    absltest.main()
