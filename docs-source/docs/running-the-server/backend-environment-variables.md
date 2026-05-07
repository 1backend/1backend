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

This page lists the `OB_*` environment variables used by the 1Backend server, SDK boot helpers, telemetry setup, and local test harness.

## `OB_AZ`

Sets the node availability zone metadata used by server/node registration code.

For clustered or multi-node deployments, set this to the physical or cloud availability zone where the server instance runs, for example `us-east-1a`.

## `OB_BOOTSTRAP_FOLDER`

Points 1Backend at a file or directory of startup manifests to apply after the built-in services have started.

```sh
OB_BOOTSTRAP_FOLDER=/etc/1backend/bootstrap
```

This is intended for infrastructure-as-code bootstrapping, such as preview environments that need service-to-service permits, routes, configs, secrets, or policy instances before the application stack starts serving traffic.

The loader applies supported manifests recursively. It understands folder-level entity defaults from `_meta.yaml` files and file-level `_meta` blocks.

Folder names do not set `appHost`. App-scoped manifests must specify `appHost` explicitly. Use `appHost: "*"` for all apps. Apps referenced by `appHost` are created automatically when needed.

Supported entities:

- `user-svc:permit`
- `user-svc:enroll`
- `proxy-svc:route`
- `proxy-svc:redirect`
- `config-svc:config`
- `secret-svc:secret`
- `policy-svc:instance`

Secrets are loaded through Secret Svc. Plaintext values are encrypted before storage; encrypted values can be supplied directly with their checksum fields. Keep bootstrap folders protected and prefer encrypted secret manifests for production deployments.

Example:

```yaml
# /etc/1backend/bootstrap/permits/_meta.yaml
entity: "user-svc:permit"
```

```yaml
# /etc/1backend/bootstrap/permits/report-svc-list-users.yaml
id: "report-svc-list-users"
appHost: "*"
permission: "user-svc:user:view"
slugs:
  - "report-svc"
```

```yaml
# /etc/1backend/bootstrap/apps/tenant.example.com/enrolls/site-owner.yaml
id: "site-owner"
appHost: "tenant.example.com"
contactId: "owner@example.com"
role: "site-svc:admin"
```

```yaml
# /etc/1backend/bootstrap/redirects/old-docs.yaml
id: "example.com/old-docs"
target: "https://example.com/docs"
statusCode: 308
```

```yaml
# /etc/1backend/bootstrap/apps/tenant.example.com/configs/payment.yaml
id: "paymentSvc"
appHost: "tenant.example.com"
data:
  publicKey: "pk_test"
```

```yaml
# /etc/1backend/bootstrap/apps/tenant.example.com/secrets/api-key.yaml
id: "api-key"
appHost: "tenant.example.com"
value: "encrypted-secret"
encrypted: true
checksum: "12345678"
checksumAlgorithm: "CRC32"
```

```yaml
# /etc/1backend/bootstrap/policies/login-rate-limit.yaml
id: "login-rate-limit"
endpoint: "/user-svc/login"
templateId: "rate-limit"
parameters:
  rateLimit:
    maxRequests: 20
    timeWindow: "1m"
    entity: "ip"
    scope: "endpoint"
```

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

## `OB_DB_CONNECTION_STRING`

Connection string for the primary/write database.

For PostgreSQL:

