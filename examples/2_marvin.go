package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	. "github.com/deluan/flowllm"
	"github.com/deluan/flowllm/llms/openai"
	"github.com/deluan/flowllm/memory"
)

func init() {
	registerExample("marvin", "A simple chatbot using the GPT-3.5 model, using m