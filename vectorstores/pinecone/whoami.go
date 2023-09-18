package pinecone

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type whoAmIResponse struct {
	ProjectName string `json:"project_name"`
	UserLabe