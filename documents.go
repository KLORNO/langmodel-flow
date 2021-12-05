package flowllm

import (
	"context"
	"errors"
	"io"
)

// VectorStore is a particular type of database optimized for storing documents and their embeddings,
// and then fetching of the most relevant documents for a particular query, i.e. those whose embeddings
// are most similar to the embedding of the query.
type VectorStore interface {
	// AddDocuments adds the given documents to the store
	AddDocuments(context.Co