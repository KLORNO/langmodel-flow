package pinecone

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

type pineconeItem struct {
	Values   []float32         `json:"values"`
	Metadata map[string]string `json:"metadata"`
	ID       string            `json:"id"`
}

type upsertPayload struct {
	Vectors   []pineconeItem `json:"vector