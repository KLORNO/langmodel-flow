package memory

import (
	"github.com/deluan/flowllm"
)

type ChatMessageHistory struct {
	messages []flowllm.ChatMessage
}

func (h *ChatMessageHistory) GetMessages() flowllm.ChatMessages {
	copyMessages := make(flowllm.ChatMessages, len(h.messages))
	copy(copyMessages, h.messages)
	return copyMessages
}

fun