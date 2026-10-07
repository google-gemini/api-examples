#!/bin/bash
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

set -e

sample_upload_api_version_environments_environment_files_path_put_start_upload() {
# [START paths['/upload/{api_version}/environments/{environment}/files/{path}'].put:start_upload]
curl -i -X PUT \
  'https://generativelanguage.googleapis.com/upload/v1beta/environments/env_abc123/files/main.py?overwrite=true' \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H 'X-Goog-Upload-Protocol: resumable' \
  -H 'X-Goog-Upload-Command: start' \
  -H "X-Goog-Upload-Header-Content-Length: $(wc -c < main.py)" \
  -H 'X-Goog-Upload-Header-Content-Type: text/x-python'
# [END paths['/upload/{api_version}/environments/{environment}/files/{path}'].put:start_upload]
}

sample_api_version_agents_get_list() {
# [START paths['/{api_version}/agents'].get:list]
curl -X GET https://generativelanguage.googleapis.com/v1beta/agents \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/agents'].get:list]
}

sample_api_version_agents_post_create() {
# [START paths['/{api_version}/agents'].post:create]
curl -X POST https://generativelanguage.googleapis.com/v1beta/agents \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "research-assistant-abc123",
    "base_agent": "antigravity-preview-05-2026",
    "description": "A helpful research assistant.",
    "system_instruction": "You are a helpful research assistant.",
    "base_environment": "remote",
    "tools": [{"type": "google_search"}]
  }'
# [END paths['/{api_version}/agents'].post:create]
}

sample_api_version_agents_post_with_sources() {
# [START paths['/{api_version}/agents'].post:with_sources]
curl -X POST https://generativelanguage.googleapis.com/v1beta/agents \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "data-analyst-abc123",
    "base_agent": "antigravity-preview-05-2026",
    "system_instruction": "You are a data analyst. Always include visualizations and export results as PDF.",
    "base_environment": {
      "type": "remote",
      "sources": [
        {
          "type": "inline",
          "target": ".agents/AGENTS.md",
          "content": "Always use matplotlib for charts. Include a summary table in every report."
        },
        {
          "type": "repository",
          "source": "https://github.com/my-org/analysis-templates",
          "target": "/workspace/templates"
        }
      ]
    }
  }'
# [END paths['/{api_version}/agents'].post:with_sources]
}

sample_api_version_agents_post_fork_from_env() {
# [START paths['/{api_version}/agents'].post:fork_from_env]
# Step 1: Set up the environment interactively
RESPONSE=$(curl -s -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "agent": "antigravity-preview-05-2026",
    "input": "Write a basic Hello World template to /workspace/template.py.",
    "environment": "remote"
  }')
ENV_ID=$(echo "$RESPONSE" | python3 -c "import sys,json; print(json.load(sys.stdin)['environment_id'])")

# Step 2: Fork that environment into a named agent
curl -X POST https://generativelanguage.googleapis.com/v1beta/agents \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "{
    \"id\": \"my-data-analyst\",
    \"base_agent\": \"antigravity-preview-05-2026\",
    \"system_instruction\": \"You are a data analyst. Use the template at /workspace/template.py for all reports.\",
    \"base_environment\": \"$ENV_ID\"
  }"
# [END paths['/{api_version}/agents'].post:fork_from_env]
}

sample_api_version_agents_id_delete_delete() {
# [START paths['/{api_version}/agents/{id}'].delete:delete]
curl -X DELETE https://generativelanguage.googleapis.com/v1beta/agents/ag_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/agents/{id}'].delete:delete]
}

sample_api_version_agents_id_get_get() {
# [START paths['/{api_version}/agents/{id}'].get:get]
curl -X GET https://generativelanguage.googleapis.com/v1beta/agents/ag_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/agents/{id}'].get:get]
}

sample_api_version_environments_get_list() {
# [START paths['/{api_version}/environments'].get:list]
curl -X GET https://generativelanguage.googleapis.com/v1beta/environments \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/environments'].get:list]
}

sample_api_version_environments_post_create() {
# [START paths['/{api_version}/environments'].post:create]
curl -X POST https://generativelanguage.googleapis.com/v1beta/environments \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "sources": [{
      "type": "inline",
      "target": "main.py",
      "content": "print(\"Hello, World!\")"
    }]
  }'
