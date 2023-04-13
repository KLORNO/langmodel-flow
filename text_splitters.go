package flowllm

import (
	"log"
	"strings"
)

var (
	defaultSplitterChunkSize  = 1000
	defaultSplitterLenFunc    = func(