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
    chain := Chain(
        ParallelChain(
            2,
            Chain(
                Template("What is a good name for a company that makes {product}?"),
                LLM(openai.NewCompletionModel(openai.Options{Model: "text-davinci-003", Temperature: 1})),
                MapOutputTo("name"),
            ),
            Chain(
                ChatTemplate{UserMessage("What is a good slogan for a company that makes {product}?")},
                ChatLLM(openai.NewChatModel(openai.Options{Model: "gpt-3.5-turbo", Temperature: 1})),
                MapOutputTo("slogan"),
            ),
        ),
        // You can modify the LLMs outputs using some string transformation handlers
        TrimSpace("name", "slogan"),
        TrimSuffix(".", "name"),
        Template("The company {name} makes {product} and their slogan is {slogan}."),
    )

    // Run the chain
    res, err := chain(context.Background(), Values{"product": "colorful sockets"})
    fmt.Println(res, err)

    // Output:
    // The company Rainbow Socks Co makes colorful socks and their slogan is "Life is too short for boring socks – let us add some color to your steps!". <nil>
}

```

For more sophisticated usage and features, please inspect the [examples](/examples) folder.

## Installation

To incorporate LangModel-Flow, use the command mentioned below:

```sh
go get -u github.com/KLORNO/langmodel-flow
```

## Features
- Ability to access LLMs and use them effectively
- Tools to construct prompts and parse outputs
- Database integration for smooth storage and retrieval of data
- Inspiration derived from the LangChain project with a keen focus on Go idioms and patterns

## Usage
For examples and detailed instructions on usage, please refer to the [documentation](https://pkg.go.dev/github.com/KLORNO/langmodel-flow) (Work in Progress). You can also check the [examples](/examples) folder.

## Contributing
We are open to contributions from the community! For more information on how you can contribute, kindly refer to our [contributing guidelines](https://github.com/KLORNO/langmodel-flow/blob/main/CONTRIBUTING.md).

## License
LangModel-Flow is licensed under the [MIT License](/LICENSE).
