---
sidebar_position: 20
tags:
  - secret-svc
  - secrets
  - configuration
  - services
---

# Secret Svc

The Secret Svc stores sensitive or internal (non-end-user-facing) configuration. Aims to store all configuration not required at bootstrap, both for internal and external services.

> This page provides a high-level overview of `Secret Svc`. For detailed information, refer to the [Secret Svc API documentation](/docs/openorch/list-secrets).

## Access rules

### Read

Any logged in user who is amongst a `Secret`'s `Readers` can read a secret.

### Write

#### Create

Any logged in user can create a secret. Non-admin users can only create secrets with the key prefixed by their slug, ie:

```sh
deploy-svc/EXAMPLE-KEY
```

vs non-prefixed keys such as

```sh
EMAIL_API_KEY
```

Non-prefixed keys like `EMAIL_API_KEY` can only be created by admin users.

This prefix rule serves two purposes:

- It is clear which secret keys are "static" and originating from admin users
- It can prevent issues where a user claims a key knowing that it might be used later and overwritten/populated by an admin with sensitive information

#### Update

After a key is created further write access is governed by the `Writers` block.

## Entities

### Secret

```yaml
id: "secr_eG8IvKwB0A"
key: "MY_API_KEY"
value: "nNl4X9+@95Z"

# Slugs of services and users who can read the secret
readers:
  - "alice"
  - "bob"

# Slugs of services and users who can modify the secret
writers:
  - "alice"
  - "bob"

# Slugs of services and users who can delete the secret
deleters:
  - "service-admin"

# Slugs of services and users who can change the "readers" list
canChangeReaders:
  - "alice"

# Slugs of services and users who can change the "writers" list
canChangeWriters:
  - "alice"

# Slugs of services and users who can change the "deleters" list
canChangeDeleters:
  - "alice"
```

## Design choices

The Secret Svc, like most things in OpenOrch, is designed to be simple to reason about.

Instead of the OpenOrch injecting environment variables into service containers when they are deployed, the services are left to their own devices to read secrets from the Secret Svc through normal service calls, using their credentials.

This approach also works for services that you deploy manually (e.g., Kubernetes, Docker Compose) rather than through OpenOrch.

### Encryption at rest and transit

All data is encrypted using the encryption key provided by the envar `OPENORCH_ENCRYPTION_KEY` (see Todo section).

The server encrypts the secret values before saving them to disk/DB. The secret values are transmitted to readers unencrypted.

### Tips

#### Encrypt

The encrypt command helps you create encrypted YAML files that can be safely stored in source control and integrated into Infrastructure-as-Code (IaC) or GitOps workflows. This ensures sensitive data is protected while enabling automated deployment processes.

```sh
$ oo secret encrypt example-key example-value
```

```yaml
id: "secr_eR6LbYOBK2"
key: "example-key"
value: "62bQMQf5wPMrAsJ7+bcZpKBMtA7Ap7DF6xZaioq9jU0="
encrypted: true
checksum: "45a3b25f"
checksumAlgorithm: "CRC32"
```

Save the output to a file and, in your continuous delivery pipeline, apply it:

```sh
$ oo secret save my-api-key.yaml
```

##### Checksum

Checksums are optional and serve to verify the integrity of encrypted values. When an already encrypted value is saved in the Secret Svc, the service decodes it and uses the checksum to ensure the value remains intact.

#### Is Secure

After setting up your daemon it's a good idea to check if the Secret Svc is secure:

```sh
$ oo secret is-secure
Service is secure.
```

This will return successfully if the encryption key has been changed from the default value and all necessary setup steps have been completed.
