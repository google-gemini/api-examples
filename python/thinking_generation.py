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

# Define the thinking model centrally
MODEL_ID = "gemini-2.5-pro-exp-03-25"


class ThinkingUnitTests(absltest.TestCase):

    def test_thinking_text_only_prompt(self):
        # [START thinking_text_only_prompt]
        """Generates text based on a simple reasoning prompt."""
        from google import genai

        client = genai.Client()
        prompt = "Explain the concept of Occam's Razor and provide a simple, everyday example."
        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt
        )

        print(response.text)
        # [END thinking_text_only_prompt]

    def test_thinking_text_only_prompt_streaming(self):
        # [START thinking_text_only_prompt_streaming]
        """Generates text based on a simple reasoning prompt using streaming."""
        from google import genai

        client = genai.Client()
        prompt = "Explain the concept of Occam's Razor and provide a simple, everyday example."
        response = client.models.generate_content_stream(
            model=MODEL_ID,
            contents=prompt
        )

        for chunk in response:
            print(chunk.text)
            print("_" * 80)
        # [END thinking_text_only_prompt_streaming]

    def test_thinking_logic_puzzle(self):
        # [START thinking_logic_puzzle]
        """Tests the model's ability to solve a classic logic puzzle."""
        from google import genai

        client = genai.Client()
        prompt = """
        Solve this logic puzzle and explain your reasoning step-by-step:
        There are three boxes. One contains only apples, one contains only oranges,
        and one contains both apples and oranges. The boxes have been incorrectly
        labeled such that no label is correct. You are allowed to draw one fruit
        from only one box of your choosing (without looking inside). Which box
        would you draw from to correctly label all boxes, and why?
        """
        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt
        )

        print(response.text)
        # [END thinking_logic_puzzle]

    def test_thinking_code_explanation(self):
        # [START thinking_code_explanation]
        """Tests the model's ability to understand and explain code."""
        from google import genai

        client = genai.Client()
        prompt = """
        Explain this Python code snippet step-by-step, including what it does
        and why recursion is used here:

        def factorial(n):
            '''Calculates the factorial of a non-negative integer.'''
            if not isinstance(n, int) or n < 0:
                raise ValueError("Input must be a non-negative integer")
            if n == 0:
                return 1
            else:
                return n * factorial(n-1)

        result = factorial(5)
        print(f"The factorial of 5 is: {result}")
        """
        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt
        )

        print(response.text)
        # [END thinking_code_explanation]

    def test_thinking_creative_writing_constraints(self):
        # [START thinking_creative_writing_constraints]
        """Tests creative writing with specific constraints."""
        from google import genai

        client = genai.Client()
        prompt = """
        Write a short story (max 150 words) about a detective investigating a
        mystery in a library, but the story must not contain the letter 'e'.
        """
        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt
        )

        print(response.text)
        # [END thinking_creative_writing_constraints]

    def test_thinking_with_search_tool(self):
        # [START thinking_with_search_tool]
        """Uses the search tool to answer a question requiring current information."""
        from google import genai
        from google.genai.types import Tool, GenerateContentConfig, GoogleSearch

        client = genai.Client()

        google_search_tool = Tool(
            google_search=GoogleSearch()
        )

        prompt = "What were the major scientific breakthroughs announced last week?"

        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt,
            config=GenerateContentConfig(
                tools=[google_search_tool],
                response_modalities=["TEXT"],
            )
        )

        print(response)
        # [END thinking_with_search_tool]

    def test_thinking_with_search_tool_streaming(self):
        # [START thinking_with_search_tool_streaming]
        """Uses the search tool with streaming output and retrieves grounding."""
        from google import genai
        from google.genai import types

        client = genai.Client()

        google_search_tool = types.Tool(
            google_search=types.GoogleSearch()
        )

        prompt = "When is the next total solar eclipse visible from mainland Europe?"

        response = client.models.generate_content_stream(
            model=MODEL_ID,
            contents=prompt,
            config=types.GenerateContentConfig(
                tools=[google_search_tool],
            )
        )

        for chunk in response:
            print(chunk.text)
            print("_" * 80)

        print(response)
        # [END thinking_with_search_tool_streaming]

    def test_thinking_code_execution(self):
        # [START thinking_code_execution]
        """Tests the model's ability to generate and execute code."""
        from google import genai
        from google.genai import types

        client = genai.Client()
        prompt = """
        What is the sum of the first 50 prime numbers?
        Generate and run Python code for the calculation, and make sure you get all 50.
        Provide the final sum clearly.
        """
        code_execution_tool = types.Tool(
            code_execution=types.ToolCodeExecution()
        )

        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt,
            config=types.GenerateContentConfig(
                tools=[code_execution_tool],
            ),
        )

        print(response)
        # [END thinking_code_execution]


    def test_thinking_structured_output_json(self):
        # [START thinking_structured_output_json]
        """Tests the model's ability to generate structured JSON output."""
        from google import genai

        client = genai.Client()

        prompt = """"
        Provide a list of 3 famous physicists and their key contributions in JSON format.
        Use this JSON schema:
            Physicist = {'name': str, 'contribution': str, 'era': str}
            Return: list[Physicist]
        """

        response = client.models.generate_content(
            model=MODEL_ID,
            contents=prompt
            # For stricter JSON mode (if structured output is supported):
            # config=types.GenerateContentConfig(
            #     response_mime_type="application/json",
            #     # json_schema=... # Define schema object here if needed
            # )
        )

        print(response.text)
        # [END thinking_structured_output_json]

if __name__ == "__main__":
    absltest.main()
