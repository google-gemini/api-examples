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
from absl.testing import absltest

media = pathlib.Path(__file__).parents[1] / "third_party"


class UnitTests(absltest.TestCase):

    def test_text_gen_text_only_prompt(self):
        # [START text_gen_text_only_prompt]
        from google import genai

        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.0-flash", contents="Write a story about a magic backpack."
        )
        print(response.text)
        # [END text_gen_text_only_prompt]

    def test_text_gen_text_only_prompt_streaming(self):
        # [START text_gen_text_only_prompt_streaming]
        from google import genai

        client = genai.Client()
        response = client.models.generate_content_stream(
            model="gemini-2.0-flash", contents="Write a story about a magic backpack."
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_text_only_prompt_streaming]

    def test_text_gen_multimodal_one_image_prompt(self):
        # [START text_gen_multimodal_one_image_prompt]
        from google import genai
        import PIL.Image

        client = genai.Client()
        organ = PIL.Image.open(media / "organ.jpg")
        response = client.models.generate_content(
            model="gemini-2.0-flash", contents=["Tell me about this instrument", organ]
        )
        print(response.text)
        # [END text_gen_multimodal_one_image_prompt]

    def test_text_gen_multimodal_one_image_prompt_streaming(self):
        # [START text_gen_multimodal_one_image_prompt_streaming]
        from google import genai
        import PIL.Image

        client = genai.Client()
        organ = PIL.Image.open(media / "organ.jpg")
        response = client.models.generate_content_stream(
            model="gemini-2.0-flash", contents=["Tell me about this instrument", organ]
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_multimodal_one_image_prompt_streaming]

    def test_text_gen_multimodal_multi_image_prompt(self):
        # [START text_gen_multimodal_multi_image_prompt]
        from google import genai
        import PIL.Image

        client = genai.Client()
        organ = PIL.Image.open(media / "organ.jpg")
        cajun_instrument = PIL.Image.open(media / "Cajun_instruments.jpg")
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=[
                "What is the difference between both of these instruments?",
                organ,
                cajun_instrument,
            ],
        )
        print(response.text)
        # [END text_gen_multimodal_multi_image_prompt]

    def test_text_gen_multimodal_multi_image_prompt_streaming(self):
        # [START text_gen_multimodal_multi_image_prompt_streaming]
        from google import genai
        import PIL.Image

        client = genai.Client()
        organ = PIL.Image.open(media / "organ.jpg")
        cajun_instrument = PIL.Image.open(media / "Cajun_instruments.jpg")
        response = client.models.generate_content_stream(
            model="gemini-2.0-flash",
            contents=[
                "What is the difference between both of these instruments?",
                organ,
                cajun_instrument,
            ],
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_multimodal_multi_image_prompt_streaming]

    def test_text_gen_multimodal_audio(self):
        # [START text_gen_multimodal_audio]
        from google import genai

        client = genai.Client()
        sample_audio = client.files.upload(file=media / "sample.mp3")
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this audio file.", sample_audio],
        )
        print(response.text)
        # [END text_gen_multimodal_audio]

    def test_text_gen_multimodal_audio_streaming(self):
        # [START text_gen_multimodal_audio_streaming]
        from google import genai

        client = genai.Client()
        sample_audio = client.files.upload(file=media / "sample.mp3")
        response = client.models.generate_content_stream(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this audio file.", sample_audio],
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_multimodal_audio_streaming]

    def test_text_gen_multimodal_video_prompt(self):
        # [START text_gen_multimodal_video_prompt]
        from google import genai
        import time

        client = genai.Client()
        # Video clip (CC BY 3.0) from https://peach.blender.org/download/
        myfile = client.files.upload(file=media / "Big_Buck_Bunny.mp4")
        print(f"{myfile=}")

        # Poll until the video file is completely processed (state becomes ACTIVE).
        while not myfile.state or myfile.state.name != "ACTIVE":
            print("Processing video...")
            print("File state:", myfile.state)
            time.sleep(5)
            myfile = client.files.get(name=myfile.name)

        response = client.models.generate_content(
            model="gemini-2.0-flash", contents=[myfile, "Describe this video clip"]
        )
        print(f"{response.text=}")
        # [END text_gen_multimodal_video_prompt]

    def test_text_gen_multimodal_video_prompt_streaming(self):
        # [START text_gen_multimodal_video_prompt_streaming]
        from google import genai
        import time

        client = genai.Client()
        # Video clip (CC BY 3.0) from https://peach.blender.org/download/
        myfile = client.files.upload(file=media / "Big_Buck_Bunny.mp4")
        print(f"{myfile=}")

        # Poll until the video file is completely processed (state becomes ACTIVE).
        while not myfile.state or myfile.state.name != "ACTIVE":
            print("Processing video...")
            print("File state:", myfile.state)
            time.sleep(5)
            myfile = client.files.get(name=myfile.name)

        response = client.models.generate_content_stream(
            model="gemini-2.0-flash", contents=[myfile, "Describe this video clip"]
        )
        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_multimodal_video_prompt_streaming]

    def test_text_gen_multimodal_pdf(self):
        # [START text_gen_multimodal_pdf]
        from google import genai

        client = genai.Client()
        sample_pdf = client.files.upload(file=media / "test.pdf")
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this document:", sample_pdf],
        )
        print(f"{response.text=}")
        # [END text_gen_multimodal_pdf]

    def test_text_gen_multimodal_pdf_streaming(self):
        # [START text_gen_multimodal_pdf_streaming]
        from google import genai

        client = genai.Client()
        sample_pdf = client.files.upload(file=media / "test.pdf")
        response = client.models.generate_content_stream(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this document:", sample_pdf],
        )

        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END text_gen_multimodal_pdf_streaming]


if __name__ == "__main__":
    absltest.main()
