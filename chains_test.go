
package flowllm_test

import (
	"context"
	"errors"
	"time"

	. "github.com/deluan/flowllm"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	Describe("Chain", func() {
		It("should execute handlers in sequence and return the output of the last handler", func() {
			handler1 := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return Values{"key1": "value1"}, nil
			})
			handler2 := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				vals := Values{}.Merge(values...)
				vals["key2"] = "value2"
				return vals, nil
			})
			chain := Chain(handler1, handler2)
			result, err := chain.Call(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(Values{"key1": "value1", "key2": "value2"}))
		})

		It("should return an error if any of the handlers returns an error", func() {
			handler1 := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return Values{"key1": "value1"}, nil
			})
			handler2 := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return nil, errors.New("handler error")
			})
			chain := Chain(handler1, handler2)
			_, err := chain.Call(context.Background())
			Expect(err).To(MatchError("handler error"))
		})
	})

	Describe("MapOutputTo", func() {
		It("should rename the output of the chain to the given key", func() {
			handler := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return Values{DefaultKey: "value"}, nil
			})
			chain := Chain(handler, MapOutputTo("newKey"))
			result, err := chain.Call(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(Values{"newKey": "value"}))
		})
	})

	Describe("TrimSpace", func() {
		It("should trim all spaces from the values of the given keys", func() {
			handler := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return Values{"key1": " value1 ", "key2": " value2 "}, nil
			})
			chain := Chain(handler, TrimSpace("key1", "key2"))
			result, err := chain.Call(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(Values{"key1": "value1", "key2": "value2"}))
		})
	})

	Describe("TrimSuffix", func() {
		It("should trim the given suffix from the values of the given keys", func() {
			handler := HandlerFunc(func(ctx context.Context, values ...Values) (Values, error) {
				return Values{"key1": "value1_suffix", "key2": "value2_suffix"}, nil
			})
			chain := Chain(handler, TrimSuffix("_suffix", "key1", "key2"))
			result, err := chain.Call(context.Background())
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(Equal(Values{"key1": "value1", "key2": "value2"}))
		})