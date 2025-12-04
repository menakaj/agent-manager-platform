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

package config

// Config holds all configuration for the application
type Config struct {
	ServerHost          string
	ServerPort          int
	AuthHeader          string
	AutoMaxProcsEnabled bool
	LogLevel            string
	POSTGRESQL          POSTGRESQL
	KubeConfig          string
	// HTTP Server timeout configurations
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	IdleTimeoutSeconds  int
	MaxHeaderBytes      int
	// Database operation timeout configuration
	DbOperationTimeoutSeconds int
	HealthCheckTimeoutSeconds int
	DefaultHTTPPort           int

	APIKeyHeader string
	APIKeyValue  string
	// CORSAllowedOrigin is the single allowed origin for CORS; use "*" to allow all
	CORSAllowedOrigin string

	// OpenTelemetry configuration
	OTEL OTELConfig

	// Observer service configuration
	Observer ObserverConfig

	IsLocalDevEnv bool
}

// OTELConfig holds all OpenTelemetry related configuration
type OTELConfig struct {
	// Instrumentation configuration
	InstrumentationImage    string
	InstrumentationProvider string
	SDKVolumeName           string
	SDKMountPath            string

	// Tracing configuration
	TraceContent     bool
	MetricsEnabled   bool
	TelemetryEnabled bool

	// OTLP Exporter configuration
	ExporterInsecure bool
	ExporterEndpoint string
}
type ObserverConfig struct {
	// Observer service URL
	URL string
}

type POSTGRESQL struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string `json:"-"`
	DbConfigs
}

type DbConfigs struct {
	// gorm configs
	SlowThresholdMilliseconds int64
	SkipDefaultTransaction    bool

	// go sql configs
	MaxIdleCount       *int64 // zero means defaultMaxIdleConns (2); negative means 0
	MaxOpenCount       *int64 // <= 0 means unlimited
	MaxLifetimeSeconds *int64 // maximum amount of time a connection may be reused
	MaxIdleTimeSeconds *int64
}
