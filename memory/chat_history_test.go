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
		It("returns an emp