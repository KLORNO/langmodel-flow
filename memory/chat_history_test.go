package memory_test

import (
	"github.com/deluan/flowllm"
	"github.com/deluan/flowllm/memory"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChatMessageHistory", func() {
	var history *memory.ChatMessageHistory

	BeforeEach(func() {
		history = &memory.ChatMessageHistory{}
	})

	Context("GetMessages", func() {
		It("returns an empty slice when there are no messages", func() {
			Expect(history.GetMessages()).To(BeEmpty())
		})

		It("returns a copy of messages in the history", func() {
			history.AddUserMessage("Test user message")
			history.AddAssistantMessage("Test assistant message")

			messages := history.GetMessages()
			Expect(messages).To(HaveLen(2))
			Expect(messages[0]).To(Equal(flowllm.ChatMessage{Content