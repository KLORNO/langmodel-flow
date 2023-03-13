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
						return 0, errors.New("even number")
					}
					return i * 2, nil
				})

				Expect(<-errC).To(MatchError("even number"))
				Expect(<-outC).To(Equal(2))
				Expect(<-errC).To(MatchError("even number"))
				Expect(<-outC).To(Equal(6))

				Eventually(outC).Should(BeClosed())
				Eventually(errC).Should(BeClosed())
			})
		})
		Context("Multiple workers", func() {
			const maxWorkers = 2
			const numJobs = 100
			It("starts multiple workers, respecting the limit", func() {
				inC := make(chan int, numJobs)
				for i := 0; i < numJobs; i++ {
					inC <- i
				}
				close(inC)

				current := atomic.Int32{}
				count := atomic.Int32{}
				max := atomic.Int32{}
				outC, _ := pl.Stage(context.Background(), maxWorkers, inC, func(ctx context.Context, in int) (int, error) {
					defer current.Add(-1)
					c := current.Add(1)
					count.Add(1)
					if c > max.Load() {
						max.Store(c)
					}
					time.Sleep(10 * time.Millisecond) // Slow process
					return 0, nil
				})
				// Discard output and wait for completion
				for range outC {
				}

				Expect(count.Load()).To(Equal(int32(numJobs)))
				Expect(current.Load()).To(