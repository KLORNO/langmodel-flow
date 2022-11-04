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
		opts.ApiKey = os.Getenv("OPENAI_API_KEY")
	}
	if opts.BatchSize == 0 {
		opts.BatchSize = 512
	}
	e := &Embeddings{opts: o