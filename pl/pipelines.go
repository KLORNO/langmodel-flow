
// Package pl implements some Data Pipeline helper functions.
// Reference: https://medium.com/amboss/applying-modern-go-concurrency-patterns-to-data-pipelines-b3b5327908d4#3a80
//
// See also:
//
//	https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html#fano_fani
//	https://www.youtube.com/watch?v=f6kdp27TYZs
//	https://www.youtube.com/watch?v=QDDwwePbDtw
package pl

import (
	"context"
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/semaphore"
)

func Stage[In any, Out any](
	ctx context.Context,
	maxWorkers int,
	inputChannel <-chan In,
	fn func(context.Context, In) (Out, error),
) (chan Out, chan error) {
	outputChannel := make(chan Out)
	errorChannel := make(chan error)

	limit := int64(maxWorkers)
	sem1 := semaphore.NewWeighted(limit)
