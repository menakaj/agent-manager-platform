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

// BuildTraceQuery builds an OpenSearch query for traces
func BuildTraceQuery(params TraceQueryParams) map[string]interface{} {
	// Build the must conditions
	mustConditions := []map[string]interface{}{}

	// Add service name filter
	if params.ServiceName != "" {
		mustConditions = append(mustConditions, map[string]interface{}{
			"term": map[string]interface{}{
				"resource.attributes.service.name": params.ServiceName,
			},
		})
	}

	// Add time range filter
	if params.StartTime != "" && params.EndTime != "" {
		mustConditions = append(mustConditions, map[string]interface{}{
			"range": map[string]interface{}{
				"startTime": map[string]interface{}{
					"gte": params.StartTime,
					"lte": params.EndTime,
				},
			},
		})
	}

	// Set default limit if not provided
	limit := params.Limit
	if limit == 0 {
		limit = 100
	}

	// Set default offset
	offset := params.Offset
	if offset < 0 {
		offset = 0
	}

	// Set default sort order
	sortOrder := params.SortOrder
	if sortOrder == "" {
		sortOrder = "desc"
	}

	// Build the complete query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustConditions,
			},
		},
		"size": limit,
		"from": offset,
		"sort": []map[string]interface{}{
			{
				"startTime": map[string]string{
					"order": sortOrder,
				},
			},
		},
	}

	return query
}

// BuildTraceByIdAndServiceQuery builds a query to get spans by both traceId and serviceName
func BuildTraceByIdAndServiceQuery(params TraceByIdAndServiceParams) map[string]interface{} {
	// Build the must conditions - both traceId and serviceName must match
	mustConditions := []map[string]interface{}{
		{
			"term": map[string]interface{}{
				"traceId": params.TraceID,
			},
		},
		{
			"match": map[string]interface{}{
				"resource.attributes.service.name": params.ServiceName,
			},
		},
	}

	// Set default limit if not provided
	limit := params.Limit
	if limit == 0 {
		limit = 10000 // Get all spans for the trace by default
	}

	// Set default sort order
	sortOrder := params.SortOrder
	if sortOrder == "" {
		sortOrder = "asc"
	}

	// Build the complete query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustConditions,
			},
		},
		"size": limit,
		"sort": []map[string]interface{}{
			{
				"startTime": map[string]string{
					"order": sortOrder,
				},
			},
		},
	}

	return query
}
