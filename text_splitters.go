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
// it will be returned as a single chunk. If the text is larger than the chunk
// size, it will be split into chunks of the given size, trying to split at the
// given separators. If the text cannot be split at any of the given separators,
// it will be split at the last separator.
func RecursiveTextSplitter(opts SplitterOptions) Splitter {
	if opts.ChunkSize == 0 {
		opts.ChunkSize = defaultSplitterChunkSize
	}
	if opts.LenFunc == nil {
		