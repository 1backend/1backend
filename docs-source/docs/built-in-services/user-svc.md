---
sidebar_position: 10
tags:
  - user-svc
  - permissions
  - roles
  - authentication
  - authorization
  - service
  - service to service calls
  - s2s calls
  - multitenant
---

# User Svc

The user service is at the heart of 1Backend, managing users, tokens, organizations, permissions and more. Each service and human on an 1Backend network has an account in the `User Svc`.

> This page provides a high-level overview of `User Svc`. For detailed information, refer to the [User Svc API documentation](/docs/1backend-api/login).

User Svc supports multitenancy: while users are shared globally, tokens, organizations, permits, enrollments, and memberships are isolated by an "app" field (see `Apps` section on this page). This approach allows a single 1Backend instance to securely support multiple web applications.

> Note: Not all services included with 1Backend may support multitenancy. Please refer to the documentation for details and look for the tag `multitenant`.

## Glossary

**Token**: A JWT (JSON Web Token) issued and signed by the User Svc, used to authenticate both human and service accounts.

**Role**: A simple string identifier like `user-svc:user` or `user-svc:org:{orgId}:admin` that represents a specific capability or access level. Roles are embedded in tokens.

**Enroll**: (Enrollment) A mechanism to assign roles to users—both current and future. Enrolls allow roles to be claimed later, once the user joins or logs in.

**Permission**: A string such as `petstore-svc:read`, typically mapping to an API action or endpoint. Roles can bundle multiple permissions.

**Permit**: A mechanism for assigning permissions to users or roles. Permits define who can access what by connecting users or roles with specific permissions.

**Organization**: A way for users to freely associate with each other. Anyone can create organizations and grant membership to others to their organization.

**Membership**: A formal record that links a user to an organization. Memberships determine which organizations a user belongs to and enable organization-scoped roles to take effect.

## Overview

The most important thing about the User Svc is that service (machine) and user (human) accounts look and function the same.

Every service you write needs to [register](/docs/1backend-api/register) at startup, or [log in](/docs/1backend-api/login) with the credentials it saves and manages if it's already registered. Just like a human.

A service account is not an admin account, it's a simple user level account. You might wonder how service-to-service calls work then.

## Service to service calls

Most endpoints on 1Backend can only be called by administrators by default.

Let's take prompting. If you want to let your users prompt AIs you might write a wrapper service called `User Prompter Svc` with the slug `user-prompter-svc`.

If we look at the [Add Prompt endpoint API docs](/docs/1backend-api/prompt), we can see that it mentions

```
Requires the `prompt-svc:prompt:create` permission.
```

To enable your service to call the Add Prompt endpoint, we need to create a permit with your service slug and the permission mentioned above:

```yaml
id: "user-prompter-permit"
permissionId: "prompt-svc:prompt:create"
slugs:
  - "user-prompter-svc"
```

You can apply these permits with an administrator account in your CI workflow with the `oo` CLI:

```sh
oo permit save user-prompter-permit.yaml
```

## Auth patterns

### Role-based access

- **Role-Only Checks**: Authorize users based solely on their assigned roles. This is the simplest method—no need to check individual permissions.

### Permission-based access

- **API Permission Check**: Use the `Has Permission` endpoint with the user's authentication headers and a permission ID to verify access dynamically. This endpoint is designed to be easy to cache (it has no other params apart from the caller header and a permission).

Permission-based checks offer more nuanced control than simple role-only checks—permits can grant specific permissions to slugs, roles and more.

> If you are looking at restricting access to endpoints in other ways, you might be interested in: [Policy Svc](/docs/built-in-services/policy-svc).

## Tokens

The User Svc produces a JWT ([JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token)) upon [/user-svc/login](/docs/1backend-api/login) in the `token.token` field (see the response documentation).

You can either use this token as a proper JWT - parse it and inspect the contents, or you can just use the token to read the user account that belongs to the token with the [/user-svc/self](/docs/1backend-api/read-self) endpoint.

### Verifying a token

