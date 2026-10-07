// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package samples_test

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	interactions "google.golang.org/genai/interactions"
	gaos_agents "google.golang.org/genai/interactions/models/agents"
	gaos_environments "google.golang.org/genai/interactions/models/environments"
	gaos_interactions "google.golang.org/genai/interactions/models/interactions"
	"google.golang.org/genai/interactions/models/operations"
	gaos_voices "google.golang.org/genai/interactions/models/voices"
	"log"
	"testing"
	"time"
)

func TestSample_api_version_agents_get_list(t *testing.T) {
	// [START paths['/{api_version}/agents'].get:list]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Agents.List(ctx, operations.ListAgentsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.AgentListResponse != nil {
		for _, agent := range res.AgentListResponse.Agents {
			if agent.ID != nil {
				fmt.Println(*agent.ID)
			}
		}
	}
	// [END paths['/{api_version}/agents'].get:list]
}

func TestSample_api_version_agents_post_create(t *testing.T) {
	// [START paths['/{api_version}/agents'].post:create]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	agentID := fmt.Sprintf("research-assistant-%d", time.Now().UnixNano())
	req := operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:                genai.Ptr(agentID),
			BaseAgent:         genai.Ptr("antigravity-preview-05-2026"),
			Description:       genai.Ptr("A helpful research assistant."),
			SystemInstruction: genai.Ptr("You are a helpful research assistant."),
			BaseEnvironment:   genai.Ptr(gaos_agents.NewBaseEnvironment("remote")),
			Tools: []gaos_agents.AgentTool{
				gaos_agents.NewAgentTool(gaos_interactions.GoogleSearch{}),
			},
		},
	}

	res, err := client.Agents.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Agent != nil && res.Agent.ID != nil {
		fmt.Println(*res.Agent.ID)
	}

	// [cleanup]
	_, _ = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	// [/cleanup]
	// [END paths['/{api_version}/agents'].post:create]
}

func TestSample_api_version_agents_post_with_sources(t *testing.T) {
	// [START paths['/{api_version}/agents'].post:with_sources]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	agentID := fmt.Sprintf("data-analyst-%d", time.Now().UnixNano())
	env := gaos_interactions.Environment{
		Sources: []gaos_interactions.Source{
			{
				Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
				Target:  genai.Ptr(".agents/AGENTS.md"),
				Content: genai.Ptr("Always use matplotlib for charts. Include a summary table in every report."),
			},
			{
				Type:   genai.Ptr(gaos_interactions.SourceTypeRepository),
				Source: genai.Ptr("https://github.com/my-org/analysis-templates"),
				Target: genai.Ptr("/workspace/templates"),
			},
		},
	}

	req := operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:                genai.Ptr(agentID),
			BaseAgent:         genai.Ptr("antigravity-preview-05-2026"),
			SystemInstruction: genai.Ptr("You are a data analyst. Always include visualizations and export results as PDF."),
			BaseEnvironment:   genai.Ptr(gaos_agents.NewBaseEnvironment(env)),
		},
	}

	res, err := client.Agents.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Agent != nil && res.Agent.ID != nil {
		fmt.Printf("Created agent: %s\n", *res.Agent.ID)
	}

	// [cleanup]
	_, _ = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	// [/cleanup]
	// [END paths['/{api_version}/agents'].post:with_sources]
}

func TestSample_api_version_agents_post_fork_from_env(t *testing.T) {
	// [START paths['/{api_version}/agents'].post:fork_from_env]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Step 1: Set up the environment interactively.
	interactionReq := operations.CreateInteractionRequest{
		Body: operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
			Agent:       gaos_interactions.AgentOption("antigravity-preview-05-2026"),
			Input:       genai.Ptr(gaos_interactions.NewInteractionsInput("Write a basic Hello World template to /workspace/template.py.")),
			Environment: genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment("remote")),
		}),
	}
	interactionRes, err := client.Interactions.Create(ctx, interactionReq)
	if err != nil {
		log.Fatal(err)
	}
	envID := ""
	if interactionRes.Interaction != nil && interactionRes.Interaction.EnvironmentID != nil {
		envID = *interactionRes.Interaction.EnvironmentID
	}

	// Step 2: Fork that environment into a named agent.
	agentID := fmt.Sprintf("my-data-analyst-%d", time.Now().UnixNano())
	req := operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:                genai.Ptr(agentID),
			BaseAgent:         genai.Ptr("antigravity-preview-05-2026"),
			SystemInstruction: genai.Ptr("You are a data analyst. Use the template at /workspace/template.py for all reports."),
			BaseEnvironment:   genai.Ptr(gaos_agents.NewBaseEnvironment(envID)),
		},
	}
	res, err := client.Agents.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	if res.Agent != nil && res.Agent.ID != nil {
		fmt.Printf("Forked agent: %s\n", *res.Agent.ID)
	}

	// [cleanup]
	_, _ = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	// [/cleanup]
	// [END paths['/{api_version}/agents'].post:fork_from_env]
}

