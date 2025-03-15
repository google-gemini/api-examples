/**
 * @license
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import {GoogleGenAI} from '@google/genai';

export async function codeExecutionBasic() {
  // [START code_execution_basic]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-pro-exp-02-05',
    contents: `Write and execute code that calculates the sum of the first 50 prime numbers.
               Ensure that only the executable code and its resulting output are generated.`,
  });

  // Each part may contain text, executable code, or an execution result.
  for (const part of response.candidates[0].content.parts) {
    console.log(part);
    console.log('\n');
  }

  console.log('-'.repeat(80));
  // The `.text` accessor concatenates the parts into a markdown-formatted text.
  console.log('\n', response.text);
  // [END code_execution_basic]

  // [START code_execution_basic_return]
  // {
  //   text: '```python\n' +
  //     'def is_prime(n):\n' +
  //     '    if n <= 1:\n' +
  //     '        return False\n' +
  //     '    for i in range(2, int(n**0.5) + 1):\n' +
  //     '        if n % i == 0:\n' +
  //     '            return False\n' +
  //     '    return True\n' +
  //     '\n' +
  //     'sum_of_primes = 0\n' +
  //     'count = 0\n' +
  //     'num = 2\n' +
  //     'while count < 50:\n' +
  //     '    if is_prime(num):\n' +
  //     '        sum_of_primes += num\n' +
  //     '        count += 1\n' +
  //     '    num += 1\n' +
  //     '\n' +
  //     'print(sum_of_primes)\n' +
  //     '```\n' +
  //     '\n' +
  //     '```output\n' +
  //     '5117\n' +
  //     '```\n'
  // }
  
  
  // --------------------------------------------------------------------------------
  
  // ```python
  // def is_prime(n):
  //     if n <= 1:
  //         return False
  //     for i in range(2, int(n**0.5) + 1):
  //         if n % i == 0:
  //             return False
  //     return True
  
  // sum_of_primes = 0
  // count = 0
  // num = 2
  // while count < 50:
  //     if is_prime(num):
  //         sum_of_primes += num
  //         count += 1
  //     num += 1
  
  // print(sum_of_primes)
  // ```
  
  // ```output
  // 5117
  // ```
  // [END code_execution_basic_return]

  return {
    parts: response.candidates[0].content.parts,
    text: response.text,
  };
}

export async function codeExecutionRequestOverride() {
  // [START code_execution_request_override]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const response = await ai.models.generateContent({
    model: 'gemini-2.0-flash',
    contents:
      'What is the sum of the first 50 prime numbers? Generate and run code for the calculation, and make sure you get all 50.',
    config: {
      tools: [{codeExecution: {}}],
    },
  });

  console.log('\n', response.executableCode);
  console.log('\n', response.codeExecutionResult);
  // [END code_execution_request_override]

  // [START code_execution_request_override_return]
  // def is_prime(n):
  //   """Check if a number is prime."""
  //   if n <= 1:
  //       return False
  //   if n <= 3:
  //       return True
  //   if n % 2 == 0 or n % 3 == 0:
  //       return False
  //   i = 5
  //   while i * i <= n:
  //       if n % i == 0 or n % (i + 2) == 0:
  //           return False
  //       i += 6
  //   return True

  // primes = []
  // num = 2
  // while len(primes) < 50:
  //   if is_prime(num):
  //       primes.append(num)
  //   num += 1

  // print(f'{primes=}')
  // print(f'{sum(primes)=}')


  // primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
  // sum(primes)=5117
  // [END code_execution_request_override_return]

  return {
    executableCode: response.executableCode,
    codeExecutionResult: response.codeExecutionResult,
  };
}

export async function codeExecutionChat() {
  // [START code_execution_chat]
  // Make sure to include the following import:
  // import {GoogleGenAI} from '@google/genai';
  const ai = new GoogleGenAI({apiKey: process.env.GEMINI_API_KEY});

  const chat = ai.chats.create({
    model: 'gemini-2.0-flash'
  });

  const response = await chat.sendMessage({
    message: 'What is the sum of the first 50 prime numbers? Generate and run code for the calculation, and make sure you get all 50.',
    config: {
      tools: [{codeExecution: {}}],
    },
  });

  console.log('\n', response.executableCode);
  console.log('\n', response.codeExecutionResult);
  // [END code_execution_chat]

  // [START code_execution_chat_return]
  // def is_prime(n):
  //   """Check if a number is prime."""
  //   if n <= 1:
  //       return False
  //   if n <= 3:
  //       return True
  //   if n % 2 == 0 or n % 3 == 0:
  //       return False
  //   i = 5
  //   while i * i <= n:
  //       if n % i == 0 or n % (i + 2) == 0:
  //           return False
  //       i += 6
  //   return True

  // primes = []
  // num = 2
  // while len(primes) < 50:
  //   if is_prime(num):
  //       primes.append(num)
  //   num += 1

  // print(f'{primes=}')
  // print(f'{sum(primes)=}')


  // primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
  // sum(primes)=5117
  // [END code_execution_request_chat_return]

  return {
    executableCode: response.executableCode,
    codeExecutionResult: response.codeExecutionResult,
  };
}
