# Port Forwarding Fix Summary

## Problem
Port 21893 (and potentially other ports) were not accessible when using the container-based setup with Kind cluster on Colima/Docker.

## Root Causes Identified

1. **Wrong Method**: Using `kubectl port-forward` instead of `socat`
   - kubectl port-forward is fragile and dies on pod restarts
   - Requires pods to be running
   - More complex dependency chain

2. **Missing Services**: data-prepper and traces-observer services didn't exist
   - Services weren't deployed
   - Port forwarding failed silently

3. **Wrong Service Type**: Services were ClusterIP instead of NodePort
   - ClusterIP is only accessible within cluster
   - socat cannot forward to ClusterIP addresses

4. **kubectl Context Issue**: kubectl was pointing to rancher-desktop instead of kind cluster
   - Fixed by running: `kind export kubeconfig --name openchoreo-local`

## Solution Implemented

### Adopted OpenChoreo's Approach

Instead of using `kubectl port-forward`, we now use **socat + NodePort** (exactly like OpenChoreo):

```bash
# Old approach (broken)
kubectl port-forward --address=0.0.0.0 svc/data-prepper 21893:21893

# New approach (reliable)
nodeport=$(kubectl get svc data-prepper -n observability -o jsonpath='{.spec.ports[0].nodePort}')
socat TCP-LISTEN:21893,fork,reuseaddr TCP:openchoreo-local-worker:$nodeport &
```

### Files Modified

1. **quick-start/port-forward.sh** (complete rewrite)
   - Now uses socat approach
   - Dynamic NodePort discovery
   - Robust error handling
   - Waits for services to be ready

2. **quick-start/install-helpers.sh** (added function)
   - New function: `setup_amp_port_forwarding()`
   - Implements OpenChoreo's socat approach
   - Handles all 5 ports (3000, 8080, 9098, 21893, 8443)

3. **quick-start/bootstrap.sh** (updated)
   - Calls `setup_amp_port_forwarding()` function
   - No longer runs port-forward.sh as background script
   - Better integration

4. **quick-start/stop-port-forward.sh** (updated)
   - Now kills socat processes
   - Also handles kubectl port-forward (backward compatibility)
   - Clean shutdown

## Requirements

### 1. Services Must Be NodePort Type

**CRITICAL**: Update your helm charts or values to deploy services as NodePort:

```yaml
# In helm values.yaml or chart templates
console:
  service:
    type: NodePort

agentManagerService:
  service:
    type: NodePort

dataPrepper:
  service:
    type: NodePort

tracesObserver:
  service:
    type: NodePort
```

### 2. Install socat

socat must be installed in the container or on the host:

```bash
# macOS
brew install socat

# Ubuntu/Debian
apt-get install socat

# Alpine (for Docker containers)
apk add socat
```

### 3. Ensure kubectl Context is Correct

Make sure kubectl is pointing to the kind cluster:

```bash
kind export kubeconfig --name openchoreo-local
kubectl config current-context  # Should show: kind-openchoreo-local
```

## How It Works

### Network Flow

```
Mac (localhost:3000)
  ↓
  [Docker -p 3000:3000]
  ↓
Colima VM → Container port 3000
  ↓
  [socat TCP-LISTEN:3000]
  ↓
Container → Kind Worker Node (openchoreo-local-worker:30XXX)
  ↓
  [NodePort Service]
  ↓
Kubernetes Service (ClusterIP internally)
  ↓
Pod
```

### Why This Works

✅ **Direct TCP Proxy**: socat creates a direct TCP connection to the worker node
✅ **Container on kind Network**: Container can reach worker node by hostname
✅ **NodePort Exposes Service**: Service is accessible on worker node's IP:NodePort
✅ **Survives Pod Restarts**: Forwarding continues even if pods restart
✅ **No kubectl Needed**: After setup, no kubectl dependency during forwarding
✅ **Proven Pattern**: Same approach used by OpenChoreo successfully

## Testing the Fix

### 1. Deploy Services as NodePort

Update and deploy your helm charts with NodePort service types.

### 2. Run Port Forwarding

```bash
# Option 1: Via bootstrap (recommended)
./bootstrap.sh

# Option 2: Standalone script
./port-forward.sh
```

### 3. Verify Ports Are Accessible

```bash
# From your Mac
curl http://localhost:3000      # Console
curl http://localhost:8080      # Agent Manager
curl http://localhost:9098      # Traces Observer
curl http://localhost:21893     # Data Prepper
curl https://localhost:8443     # External Gateway
```

### 4. Check socat Processes

```bash
# See running socat processes
ps aux | grep socat
```

### 5. Stop Port Forwarding

```bash
./stop-port-forward.sh
```

## Comparison: Before vs After

| Aspect | Before (kubectl port-forward) | After (socat + NodePort) |
|--------|------------------------------|--------------------------|
| **Method** | kubectl port-forward | socat TCP proxy |
| **Target** | ClusterIP service | NodePort on worker node |
| **Reliability** | Low - dies on pod restart | High - survives pod restarts |
| **Dependencies** | kubectl + valid kubeconfig | Direct TCP, minimal deps |
| **Error Handling** | Silent failures | Explicit checks and retries |
| **Used By** | Our old approach | OpenChoreo (proven) |

## Next Steps

1. ✅ **Update Helm Charts**: Modify service types to NodePort
2. ✅ **Install socat**: Ensure socat is available
3. ✅ **Test Port Forwarding**: Run port-forward.sh and verify
4. ✅ **Update Documentation**: Document NodePort requirement

## References

- OpenChoreo's approach: `quick-start/install-helpers-oc.sh:302-353`
- Network analysis: `quick-start/NETWORKING_ANALYSIS.md`
- Kind networking: Container is on `kind` Docker network, can reach worker nodes directly
