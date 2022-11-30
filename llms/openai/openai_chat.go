package openai

import (
	"context"

	"github.com/deluan/flowllm"
	"github.com/sashabaranov/go-openai"
)

const defaultChatModel = "gpt-3.5-turbo"

// ChatModel is a LLM implementation that uses the Chat Completions API with the chat style models, like gpt-3.5-turbo and gpt-4.
// It uses a special prompt, Chat, to format the messages as expected by the chat completion API.
// If you use a d