# [END paths['/{api_version}/environments'].post:create]
}

sample_api_version_environments_post_copy() {
# [START paths['/{api_version}/environments'].post:copy]
curl -X POST https://generativelanguage.googleapis.com/v1beta/environments \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "from_environment": "environments/env_abc123"
  }'
# [END paths['/{api_version}/environments'].post:copy]
}

sample_api_version_environments_environment_files_path_get_list_files() {
# [START paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
curl -X GET 'https://generativelanguage.googleapis.com/v1beta/environments/env_abc123/files/src' \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/environments/{environment}/files/{path}'].get:list_files]
}

sample_api_version_environments_environment_files_path_get_get_file() {
# [START paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
curl -X GET 'https://generativelanguage.googleapis.com/v1beta/environments/env_abc123/files/main.py' \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/environments/{environment}/files/{path}'].get:get_file]
}

sample_api_version_environments_environment_files_path_get_download_file() {
# [START paths['/{api_version}/environments/{environment}/files/{path}'].get:download_file]
curl -X GET 'https://generativelanguage.googleapis.com/v1beta/environments/env_abc123/files/src/main.py?alt=media' \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  --output main.py
# [END paths['/{api_version}/environments/{environment}/files/{path}'].get:download_file]
}

sample_api_version_environments_id_delete_delete() {
# [START paths['/{api_version}/environments/{id}'].delete:delete]
curl -X DELETE https://generativelanguage.googleapis.com/v1beta/environments/env_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/environments/{id}'].delete:delete]
}

sample_api_version_environments_id_get_get() {
# [START paths['/{api_version}/environments/{id}'].get:get]
curl -X GET https://generativelanguage.googleapis.com/v1beta/environments/env_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/environments/{id}'].get:get]
}

sample_api_version_interactions_post_simple() {
# [START paths['/{api_version}/interactions'].post:simple]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "input": "Hello, how are you?"
  }'
# [END paths['/{api_version}/interactions'].post:simple]
}

sample_api_version_interactions_post_multi_turn() {
# [START paths['/{api_version}/interactions'].post:multi_turn]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "input": [
      { "type": "user_input", "content": [{ "type": "text", "text": "Hello!" }] },
      { "type": "model_output", "content": [{ "type": "text", "text": "Hi there! How can I help you today?" }] },
      { "type": "user_input", "content": [{ "type": "text", "text": "What is the capital of France?" }] }
    ]
  }'
# [END paths['/{api_version}/interactions'].post:multi_turn]
}

sample_api_version_interactions_post_multimodal_image() {
# [START paths['/{api_version}/interactions'].post:multimodal_image]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "input": [
      {
        "type": "text",
        "text": "What is in this picture?"
      },
      {
        "type": "image",
        "data": "BASE64_ENCODED_IMAGE",
        "mime_type": "image/png"
      }
    ]
  }'
# [END paths['/{api_version}/interactions'].post:multimodal_image]
}

sample_api_version_interactions_post_function_calling() {
# [START paths['/{api_version}/interactions'].post:function_calling]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [
      {
        "type": "function",
        "name": "get_weather",
        "description": "Get the current weather in a given location",
        "parameters": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "The city and state, e.g. San Francisco, CA"
            }
          },
          "required": [
            "location"
          ]
        }
      }
    ],
    "input": "What is the weather like in Boston, MA?"
  }'
# [END paths['/{api_version}/interactions'].post:function_calling]
}

sample_api_version_interactions_post_deep_research() {
# [START paths['/{api_version}/interactions'].post:deep_research]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "agent": "deep-research-pro-preview-12-2025",
    "input": "Find a cure to cancer",
    "background": true
  }'
# [END paths['/{api_version}/interactions'].post:deep_research]
}

sample_api_version_interactions_post_antigravity() {
# [START paths['/{api_version}/interactions'].post:antigravity]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "agent": "antigravity-preview-05-2026",
    "input": "Read Hacker News, summarize the top 5 stories, and save results as a markdown file.",
    "environment": "remote"
  }'
# [END paths['/{api_version}/interactions'].post:antigravity]
}

