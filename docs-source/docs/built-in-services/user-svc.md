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
---

# User Svc

The user service is at the heart of 1Backend, managing users, tokens, organizations, permissions and more. Each service and human on an 1Backend network has an account in the `User Svc`.

> This page provides a high-level overview of `User Svc`. For detailed information, refer to the [User Svc API documentation](/docs/1backend/login).

## Glossary

**Token**: A JWT (JSON Web Token) issued and signed by the User Svc, used to authenticate both human and service accounts.

**Role**: A simple string identifier like `user-svc:user` or `user-svc:org:{orgId}:admin` that represents a specific capability or access level. Roles are embedded in tokens.

**Invite**: A way to assign roles to users—both current and future. Invites allow roles to be claimed later, once the user joins or logs in.

**Permission**: A string such as `petstore-svc:read`, typically mapping to an API action or endpoint. Roles can bundle multiple permissions.

**Grant**: A mechanism for assigning permissions to users or roles. Grants define who can access what by connecting users or roles with specific permissions.

## Overview

The most important thing about the User Svc is that service (machine) and user (human) accounts look and function the same.

Every service you write needs to [register](/docs/1backend/register) at startup, or [log in](/docs/1backend/login) with the credentials it saves and manages if it's already registered. Just like a human.

A service account is not an admin account, it's a simple user level account. You might wonder how service-to-service calls work then.

## Service to service calls

Most endpoints on 1Backend can only be called by administrators by default.

Let's take prompting. If you want to let your users prompt AIs you might write a wrapper service called `User Prompter Svc` with the slug `user-prompter-svc`.

If we look at the [Add Prompt endpoint API docs](/docs/1backend/prompt), we can see that it mentions

```
Requires the `prompt-svc:prompt:create` permission.
```

To enable your service to call the Add Prompt endpoint, we need to create a grant with your service slug and the permission mentioned above:

```yaml
id: "user-prompter-grant"
permissionId: "prompt-svc:prompt:create"
slugs:
  - "user-prompter-svc"
```

You can apply these grants with an administrator account in your CI workflow with the `oo` CLI:

```sh
oo grant save user-prompter-grant.yaml
```

## Auth patterns

### Role-Based Access

- **Role-Only Checks**: Authorize users based solely on their assigned roles. This is the simplest method—no need to check individual permissions.

### Permission-Based Access

- **API Permission Check**: Use the `Has Permission` endpoint with the user's authentication headers and a permission ID to verify access dynamically.

- **Cached Role Permissions**: Store `Grant`s locally and check only the user's token for the required role. This is faster and avoids API calls but requires a bit more setup.
  - An SDK can help simplify this.
  - If you need to verify a token manually, refer to the `Get Public Key` endpoint.

> If you are looking at restricting access to endpoints in other ways, you might be interested in: [Policy Svc](/docs/built-in-services/policy-svc).

## Tokens

The User Svc produces a JWT ([JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token)) upon [/user-svc/login](/docs/1backend/login) in the `token.token` field (see the response documentation).

You can either use this token as a proper JWT - decode it and inspect the contents, or you can just use the token to read the user account that belongs to the token with the [/user-svc/user/by-token](/docs/1backend/read-user-by-token) endpoint.

### Decoding a token

The [`/user-svc/public-key`](/docs/1backend/get-public-key) will return you the public key of the User Svc which then you can use that to decode the token.

Use the JWT libraries that are available in your programming language to do that, or use the Singularon [SDK](https://github.com/1backend/1backend/tree/main/sdk) if your language is supported.

### Token structure

The structure of the JWT is the following:

```yaml
# User Id
oui: usr_dC4K75Cbp6

# Slug
olu: test-user-slug-0

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

In many endpoints such `SaveInvites`, the topic of "role ownership" comes up. The basic problem is simple: just because someone has a role, it doesn't mean they can also bestow that role upon other users. In simple terms, if an admin makes someone a user, that user should not be able to make others users, as that is the privilege of the admins.

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
