package main

import (
	"context"
	"fmt"

	. "github.com/deluan/flowllm"
	"github.com/deluan/flowllm/llms/openai"
)

func init() {
	registerExample("simple", "A simple example with only one chain", simple)
}

func simple() {
	// Build a simple chain that will generate a joke about a given topic
	chain := Chain(
		ChatTemplate{UserMessage("Tell me a joke about {topic}?")},
		ChatLLM(openai.NewChatModel(openai.Options{})),
	)

	// Run the chain for