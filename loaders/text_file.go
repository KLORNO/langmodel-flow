package loaders

import (
	"context"
	"io"
	"os"

	"github.com/deluan/flowllm"
)

func TextFile(path string, splitter ...flowllm.Splitter) flowllm.DocumentLoaderFunc {
	var docs []flowllm.Document
	var idx int
	var spl flowllm.Splitter
	if len(splitter) > 0 {
		spl = splitter[0]
	}

	return func(context.Context) (flowllm.Document, error) {
		// Return next document if already loaded
		if len(docs) > 0 {
			if idx < len(docs) {
				idx++
				return docs[idx-1], nil
			}
			re