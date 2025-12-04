#!/usr/bin/env bash

# Check status of all required services

AMP_NS="${AMP_NS:-agent-management-platform}"
OBSERVABILITY_NS="${OBSERVABILITY_NS:-openchoreo-observability-plane}"
DATA_PLANE_NS="${DATA_PLANE_NS:-openchoreo-data-plane}"

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Service Status Check"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

echo "Checking Agent Management Platform services..."
echo ""

# Console (3000)
echo -n "Console (3000): "
if kubectl get svc agent-management-platform-console -n "$AMP_NS" >/dev/null 2>&1; then
    echo "✓ Service exists"
    kubectl get svc agent-management-platform-console -n "$AMP_NS" | tail -n 1
else
    echo "✗ Service not found"
fi
echo ""

# Agent Manager (8080)
echo -n "Agent Manager (8080): "
if kubectl get svc agent-management-platform-agent-manager-service -n "$AMP_NS" >/dev/null 2>&1; then
    echo "✓ Service exists"
    kubectl get svc agent-management-platform-agent-manager-service -n "$AMP_NS" | tail -n 1
else
    echo "✗ Service not found"
fi
echo ""

# Traces Observer (9098)
echo -n "Traces Observer (9098): "
if kubectl get svc traces-observer-service -n "$OBSERVABILITY_NS" >/dev/null 2>&1; then
    echo "✓ Service exists"
    kubectl get svc traces-observer-service -n "$OBSERVABILITY_NS" | tail -n 1
else
    echo "✗ Service not found"
fi
echo ""

# Data Prepper (21893)
echo -n "Data Prepper (21893): "
if kubectl get svc data-prepper -n "$OBSERVABILITY_NS" >/dev/null 2>&1; then
    echo "✓ Service exists"
    kubectl get svc data-prepper -n "$OBSERVABILITY_NS" | tail -n 1
else
    echo "✗ Service not found"
fi
echo ""

# Gateway (8443)
echo -n "External Gateway (8443): "
if kubectl get svc gateway-external -n "$DATA_PLANE_NS" >/dev/null 2>&1; then
    echo "✓ Service exists"
    kubectl get svc gateway-external -n "$DATA_PLANE_NS" | tail -n 1
else
    echo "✗ Service not found"
fi
echo ""

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Pod Status"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
kubectl get pods -n "$AMP_NS"
echo ""

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Port Forward Status"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""
ps aux | grep 'port-forward' | grep -v grep || echo "No port-forward processes running"
echo ""
