
package pinecone

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/deluan/flowllm"
	"github.com/deluan/flowllm/vectorstores"
	"golang.org/x/exp/slices"
)

// Options for the Pinecone vector store.
type Options struct {
	ApiKey      string
	Environment string
	Index       string
	NameSpace   string
	Pods        int
	Replicas    int
	PodType     string
	Metric      Metric
}

// VectorStore is a vector store backed by Pinecone. It requires an already created Pinecone index,
// with the same dimensionality as the embeddings used to create the store.
type VectorStore struct {
	client     *client
	embeddings flowllm.Embeddings
	textKey    string
}

// NewVectorStore creates a new Pinecone vector store.
func NewVectorStore(ctx context.Context, embeddings flowllm.Embeddings, opts Options) (*VectorStore, error) {
	if opts.ApiKey == "" {
		opts.ApiKey = os.Getenv("PINECONE_API_KEY")
	}
	if opts.Environment == "" {
		opts.Environment = os.Getenv("PINECONE_ENVIRONMENT")
	}
	if opts.Index == "" {
		opts.Index = os.Getenv("PINECONE_INDEX")
	}
	if opts.Pods == 0 {
		opts.Pods = 1
	}
	if opts.Replicas == 0 {
		opts.Replicas = 1
	}
	if opts.PodType == "" {
		opts.PodType = "s1"
	}
	if opts.Metric == "" {
		opts.Metric = Cosine
	}
	c := &client{
		index:       opts.Index,
		pods:        opts.Pods,
		replicas:    opts.Replicas,
		podType:     opts.PodType,
		metric:      string(opts.Metric),
		apiKey:      opts.ApiKey,
		environment: opts.Environment,
		namespace:   opts.NameSpace,
	}
	s := VectorStore{
		embeddings: embeddings,
		client:     c,
		textKey:    "text",
	}
	err := connect(ctx, c)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *VectorStore) AddDocuments(ctx context.Context, documents ...flowllm.Document) error {
	var texts []string
	for i := 0; i < len(documents); i++ {
		texts = append(texts, documents[i].PageContent)
	}

	vectors, err := s.embeddings.EmbedStrings(ctx, texts)
	if err != nil {
		return err
	}

	var items []pineconeItem
	for i := 0; i < len(vectors); i++ {
		curMetadata := make(map[string]string)
		for key, value := range documents[i].Metadata {
			curMetadata[key] = fmt.Sprintf("%s", value)