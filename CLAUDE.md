# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

WSO2 AI Agent Management Platform is an enterprise control plane for deploying, managing, and governing AI agents at scale. Built on OpenChoreo for internal agent deployments with OpenTelemetry-based instrumentation.

**Architecture**: Multi-service platform with Go backend (Agent Manager Service), Go traces observer, React/TypeScript console, and Kubernetes/Helm deployment infrastructure.

## Repository Structure

```
ai-agent-management-platform/
├── agent-manager-service/      # Go backend - main control plane API
├── traces-observer-service/    # Go service - OpenSearch traces query API
├── console/                    # React/TypeScript frontend (Rush monorepo)
├── deployments/                # Docker Compose, Helm charts, K8s config
├── quick-start/                # Installation scripts and guides
└── samples/                    # Example agent implementations
```

## Development Commands

### Complete Platform Setup (One-Time)

From repository root:

```bash
# Complete local development setup (Colima + Kind + OpenChoreo + Platform)
make setup

# Daily development workflow
make dev-up           # Start all services (console, API, database)
make dev-down         # Stop all services
make dev-restart      # Restart services
make dev-rebuild      # Rebuild Docker images and restart
make dev-logs         # View all service logs
make dev-migrate      # Run database migrations
```

**Access URLs after `make dev-up`:**
- Console: http://localhost:3000
- API: http://localhost:8080
- Traces Observer: http://localhost:9098
- Database: localhost:5432

### Agent Manager Service (Go)

Working directory: `agent-manager-service/`

**Requirements**: Go 1.25+, PostgreSQL 12+, air (hot-reload), moq (mocking)

```bash
# Development
make run              # Start with hot-reload (uses air)
make test             # Run tests
make fmt              # Format code
make lint             # Run linters

# Code generation (run after API spec or wiring changes)
make wire             # Generate dependency injection code
make spec             # Generate models from OpenAPI spec (docs/api_v1_openapi.yaml)
make codegen          # Generate all code (wire + models)

# Verify before commit
make codegenfmt-check # Ensure codegen + formatting is up-to-date
```

**Database migrations**: Located in `db_migrations/`. Run with `go run . -migrate` or `make dev-migrate` (Docker).

**API Documentation**: OpenAPI 3.0 spec at `docs/api_v1_openapi.yaml`

**Architecture**:
- **api/**: HTTP handlers and routing
- **controllers/**: Request controllers
- **services/**: Business logic layer
- **repositories/**: Data access layer
- **wiring/**: Dependency injection (Wire framework)
- **clients/**: External service clients (OpenChoreo, OpenSearch)
- **middleware/**: Auth, logging, recovery

### Traces Observer Service (Go)

Working directory: `traces-observer-service/`

Queries traces from OpenSearch for the console. **Note**: Planned migration to OpenChoreo Observability Plane Observer.

```bash
# Local development
go run .              # Start service
make build            # Build binary
make run              # Build and run

# Docker
make docker-build     # Build image
make docker-run       # Run container
make docker-load-kind # Load image into Kind cluster
```

**Configuration**: Via environment variables (see traces-observer-service/README.MD)
- `TRACES_OBSERVER_PORT`: Service port (default: 9098)
- `OPENSEARCH_ADDRESS`: OpenSearch endpoint
- `OPENSEARCH_USERNAME`, `OPENSEARCH_PASSWORD`: Auth credentials
- `OPENSEARCH_TRACE_INDEX`: Trace index name

### Console (React/TypeScript)

Working directory: `console/`

**Tech Stack**: React 19, TypeScript, Vite, Rush monorepo, pnpm

**Requirements**: Node.js 18.20.3+ or 20.14.0+, Rush 5.157.0

```bash
# Install Rush globally (if not installed)
npm install -g @microsoft/rush@5.157.0

# Dependencies and build
rush install          # Install all dependencies
rush update           # Update dependencies (when rush.json changes)
rush build            # Build all packages
rush purge            # Clean all build outputs

# Development
rushx --to @agent-management-platform/webapp dev  # Start dev server

# Or from apps/webapp:
cd apps/webapp
rushx dev             # Start dev server (http://localhost:5173)
rushx build           # Production build
rushx lint            # Run linting
rushx lint:fix        # Fix linting issues
rushx preview         # Preview production build

# Create new page component
rush create-page      # Interactive generator
```

**Monorepo Structure**:
- **apps/webapp**: Main application
- **workspaces/libraries/**: Shared libraries (auth, types, api-client, views, eslint-config)
- **workspaces/pages/**: Page components (use .template for new pages)

**Configuration**: Edit `apps/webapp/public/config.js` to set API_URL.

### Deployment & Infrastructure

Working directory: `deployments/`

**Docker Compose** (local development):
```bash
cd deployments
docker compose up -d       # Start services
docker compose down        # Stop services
docker compose logs -f     # View logs
```

**OpenChoreo Cluster Management**:
```bash
make openchoreo-up         # Start OpenChoreo Kind cluster
make openchoreo-down       # Stop cluster (preserves containers)
make openchoreo-status     # Check cluster status
make port-forward          # Forward OpenChoreo services to localhost
```

**Debugging**:
```bash
make db-connect            # Connect to PostgreSQL
make db-logs               # View database logs
make service-logs          # View agent-manager-service logs
make service-shell         # Shell into service container
make console-logs          # View console logs
```

**Helm Charts**: Located in `deployments/helm-charts/`
- `wso2-ai-agent-management-platform`: Main platform
- `wso2-amp-build-extension`: Build extension for OpenChoreo
- `wso2-amp-observability-extension`: Observability stack extension

**Installation on Existing Cluster**: See `quick-start/README.md` for installing on existing OpenChoreo clusters.

### Cleanup

```bash
make teardown              # Remove Kind cluster + platform
```

## Key Conventions

### Go Services (agent-manager-service, traces-observer-service)

- **Dependency Injection**: Use Wire framework (`wiring/` directory)
- **After API spec changes**: Run `make spec` then `make codegen`
- **Before committing Go code**: Run `make codegenfmt-check` to verify code generation and formatting
- **Environment variables**: Defined in `.env` files (see service READMEs)
- **Testing**: Standard Go testing with moq for mocks

### Console (React/TypeScript)

- **Rush commands**: Always use `rush` or `rushx` (never `npm` or `pnpm` directly)
- **New pages**: Use `rush create-page` generator, then add to `rush.json` projects list
- **Shared code**: Place in `workspaces/libraries/`, not `apps/webapp`
- **After rush.json changes**: Run `rush update`

### Docker Development

- **Kubeconfig**: Services use `~/.kube/config-docker` (generated by setup scripts)
- **Hot-reload**: Source code is volume-mounted for live updates
- **Network**: Services connect to `kind` network for OpenChoreo access

## Important Files

- **Makefile** (root): Main development commands
- **agent-manager-service/docs/api_v1_openapi.yaml**: API specification
- **console/rush.json**: Monorepo configuration
- **deployments/docker-compose.yml**: Local development stack
- **deployments/kind-config.yaml**: Kind cluster configuration
- **quick-start/install.sh**: Platform installation script

## Release Process

Releases are managed via GitHub workflows (see `github-workflows.md`):

**Release Tag Format**: `amp/v{version}` (e.g., `amp/v0.1.0`)

**Components Released**:
- Docker images: `amp-console`, `amp-api`, `amp-trace-observer`, `amp-python-instrumentation-provider`, `amp-quick-start`
- Helm charts: All charts in `deployments/helm-charts/`
- Quick-start archive: Packaged `quick-start/` directory

**Image Registry**: `ghcr.io/wso2`