sample_api_version_interactions_post_reuse_env() {
# [START paths['/{api_version}/interactions'].post:reuse_env]
# Step 1: Create an interaction with a fresh remote environment.
RESPONSE=$(curl -s -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -H "Api-Revision: 2026-05-20" \
  -d '{
    "agent": "antigravity-preview-05-2026",
    "input": "Write a hello world script at /workspace/hello.py.",
    "environment": "remote"
  }')
INTERACTION_ID=$(echo "$RESPONSE" | python3 -c "import sys,json; print(json.load(sys.stdin)['id'])")
ENV_ID=$(echo "$RESPONSE" | python3 -c "import sys,json; print(json.load(sys.stdin)['environment_id'])")

# Step 2: Reuse the same environment in a follow-up interaction.
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d "{
    \"agent\": \"antigravity-preview-05-2026\",
    \"input\": \"Modify the script to accept a name argument and greet the user.\",
    \"environment\": \"$ENV_ID\",
    \"previous_interaction_id\": \"$INTERACTION_ID\"
  }"
# [END paths['/{api_version}/interactions'].post:reuse_env]
}

sample_api_version_interactions_post_with_sources() {
# [START paths['/{api_version}/interactions'].post:with_sources]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "agent": "antigravity-preview-05-2026",
    "input": "List all files under /workspace and summarize what you find.",
    "environment": {
      "type": "remote",
      "sources": [
        {
          "type": "repository",
          "source": "https://github.com/octocat/Spoon-Knife",
          "target": "/workspace/repo"
        },
        {
          "type": "inline",
          "content": "Focus on Python files only.",
          "target": "/workspace/notes.txt"
        }
      ]
    }
  }'
# [END paths['/{api_version}/interactions'].post:with_sources]
}

sample_api_version_interactions_post_custom_agent() {
# [START paths['/{api_version}/interactions'].post:custom_agent]
# Step 1: Create a custom agent.
curl -X POST https://generativelanguage.googleapis.com/v1beta/agents \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "code-reviewer",
    "base_agent": "antigravity-preview-05-2026",
    "system_instruction": "You are a senior code reviewer. Check every file for bugs, style issues, and security vulnerabilities.",
    "base_environment": {
      "type": "remote",
      "sources": [{
        "type": "repository",
        "source": "https://github.com/octocat/Spoon-Knife",
        "target": "/workspace/repo"
      }]
    }
  }'

# Step 2: Use the custom agent.
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "agent": "code-reviewer",
    "input": "Review the latest changes in /workspace/repo/src and file a summary.",
    "environment": "remote"
  }'
# [END paths['/{api_version}/interactions'].post:custom_agent]
}

sample_api_version_interactions_id_delete_delete() {
# [START paths['/{api_version}/interactions/{id}'].delete:delete]
# [setup]
INTERACTION_ID=$(curl -s -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \ -d '{"model": "gemini-3.6-flash", "input": "Hello"}' \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['id'])")
# [/setup]

curl -X DELETE "https://generativelanguage.googleapis.com/v1beta/interactions/$INTERACTION_ID" \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/interactions/{id}'].delete:delete]
}

sample_api_version_interactions_id_get_get() {
# [START paths['/{api_version}/interactions/{id}'].get:get]
# [setup]
INTERACTION_ID=$(curl -s -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -H "Api-Revision: 2026-05-20" \
  -d '{"model": "gemini-3.6-flash", "input": "Say hello."}' \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['id'])")
# [/setup]

curl -X GET "https://generativelanguage.googleapis.com/v1beta/interactions/$INTERACTION_ID" \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Api-Revision: 2026-05-20"
# [END paths['/{api_version}/interactions/{id}'].get:get]
}

sample_api_version_interactions_id_cancel_post_cancel() {
# [START paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
# [setup]
INTERACTION_ID=$(curl -s -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \ -d '{"model": "gemini-3.6-flash", "input": "Write a long essay about the history of computing.", "background": true}' \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['id'])")
# [/setup]

curl -X POST "https://generativelanguage.googleapis.com/v1beta/interactions/$INTERACTION_ID/cancel" \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/interactions/{id}/cancel'].post:cancel]
}

sample_schema_CodeExecution_code_execution() {
# [START components.schemas.CodeExecution:code_execution]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "code_execution"
    }],
    "input": "Calculate the first 10 Fibonacci numbers"
  }'
# [END components.schemas.CodeExecution:code_execution]
}

