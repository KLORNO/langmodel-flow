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
	AddDocuments(context.Context, ...Document) error
	// SimilaritySearch returns the k most similar documents to the query
	SimilaritySearch(ctx context.Context, query string, k int) ([]Document, error)
	// SimilaritySearchVectorWithScore returns the k most similar documents to the query, along with their similarity score
	SimilaritySearchVectorWithScore(ctx context.Context, query []float32, k int) ([]ScoredDocument, error)
}

// Document represents a document to be stored in a VectorStore.
type Document struct {
	ID          