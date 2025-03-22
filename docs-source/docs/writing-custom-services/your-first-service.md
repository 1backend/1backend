---
sidebar_position: 3
tags:
  - test
---

# Your first service

While 1Backend itself is written in Go, services that run on 1Backend can be written in any language.
A service only needs a few things to fully function:

- Register a user account, just like a human user. For details, see the [User Svc](/docs/built-in-services/user-svc).
- Register its instance in the registry so 1Backend knows where to route requests.

## A Go example

The following Go service demonstrates these steps:

- Registers itself as a user with the slug `skeleton-svc`
- Registers or updates its URL (`http://127.0.0.1:9311`) in the [Registry](/docs/built-in-services/registry-svc).

You may notice that the following code uses a "Go SDK," but it's simply a set of convenience functions built on top of the 1Backend API.
1Backend is language-agnostic and can be used with any language, even if no SDK is available in the repository.

```go
// <!-- INCLUDE: ./first-service-go/main.go -->

// <!-- /INCLUDE -->
```

Just make sure you run it with the appropriate envars:

```sh
OB_URL=http://127.0.0.1:58231 SELF_URL=http://127.0.0.1:9311 go run main.go
```

Once it's running you will be able to call the 1Backend server proxy and that will proxy to your skeleton service:

```sh
# 127.0.0.1:58231 here is the address of the 1Backend server
$ curl 127.0.0.1:58231/skeleton-svc/hello
{"hello": "world"}
```

This is so you don't have to expose your skeleton service to the outside world, only your 1Backend server.

Let's recap how the proxying works:

- Service registers an account, acquires the `skeleton-svc` slug.
- Service calls the 1Backend [Registry Svc](/docs/built-in-services/registry-svc) to tell the system an instance of the Skeleton service is available under the URL `http://127.0.0.1:9311`
- When you curl the 1Backend server with a path like `127.0.0.1:58231/skeleton-svc/hello`, the first section of the path will be a user account slug. The daemon checks what instances are owned by that slug and routes the request to one of the instances.

```sh
$ oo instance ls
ID                URL                     STATUS    OWNER SLUG       LAST HEARTBEAT
inst_eHFTNvAlk9   http://127.0.0.1:9311   Healthy   skeleton-svc     10s ago
```

## Things to understand

### Instance registration

Like most other things on the platform, service instances become owned by a user account slug. When the skeleton service calls [RegisterInstance](/docs/1backend/register-instance), the host will be associated with the `skeleton-svc` slug.

Updates to this host won't be possible unless the caller is the skeleton service (or the caller is an admin). The service becomes the owner of that URL essentially.

This is the same ownership model like in other parts of the 1Backend system.
