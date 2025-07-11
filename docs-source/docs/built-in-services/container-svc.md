---
sidebar_position: 110
tags:
  - container-svc
  - container
  - containers
  - services
  - docker
  - gpu
  - cuda
---

# Container Svc

The Container Svc manages Docker containers on individual nodes, providing a unified API for container lifecycle management, image building, and resource monitoring.

> This page provides a comprehensive overview of `Container Svc`. For detailed API information, refer to the [Container Svc API documentation](/docs/1backend-api/run-container).

## Architecture

The Container Svc is **node-local** and **non-distributed**—each instance only manages containers on its own node. This design simplifies container management while supporting advanced features like GPU acceleration and persistent storage.

### Key Components
- **Docker Backend**: Primary container runtime (Docker socket required)
- **Log Accumulator**: Collects and stores container logs
- **Container Tracker**: Monitors container state and health
- **Asset Manager**: Downloads and mounts files from URLs

## Core Features

### 🐳 Container Management
- **Idempotent Operations**: Running the same container request multiple times is safe
- **Container Lifecycle**: Create, start, stop, and monitor containers
- **Health Checking**: Automatic container health monitoring
- **Log Collection**: Real-time log streaming and storage

### 🖥️ GPU Support
- **CUDA Integration**: Automatic GPU detection and container configuration
- **Dynamic Image Selection**: Chooses appropriate CUDA images based on host GPU version
- **Resource Allocation**: GPU memory and compute resource management

### 📦 Advanced Features
- **Asset Downloading**: Automatic file download and mounting from URLs
- **Persistent Storage**: "Keeps" for data persistence across container restarts
- **Environment Variables**: Dynamic environment configuration
- **Port Mapping**: Flexible host-to-container port mapping
- **Image Building**: Build containers from source repositories

## API Endpoints

### Container Operations

#### Run Container
```bash
POST /container-svc/container
```
Starts a new container or ensures an existing one matches the specification.

**Request Body:**
```json
{
  "image": "nginx:latest",
  "names": ["my-web-server"],
  "ports": [{"internal": 80, "host": 8080}],
  "envs": [{"key": "NODE_ENV", "value": "production"}],
  "assets": [{"envVarKey": "MODEL_PATH", "url": "https://example.com/model.bin"}],
  "capabilities": {"gpuEnabled": true}
}
```

#### Stop Container
```bash
PUT /container-svc/container/stop
```
Stops a running container by ID or name.

#### List Containers
```bash
POST /container-svc/containers
```
Lists all containers with optional filtering by node or container ID.

### Monitoring & Diagnostics

#### Container Status
```bash
GET /container-svc/container/is-running?hash=abc123
```
Checks if a container is currently running.

#### Container Summary
```bash
GET /container-svc/container/summary?name=my-container&lines=100
```
Gets container status and recent logs.

#### Container Logs
```bash
POST /container-svc/logs
```
Retrieves container logs with filtering options.

### System Information

#### Daemon Info
```bash
GET /container-svc/daemon/info
```
Returns Docker daemon status and system information.

#### Host Information
```bash
GET /container-svc/host
```
Gets container host address and configuration.

#### Image Availability
```bash
GET /container-svc/image/{imageName}/pullable
```
Checks if a Docker image exists and can be pulled.

### Image Building

#### Build Image
```bash
PUT /container-svc/image
```
Builds a Docker image from source code.

**Request Body:**
```json
{
  "contextPath": "/path/to/build/context",
  "dockerfilePath": "Dockerfile",
  "name": "my-custom-image"
}
```

## Integration with Other Services

### Model Svc Integration

The Model Svc uses Container Svc to launch AI model containers:

```json
{
  "image": "crufter/llama-cpp-python:cuda-12.8.0-latest",
  "names": ["the-1backend-container"],
  "ports": [{"internal": 8000, "host": 8001}],
  "assets": [
    {
      "envVarKey": "MODEL",
      "url": "https://huggingface.co/model.gguf"
    }
  ],
  "capabilities": {"gpuEnabled": true},
  "envs": [{"key": "NVIDIA_VISIBLE_DEVICES", "value": "all"}]
}
```

### Deploy Svc Integration

The Deploy Svc uses Container Svc for service deployment:

```json
{
  "image": "my-service:latest",
  "names": ["1backend-my-service"],
  "ports": [{"internal": 3000, "host": 3000}],
  "envs": [
    {"key": "NODE_ENV", "value": "production"},
    {"key": "API_KEY", "value": "secret"}
  ]
}
```

## Container Configuration

### Assets & File Mounting

Assets automatically download files and mount them in containers:

```json
{
  "assets": [
    {
      "envVarKey": "MODEL_PATH",
      "url": "https://huggingface.co/model.bin"
    },
    {
      "envVarKey": "CONFIG_FILE", 
      "url": "https://example.com/config.json"
    }
  ]
}
```

