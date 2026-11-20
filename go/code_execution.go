package examples

import (
	"context"
	"fmt"
	"log"
	"os"

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
		"gemini-3.8-flash",
		genai.Text(
			"What is the sum of the first 50 prime numbers? Generate and run code for the calculation, and make sure you get all 50.",
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
	// [END code_execution_basic]

	// [START code_execution_basic_return]
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

	// [END code_execution_basic_return]

	return response, err
}
