package pinecone

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type whoAmIResponse struct {
	ProjectName string `json:"project_name"`
	UserLabel   string `json:"user_label"`
	UserName    string `json:"user_name"`
}

func (c *cli