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

## Quick Start

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

## CLI Reference

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

# List secrets for specific namespace
oo secret list --namespace production
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

# Backup current secrets to directory
mkdir backup-$(date +%Y%m%d)
oo secret list --show > backup-$(date +%Y%m%d)/secrets-backup.yaml
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

#### User vs Admin Access

**Regular Users:**
- Can only create secrets with keys prefixed by their user slug
- Automatically granted all permissions on their own secrets
- Must be explicitly granted access to others' secrets

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

#### Production Deployment

```bash
# Encrypt production secrets
oo secret encrypt prod/DATABASE_URL > secrets/prod-database.yaml
oo secret encrypt prod/API_KEY > secrets/prod-api.yaml

# Deploy with proper access controls
cat > secrets/prod-complete.yaml << EOF
- key: "prod/DATABASE_URL"
  value: "$(oo secret encrypt prod/DATABASE_URL | grep value:)"
  encrypted: true
  readers: ["api-svc", "worker-svc"]
  writers: ["admin-team"]
  deleters: ["admin-team"]

- key: "prod/MONITORING_KEY"
  value: "$(oo secret encrypt prod/MONITORING_KEY | grep value:)"
  encrypted: true
  readers: ["monitoring-svc"]
  writers: ["sre-team"]
  deleters: ["admin-team"]
EOF

oo secret save secrets/prod-complete.yaml
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

### Infrastructure as Code Integration

#### Terraform Integration

```bash
# Generate secrets for Terraform
oo secret list --show --namespace terraform > terraform-secrets.env

# Or use in Terraform data sources
data "external" "secrets" {
  program = ["bash", "-c", "oo secret list --show --namespace terraform | yq -o=json"]
}
```

#### Kubernetes Integration

```bash
# Export secrets for Kubernetes
oo secret list --show --namespace k8s-production | \
  yq -r '.[] | "kubectl create secret generic " + .key + " --from-literal=value=" + .value' | \
  bash
```

#### Docker Compose Integration

```bash
# Generate .env file for Docker Compose
oo secret list --show --namespace docker | \
  yq -r '.[] | .key + "=" + .value' > .env

# Use in docker-compose.yml
docker-compose --env-file .env up -d
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

2. **Regular Security Audits**
   ```bash
   # Audit secret access patterns
   oo secret list | awk '{print $3}' | sort | uniq -c
   
   # Check for secrets without proper access controls
   oo secret list --show | grep -E "(readers|writers|deleters): \[\]"
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

2. **Regular Access Review**
   ```bash
   # Review all secrets and their permissions
   oo secret list --show | grep -A 10 -B 5 "readers\|writers\|deleters"
   
   # Find secrets with overly broad access
   oo secret list --show | grep -E "readers.*\[.*,.*,.*\]"
   ```

### Rotation and Lifecycle Management

#### Automated Secret Rotation

```bash
#!/bin/bash
# rotate-database-password.sh

# Generate new password
NEW_PASSWORD=$(openssl rand -base64 32)

# Encrypt and save
oo secret encrypt DATABASE_PASSWORD "$NEW_PASSWORD" > temp-secret.yaml
oo secret save temp-secret.yaml

# Update database
mysql -h $DB_HOST -u root -p"$OLD_PASSWORD" -e "SET PASSWORD FOR 'app'@'%' = PASSWORD('$NEW_PASSWORD')"

# Cleanup
rm temp-secret.yaml
echo "Database password rotated successfully"
```

#### Secret Expiration Tracking

```bash
# Add expiration metadata to secrets
cat > expiring-secret.yaml << EOF
key: "TEMPORARY_API_KEY"
value: "$(oo secret encrypt TEMP_KEY temp-value | grep value:)"
encrypted: true
metadata:
  expires: "2024-12-31"
  rotation_interval: "90d"
  owner: "security-team"
EOF

oo secret save expiring-secret.yaml
```

## Troubleshooting

### Common Issues

#### Permission Denied Errors

```bash
# Check your current user permissions
oo user whoami

# Verify secret access permissions
oo secret list YOUR_SECRET_KEY

# Request access from secret owner
echo "Need access to secret: YOUR_SECRET_KEY" | \
  mail -s "Secret Access Request" security-team@company.com
```

#### Encryption/Decryption Failures

```bash
# Verify service security status
oo secret is-secure

# Check encryption key configuration
echo $OB_ENCRYPTION_KEY | wc -c  # Should be 32 characters

# Test encryption/decryption cycle
oo secret encrypt test-key test-value > test.yaml
oo secret save test.yaml
oo secret list test-key --show
oo secret remove --key test-key
```

#### Import/Export Issues

```bash
# Validate YAML syntax before import
yamllint secrets/production.yaml

