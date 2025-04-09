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
import com.google.genai.types.GenerateContentConfig;
import com.google.genai.types.GenerateContentResponse;
import com.google.genai.types.Part;
import com.google.genai.types.Tool;
import com.google.genai.types.ToolCodeExecution;

import java.util.Collections;

@SuppressWarnings("resource")
public class CodeExecution {
    public static GenerateContentResponse codeExecutionBasic() throws Exception {
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

        System.out.println("-".repeat(80));
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
        return response;
    }

    public static GenerateContentResponse codeExecutionRequestOverride() throws Exception {
        // [START code_execution_request_override]
        Client client = new Client();

        String prompt = """
                What is the sum of the first 50 prime numbers?
                Generate and run code for the calculation, and make sure you get all 50.
                """;

        GenerateContentConfig config =
                GenerateContentConfig.builder()
                        .tools(Collections.singletonList(
                                        Tool.builder()
                                                .codeExecution(ToolCodeExecution.builder().build())
                                                .build()
                                )
                        ).build();

        GenerateContentResponse response =
                client.models.generateContent(
                        "gemini-2.0-flash",
                        prompt,
                        config);

        System.out.println("-".repeat(80));
        System.out.println(response.executableCode());
        System.out.println(response.codeExecutionResult());
        // [END code_execution_request_override]

        /*
           [START code_execution_request_override_return]
           Expected output:
           --------------------------------------------------------------------------------
           def is_prime(n):
                if n <= 1:
                    return False
                if n <= 3:
                    return True
                if n % 2 == 0 or n % 3 == 0:
                    return False
                i = 5
                while i * i <= n:
                    if n % i == 0 or n % (i + 2) == 0:
                        return False
                    i += 6
                return True

            primes = []
            num = 2
            while len(primes) < 50:
                if is_prime(num):
                    primes.append(num)
                num += 1

            print(f'{primes=}')
            print(f'{sum(primes)=}')

            primes=[2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229]
            sum(primes)=5117
           [END code_execution_request_override_return]
         */
        return response;
    }
}
