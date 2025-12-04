# Dev Container Guide

Run the complete Agent Management Platform in a single Docker container using Docker-in-Docker.

## 🎯 What is the Dev Container?

The dev container is a lightweight Docker environment that includes:
- **Docker CLI** - Uses your host's Docker daemon (no Docker-in-Docker!)
- **Kind cluster** - Kubernetes cluster created automatically
- **OpenChoreo** - Complete platform installation
- **Agent Management Platform** - Your application with observability
- **All tools** - kubectl, helm, kind pre-installed

**Key Advantage:** Uses the host's Docker daemon via socket mounting - more secure, faster, and smaller than Docker-in-Docker!

## ⚡ Quick Start

### Step 1: Build the Image

```bash
cd quick-start
./docker-run.sh --build
```

**Time:** ~5-10 minutes (one-time build)

### Step 2: Run the Container

```bash
./docker-run.sh
```

**Time:** ~15-20 minutes for complete installation

That's it! The container will:
1. Start Docker daemon inside
2. Create Kind cluster
3. Install OpenChoreo
4. Install Agent Management Platform
5. Set up port forwarding

## 🌐 Access Your Platform

Once installation completes, access from your host machine:

| Service | URL | Description |
|---------|-----|-------------|
| **Console** | http://localhost:3000 | Web UI |
| **API** | http://localhost:8080 | REST API |
| **Traces Observer** | http://localhost:9098 | Trace analysis |
| **Data Prepper** | http://localhost:21893 | Trace ingestion |

## 📋 Prerequisites

### Required
- **Docker** - Version 20.10+ running on host
- **RAM** - 4GB+ allocated to Docker
- **Disk** - 10GB+ free space
- **Ports** - 3000, 8080, 9098, 21893 available

**Note:** The container uses your host's Docker daemon, so Docker must be running before starting the container.

### Check Docker Resources

**macOS/Windows (Docker Desktop):**
```bash
# Check current allocation
docker info | grep Memory

# Increase if needed:
# Docker Desktop → Settings → Resources → Memory (set to 4GB+)
```

**Linux:**
```bash
# Docker uses host resources directly
free -h
```

## 🚀 Usage Examples

### Full Installation (Default)

```bash
./docker-run.sh
```

Installs everything including optional OpenChoreo components.

### Minimal Installation (Faster)

```bash
./docker-run.sh --minimal
```

Installs only core components. **Time:** ~10-12 minutes

### Verbose Output

```bash
./docker-run.sh --verbose
```

Shows detailed installation progress. Useful for debugging.

### Background Mode

```bash
./docker-run.sh --detach

# Then follow logs
docker logs -f amp-devcontainer
```

Runs container in background (detached mode).

### Interactive Shell

```bash
./docker-run.sh --shell
```

Starts an interactive bash shell instead of running bootstrap.

### Build and Run

```bash
./docker-run.sh --build
```

Builds the image and runs the container in one command.

## 🎛️ Advanced Usage

### Custom Container Name

```bash
./docker-run.sh --name my-amp-container
```

### Combine Options

```bash
# Minimal installation with verbose output in background
./docker-run.sh --minimal --verbose --detach
```

### Manual Docker Run

If you prefer to use `docker run` directly:

```bash
# Build image first
docker build -f Dockerfile.devcontainer -t amp-devcontainer .

# Run container (note: mounts Docker socket, no --privileged needed!)
docker run \
  --name amp-devcontainer \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -p 3000:3000 \
  -p 8080:8080 \
  -p 9098:9098 \
  -p 21893:21893 \
  -e BOOTSTRAP_MODE=full \
  -e BOOTSTRAP_VERBOSE=false \
  amp-devcontainer
```

### Environment Variables

| Variable | Values | Default | Description |
|----------|--------|---------|-------------|
| `BOOTSTRAP_MODE` | `full`, `minimal` | `full` | Installation mode |
| `BOOTSTRAP_VERBOSE` | `true`, `false` | `false` | Verbose output |

