---
sidebar_position: 15
tags:
  - proxy-svc
  - routing
  - reverse-proxy
  - edge-proxy
  - load-balancer
  - certificates
  - https
  - tls
  - acme
  - lets-encrypt
  - multitenant
---

# Proxy Svc

The Proxy Svc is a sophisticated reverse proxy and load balancer that handles both internal service routing and edge traffic management with automatic HTTPS certificate provisioning.

> This page provides a comprehensive overview of `Proxy Svc`. For detailed API information, refer to the [Proxy Svc API documentation](/docs/1backend-api/list-routes).

## Architecture & Purpose

Proxy Svc serves as the **traffic management layer** for 1Backend, providing:

- **Service Proxying**: Routes requests to custom services registered in the Registry Svc
- **Edge Proxying**: Handles external HTTPS traffic with automatic TLS termination
- **Load Balancing**: Intelligent distribution with health-aware routing
- **Certificate Management**: Automatic Let's Encrypt certificate provisioning and renewal
- **Multi-Tenant Routing**: Domain-based routing for multiple applications

### Dual Operation Modes

```mermaid
graph TD
    Client[External Client]

    subgraph "Edge Mode"
        Edge[Edge Proxy :80/:443]
        Routes["Saved Routes: domain/path -> target URL"]
        Targets[Target URLs / Frontends]
        Edge --> Routes --> Targets
    end

    subgraph "Service Mode"
        Internal[Internal Proxy :11337]
        Registry[Registry Svc]
        Services[Backend Services]
        Internal --> |Lookup service slug| Registry --> |Discovered instances| Services
    end

    Client --> Edge
    Client --> Internal
```

## Microfrontends by Path

Proxy Svc supports hosting **multiple frontends under a single domain** by using path-prefix routes. This makes it easy to run a microfrontend architecture without extra gateways.

### Example Routes

- id: "x.localhost"  
  target: "http://root-frontend:8080"

- id: "x.localhost/app"  
  target: "http://app-frontend:3000"

- id: "x.localhost/app/admin"  
  target: "http://admin-frontend:4000"

### How Lookup Works

When a request comes in, Proxy Svc tries the **longest matching prefix**:

1. `/app/admin/settings` → matches `x.localhost/app/admin`  
   → goes to `http://admin-frontend:4000/app/admin/settings`

2. `/app/page` → matches `x.localhost/app`  
   → goes to `http://app-frontend:3000/app/page`

3. `/` or anything else → falls back to `x.localhost`  
   → goes to `http://root-frontend:8080/`

### Visual Flow

```mermaid
flowchart TD
    Client["Browser Request to x.localhost/..."]

    subgraph ProxySvc["Proxy Svc Route Matching"]
      A1["Check host+path: x.localhost/app/admin"]
      A2["Check host+path: x.localhost/app"]
      A3["Fallback: x.localhost"]
    end

    Client --> A1
    A1 -->|If match| Admin["Admin Frontend (4000)"]
    A1 -->|Else| A2
    A2 -->|If match| App["App Frontend (3000)"]
    A2 -->|Else| A3
    A3 --> Root["Root Frontend (8080)"]
```

Benefits:

- Multiple independent frontends can live under one domain
- Deepest path prefix always takes precedence, ensuring /app/admin resolves correctly
- Easy to extend: just add more host+path routes

## CLI Usage

### Route Management

```bash
# List all configured routes
oo routes list

# Save routes from YAML file
oo routes save routes.yaml

# Save routes from directory (processes all .yaml files)
oo routes save ./config/routes/
```

### Route Configuration Files

#### Single Route YAML

```yaml
# api-route.yaml
id: "api.example.com"
target: "http://1backend:11337"
```

#### Multiple Routes YAML

```yaml
# routes.yaml
- id: "api.example.com"
  target: "http://1backend:11337"
- id: "example.com"
  target: "http://frontend:8080"
- id: "cdn.example.com"
  target: "http://nginx:80"
- id: "app.example.com"
  target: "http://react-app:3000"
```

## Service Proxying (Internal Routing)

### How Service Routing Works

When you send a request to 1Backend:

```bash
curl http://127.0.0.1:11337/user-svc/login
```

The routing decision follows this flow:

```mermaid
flowchart TD
    Request["`**Client Request**
    /user-svc/login`"]
    BuiltIn{"`Is 'user-svc'
    a built-in service?`"}
    Registry["`**Registry Lookup**
    Find 'user-svc' instances`"]
    Health{"`Any healthy
    instances?`"}
    LoadBalance["`**Load Balancing**
    Select healthy instance`"]
    Proxy["`**Proxy Request**
    Forward to target`"]
    Error404["`**404 Not Found**
    No instances available`"]
    Direct["`**Direct Route**
    Built-in service`"]

    Request --> BuiltIn
    BuiltIn -->|Yes| Direct
    BuiltIn -->|No| Registry
    Registry --> Health
    Health -->|Yes| LoadBalance
    Health -->|No| Error404
    LoadBalance --> Proxy
```

### Edge Proxy Flow

```mermaid
sequenceDiagram
    participant Browser
    participant EdgeProxy as Edge Proxy :443
    participant ACME as Let's Encrypt
    participant Backend as Target Service

    Note over Browser,Backend: First-time domain setup
    Browser->>EdgeProxy: HTTPS request to myapp.com
    EdgeProxy->>ACME: Request certificate for myapp.com
    ACME-->>EdgeProxy: Certificate + Private Key
    EdgeProxy->>EdgeProxy: Store encrypted certificate
    EdgeProxy->>Backend: Forward request to http://frontend:8080
    Backend-->>EdgeProxy: Response
    EdgeProxy-->>Browser: HTTPS response

    Note over Browser,Backend: Subsequent requests
    Browser->>EdgeProxy: HTTPS request to myapp.com
    EdgeProxy->>EdgeProxy: TLS termination with stored cert
    EdgeProxy->>Backend: Forward request to http://frontend:8080
    Backend-->>EdgeProxy: Response
    EdgeProxy-->>Browser: HTTPS response
```

### Automatic HTTPS Certificates

The edge proxy automatically handles:

- **Certificate Provisioning**: Requests certificates from Let's Encrypt
- **Challenge Handling**: Responds to ACME HTTP-01 challenges on port 80
- **TLS Termination**: Handles SSL/TLS encryption/decryption
- **Certificate Renewal**: Automatically renews certificates before expiry
- **Certificate Storage**: Encrypted storage in the database

```bash
crufter@cruftop:~/1backend$ oo certs ls
CERT ID                     CREATED AT   UPDATED AT            COMMON NAME                 ISSUER   NOT BEFORE   NOT AFTER    SERIAL
singulatron.com             2025-06-08   2025-08-09 14:22:44   singulatron.com             E6       2025-08-09   2025-11-07   34394815…
api.singulatron.com         2025-06-08   2025-08-09 14:15:47   api.singulatron.com         E5       2025-08-09   2025-11-07   83840157…
acme_account+key            2025-06-08   2025-06-08 14:12:42

```
