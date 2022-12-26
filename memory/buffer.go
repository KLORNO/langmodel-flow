package memory

import (
	"context"

	"github.com/deluan/flowllm"
)

type Buffer struct {
	chatHistory *ChatMessageHistory
	windowSize  int
}

func NewBuffer(windowSize int, history *flowllm.ChatMessages) *Buffer {
	chatHistory := &ChatMessageHistory