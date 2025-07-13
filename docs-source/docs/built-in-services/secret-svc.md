---
sidebar_position: 20
tags:
  - secret-svc
  - secrets
  - configuration
  - services
  - multitenant
  - encryption
  - security
---

# Secret Svc

The Secret Svc provides secure, encrypted storage for sensitive configuration data in 1Backend. It offers enterprise-grade secret management with multi-tenant access controls, encryption at rest, GitOps integration, and comprehensive audit capabilities.

> This page provides comprehensive usage examples and advanced features. For API details, see [Secret Svc API documentation](/docs/1backend-api/list-secrets).

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

The Secret Svc, like most things in 1Backend, is designed to be simple to reason about.

Instead of the 1Backend injecting environment variables into service containers when they are deployed, the services are left to their own devices to read secrets from the Secret Svc through normal service calls, using their credentials.

This approach also works for services that you deploy manually (e.g., Kubernetes, Docker Compose) rather than through 1Backend.

### Encryption at rest and transit

All data is encrypted using the encryption key provided by the envar `OB_ENCRYPTION_KEY` (see Todo section).

The server encrypts the secret values before saving them to disk/DB. The secret values are transmitted to readers unencrypted.

### Tips

#### Encrypt

The encrypt command helps you create encrypted YAML files that can be safely stored in source control and integrated into Infrastructure-as-Code (IaC) or GitOps workflows. This ensures sensitive data is protected while enabling automated deployment processes.

```sh
oo secret encrypt example-key example-value
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
oo secret save my-api-key.yaml
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

## CLI Reference

### Basic Secret Management

```bash
# Save a simple API key
oo secret save DATABASE_URL "postgresql://user:pass@host:5432/db"

# List all accessible secrets
oo secret list

# Show secret values (unmasked)
oo secret list --show

# Remove a secret by key
oo secret remove --key DATABASE_URL
```

### Secure Encryption Workflow

```bash
# Encrypt a secret for GitOps (interactive - hides value from terminal)
oo secret encrypt API_KEY
Enter secret value: [hidden]

# Direct encryption (for automation)
oo secret encrypt API_KEY "super-secret-value"

# Check if service is properly secured
oo secret is-secure
```

### Secret Management Commands

#### `oo secret save` - Store Secrets

**Basic Usage:**

```bash
# Save key-value pair
oo secret save <key> <value>

# Save from YAML file
oo secret save <file.yaml>

# Save multiple secrets from directory
oo secret save <directory>
```

**Examples:**

```bash
# API credentials
oo secret save STRIPE_SECRET_KEY "sk_test_abc123"
oo secret save OPENAI_API_KEY "sk-proj-xyz789"

# Database connection
oo secret save DB_PASSWORD "complex-password-123"

# Service-specific secrets (auto-prefixed for non-admin users)
oo secret save auth-svc/JWT_SECRET "jwt-signing-key"
oo secret save payment-svc/WEBHOOK_SECRET "webhook-validation-key"
```

#### `oo secret list` - View Secrets

**Usage:**

```bash
oo secret list [options]
oo secret list [key-pattern]
```

**Options:**

- `--show` - Display actual values (unmasked)
- `--namespace`, `-n` - Filter by namespace/app

**Examples:**

```bash
# List all secrets (values masked)
oo secret list

# Show actual secret values
oo secret list --show

# Filter by key pattern
oo secret list DATABASE
```

#### `oo secret remove` - Delete Secrets

**Usage:**

```bash
oo secret remove --key <key> [--key <key2>...]
oo secret remove --id <id> [--id <id2>...]
```

**Examples:**

```bash
# Remove by key
oo secret remove --key API_KEY

# Remove multiple by key
oo secret remove --key DB_USER --key DB_PASS

# Remove by ID
oo secret remove --id secr_abc123
```

#### `oo secret encrypt` - GitOps Encryption

**Usage:**

```bash
oo secret encrypt <key> [value]
```

**Examples:**

```bash
# Interactive encryption (secure - no terminal history)
oo secret encrypt PRODUCTION_API_KEY
Enter secret value: [hidden]

