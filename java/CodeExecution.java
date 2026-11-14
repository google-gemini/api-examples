
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

import com.google.genai.Client;
import com.google.genai.types.*;
import com.google.common.collect.ImmutableList;
import java.io.IOException;
import org.apache.http.HttpException;
import java.util.List;

public class CodeExecution {
    public static String codeExecutionBasic(Client client) throws IOException, HttpException {
        GenerateContentResponse response = client.models.generateContent("gemini-2.0-flash-001",
                "Write and execute code that calculates the sum of the first 50 prime numbers.Ensure that only the executable code and its resulting output are generated",
                null);

        System.out.println("Basic Code Execution Response:");

        // [start code_execution response]
        // Basic Code Execution Response:
        // ```python
        // def is_prime(n):
        // """Returns True if n is a prime number, False otherwise."""
        // if n <= 1:
        // return False
        // for i in range(2, int(n**0.5) + 1):
        // if n % i == 0:
        // return False
        // return True

        // def sum_of_first_n_primes(n):
        // """Calculates the sum of the first n prime numbers."""
        // primes = []
        // num = 2
        // while len(primes) < n:
        // if is_prime(num):
        // primes.append(num)
        // num += 1
        // return sum(primes)

        // # Calculate the sum of the first 50 prime numbers.
        // sum_of_primes = sum_of_first_n_primes(50)
        // print(sum_of_primes)
        // ```

        // ```
        // 2487
        // ```

        // [end code_execution response]

        String textResponse = response.text();
        System.out.println(textResponse);
        return textResponse;
    }

    public static String[] codeExecutionRequestOverride(Client client) throws IOException, HttpException {
        GenerateContentResponse response = client.models.generateContent(
                "gemini-2.0-flash-001",
                "Write and execute code that calculates the sum of the first 50 prime numbers. Ensure that only the executable code and its resulting output are generated.",
                GenerateContentConfig.builder()
                        .tools(ImmutableList.of(
                                Tool.builder()
                                        .codeExecution(ToolCodeExecution.builder().build())
                                        .build()))
                        .build());

        String executableCode = "";
        String executionResult = "";

        List<Part> parts = response.parts();
        for (Part part : parts) {
            if (part.executableCode().isPresent()) {
                executableCode = part.executableCode().get().code().get();
                System.out.println("\nExecutable Code:");
                System.out.println(executableCode);
            }
            if (part.codeExecutionResult().isPresent()) {
                executionResult = part.codeExecutionResult().get().output().get();
                System.out.println("\nExecution Result:");
                System.out.println(executionResult);
            }
        }

        // [start code_execution_request_override response]
        // Executable Code:
        // def is_prime(n):
        // if n <= 1:
        // return False
        // if n <= 3:
        // return True
        // if n % 2 == 0 or n % 3 == 0:
        // return False
        // i = 5
        // while i * i <= n:
        // if n % i == 0 or n % (i + 2) == 0:
        // return False
        // i += 6
        // return True

        // count = 0
        // num = 2
        // sum_of_primes = 0
        // while count < 50:
        // if is_prime(num):
        // sum_of_primes += num
        // count += 1
        // num += 1

        // print(sum_of_primes)

        // Execution Result:
        // 5117

        // Executable Code:
        // def is_prime(n):
        // if n <= 1:
        // return False
        // if n <= 3:
        // return True
        // if n % 2 == 0 or n % 3 == 0:
        // return False
        // i = 5
        // while i * i <= n:
        // if n % i == 0 or n % (i + 2) == 0:
        // return False
        // i += 6
        // return True

        // primes = []
        // num = 2
        // while len(primes) < 50:
        // if is_prime(num):
        // primes.append(num)
        // num += 1

        // total = sum(primes)
        // print(total)

        // Execution Result:
        // 5117

        // [end code_execution_request_override response]

        return new String[] { executableCode, executionResult };
    }

    public static void main(String[] args) throws IOException, HttpException {
        String apiKey = System.getenv("GOOGLE_API_KEY");
        if (apiKey == null || apiKey.isEmpty()) {
            System.err.println("Error: GOOGLE_API_KEY environment variable not set");
            System.exit(1);
        }

        Client client = new Client();
        codeExecutionBasic(client);
        System.out.println("\n" + "=".repeat(80) + "\n");
        codeExecutionRequestOverride(client);
    }
}
