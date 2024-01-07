package vectorstores_test

import (
	. "github.com/deluan/flowllm/vectorstores"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CosineSimilarity", func() {
	It("should return 0 when both input vectors are empty", func() {
		var a []float32
		var b []float32
		Expect(CosineSimilarity(a, b)).To(Equal(float32(0)))
	})

	It("should return 0 when one of the input vectors is empty", func() {
		a := []float32{1, 2, 3}