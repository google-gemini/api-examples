# Gemini API Examples

This repository contains example code for key features of the Gemini API.
The repo is organized by programming language.

The examples are embedded in the
[Gemini API reference](https://ai.google.dev/api) and other places in the developer documentation.
Every example is:

* **Runnable** — all examples execute as real tests or standalone scripts
* **Minimal & pedagogical** — each file teaches one concept
* **Multi-language** — Go, Java, JavaScript/TypeScript, Python, and REST
* **Documentation-aware** — region tags allow automatic snippet extraction
* **Kept up-to-date** with evolving Gemini features


# **Repository Structure**

```
api-examples/
├── go/                # Go examples + tests
├── java/              # Java examples (Maven project)
├── javascript/        # JS/TS examples + tests
├── python/            # Python runnable examples
├── rest/              # cURL-style shell examples
├── third_party/       # Multimedia files for multimodal examples
├── CONTRIBUTING.md
├── LICENSE
└── README.md          # (this file)
```

Each language folder includes:

* Example files (e.g., `chat.go`, `embed.py`, `text_generation.js`)
* Matching tests (e.g., `chat_test.go`, `chat.test.js`)
* Language-specific README instructions


# **What’s Included**

The repository covers a wide set of Gemini capabilities:

## **Core Text & Chat**

* Text generation
* Thinking mode
* Controlled generation
* Chat sessions
* System instructions
* Safety settings

## **Structured & Programmatic Generation**

* Function calling
* Code execution

## **Data & Tokens**

* Embeddings
* Token counting
* File uploads & retrieval

## **Advanced Features**

* Grounding
* Model parameter configuration
* Caching

## **Multimodal**

* Examples using images, videos, and audio (in `third_party/`)


# **Quickstart by Language**

Below are short, ready-to-run instructions for each language folder.


## **Go**

### Install & Run

```bash
cd go
go mod tidy
export GOOGLE_API_KEY="YOUR_KEY"
go test ./...
```

Each example is a Go test file such as `chat_test.go`, `code_execution_test.go`, etc.


## **Java**

### Set Up

1. Ensure you have a Java JDK (version 11 or higher) and Maven installed
2. Open the `java/` folder in your preferred IDE or command line
3. Run `mvn clean install` to download dependencies

### Set API Key

Export the API key to your environment:

```bash
export GOOGLE_API_KEY="YOUR_KEY"
```

### Run a Test

```bash
mvn -Dtest=<TestClassName> test
```

Example:

```bash
mvn -Dtest=CodeExecutionTest test
```


## **JavaScript / TypeScript**

### Install

```bash
cd javascript
npm install
export GOOGLE_API_KEY="YOUR_KEY"
```

### Run Tests

```bash
npm test
```

To run a specific test file:

```bash
node <feature>.test.js
```

Example:

```bash
node text_generation.test.js
```

### Format Code

```bash
npm run format
```

---

## **Python**

### Install

```bash
cd python
pip install absl-py google-genai Pillow pyink
export GOOGLE_API_KEY="YOUR_KEY"
```

### Run an Example

```bash
python <filename>.py
```

Example:

```bash
python text_generation.py
```

### Format Code

```bash
pyink .
```


## **REST (Shell)**

```bash
cd rest
export GOOGLE_API_KEY="YOUR_KEY"
./embed.sh
```


# **Region Tags**

All examples use **region tags** that enable automated extraction into Gemini documentation.

```python
# [START gemini_text_gen]
response = client.generate_text("Hello!")
# [END gemini_text_gen]
```

When contributing examples:

* Region tags **must** remain correctly paired
* Only example code should be inside the tags
* Test scaffolding should remain outside

This keeps documentation perfectly aligned with working code.


# **Testing Philosophy**

Every example in this repo is **automatically verified**.

* **Go** → `go test`
* **Java** → `mvn test`
* **JavaScript/TypeScript** → Node test files
* **Python** → runnable scripts
* **REST** → executable shell scripts

Tests ensure examples remain:

* Correct
* Up-to-date
* Documentation-ready
* Idiomatic for each language


# **Contributing**

If you're contributing, please ensure that code inside region tags follows best practices:

* **Clear** — easy for beginners to understand
* **Complete** — runnable with minimal setup
* **Concise** — no unnecessary lines

### You can contribute by:

* Adding missing examples
* Enhancing clarity of existing examples
* Expanding test coverage
* Adding new model features
* Improving region tags
* Submitting bug fixes
* Improving language-specific READMEs

### PR Checklist

Before submitting:

* Code runs successfully
* Region tags are correct
* Tests are added or updated
* Code is minimal & idiomatic
* Commit message is descriptive
