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
	BatchSize    int
}
type Embeddings struct {
	client *openai.Client
	opts   EmbeddingsOptions
}

func NewEmbeddings(opts EmbeddingsOptions) (*Embeddings, error) {
	if opts.ApiKey == "" {
		opts.