---
sidebar_position: 80
tags:
  - config-svc
  - configuration
  - services
---

# Config Svc

The Config Svc stores public, non-sensitive and end-user-facing data.

> This page provides a high-level overview of `Config Svc`. For detailed information, refer to the [Secret Svc API documentation](/docs/openorch/get-config).

The Config Svc is less critical than it might seem—most configuration happens internally through the [Secret Svc](/docs/built-in-services/secret-svc).

At the moment, it functions more like a minimal feature-flag service.

## Access Rules

### Read

All configs are publicly readable even without a user account.

### Write

Any logged in user can write to the config, but only to the key that matches their slug, ie. if a user's slug is `jane-doe`:

```
{
  "jane-doe": {"janesConfig": 5},
  "someOtherKey": "hi"
}
```

Only the key `jane-doe` will be written to the `Config`, all other keys (such as `someOtherKey`) will be ignored.

## Related

- [Secret Svc](/docs/built-in-services/secret-svc) to store sensitive data like internal configuration and secrets