func TestSample_api_version_agents_id_delete_delete(t *testing.T) {
	// [START paths['/{api_version}/agents/{id}'].delete:delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	agentID := fmt.Sprintf("delete-test-%d", time.Now().UnixNano())
	_, err = client.Agents.Create(ctx, operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:              genai.Ptr(agentID),
			BaseAgent:       genai.Ptr("antigravity-preview-05-2026"),
			Description:     genai.Ptr("Temporary agent for deletion."),
			BaseEnvironment: genai.Ptr(gaos_agents.NewBaseEnvironment("remote")),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// [/setup]

	_, err = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Agent deleted successfully.")
	// [END paths['/{api_version}/agents/{id}'].delete:delete]
}

func TestSample_api_version_agents_id_get_get(t *testing.T) {
	// [START paths['/{api_version}/agents/{id}'].get:get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	agentID := fmt.Sprintf("test-agent-%d", time.Now().UnixNano())
	_, err = client.Agents.Create(ctx, operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:              genai.Ptr(agentID),
			BaseAgent:       genai.Ptr("antigravity-preview-05-2026"),
			Description:     genai.Ptr("A test agent."),
			BaseEnvironment: genai.Ptr(gaos_agents.NewBaseEnvironment("remote")),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// [/setup]

	res, err := client.Agents.Get(ctx, operations.GetAgentRequest{ID: agentID})
	if err != nil {
		log.Fatal(err)
	}

	if res.Agent != nil && res.Agent.ID != nil {
		fmt.Println(*res.Agent.ID)
	}

	// [cleanup]
	_, _ = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	// [/cleanup]
	// [END paths['/{api_version}/agents/{id}'].get:get]
}

func TestSample_api_version_environments_get_list(t *testing.T) {
	// [START paths['/{api_version}/environments'].get:list]
	ctx := context.Background()
	client := interactions.New()

	res, err := client.Environments.ListEnvironments(ctx, operations.ListEnvironmentsRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.ListEnvironmentsResponse != nil {
		for _, env := range res.ListEnvironmentsResponse.Environments {
			if env.ID != "" {
				fmt.Println(env.ID)
			}
		}
	}
	// [END paths['/{api_version}/environments'].get:list]
}

func TestSample_api_version_environments_post_create(t *testing.T) {
	// [START paths['/{api_version}/environments'].post:create]
	ctx := context.Background()
	client := interactions.New()

	req := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("main.py"),
					Content: genai.Ptr("print('Hello, World!')"),
				},
			},
		},
	}

	res, err := client.Environments.CreateEnvironment(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Environment != nil && res.Environment.ID != "" {
		fmt.Println(res.Environment.ID)

		// [cleanup]
		_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{ID: res.Environment.ID})
		// [/cleanup]
	}
	// [END paths['/{api_version}/environments'].post:create]
}

func TestSample_api_version_environments_post_copy(t *testing.T) {
	// [START paths['/{api_version}/environments'].post:copy]
	ctx := context.Background()
	client := interactions.New()

	// [setup]
	sourceReq := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("main.py"),
					Content: genai.Ptr("print('Hello')"),
				},
			},
		},
	}
	sourceRes, err := client.Environments.CreateEnvironment(ctx, sourceReq)
	if err != nil {
		log.Fatal(err)
	}
	if sourceRes.Environment == nil || sourceRes.Environment.ID == "" {
		log.Fatal("missing source environment ID")
	}
	sourceEnvID := sourceRes.Environment.ID
	// [/setup]

	res, err := client.Environments.CreateEnvironment(ctx, operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			FromEnvironment: genai.Ptr(sourceEnvID),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Environment != nil && res.Environment.ID != "" {
		fmt.Println(res.Environment.ID)

		// [cleanup]
		_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{ID: res.Environment.ID})
		_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{ID: sourceEnvID})
		// [/cleanup]
	}
	// [END paths['/{api_version}/environments'].post:copy]
}

