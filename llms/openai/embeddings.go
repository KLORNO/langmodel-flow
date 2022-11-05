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
	e := &Embeddings{opts: opts}
	e.client = openai.NewClient(opts.ApiKey)

	return e, nil
}

type Option func(*Embeddings)

func (o *Embeddings) EmbedString(ctx context.Context, text string) ([]float32, error) {
	texts := o.prepareTexts([]string{text})
	embeddings, err := o.embedTexts(ctx, []string{texts[0]})
	if err != nil {
		return nil, err
	}
	return embeddings[0], nil
}

func (o *Embeddings) EmbedStrings(ctx context.Context, texts []string) ([][]float32, error) {
	chunks := chunkArray(o.prepareTexts(texts), o.opts.BatchSize