# Type Synchronization Analysis: Backend (Go) ↔ Frontend (TypeScript)

**Date**: November 18, 2025  
**Purpose**: Document the type definitions between agent-manager-service (Go) and console types (TypeScript)

---

## Overview

The **agent-manager-service** (Go backend) uses OpenAPI code generation to create type-safe API models in the `spec/` directory. These types are generated from an OpenAPI specification and should be the source of truth for the frontend TypeScript types in `console/workspaces/libs/types/`.

---

## Type Generation Strategy

### Backend (Go)
- **Location**: `agent-manager-service/spec/`
- **Generation Method**: OpenAPI Generator (generates Go structs from OpenAPI spec)
- **Command**: `make spec` (as documented in README)
- **Characteristics**:
  - Auto-generated code with strict type safety
  - Uses `json` tags for serialization
  - Includes validation and helper methods
  - All types have `*` prefix for optional fields (pointer types)

### Frontend (TypeScript)
- **Location**: `console/workspaces/libs/types/src/api/`
- **Generation Method**: Manual TypeScript definitions
- **Build Command**: `rushx build` (compiles TypeScript to JS + type declarations)
- **Characteristics**:
  - Manually maintained
  - Uses `?` for optional fields
  - ISO date-time strings (not native Date objects)
  - More flexible with union types

---

## Key Type Mappings

### Agent Types

#### Backend (Go) - `spec/model_agent_response.go`
```go
type AgentResponse struct {
    AgentName    string       `json:"agentName"`
    Description  string       `json:"description"`
    CreatedAt    time.Time    `json:"createdAt"`
    ProjectName  string       `json:"projectName"`
    Status       *string      `json:"status,omitempty"`
    Provisioning Provisioning `json:"provisioning"`
}

type Provisioning struct {
    Type       string            `json:"type"`
    Repository *RepositoryConfig `json:"repository,omitempty"`
}
```

#### Frontend (TypeScript) - `api/agents.ts` ✅ **UPDATED**
```typescript
export interface AgentResponse {
  agentName: string;
  description: string;
  createdAt: string; // ISO date-time
  projectName: string;
  status?: string;
  provisioning: Provisioning;
}

export interface Provisioning {
  type: string;
  repository?: RepositoryConfig;
}
```

**Changes Made**:
- ✅ Restructured `AgentResponse` to match backend (removed `AgentMetaDataResponse`)
- ✅ Added `Provisioning` interface
- ✅ Made `provisioning` a required field in `AgentResponse`
- ✅ Made `status` optional

---

### Create Agent Request

#### Backend (Go) - `spec/model_create_agent_request.go`
```go
type CreateAgentRequest struct {
    Provisioning   Provisioning          `json:"provisioning"`
    AgentName      string                `json:"agentName"`
    Description    *string               `json:"description,omitempty"`
    RuntimeConfigs *RuntimeConfiguration `json:"runtimeConfigs,omitempty"`
    InputInterface *InputInterface       `json:"inputInterface,omitempty"`
}
```

#### Frontend (TypeScript) - `api/agents.ts` ✅ **UPDATED**
```typescript
export interface CreateAgentRequest {
  provisioning: Provisioning;
  agentName: string;
  description?: string;
  runtimeConfigs?: ContainerSpec;
  inputInterface?: InputInterface;
}
```

**Changes Made**:
- ✅ Made `provisioning` the first required field (matches backend)
- ✅ Made `description`, `runtimeConfigs`, and `inputInterface` optional
- ✅ Removed old structure where `repository` was required

---

### Agent List Response

#### Backend (Go) - `spec/model_agent_list_response.go`
```go
type AgentListResponse struct {
    Agents []AgentResponse `json:"agents"`
    Total  int32           `json:"total"`
    Limit  int32           `json:"limit"`
    Offset int32           `json:"offset"`
}
```

#### Frontend (TypeScript) - `api/agents.ts` ✅ **UPDATED**
```typescript
export interface AgentListResponse extends PaginationMeta {
  agents: AgentResponse[]; // Now using AgentResponse instead of AgentMetaDataResponse
}
```

**Changes Made**:
- ✅ Changed to use `AgentResponse[]` array (was using a separate `AgentMetaDataResponse`)

---

### Build Types

#### Backend (Go) - `spec/model_build_response.go`
```go
type BuildResponse struct {
    BuildId     *string    `json:"buildId,omitempty"`
    BuildName   string     `json:"buildName"`
    ProjectName string     `json:"projectName"`
    AgentName   string     `json:"agentName"`
    CommitId    string     `json:"commitId"`
    StartedAt   time.Time  `json:"startedAt"`
    EndedAt     *time.Time `json:"endedAt,omitempty"`
    ImageId     *string    `json:"imageId,omitempty"`
    Status      *string    `json:"status,omitempty"`
    Branch      string     `json:"branch"`
}
```

#### Frontend (TypeScript) - `api/builds.ts` ✅ **UPDATED**
```typescript
export interface BuildResponse {
  buildId?: string;
  buildName: string;
  projectName: string;
  agentName: string;
  commitId: string;
  startedAt: string; // ISO date-time
  endedAt?: string; // ISO date-time
  imageId?: string;
  status?: string; // Changed from union type to string for flexibility
  branch: string;
}
```

