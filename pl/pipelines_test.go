package pl_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/deluan/flowllm/pl"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPipeline(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pipeline Tests Suite")
}

var _ = Describe("Pipeline", func() {
	Describe("Stage", func() {
		Context("happy path", func() {
			It("calls the 'transform' function and returns values and errors", func() {
				inC := make(chan int, 4)
				for i := 0; i < 4; i++ {
					inC <- i
				}
				close(inC)

				outC, errC := pl.Stage(context.Background(), 1, inC, func(ctx context.Context, i int) (int, error) {
					if i%2 == 0 {
						return 0, errors.N