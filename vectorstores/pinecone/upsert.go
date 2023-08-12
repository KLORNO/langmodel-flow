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
	Vectors   []pineconeItem `json:"vectors"`
	Namespace string         `json:"namespace,omitempty"`
}

func errorMessageFromErrorResponse(task string, body io.Reader) error {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, body)
	if err !=