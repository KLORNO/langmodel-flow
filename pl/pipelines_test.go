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
			It("calls the 'tra