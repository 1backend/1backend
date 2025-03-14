---
sidebar_position: 2
tags:
  - run
  - deploy
---

# Running the OpenOrch server with Docker Compose and prebuilt images

This deployment method is one step above local development in terms of sophistication. It’s suitable for a development server or simple production environments.

This snippet will give you a quick idea about how to deploy the frontend and backend containers so they play nicely together:

```yaml
version: "3.8"

volumes:
  openorch-data:
    name: openorch-data
    driver: local

services:
  openorch-frontend:
    image: crufter/openorch-frontend:latest
    ports:
      - "3901:80"
    environment:
      # `BACKEND_ADDRESS` must be reachable from the browser.
      # This is the API the browser will communicate with, not an internal address.
      - BACKEND_ADDRESS=http://127.0.0.1:58231

  openorch-backend:
    image: crufter/openorch-backend:default-1-latest
    # Use a version that matches your GPU architecture for GPU acceleration, e.g.:
    # crufter/openorch-backend:cuda-12.2.0-latest
    # For available versions, see:
    # - https://hub.docker.com/r/crufter/openorch-backend/tags
    # - The build file `openorch-docker-build.yaml`
    ports:
      - "58231:58231"
    volumes:
      # We mount the hostname to have a sensible fallback node URL
      - /etc/hostname:/etc/host_hostname:ro
      # We mount the docker socket so the backend can start   containers
      - /var/run/docker.sock:/var/run/docker.sock
      # We mount a volume so data will be persisted
      - openorch-data:/root/.openorch
    # Grants OpenOrch access to GPU metrics.
    # Containers launched by OpenOrch can still use GPU acceleration even if OpenOrch lacks direct GPU access.
    # deploy:
    #   resources:
    #     reservations:
    #       devices:
    #         - driver: nvidia
    #           count: all
    #           capabilities: [gpu]
    environment:
      # Volume mounted by AI containers launched by OpenOrch to access models downloaded by the OpenOrch File Svc.
      - OPENORCH_VOLUME_NAME=openorch-data
      #
      # Enables GPU acceleration for NVIDIA GPUs.
      # This flag controls GPU access for AI containers launched by OpenOrch.
      #
      # - OPENORCH_GPU_PLATFORM=cuda
```

Put the above into a file called `docker-compose.yaml` in a folder on your computer and run it with the following command:

```sh
docker compose up
```

## Once it's running

After the containers successfully start, you can go to `127.0.0.1:3901` and log in with the [Default Credentials](/docs/running-the-server/using#default-credentials).

## Configuring

See the [Backend Environment Variables](./backend-environment-variables/) and [Frontend Environment Variables](./backend-environment-variables/).
