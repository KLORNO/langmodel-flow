package memory

import (
	"context"

	"github.com/deluan/flowllm"
)

type Buffer struct {
	chatHistory *ChatMessageHistory
	windowSize 