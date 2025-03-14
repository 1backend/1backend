---
sidebar_position: 40
tags:
  - registry-service
  - registry
  - microservices
  - addresses
  - authentication
---

# Registry Svc

The registry service is designed to maintain a database of service definitions, service instances and nodes.

> This page provides a high-level overview of `Registry Svc`. For detailed information, refer to the [Registry Svc API documentation](/docs/openorch/register-instance).

## Entities

### Definition

A `Definition` or service definition consists of the following things:

- An ID
- A container image to run and some additional information like the internal port exposed etc.
- A set of endpoint definitions (OpenAPI etc.)
- The URL of different clients (JS, Go etc.)

A `Definition` is an abstract concept that can not be called. For a callable entity look at `Instance`s. Definitions are basically things you can deploy as an instance with a deployment.

#### Container based definition

```yaml
id: test-a
image:
  name: hashicorp/http-echo
  port: 8080
hostPort: 8887
```

#### Notes

HostPorts are a temporary requirement until support for dynamic port assignment lands.

#### Source code based definition

```yaml
id: test-b
repository:
  url: https://github.com/openorch/openorch.git
  containerFile: server/docker/Dockerfile
hostPort: 9998
```

### Instance

A `Instance` or a service instance is an actual running, callable instance of a `Definition`.

A `Instance` consists of the following things:

- Address information that can be used to internally address the `Instance`.
- A Deployment ID.

### Deployment

Definitions become instances through the [Deployment entity of the Deploy Service](/docs/built-in-services/deploy-svc).

### Node

A `Node` is a physical or virtual machine that runs a OpenOrch server.
Maintaining a list of nodes is important so the daemon can efficiently distribute workload across the nodes. It's the basis for all distributed features.

This is how a well configured node should look like:

```sh
$ oo nodes lis
NODE ID    URL                                  LAST HEARTBEAT
myNodeId   http://myNetworkInternalHost:58231   8s ago
```

For well configured nodes, the following must be present

- Each node should have a unique URL (eg. 127.0.0.1 for all nodes is not unique...)
- Each node URL should be addressable by every node including themselves
- A node ID should be defined to avoid the more error prone ID generation

Here is how a suboptimally configured node looks like:

```sh
$ oo nodes lis
NODE ID           URL                         LAST HEARTBEAT
node_eIfnt9CGJV   http://127.0.0.1:58231      8s ago
```

For single node setups that can work, but not if you plan to use distributed features.

To configure the nodes, please see the `OPENORCH_URL` and `OPENORCH_NODE_ID` envars [here](/docs/running-the-server/backend-environment-variables)

## How It Works

The registry is needed when you want to call services not included in the OpenOrch server. You can think of the daemon as the standard library and services in the registry as third party libraries.

When you want to call a service, you can ask the registry to provide you with a list of instance addresses for a service by service slug. Then you can use any of those instance addresses to make a call.

Things like load balancing should be done on the client side at the moment, the damon does not provide a Proxy Svc yet.
