#!/bin/bash
set -e

# ============================================================================
# Agent Management Platform - Dev Container Runner
# ============================================================================
# Simple wrapper script to run the dev container with proper configuration
#
# Usage:
#   ./docker-run.sh              # Full installation
#   ./docker-run.sh --minimal    # Minimal installation
#   ./docker-run.sh --shell      # Interactive shell
#   ./docker-run.sh --build      # Build the image first
# ============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
IMAGE_NAME="agent-management-platform"
IMAGE_TAG="latest"
CONTAINER_NAME="amp-devcontainer"

# Default settings
BUILD_IMAGE="false"
DETACHED="false"
INTERACTIVE="true"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RESET='\033[0m'

# Helper functions
log_info() {
    echo -e "${BLUE}[INFO]${RESET} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${RESET} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${RESET} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${RESET} $1"
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        --build|-b)
            BUILD_IMAGE="true"
            shift
            ;;
        --detach|-d)
            DETACHED="true"
            INTERACTIVE="false"
            shift
            ;;
        --name)
            CONTAINER_NAME="$2"
            shift 2
            ;;
        --help|-h)
            cat << EOF

🐳 Agent Management Platform - Dev Container Runner

Run the complete platform in a single Docker container using the host's Docker daemon.

Usage:
  $0 [OPTIONS]

Options:
  --build, -b          Build the Docker image before running
  --detach, -d         Run container in detached mode (background)
  --name NAME          Custom container name (default: amp-devcontainer)
  --help, -h           Show this help message

Examples:
  # Build and run (interactive shell)
  $0 --build

  # Run existing image
  $0

  # Run in background
  $0 --detach

  # Custom container name
  $0 --name my-container

Prerequisites:
  • Docker running on host
  • At least 4GB RAM allocated to Docker
  • Ports available: 3000, 8080, 9098, 21893, 8443

Once Inside Container:
  • Run: ./bootstrap.sh
  • Or:  ./bootstrap.sh --minimal
  • Help: ./bootstrap.sh --help

Access After Installation:
  • Console:         http://localhost:3000
  • API:             http://localhost:8080
  • Traces Observer: http://localhost:9098
  • Data Prepper:    http://localhost:21893
  • Gateway:         https://localhost:8443

EOF
            exit 0
            ;;
        *)
            log_error "Unknown option: $1"
            echo "Use --help for usage information"
            exit 1
            ;;
    esac
done

# ============================================================================
# Pre-flight Checks
# ============================================================================

log_info "Running pre-flight checks..."

# Check if Docker is running
if ! docker info >/dev/null 2>&1; then
    log_error "Docker is not running"
    echo ""
    echo "   Please start Docker Desktop or Docker daemon"
    echo ""
    exit 1
fi

# Check Docker resources
DOCKER_MEM=$(docker info --format '{{.MemTotal}}' 2>/dev/null || echo 0)
DOCKER_MEM_GB=$((DOCKER_MEM / 1024 / 1024 / 1024))

if [ "$DOCKER_MEM_GB" -lt 4 ]; then
    log_warning "Docker has ${DOCKER_MEM_GB}GB RAM allocated"
    echo ""
    echo "   Recommended: 4GB+ for stable operation"
    echo "   Increase in Docker Desktop → Settings → Resources"
    echo ""
    read -p "   Continue anyway? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Check if ports are available
PORTS=(3000 8080 9098 21893)
PORTS_IN_USE=()

for port in "${PORTS[@]}"; do
    if lsof -Pi :$port -sTCP:LISTEN -t >/dev/null 2>&1; then
        PORTS_IN_USE+=($port)
    fi
done

