# Gemini API REST examples

This directory contains curl examples for key features of the Gemini API,
organized by feature. They're embedded in the
[Gemini API reference](https://ai.google.dev/api).

Each file is a runnable script. Region tags (`# [START tag]` / `# [END tag]`)
demarcate the example code from the surrounding script. Code within region tags
should be clear, complete, and concise.

Scripts that upload media read it from [`../third_party`](../third_party).

| File | Description |
| ---- | ----------- |
| [cache.sh](./cache.sh) | Context caching |
| [chat.sh](./chat.sh) | Multi-turn chat conversations |
| [code_execution.sh](./code_execution.sh) | Executing code |
| [configure_model_parameters.sh](./configure_model_parameters.sh) | Setting model parameters |
| [controlled_generation.sh](./controlled_generation.sh) | Generating content with output constraints (e.g. JSON mode) |
| [count_tokens.sh](./count_tokens.sh) | Counting input and output tokens |
| [embed.sh](./embed.sh) | Generating embeddings |
| [files.sh](./files.sh) | Managing files with the File API |
| [function_calling.sh](./function_calling.sh) | Using function calling |
| [grounding.sh](./grounding.sh) | Grounding with Google Maps |
| [models.sh](./models.sh) | Listing models and model metadata |
| [safety_settings.sh](./safety_settings.sh) | Setting and using safety controls |
| [system_instruction.sh](./system_instruction.sh) | Setting system instructions |
| [text_generation.sh](./text_generation.sh) | Generating text |
