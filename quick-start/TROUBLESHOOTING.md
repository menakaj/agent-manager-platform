# Troubleshooting Guide

Common issues and solutions for Agent Management Platform installation.

## Installation Issues

### Prerequisites Check Failed

**Error**: `OpenChoreo Observability Plane not found`

**Solution**:
```bash
# Install OpenChoreo Observability Plane
helm install observability-plane oci://ghcr.io/openchoreo/helm-charts/openchoreo-observability-plane \
  --version 0.3.2 \
  --namespace openchoreo-observability-plane \
  --create-namespace \
  --wait

# Verify installation
kubectl get pods -n openchoreo-observability-plane
```

**Documentation**: https://openchoreo.dev/docs/v0.3.x/observability/

---

### Helm Chart Pull Failed

**Error**: `Error: failed to download "oci://ghcr.io/..."`

**Possible Causes**:
1. Network connectivity issues
2. GHCR rate limiting
3. Invalid chart version

**Solution**:
```bash
# Check network connectivity
curl -I https://ghcr.io

# Verify Helm can access GHCR
helm pull oci://ghcr.io/openchoreo/helm-charts/cilium --version 0.3.2

# If rate limited, wait a few minutes and retry
./install.sh
```

---

### Pods Not Ready

**Error**: Pods stuck in `Pending`, `CrashLoopBackOff`, or `ImagePullBackOff`

**Solution**:
```bash
# Check pod status
kubectl get pods -n agent-management-platform
kubectl get pods -n openchoreo-observability-plane

# Describe problematic pod
kubectl describe pod <pod-name> -n agent-management-platform

# Check logs
kubectl logs <pod-name> -n agent-management-platform

# Common issues:
# 1. Insufficient resources - check cluster capacity
kubectl top nodes

# 2. Image pull issues - check image pull secrets
kubectl get events -n agent-management-platform --sort-by='.lastTimestamp'

# 3. Storage issues - check PVC status
kubectl get pvc -n agent-management-platform
```

---

### PostgreSQL Not Starting

**Error**: PostgreSQL pod in `CrashLoopBackOff` or `Pending`

**Solution**:
```bash
# Check PostgreSQL pod
kubectl get pods -n agent-management-platform -l app.kubernetes.io/name=postgresql

# Check PVC
kubectl get pvc -n agent-management-platform

# If PVC is pending, check storage class
kubectl get storageclass

# If no default storage class, create one or specify in config
# See: https://kubernetes.io/docs/concepts/storage/storage-classes/
```

---

## Access Issues

### Port Forwarding Not Working

**Error**: Cannot access services at localhost

**Solution**:
```bash
# Check if port forwarding is running
ps aux | grep "kubectl port-forward"

# Stop existing port forwarding
./stop-port-forward.sh

# Restart port forwarding
./port-forward.sh

# Or manually forward each service
kubectl port-forward -n agent-management-platform svc/agent-management-platform-console 3000:3000 &
kubectl port-forward -n agent-management-platform svc/agent-management-platform-agent-manager-service 8080:8080 &
kubectl port-forward -n openchoreo-observability-plane svc/traces-observer-service 9098:9098 &
kubectl port-forward -n openchoreo-observability-plane svc/data-prepper 21893:21893 &
```

---

### Console Shows "API Connection Error"

**Error**: Console cannot connect to backend API

**Solution**:
```bash
# Check if Agent Manager Service is running
kubectl get pods -n agent-management-platform -l app=agent-manager-service

# Check service endpoints
kubectl get svc -n agent-management-platform agent-management-platform-agent-manager-service

# Verify port forwarding for API (8080)
curl http://localhost:8080/health

# Check CORS configuration
kubectl get apiclass default-with-cors -n default -o yaml

# If CORS issue, re-patch
kubectl patch apiclass default-with-cors -n default --type json \
  -p '[{"op":"add","path":"/spec/restPolicy/defaults/cors/allowOrigins/-","value":"http://localhost:3000"}]'
```

---

## Observability Issues

### Traces Not Appearing

**Error**: No traces visible in the observability dashboard

**Solution**:
```bash
# Check Data Prepper is running
kubectl get pods -n openchoreo-observability-plane -l app=data-prepper

# Check Traces Observer is running
kubectl get pods -n openchoreo-observability-plane -l app=traces-observer-service

# Check OpenSearch is accessible
kubectl get pods -n openchoreo-observability-plane -l app=opensearch

# Test Data Prepper endpoint
curl http://localhost:21893/metrics

# Check Data Prepper logs
kubectl logs -n openchoreo-observability-plane deployment/data-prepper

# Verify agent is sending traces (check agent logs)
```

