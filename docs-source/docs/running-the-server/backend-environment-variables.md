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

## `OB_CONTACT_EMAIL`

Specifies the system-wide contact email address for operational and administrative use.

This email is used in various backend components, including:

- ACME TLS certificate registration (e.g., Let's Encrypt)
- System notifications or future alerting mechanisms
- Recovery options for certain services that require a fallback contact

While optional, it is strongly recommended to set this in production environments so you can receive:

- Expiration or renewal notices for HTTPS certificates
- Warnings about rate limits or configuration issues
- Future administrative alerts from the system

This address will not be used for login or user account purposes.

## `OB_DB`

You can use this envar to make 1Backend actually use a database instead of local file storage to store data.

## `OB_DB_PREFIX`

When specified, all tables in the database will be prefixed by this strings. Mostly useful for testing.

### PostgreSQL

```sh
OB_DB=postgres
OB_DB_DRIVER="postgres" # or "mysql"
OB_DB_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.

## `OB_EDGE_PROXY`

When set to true, 1Backend will enable the edge proxy feature. This feature configures the system to listen for incoming HTTP and HTTPS traffic on ports 80 and 443, respectively.

The edge proxy acts as a public-facing reverse proxy, handling domain-based routing and TLS termination for external requests. It is typically used to:

- Serve ACME HTTP-01 challenges (for automated TLS certificates, e.g., via Let's Encrypt) on port 80.
- Handle public HTTPS traffic on port 443, routing incoming domain-based requests to appropriate backends or services based on their domain.

When `OB_EDGE_PROXY` is not set to true, 1Backend will not start these public-facing routers. Only the internal API server on the `OB_SERVER_URL` port (default: 11337) will be active.

**Typical Use Case**:

Use this flag when 1Backend is running as a publicly accessible server that needs to:

- Terminate TLS (HTTPS) at the edge.
- Serve automated certificates via ACME.
- Route external requests based on domain names.

## `OB_EDGE_PROXY_TEST_MODE`

Enables test mode for the edge proxy:

- Turns off autocert.
- Uses HTTP for both the HTTP and HTTPS port.

## `OB_EDGE_PROXY_HTTP_PORT`

Sets the HTTP port for the edge proxy.
In production, this must be `80`. Any other value is only for testing purposes.

## `OB_EDGE_PROXY_HTTPS_PORT`

Sets the HTTPS port for the edge proxy.
In production, this must be `443`. Any other value is only useful in testing.

## `OB_ENCRYPTION_KEY`

This key is used in the Secret Svc so secrets are encrypted at rest.

## `OB_FOLDER`

When specified, all data (uploads, downloads, image resize caches, models etc.) will be stored in this folder. Defaults to `~/.1backend`.

## `OB_GPU_PLATFORM`

This envar is used to enabel GPU acceleration.
Supported platforms:

- `cuda`

Do not set this if your card doesn't support the given architecture or things will break.

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

## `OB_NODE_ID`

For information about this, please refer to the [Registry Svc Node section](/docs/built-in-services/registry-svc#node)

## `OB_SERVER_URL`

The OB_SERVER_URL is the internally addressable (non-public-facing) URL of an 1Backend server. It should point to the local 1Backend instance on each physical node. Ideally, every node should have its own 1Backend instance.

This envar should be set only for microservices built on 1Backend. The 1Backend server itself should use `OB_SELF_URL`.

## `OB_SELF_URL`

Microservices use this to register themselves in the 1Backend registry. The 1Backend server uses this to address itself.

## `OB_TEST`

Microservices and the 1Backend server uses this envar to detect if they are running as part of a test.

When set to true, subsystems act accordingly: for example the datastore will prefix tables with random numbers to provide a unique and clean environment for each test.

First startup is also significantly faster when this flag is enabled, as 1Backend uses bcrypt.MinCost instead of bcrypt.DefaultCost for password generation, and sets the RSA key size to 512 bits in test mode instead of the default 4096.

## `OB_TOKEN_AUTO_REFRESH_OFF`

When set to true, clients are responsible for handling the refresh of expired tokens themselves. This overrides the default behavior, where expired tokens are automatically accepted and refreshed by the system.

## `OB_TOKEN_EXPIRATION`

Specifies the duration before a token expires. Use formats like 5m for five minutes, 10h for ten hours, etc.

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
