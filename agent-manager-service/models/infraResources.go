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

type EnvironmentResponse struct {
	Name         string    `json:"name"`
	Namespace    string    `json:"namespace"`
	DisplayName  string    `json:"displayName,omitempty"`
	IsProduction bool      `json:"isProduction"`
	DNSPrefix    string    `json:"dnsPrefix,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
}

type DeploymentPipelineResponse struct {
	Name           string          `json:"name"`
	DisplayName    string          `json:"displayName,omitempty"`
	Description    string          `json:"description,omitempty"`
	OrgName        string          `json:"orgName"`
	CreatedAt      time.Time       `json:"createdAt"`
	PromotionPaths []PromotionPath `json:"promotionPaths,omitempty"`
}

type PromotionPath struct {
	SourceEnvironmentRef  string                 `json:"sourceEnvironmentRef"`
	TargetEnvironmentRefs []TargetEnvironmentRef `json:"targetEnvironmentRefs"`
}
type TargetEnvironmentRef struct {
	Name             string `json:"name"`
	RequiresApproval bool   `json:"requiresApproval,omitempty"`
}

type LogEntry struct {
	Timestamp     time.Time         `json:"timestamp"`
	Log           string            `json:"log"`
	LogLevel      string            `json:"logLevel"` // ERROR, WARN, INFO, DEBUG
	ComponentId   string            `json:"componentId"`
	EnvironmentId string            `json:"environmentId"`
	ProjectId     string            `json:"projectId"`
	Version       string            `json:"version"`
	VersionId     string            `json:"versionId"`
	Namespace     string            `json:"namespace"`
	PodId         string            `json:"podId"`
	ContainerName string            `json:"containerName"`
	Labels        map[string]string `json:"labels"`
}

type BuildLogsResponse struct {
	Logs       []LogEntry `json:"logs"`
	TotalCount int32      `json:"totalCount"`
	TookMs     float32    `json:"tookMs"`
}
