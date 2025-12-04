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

package models

import "time"

// DeploymentResponse represents deployment information
type DeploymentResponse struct {
	AgentName                  string                      `json:"agentName"`
	ProjectName                string                      `json:"projectName"`
	ImageId                    string                      `json:"imageId"`
	Status                     string                      `json:"status"`
	Environment                string                      `json:"environment"`
	EnvironmentDisplayName     string                      `json:"environmentDisplayName"`
	PromotionTargetEnvironment *PromotionTargetEnvironment `json:"promotionTargetEnvironment,omitempty"`
	LastDeployedAt             time.Time                   `json:"lastDeployedAt"`
	Endpoints                  []Endpoint                  `json:"endpoints"`
}

// PromotionTargetEnvironment represents environment promotion targets
type PromotionTargetEnvironment struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// EndpointsResponse represents detailed endpoint information
type EndpointsResponse struct {
	Endpoint
	Schema EndpointSchema `json:"schema"`
}

// EndpointSchema represents the schema for an endpoint
type EndpointSchema struct {
	Content string `json:"content"`
}

// Endpoint represents endpoint configuration
type Endpoint struct {
	URL        string `json:"url"`
	Name       string `json:"name,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

// EnvVars represents environment variables
type EnvVars struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Build represents a build instance
type BuildResponse struct {
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	AgentName   string    `json:"agentName"`
	ProjectName string    `json:"projectName"`
	CommitID    string    `json:"commitId"`
	Status      string    `json:"status"`
	StartedAt   time.Time `json:"startedAt"`
	Image       string    `json:"image,omitempty"`
	Branch      string    `json:"branch,omitempty"`
	EndedAt     time.Time `json:"endedAt,omitempty"`
}

// BuildStep represents a step in the build process
type BuildStep struct {
	Type    string    `json:"type"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	At      time.Time `json:"at"`
}

// BuildDetails represents detailed build information
type BuildDetailsResponse struct {
	BuildResponse
	Percent         float32     `json:"percent,omitempty"`
	Steps           []BuildStep `json:"steps,omitempty"`
	DurationSeconds int32       `json:"durationSeconds,omitempty"`
	EndedAt         time.Time   `json:"endedAt,omitempty"`
}
