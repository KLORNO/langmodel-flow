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
	//