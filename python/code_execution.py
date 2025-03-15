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

    def test_code_execution_basic(self):
        # [START code_execution_basic]
        from google import genai
        from google.genai import types

        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.0-pro-exp-02-05",
            contents=(
                "Using Python, write and execute code that calculates the sum of the first 50 prime numbers. "
                "Ensure that only the executable code and its resulting output are generated."
            ),
        )
        # Each part may contain text, executable code, or an execution result.
        for part in response.candidates[0].content.parts:
            print(part, "\n")

        print("-" * 80)
        # The .text accessor concatenates the parts into a markdown-formatted text.
        print("\n", response.text)
        # [END code_execution_basic]

        # [START code_execution_basic_return]
        # Expected output:
        # video_metadata=None thought=None code_execution_result=None executable_code=None file_data=None function_call=None function_response=None inline_data=None text='```python\ndef is_prime(n):\n    if n <= 1:\n        return False\n    for i in range(2, int(n**0.5) + 1):\n        if n % i == 0:\n            return False\n    return True\n\nsum_of_primes = 0\ncount = 0\nnum = 2\nwhile count < 50:\n    if is_prime(num):\n        sum_of_primes += num\n        count += 1\n    num += 1\n\nprint(sum_of_primes)\n```\n\nOutput:\n\n```\n5117\n```\n' 

        # --------------------------------------------------------------------------------

        #  ```python
        # def is_prime(n):
        #     if n <= 1:
        #         return False
        #     for i in range(2, int(n**0.5) + 1):
        #         if n % i == 0:
        #             return False
        #     return True

        # sum_of_primes = 0
        # count = 0
        # num = 2
        # while count < 50:
        #     if is_prime(num):
        #         sum_of_primes += num
        #         count += 1
        #     num += 1

        # print(sum_of_primes)
        # ```

        # Output:

        # ```
        # 5117
        # ```
        # [END code_execution_basic_return]

    def test_code_execution_request_override(self):
        # [START code_execution_request_override]
        from google import genai
        from google.genai import types

        client = genai.Client()
        response = client.models.generate_content(
            model="gemini-2.0-flash",
            contents=(
                "What is the sum of the first 50 prime numbers? "
                "Generate and run code for the calculation, and make sure you get all 50."
            ),
            config=types.GenerateContentConfig(
                tools=[types.Tool(code_execution=types.ToolCodeExecution())],
            ),
        )
        print(response.executable_code)
        print(response.code_execution_result)
        # [END code_execution_request_override]

        # [START code_execution_request_override_return]
        # Expected output:
        # def is_prime(n):
        #   """Efficiently checks if a number is prime."""
        #   if n <= 1:
        #     return False
        #   if n <= 3:
        #     return True
        #   if n % 2 == 0 or n % 3 == 0:
        #     return False
        #   i = 5
        #   while i * i <= n:
        #     if n % i == 0 or n % (i + 2) == 0:
        #       return False
        #     i += 6
        #   return True


        # def get_first_n_primes(n):
        #   """Generates the first n prime numbers."""
        #   primes = []
        #   num = 2
        #   while len(primes) < n:
        #     if is_prime(num):
        #       primes.append(num)
        #     num += 1
        #   return primes


        # first_50_primes = get_first_n_primes(50)
        # sum_of_primes = sum(first_50_primes)

        # print(f"{first_50_primes=}")
        # print(f"{sum_of_primes=}")


        # first_50_primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
        # sum_of_primes=5117
        # [END code_execution_request_override_return]

    def test_code_execution_chat(self):
        # [START code_execution_chat]
        from google import genai
        from google.genai import types

        client = genai.Client()
        chat = client.chats.create(
            model="gemini-2.0-flash",
            config=types.GenerateContentConfig(
                tools=[types.Tool(code_execution=types.ToolCodeExecution())],
            ),
        )
        # First, a simple chat message.
        response = chat.send_message(message='Can you print "Hello world!"?')
        # Then, a code-execution request.
        response = chat.send_message(
            message=(
                "What is the sum of the first 50 prime numbers? "
                "Generate and run code for the calculation, and make sure you get all 50."
            )
        )
        print(response.executable_code)
        print(response.code_execution_result)
        # [END code_execution_chat]

        # [START code_execution_chat_return]
        # Expected output:
        # def is_prime(n):
        #   """Return True if n is prime, False otherwise."""
        #   if n <= 1:
        #     return False
        #   for i in range(2, int(n**0.5) + 1):
        #     if n % i == 0:
        #       return False
        #   return True

        # primes = []
        # num = 2
        # while len(primes) < 50:
        #   if is_prime(num):
        #     primes.append(num)
        #   num += 1

        # print(f'{primes=}')

        # sum_of_primes = sum(primes)
        # print(f'{sum_of_primes=}')

        # primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
        # sum_of_primes=5117
        # [END code_execution_chat_return]


if __name__ == "__main__":
    absltest.main()