```sh
OB_DB=postgres
OB_DB_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.

## `OB_DB_PREFIX`

When specified, all tables in the database will be prefixed by this strings. Mostly useful for testing.

## `OB_DB_READ_CONNECTION_STRING`

Optional connection string for a read-replica database.

When set, datastore reads can use this connection while writes continue to use `OB_DB_CONNECTION_STRING`. Leave it unset when your environment only has a primary database.

### PostgreSQL

```sh
OB_DB=postgres
OB_DB_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
OB_DB_READ_CONNECTION_STRING="postgres://postgres:mysecretpassword@localhost:5432/mydatabase?sslmode=disable"
```

Naturally, you should change the details of the connection string to reflect your environment.

## `OB_EDGE_CACHE_ITEM_MAX_SIZE`

Sets the maximum size of a single edge-proxy response that may be cached.

Use byte-size strings understood by Docker's unit parser, such as `512KB`, `2MB`, or `50MB`.

## `OB_EDGE_CACHE_MAX_SIZE`

Sets the total edge-proxy cache size.

Use byte-size strings such as `100MB` or `1GB`. Leave it unset to use the server default.

## `OB_EDGE_PROXY`

When set to true, 1Backend will enable the edge proxy feature. This feature configures the system to listen for incoming HTTP and HTTPS traffic on ports 80 and 443, respectively.

```sh
OB_EDGE_PROXY=true
```

The edge proxy acts as a public-facing reverse proxy, handling domain-based routing and TLS termination for external requests. It is typically used to:

- Serve ACME HTTP-01 challenges (for automated TLS certificates, e.g., via Let's Encrypt) on port 80.
- Handle public HTTPS traffic on port 443, routing incoming domain-based requests to appropriate backends or services based on their domain.

When `OB_EDGE_PROXY` is not set to true, 1Backend will not start these public-facing routers. Only the internal API server on `OB_PORT` (default: 11337) will be active.

**Typical Use Case**:

Use this flag when 1Backend is running as a publicly accessible server that needs to:

- Terminate TLS (HTTPS) at the edge.
- Serve automated certificates via ACME.
- Route external requests based on domain names.

## `OB_EDGE_PROXY_TEST_MODE`

Enables test mode for the edge proxy. Contrary to the production edge proxy, the dev mode:

- Does not try to redirect to 443
- Does no certificate requests and renewal
- Routes HTTP requests (vs HTTP for prod).

## `OB_EDGE_PROXY_HTTP_PORT`

Sets the HTTP port for the edge proxy.
In production, this must be `80`. Any other value is only for testing purposes.

## `OB_EDGE_PROXY_HTTPS_PORT`

Sets the HTTPS port for the edge proxy.
In production, this must be `443`. Any other value is only useful in testing.

## `OB_ENCRYPTION_KEY`

This key is used in the Secret Svc so secrets are encrypted at rest.

## `OB_ENV`

Sets the deployment environment label used by telemetry resource attributes.

Examples: `dev`, `preview`, `staging`, `prod`.

## `OB_FILE_GCS`

When set to true, enables Google Cloud Storage-backed file storage.

This requires `OB_GCP_SA_KEY` and `OB_GCS_BUCKET`.

## `OB_FOLDER`

When specified, all data (uploads, downloads, image resize caches, models etc.) will be stored in this folder. Defaults to `~/.1backend`.

## `OB_GCP_SA_KEY`

Path to the Google Cloud service account key file used when `OB_FILE_GCS=true`.

Treat the file contents as secret. Prefer mounting it from your secret manager rather than committing it into configuration files.

## `OB_GCS_BUCKET`

Google Cloud Storage bucket name used when `OB_FILE_GCS=true`.

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

For information about this, please refer to the [Registry Svc Node section](/docs/built-in-services/registry-svc#node-management--cluster-topology)

## `OB_OTEL_DISABLED`

Disables OpenTelemetry setup and HTTP metrics instrumentation when set to true.

```sh
OB_OTEL_DISABLED=true
```

## `OB_OTEL_EXPORTER_OTLP_ENDPOINT`

Sets the OTLP HTTP exporter endpoint used by telemetry.

If the standard `OTEL_EXPORTER_OTLP_ENDPOINT` variable is already set, that standard variable takes precedence.

## `OB_OTEL_METRICS_PATH`

Sets the Prometheus-format metrics route. Defaults to `/metrics`.

For service-specific SDK instrumentation, this path is combined with the service name where appropriate.

## `OB_OTEL_TRACES`

When set to true, enables OTLP trace export.

## `OB_OTEL_TRACE_SAMPLE_RATIO`

Sets the trace sampling ratio when traces are enabled.

Use a number from `0` to `1`, for example `0.1` for 10% sampling or `1` for all traces.

## `OB_PORT`

Sets the port used by the 1Backend server's internal API listener. Defaults to `11337`.

If `OB_PORT` is not set, the server also derives the listen port from `OB_SELF_URL` when that URL includes a port.

## `OB_REGION`

Region metadata used by test-launched services and custom service processes that consume it.

The built-in 1Backend server currently reads `OB_AZ` for its own region-related option as well as availability-zone metadata, so set `OB_AZ` for the server itself.

## `OB_VERIFY_CONTACTS`

When set to true, registration of new users with a contact address requires an OTP to prove ownership of the contact.

## `OB_VERSION`

Sets the service version label used by telemetry.

## `OB_SERVER_URL`

The OB_SERVER_URL is the internally addressable (non-public-facing) URL of an 1Backend server. It should point to the local 1Backend instance on each physical node. Ideally, every node should have its own 1Backend instance.

This envar should be set only for microservices built on 1Backend. The 1Backend server itself should use `OB_SELF_URL`.

## `OB_SELF_URL`

Microservices use this to register themselves in the 1Backend registry. The 1Backend server uses this to address itself.

## `OB_SYNC_CERTS_TO_FILES`

Also save certificates to disk.

By default, certificates are stored in the internal datastore managed by the proxy service. However, in some cases — for example, when other containers need to access the certificates — it's useful to have them available on the filesystem as well.

When this flag is set to true, the server will write certificates to the `OB_FOLDER/certs` directory (by default, `~/.1backend/certs`) using a standard directory layout compatible with common infrastructure software like Let's Encrypt (Certbot), NGINX, and Apache.

### Directory Structure

Certificates are organized under:

```sh
<OB_FOLDER>/certs/live/<domain>/
```

Within each domain directory, files are saved as:

- `cert.pem` — the domain’s public certificate
- `privkey.pem` — the private key corresponding to the certificate
- `chain.pem` — intermediate certificates forming the chain (if any)
- `fullchain.pem` — concatenation of cert.pem and chain.pem

This structure aligns with what many TLS tools and web servers expect, making it easy to integrate with existing systems that rely on file-based certificates.

### Example

For the domain `singulatron.com`, files will be saved as:

```
~/.1backend/certs/live/singulatron.com/cert.pem
~/.1backend/certs/live/singulatron.com/privkey.pem
~/.1backend/certs/live/singulatron.com/chain.pem
~/.1backend/certs/live/singulatron.com/fullchain.pem
```

This standard layout ensures compatibility with web servers like NGINX and Apache, as well as automation tools like Certbot, enabling easy mounting of certificates into containers or configuring HTTPS endpoints.

## `OB_TEST`

Microservices and the 1Backend server uses this envar to detect if they are running as part of a test.

When set to true, subsystems act accordingly: for example the datastore will prefix tables with random numbers to provide a unique and clean environment for each test.

First startup is also significantly faster when this flag is enabled, as 1Backend uses bcrypt.MinCost instead of bcrypt.DefaultCost for password generation, and sets the RSA key size to 512 bits in test mode instead of the default 4096.

## `OB_TOKEN_AUTO_REFRESH_OFF`

When set to true, clients are responsible for handling the refresh of expired tokens themselves. This overrides the default behavior, where expired tokens are automatically accepted and refreshed by the system.

## `OB_TOKEN_EXPIRATION`

Specifies the duration before a token expires. Use formats like 5m for five minutes, 10h for ten hours, etc.

## `OB_TEST_SERVER_BIN`

Test harness only.

When integration tests launch a 1Backend server binary, this can point at the exact binary path to execute.

## `OB_TEST_<SERVICE>_BIN`

Test harness only.

Overrides the binary path for one named service when tests launch service processes. The service name is uppercased and dashes are replaced with underscores.

Example:

```sh
OB_TEST_USER_SVC_BIN=/tmp/user-svc
```

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
