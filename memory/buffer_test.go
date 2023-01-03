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

	BeforeEach(func() {
		ctx = context.Background()
	})

	It("saves user and assistant messages to chat history", func() {
		buf = memory.NewBuffer(0, nil)
		input := "User input message"
		output := "Assistant output message"
		err := buf.Save(ctx, input, output)
		Expect(err).NotTo(HaveOccurred())

		messages, err := buf.Load(ctx)
		Expect(err).NotTo(HaveOccurred())
		Expect(messages).To(HaveLen(2))
		Expect(messages[0])