---

### OpenSearch Connection Failed

**Error**: Data Prepper cannot connect to OpenSearch

**Solution**:
```bash
# Check OpenSearch status
kubectl get pods -n openchoreo-observability-plane -l app=opensearch

# Check OpenSearch service
kubectl get svc -n openchoreo-observability-plane opensearch

# Port forward to OpenSearch for testing
kubectl port-forward -n openchoreo-observability-plane svc/opensearch 9200:9200

# Test OpenSearch connection
curl http://localhost:9200/_cluster/health

# Check Data Prepper configuration
kubectl get configmap -n openchoreo-observability-plane data-prepper-config -o yaml
```

---

## Permission Issues

### Service Account Cannot Access OpenChoreo Resources

**Error**: `User "system:serviceaccount:agent-management-platform:agent-management-platform" cannot get resource "projects" in API group "openchoreo.dev"`

**Cause**: The service account doesn't have permissions to access OpenChoreo custom resources.

**Solution**:

The Helm chart should include OpenChoreo API group permissions. Verify RBAC rules:

```bash
# Check the ClusterRole
kubectl get clusterrole agent-management-platform -o yaml

# Should include openchoreo.dev API group
```

If missing, the chart needs to be upgraded with correct RBAC rules. The values should include:

```yaml
rbac:
  create: true
  rules:
    - apiGroups: ["*"]
      resources: ["*"]
      verbs: ["*"]
```

**Note**: The service account is granted full cluster permissions to manage all Kubernetes and OpenChoreo resources.

**Fix**:
```bash
# Upgrade the installation with correct RBAC
helm upgrade agent-management-platform \
  oci://ghcr.io/agent-mgt-platform/agent-management-platform \
  --version 0.1.0 \
  --namespace agent-management-platform \
  --reuse-values

# Or reinstall
./uninstall.sh --force
./install.sh
```

---

## Cluster Issues

### Insufficient Resources

**Error**: Pods pending due to insufficient CPU/memory

**Solution**:
```bash
# Check cluster resources
kubectl top nodes
kubectl describe nodes

# Check resource requests
kubectl get pods -n agent-management-platform -o json | \
  jq '.items[] | {name: .metadata.name, resources: .spec.containers[].resources}'

# Scale down replicas if needed (temporary)
kubectl scale deployment agent-manager-service -n agent-management-platform --replicas=1
kubectl scale deployment console -n agent-management-platform --replicas=1

# For production, increase cluster capacity
```

---

### Namespace Stuck in Terminating

**Error**: Cannot delete namespace, stuck in `Terminating` state

**Solution**:
```bash
# Check what's blocking deletion
kubectl api-resources --verbs=list --namespaced -o name | \
  xargs -n 1 kubectl get --show-kind --ignore-not-found -n agent-management-platform

# Force delete finalizers (use with caution)
kubectl get namespace agent-management-platform -o json | \
  jq '.spec.finalizers = []' | \
  kubectl replace --raw "/api/v1/namespaces/agent-management-platform/finalize" -f -
```

---

## Getting Help

### Verbose Installation

For detailed installation logs:
```bash
./install.sh --verbose
```

### Collect Diagnostic Information

```bash
# Get all resources
kubectl get all -n agent-management-platform
kubectl get all -n openchoreo-observability-plane

# Get events
kubectl get events -n agent-management-platform --sort-by='.lastTimestamp'
kubectl get events -n openchoreo-observability-plane --sort-by='.lastTimestamp'

# Get logs from all pods
kubectl logs -n agent-management-platform --all-containers=true --selector app=agent-manager-service
kubectl logs -n openchoreo-observability-plane --all-containers=true --selector app=data-prepper
```

### Report an Issue

If you're still experiencing issues:

1. Collect diagnostic information (see above)
2. Check existing issues: https://github.com/wso2/agent-management-platform/issues
3. Create a new issue with:
   - Installation command used
   - Error messages
   - Diagnostic information
   - Kubernetes version and cluster type

### Additional Resources

- [Quick Start Guide](./QUICK_START.md)
- [Detailed Installation Guide](./README.md)
- [OpenChoreo Documentation](https://openchoreo.dev/docs/)
- [GitHub Issues](https://github.com/wso2/agent-management-platform/issues)

