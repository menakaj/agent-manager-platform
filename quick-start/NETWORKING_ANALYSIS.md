# Networking Analysis: OpenChoreo vs AMP

## Problem Statement
Port forwarding from Mac → Colima VM → Container → Kind cluster is failing for all services, not just port 21893.

## Root Cause

### Current Architecture
```
Mac (localhost:3000)
  → Colima VM forwards to Container port 3000 (via -p 3000:3000)
    → Container runs: kubectl port-forward --address=0.0.0.0 svc/console 3000:3000
      → Tries to forward to K8s service
        → FAILS if service doesn't exist, pod not ready, or kubectl dies
```

### OpenChoreo's Working Architecture
```
Mac (localhost:8443)
  → Colima VM forwards to Container port 8443 (via -p 8443:8443)
    → Container runs: socat TCP-LISTEN:8443,fork TCP:openchoreo-quick-start-worker:$nodeport
      → Direct TCP connection to Kind worker node
        → NodePort service on worker
          → Pod in cluster
```

## Key Differences

| Aspect | OpenChoreo (Working) | AMP (Broken) |
|--------|---------------------|--------------|
| **Method** | `socat` TCP proxy | `kubectl port-forward` |
| **Target** | Kind worker node IP:NodePort | K8s ClusterIP service |
| **Service Type** | NodePort | ClusterIP (default) |
| **Reliability** | High - survives pod restarts | Low - dies on pod restart |
| **Dependencies** | Direct TCP, no kubectl needed | Requires kubectl + valid kubeconfig |
| **Network Path** | Container → Worker Node (same network) | Container → API Server → Service → Pod |

## Why kubectl port-forward Fails

1. **Service Must Exist**: If helm installation hasn't created the service, port-forward fails
2. **Pod Must Be Running**: Port-forward dies if pod isn't ready or restarts
3. **Kubeconfig Issues**: Kubeconfig might point to 127.0.0.1:6443 (won't work from Mac)
4. **Silent Failures**: kubectl port-forward dies without proper error handling
5. **Container Network**: Container must be on `kind` network to reach API server

## The Fix: Two Approaches

### Approach 1: Improve kubectl port-forward (Quick Fix)
- Add service existence checks
- Add pod readiness checks
- Add retry logic
- Add proper error handling
- Fix kubeconfig for container access

**Pros**: Minimal changes to helm charts
**Cons**: Still fragile, can break on pod restarts

### Approach 2: Adopt socat + NodePort (Robust Fix, Recommended)
- Change services to NodePort type
- Use socat for TCP forwarding
- Follow OpenChoreo's proven approach

**Pros**: Reliable, survives pod restarts, no kubectl dependency
**Cons**: Requires helm chart changes or values overrides

## Recommended Solution

Use **Approach 2** (socat + NodePort) because:
1. ✅ Proven to work (OpenChoreo uses it)
2. ✅ More reliable
3. ✅ Easier to debug
4. ✅ Works in container environments
5. ✅ No kubectl dependency after setup

## Implementation Steps

1. Check if `socat` is installed in container
2. Get NodePort for each service
3. Use socat to forward from container port to worker node:port
4. Update helm values to use NodePort type for services

## Example Commands

### Check if service is NodePort
```bash
kubectl get svc -n agent-management-platform agent-management-platform-console -o jsonpath='{.spec.type}'
```

### Get NodePort
```bash
kubectl get svc -n agent-management-platform agent-management-platform-console -o jsonpath='{.spec.ports[0].nodePort}'
```

### Setup socat forwarding
```bash
socat TCP-LISTEN:3000,fork TCP:openchoreo-local-worker:30080 &
```

Where:
- `3000` = Local port in container
- `openchoreo-local-worker` = Kind worker node hostname
- `30080` = NodePort assigned by K8s

---

## SOLUTION IMPLEMENTED ✅

### Changes Made

1. **port-forward.sh** - Rewritten to use socat approach
   - Uses `socat TCP-LISTEN` instead of `kubectl port-forward`
   - Gets NodePort dynamically for each service
   - Robust error handling and retry logic
   - Matches OpenChoreo's proven approach

2. **install-helpers.sh** - Added `setup_amp_port_forwarding()` function
   - Implements socat-based port forwarding
   - Waits for services to be ready
   - Gracefully handles missing services
   - Can be called from bootstrap.sh

3. **bootstrap.sh** - Updated to call new function
   - Calls `setup_amp_port_forwarding()` instead of running script
   - Uses socat approach automatically
   - Better integration with installation flow

4. **stop-port-forward.sh** - Updated to handle socat
   - Kills socat processes in addition to kubectl
   - Supports both methods for backward compatibility
   - Clean shutdown

### Requirements for Services

**CRITICAL**: Your services MUST be deployed as **NodePort** type for this to work.

Update your helm values:
```yaml
console:
  service:
    type: NodePort

agentManagerService:
  service:
    type: NodePort

# For observability services in your helm chart
dataPrepper:
  service:
    type: NodePort

tracesObserver:
  service:
    type: NodePort
```

### How to Use

1. **Deploy services as NodePort**:
   ```bash
   helm install/upgrade with values specifying NodePort
   ```

2. **Install socat** (if not already installed):
   ```bash
   # macOS
   brew install socat

   # Ubuntu/Debian
   apt-get install socat

   # Alpine (for containers)
   apk add socat
   ```

3. **Run bootstrap** (automatically sets up port forwarding):
   ```bash
   ./bootstrap.sh
   ```

4. **Or run port-forward script manually**:
   ```bash
   ./port-forward.sh
   ```

5. **Stop port forwarding**:
   ```bash
   ./stop-port-forward.sh
   ```

### Why This Works

✅ **Direct TCP Connection**: socat creates direct TCP proxy to worker node
✅ **Survives Pod Restarts**: Forwarding continues even if pods restart
✅ **No kubectl Dependency**: Works without kubectl after initial setup
✅ **Proven Approach**: Same method used by OpenChoreo
✅ **Container-Friendly**: Works perfectly in container on `kind` network

### Network Flow

```
Mac (localhost:3000)
  ↓ [Docker port mapping -p 3000:3000]
Colima VM forwards to Container port 3000
  ↓ [socat TCP-LISTEN:3000]
Container's socat forwards to openchoreo-local-worker:30XXX
  ↓ [Kind worker node NodePort]
Service in Kubernetes cluster
  ↓
Pod
```

This is the same approach OpenChoreo uses and is **much more reliable** than kubectl port-forward.
