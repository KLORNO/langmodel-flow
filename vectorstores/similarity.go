package vectorstores

import (
	"context"
	"math"

	"github.com/deluan/flowllm"
)

// CosineSimilarity calculates the cosine similarity between two vectors.
func CosineSimilarity(a, b []float32) float32 {
	var p, p2, q2 float32
	for i := 0; i < len(a) && i < len(b); i++ {
		p += a[i] * b[i]
		p2 += a[i] * a[i]
		q2 += b[i] * b[i]
	}