sample_schema_ComputerUse_computer_use() {
# [START components.schemas.ComputerUse:computer_use]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-2.5-computer-use-preview-10-2025",
    "tools": [{
      "type": "computer_use"
    }],
    "input": "Find a flight to Tokyo"
  }'
# [END components.schemas.ComputerUse:computer_use]
}

sample_schema_FileSearch_file_search() {
# [START components.schemas.FileSearch:file_search]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "file_search",
      "file_search_store_names": ["fileSearchStores/m64d1sevsr4y-xfyawui3fxqg"]
    }],
    "input": "Who is the author of the book?"
  }'
# [END components.schemas.FileSearch:file_search]
}

sample_schema_Function_function_calling() {
# [START components.schemas.Function:function_calling]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "function",
      "name": "get_weather",
      "description": "Get the current weather in a given location",
      "parameters": {
        "type": "object",
        "properties": {
          "location": {
            "type": "string",
            "description": "The city and state, e.g. San Francisco, CA"
          }
        },
        "required": ["location"]
      }
    }],
    "input": "What is the weather like in Boston, MA?"
  }'
# [END components.schemas.Function:function_calling]
}

sample_schema_GoogleMaps_google_maps() {
# [START components.schemas.GoogleMaps:google_maps]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "google_maps",
      "latitude": 37.7749,
      "longitude": -122.4194
    }],
    "input": "What is the best food near me?"
  }'
# [END components.schemas.GoogleMaps:google_maps]
}

sample_schema_GoogleSearch_google_search() {
# [START components.schemas.GoogleSearch:google_search]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "google_search"
    }],
    "input": "Who is the current president of France?"
  }'
# [END components.schemas.GoogleSearch:google_search]
}

sample_schema_McpServer_mcp_server() {
# [START components.schemas.McpServer:mcp_server]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "mcp_server",
      "name": "weather_service",
      "url": "https://gemini-api-demos.uc.r.appspot.com/mcp"
    }],
    "input": "Today is 12-05-2025, what is the temperature today in London?"
  }'
# [END components.schemas.McpServer:mcp_server]
}

sample_schema_UrlContext_url_context() {
# [START components.schemas.UrlContext:url_context]
curl -X POST https://generativelanguage.googleapis.com/v1beta/interactions \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "gemini-3.6-flash",
    "tools": [{
      "type": "url_context"
    }],
    "input": "Summarize https://www.example.com"
  }'
# [END components.schemas.UrlContext:url_context]
}

sample_api_version_voices_get_list() {
# [START paths['/{api_version}/voices'].get:list]
curl -X GET https://generativelanguage.googleapis.com/v1beta/voices \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/voices'].get:list]
}

sample_api_version_voices_post_create() {
# [START paths['/{api_version}/voices'].post:create]
curl -X POST https://generativelanguage.googleapis.com/v1beta/voices \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "store": true,
    "voice": {
      "type": "prompted",
      "display_name": "Warm Narrator",
      "language_code": "en-US",
      "prompted": {
        "input": "A warm, friendly narrator voice with a calm pace."
      }
    }
  }'
# [END paths['/{api_version}/voices'].post:create]
}

sample_api_version_voices_voicesid_get_get() {
# [START paths['/{api_version}/voices/{voicesId}'].get:get]
curl -X GET https://generativelanguage.googleapis.com/v1beta/voices/voice_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/voices/{voicesId}'].get:get]
}

sample_api_version_voices_voicesid_delete_delete() {
# [START paths['/{api_version}/voices/{voicesId}'].delete:delete]
curl -X DELETE https://generativelanguage.googleapis.com/v1beta/voices/voice_abc123 \
  -H "x-goog-api-key: $GEMINI_API_KEY"
# [END paths['/{api_version}/voices/{voicesId}'].delete:delete]
}

sample_api_version_triggers_post_create() {
# [START paths['/{api_version}/triggers'].post:create]
curl -X POST https://generativelanguage.googleapis.com/v1beta/triggers \
  -H "x-goog-api-key: $GEMINI_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "schedule": "0 9 * * *",
    "time_zone": "America/New_York",
    "interaction": {
      "agent": "antigravity-preview-05-2026",
      "input": "Summarize top news stories.",
      "environment": "remote"
    }
  }'
# [END paths['/{api_version}/triggers'].post:create]
}
