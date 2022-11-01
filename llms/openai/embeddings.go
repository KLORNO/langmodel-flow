package openai

import (
	"context"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type EmbeddingsOptions struct {
	ApiKey       string
	KeepNewLines bool
