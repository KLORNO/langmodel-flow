package openai

import (
	"context"

	"github.com/deluan/flowllm"
	"github.com/sashabaranov/go-openai"
)

const defaultChatModel = "gpt-3.5-turbo"

// ChatModel is a LLM implementation t