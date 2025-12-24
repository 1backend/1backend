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
- **[File Svc](/docs/built-in-services/file-svc)** - Asset downloading and file management
- **[Registry Svc](/docs/built-in-services/registry-svc)** - Node and instance registry
