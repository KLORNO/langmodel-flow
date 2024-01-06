package vectorstores_test

import (
	. "github.com/deluan/flowllm/vectorstores"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CosineSimilarity", func() {
	It("should return 0 when both input vectors are empt