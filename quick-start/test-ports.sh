#!/usr/bin/env bash

# Port Testing Script for Agent Management Platform
# Tests all exposed ports to verify they are accessible and responding

set -eo pipefail

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
RESET='\033[0m'

log_info() {
    echo -e "${BLUE}[INFO]${RESET} $1"
}

log_success() {
    echo -e "${GREEN}[✓]${RESET} $1"
}

log_warning() {
    echo -e "${YELLOW}[⚠]${RESET} $1"
}

log_error() {
    echo -e "${RED}[✗]${RESET} $1"
}

# Test a TCP port connection
test_tcp_port() {
    local port="$1"
    local description="$2"
    local timeout="${3:-3}"

    if nc -z -w "$timeout" localhost "$port" 2>/dev/null; then
        return 0
    else
        return 1
    fi
}

# Test an HTTP endpoint
test_http_endpoint() {
    local url="$1"
    local description="$2"
    local expected_pattern="${3:-.*}"
    local timeout="${4:-5}"

    local response
    response=$(curl -s -m "$timeout" "$url" 2>&1)
    local exit_code=$?

    if [ $exit_code -eq 0 ]; then
        if echo "$response" | grep -q "$expected_pattern"; then
            return 0
        else
            # Got response but not expected pattern
            return 2
        fi
    else
        return 1
    fi
}

# Test OTLP endpoint (expects connection but may return error for invalid data)
test_otlp_endpoint() {
    local url="$1"
    local description="$2"
    local timeout="${3:-3}"

    # Send empty POST request - should get connection established
    local response
    response=$(curl -v -X POST "$url" \
        -H "Content-Type: application/json" \
        -m "$timeout" 2>&1)

    # Check if connection was established
    if echo "$response" | grep -q "Connected to"; then
        return 0
    else
        return 1
    fi
}

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "   Agent Management Platform - Port Test"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

# Check if nc (netcat) is available
if ! command -v nc >/dev/null 2>&1; then
    log_warning "netcat (nc) not found - TCP tests will be skipped"
    HAS_NC=false
else
    HAS_NC=true
fi

# Check if curl is available
if ! command -v curl >/dev/null 2>&1; then
    log_error "curl is required but not found"
    exit 1
fi

TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Test Console (3000)
echo "Testing Console (port 3000)..."
TOTAL_TESTS=$((TOTAL_TESTS + 1))
if test_http_endpoint "http://localhost:3000" "Console" "<!DOCTYPE html>|<html" 5; then
    log_success "Console is accessible at http://localhost:3000"
    PASSED_TESTS=$((PASSED_TESTS + 1))
elif [ $? -eq 2 ]; then
    log_warning "Console port 3000 is accessible but returned unexpected response"
    PASSED_TESTS=$((PASSED_TESTS + 1))
else
    log_error "Console is NOT accessible at http://localhost:3000"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
echo ""

# Test Agent Manager Service (8080 -> 8081 mapped)
echo "Testing Agent Manager Service (port 8080)..."
TOTAL_TESTS=$((TOTAL_TESTS + 1))
if test_http_endpoint "http://localhost:8080/health" "Agent Manager" "UP|status|healthy" 5; then
    log_success "Agent Manager Service is accessible at http://localhost:8080"
    PASSED_TESTS=$((PASSED_TESTS + 1))
elif test_tcp_port 8080 "Agent Manager" 3 2>/dev/null && [ "$HAS_NC" = true ]; then
    log_warning "Agent Manager port 8080 is accessible but health endpoint returned unexpected response"
    PASSED_TESTS=$((PASSED_TESTS + 1))
else
    log_error "Agent Manager Service is NOT accessible at http://localhost:8080"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
echo ""

# Test Traces Observer Service (9098)
echo "Testing Traces Observer Service (port 9098)..."
TOTAL_TESTS=$((TOTAL_TESTS + 1))
if test_http_endpoint "http://localhost:9098" "Traces Observer" ".*" 5; then
    log_success "Traces Observer Service is accessible at http://localhost:9098"
    PASSED_TESTS=$((PASSED_TESTS + 1))
elif test_tcp_port 9098 "Traces Observer" 3 2>/dev/null && [ "$HAS_NC" = true ]; then
    log_warning "Traces Observer port 9098 is accessible (TCP connection OK)"
    PASSED_TESTS=$((PASSED_TESTS + 1))
else
    log_error "Traces Observer Service is NOT accessible at http://localhost:9098"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
echo ""

# Test Data Prepper (21893)
echo "Testing Data Prepper OTLP endpoint (port 21893)..."
TOTAL_TESTS=$((TOTAL_TESTS + 1))
if test_otlp_endpoint "http://localhost:21893/v1/traces" "Data Prepper" 5; then
    log_success "Data Prepper is accessible at http://localhost:21893"
    PASSED_TESTS=$((PASSED_TESTS + 1))
elif test_tcp_port 21893 "Data Prepper" 3 2>/dev/null && [ "$HAS_NC" = true ]; then
    log_warning "Data Prepper port 21893 is accessible (TCP connection OK, but HTTP may be failing)"
    log_info "Testing direct TCP connection to port 21893..."
    if curl -v http://localhost:21893/v1/traces -X POST -m 3 2>&1 | grep -q "Connected to"; then
        log_success "TCP connection to Data Prepper successful"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    else
        log_error "TCP connection to Data Prepper failed"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
else
    log_error "Data Prepper is NOT accessible at http://localhost:21893"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
echo ""

# Test External Gateway (8443 -> 8444 mapped)
echo "Testing External Gateway (port 8444)..."
TOTAL_TESTS=$((TOTAL_TESTS + 1))
if curl -k -s -m 5 https://localhost:8444 >/dev/null 2>&1; then
    log_success "External Gateway is accessible at https://localhost:8444"
    PASSED_TESTS=$((PASSED_TESTS + 1))
elif test_tcp_port 8444 "External Gateway" 3 2>/dev/null && [ "$HAS_NC" = true ]; then
    log_warning "External Gateway port 8444 is accessible (TCP connection OK)"
    PASSED_TESTS=$((PASSED_TESTS + 1))
else
    log_error "External Gateway is NOT accessible at https://localhost:8444"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi
echo ""

# Summary
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "   Test Summary"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
echo "Total tests:  $TOTAL_TESTS"
echo -e "${GREEN}Passed:       $PASSED_TESTS${RESET}"
if [ $FAILED_TESTS -gt 0 ]; then
    echo -e "${RED}Failed:       $FAILED_TESTS${RESET}"
else
    echo "Failed:       $FAILED_TESTS"
fi
echo ""

if [ $FAILED_TESTS -eq 0 ]; then
    log_success "All port tests passed!"
    echo ""
    echo "Your services are accessible at:"
    echo "  • Console:           http://localhost:3000"
    echo "  • Agent Manager:     http://localhost:8081"
    echo "  • Traces Observer:   http://localhost:9098"
    echo "  • Data Prepper:      http://localhost:21893"
    echo "  • External Gateway:  https://localhost:8444"
    echo ""
    exit 0
else
    log_error "Some port tests failed"
    echo ""
    echo "Troubleshooting:"
    echo "  1. Check if port forwarding is running:"
    echo "     docker exec amp ps aux | grep socat"
    echo ""
    echo "  2. Restart port forwarding inside the container:"
    echo "     docker exec amp /app/port-forward.sh"
    echo ""
    echo "  3. Check logs:"
    echo "     docker logs amp"
    echo ""
    exit 1
fi
