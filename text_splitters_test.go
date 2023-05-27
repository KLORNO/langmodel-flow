
package flowllm_test

import (
	. "github.com/deluan/flowllm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Splitters", func() {

	Describe("RecursiveTextSplitter", func() {
		var (
			text           string
			splitter       Splitter
			expectedOutput []string
		)

		BeforeEach(func() {
			text = "This is a sample text for testing the RecursiveTextSplitter function."
		})

		Context("with default options", func() {
			BeforeEach(func() {
				splitter = RecursiveTextSplitter(SplitterOptions{})
				expectedOutput = []string{
					"This is a sample text for testing the RecursiveTextSplitter function.",
				}
			})

			It("splits the text into chunks based on the default chunk size and overlap", func() {
				chunks, err := splitter(text)
				Expect(err).NotTo(HaveOccurred())
				Expect(chunks).To(Equal(expectedOutput))