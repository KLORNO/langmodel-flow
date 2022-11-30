package openai

import (
	"context"

	"github.com/deluan/flowllm"
	"github.com/sashabaranov/go-openai"
)

const defaultChatModel = "gpt-3.5-turbo"

// ChatModel is a LLM implementation that uses the Chat Completions API with the chat style models, like gpt-3.5-turbo and gpt-4.
// It uses a spec