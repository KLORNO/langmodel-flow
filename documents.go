package flowllm

import (
	"context"
	"errors"
	"io"
)

// VectorStore is a particular type of database optimized for storing documents and their embeddings,
// and then fetching of the most relevant documents