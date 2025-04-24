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

    def test_files_create_text(self):
        # [START files_create_text]
        from google import genai

        client = genai.Client()
        myfile = client.files.upload(file=media / "poem.txt")
        print(f"{myfile=}")

        result = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=[myfile, "\n\n", "Can you add a few more lines to this poem?"],
        )
        print(f"{result.text=}")
        # [END files_create_text]

    def test_files_create_image(self):
        # [START files_create_image]
        from google import genai

        client = genai.Client()
        myfile = client.files.upload(file=media / "Cajun_instruments.jpg")
        print(f"{myfile=}")

        result = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=[
                myfile,
                "\n\n",
                "Can you tell me about the instruments in this photo?",
            ],
        )
        print(f"{result.text=}")
        # [END files_create_image]

    def test_files_create_audio(self):
        # [START files_create_audio]
        from google import genai

        client = genai.Client()
        myfile = client.files.upload(file=media / "sample.mp3")
        print(f"{myfile=}")

        result = client.models.generate_content(
            model="gemini-2.0-flash", contents=[myfile, "Describe this audio clip"]
        )
        print(f"{result.text=}")
        # [END files_create_audio]

    def test_files_create_video(self):
        # [START files_create_video]
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

        result = client.models.generate_content(
            model="gemini-2.0-flash", contents=[myfile, "Describe this video clip"]
        )
        print(f"{result.text=}")
        # [END files_create_video]

    def test_files_create_pdf(self):
        # [START files_create_pdf]
        from google import genai

        client = genai.Client()
        sample_pdf = client.files.upload(file=media / "test.pdf")
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this pdf file.", sample_pdf],
        )
        print(response.text)
        # [END files_create_pdf]

    def test_files_create_from_IO(self):
        # [START files_create_io]
        from google import genai
        from google.genai import types

        client = genai.Client()
        fpath = media / "test.pdf"
        with open(fpath, "rb") as f:
            sample_pdf = client.files.upload(
                file=f, config=types.UploadFileConfig(mime_type="application/pdf")
            )
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=["Give me a summary of this pdf file.", sample_pdf],
        )
        print(response.text)
        # [END files_create_io]

    def test_files_list(self):
        # [START files_list]
        from google import genai

        client = genai.Client()
        print("My files:")
        for f in client.files.list():
            print("  ", f.name)
        # [END files_list]

    def test_files_get(self):
        # [START files_get]
        from google import genai

        client = genai.Client()
        myfile = client.files.upload(file=media / "poem.txt")
        file_name = myfile.name
        print(file_name)  # "files/*"

        myfile = client.files.get(name=file_name)
        print(myfile)
        # [END files_get]

    def test_files_delete(self):
        # [START files_delete]
        from google import genai

        client = genai.Client()
        myfile = client.files.upload(file=media / "poem.txt")

        client.files.delete(name=myfile.name)

        try:
            result = client.models.generate_content(
                model="gemini-2.0-flash", contents=[myfile, "Describe this file."]
            )
            print(result)
        except genai.errors.ClientError:
            pass
        # [END files_delete]


if __name__ == "__main__":
    absltest.main()
