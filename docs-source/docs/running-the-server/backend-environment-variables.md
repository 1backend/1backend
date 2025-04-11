---
sidebar_position: 3
tags:
  - configuration
  - envars
  - backend
  - daemon
  - setup
---

# Backend Environment Variables

## `OB_SERVER_URL`

The OB_SERVER_URL is the internally addressable (non-public-facing) URL of an 1Backend server. It should point to the local 1Backend instance on each physical node. Ideally, every node should have its own 1Backend instance.

This envar should be set only for microservices built on 1Backend. The 1Backend server itself should use `OB_SELF_URL`.

## `OB_SELF_URL`

Microservices use this to register themselves in the 1Backend registry. The 1Backend server uses this to address itself.

## `OB_TEST`

Microservices and the 1Backend server uses this envar to detect if they are running as part of a test.

When set to true, subsystems act accordingly: for example the datastore will prefix tables with random numbers to provide a unique and clean environment for each test.

## `OB_NODE_ID`

For information about this, please refer to the [Registry Svc Node section](/docs/built-in-services/registry-svc#node)

## `OB_GPU_PLATFORM`

This envar is used to enabel GPU acceleration.
Supported platforms:

- `cuda`

Do not set this if your card doesn't support the given architecture or things will break.

## `OB_VOLUME_NAME`

**This flag is typically unnecessary since 1Backend automatically detects the volume that is bound to `/root/.1backend`. Use it only as a corrective action.**

This envar is needed when 1Backend runs as a container next to containers it starts:

```sh
Host
 |
 |-> 1Backend Container
 |-> Container Launched By 1Backend
```

For the containers like `llama-cpp` to be able to read the models downloaded by 1Backend we they must both mount the same docker volume.

An example of this can be seen in the root `docker-compose.yaml` file: `OB_VOLUME_NAME=singulatron-data`.

So cycle goes like this:

- 1Backend container writes to `/root/.1backend`, which is mounted to the volume `singulatron-data`
- Assets (which are basically downloaded files) will be passed to containers created by 1Backend by mounting files in `singulatron-data`.

## `OB_LLM_HOST`

**This flag is typically unnecessary since 1Backend retrieves the IP of the Docker bridge. Use it only as a corrective action.**

When 1Backend is running in a container, it needs to know how to address its siblings (other containers it started):

```sh
Host
 |
 |-> 1Backend Container
 |-> Container Launched By 1Backend
```

The `1Backend Container` uses the envar `OB_LLM_HOST` to address `Container Launched By 1Backend`.

Typically this value should be `172.17.0.1` if you are using the default docker network.

If you are using an other network than default, use `docker network inspect` to find out the IP of your docker bridge for that network.
Usually it's going to be `172.18.0.1`.

This envar is not needed if 1Backend runs directly on the host:

```sh
Host With 1Backend
 |
 |-> Container Launched By 1Backend
```

## `OB_DB`

You can use this envar to make 1Backend actually use a database instead of local file storage to store data.

### PostgreSQL

```sh
OB_DB=postgres
OB_DB_DRIVER="postgres" # or "mysql"
OB_DB_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.
