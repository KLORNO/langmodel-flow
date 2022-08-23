package integration_tests

import (
	"context"
	"os"
	"strconv"

	"github.com/deluan/flowllm"
	"github.com/deluan/flowllm/vectorstores"
	"github.com/deluan/flowllm/vectorstores/bolt"
	"github.com/deluan/flowllm/vectorstores/pinecone"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vector Stores Integration Tests", func() {
	var (
		boltVS