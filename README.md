# LangModel-Flow

[![Build](https://img.shields.io/github/actions/workflow/status/KLORNO/langmodel-flow/go.yml?branch=main&logo=github)](https://github.com/KLORNO/langmodel-flow/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/KLORNO/langmodel-flow)](https://goreportcard.com/report/github.com/KLORNO/langmodel-flow)
[![GoDoc](https://pkg.go.dev/badge/github.com/KLORNO/langmodel-flow)](https://pkg.go.dev/github.com/KLORNO/langmodel-flow)
[![License](https://img.shields.io/github/license/KLORNO/langmodel-flow)](/LICENSE)

> **NOTICE**: This is still a work-in-progress. The interfaces provided by this project are still subject to change.

LangModel-Flow is a Go-based framework for building applications drawing on the strength of language models. It leverages the patterns of composability and chain of responsibility, and provides tools for accessing Language Model Learning (LLMs), building prompts, and chaining calls together. It even encompasses parsers and database integrations, making it ideal for developers working with language models.

LangModel-Flow draws significant inspiration from the [LangChain](https://docs.langchain.com/docs) project.

## Usage

In the example below, we employ LangModel-Flow to construct a basic chain which generates a company name and slogan based on a product name. The application exploits two distinct LLMs, one for the company name and the other for the slogan. The LLMs are called in parallel, and the results are subsequently amalgamated into a single output.

```go
package main

import (
    "context"
    "fmt"

    . "github.com/KLORNO/langmodel-flow"
    "github.com/KLORNO/langmodel-flow/llms/openai"
)

func main() {
    // Build a chain that will generate a company name and slogan. Calls to the OpenAI API are made in parallel, and the 
    // results are merged into a single result.
   