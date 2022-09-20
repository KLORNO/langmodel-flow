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
		var closeDB func()
		boltVS, closeDB, err = bolt.NewVectorStore(mockEmbeddings, bolt.Options{Path: boltTmpDB.Name()})
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(closeDB)
		DeferCleanup(func() { _ = os.Remove(boltTmpDB.Name()) })

		if os.Getenv("PINECONE_API_KEY") != "" {
			// Create a Pinecone VectorStore
			pineconeVS, err = pinecone.NewVectorStore(ctx, mockEmbeddings,
				pinecone.Options{
					Index:     os.Getenv("PINECONE_INDEX_INTEGRATION_TEST"),
					NameSpace: "flowllm-integration-tests-" + uuid.NewString(),
				},
			)
			Expect(err).ToNot(HaveOccurred())
		}
	})

	DescribeTable("It should perform a similarity search using the query string and return correct results",
		func(getStore func() flowllm.VectorStore) {
			store := getStore()
			if store == nil {
				Skip("Skipping test. No VectorStore found.")
			}
			documents := []flowllm.Document{
				{
					PageContent: "first document",
					Metadata:    map[string]any{"key1": "value1"},
				},
				{
					PageContent: "second document",
					Metadata:    m