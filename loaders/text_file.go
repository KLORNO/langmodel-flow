package loaders

import (
	"context"
	"io"
	"os"

	"github.com/deluan/flowllm"
)

func TextFile(path string, splitter ...flowllm.Splitter) flowllm.DocumentLoaderF