func TestSample_api_version_environments_environment_files_path_get_list_files(t *testing.T) {
	// [START paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
	ctx := context.Background()
	client := interactions.New()

	// [setup]
	createReq := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("src/main.py"),
					Content: genai.Ptr("print('Hello')"),
				},
			},
		},
	}
	createRes, err := client.Environments.CreateEnvironment(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if createRes.Environment == nil || createRes.Environment.ID == "" {
		log.Fatal("missing environment ID")
	}
	envID := createRes.Environment.ID
	// [/setup]

	res, err := client.Environments.Files.List(ctx, operations.GetEnvironmentFilesRequest{
		Environment: envID,
		Path:        "src",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.GetEnvironmentFilesResponse != nil {
		for _, file := range res.GetEnvironmentFilesResponse.Files {
			name := ""
			if file.Name != nil {
				name = *file.Name
			}
			fileType := ""
			if file.Type != nil {
				fileType = string(*file.Type)
			}
			var size int64
			if file.SizeBytes != nil {
				size = *file.SizeBytes
			}
			fmt.Println(name, fileType, size)
		}
	}

	// [cleanup]
	_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{
		ID: envID,
	})
	// [/cleanup]
	// [END paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
}

func TestSample_api_version_environments_environment_files_path_get_get_file(t *testing.T) {
	// [START paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
	ctx := context.Background()
	client := interactions.New()

	// [setup]
	createReq := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("main.py"),
					Content: genai.Ptr("print('Hello')"),
				},
			},
		},
	}
	createRes, err := client.Environments.CreateEnvironment(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if createRes.Environment == nil || createRes.Environment.ID == "" {
		log.Fatal("missing environment ID")
	}
	envID := createRes.Environment.ID
	// [/setup]

	res, err := client.Environments.Files.List(ctx, operations.GetEnvironmentFilesRequest{
		Environment: envID,
		Path:        "main.py",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.GetEnvironmentFilesResponse != nil {
		for _, file := range res.GetEnvironmentFilesResponse.Files {
			name := ""
			if file.Name != nil {
				name = *file.Name
			}
			var size int64
			if file.SizeBytes != nil {
				size = *file.SizeBytes
			}
			fmt.Println(name, size)
		}
	}

	// [cleanup]
	_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{
		ID: envID,
	})
	// [/cleanup]
	// [END paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
}

func TestSample_api_version_environments_id_delete_delete(t *testing.T) {
	// [START paths['/{api_version}/environments/{id}'].delete:delete]
	ctx := context.Background()
	client := interactions.New()

	// [setup]
	createReq := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("main.py"),
					Content: genai.Ptr("print('Hello')"),
				},
			},
		},
	}
	createRes, err := client.Environments.CreateEnvironment(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if createRes.Environment == nil || createRes.Environment.ID == "" {
		log.Fatal("missing environment ID")
	}
	envID := createRes.Environment.ID
	// [/setup]

	_, err = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{
		ID: envID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Environment deleted successfully.")
	// [END paths['/{api_version}/environments/{id}'].delete:delete]
}

func TestSample_api_version_environments_id_get_get(t *testing.T) {
	// [START paths['/{api_version}/environments/{id}'].get:get]
	ctx := context.Background()
	client := interactions.New()

	// [setup]
	createReq := operations.CreateEnvironmentRequest{
		Body: gaos_environments.CreateEnvironmentRequest{
			Sources: []gaos_interactions.Source{
				{
					Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
					Target:  genai.Ptr("main.py"),
					Content: genai.Ptr("print('Hello')"),
				},
			},
		},
	}
	createRes, err := client.Environments.CreateEnvironment(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if createRes.Environment == nil || createRes.Environment.ID == "" {
		log.Fatal("missing environment ID")
	}
	envID := createRes.Environment.ID
	// [/setup]

	res, err := client.Environments.GetEnvironment(ctx, operations.GetEnvironmentRequest{
		ID: envID,
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Environment != nil && res.Environment.ID != "" {
		fmt.Println(res.Environment.ID)
	}

	// [cleanup]
	_, _ = client.Environments.DeleteEnvironment(ctx, operations.DeleteEnvironmentRequest{
		ID: envID,
	})
	// [/cleanup]
	// [END paths['/{api_version}/environments/{id}'].get:get]
}

func TestSample_api_version_interactions_post_simple(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:simple]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Hello, how are you?")),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:simple]
}

func TestSample_api_version_interactions_post_multi_turn(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:multi_turn]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	history := []gaos_interactions.Step{
		gaos_interactions.NewStep(gaos_interactions.UserInputStep{
			Content: []gaos_interactions.Content{
				gaos_interactions.NewContent(gaos_interactions.TextContent{
					Text: "Hello!",
				}),
			},
		}),
		gaos_interactions.NewStep(gaos_interactions.ModelOutputStep{
			Content: []gaos_interactions.Content{
				gaos_interactions.NewContent(gaos_interactions.TextContent{
					Text: "Hi there! How can I help you today?",
				}),
			},
		}),
		gaos_interactions.NewStep(gaos_interactions.UserInputStep{
			Content: []gaos_interactions.Content{
				gaos_interactions.NewContent(gaos_interactions.TextContent{
					Text: "What is the capital of France?",
				}),
			},
		}),
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput(history)),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:multi_turn]
}

func TestSample_api_version_interactions_post_multimodal_image(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:multimodal_image]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	contents := []gaos_interactions.Content{
		gaos_interactions.NewContent(gaos_interactions.TextContent{
			Text: "What is in this picture?",
		}),
		gaos_interactions.NewContent(gaos_interactions.ImageContent{
			Data:     genai.Ptr("BASE64_ENCODED_IMAGE"),
			MimeType: gaos_interactions.ImageContentMimeTypeImagePng.ToPointer(),
		}),
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput(contents)),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:multimodal_image]
}

