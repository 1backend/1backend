---
sidebar_position: 40
tags:
  - registry-svc
  - service-discovery
  - microservices
  - instances
  - nodes
  - distributed
  - service-mesh
  - authentication
  - load-balancing
---

# Registry Svc

Registry Svc is the service discovery backbone for 1Backend. It tracks running service instances and distributed nodes so internal routing, health checks, and node-aware operations can find the right target.

> For detailed API information, refer to the [Registry Svc API documentation](/docs/1backend-api/register-instance).

## Core Responsibilities

- **Instance Management**: Registry of running service instances with heartbeat and status data.
- **Node Discovery**: Topology and resource visibility for 1Backend nodes.
- **Service Discovery**: Instance lookup for services that need to call or proxy other services.
- **Health Monitoring**: Status tracking through periodic instance scans and node heartbeats.

## CLI Usage

### Instance Management

```bash
# List all service instances
oo instances list
oo instances ls
oo i ls

# Remove a service instance
oo instances remove inst_12345
oo instances rm inst_12345
```

### Node Management

```bash
# List all nodes
oo nodes list
oo nodes ls
oo n ls

# Remove a node
oo nodes delete https://node.example.com:11337
oo nodes rm https://node.example.com:11337
```

## API Usage

```bash
# Register or update a service instance
oo put /registry-svc/instance \
  --id=inst_myservice_001 \
  --url=http://my-service:8080 \
  --slug=my-service \
  --status=Healthy

# List registered instances
oo get /registry-svc/instances

# Get current node information
oo get /registry-svc/node/self

# List cluster nodes
oo post /registry-svc/nodes
```

## Instance Records

Instances identify live service processes. A service can register itself with a full URL or with separate scheme, host/IP, port, and path fields. Registry Svc stores heartbeat timestamps, status, details, slug, tags, and the node URL where the instance runs.

Services that want to be routed by Proxy Svc should register an instance with their service slug. Proxy Svc uses Registry Svc to resolve service slugs to live instance URLs.

## Node Records

Nodes describe 1Backend servers in the cluster. Node records include host information, resource usage, GPU information, known processes, and update timestamps. Node discovery is used by built-in services that need node-level awareness.

## Permissions

| Permission | Purpose |
| --- | --- |
| `registry-svc:instance:view` | List service instances |
| `registry-svc:instance:edit` | Register or update service instances |
| `registry-svc:instance:delete` | Remove service instances |
| `registry-svc:node:view` | List and inspect nodes |
| `registry-svc:node:delete` | Remove node records |

Admin users receive registry administration permissions. Regular users can view registry data according to the default permit setup.