Example:
```bash
docker run \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e BOOTSTRAP_MODE=minimal \
  -e BOOTSTRAP_VERBOSE=true \
  -p 3000:3000 -p 8080:8080 -p 9098:9098 -p 21893:21893 \
  amp-devcontainer
```

## 🔍 Container Management

### View Logs

```bash
# Follow logs in real-time
docker logs -f amp-devcontainer

# View last 100 lines
docker logs --tail 100 amp-devcontainer
```

### Enter Container Shell

```bash
docker exec -it amp-devcontainer bash
```

Once inside:
```bash
# Check cluster status
kubectl get nodes
kubectl get pods --all-namespaces

# Check platform status
kubectl get pods -n agent-management-platform

# Run commands
./bootstrap.sh --help
```

### Check Container Status

```bash
# Is container running?
docker ps | grep amp-devcontainer

# Container details
docker inspect amp-devcontainer

# Resource usage
docker stats amp-devcontainer
```

### Stop Container

```bash
docker stop amp-devcontainer
```

**Note:** This stops the container but preserves the state.

### Start Stopped Container

```bash
docker start amp-devcontainer

# Follow logs
docker logs -f amp-devcontainer
```

### Remove Container

```bash
# Stop and remove
docker rm -f amp-devcontainer

# Remove image too
docker rmi amp-devcontainer
```

## 🛠️ Troubleshooting

### Container Fails to Start

**Check Docker is running:**
```bash
docker info
```

**Check available resources:**
```bash
docker info | grep -E 'Memory|CPUs'
```

**View container logs:**
```bash
docker logs amp-devcontainer
```

### Cannot Connect to Docker Daemon

**Symptom:** "Cannot connect to the Docker daemon"

**Solution:**
```bash
# Ensure Docker socket is mounted
docker run -v /var/run/docker.sock:/var/run/docker.sock ...

# Verify Docker is running on host
docker info

# Check socket permissions
ls -la /var/run/docker.sock

# On some systems, you may need to add your user to docker group
sudo usermod -aG docker $USER
```

### Ports Already in Use

**Check what's using the ports:**
```bash
# macOS/Linux
lsof -i :3000
lsof -i :8080

# Windows
netstat -ano | findstr :3000
```

**Solution:**
- Stop the conflicting service, or
- Use different ports:
  ```bash
  docker run --privileged \
    -p 3001:3000 \
    -p 8081:8080 \
    ...
  ```

### Installation Fails

**Run with verbose mode:**
```bash
./docker-run.sh --verbose
```

**Check inside container:**
```bash
docker exec -it amp-devcontainer bash

# Check Docker
docker ps

# Check Kind cluster
kind get clusters
kubectl get nodes

# Check pods
kubectl get pods --all-namespaces
```

### Out of Memory

**Symptom:** Pods stuck in `Pending` or `CrashLoopBackOff`

**Solution:**
1. Increase Docker memory allocation to 6GB+
2. Use minimal mode: `./docker-run.sh --minimal`
3. Close other applications

### Slow Performance

**Causes:**
- Insufficient resources
- Multiple layers of virtualization
- Disk I/O bottleneck

**Solutions:**
1. Increase Docker resources (CPU + RAM)
2. Use SSD for Docker storage
3. Close unnecessary applications
4. Consider using `bootstrap.sh` directly on host instead

## 🧹 Cleanup

### Remove Everything

```bash
# Stop and remove container
docker rm -f amp-devcontainer

# Remove image
docker rmi amp-devcontainer

# Clean up Docker system
docker system prune -a
```

### Rebuild from Scratch

```bash
# Remove old container and image
docker rm -f amp-devcontainer
docker rmi amp-devcontainer

# Rebuild and run
./docker-run.sh --build
```

## 📊 What's Inside the Container?