The [`/user-svc/public-key`](/docs/1backend-api/get-public-key) will return you the public key of the User Svc which then you can use that to verify the token.

Use the JWT libraries that are available in your programming language to do that, or use the Singularon [SDK](https://github.com/1backend/1backend/tree/main/sdk) if your language is supported.

### Automatic token refresh

1Backend tokens are valid for a limited time (see `OB_TOKEN_EXPIRATION`). Once a token expires, 1Backend can either automatically refresh it (this is the default behaviour) or reject it based on configuration (see `OB_TOKEN_AUTO_REFRESH_OFF`).

If automatic token refresh is disabled, clients are responsible for detecting expiration and refreshing the token themselves.

If automatic refresh is enabled, expired tokens can still be reused indefinitely. Behind the scenes, 1Backend maps old tokens to their most recent valid version.

#### Example flow

To understand how automatic token refresh works in practice, consider the following scenario:

- A user acquires a token.
- The token is valid for X minutes.
- After expiration, Service A receives a request with the old token.
- Service A then has two options:
  1. Call the User Svc RefreshToken endpoint on every request to get a new token — which undermines the stateless nature of JWTs.
  2. Cache the refreshed token and continue accepting the expired one, internally mapping it to the latest valid token without calling 1Backend. When the refreshed token expires, this process repeats.

#### Token Pruning

You might wonder: if an old token keeps getting refreshed indefinitely, does that mean a new token is minted every `OB_TOKEN_EXPIRATION` interval — and do they pile up forever?

While a new token is issued on each refresh, the system keeps track of which tokens are actively being refreshed and discards the rest. At any given time, a maximum of **three tokens per device** (see the `device` field in the token) are retained:

- The currently active token
- The two most recently refreshed tokens (kept as a buffer to handle clock drift or retries)

All other older tokens are pruned to avoid unbounded growth.

### Token structure

The structure of the JWT is the following:

```yaml
# User Id
oui: usr_dC4K75Cbp6

# Slug
osl: test-user-slug-0

# Roles
oro:
  - user-svc:user
  - user-svc:org:{org_dC4K7NNDCG}:user
```

The field names are kept short to save space, so perhaps the Go definition is also educational:

```go
type Claims struct {
 UserId  string   `json:"oui"` // `oui`: 1backend user ids
 Slug    string   `json:"osl"` // `osl`: 1backend slug
 Roles []string   `json:"oro"` // `oro`: 1backend roles
 jwt.RegisteredClaims
}
```

## Roles

Roles are simply strings. They are not a database record, they don't have an ID, name etc. They are simple strings, such as `user-svc:admin`.

A user token produced upon login contains all the roles a user has.

> Efficiency Tip: JWT tokens are sent with every request. Keeping the number of roles minimal improves performance.

When checking if a user is authorized, there are a few common patterns to choose from:

### Roles without permissions

Roles are powerful, even without specific permissions attached. One common use case is managing product subscriptions.

Suppose you launch a new product called Funny Cats Newsletter with two subscription tiers: Pro and Ultra. You could create a service with the slug funny-cats-newsletter-svc and define custom static roles for each tier:

```yaml
funny-cats-newsletter-svc:pro
funny-cats-newsletter-svc:ultra
```

By checking if these roles exist in a user's token, you can authorize access to product-specific features. These roles can be created dynamically by calling the Create Role endpoint.

### Roles containing dynamic data

You are free to make up your own roles which might even have dynamic data, just like the User Svc did with the organization ids.

Example:

```sh
user-svc:org:{org_dBZRCej3fo}:admin
user-svc:org:{org_dBZRCej3fo}:user
```

By convention these dynamic values are enclosed in {}. In this example, roles are assigned per organization. For more details, see the Organizations section.

### Owning roles vs. having roles

In many endpoints such `SaveEnrolls`, the topic of "role ownership" comes up. The basic problem is simple: just because someone has a role, it doesn't mean they can also bestow that role upon other users. In simple terms, if an admin makes someone a user, that user should not be able to make others users, as that is the privilege of the admins.

A user "owns" a role in the following cases:

- The role is prefixed with the caller's slug. For example, a user with the slug `joe-doe` owns roles like `joe-doe:any-custom-role`.
- User "owns" all roles that share a prefix if they hold the corresponding `:admin` role. For example, having `user-svc:org:{org_id}:admin` means the user owns roles like `user-svc:org:{org_id}:user` and `user-svc:org:{org_id}:viewer`.

By enforcing role ownership rules, the system ensures that roles are only assigned by authorized users, preventing privilege escalation and maintaining security within the organization.

### Conventions

Each role created must by prefixed by the slug of the account that created it. Said account becomes the owner of the role and only that account can edit the role.

## Organizations

In the dynamic roles section we already talked about organizations, lets elaborate on them here a bit.

```yaml
id: "org_eZqC0BbdG2"
name: "Acme Corporation" # Full name of the organization
slug: "acme-corporation" # URL-friendly unique identifier for the organization
createdAt: "2025-01-15T12:00:00Z" # Example ISO 8601 timestamp
```

### Access rules

#### Create

Any logged in user can create an organization, provided the `Organization` slug is not taken yet. The creator becomes the first admin of the organization, acquiring the role of `user-svc:org:{orgId}:admin` role.

#### Membership

Admins can assign other users member (`user-svc:org:{orgId}:user`) or admin roles (`user-svc:org:{orgId}:admin`) for the organizations they administer.

### Use cases

Organizations let users to freely associate with each other and create groups. Think about Discord servers, Slack workspaces, Github organizations etc.

## Permissions

### Conventions

Each permission created must by prefixed by the slug of the account that created it. Said account becomes the owner of the permission and only that account can add the permission to a role.

> Once you (your service) own a permission (by creating it, and it being prefixed by your account slug), you can add it to any role, not just roles owned by you.

Example; let's say your service is `petstore-svc`. 1Backend prefers fine-grained access control, so you are free to define your own permissions, such as `petstore-svc:read` or `petstore-svc:pet:read`.

## Services with multiple nodes

You might now wonder what happens when a service has multiple instances/nodes. Won't their user accounts "clash" in the `User Svc`? The answer to this is that from the `User Svc` point of view, each node/instance of a service is the same account.

This is possible because the platform is designed with services having a "Shared Database Access".

Let's say you have a Cassandra network that spans multiple Availability Zones/Regions. Your nodes will also span multiple AZs/Regions and each instance of them will log in as `X Svc`.

## Apps

Apps are the foundation of multitenancy in 1Backend. They act as namespaces that allow you to serve multiple clients, projects, or websites using the same set of microservices.

It’s recommended to use the website's hostname (e.g., `socks.com`, `shoes.com`) as the app name. This helps clearly separate environments while remaining easy to manage.

The current app is typically inferred from the authentication token. For public endpoints, it’s derived from the request host. As a result, even if a service API doesn’t explicitly mention "App" in its endpoints, it might still support multitenancy.

On this documentation page, services marked with the `multitenancy` tag—such as **User Svc**, **Config Svc**, and **Secret Svc**—support apps. These services provide the essential building blocks for building robust, multi-tenant applications.

## CLI Reference

### Quick Start

```bash
# Register a new user
oo register alice
Enter password: [hidden]

# Login with existing credentials
oo login alice
Enter password: [hidden]

# Check current authentication status
oo whoami

# View authentication token
oo token

# Switch between users (if multiple logged in)
oo use bob
```

### Authentication Commands

#### `oo register` - Create New Accounts

**Usage:**

```bash
oo register [slug] [password]
# or
oo register [slug]
Enter password: [hidden]

oo register [--contact-id email] [--contact-platform email]
```

**Examples:**

```bash
# Interactive registration (secure - no terminal history)
oo register alice
Enter password: [hidden]

# User with contact information
oo register john-doe --contact-id john@company.com --contact-platform email

# Direct registration (automation only - avoid for security)
oo register test-user test-password
```

#### `oo login` - Authenticate

**Usage:**

```bash
oo login [slug] [password]
```

**Examples:**

```bash
# Interactive login (recommended)
oo login alice
Enter password: [hidden]

# Quick development login (avoid in production)
oo login dev-user dev-password
```

#### `oo whoami` - Identity Information

**Usage:**

```bash
oo whoami [--all]
```

**Examples:**

```bash
# Current user information
oo whoami

# All logged in users
oo whoami --all

# Example output:
# id: usr_abc123
# slug: alice
# roles:
#   - user-svc:user
#   - user-svc:org:org_xyz789:admin
#   - payment-svc:processor
```

#### `oo use` - Switch User Context

**Examples:**

```bash
# Switch to different authenticated user
oo use payment-svc
oo use alice
oo use admin-user

# Verify switch
oo whoami
```

#### `oo token` - Access Tokens

```bash
# Get current authentication token
oo token

# Use in API calls
curl -H "Authorization: Bearer $(oo token)" \
  https://api.1backend.com/user-svc/self
```

### User Management Commands

#### `oo user list` - Browse Users

**Usage:**

```bash
oo user list [--userId id] [--contactId email] [--limit count]
```

**Examples:**

```bash
# List all users (admin required)
oo user list

# Find specific user
oo user list --userId usr_abc123

# Find by email
oo user list --contactId alice@company.com

# Paginated results
oo user list --limit 50
```

### Permission Management

#### `oo permit save` - Grant Permissions

**Usage:**

```bash
oo permit save <permit-file.yaml>
oo permit save <permits-directory>
```

**Single Permit Example:**

```yaml
# api-access-permit.yaml
id: "payment-api-access"
permissionId: "payment-svc:process"
slugs:
  - "order-service"
  - "subscription-service"
roles:
  - "payment-svc:processor"
```

**Multiple Permits Example:**

```yaml
# service-permissions.yaml
- id: "chat-read-permit"
  permissionId: "chat-svc:message:read"
  slugs: ["frontend-app", "mobile-app"]
  
- id: "chat-write-permit"
  permissionId: "chat-svc:message:create"
  roles: ["chat-svc:user"]
  
- id: "admin-chat-permit"
  permissionId: "chat-svc:admin"
  roles: ["user-svc:admin"]
```

**Apply Permits:**

```bash
# Save single permit
oo permit save api-access-permit.yaml

# Save multiple permits
oo permit save service-permissions.yaml

# Save from directory
oo permit save permissions/production/
```

#### `oo permit list` - View Permissions

```bash
# List all permits (admin required)
oo permit list

# Example output:
# PERMIT ID          PERMISSION              SLUGS               ROLES
# payment-api-access payment-svc:process     order-service       payment-svc:processor
# chat-read-permit   chat-svc:message:read   frontend-app        
```

### Role Enrollment

#### `oo enroll save` - Assign Roles

**Usage:**

```bash
oo enroll save <role> --userId <id>
oo enroll save <role> --contactId <email>
oo enroll save <enroll-file.yaml>
```

**Direct Enrollment:**

```bash
# Enroll user to admin role
oo enroll save user-svc:admin --userId usr_abc123

# Enroll by email (future user)
oo enroll save payment-svc:processor --contactId alice@company.com

# Organization role
oo enroll save user-svc:org:org_xyz789:admin --userId usr_def456
```

**File-Based Enrollment:**

```yaml
# team-enrollments.yaml
- id: "admin-enrollment-1"
  role: "user-svc:admin"
  contactId: "admin@company.com"

- id: "dev-team-enrollment"
  role: "dev-team:developer"
  contactId: "developer@company.com"

- id: "payment-processor-enrollment"
  role: "payment-svc:processor"
  userId: "usr_abc123"
```

```bash
# Apply enrollments
oo enroll save team-enrollments.yaml
```

#### `oo enroll list` - View Enrollments

**Usage:**

```bash
oo enroll list [--role role] [--userId id] [--contactId email]
```

```bash
# Apply service permissions
oo permit save payment-service-permissions.yaml
```
