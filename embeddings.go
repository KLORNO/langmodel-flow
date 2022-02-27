package flowllm

import "context"

// Embeddings can be used to create a numerical representation of textual data.
// This numerical representation is useful when searching for similar documents.
type Embeddings interface {
	// EmbedString returns the embed