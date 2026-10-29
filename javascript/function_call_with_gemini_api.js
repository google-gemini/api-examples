//This example defines a simple function to get the current weather and uses the Gemini API to interpret a user query, call the function, and return a response.

const { GoogleGenerativeAI } = require("@google/generative-ai");

// Initialize the Gemini API client with your API key
const genAI = new GoogleGenerativeAI(process.env.API_KEY);

// Define a sample function to simulate getting the weather
function getWeather(location) {
  // Simulated weather data (in a real app, this could call an external API)
  const weatherData = {
    "New York": { temp: 65, condition: "Sunny" },
    "London": { temp: 50, condition: "Rainy" },
  };
  return weatherData[location] || { temp: "unknown", condition: "unknown" };
}

// Define the function schema for the Gemini API
const tools = [
  {
    functionDeclarations: [
      {
        name: "getWeather",
        description: "Get the current weather for a specified location",
        parameters: {
          type: "object",
          properties: {
            location: {
              type: "string",
              description: "The city or location to get the weather for",
            },
          },
          required: ["location"],
        },
      },
    ],
  },
];

// Test case for function calling
async function testFunctionCalling() {
  // [START function_calling_weather]
  const model = genAI.getGenerativeModel({
    model: "gemini-1.5-flash",
    tools: tools,
  });

  const prompt = "What's the weather like in New York today?";

  // Generate content with function calling enabled
  const result = await model.generateContent(prompt);
  const response = result.response;

  // Check if the model wants to call a function
  const functionCall = response.functionCalls && response.functionCalls()[0];
  if (functionCall && functionCall.name === "getWeather") {
    const { location } = functionCall.args;
    const weather = getWeather(location);

    // Send the function result back to the model for a final response
    const finalResult = await model.generateContent([
      prompt,
      {
        functionCall: {
          name: "getWeather",
          args: { location },
        },
      },
      {
        functionResponse: {
          name: "getWeather",
          response: weather,
        },
      },
    ]);

    console.log(finalResult.response.text());
  } else {
    console.log(response.text());
  }
  // [END function_calling_weather]
}

// Run the test
testFunctionCalling().catch(console.error);