**Changes Made**:
- ✅ Changed `status` from `BuildStatus` union type to `string` for flexibility with backend values

---

### Build Step Types

#### Backend (Go) - `spec/model_build_step.go`
```go
type BuildStep struct {
    Type    string    `json:"type"`
    Status  string    `json:"status"`
    Message string    `json:"message"`
    At      time.Time `json:"at"`
}
```

#### Frontend (TypeScript) - `api/builds.ts` ✅ **UPDATED**
```typescript
export interface BuildStep {
  type: string; // Changed from BuildStepType union to string
  status: string; // Changed from BuildStepStatus union to string
  message: string;
  at: string; // ISO date-time
}
```

**Changes Made**:
- ✅ Changed `type` from `BuildStepType` union to `string` for flexibility
- ✅ Changed `status` from `BuildStepStatus` union to `string` for flexibility

---

### Deployment Types

#### Backend (Go) - `spec/model_deployment_response.go`
```go
type DeploymentResponse struct {
    AgentName   string `json:"agentName"`
    ProjectName string `json:"projectName"`
    ImageId     string `json:"imageId"`
    Environment string `json:"environment"`
}
```

#### Frontend (TypeScript) - `api/deployments.ts` ✅ **ALIGNED**
```typescript
export interface DeploymentResponse {
  agentName: string;
  projectName: string;
  imageId: string;
  environment: string;
}
```

**Status**: Already aligned ✅

---

### Deployment Details Response

#### Backend (Go) - `spec/model_deployment_details_response.go`
```go
type DeploymentDetailsResponse struct {
    ImageId                    string                                     `json:"imageId"`
    Status                     string                                     `json:"status"`
    LastDeployed               time.Time                                  `json:"lastDeployed"`
    Endpoints                  []DeploymentEndpoint                       `json:"endpoints"`
    SourceEnvironment          EnvironmentObject                          `json:"sourceEnvironment"`
    EnvironmentDisplayName     *string                                    `json:"environmentDisplayName,omitempty"`
    PromotionTargetEnvironment *DeploymentDetailsResponsePromotionTarget  `json:"promotionTargetEnvironment,omitempty"`
}
```

#### Frontend (TypeScript) - `api/deployments.ts` ✅ **ALIGNED**
```typescript
export interface DeploymentDetailsResponse {
  imageId: string;
  status: string;
  lastDeployed: string; // ISO date-time
  endpoints: DeploymentEndpoint[];
  sourceEnvironment: EnvironmentObject;
  environmentDisplayName?: string;
  promotionTargetEnvironment?: PromotionTargetEnvironment;
}
```

**Status**: Already aligned ✅

---

### Common Types

#### Repository Config

**Backend (Go)**:
```go
type RepositoryConfig struct {
    Url     string `json:"url"`
    Branch  string `json:"branch"`
    AppPath string `json:"appPath"`
}
```

**Frontend (TypeScript)**: ✅ **ALIGNED**
```typescript
export interface RepositoryConfig {
  url: string;
  branch: string;
  appPath: string;
}
```

---

#### Runtime Configuration (ContainerSpec)

**Backend (Go)**:
```go
type RuntimeConfiguration struct {
    Env             []EnvironmentVariable `json:"env,omitempty"`
    RunCommand      *string               `json:"runCommand,omitempty"`
    LanguageVersion *string               `json:"languageVersion,omitempty"`
}
```

**Frontend (TypeScript)**: ✅ **ALIGNED**
```typescript
export interface ContainerSpec {
  runCommand?: string;
  env?: EnvironmentVariable[];
  languageVersion?: string;
}
```

**Note**: Frontend uses `ContainerSpec` name, backend uses `RuntimeConfiguration` - functionally equivalent

---

#### Environment Variable

**Backend (Go)**:
```go
type EnvironmentVariable struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
```

**Frontend (TypeScript)**: ✅ **ALIGNED**
```typescript
export interface EnvironmentVariable {
  key: string;
  value: string;
}
```

---

#### Input Interface

**Backend (Go)**:
```go
type InputInterface struct {
    Type              string        `json:"type"`
    CustomOpenAPISpec *EndpointSpec `json:"customOpenAPISpec,omitempty"`
}
```

**Frontend (TypeScript)**: ✅ **ALIGNED**
```typescript
export interface InputInterface {
  type: string;
  customOpenAPISpec?: EndpointSpec;
}
```

---

#### Endpoint Spec

**Backend (Go)**:
```go
type EndpointSpec struct {
    Port     int32          `json:"port"`
    Schema   EndpointSchema `json:"schema"`
    BasePath string         `json:"basePath"`
}
```

**Frontend (TypeScript)**: ✅ **ALIGNED**
```typescript
export interface EndpointSpec {
  port: number; // 1 - 65535
  schema: EndpointSchema;
  basePath: string;
}
```

---

## Type Alignment Summary

