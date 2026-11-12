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

    def video_gen_text_to_video_prompt(self):
        # [START video_gen_text_to_video_prompt]
        from google import genai
        from google.genai import types
        import time


        client = genai.Client()

        operation = client.models.generate_videos(
            model="veo-2.0-generate-001",
            prompt="a video of a unicorn flying through rainbows",
            config=types.GenerateVideosConfig(
                # At the moment the config must not be empty
                person_generation = "allow_adult", # allow_adult or dont_allow
                aspect_ratio = "9:16",  # 16:9 or 9:16
                number_of_videos = 1, # 1-2
                negative_prompt = "dragons",
                durationSeconds = 8, # 5-8
            ),
        )
      
        # Waiting for the videos to be generated
        while not operation.done:
            time.sleep(20)
            operation = client.operations.get(operation)
            print(operation)
      
        print(operation.result.generated_videos)
      
        for n, generated_video in enumerate(operation.result.generated_videos):
            client.files.download(file=generated_video.video)
            generated_video.video.save(f'video{n}.mp4') # Saves the video
        
        # [END video_gen_text_to_video_prompt]

    def video_gen_image_to_video_prompt(self):
        # [START video_gen_image_to_video_prompt]
        from google import genai
        from google.genai import types
        import PIL.Image
        import time
        import io

        client = genai.Client()
        organ = PIL.Image.open(media / "Cajun_instruments.jpg")
        
        # converting the image to bytes
        image_bytes_io = io.BytesIO()
        organ.save(image_bytes_io, format=organ.format)
        image_bytes = image_bytes_io.getvalue()

        operation = client.models.generate_videos(
            model="veo-2.0-generate-001",
            prompt="The instruments come to life and start dancing",
            image=types.Image(image_bytes=image_bytes, mime_type=organ.format),
            config=types.GenerateVideosConfig(
                # At the moment the config must not be empty
                person_generation = "dont_allow", # allow_adult or dont_allow
                aspect_ratio = "16:9",  # 16:9 or 9:16
                number_of_videos = 1, # 1-2
                negative_prompt = "hands",
                durationSeconds = 8, # 5-8
            ),
        )
      
        # Waiting for the videos to be generated
        while not operation.done:
            time.sleep(20)
            operation = client.operations.get(operation)
            print(operation)
      
        print(operation.result.generated_videos)
      
        for n, generated_video in enumerate(operation.result.generated_videos):
            client.files.download(file=generated_video.video)
            generated_video.video.save(f'video{n}.mp4') # Saves the video
        
        # [END video_gen_image_to_video_prompt]

if __name__ == "__main__":
    absltest.main()
