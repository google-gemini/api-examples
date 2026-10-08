# Copyright 2026 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Tests and code samples for Gemini Interactions API in Python."""

# pylint: disable=g-import-not-at-top,line-too-long,invalid-name

from absl.testing import absltest
from absl.testing import parameterized


class PythonSamplesTest(parameterized.TestCase):

  def test_api_version_agents_get_list(self):
    # [START paths['/{api_version}/agents'].get:list]
    from google import genai

    client = genai.Client()
    response = client.agents.list()
    for agent in response.agents or []:
      print(agent.id)
    # [END paths['/{api_version}/agents'].get:list]

  def test_api_version_agents_post_create(self):
    # [START paths['/{api_version}/agents'].post:create]
    import uuid
    from google import genai

    client = genai.Client()
    agent = client.agents.create(
        id=f"research-assistant-{uuid.uuid4().hex[:8]}",
        base_agent="antigravity-preview-05-2026",
        description="A helpful research assistant.",
        system_instruction="You are a helpful research assistant.",
        base_environment="remote",
        tools=[{"type": "google_search"}],
    )
    print(agent.id)

    # [cleanup]
    client.agents.delete(agent.id)
    # [/cleanup]
    # [END paths['/{api_version}/agents'].post:create]

  def test_api_version_agents_post_with_sources(self):
    # [START paths['/{api_version}/agents'].post:with_sources]
    import uuid
    from google import genai

    client = genai.Client()
    agent = client.agents.create(
        id=f"data-analyst-{uuid.uuid4().hex[:8]}",
        base_agent="antigravity-preview-05-2026",
        system_instruction=(
            "You are a data analyst. Always include visualizations and export"
            " results as PDF."
        ),
        base_environment={
            "type": "remote",
            "sources": [
                {
                    "type": "inline",
                    "target": ".agents/AGENTS.md",
                    "content": (
                        "Always use matplotlib for charts. Include a summary"
                        " table in every report."
                    ),
                },
                {
                    "type": "repository",
                    "source": "https://github.com/my-org/analysis-templates",
                    "target": "/workspace/templates",
                },
            ],
        },
    )
    print(f"Created agent: {agent.id}")

    # [cleanup]
    client.agents.delete(agent.id)
    # [/cleanup]
    # [END paths['/{api_version}/agents'].post:with_sources]

  def test_api_version_agents_post_fork_from_env(self):
    # [START paths['/{api_version}/agents'].post:fork_from_env]
    import uuid
    from google import genai

    client = genai.Client()

    # Step 1: Set up the environment interactively.
    interaction = client.interactions.create(
        agent="antigravity-preview-05-2026",
        input="Write a basic Hello World template to /workspace/template.py.",
        environment="remote",
    )

    # Step 2: Fork that environment into a named agent.
    agent_id = f"my-data-analyst-{uuid.uuid4().hex[:8]}"
    agent = client.agents.create(
        id=agent_id,
        base_agent="antigravity-preview-05-2026",
        system_instruction=(
            "You are a data analyst. Use the template at /workspace/template.py"
            " for all reports."
        ),
        base_environment=interaction.environment_id,
    )
    print(f"Forked agent: {agent.id}")

    # [cleanup]
    client.agents.delete(agent.id)
    # [/cleanup]
    # [END paths['/{api_version}/agents'].post:fork_from_env]

  def test_api_version_agents_id_delete_delete(self):
    # [START paths['/{api_version}/agents/{id}'].delete:delete]
    import uuid
    from google import genai

    client = genai.Client()

    # [setup]
    agent_id = f"delete-test-{uuid.uuid4().hex[:8]}"
    client.agents.create(
        id=agent_id,
        base_agent="antigravity-preview-05-2026",
        description="Temporary agent for deletion.",
        base_environment="remote",
    )
    # [/setup]

    client.agents.delete(agent_id)
    print("Agent deleted successfully.")
    # [END paths['/{api_version}/agents/{id}'].delete:delete]

  def test_api_version_agents_id_get_get(self):
    # [START paths['/{api_version}/agents/{id}'].get:get]
    import uuid
    from google import genai

    client = genai.Client()

    # [setup]
    agent_id = f"test-agent-{uuid.uuid4().hex[:8]}"
    client.agents.create(
        id=agent_id,
        base_agent="antigravity-preview-05-2026",
        description="A test agent.",
        base_environment="remote",
    )
    # [/setup]

    agent = client.agents.get(agent_id)
    print(agent.id)

    # [cleanup]
    client.agents.delete(agent.id)
    # [/cleanup]
    # [END paths['/{api_version}/agents/{id}'].get:get]

  def test_api_version_environments_get_list(self):
    # [START paths['/{api_version}/environments'].get:list]
    from google import genai

    client = genai.Client()
    response = client.environments.list()
    for environment in response.environments or []:
      print(environment.id)
    # [END paths['/{api_version}/environments'].get:list]

  def test_api_version_environments_post_create(self):
    # [START paths['/{api_version}/environments'].post:create]
    from google import genai

    client = genai.Client()
    environment = client.environments.create(
        sources=[{
            "type": "inline",
            "target": "main.py",
            "content": "print('Hello, World!')",
        }]
    )
    print(environment.id)

    # [cleanup]
    client.environments.delete(id=environment.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments'].post:create]

  def test_api_version_environments_post_copy(self):
    # [START paths['/{api_version}/environments'].post:copy]
    from google import genai

    client = genai.Client()

    # [setup]
    source_env = client.environments.create(
        sources=[{
            "type": "inline",
            "target": "main.py",
            "content": "print('Hello')",
        }]
    )
    # [/setup]

    environment = client.environments.create(
        from_environment=source_env.id,
    )
    print(environment.id)

    # [cleanup]
    client.environments.delete(id=environment.id)
    client.environments.delete(id=source_env.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments'].post:copy]

  def test_api_version_environments_environment_files_path_get_list_files(self):
    # [START paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.environments.create(
        sources=[{
            "type": "inline",
            "target": "src/main.py",
            "content": "print('Hello')",
        }]
    )
    # [/setup]

    response = client.environments.files.list(
        environment=created.id,
        path="src",
    )
    for file in response.files or []:
      print(file.name, file.type, file.size_bytes)

    # [cleanup]
    client.environments.delete(id=created.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]

  def test_api_version_environments_environment_files_path_get_get_file(self):
    # [START paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.environments.create(
        sources=[
            {"type": "inline", "target": "main.py", "content": "print('Hello')"}
        ]
    )
    # [/setup]

    response = client.environments.files.get(
        environment=created.id,
        path="main.py",
    )
    for file in response.files or []:
      print(file.name, file.size_bytes)

    # [cleanup]
    client.environments.delete(id=created.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]

  def test_api_version_environments_environment_files_path_get_download_file(
      self,
  ):
    # [START paths['/{api_version}/environments/{environment}/files/{path}'].get:download_file]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.environments.create(
        sources=[{
            "type": "inline",
            "target": "src/main.py",
            "content": "print('Hello')",
        }]
    )
    # [/setup]

    content = client.environments.files.download(
        environment=created.id,
        path="src/main.py",
    )
    print(content.decode("utf-8"))

    # [cleanup]
    client.environments.delete(id=created.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments/{environment}/files/{path}'].get:download_file]

  def test_api_version_environments_id_delete_delete(self):
    # [START paths['/{api_version}/environments/{id}'].delete:delete]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.environments.create(
        sources=[
            {"type": "inline", "target": "main.py", "content": "print('Hello')"}
        ]
    )
    # [/setup]

    client.environments.delete(id=created.id)
    print("Environment deleted successfully.")
    # [END paths['/{api_version}/environments/{id}'].delete:delete]

  def test_api_version_environments_id_get_get(self):
    # [START paths['/{api_version}/environments/{id}'].get:get]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.environments.create(
        sources=[
            {"type": "inline", "target": "main.py", "content": "print('Hello')"}
        ]
    )
    # [/setup]

    environment = client.environments.get(id=created.id)
    print(environment.id)

    # [cleanup]
    client.environments.delete(id=created.id)
    # [/cleanup]
    # [END paths['/{api_version}/environments/{id}'].get:get]

  def test_api_version_interactions_post_simple(self):
    # [START paths['/{api_version}/interactions'].post:simple]
    from google import genai

    client = genai.Client()
    interaction = client.interactions.create(
        model="gemini-3.6-flash",
        input="Hello, how are you?",
    )
    print(interaction.output_text)
    # [END paths['/{api_version}/interactions'].post:simple]

  def test_api_version_interactions_post_multi_turn(self):
    # [START paths['/{api_version}/interactions'].post:multi_turn]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        input=[
            {
                "type": "user_input",
                "content": [{"type": "text", "text": "Hello!"}],
            },
            {
                "type": "model_output",
                "content": [{
                    "type": "text",
                    "text": "Hi there! How can I help you today?",
                }],
            },
            {
                "type": "user_input",
                "content": [
                    {"type": "text", "text": "What is the capital of France?"}
                ],
            },
        ],
    )
    print(response.output_text)
    # [END paths['/{api_version}/interactions'].post:multi_turn]

  def test_api_version_interactions_post_multimodal_image(self):
    # [START paths['/{api_version}/interactions'].post:multimodal_image]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        input=[
            {"type": "text", "text": "What is in this picture?"},
            {
                "type": "image",
                "data": "BASE64_ENCODED_IMAGE",
                "mime_type": "image/png",
            },
        ],
    )
    print(response.output_text)
    # [END paths['/{api_version}/interactions'].post:multimodal_image]

  def test_api_version_interactions_post_function_calling(self):
    # [START paths['/{api_version}/interactions'].post:function_calling]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{
            "type": "function",
            "name": "get_weather",
            "description": "Get the current weather in a given location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": (
                            "The city and state, e.g. San Francisco, CA"
                        ),
                    }
                },
                "required": ["location"],
            },
        }],
        input="What is the weather like in Boston, MA?",
    )
    print(response.steps[-1])
    # [END paths['/{api_version}/interactions'].post:function_calling]

  def test_api_version_interactions_post_deep_research(self):
    # [START paths['/{api_version}/interactions'].post:deep_research]
    from google import genai

    client = genai.Client()
    interaction = client.interactions.create(
        agent="deep-research-pro-preview-12-2025",
        input="find a cure to cancer",
        background=True,
    )
    print(interaction.status)
    # [END paths['/{api_version}/interactions'].post:deep_research]

  def test_api_version_interactions_post_antigravity(self):
    # [START paths['/{api_version}/interactions'].post:antigravity]
    from google import genai

    client = genai.Client()
    interaction = client.interactions.create(
        agent="antigravity-preview-05-2026",
        input=(
            "Read Hacker News, summarize the top 5 stories, and save results as"
            " a markdown file."
        ),
        environment="remote",
    )
    print(interaction.output_text)
    # [END paths['/{api_version}/interactions'].post:antigravity]

  def test_api_version_interactions_post_reuse_env(self):
    # [START paths['/{api_version}/interactions'].post:reuse_env]
    from google import genai

    client = genai.Client()

    # Step 1: Create an interaction with a fresh remote environment.
    interaction = client.interactions.create(
        agent="antigravity-preview-05-2026",
        input="Write a hello world script at /workspace/hello.py.",
        environment="remote",
    )
    print(f"Environment ID: {interaction.environment_id}")

    # Step 2: Reuse the same environment in a follow-up interaction.
    interaction_2 = client.interactions.create(
        agent="antigravity-preview-05-2026",
        input="Modify the script to accept a name argument and greet the user.",
        environment=interaction.environment_id,
        previous_interaction_id=interaction.id,
    )
    print(interaction_2.output_text)
    # [END paths['/{api_version}/interactions'].post:reuse_env]

  def test_api_version_interactions_post_with_sources(self):
    # [START paths['/{api_version}/interactions'].post:with_sources]
    from google import genai

    client = genai.Client()
    interaction = client.interactions.create(
        agent="antigravity-preview-05-2026",
        input="List all files under /workspace and summarize what you find.",
        environment={
            "type": "remote",
            "sources": [
                {
                    "type": "repository",
                    "source": "https://github.com/octocat/Spoon-Knife",
                    "target": "/workspace/repo",
                },
                {
                    "type": "inline",
                    "content": "Focus on Python files only.",
                    "target": "/workspace/notes.txt",
                },
            ],
        },
    )
    print(interaction.output_text)
    # [END paths['/{api_version}/interactions'].post:with_sources]

  def test_api_version_interactions_post_custom_agent(self):
    # [START paths['/{api_version}/interactions'].post:custom_agent]
    import uuid
    from google import genai

    client = genai.Client()

    # Step 1: Create a custom agent.
    agent_id = f"code-reviewer-{uuid.uuid4().hex[:8]}"
    client.agents.create(
        id=agent_id,
        base_agent="antigravity-preview-05-2026",
        system_instruction=(
            "You are a senior code reviewer. Check every file for bugs, style"
            " issues, and security vulnerabilities."
        ),
        base_environment={
            "type": "remote",
            "sources": [{
                "type": "repository",
                "source": "https://github.com/octocat/Spoon-Knife",
                "target": "/workspace/repo",
            }],
        },
    )

    # Step 2: Use the custom agent.
    result = client.interactions.create(
        agent=agent_id,
        input=(
            "Review the latest changes in /workspace/repo/src and file a"
            " summary."
        ),
        environment="remote",
    )
    print(result.output_text)

    # [cleanup]
    client.agents.delete(agent_id)
    # [/cleanup]
    # [END paths['/{api_version}/interactions'].post:custom_agent]

  def test_api_version_interactions_id_delete_delete(self):
    # [START paths['/{api_version}/interactions/{id}'].delete:delete]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.interactions.create(
        model="gemini-3.6-flash",
        input="Hello",
    )
    # [/setup]

    client.interactions.delete(id=created.id)
    print("Interaction deleted successfully.")
    # [END paths['/{api_version}/interactions/{id}'].delete:delete]

  def test_api_version_interactions_id_get_get(self):
    # [START paths['/{api_version}/interactions/{id}'].get:get]
    from google import genai

    client = genai.Client()

    # [setup]
    created = client.interactions.create(
        model="gemini-3.6-flash", input="Say hello."
    )
    # [/setup]

    interaction = client.interactions.get(id=created.id)
    print(interaction.status)
    # [END paths['/{api_version}/interactions/{id}'].get:get]

  def test_api_version_interactions_id_cancel_post_cancel(self):
    # [START paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
    from google import genai

    client = genai.Client()

    # Start a background interaction so it stays in-progress.
    created = client.interactions.create(
        model="gemini-3.6-flash",
        input="Write a long essay about the history of computing.",
        tools=[{"type": "computer_use"}],
        background=True,
    )

    # Cancel the in-progress interaction.
    interaction = client.interactions.cancel(id=created.id)
    print(interaction.status)
    # [END paths['/{api_version}/interactions/{id}/cancel'].post:cancel]

  def test_schema_CodeExecution_code_execution(self):
    # [START components.schemas.CodeExecution:code_execution]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{"type": "code_execution"}],
        input="Calculate the first 10 Fibonacci numbers",
    )
    print(response.output_text)
    # [END components.schemas.CodeExecution:code_execution]

  def test_schema_ComputerUse_computer_use(self):
    # [START components.schemas.ComputerUse:computer_use]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-2.5-computer-use-preview-10-2025",
        tools=[{"type": "computer_use"}],
        input="Find a flight to Tokyo",
    )
    print(response.output_text)
    # [END components.schemas.ComputerUse:computer_use]

  def test_schema_FileSearch_file_search(self):
    # [START components.schemas.FileSearch:file_search]
    from google import genai

    client = genai.Client()

    # Create a file search store so we have a valid one to use.
    store = client.file_search_stores.create()

    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[
            {"type": "file_search", "file_search_store_names": [store.name]}
        ],
        input="What documents are available?",
    )
    print(response.output_text)

    # [cleanup]
    client.file_search_stores.delete(name=store.name)
    # [/cleanup]
    # [END components.schemas.FileSearch:file_search]

  def test_schema_Function_function_calling(self):
    # [START components.schemas.Function:function_calling]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{
            "type": "function",
            "name": "get_weather",
            "description": "Get the current weather in a given location",
            "parameters": {
                "type": "object",
                "properties": {
                    "location": {
                        "type": "string",
                        "description": (
                            "The city and state, e.g. San Francisco, CA"
                        ),
                    }
                },
                "required": ["location"],
            },
        }],
        input="What is the weather like in Boston?",
    )
    print(response.steps[-1])
    # [END components.schemas.Function:function_calling]

  def test_schema_GoogleMaps_google_maps(self):
    # [START components.schemas.GoogleMaps:google_maps]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[
            {"type": "google_maps", "latitude": 37.7749, "longitude": -122.4194}
        ],
        input="What is the best food near me?",
    )
    print(response.output_text)
    # [END components.schemas.GoogleMaps:google_maps]

  def test_schema_GoogleSearch_google_search(self):
    # [START components.schemas.GoogleSearch:google_search]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{"type": "google_search"}],
        input="Who is the current president of France?",
    )
    print(response.output_text)
    # [END components.schemas.GoogleSearch:google_search]

  def test_schema_McpServer_mcp_server(self):
    # [START components.schemas.McpServer:mcp_server]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{
            "type": "mcp_server",
            "name": "weather_service",
            "url": "https://gemini-api-demos.uc.r.appspot.com/mcp",
        }],
        input="Today is 12-05-2025, what is the temperature today in London?",
    )
    print(response.output_text)
    # [END components.schemas.McpServer:mcp_server]

  def test_schema_UrlContext_url_context(self):
    # [START components.schemas.UrlContext:url_context]
    from google import genai

    client = genai.Client()
    response = client.interactions.create(
        model="gemini-3.6-flash",
        tools=[{"type": "url_context"}],
        input="Summarize https://www.example.com",
    )
    print(response.output_text)
    # [END components.schemas.UrlContext:url_context]

  def test_api_version_voices_get_list(self):
    # [START paths['/{api_version}/voices'].get:list]
    from google import genai

    client = genai.Client()
    response = client.voices.list()
    for voice in response.voices or []:
      print(voice.id, voice.display_name)
    # [END paths['/{api_version}/voices'].get:list]

  def test_api_version_voices_post_create(self):
    # [START paths['/{api_version}/voices'].post:create]
    from google import genai

    client = genai.Client()
    voice = client.voices.create(
        store=True,
        voice={
            "type": "prompted",
            "display_name": "Warm Narrator",
            "language_code": "en-US",
            "prompted": {
                "input": "A warm, friendly narrator voice with a calm pace.",
            },
        },
    )
    print(voice.id)

    # [cleanup]
    if voice.id:
      client.voices.delete(id=voice.id)
    # [/cleanup]
    # [END paths['/{api_version}/voices'].post:create]

  def test_api_version_voices_voicesid_get_get(self):
    # [START paths['/{api_version}/voices/{voicesId}'].get:get]
    from google import genai

    client = genai.Client()
    voice = client.voices.get(id="voice_abc123")
    print(voice.id, voice.display_name)
    # [END paths['/{api_version}/voices/{voicesId}'].get:get]

  def test_api_version_voices_voicesid_delete_delete(self):
    # [START paths['/{api_version}/voices/{voicesId}'].delete:delete]
    from google import genai

    client = genai.Client()
    client.voices.delete(id="voice_abc123")
    # [END paths['/{api_version}/voices/{voicesId}'].delete:delete]

  def test_api_version_triggers_post_create(self):
    # [START paths['/{api_version}/triggers'].post:create]
    from google import genai

    client = genai.Client()
    trigger = client.triggers.create(
        schedule="0 9 * * *",
        time_zone="America/New_York",
        interaction={
            "agent": "antigravity-preview-05-2026",
            "input": "Summarize top news stories.",
            "environment": "remote",
        },
    )
    print(trigger.id, trigger.interaction.agent)

    # [cleanup]
    if trigger.id:
      client.triggers.delete(id=trigger.id)
    # [/cleanup]
    # [END paths['/{api_version}/triggers'].post:create]


if __name__ == "__main__":
  absltest.main()