func TestSample_api_version_interactions_post_function_calling(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:function_calling]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	function := gaos_interactions.Function{
		Name:        genai.Ptr("get_weather"),
		Description: genai.Ptr("Get the current weather in a given location"),
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"location": map[string]any{
					"type":        "string",
					"description": "The city and state, e.g. San Francisco, CA",
				},
			},
			"required": []string{"location"},
		},
	}

	tool := gaos_interactions.NewTool(function)

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("What is the weather like in Boston, MA?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil {
		for _, step := range res.Interaction.Steps {
			if step.FunctionCallStep != nil {
				fmt.Println("Function Call:", step.FunctionCallStep.Name)
				fmt.Println("Arguments:", step.FunctionCallStep.Arguments)
			}
		}
	}
	// [END paths['/{api_version}/interactions'].post:function_calling]
}

func TestSample_api_version_interactions_post_deep_research(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:deep_research]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:      gaos_interactions.AgentOption("deep-research-pro-preview-12-2025"),
		Input:      genai.Ptr(gaos_interactions.NewInteractionsInput("find a cure to cancer")),
		Background: genai.Ptr(true),
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil {
		fmt.Println(res.Interaction.Status)
	}
	// [END paths['/{api_version}/interactions'].post:deep_research]
}

func TestSample_api_version_interactions_post_antigravity(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:antigravity]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:       gaos_interactions.AgentOption("antigravity-preview-05-2026"),
		Input:       genai.Ptr(gaos_interactions.NewInteractionsInput("Read Hacker News, summarize the top 5 stories, and save results as a markdown file.")),
		Environment: genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment("remote")),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:antigravity]
}

