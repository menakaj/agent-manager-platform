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

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/config"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/controllers"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/handlers"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/middleware"
	"github.com/wso2-enterprise/agent-management-platform/traces-observer-service/opensearch"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting tracing service on port %d", cfg.Server.Port)

	// Initialize OpenSearch client
	osClient, err := opensearch.NewClient(&cfg.OpenSearch)
	if err != nil {
		log.Fatalf("Failed to create OpenSearch client: %v", err)
		// Exit if OpenSearch client cannot be created
		// Raise Error and exit
		os.Exit(1)
	}

	// Initialize service
	tracingController := controllers.NewTracingController(osClient)

	// Initialize handlers
	handler := handlers.NewHandler(tracingController)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/traces", handler.GetTraceOverviews)
	mux.HandleFunc("/api/trace", handler.GetTraceByIdAndService)
	mux.HandleFunc("/health", handler.Health)

	// Apply CORS middleware
	corsConfig := middleware.DefaultCORSConfig()
	corsHandler := middleware.CORS(corsConfig)(mux)

	// Create server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      corsHandler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on :%d", cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
