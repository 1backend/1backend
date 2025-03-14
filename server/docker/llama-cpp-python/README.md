# Llama Cpp Python Docker images

This project includes code from [Llama Cpp Python](https://github.com/abetlen/llama-cpp-python),
licensed under the [MIT License](https://opensource.org/licenses/MIT).

It is built and pushed by the `llama-cpp-python-docker-build.yaml` manual dispatch GitHub action.

## Anatomy

Example image and tag: `crufter/llama-cpp-python:cuda-12.5.0-0.3.0-rc.16`, ie. `${imageName}:${backend}-${backendVersion}-${openOrchTag}.

## default

Aimed to run on any platform, uses the CPU+GPU..

## cuda

The cuda folder version is used by OpenOrch (see the `OPENORCH_GPU_PLATFORM` on the [Server envars page](https://openorch.org/docs/running-the-server/backend-environment-variables)).

It is designed to run on multiple GPUs with CUDA support.
