
package bolt

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"time"

	"github.com/deluan/flowllm"
	"github.com/deluan/flowllm/vectorstores"
	"go.etcd.io/bbolt"
	"golang.org/x/exp/slices"
)

const (
	DefaultPath       = "vector_store.db"
	DefaultBucket     = "embeddings"
	DefaultPermission = 0600
)

// Options for the Bolt vector store.
type Options struct {
	Path       string
	Bucket     string
	Permission fs.FileMode
	Timeout    time.Duration
}

// VectorStore is a vector store backed by BoltDB. It implements the flowllm.VectorStore interface,
// and it is ideal for small to medium-sized collections of vectors.
type VectorStore struct {
	embeddings flowllm.Embeddings
	db         *bbolt.DB
	bucket     string
}

// NewVectorStore creates a new Bolt vector store.
func NewVectorStore(embeddings flowllm.Embeddings, opts Options) (*VectorStore, func(), error) {
	if opts.Path == "" {
		opts.Path = DefaultPath
	}
	if opts.Bucket == "" {
		opts.Bucket = DefaultBucket
	}
	if opts.Permission == 0 {
		opts.Permission = DefaultPermission
	}
	if opts.Timeout == 0 {
		opts.Timeout = time.Second
	}
	s := VectorStore{
		embeddings: embeddings,
		bucket:     opts.Bucket,
	}
	db, err := bbolt.Open(opts.Path, opts.Permission, &bbolt.Options{Timeout: opts.Timeout})
	if err != nil {
		return nil, func() {}, err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(s.bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, func() {}, err
	}
	s.db = db
	return &s, func() { _ = db.Close() }, nil
}

type boltItem struct {
	Vectors  []float32              `json:"vectors"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

func (d boltItem) id() string {
	return fmt.Sprintf("%x", sha256.Sum256(d.Marshall()))
}

func (d boltItem) Marshall() []byte {
	buf, _ := json.Marshal(d)
	return buf
}

func (s *VectorStore) AddDocuments(ctx context.Context, documents ...flowllm.Document) error {
	texts := make([]string, len(documents))
	for i, document := range documents {
		texts[i] = document.PageContent
	}
	vectors, err := s.embeddings.EmbedStrings(ctx, texts)
	if err != nil {
		return err
	}

	return s.db.Batch(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(s.bucket))
		for i, doc := range documents {
			item := boltItem{
				Vectors:  vectors[i],
				Content:  doc.PageContent,
				Metadata: doc.Metadata,
			}
			if err := bucket.Put([]byte(item.id()), item.Marshall()); err != nil {
				return err
			}
		}
		return nil
	})
}

type match struct {