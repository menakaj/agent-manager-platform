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
	"fmt"
	"time"
)

// ParseSpans converts OpenSearch response to Span structs
func ParseSpans(response *SearchResponse) []Span {
	spans := make([]Span, 0, len(response.Hits.Hits))

	for _, hit := range response.Hits.Hits {
		span := parseSpan(hit.Source)
		spans = append(spans, span)
	}

	return spans
}

// parseSpan extracts span information from a source document
func parseSpan(source map[string]interface{}) Span {
	span := Span{}

	// Try standard OTEL fields first
	if traceID, ok := source["traceId"].(string); ok {
		span.TraceID = traceID
	}
	if spanID, ok := source["spanId"].(string); ok {
		span.SpanID = spanID
	}
	if parentSpanID, ok := source["parentSpanId"].(string); ok {
		span.ParentSpanID = parentSpanID
	}
	if name, ok := source["name"].(string); ok {
		span.Name = name
	}
	if kind, ok := source["kind"].(string); ok {
		span.Kind = kind
	}

	// Extract service name from nested resource.attributes.service.name
	if resource, ok := source["resource"].(map[string]interface{}); ok {
		if attributes, ok := resource["attributes"].(map[string]interface{}); ok {
			if serviceName, ok := attributes["service.name"].(string); ok {
				span.Service = serviceName
			}
		}

		// Store the complete resource object
		span.Resource = resource
	}
	// Fallback to serviceName if exists
	if span.Service == "" {
		if service, ok := source["serviceName"].(string); ok {
			span.Service = service
		}
	}

	// Parse timestamps
	if startTime, ok := source["startTime"].(string); ok {
		if t, err := time.Parse(time.RFC3339Nano, startTime); err == nil {
			span.StartTime = t
		}
	}
	if endTime, ok := source["endTime"].(string); ok {
		if t, err := time.Parse(time.RFC3339Nano, endTime); err == nil {
			span.EndTime = t
		}
	}

	// Parse duration
	if duration, ok := source["durationInNanos"].(float64); ok {
		span.DurationInNanos = int64(duration)
	}

	// Parse status
	if status, ok := source["status"].(map[string]interface{}); ok {
		if code, ok := status["code"].(string); ok {
			span.Status = code
		} else if code, ok := status["code"].(float64); ok {
			span.Status = fmt.Sprintf("%d", int(code))
		}
	}

	// Parse attributes
	if attributes, ok := source["attributes"].(map[string]interface{}); ok {
		span.Attributes = attributes
	}

	return span
}
