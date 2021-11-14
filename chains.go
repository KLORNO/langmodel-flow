
package flowllm

import (
	"context"
	"fmt"
	"strings"

	"github.com/deluan/flowllm/pl"
)

// Handler is the interface implemented by all composable modules in the library.
type Handler interface {
	Call(ctx context.Context, values ...Values) (Values, error)
}

// HandlerFunc is a function that implements the Handler interface.
type HandlerFunc func(context.Context, ...Values) (Values, error)

func (f HandlerFunc) Call(ctx context.Context, values ...Values) (Values, error) {
	return f(ctx, values...)
}

// Chain is a special handler that executes a list of handlers in sequence.
// The output of each chain is passed as input to the next one.
// The output of the last chain is returned as the output of the Sequential chain.
func Chain(handlers ...Handler) HandlerFunc {
	return func(ctx context.Context, values ...Values) (Values, error) {
		vals := Values{}.Merge(values...)
		for _, chain := range handlers {
			var err error
			vals, err = chain.Call(ctx, vals)
			if err != nil {
				return nil, err
			}
		}
		return vals, nil
	}
}

// MapOutputTo renames the output of the chain (DefaultKey) to the given key.
func MapOutputTo(key string) HandlerFunc {
	return func(ctx context.Context, values ...Values) (Values, error) {
		vals := Values{}.Merge(values...)
		vals[key] = vals[DefaultKey]
		delete(vals, DefaultKey)
		return vals, nil
	}
}

// TrimSpace trims all spaces from the values of the given keys.
func TrimSpace(keys ...string) HandlerFunc {
	return func(ctx context.Context, values ...Values) (Values, error) {
		vals := Values{}.Merge(values...)
		for _, key := range keys {
			vals[key] = strings.TrimSpace(vals.Get(key))
		}
		return vals, nil
	}
}

// TrimSuffix trims the given suffix from the values of the given keys.
func TrimSuffix(suffix string, keys ...string) HandlerFunc {
	return func(ctx context.Context, values ...Values) (Values, error) {
		vals := Values{}.Merge(values...)
		for _, key := range keys {
			vals[key] = strings.TrimSuffix(vals.Get(key), suffix)
		}
		return vals, nil
	}
}

// ParallelChain executes a list of handlers in parallel, up to a maximum number of concurrent executions.