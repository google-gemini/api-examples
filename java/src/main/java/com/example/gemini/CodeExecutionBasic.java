/*
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.example.gemini;

import com.google.genai.Client;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import org.apache.http.HttpException;

import java.io.IOException;

public class CodeExecutionBasic {
    public static void main(String[] args) throws IOException, HttpException {
        // [START code_execution_basic]
        Client client = new Client();

        String prompt = """
                Write and execute code that calculates the sum of the first 50 prime numbers.
                Ensure that only the executable code and its resulting output are generated.
                """;

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-pro-exp-02-05",
                        prompt,
                        null);

        for (Part part : response.candidates().get().getFirst().content().get().parts().get()) {
            System.out.println(part + "\n");
        }

        System.out.println(response.text());
        // [END code_execution_basic]

        /*
           [START code_execution_basic_return]
           Expected output:
           --------------------------------------------------------------------------------

           ```python
           import math

           def is_prime(n):
               """Checks if a number is prime."""
               if n < 2:
                   return False
               if n == 2:
                   return True
               if n % 2 == 0:
                   return False
               # Check only odd divisors up to the square root
               for i in range(3, int(math.sqrt(n)) + 1, 2):
                   if n % i == 0:
                       return False
               return True

           count = 0
           num = 2
           prime_sum = 0
           target_count = 50

           while count < target_count:
               if is_prime(num):
                   prime_sum += num
                   count += 1
               num += 1

           print(prime_sum)
           ```

           Output:
           ```
           5117
           ```
           [END code_execution_basic_return]
         */
    }
}
