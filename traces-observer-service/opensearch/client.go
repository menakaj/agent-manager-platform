// Copyright (c) 2025, WSO2 LLC (http://www.wso2.com). All Rights Reserved.
//
// This software is the property of WSO2 LLC and its suppliers, if any.
// Dissemination of any information or reproduction of any material contained
// herein is strictly forbidden, unless permitted by WSO2 in accordance with
// the WSO2 Commercial License available at http://wso2.com/licenses.
// For specific language governing the permissions and limitations under
// this license, please see the license as well as any agreement you've
// entered into with WSO2 governing the purchase of this software and any
// associated services.

package opensearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/opensearch-project/opensearch-go"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/config"
)

// Client wraps the OpenSearch client
type Client struct {
	client *opensearch.Client
	config *config.OpenSearchConfig
}

// NewClient creates a new OpenSearch client
func NewClient(cfg *config.OpenSearchConfig) (*Client, error) {
	opensearchConfig := opensearch.Config{
		Addresses: []string{cfg.Address},
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	client, err := opensearch.NewClient(opensearchConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenSearch client: %w", err)
	}

	// Test connection
	info, err := client.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to OpenSearch: %w", err)
	} else {
		log.Printf("Connected to OpenSearch, status: %s", info.Status())
	}

	return &Client{
		client: client,
		config: cfg,
	}, nil
}

// Search executes a search query
func (c *Client) Search(ctx context.Context, query map[string]interface{}) (*SearchResponse, error) {
	// Convert query to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("failed to encode query: %w", err)
	}

	// Execute search
	res, err := c.client.Search(
		c.client.Search.WithContext(ctx),
		c.client.Search.WithIndex(c.config.Index),
		c.client.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search returned error: %s", res.Status())
	}

	// Parse response
	var response SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}

// HealthCheck checks if OpenSearch is accessible
func (c *Client) HealthCheck(ctx context.Context) error {
	_, err := c.client.Info()
	return err
}
