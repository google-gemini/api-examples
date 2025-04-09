# Java examples

This directory contains examples of working with the Gemini API using the
[Google Gen AI SDK for Java](https://github.com/googleapis/java-genai).

## Install dependencies
open the project in intellliJ IDEA, load the maven build script.

## Run a test file
Before running the tests, set GOOGLE_API_KEY as environment variable for run/debug configuration. 

    mvn -Dtest=<filename> test

For example:

    mvn -Dtest=CodeExecutionTest test

