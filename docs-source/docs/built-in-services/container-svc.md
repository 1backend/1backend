---
sidebar_position: 110
tags:
  - container-svc
  - container
  - containers
  - services
---

# Container Svc

The container service maintains containers on a node. It currently only supports Docker. It expects the Docker socket to be mounted.

For simplicity the Container Svc is only concerned with the node it resides on. In other words, the Container Svc is not distributed, it only starts and stops containers locally.

## Used By

- [Model Svc](/docs/built-in-services/file-svc) to launch containers running AI models.
- [Deploy Svc](/docs/built-in-services/file-svc) to launch containers to deploy service instances.

> This page provides a high-level overview of `Container Svc`. For detailed information, refer to the [Container Svc API documentation](/docs/openorch/run-container).

## Responsibilities

- Start and stop containers when needed - ensuring the running containers match what is expected.

## Dependencies

- [File Svc](/docs/built-in-services/file-svc) to download asset files from.
