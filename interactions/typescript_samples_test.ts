/**
 * @license
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

import {GoogleGenAI} from '@google/genai';

describe('Interactions API TypeScript Samples', () => {
  it('api version agents get list', async () => {
    // [START paths['/{api_version}/agents'].get:list]

    const ai = new GoogleGenAI({});
    const agents = await ai.agents.list();
    for (const agent of agents.agents ?? []) {
      console.log(agent.id);
    }
    // [END paths['/{api_version}/agents'].get:list]
  });

  it('api version agents post create', async () => {
    // [START paths['/{api_version}/agents'].post:create]

    const ai = new GoogleGenAI({});
    const agentId = `research-assistant-${crypto.randomUUID().slice(0, 8)}`;
    const agent = await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      description: 'A helpful research assistant.',
      system_instruction: 'You are a helpful research assistant.',
      base_environment: 'remote',
      tools: [{type: 'google_search'}],
    });
    if (!agent.id) {
      throw new Error('Agent creation failed: ID is undefined');
    }
    console.log(agent.id);

    // [cleanup]
    await ai.agents.delete(agent.id);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:create]
  });

  it('api version agents post with sources', async () => {
    // [START paths['/{api_version}/agents'].post:with_sources]

    const ai = new GoogleGenAI({});
    const agentId = `data-analyst-${crypto.randomUUID().slice(0, 8)}`;
    const agent = await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      system_instruction:
        'You are a data analyst. Always include visualizations and export results as PDF.',
      base_environment: {
        type: 'remote',
        sources: [
          {
            type: 'inline',
            target: '.agents/AGENTS.md',
            content:
              'Always use matplotlib for charts. Include a summary table in every report.',
          },
          {
            type: 'repository',
            source: 'https://github.com/my-org/analysis-templates',
            target: '/workspace/templates',
          },
        ],
      },
    });
    console.log(`Created agent: ${agent.id}`);

    // [cleanup]
    if (agent.id) await ai.agents.delete(agent.id);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:with_sources]
  });

  it('api version agents post fork from env', async () => {
    // [START paths['/{api_version}/agents'].post:fork_from_env]

    const ai = new GoogleGenAI({});

    // Step 1: Set up the environment interactively.
    const interaction = await ai.interactions.create({
      agent: 'antigravity-preview-05-2026',
      input: 'Write a basic Hello World template to /workspace/template.py.',
      environment: 'remote',
    });

    // Step 2: Fork that environment into a named agent.
    const agentId = `my-data-analyst-${crypto.randomUUID().slice(0, 8)}`;
    const agent = await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      system_instruction:
        'You are a data analyst. Use the template at /workspace/template.py for all reports.',
      base_environment: interaction.environment_id,
    });
    console.log(`Forked agent: ${agent.id}`);

    // [cleanup]
    if (agent.id) await ai.agents.delete(agent.id);
    // [/cleanup]
    // [END paths['/{api_version}/agents'].post:fork_from_env]
  });

  it('api version agents id delete delete', async () => {
    // [START paths['/{api_version}/agents/{id}'].delete:delete]

    const ai = new GoogleGenAI({});

    // [setup]
    const agentId = `delete-test-${crypto.randomUUID().slice(0, 8)}`;
    await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      description: 'Temporary agent for deletion.',
      base_environment: 'remote',
    });
    // [/setup]

    await ai.agents.delete(agentId);
    console.log('Agent deleted successfully.');
    // [END paths['/{api_version}/agents/{id}'].delete:delete]
  });

  it('api version agents id get get', async () => {
    // [START paths['/{api_version}/agents/{id}'].get:get]

    const ai = new GoogleGenAI({});

    // [setup]
    const agentId = `test-agent-${crypto.randomUUID().slice(0, 8)}`;
    await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      description: 'A test agent.',
      base_environment: 'remote',
    });
    // [/setup]

    const agent = await ai.agents.get(agentId);
    if (!agent.id) {
      throw new Error('Agent retrieval failed: ID is undefined');
    }
    console.log(agent.id);

    // [cleanup]
    await ai.agents.delete(agent.id);
    // [/cleanup]
    // [END paths['/{api_version}/agents/{id}'].get:get]
  });

  it('api version environments get list', async () => {
    // [START paths['/{api_version}/environments'].get:list]

    const ai = new GoogleGenAI({});
    const response = await ai.environments.list();
    for (const env of response.environments ?? []) {
      console.log(env.id);
    }
    // [END paths['/{api_version}/environments'].get:list]
  });

  it('api version environments post create', async () => {
    // [START paths['/{api_version}/environments'].post:create]

    const ai = new GoogleGenAI({});
    const environment = await ai.environments.create({
      sources: [
        {
          type: 'inline',
          target: 'main.py',
          content: "print('Hello, World!')",
        },
      ],
    });
    if (!environment.id) {
      throw new Error('Environment creation failed: ID is undefined');
    }
    console.log(environment.id);

    // [cleanup]
    await ai.environments.delete(environment.id);
    // [/cleanup]
    // [END paths['/{api_version}/environments'].post:create]
  });

  it('api version environments post copy', async () => {
    // [START paths['/{api_version}/environments'].post:copy]

    const ai = new GoogleGenAI({});

    // [setup]
    const sourceEnv = await ai.environments.create({
      sources: [{type: 'inline', target: 'main.py', content: "print('Hello')"}],
    });
    if (!sourceEnv.id) throw new Error('Failed to create source environment');
    // [/setup]

    const environment = await ai.environments.create({
      from_environment: sourceEnv.id,
    });
    if (!environment.id) {
      throw new Error('Environment creation failed: ID is undefined');
    }
    console.log(environment.id);

    // [cleanup]
    await ai.environments.delete(environment.id);
    await ai.environments.delete(sourceEnv.id);
    // [/cleanup]
    // [END paths['/{api_version}/environments'].post:copy]
  });

  it('api version environments environment files path get list files', async () => {
    // [START paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.environments.create({
      sources: [
        {type: 'inline', target: 'src/main.py', content: "print('Hello')"},
      ],
    });
    if (!created.id) throw new Error('Failed to create environment');
    // [/setup]

    const response = await ai.environments.files.list({
      environment: created.id,
      path: 'src',
    });
    for (const file of response.files ?? []) {
      console.log(file.name, file.type, file.size_bytes);
    }

    // [cleanup]
    await ai.environments.delete(created.id);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
  });

  it('api version environments environment files path get get file', async () => {
    // [START paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.environments.create({
      sources: [{type: 'inline', target: 'main.py', content: "print('Hello')"}],
    });
    if (!created.id) throw new Error('Failed to create environment');
    // [/setup]

    const response = await ai.environments.files.list({
      environment: created.id,
      path: 'main.py',
    });
    for (const file of response.files ?? []) {
      console.log(file.name, file.size_bytes);
    }

    // [cleanup]
    await ai.environments.delete(created.id);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
  });

  it('api version environments id delete delete', async () => {
    // [START paths['/{api_version}/environments/{id}'].delete:delete]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.environments.create({
      sources: [{type: 'inline', target: 'main.py', content: "print('Hello')"}],
    });
    if (!created.id) throw new Error('Failed to create environment');
    // [/setup]

    await ai.environments.delete(created.id);
    console.log('Environment deleted successfully.');
    // [END paths['/{api_version}/environments/{id}'].delete:delete]
  });

  it('api version environments id get get', async () => {
    // [START paths['/{api_version}/environments/{id}'].get:get]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.environments.create({
      sources: [{type: 'inline', target: 'main.py', content: "print('Hello')"}],
    });
    if (!created.id) throw new Error('Failed to create environment');
    // [/setup]

    const environment = await ai.environments.get(created.id);
    console.log(environment.id);

    // [cleanup]
    await ai.environments.delete(created.id);
    // [/cleanup]
    // [END paths['/{api_version}/environments/{id}'].get:get]
  });

  it('api version interactions post simple', async () => {
    // [START paths['/{api_version}/interactions'].post:simple]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: 'Hello, how are you?',
    });
    console.log(interaction.output_text);
    // [END paths['/{api_version}/interactions'].post:simple]
  });

  it('api version interactions post multi turn', async () => {
    // [START paths['/{api_version}/interactions'].post:multi_turn]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: [
        {type: 'user_input', content: [{type: 'text', text: 'Hello'}]},
        {
          type: 'model_output',
          content: [
            {type: 'text', text: 'Hi there! How can I help you today?'},
          ],
        },
        {
          type: 'user_input',
          content: [{type: 'text', text: 'What is the capital of France?'}],
        },
      ],
    });
    console.log(interaction.output_text);
    // [END paths['/{api_version}/interactions'].post:multi_turn]
  });

  it('api version interactions post multimodal image', async () => {
    // [START paths['/{api_version}/interactions'].post:multimodal_image]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: [
        {type: 'text', text: 'What is in this picture?'},
        {type: 'image', data: 'BASE64_ENCODED_IMAGE', mime_type: 'image/png'},
      ],
    });
    console.log(interaction.output_text);
    // [END paths['/{api_version}/interactions'].post:multimodal_image]
  });

  it('api version interactions post function calling', async () => {
    // [START paths['/{api_version}/interactions'].post:function_calling]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [
        {
          type: 'function',
          name: 'get_weather',
          description: 'Get the current weather in a given location',
          parameters: {
            type: 'object',
            properties: {
              location: {
                type: 'string',
                description: 'The city and state, e.g. San Francisco, CA',
              },
            },
            required: ['location'],
          },
        },
      ],
      input: 'What is the weather like in Boston, MA?',
    });
    console.log(interaction.steps.at(-1));
    // [END paths['/{api_version}/interactions'].post:function_calling]
  });

  it('api version interactions post deep research', async () => {
    // [START paths['/{api_version}/interactions'].post:deep_research]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      agent: 'deep-research-pro-preview-12-2025',
      input: 'find a cure to cancer',
      background: true,
    });
    console.log(interaction.status);
    // [END paths['/{api_version}/interactions'].post:deep_research]
  });

  it('api version interactions post antigravity', async () => {
    // [START paths['/{api_version}/interactions'].post:antigravity]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      agent: 'antigravity-preview-05-2026',
      input:
        'Read Hacker News, summarize the top 5 stories, and save results as a markdown file.',
      environment: 'remote',
    });
    console.log(interaction.output_text);
    // [END paths['/{api_version}/interactions'].post:antigravity]
  });

  it('api version interactions post reuse env', async () => {
    // [START paths['/{api_version}/interactions'].post:reuse_env]

    const ai = new GoogleGenAI({});

    // Step 1: Create an interaction with a fresh remote environment.
    const interaction = await ai.interactions.create({
      agent: 'antigravity-preview-05-2026',
      input: 'Write a hello world script at /workspace/hello.py.',
      environment: 'remote',
    });
    console.log(`Environment ID: ${interaction.environment_id}`);

    // Step 2: Reuse the same environment in a follow-up interaction.
    const interaction2 = await ai.interactions.create({
      agent: 'antigravity-preview-05-2026',
      input: 'Modify the script to accept a name argument and greet the user.',
      environment: interaction.environment_id,
      previous_interaction_id: interaction.id,
    });
    console.log(interaction2.output_text);
    // [END paths['/{api_version}/interactions'].post:reuse_env]
  });

  it('api version interactions post with sources', async () => {
    // [START paths['/{api_version}/interactions'].post:with_sources]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      agent: 'antigravity-preview-05-2026',
      input: 'List all files under /workspace and summarize what you find.',
      environment: {
        type: 'remote',
        sources: [
          {
            type: 'repository',
            source: 'https://github.com/octocat/Spoon-Knife',
            target: '/workspace/repo',
          },
          {
            type: 'inline',
            content: 'Focus on Python files only.',
            target: '/workspace/notes.txt',
          },
        ],
      },
    });
    console.log(interaction.output_text);
    // [END paths['/{api_version}/interactions'].post:with_sources]
  });

  it('api version interactions post custom agent', async () => {
    // [START paths['/{api_version}/interactions'].post:custom_agent]

    const ai = new GoogleGenAI({});

    // Step 1: Create a custom agent.
    const agentId = `code-reviewer-${crypto.randomUUID().slice(0, 8)}`;
    await ai.agents.create({
      id: agentId,
      base_agent: 'antigravity-preview-05-2026',
      system_instruction:
        'You are a senior code reviewer. Check every file for bugs, style issues, and security vulnerabilities.',
      base_environment: {
        type: 'remote',
        sources: [
          {
            type: 'repository',
            source: 'https://github.com/octocat/Spoon-Knife',
            target: '/workspace/repo',
          },
        ],
      },
    });

    // Step 2: Use the custom agent.
    const result = await ai.interactions.create({
      agent: agentId,
      input:
        'Review the latest changes in /workspace/repo/src and file a summary.',
      environment: 'remote',
    });
    console.log(result.output_text);

    // [cleanup]
    await ai.agents.delete(agentId);
    // [/cleanup]
    // [END paths['/{api_version}/interactions'].post:custom_agent]
  });

  it('api version interactions id delete delete', async () => {
    // [START paths['/{api_version}/interactions/{id}'].delete:delete]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: 'Hello',
    });
    // [/setup]

    await ai.interactions.delete(created.id);
    console.log('Interaction deleted successfully.');
    // [END paths['/{api_version}/interactions/{id}'].delete:delete]
  });

  it('api version interactions id get get', async () => {
    // [START paths['/{api_version}/interactions/{id}'].get:get]

    const ai = new GoogleGenAI({});

    // [setup]
    const created = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: 'Say hello.',
    });
    // [/setup]

    const interaction = await ai.interactions.get(created.id);
    console.log(interaction.status);
    // [END paths['/{api_version}/interactions/{id}'].get:get]
  });

  it('api version interactions id cancel post cancel', async () => {
    // [START paths['/{api_version}/interactions/{id}/cancel'].post:cancel]

    const ai = new GoogleGenAI({});

    // Start a background interaction so it stays in-progress.
    const created = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      input: 'Write a long essay about the history of computing.',
      tools: [{type: 'computer_use'}],
      background: true,
    });

    // Cancel the in-progress interaction.
    const interaction = await ai.interactions.cancel(created.id);
    console.log(interaction.status);
    // [END paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
  });

  it('schema CodeExecution code execution', async () => {
    // [START components.schemas.CodeExecution:code_execution]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [{type: 'code_execution'}],
      input: 'Calculate the first 10 Fibonacci numbers',
    });
    console.log(interaction.output_text);
    // [END components.schemas.CodeExecution:code_execution]
  });

  it('schema ComputerUse computer use', async () => {
    // [START components.schemas.ComputerUse:computer_use]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-2.5-computer-use-preview-10-2025',
      tools: [{type: 'computer_use'}],
      input: 'Find a flight to Tokyo',
    });
    console.log(interaction.output_text);
    // [END components.schemas.ComputerUse:computer_use]
  });

  it('schema FileSearch file search', async () => {
    // [START components.schemas.FileSearch:file_search]

    const ai = new GoogleGenAI({});

    // Create a file search store so we have a valid one to use.
    const store = await ai.fileSearchStores.create({});
    if (!store.name) {
      throw new Error('Store creation failed: Name is undefined');
    }

    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [
        {
          type: 'file_search',
          file_search_store_names: [store.name],
        },
      ],
      input: 'What documents are available?',
    });
    console.log(interaction.output_text);

    // [cleanup]
    await ai.fileSearchStores.delete({name: store.name});
    // [/cleanup]
    // [END components.schemas.FileSearch:file_search]
  });

  it('schema Function function calling', async () => {
    // [START components.schemas.Function:function_calling]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [
        {
          type: 'function',
          name: 'get_weather',
          description: 'Get the current weather in a given location',
          parameters: {
            type: 'object',
            properties: {
              location: {
                type: 'string',
                description: 'The city and state, e.g. San Francisco, CA',
              },
            },
            required: ['location'],
          },
        },
      ],
      input: 'What is the weather like in Boston?',
    });
    console.log(interaction.steps.at(-1));
    // [END components.schemas.Function:function_calling]
  });

  it('schema GoogleMaps google maps', async () => {
    // [START components.schemas.GoogleMaps:google_maps]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [
        {
          type: 'google_maps',
          latitude: 37.7749,
          longitude: -122.4194,
        },
      ],
      input: 'What is the best food near me?',
    });
    console.log(interaction.output_text);
    // [END components.schemas.GoogleMaps:google_maps]
  });

  it('schema GoogleSearch google search', async () => {
    // [START components.schemas.GoogleSearch:google_search]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [{type: 'google_search'}],
      input: 'Who is the current president of France?',
    });
    console.log(interaction.output_text);
    // [END components.schemas.GoogleSearch:google_search]
  });

  it('schema McpServer mcp server', async () => {
    // [START components.schemas.McpServer:mcp_server]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [
        {
          type: 'mcp_server',
          name: 'weather_service',
          url: 'https://gemini-api-demos.uc.r.appspot.com/mcp',
        },
      ],
      input: 'Today is 12-05-2025, what is the temperature today in London?',
    });
    console.log(interaction.output_text);
    // [END components.schemas.McpServer:mcp_server]
  });

  it('schema UrlContext url context', async () => {
    // [START components.schemas.UrlContext:url_context]

    const ai = new GoogleGenAI({});
    const interaction = await ai.interactions.create({
      model: 'gemini-3.6-flash',
      tools: [{type: 'url_context'}],
      input: 'Summarize https://www.example.com',
    });
    console.log(interaction.output_text);
    // [END components.schemas.UrlContext:url_context]
  });

  it('api version voices get list', async () => {
    // [START paths['/{api_version}/voices'].get:list]

    const ai = new GoogleGenAI({});
    const response = await ai.voices.list();
    for (const voice of response.voices ?? []) {
      console.log(voice.id, voice.display_name);
    }
    // [END paths['/{api_version}/voices'].get:list]
  });

  it('api version voices post create', async () => {
    // [START paths['/{api_version}/voices'].post:create]

    const ai = new GoogleGenAI({});
    const voice = await ai.voices.create({
      store: true,
      voice: {
        type: 'prompted',
        display_name: 'Warm Narrator',
        language_code: 'en-US',
        prompted: {
          input: 'A warm, friendly narrator voice with a calm pace.',
        },
      },
    });
    console.log(voice.id);

    // [cleanup]
    if (voice.id) {
      await ai.voices.delete(voice.id);
    }
    // [/cleanup]
    // [END paths['/{api_version}/voices'].post:create]
  });

  it('api version voices voicesId get get', async () => {
    // [START paths['/{api_version}/voices/{voicesId}'].get:get]

    const ai = new GoogleGenAI({});
    const voice = await ai.voices.get('voice_abc123');
    console.log(voice.id, voice.display_name);
    // [END paths['/{api_version}/voices/{voicesId}'].get:get]
  });

  it('api version voices voicesId delete delete', async () => {
    // [START paths['/{api_version}/voices/{voicesId}'].delete:delete]

    const ai = new GoogleGenAI({});
    await ai.voices.delete('voice_abc123');
    // [END paths['/{api_version}/voices/{voicesId}'].delete:delete]
  });

  it('api version triggers post create', async () => {
    // [START paths['/{api_version}/triggers'].post:create]

    const ai = new GoogleGenAI({});
    const trigger = await ai.triggers.create({
      schedule: '0 9 * * *',
      time_zone: 'America/New_York',
      interaction: {
        agent: 'antigravity-preview-05-2026',
        input: 'Summarize top news stories.',
        environment: 'remote',
      },
    });
    console.log(trigger.id, trigger.interaction.agent);

    // [cleanup]
    if (trigger.id) {
      await ai.triggers.delete(trigger.id);
    }
    // [/cleanup]
    // [END paths['/{api_version}/triggers'].post:create]
  });
});