func TestSample_api_version_interactions_post_reuse_env(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:reuse_env]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Step 1: Create an interaction with a fresh remote environment.
	body1 := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:       gaos_interactions.AgentOption("antigravity-preview-05-2026"),
		Input:       genai.Ptr(gaos_interactions.NewInteractionsInput("Write a hello world script at /workspace/hello.py.")),
		Environment: genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment("remote")),
	})

	res1, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body1})
	if err != nil {
		log.Fatal(err)
	}

	if res1.Interaction == nil || res1.Interaction.EnvironmentID == nil {
		log.Fatal("No environment ID returned")
	}
	fmt.Printf("Environment ID: %s\n", *res1.Interaction.EnvironmentID)

	// Step 2: Reuse the same environment in a follow-up interaction.
	body2 := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:                 gaos_interactions.AgentOption("antigravity-preview-05-2026"),
		Input:                 genai.Ptr(gaos_interactions.NewInteractionsInput("Modify the script to accept a name argument and greet the user.")),
		Environment:           genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment(*res1.Interaction.EnvironmentID)),
		PreviousInteractionID: res1.Interaction.ID,
	})

	res2, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body2})
	if err != nil {
		log.Fatal(err)
	}

	if res2.Interaction != nil && res2.Interaction.OutputText != nil {
		fmt.Println(*res2.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:reuse_env]
}

func TestSample_api_version_interactions_post_with_sources(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:with_sources]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	env := gaos_interactions.Environment{
		Sources: []gaos_interactions.Source{
			{
				Type:   genai.Ptr(gaos_interactions.SourceTypeRepository),
				Source: genai.Ptr("https://github.com/octocat/Spoon-Knife"),
				Target: genai.Ptr("/workspace/repo"),
			},
			{
				Type:    genai.Ptr(gaos_interactions.SourceTypeInline),
				Content: genai.Ptr("Focus on Python files only."),
				Target:  genai.Ptr("/workspace/notes.txt"),
			},
		},
	}

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:       gaos_interactions.AgentOption("antigravity-preview-05-2026"),
		Input:       genai.Ptr(gaos_interactions.NewInteractionsInput("List all files under /workspace and summarize what you find.")),
		Environment: genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment(env)),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END paths['/{api_version}/interactions'].post:with_sources]
}

func TestSample_api_version_interactions_post_custom_agent(t *testing.T) {
	// [START paths['/{api_version}/interactions'].post:custom_agent]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Step 1: Create a custom agent.
	agentID := fmt.Sprintf("code-reviewer-%d", time.Now().UnixNano())
	baseEnv := gaos_interactions.Environment{
		Sources: []gaos_interactions.Source{
			{
				Type:   genai.Ptr(gaos_interactions.SourceTypeRepository),
				Source: genai.Ptr("https://github.com/octocat/Spoon-Knife"),
				Target: genai.Ptr("/workspace/repo"),
			},
		},
	}

	_, err = client.Agents.Create(ctx, operations.CreateAgentRequest{
		Body: gaos_agents.Agent{
			ID:                genai.Ptr(agentID),
			BaseAgent:         genai.Ptr("antigravity-preview-05-2026"),
			SystemInstruction: genai.Ptr("You are a senior code reviewer. Check every file for bugs, style issues, and security vulnerabilities."),
			BaseEnvironment:   genai.Ptr(gaos_agents.NewBaseEnvironment(baseEnv)),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Use the custom agent.
	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateAgentInteraction{
		Agent:       gaos_interactions.AgentOption(agentID),
		Input:       genai.Ptr(gaos_interactions.NewInteractionsInput("Review the latest changes in /workspace/repo/src and file a summary.")),
		Environment: genai.Ptr(gaos_interactions.NewCreateAgentInteractionEnvironment("remote")),
	})

	res, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}

	// [cleanup]
	_, _ = client.Agents.Delete(ctx, operations.DeleteAgentRequest{ID: agentID})
	// [/cleanup]
	// [END paths['/{api_version}/interactions'].post:custom_agent]
}

func TestSample_api_version_interactions_id_delete_delete(t *testing.T) {
	// [START paths['/{api_version}/interactions/{id}'].delete:delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Hello")),
	})
	createReq := operations.CreateInteractionRequest{
		Body: body,
	}

	created, err := client.Interactions.Create(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if created.Interaction == nil || created.Interaction.ID == nil {
		log.Fatal("failed to create interaction in setup")
	}
	interactionID := *created.Interaction.ID
	// [/setup]

	req := operations.DeleteInteractionRequest{
		ID: interactionID,
	}

	_, err = client.Interactions.Delete(
		ctx,
		req,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Interaction deleted successfully.")
	// [END paths['/{api_version}/interactions/{id}'].delete:delete]
}

func TestSample_api_version_interactions_id_get_get(t *testing.T) {
	// [START paths['/{api_version}/interactions/{id}'].get:get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Say hello.")),
	})
	created, err := client.Interactions.Create(ctx, operations.CreateInteractionRequest{Body: body})
	if err != nil {
		log.Fatal(err)
	}
	if created.Interaction == nil || created.Interaction.ID == nil {
		log.Fatal("No interaction ID returned")
	}
	// [/setup]

	interaction, err := client.Interactions.Get(ctx, operations.GetInteractionByIDRequest{
		ID: *created.Interaction.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	if interaction.Interaction != nil {
		fmt.Println(interaction.Interaction.Status)
	}
	// [END paths['/{api_version}/interactions/{id}'].get:get]
}

func TestSample_api_version_interactions_id_cancel_post_cancel(t *testing.T) {
	// [START paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model:      gaos_interactions.Model("gemini-3.6-flash"),
		Input:      genai.Ptr(gaos_interactions.NewInteractionsInput("Write a long essay about the history of computing.")),
		Background: genai.Ptr(true),
	})
	createReq := operations.CreateInteractionRequest{
		Body: body,
	}

	created, err := client.Interactions.Create(ctx, createReq)
	if err != nil {
		log.Fatal(err)
	}
	if created.Interaction == nil || created.Interaction.ID == nil {
		log.Fatal("failed to create interaction in setup")
	}
	interactionID := *created.Interaction.ID
	// [/setup]

	req := operations.CancelInteractionByIDRequest{
		ID: interactionID,
	}

	res, err := client.Interactions.Cancel(
		ctx,
		req,
	)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil {
		fmt.Println(res.Interaction.Status)
	}
	// [END paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
}

