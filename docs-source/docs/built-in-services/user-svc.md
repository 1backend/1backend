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

The user service is at the heart of OpenOrch, managing users, tokens, organizations, permissions and more. Each service and human on an OpenOrch network has an account in the `User Svc`.

> This page provides a high-level overview of `User Svc`. For detailed information, refer to the [User Svc API documentation](/docs/openorch/login).

## Overview

The most important thing about the User Svc is that service (machine) and user (human) accounts look and function the same.

Every service you write needs to [register](/docs/openorch/register) at startup, or [log in](/docs/openorch/login) with the credentials it saves and manages if it's already registered. Just like a human.

A service account is not an admin account, it's a simple user level account. You might wonder how service-to-service calls work then.

## Service to service calls

Most endpoints on OpenOrch can only be called by administrators by default.

Let's take prompting. If you want to let your users prompt AIs you might write a wrapper service called `User Prompter Svc` with the slug `user-prompter-svc`.

If we look at the [Add Prompt endpoint API docs](/docs/openorch/prompt), we can see that it mentions

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

## Tokens

The User Svc produces a JWT ([JSON Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token)) upon [/user-svc/login](/docs/openorch/login) in the `token.token` field (see the response documentation).

You can either use this token as a proper JWT - decode it and inspect the contents, or you can just use the token to read the user account that belongs to the token with the [/user-svc/user/by-token](/docs/openorch/read-user-by-token) endpoint.

### Decoding a token

The [`/user-svc/public-key`](/docs/openorch/get-public-key) will return you the public key of the User Svc which then you can use that to decode the token.

Use the JWT libraries that are available in your programming language to do that, or use the Singularon [SDK](https://github.com/openorch/openorch/tree/main/sdk) if your language is supported.

### Token structure

The structure of the JWT is the following:

```js
{
   "sui":"usr_dC4K75Cbp6",
   "slu":"test-user-slug-0",
   "sri":[
      "user-svc:user",
      "user-svc:org:{org_dC4K7NNDCG}:user"
   ]
}
```

The field names are kept short to save space, so perhaps the Go definition is also educational:

```go
type Claims struct {
	UserId  string   `json:"sui"` // `sui`: openorch (previously singulatron) user ids
	Slug    string   `json:"slu"` // `slu`: openorch (previously singulatron) slug
	RoleIds []string `json:"sri"` // `sri`: openorch (previously singulatron) role ids
	jwt.RegisteredClaims
}
```

## Roles

Every user has a role, and a user token (see more about tokens on this page) produced upon login contains all the roles a user has.

```yaml
id: "user-svc:admin"
name: "User Svc - Admin Role"
```

```yaml
id: "your-svc:your-role"
name: "Your Svc - Your Role"
ownerId: "usr_eaSNcJ0BB0" # your user ID
```

In the below sections we'll refer to roles by their ID (such as `user-svc:admin`). Usually such readable strings are slugs, but in the case of roles slugs were eliminated for simplicity.

### Static roles

Static roles, such as

```sh
user-svc:admin
user-svc:user
```

defined by the `User Svc` are used for role-based access control. Each role has a list of permissions associated with it. When endpoints authorize a user it can do two things:

- Call the [Is Authorized](/docs/openorch/is-authorized) with the caller user auth headers and a permission ID to see if a given caller is authorized for that endpoint.
- Cache the list of permissions belonging to different roles and inspect only the caller's token to see if an appropriate role is present. This has the advantage of being quicker but slightly more complex (suitable SDK functions can help here). To parse and verify a token yourself, see the [Get Public Key](/docs/openorch/get-public-key) endpoint.

> If you are looking at restricting access to endpoints in other ways, you might be interested in: [Policy Svc](/docs/built-in-services/policy-svc).

#### Custom static roles

While deceptively simple, static roles can get you far, even without any permissions associated with them. A prime use case is product subscriptions.

Let's say you have a new product called "Funny Cats Newsletter" with two subscription tiers: Pro and Ultra.
You might create a service with the slug `funny-cats-newsletter-svc` for this product. You could have the following custom static roles created by your service (by calling the [Create Role](/docs/openorch/create-role) endpoint):

```yaml
funny-cats-newsletter-svc:pro
funny-cats-newsletter-svc:ultra
```

By checking the existence of these roles in a user's token you can successfully authorize product specific features.

### Dynamic roles

Dynamic roles are generated based on specific user-resource associations, offering more flexible permission management compared to static roles.

Dynamic roles look like

```sh
user-svc:org:{org_dBZRCej3fo}:admin
user-svc:org:{org_dBZRCej3fo}:user
```

The dynamic values must be surrounded by `{}` symbols. The above example is how organization roles are represented. For more about organizations see the Organizations section on this page.

These dynamic roles, like static roles are stored in the JWT tokens so it is advisable to keep them to a minimum. The organization example is an apt one here: think about how many GitHub or Google organizations you are part of. Likely even a few dozen are at the most extreme upper limit.

> JWT tokens (and the dynamic they contain) are sent with each request, so try to be efficient with dynamic roles.

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

Any logged in user can create an organization, provided the `Organization` slug is not taken yet. The creator becomes the first admin of the organization, acquiring the role of `user-svc:org:{$orgId}:admin` role.

#### Membership

Admins can assign other users member (`user-svc:org:{$orgId}:user`) or admin roles (`user-svc:org:{$orgId}:admin`) for the organizations they administer.

### Use cases

Organizations let users to freely associate with each other and create groups. Think about Discord servers, Slack workspaces, Github organizations etc.

## Permissions

### Conventions

Each permission created must by prefixed by the slug of the account that created it. Said account becomes the owner of the permission and only that account can add the permission to a role.

> Once you (your service) own a permission (by creating it, and it being prefixed by your account slug), you can add it to any role, not just roles owned by you.

Example; let's say your service is `petstore-svc`. OpenOrch prefers fine-grained access control, so you are free to define your own permissions, such as `petstore-svc:read` or `petstore-svc:pet:read`.

## Services with multiple nodes

You might now wonder what happens when a service has multiple instances/nodes. Won't their user accounts "clash" in the `User Svc`? The answer to this is that from the `User Svc` point of view, each node/instance of a service is the same account.

This is possible because the platform is designed with services having a "Shared Database Access".

Let's say you have a Cassandra network that spans multiple Availability Zones/Regions. Your nodes will also span multiple AZs/Regions and each instance of them will log in as `X Svc`.