# Direct encryption for automation
oo secret encrypt STAGING_KEY "staging-value"

# Output ready for version control
oo secret encrypt DATABASE_PASSWORD > secrets/prod-db.yaml
```

#### `oo secret is-secure` - Security Check

```bash
# Verify service security configuration
oo secret is-secure
Service is secure.

# Example warning output
oo secret is-secure
Error: secret svc is not secure: it is using the default encryption key
```

## File-Based Secret Management

### Single Secret YAML Structure

```yaml
# production-api-key.yaml
id: "secr_prod_api_001"
key: "PRODUCTION_API_KEY"
value: "a37/KUAr4SOYi6Xw9i9T8qo3QCk8WvnzONo47jHAkwk="
encrypted: true
checksum: "45a3b25f"
checksumAlgorithm: "CRC32"

# Access control
readers:
  - "payment-svc"
  - "order-svc"
  - "admin-team"
writers:
  - "admin-team"
  - "devops-team"
deleters:
  - "admin-team"

# Permission management
canChangeReaders:
  - "admin-team"
canChangeWriters:
  - "admin-team"
canChangeDeleters:
  - "admin-team"
```

### Multiple Secrets YAML Structure

```yaml
# application-secrets.yaml
- id: "secr_db_primary"
  key: "DATABASE_URL"
  value: "encrypted-connection-string"
  encrypted: true
  readers: ["api-svc", "worker-svc"]
  writers: ["admin-team"]

- id: "secr_redis_cache"
  key: "REDIS_URL"  
  value: "redis://localhost:6379"
  encrypted: false
  readers: ["api-svc", "cache-svc"]
  writers: ["devops-team"]

- id: "secr_jwt_signing"
  key: "JWT_SECRET"
  value: "encrypted-jwt-key"
  encrypted: true
  readers: ["auth-svc"]
  writers: ["security-team"]
```

### Batch Operations

```bash
# Deploy all production secrets
oo secret save secrets/production/

# Update staging environment
oo secret save environments/staging-secrets.yaml
```

## Advanced Permission Management

### Granular Access Control

The Secret Svc implements fine-grained permission management with six distinct access levels:

#### Permission Types

- **`readers`** - Can view secret values
- **`writers`** - Can modify secret values  
- **`deleters`** - Can remove secrets
- **`canChangeReaders`** - Can modify the readers list
- **`canChangeWriters`** - Can modify the writers list
- **`canChangeDeleters`** - Can modify the deleters list

#### User vs "Admin" Access

**Regular Users:**

- Can only create secrets with keys prefixed by their user slug
- Automatically granted all permissions on their own secrets
- Must be explicitly granted access to others' secrets
- Need the

**Admin Users:**

- Can create secrets with any key name
- Have access to all secrets regardless of permission lists
- Can modify any secret's permission structure

### Permission Scenarios

#### Service-to-Service Communication

```yaml
# auth-service-secrets.yaml
- key: "auth-svc/DATABASE_URL"
  value: "encrypted-db-connection"
  encrypted: true
  readers: ["auth-svc"]           # Only auth service can read
  writers: ["devops-team"]        # Only devops can update
  deleters: ["admin-team"]        # Only admins can delete

- key: "auth-svc/JWT_SECRET"
  value: "encrypted-jwt-key" 
  encrypted: true
  readers: ["auth-svc", "api-gateway"]  # Both services need access
  writers: ["security-team"]
  deleters: ["security-team"]
```

#### Cross-Team Secret Sharing

```yaml
# shared-infrastructure-secrets.yaml
- key: "MONITORING_API_KEY"
  value: "encrypted-monitoring-key"
  encrypted: true
  readers: ["platform-team", "sre-team", "monitoring-svc"]
  writers: ["platform-team"]
  deleters: ["platform-team"]
  canChangeReaders: ["platform-team", "sre-team"]  # Both teams can add readers
  canChangeWriters: ["platform-team"]              # Only platform team controls writers