### Installed Tools

- **Docker** - v24.0.7
- **kubectl** - v1.29.0
- **Helm** - v3.13.3
- **kind** - v0.20.0
- **bash-completion** - For command completion
- **git, vim, curl** - Standard utilities

### Directory Structure

```
/workspace/
├── bootstrap.sh           # Complete installation script
├── install.sh             # Platform-only installer
├── install-helpers.sh     # Helper functions
├── uninstall.sh          # Cleanup script
├── port-forward.sh       # Port forwarding
├── kind-config.yaml      # Kind cluster configuration
└── *.md                  # Documentation
```

### Kubernetes Context

The container automatically configures:
- Kind cluster named `openchoreo-local`
- Kubeconfig at `/root/.kube/config`
- Context set to `kind-openchoreo-local`

## 💡 Tips and Best Practices

### 1. Use Minimal Mode for Development

```bash
./docker-run.sh --minimal
```

Faster startup, includes everything you need for development.

### 2. Run in Background

```bash
./docker-run.sh --detach
docker logs -f amp-devcontainer
```

Free up your terminal while installation runs.

### 3. Keep Container Running

Don't remove the container between sessions:
```bash
# Stop when not in use
docker stop amp-devcontainer

# Start when needed
docker start amp-devcontainer
```

### 4. Monitor Resources

```bash
# Watch resource usage
docker stats amp-devcontainer
```

### 5. Backup Kubeconfig

```bash
# Copy kubeconfig from container
docker cp amp-devcontainer:/root/.kube/config ./kubeconfig-backup
```

## 🆚 Dev Container vs Direct Installation

| Aspect | Dev Container | Direct (bootstrap.sh) |
|--------|---------------|----------------------|
| **Setup** | Docker only | Docker + kubectl + helm + kind |
| **Isolation** | Complete | Uses host Docker |
| **Performance** | Slower (nested) | Faster (native) |
| **Portability** | High | Medium |
| **Resource Usage** | Higher | Lower |
| **Debugging** | Harder | Easier |
| **Best For** | Testing, CI/CD | Development |

## 🚦 When to Use Dev Container

**✅ Use Dev Container When:**
- You want complete isolation
- Testing in CI/CD pipelines
- Quick demos or POCs
- Don't want to install tools locally
- Need reproducible environments

**❌ Use Direct Installation When:**
- Doing active development
- Need better performance
- Debugging issues
- Long-running development environment
- Limited Docker resources

## 📚 Next Steps

1. **Access the Console**
   ```bash
   open http://localhost:3000
   ```

2. **Deploy a Sample Agent**
   ```bash
   docker exec -it amp-devcontainer bash
   cd /workspace/../runtime/sample-agents/python-agent
   # Follow the README
   ```

3. **Explore the Platform**
   - Create agents
   - View traces
   - Test observability features

4. **Read More Documentation**
   - [Bootstrap Guide](./BOOTSTRAP_GUIDE.md)
   - [Quick Start](./QUICK_START.md)
   - [Troubleshooting](./TROUBLESHOOTING.md)

## 🆘 Getting Help

If you encounter issues:

1. **Check logs:**
   ```bash
   docker logs amp-devcontainer
   ```

2. **Enter container:**
   ```bash
   docker exec -it amp-devcontainer bash
   kubectl get pods --all-namespaces
   ```

3. **Run with verbose:**
   ```bash
   ./docker-run.sh --verbose
   ```

4. **Check documentation:**
   - [TROUBLESHOOTING.md](./TROUBLESHOOTING.md)
   - [BOOTSTRAP_GUIDE.md](./BOOTSTRAP_GUIDE.md)

5. **Open an issue:**
   - GitHub: https://github.com/wso2/agent-management-platform/issues
   - Include logs and error messages

---

**Ready to start?**

```bash
cd quick-start
./docker-run.sh --build
```

Happy containerizing! 🐳

