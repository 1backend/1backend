---
sidebar_position: 15
tags:
  - proxy-svc
  - routing
  - reverse proxy
  - edge proxy
  - certifications
  - certs
  - https
  - tls
  - multitenant
---

# Proxy Svc

The Proxy Service in 1Backend handles various forms of proxying:

- **Service proxying**: Proxies requests to custom services registered in the registry.
- **Edge proxying**: When `OB_EDGE_PROXY` is set to true and the appropriate Routes are configured, the proxy service can handle HTTPS requests and perform TLS termination.

> This page provides a high-level overview of `Proxy Svc`. For detailed information, refer to the [Proxy Svc API documentation](/docs/1backend-api/list-routes).

## Service proxying

When you send a request to 1Backend, for example:

```sh
curl 127.0.0.1:11337/user-svc/login ...
```

One of two things will happen:

- If the request matches a built-in service, it is routed there.
- Otherwise, the [registry](/docs/built-in-services/registry-svc) is consulted. If a custom service exists with a slug matching the first path segment (e.g., user-svc), the request is routed to it.

```
Client Request: curl 127.0.0.1:11337/user-svc/login
                        |
                        v
      +-----------------------------------------+
      |  Is "user-svc" a built-in service?      |
      +-----------------------------------------+
                 |                      |
                Yes                    No
                 |                      |
                 v                      v
     +----------------------+   +-----------------------+
     | Route directly to    |   |      Proxy Svc        |
     | built-in service     |   +-----------------------+
     +----------------------+              |
                                           v
                          +-------------------------------+
                          | Look up "user-svc" in registry|
                          +-------------------------------+
                                           |
                                           v
                            +--------------------------+
                            | Proxy to target service  |
                            +--------------------------+
```

## Edge proxying

When edge proxying is enabled, 1Backend also listens on ports 80 (HTTP) and 443 (HTTPS). It can automatically manage HTTPS certificates (requesting and renewing them as needed) for your configured domains.

```sh
$ oo routes list
ROUTE ID                    TARGET
api.singulatron.com         http://1backend:11337
singulatron.com             http://frontend:8080
```

```sh
        +----------------------+
        |  Browser / Client    |
        |  (requests HTTPS)    |
        +----------+-----------+
                   |
           Request to domain
       (e.g. https://singulatron.com)
                   |
                   v
          +------------------------+
          |     Port 443 (TLS)     |
          |   TLS termination by   |
          |     Proxy Svc (edge)   |
          +------------------------+
                   |
           Match domain in Routes
                   |
                   v
     +------------------------------+
     | Route found:                |
     |   singulatron.com           |
     |   → http://frontend:8080    |
     +------------------------------+
                   |
                   v
         +---------------------+
         |   Target container  |
         |   (e.g. frontend)   |
         +---------------------+
```