if [ ${#PORTS_IN_USE[@]} -gt 0 ]; then
    log_warning "The following ports are already in use: ${PORTS_IN_USE[*]}"
    echo ""
    echo "   The container will start but services may not be accessible"
    echo "   Consider stopping services using these ports"
    echo ""
    read -p "   Continue anyway? (y/N) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

log_success "Pre-flight checks passed"
echo ""

# ============================================================================
# Build Image if Requested
# ============================================================================

if [ "$BUILD_IMAGE" = "true" ]; then
    log_info "Building Docker image: ${IMAGE_NAME}:${IMAGE_TAG}"
    echo ""
    
    cd "$SCRIPT_DIR"
    
    if docker build -f Dockerfile.devcontainer -t "${IMAGE_NAME}:${IMAGE_TAG}" .; then
        log_success "Image built successfully"
        echo ""
    else
        log_error "Failed to build Docker image"
        exit 1
    fi
fi

# Check if image exists
if ! docker image inspect "${IMAGE_NAME}:${IMAGE_TAG}" >/dev/null 2>&1; then
    log_error "Docker image not found: ${IMAGE_NAME}:${IMAGE_TAG}"
    echo ""
    echo "   Build the image first:"
    echo "   $0 --build"
    echo ""
    exit 1
fi

# ============================================================================
# Stop Existing Container
# ============================================================================

if docker ps -a --format '{{.Names}}' | grep -q "^${CONTAINER_NAME}$"; then
    log_warning "Container '${CONTAINER_NAME}' already exists"
    echo ""
    read -p "   Remove and recreate? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        log_info "Stopping and removing existing container..."
        docker stop "$CONTAINER_NAME" >/dev/null 2>&1 || true
        docker rm "$CONTAINER_NAME" >/dev/null 2>&1 || true
        log_success "Existing container removed"
        echo ""
    else
        log_error "Cannot proceed with existing container"
        exit 1
    fi
fi

# ============================================================================
# Run Container
# ============================================================================

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
log_info "Starting dev container..."
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Configuration:"
echo "  • Container: $CONTAINER_NAME"
echo "  • Interactive: $INTERACTIVE"
echo ""

if [ "$DETACHED" = "true" ]; then
    echo "  • Running in detached mode"
    log_info "Starting container in background..."
else
    echo "  • Running in interactive mode"
    log_info "Starting interactive container..."
fi
echo ""

# Build docker run command
# Note: We mount the Docker socket instead of using --privileged
# This is more secure and follows Docker best practices
DOCKER_RUN_ARGS=(
    "run"
    "--rm"
    "--name" "$CONTAINER_NAME"
    "-v" "/var/run/docker.sock:/var/run/docker.sock"
    "-p" "3000:3000"
    "-p" "8080:8080"
    "-p" "9098:9098"
    "-p" "21893:21893"
    "-p" "8443:8443"
)

if [ "$INTERACTIVE" = "true" ]; then
    DOCKER_RUN_ARGS+=("-it")
fi

if [ "$DETACHED" = "true" ]; then
    DOCKER_RUN_ARGS+=("-d")
fi

DOCKER_RUN_ARGS+=(
    "${IMAGE_NAME}:${IMAGE_TAG}"
)

# Run the container
if docker "${DOCKER_RUN_ARGS[@]}"; then
    if [ "$DETACHED" = "true" ]; then
        echo ""
        log_success "Container started in background"
        echo ""
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        echo "✅ Dev Container Running"
        echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
        echo ""
        echo "📊 Useful commands:"
        echo ""
        echo "   Enter shell:    docker exec -it $CONTAINER_NAME bash"
        echo "   View logs:      docker logs -f $CONTAINER_NAME"
        echo "   Stop:           docker stop $CONTAINER_NAME"
        echo ""
        echo "🚀 Inside the container, run:"
        echo "   ./bootstrap.sh"
        echo ""
        echo "🌐 After installation, access at:"
        echo "   • Console:         http://localhost:3000"
        echo "   • API:             http://localhost:8080"
        echo "   • Traces Observer: http://localhost:9098"
        echo ""
    fi
else
    log_error "Failed to start container"
    exit 1
fi