```

## Encryption and Security

### Supported Checksum Algorithms

The Secret Svc supports multiple checksum algorithms for data integrity verification:

- **CRC32** (default) - Fast, good for basic integrity checks
- **SHA-256** - Cryptographically secure, recommended for production
- **SHA-512** - Maximum security, slower performance
- **BLAKE2s** - Fast and secure alternative to SHA family

### Encryption Workflow for GitOps

#### 1. Encrypt Locally

```bash
# Create encrypted secret file
oo secret encrypt PRODUCTION_DATABASE_URL > secrets/prod-db.yaml

# Example output (safe for version control):
# id: "secr_ABC123"
# key: "PRODUCTION_DATABASE_URL"
# value: "a37/KUAr4SOYi6Xw9i9T8qo3QCk8WvnzONo47jHAkwk="
# encrypted: true
# checksum: "45a3b25f"
# checksumAlgorithm: "CRC32"
```

#### 2. Version Control Integration

```bash
# Commit encrypted secrets safely
git add secrets/
git commit -m "Add production database credentials"
git push origin main
```

#### 3. Deployment Pipeline

```bash
# In CI/CD pipeline
oo secret save secrets/production/
oo secret save secrets/staging/

# Verify deployment
oo secret is-secure
oo secret list --namespace production
```

## Real-World Usage Patterns

### Multi-Environment Management

#### Development Environment Setup

```bash
# Development secrets (unencrypted for local debugging)
oo secret save dev/DATABASE_URL "postgresql://localhost:5432/myapp_dev"
oo secret save dev/REDIS_URL "redis://localhost:6379"
oo secret save dev/API_KEY "dev-api-key-12345"

# List development secrets
oo secret list --namespace dev
```

### Microservices Architecture

#### Service-Specific Secret Management

```bash
# Authentication service secrets
oo secret save auth-svc/JWT_SECRET "$(openssl rand -base64 32)"
oo secret save auth-svc/OAUTH_CLIENT_SECRET "oauth-client-secret"
oo secret save auth-svc/PASSWORD_SALT "$(openssl rand -base64 16)"

# Payment service secrets  
oo secret save payment-svc/STRIPE_SECRET_KEY "sk_live_abc123"
oo secret save payment-svc/WEBHOOK_SECRET "whsec_xyz789"

# Email service secrets
oo secret save email-svc/SENDGRID_API_KEY "SG.abc123"
oo secret save email-svc/SMTP_PASSWORD "smtp-password"
```

#### Cross-Service Dependencies

```yaml
# shared-service-secrets.yaml
- key: "shared/DATABASE_URL"
  value: "encrypted-shared-db-connection"
  encrypted: true
  readers: 
    - "auth-svc"
    - "user-svc" 
    - "payment-svc"
    - "order-svc"
  writers: ["dba-team"]
  deleters: ["admin-team"]

- key: "shared/REDIS_URL"
  value: "encrypted-redis-connection"
  encrypted: true
  readers:
    - "cache-svc"
    - "session-svc"
    - "rate-limiter-svc"
  writers: ["infrastructure-team"]
```

## Security Best Practices

### Encryption Key Management

1. **Change Default Encryption Key**

   ```bash
   # Always verify encryption is properly configured
   oo secret is-secure
   
   # Set strong encryption key in environment
   export OB_ENCRYPTION_KEY="$(openssl rand -base64 32)"
   ```

### Access Control Best Practices

1. **Principle of Least Privilege**

   ```yaml
   # Give minimum necessary access
   - key: "payment-svc/STRIPE_KEY"
     readers: ["payment-svc"]           # Only payment service
     writers: ["payment-admin"]         # Only payment administrators
     deleters: ["security-team"]        # Only security team can delete
   ```
