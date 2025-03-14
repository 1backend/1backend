---
sidebar_position: 50
tags:
  - deploy-svc
  - deploy
  - containers
  - services
---

# Deploy Svc

The deploy service is responsible of launching containers on whatever infrastructure the 1Backend is running on (eg. [Docker Svc](/docs/built-in-services/container-svc)) and registering them into the [Registry Svc](/docs/built-in-services/container-svc).

> This page provides a high-level overview of `Deploy Svc`. For detailed information, refer to the [Deploy Svc API documentation](/docs/1backend/save-deployment).

## Warning

Deployment capabilities are unfinished. This section is only for contributors. Deploy your services manually for now.

## Types of services

On 1Backend, services generally fall into the following categories:

- Services deployed by the `Deploy Svc`: These can include services built with the 1Backend SDK or standard containers (e.g., NGINX) that are not 1Backend-specific.
- Services built with the 1Backend SDK but deployed through other methods; for example, using Docker Compose or Kubernetes. These services self-register into the `Registry Svc`.

## Entities

### Deployment

```yaml
id: depl_dbOdi5eLQK
definitionId: def_deBXZMpxrQ
name: user-service-v2
description: Handles user service requests
replicas: 3
strategy:
  type: RollingUpdate
  maxUnavailable: 1
  maxSurge: 2
resources:
  cpu: "500m"
  memory: "256Mi"
  vram: "24GB"
autoScaling:
  minReplicas: 2
  maxReplicas: 10
  cpuThreshold: 80
targetRegions:
  - cluster: us-west1
    zone: us-west1-b
  - cluster: local-docker
status: OK
details: Deployment is in progress
envars:
  ENVIRONMENT: production
  LOG_LEVEL: debug
  FEATURE_FLAG_X: "true"
```

## Dependencies

- [Container Svc](/docs/built-in-services/container-svc) to start containers of services
- [Registry Svc](/docs/built-in-services/registry-svc) to start containers of services
