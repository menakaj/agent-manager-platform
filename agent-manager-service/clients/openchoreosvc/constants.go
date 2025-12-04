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

package openchoreosvc

type LabelKeys string

const (
	LabelKeyOrganizationName LabelKeys = "openchoreo.dev/organization"
	LabelKeyProjectName      LabelKeys = "openchoreo.dev/project"
	LabelKeyComponentName    LabelKeys = "openchoreo.dev/component"
	LabelKeyComponentType    LabelKeys = "agent-manager/component-type"
)

type AnnotationKeys string

const (
	AnnotationKeyDisplayName AnnotationKeys = "openchoreo.dev/display-name"
	AnnotationKeyDescription AnnotationKeys = "openchoreo.dev/description"
)

type BuildTemplateNames string

const (
	GoogleBuildpackBuildTemplate    BuildTemplateNames = "buildpack-ci"
	BallerinaBuildpackBuildTemplate BuildTemplateNames = "ballerina-buildpack-ci"
)

const (
	AgentComponentType string = "agent-component"
	GoogleEntryPoint   string = "google-entry-point"
	LanguageVersion    string = "language-version"
	LanguageVersionKey string = "language-version-key"
)

// Build condition types
type BuildConditionType string

const (
	ConditionBuildInitiated  BuildConditionType = "BuildInitiated"
	ConditionBuildTriggered  BuildConditionType = "BuildTriggered"
	ConditionBuildCompleted  BuildConditionType = "BuildCompleted"
	ConditionWorkloadUpdated BuildConditionType = "WorkloadUpdated"
)

const (
	statusUnknown   = "Unknown"
	statusCompleted = "Completed"
)

// ServiceBinding condition types
const (
	ConditionActive         = "Active"
	ConditionFailed         = "Failed"
	ConditionInProgress     = "InProgress"
	ConditionNotYetDeployed = "NotYetDeployed"
	ConditionSuspended      = "Suspended"
)

// Deployment status values
const (
	DeploymentStatusFailed      = "failed"
	DeploymentStatusNotDeployed = "not-deployed"
	DeploymentStatusSuspended   = "suspended"
	DeploymentStatusInProgress  = "in-progress"
	DeploymentStatusActive      = "active"
)

const (
	EndpointTypeDefault = "DEFAULT"
	EndpointTypeCustom  = "CUSTOM"
)

const (
	MainContainerName                    = "main"
	DevEnvironmentName                   = "development"
	DevEnvironmentDisplayName            = "Development"
	DefaultDisplayName                   = "Default"
	DefaultName                          = "default"
	DefaultAPIClassNameWithCORS          = "default-with-cors"
	ObservabilityEnabledServiceClassName = "default-otel-supported"
	DefaultServiceClassName              = "default"
)

// Resource constants
const (
	DefaultCPURequest    = "100m"
	DefaultMemoryRequest = "64Mi"
	DefaultCPULimit      = "400m"
	DefaultMemoryLimit   = "256Mi"
)

const (
	BuildPlaneKind = "BuildPlane"
	DataPlaneKind  = "DataPlane"
)

// Environment variable names for otel and tracing configuration
const (
	EnvPythonPath                   = "PYTHONPATH"
	EnvAMPTraceloopTraceContent     = "AMP_TRACELOOP_TRACE_CONTENT"
	EnvAMPOTELExporterOTLPInsecure  = "AMP_OTEL_EXPORTER_OTLP_INSECURE"
	EnvAMPTraceloopMetricsEnabled   = "AMP_TRACELOOP_METRICS_ENABLED"
	EnvAMPTraceloopTelemetryEnabled = "AMP_TRACELOOP_TELEMETRY_ENABLED"
	EnvAMPOTELExporterOTLPEndpoint  = "AMP_OTEL_EXPORTER_OTLP_ENDPOINT"
	EnvInstrumentationProvider      = "INSTRUMENTATION_PROVIDER"
	EnvAMPComponentID               = "AMP_COMPONENT_ID"
	EnvAMPAppName                   = "AMP_APP_NAME"
	EnvAMPAppVersion                = "AMP_APP_VERSION"
	EnvAMPEnv                       = "AMP_ENV"
)
