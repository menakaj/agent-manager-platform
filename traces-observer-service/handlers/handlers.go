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

package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/controllers"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/opensearch"
)

// Handler handles HTTP requests for tracing
type Handler struct {
	controllers *controllers.TracingController
}

// NewHandler creates a new handler
func NewHandler(controllers *controllers.TracingController) *Handler {
	return &Handler{
		controllers: controllers,
	}
}

// TraceRequest represents the request body for getting traces
type TraceRequest struct {
	ServiceName string `json:"serviceName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
	Limit       int    `json:"limit,omitempty"`
	SortOrder   string `json:"sortOrder,omitempty"`
}

// TraceByIdAndServiceRequest represents the request body for getting traces by ID and service
type TraceByIdAndServiceRequest struct {
	TraceID     string `json:"traceId"`
	ServiceName string `json:"serviceName"`
	SortOrder   string `json:"sortOrder,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// GetTraceOverviews handles GET /api/traces with query parameters
func (h *Handler) GetTraceOverviews(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query()

	serviceName := query.Get("serviceName")
	if serviceName == "" {
		h.writeError(w, http.StatusBadRequest, "serviceName is required")
		return
	}

	startTime := query.Get("startTime")
	endTime := query.Get("endTime")

	// Parse limit (default: 10)
	limit := 10
	if limitStr := query.Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			h.writeError(w, http.StatusBadRequest, "limit must be a positive integer")
			return
		}
		limit = parsedLimit
	}

	// Parse offset for pagination (default: 0)
	offset := 0
	if offsetStr := query.Get("offset"); offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil || parsedOffset < 0 {
			h.writeError(w, http.StatusBadRequest, "offset must be a non-negative integer")
			return
		}
		offset = parsedOffset
	}

	// Parse sortOrder (default: desc for traces - newest first)
	sortOrder := query.Get("sortOrder")
	if sortOrder == "" {
		sortOrder = "desc"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		h.writeError(w, http.StatusBadRequest, "sortOrder must be 'asc' or 'desc'")
		return
	}

	// Build query parameters
	params := opensearch.TraceQueryParams{
		ServiceName: serviceName,
		StartTime:   startTime,
		EndTime:     endTime,
		Limit:       limit,
		Offset:      offset,
		SortOrder:   sortOrder,
	}

	// Execute query
	ctx := r.Context()
	result, err := h.controllers.GetTraceOverviews(ctx, params)
	if err != nil {
		log.Printf("Failed to get trace overviews: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to retrieve trace overviews")
		return
	}

	// Write response
	h.writeJSON(w, http.StatusOK, result)
}

// GetTraceByIdAndService handles GET /api/trace with query parameters
func (h *Handler) GetTraceByIdAndService(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query()

	traceID := query.Get("traceId")
	if traceID == "" {
		h.writeError(w, http.StatusBadRequest, "traceId is required")
		return
	}

	serviceName := query.Get("serviceName")
	if serviceName == "" {
		h.writeError(w, http.StatusBadRequest, "serviceName is required")
		return
	}

	// Parse sortOrder (default: desc)
	sortOrder := query.Get("sortOrder")
	if sortOrder == "" {
		sortOrder = "desc"
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		h.writeError(w, http.StatusBadRequest, "sortOrder must be 'asc' or 'desc'")
		return
	}

	// Parse limit (default: 100 for spans)
	limit := 100
	if limitStr := query.Get("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit <= 0 {
			h.writeError(w, http.StatusBadRequest, "limit must be a positive integer")
			return
		}
		limit = parsedLimit
	}

	// Build query parameters
	params := opensearch.TraceByIdAndServiceParams{
		TraceID:     traceID,
		ServiceName: serviceName,
		SortOrder:   sortOrder,
		Limit:       limit,
	}

	// Execute query
	ctx := r.Context()
	result, err := h.controllers.GetTraceByIdAndService(ctx, params)
	if err != nil {
		log.Printf("Failed to get trace by ID and service: %v", err)
		h.writeError(w, http.StatusInternalServerError, "Failed to retrieve traces")
		return
	}

	// Write response
	h.writeJSON(w, http.StatusOK, result)
}

// Health handles GET /health
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := h.controllers.HealthCheck(ctx); err != nil {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}

	h.writeJSON(w, http.StatusOK, map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// Helper functions
func (h *Handler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode JSON: %v", err)
	}
}

func (h *Handler) writeError(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, ErrorResponse{
		Error:   "error",
		Message: message,
	})
}
