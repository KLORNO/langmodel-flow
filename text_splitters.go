package flowllm

import (
	"log"
	"strings"
)

var (
	defaultSplitterChunkSize  = 1000
	defaultSplitterLenFunc    = func(s string) int { return len(s) }
	defaultSplitterSeparators = []string{"\n\n", "\n", " ", ""}
)

// SplitterOptions for the RecursiveTextSplitter splitter
type SplitterOptions struct {
	// ChunkSize is the maximum size of each chunk
	ChunkSize int
	// ChunkOverlap is the number of characters that will be repeated in each
	ChunkOverlap int
	// LenFunc is the length function to be used to calculate the chunk size
	LenFunc func(string) int
	// Separators is a list of strings that will be used to split the text
	Separators []string
}

// RecursiveTextSplitter splits a text into chunks of a given size, trying to
// split at the given separators. If the text is smaller than the chunk size,
// it will be re