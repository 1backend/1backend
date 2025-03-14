---
sidebar_position: 70
tags:
  - policy-svc
  - blocklist
  - rate-limiting
  - policy
  - policies
  - reliability
---

# Policy Svc

The policy service provides features such as rate limiting of endpoint calls by user ip, user id, organization id and more.

> This page provides a high-level overview of `Policy Svc`. For detailed information, refer to the [Policy Svc API documentation](/docs/openorch/upsert-instance).

## Responsibilities

- Provide tools for service writers to prevent adversarial users from causing service degradation

## How It Works

The Policy Svc requires you to explicitly call the [/policy-svc/check](/docs/openorch/check) endpoint in every service endpoint you are building. There is no magic or framework feature involved.

## Usage

The policy service has two endpoint:

- You can create policy instances with [/policy-svc/upsert-instance](/docs/openorch/upsert-instance) - use this to define rate limits, block IPs etc.
- A [/policy-svc/check](/docs/openorch/check) endpoint that you should call for every request in your endpoint you want to rate limit.

While the documentation should be
thorough, it might be not be the easiest to understand at first glance due to the presence of `*Parameters` fields which are specific to `Policy Templates`.

## Terms

### Policy Templates

`Policy Templates` are hardcoded features of the `Policy Svc`:

#### Rate Limit

Rate Limit (templateId: `rate-limit`) provides rate limiting various entities and scopes, see the `rateLimitParameters` in the [api doc](/docs/openorch/upsert-instance)).

#### Blocklist

Blocklist provides blocking of access by `ip` addresses, see the `blocklistParameters` in the [api doc](/docs/openorch/upsert-instance)).

### Policy Instance

A policy instance is a specific application of a policy template to certain data like endpoints, user ids, ip addresses etc.

## Examples

### Rate Limit

The following [`/policy-svc/upsert-instance`](/docs/openorch/upsert-instance) payload rate limits calls to the `register` endpoint by caller IP: maximum 5 calls are permitted per IP per day:

```js
{
  "instance": {
    "endpoint": "/user-svc/register",
    "id": "insta_dBZRCej3fo",
    "rateLimitParameters": {
      "entity": "ip",
      "maxRequests": 5,
      "scope": "endpoint",
      "timeWindow": "1d"
    },
    "templateId": "rate-limit"
  }
}
```

### Block by IP

The following [`/policy-svc/upsert-instance`](/docs/openorch/upsert-instance) payload blocks access to the register endpoint by ip address.

```js
{
  "instance": {
    "endpoint": "/user-svc/register",
    "id": "insta_dBZRCej3fo",
    "blocklistParameters": {
      "blockedIPs": ["8.8.8.8"]
    },
    "templateId": "rate-limit"
  }
}
```
