package memory_test

import (
	"context"
	"strconv"

	"github.com/deluan/flowllm"
	"github.com/deluan/flowllm/memory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Buffer", func() {
	var ctx context.Context
	var buf *memory.Buffer

	BeforeEach(func()