| Type Category | Backend (Go) | Frontend (TS) | Status |
|--------------|--------------|---------------|--------|
| **Agent Response** | `spec/model_agent_response.go` | `api/agents.ts` | ✅ **UPDATED** |
| **Create Agent Request** | `spec/model_create_agent_request.go` | `api/agents.ts` | ✅ **UPDATED** |
| **Agent List Response** | `spec/model_agent_list_response.go` | `api/agents.ts` | ✅ **UPDATED** |
| **Build Response** | `spec/model_build_response.go` | `api/builds.ts` | ✅ **UPDATED** |
| **Build Step** | `spec/model_build_step.go` | `api/builds.ts` | ✅ **UPDATED** |
| **Build Details Response** | `spec/model_build_details_response.go` | `api/builds.ts` | ✅ ALIGNED |
| **Deployment Response** | `spec/model_deployment_response.go` | `api/deployments.ts` | ✅ ALIGNED |
| **Deployment Details** | `spec/model_deployment_details_response.go` | `api/deployments.ts` | ✅ ALIGNED |
| **Repository Config** | `spec/model_repository_config.go` | `api/common.ts` | ✅ ALIGNED |
| **Runtime Configuration** | `spec/model_runtime_configuration.go` | `api/common.ts` | ✅ ALIGNED |
| **Environment Variable** | `spec/model_environment_variable.go` | `api/common.ts` | ✅ ALIGNED |
| **Input Interface** | `spec/model_input_interface.go` | `api/agents.ts` | ✅ ALIGNED |
| **Endpoint Spec** | `spec/model_endpoint_spec.go` | `api/common.ts` | ✅ ALIGNED |

---

## Key Differences & Design Decisions

### 1. **Date/Time Handling**
- **Backend**: Uses Go's `time.Time` type
- **Frontend**: Uses ISO 8601 string format (`string // ISO date-time`)
- **Rationale**: JSON serialization naturally converts to ISO strings; frontend can parse when needed

### 2. **Optional Fields**
- **Backend**: Uses pointer types (`*string`, `*int32`) with `omitempty` JSON tag
- **Frontend**: Uses TypeScript's optional syntax (`field?: type`)
- **Rationale**: Different language idioms for expressing optionality

### 3. **Type Flexibility**
- **Backend**: Strongly typed with exact string values
- **Frontend**: More flexible (e.g., `string` instead of `'Value1' | 'Value2'`)
- **Rationale**: 
  - Frontend is more permissive to handle backend evolution
  - Specific union types can be added when validation is critical
  - Reduces breaking changes when backend adds new enum values

### 4. **Naming Conventions**
- **Backend**: PascalCase for types (Go convention)
- **Frontend**: camelCase for properties, PascalCase for interfaces (TypeScript convention)
- **JSON Fields**: Both use camelCase (API contract)

---

## Synchronization Workflow

### When Backend Types Change:

1. **Update OpenAPI Spec** (if applicable)
2. **Regenerate Go Types**:
   ```bash
   cd agent-manager-service
   make spec
   ```

3. **Review Generated Types** in `spec/` directory

4. **Update Frontend Types**:
   ```bash
   cd console/workspaces/libs/types
   ```
   - Manually update `src/api/*.ts` files to match
   - Keep optional fields aligned
   - Maintain ISO date-time string format

5. **Build Types Package**:
   ```bash
   rushx build
   ```

6. **Update Dependent Packages**:
   ```bash
   cd ../../..
   rush build
   ```

### Best Practices:

1. **Use Backend as Source of Truth**: Always reference Go `spec/` types when making changes
2. **Document Differences**: If frontend types diverge intentionally, document why
3. **Test API Contracts**: Ensure both sides can serialize/deserialize successfully
4. **Version Compatibility**: Keep backwards compatibility in mind
5. **Type Comments**: Add comments explaining complex types or business logic

---

## Files Modified

### Console (Frontend) ✅
1. **`console/workspaces/libs/types/src/api/agents.ts`**
   - Restructured `AgentResponse` to match backend
   - Added `Provisioning` interface
   - Updated `CreateAgentRequest` structure
   - Updated `AgentListResponse` to use new `AgentResponse`

2. **`console/workspaces/libs/types/src/api/builds.ts`**
   - Changed `BuildResponse.status` to `string` (from union type)
   - Changed `BuildStep.type` and `BuildStep.status` to `string` (from union types)

### Backend (Go) - No changes needed ✅
- All types in `agent-manager-service/spec/` are auto-generated and correct

---

## Next Steps

1. ✅ **Types Updated**: Agent and Build types synchronized
2. 🔄 **Build & Test**: 
   ```bash
   cd console/workspaces/libs/types
   rushx build
   ```
3. 🔄 **Rebuild Dependent Packages**:
   ```bash
   cd ../../../..
   rush build
   ```
4. 🧪 **Test API Integration**: Verify actual API calls work with updated types
5. 📝 **Update Components**: Check if any React components need updates due to type changes

---

## Maintenance Notes

- **Auto-generated Code**: Never manually edit files in `agent-manager-service/spec/`
- **Type Safety**: TypeScript types should mirror Go types as closely as possible
- **API Versioning**: Consider API versioning strategy for breaking changes
- **Documentation**: Keep this document updated when types change significantly