The files are downloaded by [File Svc](/docs/built-in-services/file-svc) and mounted in the container. Environment variables point to the local file paths.

### Persistent Storage (Keeps)

"Keeps" provide persistent storage across container restarts:

```json
{
  "keeps": [
    {"path": "/app/data", "readOnly": false},
    {"path": "/app/logs", "readOnly": false}
  ]
}
```

Unlike traditional volumes, keeps abstract away the host storage location—you only specify what needs to persist.

### GPU Configuration

GPU support is automatically configured based on the host system:

```json
{
  "capabilities": {"gpuEnabled": true},
  "envs": [
    {"key": "NVIDIA_VISIBLE_DEVICES", "value": "all"}
  ]
}
```

The system automatically:
- Detects available GPUs
- Selects appropriate CUDA images
- Configures GPU resource allocation

## Usage Examples

### Running a Web Application

```bash
# API call to run a containerized web app
curl -X PUT "http://localhost:11337/container-svc/container" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "nginx:alpine",
    "names": ["my-web-app"],
    "ports": [{"internal": 80, "host": 8080}],
    "envs": [{"key": "ENV", "value": "production"}],
    "keeps": [{"path": "/usr/share/nginx/html"}]
  }'
```

### Running an AI Model

```bash
# Start a GPU-enabled AI model container
curl -X PUT "http://localhost:11337/container-svc/container" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "image": "crufter/llama-cpp-python:cuda-12.8.0-latest",
    "names": ["ai-model"],
    "ports": [{"internal": 8000, "host": 8001}],
    "assets": [
      {
        "envVarKey": "MODEL",
        "url": "https://huggingface.co/model.gguf"
      }
    ],
    "capabilities": {"gpuEnabled": true},
    "envs": [{"key": "NVIDIA_VISIBLE_DEVICES", "value": "all"}]
  }'
```

### Building a Custom Image

```bash
# Build an image from source
curl -X PUT "http://localhost:11337/container-svc/image" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "contextPath": "/tmp/my-app",
    "dockerfilePath": "Dockerfile",
    "name": "my-custom-service"
  }'
```

### Monitoring Containers

```bash
# Check container status
curl "http://localhost:11337/container-svc/container/is-running?hash=abc123" \
  -H "Authorization: Bearer $TOKEN"

# Get container logs
curl "http://localhost:11337/container-svc/container/summary?name=my-container&lines=50" \
  -H "Authorization: Bearer $TOKEN"

# List all containers
curl -X POST "http://localhost:11337/container-svc/containers" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{}'
```

## Access Control

Container Svc uses permission-based access control:

- **`container-svc:container:view`** - View container information
- **`container-svc:container:run`** - Create and start containers  
- **`container-svc:container:stop`** - Stop containers
- **`container-svc:log:view`** - View container logs
- **`container-svc:image:build`** - Build custom images

By default, these permissions are granted to:
- **model-svc** and **deploy-svc** for container operations
- **Administrators** for all container management operations

## System Requirements

### Docker Configuration
- Docker daemon must be running and accessible
- Docker socket must be mounted (typically `/var/run/docker.sock`)
- User must have Docker permissions

### GPU Support (Optional)
- NVIDIA GPU with CUDA support
- NVIDIA Container Toolkit installed
- Appropriate CUDA drivers

### Storage Requirements
- Sufficient disk space for container images and persistent data
- Fast storage recommended for asset downloads and container I/O

## Advanced Features

### Container Hashing
Containers are identified by content hashes to ensure idempotent operations. The same configuration always produces the same hash, preventing duplicate containers.

### Dynamic Port Allocation
Host ports can be dynamically allocated if not specified, with the actual assigned ports returned in the response.

### Log Streaming
Container logs are collected in real-time and stored for later retrieval, with support for filtering and pagination.

### Health Monitoring
Containers are continuously monitored for health status, with automatic restart capabilities for failed containers.

## Troubleshooting

### Common Issues

**Docker Permission Denied**
```bash
# Add user to docker group
sudo usermod -aG docker $USER
# Restart session or run
newgrp docker
```

**GPU Not Detected**
```bash
# Check NVIDIA drivers
nvidia-smi
# Verify Docker GPU support
docker run --rm --gpus all nvidia/cuda:11.0-base nvidia-smi
```

**Container Won't Start**
- Check Docker daemon status: `systemctl status docker`
- Verify image availability: `docker images`
- Review container logs via Container Svc API

**Port Conflicts**
- Use dynamic port allocation or check for port conflicts
- Review existing container port mappings

## Related Services

- **[Model Svc](/docs/built-in-services/model-svc)** - AI model deployment and management
- **[Deploy Svc](/docs/built-in-services/deploy-svc)** - Service deployment orchestration
- **[File Svc](/docs/built-in-services/file-svc)** - Asset downloading and file management
- **[Registry Svc](/docs/built-in-services/registry-svc)** - Node and instance registry
