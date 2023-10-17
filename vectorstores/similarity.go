package vectorstores

import (
	"context"
	"math"

	"github.com/deluan/flowllm"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(a, b []float32) float32 {
	var p, p2, q2 float32
	for i