package openai

import (
	"context"

	"github.com/deluan/flowllm"
	"github.com/sashabaranov/go-openai"
)

const defaultChatModel = "gpt-3.5-turbo"

// ChatModel is a LLM implementation that uses the Chat Completions API with the chat style models, like gpt-3.5-turbo and gpt-4.
// It uses a special prompt, Chat, to format the messages as expected by the chat completion API.
// If you use a different prompt, it will be wrapped in a Chat with a single user message.
type ChatModel struct {
	*CompletionModel
}

func NewChatModel(opts Options) *ChatModel {
	if opts.Model == "" {
		opts.Model = defaultChatModel
	}
	llm := NewCompletionModel(opts)
	return &ChatModel{CompletionModel: llm}
}

func (m *ChatModel) Call(ctx context.Context, input string) (string, error) {
	req := m.makeRequest([]flowllm.ChatMessage{{Role: "user", Content: input}})
	resp, err := m.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return 