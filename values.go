package flowllm

import (
	"encoding/json"
	"fmt"

	"golang.org/x/exp/maps"
)

const (
	DefaultKey     = "text"
	DefaultChatKey = "_chat_messages"
)

// Values is a map of string to any value. This is