---
sidebar_position: 90
tags:
  - model-svc
  - models
  - ai
  - services
---

# Model Svc

The model service can start, stop AI models across multiple runtimes (eg. Docker) and maintains a database of available models on the platform.

> This page provides a high-level overview of `Model Svc`. For detailed information, refer to the [Model Svc API documentation](/docs/openorch/start-default-model).

## Responsibilities

- Start and stop models
- Maintain a database of models and other related information such as the default model

## Dependencies

- [Container Svc](/docs/built-in-services/container-svc) to start containerized AI models (eg. Llama, Stabel Diffusion etc.)

## Current Limitations

- Stop model endpoint is missing
