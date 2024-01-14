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
		var b []float32
		Expect(CosineSimilarity(a, b)).To(Equal(float32(0)))
	})

	It("should return 1 when both input vectors are the same", func() {
		a := []float32{1, 2, 3}
		b := []float32{1, 2, 3}
		Expect(CosineSimilarity(a, b)).To(BeNumerically("~", float32(1), 1e-6))
	})

	It("should return 0 when input vectors are orthogonal", func() {
		a := []float32{1, 0, 0}
		b := []float32{0, 1, 0}
		Expect(CosineSimilarity(a, b)).To(BeNu