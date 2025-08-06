# Gemini API examples

This repository contains example code for key features of the Gemini API. The
repo is organized by programming language.

The examples are embedded in the
[Gemini API reference](https://ai.google.dev/api) and other places in the
developer documentation.

Each file is structured as a runnable test case, ensuring that examples are
executable and functional. Each test demonstrates a single concept and contains
special comments called _region tags_ that demarcate the test scaffolding from
the example code snippet. Here's a Python example:

```python
def test_text_gen_text_only_prompt(self):
    # [START text_gen_text_only_prompt]
    from google import genai

    client = genai.Client()
    response = client.models.generate_content(
        model="gemini-2.5-flash", contents="Write a story about a magic backpack."
    )
    print(response.text)
    # [END text_gen_text_only_prompt]
```

The API reference can be configured to show the code between the region tags.

If you're contributing, please make sure that the code within region tags
follows best practices for example code: clear, complete, and concise.