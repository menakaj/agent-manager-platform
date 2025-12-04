#!/usr/bin/env bash
# check-resources.sh - Check Docker and system resources for agent-management-platform
set -eo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging functions
log_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

log_success() {
    echo -e "${GREEN}✓${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

log_error() {
    echo -e "${RED}✗${NC} $1"
}

# Recommended minimums
MIN_MEMORY_GB=4
MIN_CPU=2
RECOMMENDED_MEMORY_GB=8
RECOMMENDED_CPU=4

echo ""
echo "======================================"
echo "  Docker Resource Check"
echo "======================================"
echo ""

# Check if Docker is running
if ! docker info &>/dev/null; then
    log_error "Docker is not running or not accessible"
    echo ""
    echo "Please start Docker Desktop and try again."
    exit 1
fi

log_success "Docker is running"
echo ""

# Get Docker resources
DOCKER_INFO=$(docker info --format '{{json .}}')
MEM_TOTAL=$(echo "$DOCKER_INFO" | jq -r '.MemTotal')
NCPU=$(echo "$DOCKER_INFO" | jq -r '.NCPU')
OS=$(echo "$DOCKER_INFO" | jq -r '.OperatingSystem')
VERSION=$(echo "$DOCKER_INFO" | jq -r '.ServerVersion')

# Convert memory to GB
MEM_GB=$(echo "scale=2; $MEM_TOTAL / 1024 / 1024 / 1024" | bc)

echo "Current Docker Configuration:"
echo "------------------------------------"
echo "  Memory:         ${MEM_GB} GB"
echo "  CPUs:           ${NCPU}"
echo "  OS:             ${OS}"
echo "  Docker Version: ${VERSION}"
echo ""

# Check memory
echo "Resource Assessment:"
echo "------------------------------------"

# Memory check
if (( $(echo "$MEM_GB >= $RECOMMENDED_MEMORY_GB" | bc -l) )); then
    log_success "Memory: ${MEM_GB} GB (Excellent - meets recommended ${RECOMMENDED_MEMORY_GB}+ GB)"
elif (( $(echo "$MEM_GB >= $MIN_MEMORY_GB" | bc -l) )); then
    log_warning "Memory: ${MEM_GB} GB (Acceptable - minimum ${MIN_MEMORY_GB} GB met, but ${RECOMMENDED_MEMORY_GB}+ GB recommended)"
else
    log_error "Memory: ${MEM_GB} GB (Insufficient - minimum ${MIN_MEMORY_GB} GB required)"
    INSUFFICIENT=true
fi

# CPU check
if (( NCPU >= RECOMMENDED_CPU )); then
    log_success "CPUs: ${NCPU} (Excellent - meets recommended ${RECOMMENDED_CPU}+ CPUs)"
elif (( NCPU >= MIN_CPU )); then
    log_warning "CPUs: ${NCPU} (Acceptable - minimum ${MIN_CPU} CPUs met, but ${RECOMMENDED_CPU}+ CPUs recommended)"
else
    log_error "CPUs: ${NCPU} (Insufficient - minimum ${MIN_CPU} CPUs required)"
    INSUFFICIENT=true
fi

echo ""

# Check disk space
echo "Disk Space:"
echo "------------------------------------"
DOCKER_ROOT=$(echo "$DOCKER_INFO" | jq -r '.DockerRootDir // "/var/lib/docker"')
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS - check available space on root
    DISK_AVAIL=$(df -h / | awk 'NR==2 {print $4}')
    log_info "Available disk space: ${DISK_AVAIL}"
else
    # Linux - check Docker root directory
    DISK_AVAIL=$(df -h "$DOCKER_ROOT" | awk 'NR==2 {print $4}')
    log_info "Available disk space (Docker root): ${DISK_AVAIL}"
fi

echo ""

# Summary and recommendations
echo "======================================"
echo "  Summary"
echo "======================================"
echo ""

if [[ "${INSUFFICIENT}" == "true" ]]; then
    log_error "Docker resources are insufficient for running the platform"
    echo ""
    echo "To increase Docker resources:"
    echo ""
    if [[ "$OSTYPE" == "darwin"* ]]; then
        echo "  macOS (Docker Desktop):"
        echo "  1. Open Docker Desktop"
        echo "  2. Go to Settings (⚙️) → Resources"
        echo "  3. Adjust Memory to at least ${MIN_MEMORY_GB} GB (${RECOMMENDED_MEMORY_GB}+ GB recommended)"
        echo "  4. Adjust CPUs to at least ${MIN_CPU} (${RECOMMENDED_CPU}+ recommended)"
        echo "  5. Click 'Apply & Restart'"
    else
        echo "  Linux:"
        echo "  - Docker on Linux uses host resources directly"
        echo "  - Ensure your system has at least ${MIN_MEMORY_GB} GB RAM and ${MIN_CPU} CPUs"
        echo "  - For Docker Desktop on Linux, adjust in Settings → Resources"
    fi
    echo ""
    exit 1
elif (( $(echo "$MEM_GB < $RECOMMENDED_MEMORY_GB" | bc -l) )) || (( NCPU < RECOMMENDED_CPU )); then
    log_warning "Docker resources meet minimum requirements but are below recommended"
    echo ""
    echo "Recommended configuration:"
    echo "  - Memory: ${RECOMMENDED_MEMORY_GB}+ GB (current: ${MEM_GB} GB)"
    echo "  - CPUs: ${RECOMMENDED_CPU}+ (current: ${NCPU})"
    echo ""
    echo "The platform will run, but you may experience slower performance."
    echo "Consider increasing resources for optimal experience."
    echo ""
    exit 0
else
    log_success "Docker resources are optimal for running the platform"
    echo ""
    echo "Your Docker configuration meets all recommended requirements."
    echo "You can proceed with the installation."
    echo ""
    exit 0
fi