func TestSample_schema_CodeExecution_code_execution(t *testing.T) {
	// [START components.schemas.CodeExecution:code_execution]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tool := gaos_interactions.NewTool(gaos_interactions.CodeExecution{})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Calculate the first 10 Fibonacci numbers")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.CodeExecution:code_execution]
}

func TestSample_schema_ComputerUse_computer_use(t *testing.T) {
	// [START components.schemas.ComputerUse:computer_use]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tool := gaos_interactions.NewTool(gaos_interactions.ComputerUse{})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-2.5-computer-use-preview-10-2025"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Find a flight to Tokyo")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.ComputerUse:computer_use]
}

func TestSample_schema_FileSearch_file_search(t *testing.T) {
	// [START components.schemas.FileSearch:file_search]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// [setup]
	store, err := client.FileSearchStores.Create(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	// [/setup]

	tool := gaos_interactions.NewTool(gaos_interactions.FileSearch{
		FileSearchStoreNames: []string{store.Name},
	})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("What documents are available?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}

	// [cleanup]
	_ = client.FileSearchStores.Delete(ctx, store.Name, nil)
	// [/cleanup]
	// [END components.schemas.FileSearch:file_search]
}

func TestSample_schema_Function_function_calling(t *testing.T) {
	// [START components.schemas.Function:function_calling]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	function := gaos_interactions.Function{
		Name:        genai.Ptr("get_weather"),
		Description: genai.Ptr("Get the current weather in a given location"),
		Parameters: map[string]any{
			"type": "object",
			"properties": map[string]any{
				"location": map[string]any{
					"type":        "string",
					"description": "The city and state, e.g. San Francisco, CA",
				},
			},
			"required": []string{"location"},
		},
	}

	tool := gaos_interactions.NewTool(function)

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("What is the weather like in Boston, MA?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil {
		for _, step := range res.Interaction.Steps {
			if step.FunctionCallStep != nil {
				fmt.Println("Function Call:", step.FunctionCallStep.Name)
				fmt.Println("Arguments:", step.FunctionCallStep.Arguments)
			}
		}
	}
	// [END components.schemas.Function:function_calling]
}

func TestSample_schema_GoogleMaps_google_maps(t *testing.T) {
	// [START components.schemas.GoogleMaps:google_maps]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tool := gaos_interactions.NewTool(gaos_interactions.GoogleMaps{
		Latitude:  genai.Ptr(37.7749),
		Longitude: genai.Ptr(-122.4194),
	})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("What is the best food near me?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.GoogleMaps:google_maps]
}

