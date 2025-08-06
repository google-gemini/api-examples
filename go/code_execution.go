package examples

import (
	"context"
	"fmt"
	"os"
	"log"

	"google.golang.org/genai"
)

func CodeExecutionBasic() (*genai.GenerateContentResponse, error) {
	// [START code_execution_basic]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-pro",
		genai.Text(
			`Write and execute code that calculates the sum of the first 50 prime numbers.
			 Ensure that only the executable code and its resulting output are generated.`,
		),
		&genai.GenerateContentConfig{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response.
	printResponse(response)

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(response.Text())
	// [END code_execution_basic]

	// [START code_execution_basic_return]
	// Expected output:

	// --------------------------------------------------------------------------------
	// ```python
	// def is_prime(n):
	// 	if n <= 1:
	// 		return False
	// 	for i in range(2, int(n**0.5) + 1):
	// 		if n % i == 0:
	// 			return False
	// 	return True

	// sum_of_primes = 0
	// count = 0
	// num = 2
	// while count < 50:
	// 	if is_prime(num):
	// 		sum_of_primes += num
	// 		count += 1
	// 	num += 1

	// print(sum_of_primes)
	// ```

	// ```output
	// 5117
	// ```
	// [END code_execution_basic_return]

	return response, err
}

func CodeExecutionRequestOverride() (*genai.GenerateContentResponse, error) {
	// [START code_execution_request_override]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-pro",
		genai.Text(
			`What is the sum of the first 50 prime numbers?
Generate and run code for the calculation, and make sure you get all 50.`,
		),
		&genai.GenerateContentConfig{
			Tools: []*genai.Tool{
				{CodeExecution: &genai.ToolCodeExecution{}},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response.
	printResponse(response)

	fmt.Println("--------------------------------------------------------------------------------")
	
	fmt.Println(response.ExecutableCode())
	fmt.Println(response.CodeExecutionResult())
	// [END code_execution_request_override]

	// [START code_execution_request_override_return]
	// Expected output:
	// --------------------------------------------------------------------------------
	// def is_prime(n):
	// 	if n <= 1:
	// 		return False
	// 	if n <= 3:
	// 		return True
	// 	if n % 2 == 0 or n % 3 == 0:
	// 		return False
	// 	i = 5
	// 	while i * i <= n:
	// 		if n % i == 0 or n % (i + 2) == 0:
	// 			return False
	// 		i += 6
	// 	return True

	// def sum_of_first_n_primes(n):
	// 	primes = []
	// 	num = 2
	// 	while len(primes) < n:
	// 		if is_prime(num):
	// 			primes.append(num)
	// 		num += 1
	// 	return sum(primes)
	// print(sum_of_first_n_primes(50))


	// 5117

	// [END code_execution_request_override_return]

	return response, err
}
