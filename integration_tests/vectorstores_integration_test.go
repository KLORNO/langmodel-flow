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
		boltVS         flowllm.VectorStore
		memoryVS       flowllm.VectorStore
		pineconeVS     flowllm.VectorStore
		ctx            context.Context
		mockEmbeddings *FakeEmbeddings
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		mockEmbeddings = &FakeEmbeddings{}

		// Create a Memory VectorStore
		memoryVS = vectorstores.NewMemoryVectorStore(mockEmbeddings)

		// Create a BoltDB VectorStore
		boltTmpDB, err := os.CreateTemp("", "flowllm_bolt_*_.db")
		Expect(err).ToNot(HaveOccurred())
		_ = boltTmpDB.Close()