# Test import with single secret first
head -n 10 secrets/production.yaml > test-import.yaml
oo secret save test-import.yaml

# Backup before bulk operations
oo secret list --show > backup-$(date +%Y%m%d).yaml
```

### Performance Optimization

#### Batch Operations

```bash
# Instead of individual saves:
# oo secret save KEY1 value1
# oo secret save KEY2 value2
# oo secret save KEY3 value3

# Use batch file:
cat > batch-secrets.yaml << EOF
- key: "KEY1"
  value: "value1"
- key: "KEY2"  
  value: "value2"
- key: "KEY3"
  value: "value3"
EOF

oo secret save batch-secrets.yaml
```

#### Large Secret Management

```bash
# For environments with 100+ secrets, use directory organization
mkdir -p secrets/{production,staging,development}
mkdir -p secrets/services/{auth,payment,email}

# Batch save by category
oo secret save secrets/production/
oo secret save secrets/services/auth/
```

## Integration Examples

### CI/CD Pipeline Integration

#### GitHub Actions

```yaml
# .github/workflows/deploy-secrets.yml
name: Deploy Secrets
on:
  push:
    paths: ['secrets/**']
    branches: [main]

jobs:
  deploy-secrets:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Setup 1Backend CLI
        run: |
          curl -L https://releases.1backend.com/oo-linux -o oo
          chmod +x oo
          
      - name: Deploy Production Secrets
        env:
          OO_TOKEN: ${{ secrets.ONEBACKEND_TOKEN }}
          OO_URL: ${{ secrets.ONEBACKEND_URL }}
        run: |
          ./oo secret save secrets/production/
          ./oo secret is-secure
```

#### GitLab CI

```yaml
# .gitlab-ci.yml
deploy-secrets:
  stage: deploy
  script:
    - oo secret save secrets/$CI_ENVIRONMENT_NAME/
    - oo secret is-secure
  only:
    changes:
      - secrets/**
  environment:
    name: $CI_ENVIRONMENT_NAME
```

### Application Integration

#### Go Application

```go
// Load secrets in Go application
package main

import (
    "context"
    "log"
    "github.com/1backend/1backend/clients/go"
    "github.com/1backend/1backend/sdk/go/client"
)

func loadSecrets() map[string]string {
    cf := client.NewApiClientFactory(os.Getenv("ONEBACKEND_URL"))
    client := cf.Client(client.WithToken(os.Getenv("ONEBACKEND_TOKEN")))
    
    secrets := []string{
        "DATABASE_URL",
        "REDIS_URL", 
        "API_KEY",
    }
    
    req := openapi.SecretSvcListSecretsRequest{
        Keys: secrets,
    }
    
    resp, _, err := client.SecretSvcAPI.ListSecrets(context.Background()).
        Body(req).Execute()
    if err != nil {
        log.Fatal("Failed to load secrets:", err)
    }
    
    secretMap := make(map[string]string)
    for _, secret := range resp.Secrets {
        secretMap[*secret.Key] = *secret.Value
    }
    
    return secretMap
}
```

#### Node.js Application

```javascript
// Load secrets in Node.js application
const { SecretSvcApi, Configuration } = require('@1backend/client');

async function loadSecrets() {
    const config = new Configuration({
        basePath: process.env.ONEBACKEND_URL,
        accessToken: process.env.ONEBACKEND_TOKEN
    });
    
    const secretApi = new SecretSvcApi(config);
    
    const response = await secretApi.listSecrets({
        keys: ['DATABASE_URL', 'REDIS_URL', 'API_KEY']
    });
    
    const secrets = {};
    response.secrets.forEach(secret => {
        secrets[secret.key] = secret.value;
    });
    
    return secrets;
}

// Usage
loadSecrets().then(secrets => {
    process.env.DATABASE_URL = secrets.DATABASE_URL;
    process.env.REDIS_URL = secrets.REDIS_URL;
    process.env.API_KEY = secrets.API_KEY;
    
    // Start application with loaded secrets
    require('./app');
});
```

## API Reference Summary

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/secret-svc/secrets` | POST | List secrets by key(s) |
| `/secret-svc/secrets` | PUT | Save/update secrets |
| `/secret-svc/secrets` | DELETE | Remove secrets |
| `/secret-svc/encrypt` | POST | Encrypt values for GitOps |
| `/secret-svc/decrypt` | POST | Decrypt values |
| `/secret-svc/is-secure` | GET | Check security status |

## Related Documentation

- [Secret Svc API Reference](/docs/1backend-api/list-secrets) - Complete API documentation
- [Config Svc](/docs/built-in-services/config-svc) - For non-sensitive configuration
- [User Svc](/docs/built-in-services/user-svc) - User and role management
- [Container Svc](/docs/built-in-services/container-svc) - Service deployment and secrets injection
