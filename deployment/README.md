# Agent Management Platform - Docker Development Environment

This Docker image provides a quick development environment for the Agent Management Platform, built on top of OpenChoreo Quick Start v0.3.2.

## What's Included

- **Base Image**: `ghcr.io/openchoreo/quick-start:v0.3.2` (includes Kubernetes cluster and OpenChoreo components)
- **Helm Charts**: Agent Management Platform charts
- **Installation Script**: `install-amp.sh` for automated setup

## Prerequisites

- Docker installed and running
- Sufficient resources allocated to Docker (recommended: 4GB RAM, 2 CPUs minimum)

## Building the Image

```bash
cd deployments
docker build -f Dockerfile.dev -t agent-management-platform-dev:latest .
```

## Usage

### 1. Start the Container

The base image sets up a Kubernetes cluster automatically. Simply run:

```bash
docker run -d --name amp-dev --privileged agent-management-platform-dev:latest
```

**Note**: `--privileged` flag is required for the base image to run Kubernetes cluster.

### 2. Access the Container

```bash
docker exec -it amp-dev bash
```

### 3. Install Agent Management Platform

Once inside the container:

```bash
# Navigate to the installation directory
cd /opt/agent-management-platform

# Install with default settings
./install-amp.sh

# Or install with all components (observability + build-ci)
./install-amp.sh --with-observability --with-build-ci

# Or customize the installation
./install-amp.sh --namespace my-namespace --values values-dev.yaml
```

### 4. Verify Installation

```bash
# Check helm releases
helm list -A

# Check pods
kubectl get pods -n agent-platform

# Check services
kubectl get svc -n agent-platform
```

### 5. Access the Services

From inside the container:

```bash
# Port forward the console
kubectl port-forward -n agent-platform svc/agent-platform-console 3000:80 &

# Port forward the API
kubectl port-forward -n agent-platform svc/agent-platform-agent-manager-service 8080:8080 &
```

To access from your host machine, you'll need to expose ports when starting the container:

```bash
docker run -d --name amp-dev --privileged \
  -p 3000:3000 \
  -p 8080:8080 \
  agent-management-platform-dev:latest
```

Then run the port-forward commands inside the container.

## Installation Script Options

The `install-amp.sh` script supports the following options:

```
Options:
  -n, --namespace NAMESPACE              Target namespace (default: agent-platform)
  -r, --release RELEASE_NAME             Helm release name (default: agent-platform)
  -f, --values VALUES_FILE               Values file (default: values.yaml)
  --with-observability                   Install observability-dataprepper
  --with-build-ci                        Install build-ci
  --observability-namespace NAMESPACE    Observability namespace (default: openchoreo-observability-plane)
  --build-ci-namespace NAMESPACE         Build CI namespace (default: openchoreo-build-plane)
  -h, --help                             Display help message
```

## Examples

### Example 1: Basic Installation

```bash
docker run -d --name amp-dev --privileged agent-management-platform-dev:latest
docker exec -it amp-dev bash
cd /opt/agent-management-platform
./install-amp.sh
```

### Example 2: Full Installation with All Components

```bash
docker run -d --name amp-dev --privileged agent-management-platform-dev:latest
docker exec -it amp-dev bash
cd /opt/agent-management-platform
./install-amp.sh --with-observability --with-build-ci
```

### Example 3: Development Mode Installation

```bash
docker run -d --name amp-dev --privileged agent-management-platform-dev:latest
docker exec -it amp-dev bash
cd /opt/agent-management-platform
./install-amp.sh --values helm-charts/agent-management-platform/values-dev.yaml
```

### Example 4: With Port Forwarding for Host Access

```bash
# Start container with exposed ports
docker run -d --name amp-dev --privileged \
  -p 3000:3000 \
  -p 8080:8080 \
  agent-management-platform-dev:latest

# Install the platform
docker exec -it amp-dev bash -c "cd /opt/agent-management-platform && ./install-amp.sh"

# Set up port forwarding (in another terminal)
docker exec -d amp-dev kubectl port-forward -n agent-platform svc/agent-platform-console 3000:80 --address 0.0.0.0
docker exec -d amp-dev kubectl port-forward -n agent-platform svc/agent-platform-agent-manager-service 8080:8080 --address 0.0.0.0
```

Now access from your host:
- Console: http://localhost:3000
- API: http://localhost:8080

## Uninstalling

To remove the Agent Management Platform:

```bash
# Uninstall helm releases
helm uninstall agent-platform -n agent-platform
helm uninstall observability-dataprepper -n openchoreo-observability-plane
helm uninstall build-ci -n openchoreo-build-plane

# Delete namespaces
kubectl delete namespace agent-platform
kubectl delete namespace openchoreo-observability-plane
kubectl delete namespace openchoreo-build-plane
```

## Stopping and Removing the Container

```bash
# Stop the container
docker stop amp-dev

# Remove the container
docker rm amp-dev

# Remove the image
docker rmi agent-management-platform-dev:latest
```

## Troubleshooting

### Container Won't Start

Make sure Docker has sufficient resources and the `--privileged` flag is used:

```bash
docker run -d --name amp-dev --privileged agent-management-platform-dev:latest
```

### Kubernetes Not Ready

Wait a few minutes for the Kubernetes cluster to initialize. Check status:

```bash
docker exec -it amp-dev kubectl cluster-info
docker exec -it amp-dev kubectl get nodes
```

### Installation Fails

Check logs and pod status:

```bash
# Check pod status
kubectl get pods -n agent-platform

# Check pod logs
kubectl logs -n agent-platform <pod-name>

# Check helm release status
helm status agent-platform -n agent-platform
```

### Port Forwarding Not Working

Ensure the services are running and use `--address 0.0.0.0` for external access:

```bash
kubectl port-forward -n agent-platform svc/agent-platform-console 3000:80 --address 0.0.0.0
```

## Directory Structure

Inside the container:

```
/opt/agent-management-platform/
├── helm-charts/
│   ├── agent-management-platform/    # Main platform chart
│   ├── observability-dataprepper/    # Observability chart
│   └── build-ci/                     # Build CI chart
└── install-amp.sh                    # Installation script
```

## Base Image Information

The base image (`ghcr.io/openchoreo/quick-start:v0.3.2`) provides:
- Kubernetes cluster (using k3s)
- OpenChoreo Control Plane
- OpenChoreo Data Plane
- OpenChoreo Observability Plane
- Pre-configured networking and storage

For more information about the base image, visit: https://github.com/openchoreo/openchoreo

## Development Tips

1. **Use volumes for development**: Mount your local helm charts for live updates:
   ```bash
   docker run -d --name amp-dev --privileged \
     -v $(pwd)/helm-charts:/opt/agent-management-platform/helm-charts \
     agent-management-platform-dev:latest
   ```

2. **Keep container running**: The container will continue running in the background. Use `docker exec` to access it anytime.

3. **Check logs**: View container logs with:
   ```bash
   docker logs amp-dev
   ```

4. **Quick reinstall**: To reinstall quickly, uninstall and run the script again instead of rebuilding the container.

## License

See the main project LICENSE file for details.
