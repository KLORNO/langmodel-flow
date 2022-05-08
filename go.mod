module github.com/deluan/flowllm

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/onsi/ginkgo/v2 v2.9.2
	github.com/onsi/gomega v1.27.6
	github.com/sashabaranov/go-openai v1.9.0
	github.com/tiktoken-go/tokenizer v0.1.0
	go.etcd.io/bbolt v1.3.7
	golang.org/x/exp v0.0.0-20230425010034-47ecfdc1ba53
)

require github.com/dlclark/regexp2 v1.9.0 // indirect

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230