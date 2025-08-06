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

    def test_json_controlled_generation(self):
        # [START json_controlled_generation]
        from google import genai
        from google.genai import types
        from typing_extensions import TypedDict

        class Recipe(TypedDict):
            recipe_name: str
            ingredients: list[str]

        client = genai.Client()
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents="List a few popular cookie recipes.",
            config=types.GenerateContentConfig(
                response_mime_type="application/json", response_schema=list[Recipe]
            ),
        )
        print(result)
        # [END json_controlled_generation]

    def test_json_no_schema(self):
        # [START json_no_schema]
        from google import genai

        client = genai.Client()
        prompt = (
            "List a few popular cookie recipes in JSON format.\n\n"
            "Use this JSON schema:\n\n"
            "Recipe = {'recipe_name': str, 'ingredients': list[str]}\n"
            "Return: list[Recipe]"
        )
        result = client.models.generate_content(
            model="gemini-2.5-flash", contents=prompt
        )
        print(result)
        # [END json_no_schema]

    def test_json_enum(self):
        # [START json_enum]
        from google import genai
        from google.genai import types
        import enum

        class Choice(enum.Enum):
            PERCUSSION = "Percussion"
            STRING = "String"
            WOODWIND = "Woodwind"
            BRASS = "Brass"
            KEYBOARD = "Keyboard"

        client = genai.Client()
        organ = client.files.upload(file=media / "organ.jpg")
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents=["What kind of instrument is this:", organ],
            config=types.GenerateContentConfig(
                response_mime_type="application/json", response_schema=Choice
            ),
        )
        print(result)  # Expected output: "Keyboard" (or another appropriate enum value)
        # [END json_enum]

    def test_enum_in_json(self):
        # [START enum_in_json]
        from google import genai
        from google.genai import types
        import enum
        from typing_extensions import TypedDict

        class Grade(enum.Enum):
            A_PLUS = "a+"
            A = "a"
            B = "b"
            C = "c"
            D = "d"
            F = "f"

        class Recipe(TypedDict):
            recipe_name: str
            grade: Grade

        client = genai.Client()
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents="List about 10 cookie recipes, grade them based on popularity",
            config=types.GenerateContentConfig(
                response_mime_type="application/json", response_schema=list[Recipe]
            ),
        )
        # Expected output: a JSON-parsed list with recipe names and grades (e.g., "a+")
        print(result)
        # [END enum_in_json]

    def test_json_enum_raw(self):
        # [START json_enum_raw]
        from google import genai
        from google.genai import types

        client = genai.Client()
        organ = client.files.upload(file=media / "organ.jpg")
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents=["What kind of instrument is this:", organ],
            config=types.GenerateContentConfig(
                response_mime_type="application/json",
                response_schema={
                    "type": "STRING",
                    "enum": ["Percussion", "String", "Woodwind", "Brass", "Keyboard"],
                },
            ),
        )
        print(result)  # Expected output: "Keyboard"
        # [END json_enum_raw]

    def test_x_enum(self):
        # [START x_enum]
        from google import genai
        from google.genai import types
        import enum

        class Choice(enum.Enum):
            PERCUSSION = "Percussion"
            STRING = "String"
            WOODWIND = "Woodwind"
            BRASS = "Brass"
            KEYBOARD = "Keyboard"

        client = genai.Client()
        organ = client.files.upload(file=media / "organ.jpg")
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents=["What kind of instrument is this:", organ],
            config=types.GenerateContentConfig(
                response_mime_type="text/x.enum", response_schema=Choice
            ),
        )
        print(result)  # Expected output: "Keyboard"
        # [END x_enum]

    def test_x_enum_raw(self):
        # [START x_enum_raw]
        from google import genai
        from google.genai import types

        client = genai.Client()
        organ = client.files.upload(file=media / "organ.jpg")
        result = client.models.generate_content(
            model="gemini-2.5-flash",
            contents=["What kind of instrument is this:", organ],
            config=types.GenerateContentConfig(
                response_mime_type="text/x.enum",
                response_schema={
                    "type": "STRING",
                    "enum": ["Percussion", "String", "Woodwind", "Brass", "Keyboard"],
                },
            ),
        )
        print(result)  # Expected output: "Keyboard"
        # [END x_enum_raw]


if __name__ == "__main__":
    absltest.main()
