
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
