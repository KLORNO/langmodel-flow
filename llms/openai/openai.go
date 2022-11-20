
package openai

import (
	"context"
	"errors"
	"os"

	"github.com/sashabaranov/go-openai"
)

const (
	defaultModel     = "text-ada-001"
	defaultMaxTokens = 256
)

// Options for OpenAI Completions models
type Options struct {
	ApiKey           string
	Model            string
	Temperature      float32
	MaxTokens        int
	TopP             float32