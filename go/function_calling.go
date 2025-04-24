package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

// Arithmetic functions.
func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

// ArithmeticArgs represents the expected arguments for our arithmetic operations.
type ArithmeticArgs struct {
	FirstParam  float64 `json:"firstParam"`
	SecondParam float64 `json:"secondParam"`
}

// createArithmeticToolDeclaration creates a function declaration with the given name and description.
// The parameters schema includes "firstParam" and "secondParam" as required numbers.
func createArithmeticToolDeclaration(name, description string) *genai.FunctionDeclaration {
	paramSchema := &genai.Schema{
		Type: genai.TypeObject,
		Description: "The result of the arithmetic operation.",
		Properties: map[string]*genai.Schema{
			"firstParam": {
				Type:        genai.TypeNumber,
				Description: "The first parameter which can be an integer or a floating point number.",
			},
			"secondParam": {
				Type:        genai.TypeNumber,
				Description: "The second parameter which can be an integer or a floating point number.",
			},
		},
		Required: []string{"firstParam", "secondParam"},
	}
	return &genai.FunctionDeclaration{
		Name:        name,
		Description: description,
		Parameters:  paramSchema,
	}
}

func FunctionCalling() error {
	// [START function_calling]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	modelName := "gemini-2.0-flash"

	// Create the function declarations for arithmetic operations.
	addDeclaration := createArithmeticToolDeclaration("addNumbers", "Return the result of adding two numbers.")
	subtractDeclaration := createArithmeticToolDeclaration("subtractNumbers", "Return the result of subtracting the second number from the first.")
	multiplyDeclaration := createArithmeticToolDeclaration("multiplyNumbers", "Return the product of two numbers.")
	divideDeclaration := createArithmeticToolDeclaration("divideNumbers", "Return the quotient of dividing the first number by the second.")

	// Group the function declarations as a tool.
	tools := []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				addDeclaration,
				subtractDeclaration,
				multiplyDeclaration,
				divideDeclaration,
			},
		},
	}

	// Create the content prompt.
	contents := []*genai.Content{
		genai.NewContentFromText(
			"I have 57 cats, each owns 44 mittens, how many mittens is that in total?", genai.RoleUser,
		),
	}

	// Set up the generate content configuration with function calling enabled.
	config := &genai.GenerateContentConfig{
		Tools: tools,
		ToolConfig: &genai.ToolConfig{
			FunctionCallingConfig: &genai.FunctionCallingConfig{
				// The mode equivalent to FunctionCallingConfigMode.ANY in JS.
				Mode: genai.FunctionCallingConfigModeAny,
			},
		},
	}

	genContentResp, err := client.Models.GenerateContent(ctx, modelName, contents, config)
	if err != nil {
		log.Fatal(err)
	}

	// Assume the response includes a list of function calls.
	if len(genContentResp.FunctionCalls()) == 0 {
		log.Println("No function call returned from the AI.")
		return nil
	}
	functionCall := genContentResp.FunctionCalls()[0]
	log.Printf("Function call: %+v\n", functionCall)

	// Marshal the Args map into JSON bytes.
	argsMap, err := json.Marshal(functionCall.Args)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON bytes into the ArithmeticArgs struct.
	var args ArithmeticArgs
	if err := json.Unmarshal(argsMap, &args); err != nil {
		log.Fatal(err)
	}

	// Map the function name to the actual arithmetic function.
	var result float64
	switch functionCall.Name {
		case "addNumbers":
			result = add(args.FirstParam, args.SecondParam)
		case "subtractNumbers":
			result = subtract(args.FirstParam, args.SecondParam)
		case "multiplyNumbers":
			result = multiply(args.FirstParam, args.SecondParam)
		case "divideNumbers":
			result = divide(args.FirstParam, args.SecondParam)
		default:
			return fmt.Errorf("unimplemented function: %s", functionCall.Name)
	}
	log.Printf("Function result: %v\n", result)

	// Prepare the final result message as content.
	resultContents := []*genai.Content{
		genai.NewContentFromText("The final result is " + fmt.Sprintf("%v", result), genai.RoleUser),
	}

	// Use GenerateContent to send the final result.
	finalResponse, err := client.Models.GenerateContent(ctx, modelName, resultContents, &genai.GenerateContentConfig{})
	if err != nil {
		log.Fatal(err)
	}

	printResponse(finalResponse)
	// [END function_calling]
	return err
}
