---
sidebar_position: 2
tags:
  - run
  - deploy
---

# Running the 1Backend server with Docker Compose and prebuilt images

This deployment method is one step above local development in terms of sophistication. It’s suitable for a development server or simple production environments.

This snippet will give you a quick idea about how to deploy the frontend and backend containers so they play nicely together:

```yaml
version: "3.8"

volumes:
  1backend-data:
    name: 1backend-data
    driver: local

services:
  1backend-ui:
    image: crufter/1backend-ui:latest
    ports:
      - "3901:80"
    environment:
      # `BACKEND_ADDRESS` must be reachable from the browser.
      # This is the API the browser will communicate with, not an internal address.
      - BACKEND_ADDRESS=http://127.0.0.1:58231

  1backend:
    image: crufter/1backend:default-1-latest
    # Use a version that matches your GPU architecture for GPU acceleration, e.g.:
    # crufter/1backend:cuda-12.2.0-latest
    # For available versions, see:
    # - https://hub.docker.com/r/crufter/1backend/tags
    # - The build file `1backend-docker-build.yaml`
    ports:
      - "58231:58231"
    volumes:
      # We mount the hostname to have a sensible fallback node URL
      - /etc/hostname:/etc/host_hostname:ro
      # We mount the docker socket so the backend can start   containers
      - /var/run/docker.sock:/var/run/docker.sock
      # We mount a volume so data will be persisted
      - 1backend-data:/root/.1backend
    # Grants 1Backend access to GPU metrics.
    # Containers launched by 1Backend can still use GPU acceleration even if 1Backend lacks direct GPU access.
    # deploy:
    #   resources:
    #     reservations:
    #       devices:
    #         - driver: nvidia
    #           count: all
    #           capabilities: [gpu]
    environment:
      # Volume mounted by AI containers launched by 1Backend to access models downloaded by the 1Backend File Svc.
      - OB_VOLUME_NAME=1backend-data
      #
      # Enables GPU acceleration for NVIDIA GPUs.
      # This flag controls GPU access for AI containers launched by 1Backend.
      #
      # - OB_GPU_PLATFORM=cuda
```

Put the above into a file called `docker-compose.yaml` in a folder on your computer and run it with the following command:

```sh
docker compose up
```

## Once it's running

After the containers successfully start, you can go to `127.0.0.1:3901` and log in with the [Default Credentials](/docs/running-the-server/using#default-credentials).

## Configuring

See the [Backend Environment Variables](./backend-environment-variables/) and [Frontend Environment Variables](./backend-environment-variables/).
