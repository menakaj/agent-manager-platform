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

import "time"

// TraceQueryParams holds parameters for trace queries
type TraceQueryParams struct {
	ServiceName string
	StartTime   string
	EndTime     string
	Limit       int
	Offset      int
	SortOrder   string
}

// TraceByIdAndServiceParams holds parameters for querying by both traceId and serviceName
type TraceByIdAndServiceParams struct {
	TraceID     string
	ServiceName string
	SortOrder   string
	Limit       int
}

// Span represents a single trace span
type Span struct {
	TraceID         string                 `json:"traceId"`
	SpanID          string                 `json:"spanId"`
	ParentSpanID    string                 `json:"parentSpanId,omitempty"`
	Name            string                 `json:"name"`
	Service         string                 `json:"service"`
	StartTime       time.Time              `json:"startTime"`
	EndTime         time.Time              `json:"endTime,omitempty"`
	DurationInNanos int64                  `json:"durationInNanos"` // in nanoseconds
	Kind            string                 `json:"kind,omitempty"`
	Status          string                 `json:"status,omitempty"`
	Attributes      map[string]interface{} `json:"attributes,omitempty"`
	Resource        map[string]interface{} `json:"resource,omitempty"`
}

// TraceResponse represents the response for trace queries
type TraceResponse struct {
	Spans      []Span `json:"spans"`
	TotalCount int    `json:"totalCount"`
}

// TraceDetailResponse represents detailed information for a single trace
type TraceDetailResponse struct {
	TraceID    string   `json:"traceId"`
	Spans      []Span   `json:"spans"`
	TotalSpans int      `json:"totalSpans"`
	Duration   int64    `json:"duration"` // Total trace duration in microseconds
	Services   []string `json:"services"` // List of services involved
}

// TraceOverview represents a single trace overview with root span info
type TraceOverview struct {
	TraceID         string `json:"traceId"`
	RootSpanID      string `json:"rootSpanId"`
	RootSpanName    string `json:"rootSpanName"`
	StartTime       string `json:"startTime"`
	EndTime         string `json:"endTime"`
	DurationInNanos int64  `json:"durationInNanos"` // Total trace duration in nanoseconds
	SpanCount       int    `json:"spanCount"`
}

// TraceOverviewResponse represents the response for trace overview queries
type TraceOverviewResponse struct {
	Traces     []TraceOverview `json:"traces"`
	TotalCount int             `json:"totalCount"`
}

// SearchResponse represents OpenSearch search response
type SearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source map[string]interface{} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
