# Container Networking Guide

## Understanding Kind Cluster Access from Containers

When running the dev container, there's an important networking consideration: **Kind clusters created inside a container need special configuration to be accessible**.

## The Problem

Kind creates Kubernetes clusters using Docker containers. By default, Kind's kubeconfig uses `https://127.0.0.1:6443` (localhost) as the API server endpoint. This works fine on the host machine, but **fails inside a container** because:

1. The Kind control plane runs in a Docker container on the host
2. The dev container is also a Docker container
3. `localhost` inside the dev container points to itself, not the host
4. Therefore, kubectl cannot reach the Kind API server

## The Solution

We use **Kind's `--internal` flag** which configures kubectl to use the Docker network IP instead of localhost:

```bash
# This generates a kubeconfig that works from inside containers
kind export kubeconfig --name openchoreo-local --internal
```

This command:
- Gets the internal Docker network IP of the Kind control plane
- Updates the kubeconfig to use `https://<docker-ip>:6443`
- Works from both host and container

## Automatic Configuration

Our scripts automatically handle this:

### 1. During Cluster Creation

When `bootstrap.sh` creates a Kind cluster inside the container:

```bash
# In install-helpers.sh
setup_kind_cluster() {
    # ... create cluster ...
    
    # Automatically fix kubeconfig if in container
    fix_kubeconfig_for_container "$cluster_name"
}
```

### 2. On Container Start

When you start the dev container with an existing cluster:

```bash
# In entrypoint.sh
if kind get clusters | grep -q "openchoreo-local"; then
    kind export kubeconfig --name openchoreo-local --internal
fi
```

## Manual Fix

If you encounter the error:

```
Error: Kubernetes cluster unreachable: Get "http://localhost:8080/version": 
dial tcp [::1]:8080: connect: connection refused
```

**Inside the container, run:**

```bash
# Fix kubeconfig for container access
kind export kubeconfig --name openchoreo-local --internal

# Verify it works
kubectl get nodes
```

## Network Architecture

```
┌─────────────────────────────────────────────┐
│  Host Machine                                │
│                                              │
│  Docker Network (bridge)                    │
│  ┌────────────────────────────────────┐    │
│  │                                     │    │
│  │  Kind Control Plane Container      │    │
│  │  IP: 172.18.0.2                    │    │
│  │  Port: 6443                        │    │
│  │                                     │    │
│  └────────────────────────────────────┘    │
│         ▲                                    │
│         │ Docker network                    │
│         │                                    │
│  ┌──────┴──────────────────────────────┐   │
│  │  Dev Container                       │   │
│  │  IP: 172.18.0.3                     │   │
│  │                                      │   │
│  │  kubectl → https://172.18.0.2:6443 │   │
│  │  ✅ Works!                          │   │
│  │                                      │   │
│  │  kubectl → https://localhost:6443   │   │
│  │  ❌ Fails! (localhost = self)      │   │
│  └──────────────────────────────────────┘   │
└─────────────────────────────────────────────┘
```

## Verification

After fixing the kubeconfig, verify access:

```bash
# Check cluster info
kubectl cluster-info

# Should show something like:
# Kubernetes control plane is running at https://172.18.0.2:6443

# Check nodes
kubectl get nodes

# Should show:
# NAME                            STATUS   ROLES           AGE   VERSION
# openchoreo-local-control-plane  Ready    control-plane   5m    v1.32.0
# openchoreo-local-worker         Ready    <none>          5m    v1.32.0
```

## Common Issues

### Issue 1: "connection refused" on localhost:8080

**Cause:** kubectl is using default config (localhost:8080)

**Solution:**
```bash
# Export kubeconfig with internal endpoint
kind export kubeconfig --name openchoreo-local --internal

# Or set KUBECONFIG explicitly
export KUBECONFIG=/root/.kube/config
```

### Issue 2: "x509: certificate is valid for..., not 172.18.0.X"

**Cause:** Certificate doesn't include the Docker network IP

**Solution:** This shouldn't happen with Kind, but if it does:
```bash
# Recreate the cluster
kind delete cluster --name openchoreo-local
./bootstrap.sh
```

### Issue 3: "Unable to connect to the server"

**Cause:** Kind cluster not running or network issue

**Solution:**
```bash
# Check if cluster is running
kind get clusters
docker ps | grep openchoreo-local

# Restart cluster containers if needed
docker start openchoreo-local-control-plane openchoreo-local-worker

# Re-export kubeconfig
kind export kubeconfig --name openchoreo-local --internal
```

## Best Practices

### 1. Always Use --internal in Containers

When working inside a container:
```bash
# ✅ Good
kind export kubeconfig --name openchoreo-local --internal

# ❌ Bad (won't work in container)
kind export kubeconfig --name openchoreo-local
```

### 2. Check Your Environment

Determine if you're in a container:
```bash
# Check for container indicators
if [ -f /.dockerenv ]; then
    echo "Running in container"
    kind export kubeconfig --internal
fi
```

### 3. Use Full Paths

Always use absolute paths for kubeconfig:
```bash
export KUBECONFIG=/root/.kube/config
```

## Why This Matters

This networking consideration is **crucial** for:
- ✅ Running bootstrap inside dev container
- ✅ Using kubectl from container
- ✅ Helm installations
- ✅ Any Kubernetes operations

Without proper configuration:
- ❌ kubectl commands fail
- ❌ Helm installations fail
- ❌ Cannot deploy applications
- ❌ Cannot access cluster resources

## References

- [Kind Documentation - Working with Clusters](https://kind.sigs.k8s.io/docs/user/working-with-clusters/)
- [Docker Networking](https://docs.docker.com/network/)
- [Kubernetes API Server](https://kubernetes.io/docs/reference/command-line-tools-reference/kube-apiserver/)

---

**TL;DR:** When using Kind inside a container, always run:
```bash
kind export kubeconfig --name openchoreo-local --internal
```

This is automatically handled by our scripts! 🎉

