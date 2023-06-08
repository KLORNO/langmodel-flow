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

// Values is a map of string to any value. This is the type used to pass values between handlers.
type Values map[string]any

// Merge merges multiple Values into one.
func (value Values) Merge(values ...Values) Values {
	res := Values{}
	for k, v := 