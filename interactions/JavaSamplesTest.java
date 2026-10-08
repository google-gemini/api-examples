/**
 * Copyright 2026 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.google.genai.examples;

import com.google.genai.Client;
import com.google.genai.gaos.models.agents.Agent;
import com.google.genai.gaos.models.agents.AgentListResponse;
import com.google.genai.gaos.models.agents.BaseEnvironment;
import com.google.genai.gaos.models.environments.CreateEnvironmentRequest;
import com.google.genai.gaos.models.environments.EnvironmentFile;
import com.google.genai.gaos.models.environments.GetEnvironmentFilesResponse;
import com.google.genai.gaos.models.interactions.AgentOption;
import com.google.genai.gaos.models.interactions.CodeExecution;
import com.google.genai.gaos.models.interactions.ComputerUse;
import com.google.genai.gaos.models.interactions.Content;
import com.google.genai.gaos.models.interactions.CreateAgentInteraction;
import com.google.genai.gaos.models.interactions.CreateAgentInteractionEnvironment;
import com.google.genai.gaos.models.interactions.CreateModelInteraction;
import com.google.genai.gaos.models.interactions.FileSearch;
import com.google.genai.gaos.models.interactions.Function;
import com.google.genai.gaos.models.interactions.GoogleMaps;
import com.google.genai.gaos.models.interactions.GoogleSearch;
import com.google.genai.gaos.models.interactions.ImageContent;
import com.google.genai.gaos.models.interactions.ImageContentMimeType;
import com.google.genai.gaos.models.interactions.Interaction;
import com.google.genai.gaos.models.interactions.InteractionStatus;
import com.google.genai.gaos.models.interactions.InteractionsInput;
import com.google.genai.gaos.models.interactions.MCPServer;
import com.google.genai.gaos.models.interactions.ModelOutputStep;
import com.google.genai.gaos.models.interactions.Source;
import com.google.genai.gaos.models.interactions.SourceType;
import com.google.genai.gaos.models.interactions.Step;
import com.google.genai.gaos.models.interactions.TextContent;
import com.google.genai.gaos.models.interactions.URLContext;
import com.google.genai.gaos.models.interactions.UserInputStep;
import com.google.genai.gaos.models.operations.CancelInteractionByIdResponse;
import com.google.genai.gaos.models.operations.CreateAgentResponse;
import com.google.genai.gaos.models.operations.CreateInteractionRequestBody;
import com.google.genai.gaos.models.operations.CreateInteractionResponse;
import com.google.genai.gaos.models.operations.CreateTriggerResponse;
import com.google.genai.gaos.models.operations.GetAgentResponse;
import com.google.genai.gaos.models.operations.GetEnvironmentFilesRequest;
import com.google.genai.gaos.models.operations.GetInteractionByIdRequest;
import com.google.genai.gaos.models.operations.GetInteractionByIdResponse;
import com.google.genai.gaos.models.operations.ListAgentsResponse;
import com.google.genai.gaos.models.operations.ListEnvironmentsResponse;
import com.google.genai.gaos.models.triggers.Trigger;
import com.google.genai.gaos.models.triggers.TriggerCreateParams;
import com.google.genai.types.CreateFileSearchStoreConfig;
import com.google.genai.types.FileSearchStore;
import java.util.List;
import java.util.Map;
import java.util.UUID;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

@RunWith(JUnit4.class)
public class JavaSamplesTest {

  @Test
  public void testApiVersionAgentsGetList() throws Exception {
    // [START paths['/{api_version}/agents'].get:list]

    Client client = new Client();
    ListAgentsResponse response = client.agents.list().call();
    List<Agent> agents =
        response.agentListResponse().flatMap(AgentListResponse::agents).orElse(List.of());
    for (Agent agent : agents) {
      agent.id().ifPresent(System.out::println);
    }
    // [END paths['/{api_version}/agents'].get:list]
  }

  @Test
  public void testApiVersionAgentsPostCreate() throws Exception {
    // [START paths['/{api_version}/agents'].post:create]

    Client client = new Client();
    String agentId = "research-assistant-" + UUID.randomUUID().toString().substring(0, 8);
    Agent agent =
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .description("A helpful research assistant.")
            .systemInstruction("You are a helpful research assistant.")
            .baseEnvironment(BaseEnvironment.of("remote"))
            .tools(List.of(new GoogleSearch()))
            .build();
    CreateAgentResponse response = client.agents.create(agent);
    System.out.println(response.agent().flatMap(Agent::id).orElse(""));

    // [cleanup]
    client.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:create]
  }

  @Test
  public void testApiVersionAgentsPostWithSources() throws Exception {
    // [START paths['/{api_version}/agents'].post:with_sources]

    Client client = new Client();
    String agentId = "data-analyst-" + UUID.randomUUID().toString().substring(0, 8);
    com.google.genai.gaos.models.interactions.Environment env =
        com.google.genai.gaos.models.interactions.Environment.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target(".agents/AGENTS.md")
                        .content(
                            "Always use matplotlib for charts. Include a summary table in every"
                                + " report.")
                        .build(),
                    Source.builder()
                        .type(SourceType.REPOSITORY)
                        .source("https://github.com/my-org/analysis-templates")
                        .target("/workspace/templates")
                        .build()))
            .build();
    Agent agent =
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .systemInstruction(
                "You are a data analyst. Always include visualizations and export results as PDF.")
            .baseEnvironment(BaseEnvironment.of(env))
            .build();
    CreateAgentResponse response = client.agents.create(agent);
    System.out.println("Created agent: " + response.agent().flatMap(Agent::id).orElse(""));

    // [cleanup]
    client.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:with_sources]
  }

  @Test
  public void testApiVersionAgentsPostForkFromEnv() throws Exception {
    // [START paths['/{api_version}/agents'].post:fork_from_env]

    Client client = new Client();

    // Step 1: Set up the environment interactively.
    CreateAgentInteraction params =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(
                InteractionsInput.of(
                    "Write a basic Hello World template to /workspace/template.py."))
            .environment(CreateAgentInteractionEnvironment.of("remote"))
            .build();
    CreateInteractionResponse interactionResponse =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        interactionResponse
            .interaction()
            .orElseThrow(() -> new RuntimeException("No interaction returned"));

    // Step 2: Fork that environment into a named agent.
    String agentId = "my-data-analyst-" + UUID.randomUUID().toString().substring(0, 8);
    Agent agent =
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .systemInstruction(
                "You are a data analyst. Use the template at /workspace/template.py for all"
                    + " reports.")
            .baseEnvironment(BaseEnvironment.of(interaction.environmentId().orElse("")))
            .build();
    CreateAgentResponse agentResponse = client.agents.create(agent);
    System.out.println("Forked agent: " + agentResponse.agent().flatMap(Agent::id).orElse(""));

    // [cleanup]
    client.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:fork_from_env]
  }

  @Test
  public void testApiVersionAgentsIdDeleteDelete() throws Exception {
    // [START paths['/{api_version}/agents/{id}'].delete:delete]

    Client client = new Client();

    // [setup]
    String agentId = "delete-test-" + UUID.randomUUID().toString().substring(0, 8);
    client.agents.create(
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .description("Temporary agent for deletion.")
            .baseEnvironment(BaseEnvironment.of("remote"))
            .build());
    // [/setup]

    client.agents.delete(agentId);
    System.out.println("Agent deleted successfully.");
    // [END paths['/{api_version}/agents/{id}'].delete:delete]
  }

  @Test
  public void testApiVersionAgentsIdGetGet() throws Exception {
    // [START paths['/{api_version}/agents/{id}'].get:get]

    Client client = new Client();

    // [setup]
    String agentId = "test-agent-" + UUID.randomUUID().toString().substring(0, 8);
    client.agents.create(
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .description("A test agent.")
            .baseEnvironment(BaseEnvironment.of("remote"))
            .build());
    // [/setup]

    GetAgentResponse response = client.agents.get(agentId);
    Agent agent = response.agent().orElseThrow(() -> new RuntimeException("No agent returned"));
    System.out.println(agent.id().orElse(""));

    // [cleanup]
    client.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/agents/{id}'].get:get]
  }

  @Test
  public void testApiVersionEnvironmentsGetList() throws Exception {
    // [START paths['/{api_version}/environments'].get:list]

    Client client = new Client();
    ListEnvironmentsResponse response = client.environments.listEnvironmentsDirect();
    for (com.google.genai.gaos.models.environments.Environment env :
        response.listEnvironmentsResponse().flatMap(res -> res.environments()).orElse(List.of())) {
      System.out.println(env.id().orElse(""));
    }
    // [END paths['/{api_version}/environments'].get:list]
  }

  @Test
  public void testApiVersionEnvironmentsPostCreate() throws Exception {
    // [START paths['/{api_version}/environments'].post:create]

    Client client = new Client();
    CreateEnvironmentRequest request =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("main.py")
                        .content("print('Hello, World!')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment environment =
        client.environments.createEnvironment(request).environment().orElseThrow();
    System.out.println(environment.id().orElse(""));

    // [cleanup]
    environment.id().ifPresent(id -> client.environments.deleteEnvironment(id));
    // [/cleanup]
    // [END paths['/{api_version}/environments'].post:create]
  }

  @Test
  public void testApiVersionEnvironmentsPostCopy() throws Exception {
    // [START paths['/{api_version}/environments'].post:copy]

    Client client = new Client();

    // [setup]
    CreateEnvironmentRequest setupReq =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("main.py")
                        .content("print('Hello')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment sourceEnv =
        client.environments.createEnvironment(setupReq).environment().orElseThrow();
    String sourceEnvId = sourceEnv.id().orElseThrow();
    // [/setup]

    CreateEnvironmentRequest request =
        CreateEnvironmentRequest.builder().fromEnvironment(sourceEnvId).build();
    com.google.genai.gaos.models.environments.Environment environment =
        client.environments.createEnvironment(request).environment().orElseThrow();
    System.out.println(environment.id().orElse(""));

    // [cleanup]
    environment.id().ifPresent(id -> client.environments.deleteEnvironment(id));
    client.environments.deleteEnvironment(sourceEnvId);
    // [/cleanup]
    // [END paths['/{api_version}/environments'].post:copy]
  }

  @Test
  public void testApiVersionEnvironmentsEnvironmentFilesPathGetListFiles() throws Exception {
    // [START paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]

    Client client = new Client();

    // [setup]
    CreateEnvironmentRequest createReq =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("src/main.py")
                        .content("print('Hello')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment created =
        client.environments.createEnvironment(createReq).environment().orElseThrow();
    String envId = created.id().orElseThrow();
    // [/setup]

    com.google.genai.gaos.models.operations.GetEnvironmentFilesResponse response =
        client
            .environments
            .files()
            .list(GetEnvironmentFilesRequest.builder().environment(envId).path("src").build());
    GetEnvironmentFilesResponse filesResponse =
        response.getEnvironmentFilesResponse().orElseThrow();
    for (EnvironmentFile file : filesResponse.files().orElse(List.of())) {
      System.out.println(
          file.name().orElse("")
              + " "
              + file.type().map(Object::toString).orElse("")
              + " "
              + file.sizeBytes().orElse(""));
    }

    // [cleanup]
    client.environments.deleteEnvironment(envId);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
  }

  @Test
  public void testApiVersionEnvironmentsEnvironmentFilesPathGetGetFile() throws Exception {
    // [START paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]

    Client client = new Client();

    // [setup]
    CreateEnvironmentRequest createReq =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("main.py")
                        .content("print('Hello')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment created =
        client.environments.createEnvironment(createReq).environment().orElseThrow();
    String envId = created.id().orElseThrow();
    // [/setup]

    com.google.genai.gaos.models.operations.GetEnvironmentFilesResponse response =
        client
            .environments
            .files()
            .list(GetEnvironmentFilesRequest.builder().environment(envId).path("main.py").build());
    GetEnvironmentFilesResponse filesResponse =
        response.getEnvironmentFilesResponse().orElseThrow();
    for (EnvironmentFile file : filesResponse.files().orElse(List.of())) {
      System.out.println(file.name().orElse("") + " " + file.sizeBytes().orElse(""));
    }

    // [cleanup]
    client.environments.deleteEnvironment(envId);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
  }

  @Test
  public void testApiVersionEnvironmentsIdDeleteDelete() throws Exception {
    // [START paths['/{api_version}/environments/{id}'].delete:delete]

    Client client = new Client();

    // [setup]
    CreateEnvironmentRequest request =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("main.py")
                        .content("print('Hello')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment created =
        client.environments.createEnvironment(request).environment().orElseThrow();
    String envId = created.id().orElseThrow();
    // [/setup]

    client.environments.deleteEnvironment(envId);
    System.out.println(
        "com.google.genai.gaos.models.environments.Environment deleted successfully.");
    // [END paths['/{api_version}/environments/{id}'].delete:delete]
  }

  @Test
  public void testApiVersionEnvironmentsIdGetGet() throws Exception {
    // [START paths['/{api_version}/environments/{id}'].get:get]

    Client client = new Client();

    // [setup]
    CreateEnvironmentRequest request =
        CreateEnvironmentRequest.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.INLINE)
                        .target("main.py")
                        .content("print('Hello')")
                        .build()))
            .build();
    com.google.genai.gaos.models.environments.Environment created =
        client.environments.createEnvironment(request).environment().orElseThrow();
    String envId = created.id().orElseThrow();
    // [/setup]

    com.google.genai.gaos.models.environments.Environment environment =
        client.environments.getEnvironment(envId).environment().orElseThrow();
    System.out.println(environment.id().orElse(""));

    // [cleanup]
    client.environments.deleteEnvironment(envId);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{id}'].get:get]
  }

  @Test
  public void testApiVersionInteractionsPostSimple() throws Exception {
    // [START paths['/{api_version}/interactions'].post:simple]

    Client client = new Client();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.of("Hello, how are you?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:simple]
  }

  @Test
  public void testApiVersionInteractionsPostMultiTurn() throws Exception {
    // [START paths['/{api_version}/interactions'].post:multi_turn]

    Client client = new Client();
    List<Step> conversation =
        List.of(
            UserInputStep.builder()
                .content(List.of(TextContent.builder().text("Hello!").build()))
                .build(),
            ModelOutputStep.builder()
                .content(
                    List.of(
                        TextContent.builder().text("Hi there! How can I help you today?").build()))
                .build(),
            UserInputStep.builder()
                .content(
                    List.of(TextContent.builder().text("What is the capital of France?").build()))
                .build());
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.ofStep(conversation))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:multi_turn]
  }

  @Test
  public void testApiVersionInteractionsPostMultimodalImage() throws Exception {
    // [START paths['/{api_version}/interactions'].post:multimodal_image]

    Client client = new Client();
    List<Content> content =
        List.of(
            TextContent.builder().text("What is in this picture?").build(),
            ImageContent.builder()
                .data("BASE64_ENCODED_IMAGE")
                .mimeType(ImageContentMimeType.IMAGE_PNG)
                .build());
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.ofContent(content))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:multimodal_image]
  }

  @Test
  public void testApiVersionInteractionsPostFunctionCalling() throws Exception {
    // [START paths['/{api_version}/interactions'].post:function_calling]

    Client client = new Client();
    Map<String, Object> parameters =
        Map.of(
            "type", "object",
            "properties",
                Map.of(
                    "location",
                    Map.of(
                        "type", "string",
                        "description", "The city and state, e.g. San Francisco, CA")),
            "required", List.of("location"));
    Function functionTool =
        Function.builder()
            .name("get_weather")
            .description("Get the current weather in a given location")
            .parameters(parameters)
            .build();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(functionTool))
            .input(InteractionsInput.of("What is the weather like in Boston, MA?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    List<Step> steps = interaction.steps().orElse(List.of());
    if (!steps.isEmpty()) {
      System.out.println(steps.get(steps.size() - 1));
    }
    // [END paths['/{api_version}/interactions'].post:function_calling]
  }

  @Test
  public void testApiVersionInteractionsPostDeepResearch() throws Exception {
    // [START paths['/{api_version}/interactions'].post:deep_research]

    Client client = new Client();
    CreateAgentInteraction params =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("deep-research-pro-preview-12-2025"))
            .input(InteractionsInput.of("find a cure to cancer"))
            .background(true)
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.status().map(InteractionStatus::value).orElse(""));
    // [END paths['/{api_version}/interactions'].post:deep_research]
  }

  @Test
  public void testApiVersionInteractionsPostAntigravity() throws Exception {
    // [START paths['/{api_version}/interactions'].post:antigravity]

    Client client = new Client();
    CreateAgentInteraction params =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(
                InteractionsInput.of(
                    "Read Hacker News, summarize the top 5 stories, and save results as a markdown"
                        + " file."))
            .environment(CreateAgentInteractionEnvironment.of("remote"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:antigravity]
  }

  @Test
  public void testApiVersionInteractionsPostReuseEnv() throws Exception {
    // [START paths['/{api_version}/interactions'].post:reuse_env]

    Client client = new Client();

    // Step 1: Create an interaction with a fresh remote environment.
    CreateAgentInteraction params1 =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(InteractionsInput.of("Write a hello world script at /workspace/hello.py."))
            .environment(CreateAgentInteractionEnvironment.of("remote"))
            .build();
    CreateInteractionResponse response1 =
        client.interactions.create(CreateInteractionRequestBody.of(params1));
    Interaction interaction1 =
        response1.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println("Environment ID: " + interaction1.environmentId().orElse(""));

    // Step 2: Reuse the same environment in a follow-up interaction.
    CreateAgentInteraction params2 =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(
                InteractionsInput.of(
                    "Modify the script to accept a name argument and greet the user."))
            .environment(
                CreateAgentInteractionEnvironment.of(interaction1.environmentId().orElse("")))
            .previousInteractionId(interaction1.id().orElse(null))
            .build();
    CreateInteractionResponse response2 =
        client.interactions.create(CreateInteractionRequestBody.of(params2));
    Interaction interaction2 =
        response2.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction2.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:reuse_env]
  }

  @Test
  public void testApiVersionInteractionsPostWithSources() throws Exception {
    // [START paths['/{api_version}/interactions'].post:with_sources]

    Client client = new Client();
    com.google.genai.gaos.models.interactions.Environment env =
        com.google.genai.gaos.models.interactions.Environment.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.REPOSITORY)
                        .source("https://github.com/octocat/Spoon-Knife")
                        .target("/workspace/repo")
                        .build(),
                    Source.builder()
                        .type(SourceType.INLINE)
                        .content("Focus on Python files only.")
                        .target("/workspace/notes.txt")
                        .build()))
            .build();
    CreateAgentInteraction params =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(
                InteractionsInput.of(
                    "List all files under /workspace and summarize what you find."))
            .environment(CreateAgentInteractionEnvironment.of(env))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END paths['/{api_version}/interactions'].post:with_sources]
  }

  @Test
  public void testApiVersionInteractionsPostCustomAgent() throws Exception {
    // [START paths['/{api_version}/interactions'].post:custom_agent]

    Client client = new Client();

    // Step 1: Create a custom agent.
    String agentId = "code-reviewer-" + UUID.randomUUID().toString().substring(0, 8);
    com.google.genai.gaos.models.interactions.Environment baseEnv =
        com.google.genai.gaos.models.interactions.Environment.builder()
            .sources(
                List.of(
                    Source.builder()
                        .type(SourceType.REPOSITORY)
                        .source("https://github.com/octocat/Spoon-Knife")
                        .target("/workspace/repo")
                        .build()))
            .build();
    Agent customAgent =
        Agent.builder()
            .id(agentId)
            .baseAgent("antigravity-preview-05-2026")
            .systemInstruction(
                "You are a senior code reviewer. Check every file for bugs, style issues, and"
                    + " security vulnerabilities.")
            .baseEnvironment(BaseEnvironment.of(baseEnv))
            .build();
    client.agents.create(customAgent);

    // Step 2: Use the custom agent.
    CreateAgentInteraction params =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of(agentId))
            .input(
                InteractionsInput.of(
                    "Review the latest changes in /workspace/repo/src and file a summary."))
            .environment(CreateAgentInteractionEnvironment.of("remote"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));

    // [cleanup]
    client.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/interactions'].post:custom_agent]
  }

  @Test
  public void testApiVersionInteractionsIdDeleteDelete() throws Exception {
    // [START paths['/{api_version}/interactions/{id}'].delete:delete]

    Client client = new Client();

    // [setup]
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.of("Hello"))
            .build();
    CreateInteractionResponse created =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    String interactionId = created.interaction().flatMap(Interaction::id).orElseThrow();
    // [/setup]

    client.interactions.delete(interactionId);
    System.out.println("Interaction deleted successfully.");
    // [END paths['/{api_version}/interactions/{id}'].delete:delete]
  }

  @Test
  public void testApiVersionInteractionsIdGetGet() throws Exception {
    // [START paths['/{api_version}/interactions/{id}'].get:get]

    Client client = new Client();

    // [setup]
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.of("Say hello."))
            .build();
    CreateInteractionResponse created =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    String interactionId = created.interaction().flatMap(Interaction::id).orElseThrow();
    // [/setup]

    GetInteractionByIdResponse getResponse =
        client.interactions.get(new GetInteractionByIdRequest(interactionId));
    Interaction interaction =
        getResponse
            .interaction()
            .orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.status().map(InteractionStatus::value).orElse(""));
    // [END paths['/{api_version}/interactions/{id}'].get:get]
  }

  @Test
  public void testApiVersionInteractionsIdCancelPostCancel() throws Exception {
    // [START paths['/{api_version}/interactions/{id}/cancel'].post:cancel]

    Client client = new Client();

    // Start a background interaction so it stays in-progress.
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .input(InteractionsInput.of("Write a long essay about the history of computing."))
            .tools(List.of(new ComputerUse()))
            .background(true)
            .build();
    CreateInteractionResponse created =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    String interactionId = created.interaction().flatMap(Interaction::id).orElseThrow();

    // Cancel the in-progress interaction.
    CancelInteractionByIdResponse cancelResponse = client.interactions.cancel(interactionId);
    Interaction interaction =
        cancelResponse
            .interaction()
            .orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.status().map(InteractionStatus::value).orElse(""));
    // [END paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
  }

  @Test
  public void testSchemaCodeexecutionCodeExecution() throws Exception {
    // [START components.schemas.CodeExecution:code_execution]

    Client client = new Client();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(new CodeExecution()))
            .input(InteractionsInput.of("Calculate the first 10 Fibonacci numbers"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.CodeExecution:code_execution]
  }

  @Test
  public void testSchemaComputeruseComputerUse() throws Exception {
    // [START components.schemas.ComputerUse:computer_use]

    Client client = new Client();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-2.5-computer-use-preview-10-2025")
            .tools(List.of(new ComputerUse()))
            .input(InteractionsInput.of("Find a flight to Tokyo"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.ComputerUse:computer_use]
  }

  @Test
  public void testSchemaFilesearchFileSearch() throws Exception {
    // [START components.schemas.FileSearch:file_search]

    Client client = new Client();

    // Create a file search store so we have a valid one to use.
    FileSearchStore store =
        client.fileSearchStores.create(CreateFileSearchStoreConfig.builder().build());
    String storeName = store.name().orElseThrow();

    FileSearch tool = FileSearch.builder().fileSearchStoreNames(List.of(storeName)).build();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(tool))
            .input(InteractionsInput.of("What documents are available?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));

    // [cleanup]
    client.fileSearchStores.delete(storeName, null);
    // [/cleanup]
    // [END components.schemas.FileSearch:file_search]
  }

  @Test
  public void testSchemaFunctionFunctionCalling() throws Exception {
    // [START components.schemas.Function:function_calling]

    Client client = new Client();
    Map<String, Object> parameters =
        Map.of(
            "type", "object",
            "properties",
                Map.of(
                    "location",
                    Map.of(
                        "type", "string",
                        "description", "The city and state, e.g. San Francisco, CA")),
            "required", List.of("location"));
    Function functionTool =
        Function.builder()
            .name("get_weather")
            .description("Get the current weather in a given location")
            .parameters(parameters)
            .build();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(functionTool))
            .input(InteractionsInput.of("What is the weather like in Boston?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    List<Step> steps = interaction.steps().orElse(List.of());
    if (!steps.isEmpty()) {
      System.out.println(steps.get(steps.size() - 1));
    }
    // [END components.schemas.Function:function_calling]
  }

  @Test
  public void testSchemaGooglemapsGoogleMaps() throws Exception {
    // [START components.schemas.GoogleMaps:google_maps]

    Client client = new Client();
    GoogleMaps tool = GoogleMaps.builder().latitude(37.7749).longitude(-122.4194).build();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(tool))
            .input(InteractionsInput.of("What is the best food near me?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.GoogleMaps:google_maps]
  }

  @Test
  public void testSchemaGooglesearchGoogleSearch() throws Exception {
    // [START components.schemas.GoogleSearch:google_search]

    Client client = new Client();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(new GoogleSearch()))
            .input(InteractionsInput.of("Who is the current president of France?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.GoogleSearch:google_search]
  }

  @Test
  public void testSchemaMcpserverMcpServer() throws Exception {
    // [START components.schemas.McpServer:mcp_server]

    Client client = new Client();
    MCPServer mcpTool =
        MCPServer.builder()
            .name("weather_service")
            .url("https://gemini-api-demos.uc.r.appspot.com/mcp")
            .build();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(mcpTool))
            .input(
                InteractionsInput.of(
                    "Today is 12-05-2025, what is the temperature today in London?"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.McpServer:mcp_server]
  }

  @Test
  public void testSchemaUrlcontextUrlContext() throws Exception {
    // [START components.schemas.UrlContext:url_context]

    Client client = new Client();
    CreateModelInteraction params =
        CreateModelInteraction.builder()
            .model("gemini-3.6-flash")
            .tools(List.of(new URLContext()))
            .input(InteractionsInput.of("Summarize https://www.example.com"))
            .build();
    CreateInteractionResponse response =
        client.interactions.create(CreateInteractionRequestBody.of(params));
    Interaction interaction =
        response.interaction().orElseThrow(() -> new RuntimeException("No interaction returned"));
    System.out.println(interaction.outputText().orElse(""));
    // [END components.schemas.UrlContext:url_context]
  }

  @Test
  public void testApiVersionTriggersPostCreate() throws Exception {
    // [START paths['/{api_version}/triggers'].post:create]

    Client client = new Client();
    CreateAgentInteraction interaction =
        CreateAgentInteraction.builder()
            .agent(AgentOption.of("antigravity-preview-05-2026"))
            .input(InteractionsInput.of("Summarize top news stories."))
            .environment(CreateAgentInteractionEnvironment.of("remote"))
            .build();
    TriggerCreateParams params =
        TriggerCreateParams.builder()
            .schedule("0 9 * * *")
            .timeZone("America/New_York")
            .interaction(interaction)
            .build();
    CreateTriggerResponse response = client.triggers.create(params);
    Trigger trigger =
        response.trigger().orElseThrow(() -> new RuntimeException("No trigger returned"));
    System.out.println(
        trigger.id().orElse("")
            + " "
            + trigger.interaction().flatMap(CreateAgentInteraction::agent).orElse(null));

    // [cleanup]
    trigger.id().ifPresent(client.triggers::delete);
    // [/cleanup]
    // [END paths['/{api_version}/triggers'].post:create]
  }
}