func TestSample_schema_GoogleSearch_google_search(t *testing.T) {
	// [START components.schemas.GoogleSearch:google_search]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tool := gaos_interactions.NewTool(gaos_interactions.GoogleSearch{})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Who is the current president of France?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.GoogleSearch:google_search]
}

func TestSample_schema_McpServer_mcp_server(t *testing.T) {
	// [START components.schemas.McpServer:mcp_server]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mcpServer := gaos_interactions.MCPServer{
		Name: genai.Ptr("weather_service"),
		URL:  genai.Ptr("https://gemini-api-demos.uc.r.appspot.com/mcp"),
	}
	tool := gaos_interactions.NewTool(mcpServer)

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Today is 12-05-2025, what is the temperature today in London?")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.McpServer:mcp_server]
}

func TestSample_schema_UrlContext_url_context(t *testing.T) {
	// [START components.schemas.UrlContext:url_context]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tool := gaos_interactions.NewTool(gaos_interactions.URLContext{})

	body := operations.NewCreateInteractionRequestBody(gaos_interactions.CreateModelInteraction{
		Model: gaos_interactions.Model("gemini-3.6-flash"),
		Input: genai.Ptr(gaos_interactions.NewInteractionsInput("Summarize https://www.example.com")),
		Tools: []gaos_interactions.Tool{tool},
	})

	req := operations.CreateInteractionRequest{
		Body: body,
	}

	res, err := client.Interactions.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Interaction != nil && res.Interaction.OutputText != nil {
		fmt.Println(*res.Interaction.OutputText)
	}
	// [END components.schemas.UrlContext:url_context]
}

func TestSample_api_version_voices_get_list(t *testing.T) {
	// [START paths['/{api_version}/voices'].get:list]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Voices.List(ctx, operations.ListVoicesRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.ListVoicesResponse != nil {
		for _, voice := range res.ListVoicesResponse.Voices {
			if voice.ID != nil && voice.DisplayName != nil {
				fmt.Println(*voice.ID, *voice.DisplayName)
			}
		}
	}
	// [END paths['/{api_version}/voices'].get:list]
}

func TestSample_api_version_voices_post_create(t *testing.T) {
	// [START paths['/{api_version}/voices'].post:create]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	req := operations.CreateVoiceRequest{
		Body: gaos_voices.CreateVoiceRequest{
			Store: genai.Ptr(true),
			Voice: gaos_voices.VoiceInput{
				Type:         gaos_voices.VoiceTypePrompted,
				DisplayName:  genai.Ptr("Warm Narrator"),
				LanguageCode: genai.Ptr("en-US"),
				Prompted: &gaos_voices.PromptedVoice{
					Input: "A warm, friendly narrator voice with a calm pace.",
				},
			},
		},
	}

	res, err := client.Voices.Create(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Voice != nil && res.Voice.ID != nil {
		fmt.Println(*res.Voice.ID)
	}

	// [cleanup]
	if res.Voice != nil && res.Voice.ID != nil {
		_, _ = client.Voices.Delete(ctx, operations.DeleteVoiceRequest{ID: *res.Voice.ID})
	}
	// [/cleanup]
	// [END paths['/{api_version}/voices'].post:create]
}

func TestSample_api_version_voices_voicesid_get_get(t *testing.T) {
	// [START paths['/{api_version}/voices/{voicesId}'].get:get]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Voices.Get(ctx, operations.GetVoiceRequest{
		ID: "voice_abc123",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Voice != nil && res.Voice.ID != nil && res.Voice.DisplayName != nil {
		fmt.Println(*res.Voice.ID, *res.Voice.DisplayName)
	}
	// [END paths['/{api_version}/voices/{voicesId}'].get:get]
}

func TestSample_api_version_voices_voicesid_delete_delete(t *testing.T) {
	// [START paths['/{api_version}/voices/{voicesId}'].delete:delete]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Voices.Delete(ctx, operations.DeleteVoiceRequest{
		ID: "voice_abc123",
	})
	if err != nil {
		log.Fatal(err)
	}
	// [END paths['/{api_version}/voices/{voicesId}'].delete:delete]
}
