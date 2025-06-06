# 1Backend backend Docker images

This folder contains various GPU-optimized (CUDA, etc.) 1Backend images, each with specific versions and configurations.

Although AI engines run in separate containers managed by 1Backend, certain features—such as accessing GPU statistics—still require direct integration within these backend images. Tools like nvidia-smi, for example, are needed to support these functionalities.

## Build and push

These images are pushed by the `backend-container-build.yaml` manual dispatch GitHub workflows.