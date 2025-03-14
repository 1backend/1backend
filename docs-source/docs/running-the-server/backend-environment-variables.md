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

## `OPENORCH_URL`

The OPENORCH_URL is the internally addressable (non-public-facing) URL of an OpenOrch server. It should point to the local OpenOrch instance on each physical node. Ideally, every node should have its own OpenOrch instance.

This local address serves two purposes:

- It acts as the "self address" for the OpenOrch server when calling built-in services (though it can default to 127.0.0.1).
- It is also used by custom services when registering themselves, among other tasks.

## `OPENORCH_NODE_ID`

For information about this, please refer to the [Registry Svc Node section](/docs/built-in-services/registry-svc#node)

## `OPENORCH_GPU_PLATFORM`

This envar is used to enabel GPU acceleration.
Supported platforms:

- `cuda`

Do not set this if your card doesn't support the given architecture or things will break.

## `OPENORCH_VOLUME_NAME`

**This flag is typically unnecessary since OpenOrch automatically detects the volume that is bound to `/root/.openorch`. Use it only as a corrective action.**

This envar is needed when OpenOrch runs as a container next to containers it starts:

```sh
Host
 |
 |-> OpenOrch Container
 |-> Container Launched By OpenOrch
```

For the containers like `llama-cpp` to be able to read the models downloaded by OpenOrch we they must both mount the same docker volume.

An example of this can be seen in the root `docker-compose.yaml` file: `OPENORCH_VOLUME_NAME=singulatron-data`.

So cycle goes like this:

- OpenOrch container writes to `/root/.openorch`, which is mounted to the volume `singulatron-data`
- Assets (which are basically downloaded files) will be passed to containers created by OpenOrch by mounting files in `singulatron-data`.

## `OPENORCH_LLM_HOST`

**This flag is typically unnecessary since OpenOrch retrieves the IP of the Docker bridge. Use it only as a corrective action.**

When OpenOrch is running in a container, it needs to know how to address its siblings (other containers it started):

```sh
Host
 |
 |-> OpenOrch Container
 |-> Container Launched By OpenOrch
```

The `OpenOrch Container` uses the envar `OPENORCH_LLM_HOST` to address `Container Launched By OpenOrch`.

Typically this value should be `172.17.0.1` if you are using the default docker network.

If you are using an other network than default, use `docker network inspect` to find out the IP of your docker bridge for that network.
Usually it's going to be `172.18.0.1`.

This envar is not needed if OpenOrch runs directly on the host:

```sh
Host With OpenOrch
 |
 |-> Container Launched By OpenOrch
```

## `OPENORCH_DB`

You can use this envar to make OpenOrch actually use a database instead of local file storage to store data.

### PostgreSQL

```sh
OPENORCH_DB=postgres
OPENORCH_DB_DRIVER="postgres" # or "mysql"
OPENORCH_DB_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.
