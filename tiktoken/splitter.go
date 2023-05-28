package tiktoken

import "github.com/deluan/flowllm"

func Splitter(model string, options flowllm.SplitterOptions) flowllm.Splitter {
	lenFunc := Len(mod