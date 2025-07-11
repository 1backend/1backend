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

The User Svc is the authentication and authorization backbone of 1Backend, providing comprehensive identity management, role-based access control, multi-tenant isolation, and service-to-service authentication. It manages users, tokens, organizations, permissions, and provides the foundation for secure, scalable applications.

> This page provides comprehensive usage examples and advanced patterns. For API details, see [User Svc API documentation](/docs/1backend-api/login).

## Quick Start

### Basic Authentication

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

### Service Account Setup

```bash
# Register a service account
oo register payment-svc
Enter password: [secure-service-password]

# Login as service for automation
oo login payment-svc secure-service-password

# Verify service authentication
oo whoami
```

## CLI Reference

### Authentication Commands

#### `oo register` - Create New Accounts

**Usage:**
```bash
oo register [slug] [password]
oo register [--contact-id email] [--contact-platform email]
```

**Examples:**
```bash
# Interactive registration (secure - no terminal history)
oo register alice
Enter password: [hidden]

# Service account registration
oo register payment-service
oo register order-processor
oo register notification-service

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

# Service authentication in CI/CD
oo login payment-svc $SERVICE_PASSWORD

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

**Examples:**
```bash
# List all enrollments
oo enroll list

# Filter by role
oo enroll list --role user-svc:admin

# Filter by user
oo enroll list --userId usr_abc123

# Filter by contact
oo enroll list --contactId alice@company.com
```

## Service-to-Service Authentication

### Service Registration & Authentication

#### 1. Service Account Setup

```bash
# Register service account
oo register payment-processor
Enter password: [secure-service-password]

# Login as service
oo login payment-processor
Enter password: [secure-service-password]

# Verify service identity
oo whoami
# Output:
# id: usr_svc_001
# slug: payment-processor
# roles:
#   - user-svc:user
```

#### 2. Permission Configuration

```yaml
# payment-service-permissions.yaml
- id: "payment-database-access"
  permissionId: "data-svc:read"
  slugs:
    - "payment-processor"
    
- id: "payment-write-access"
  permissionId: "data-svc:write"
  slugs:
    - "payment-processor"
    
- id: "notification-send"
  permissionId: "email-svc:send"
  slugs:
    - "payment-processor"
    
- id: "secret-access"
  permissionId: "secret-svc:secret:list"
  slugs:
    - "payment-processor"
```

```bash
# Apply service permissions
oo permit save payment-service-permissions.yaml
```

#### 3. Service Integration Code

**Go Service Implementation:**
```go
// payment-service/main.go
package main

import (
    "context"
    "log"
    "os"
    
    "github.com/1backend/1backend/clients/go"
    "github.com/1backend/1backend/sdk/go/client"
    "github.com/1backend/1backend/sdk/go/boot"
)

func main() {
    // Service authentication
    clientFactory := client.NewApiClientFactory(os.Getenv("ONEBACKEND_URL"))
    
    // Register or login service account
    token, err := boot.RegisterServiceAccount(
        clientFactory.Client().UserSvcAPI,
        "payment-processor",
        "Payment Processing Service",
        credentialStore, // Your credential storage
    )
    if err != nil {
        log.Fatal("Failed to authenticate service:", err)
    }
    
    // Use authenticated client for service calls
    authenticatedClient := clientFactory.Client(client.WithToken(token.Token))
    
    // Example: Access secrets
    secrets, _, err := authenticatedClient.SecretSvcAPI.ListSecrets(context.Background()).
        Body(openapi.SecretSvcListSecretsRequest{
            Keys: []string{"STRIPE_SECRET_KEY", "DATABASE_URL"},
        }).Execute()
    if err != nil {
        log.Fatal("Failed to load secrets:", err)
    }
    
    // Example: Check permissions
    hasPermission, _, err := authenticatedClient.UserSvcAPI.HasPermission(context.Background(), "data-svc:write").Execute()
    if err != nil || !hasPermission.Authorized {
        log.Fatal("Service lacks required permissions")
    }
    
    // Start service with authenticated context
    startPaymentProcessor(authenticatedClient)
}
```

**Node.js Service Implementation:**
```javascript
// notification-service/index.js
const { UserSvcApi, SecretSvcApi, Configuration } = require('@1backend/client');

async function authenticateService() {
    const config = new Configuration({
        basePath: process.env.ONEBACKEND_URL,
    });
    
    const userApi = new UserSvcApi(config);
    
    // Service login
    const loginResponse = await userApi.login({
        slug: 'notification-service',
        password: process.env.SERVICE_PASSWORD,
        device: 'server'
    });
    
    // Update config with token
    config.accessToken = loginResponse.token.token;
    
    return config;
}

async function loadServiceSecrets(config) {
    const secretApi = new SecretSvcApi(config);
    
    const secrets = await secretApi.listSecrets({
        keys: ['SENDGRID_API_KEY', 'NOTIFICATION_TEMPLATES']
    });
    
    return secrets.secrets.reduce((acc, secret) => {
        acc[secret.key] = secret.value;
        return acc;
    }, {});
}

async function main() {
    try {
        const config = await authenticateService();
        const secrets = await loadServiceSecrets(config);
        
        // Start notification service with authenticated context
        startNotificationService(config, secrets);
    } catch (error) {
        console.error('Service authentication failed:', error);
        process.exit(1);
    }
}

main();
```

### CI/CD Service Setup

#### GitHub Actions Service Authentication

```yaml
# .github/workflows/deploy-payment-service.yml
name: Deploy Payment Service
on:
  push:
    branches: [main]
    paths: ['payment-service/**']

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Setup 1Backend CLI
        run: |
          curl -L https://releases.1backend.com/oo-linux -o oo
          chmod +x oo
          
      - name: Configure Environment
        run: |
          ./oo env add production ${{ secrets.ONEBACKEND_URL }}
          ./oo env use production
          
      - name: Service Authentication
        run: |
          ./oo login payment-processor ${{ secrets.SERVICE_PASSWORD }}
          
      - name: Deploy Service Permissions
        run: |
          ./oo permit save deploy/permissions/
          ./oo enroll save deploy/enrollments/
          
      - name: Verify Service Setup
        run: |
          ./oo whoami
          ./oo permit list
```

#### Docker Service Startup

```dockerfile
# payment-service/Dockerfile
FROM node:18-alpine

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .

# Service authentication in container
ENV ONEBACKEND_URL=${ONEBACKEND_URL}
ENV SERVICE_PASSWORD=${SERVICE_PASSWORD}

CMD ["node", "src/index.js"]
```

```bash
# Docker Compose with service authentication
version: '3.8'
services:
  payment-service:
    build: ./payment-service
    environment:
      - ONEBACKEND_URL=${ONEBACKEND_URL}
      - SERVICE_PASSWORD=${PAYMENT_SERVICE_PASSWORD}
    restart: unless-stopped
    
  notification-service:
    build: ./notification-service
    environment:
      - ONEBACKEND_URL=${ONEBACKEND_URL}
      - SERVICE_PASSWORD=${NOTIFICATION_SERVICE_PASSWORD}
    restart: unless-stopped
```

## Organization Management

### Creating Organizations

```bash
# Create organization (via API or SDK)
# Organizations are typically created programmatically
```

**Create Organization via SDK:**
```go
// Create organization
org, _, err := client.UserSvcAPI.SaveOrganization(context.Background()).
    Body(openapi.UserSvcSaveOrganizationRequest{
        Organization: &openapi.UserSvcOrganizationInput{
            Name: openapi.PtrString("Acme Corporation"),
            Slug: openapi.PtrString("acme-corp"),
        },
    }).Execute()
```

### Organization Role Management

#### Administrative Roles

```yaml
# organization-admin-enrollments.yaml
- id: "acme-corp-ceo"
  role: "user-svc:org:org_acme123:admin"
  contactId: "ceo@acme.com"
  
- id: "acme-corp-cto"
  role: "user-svc:org:org_acme123:admin"
  contactId: "cto@acme.com"
```

#### Member Roles

```yaml
# organization-member-enrollments.yaml
- id: "acme-developer-1"
  role: "user-svc:org:org_acme123:user"
  contactId: "developer1@acme.com"
  
- id: "acme-developer-2"
  role: "user-svc:org:org_acme123:user"
  contactId: "developer2@acme.com"
  
- id: "acme-designer"
  role: "user-svc:org:org_acme123:user"
  contactId: "designer@acme.com"
```

#### Apply Organization Enrollments

```bash
# Setup organization structure
oo enroll save organization-admin-enrollments.yaml
oo enroll save organization-member-enrollments.yaml

# Verify organization setup
oo enroll list --role user-svc:org:org_acme123:admin
oo enroll list --role user-svc:org:org_acme123:user
```

### Organization-Scoped Permissions

```yaml
# organization-scoped-permissions.yaml
- id: "acme-project-management"
  permissionId: "project-svc:manage"
  roles:
    - "user-svc:org:org_acme123:admin"
    
- id: "acme-project-view"
  permissionId: "project-svc:view"
  roles:
    - "user-svc:org:org_acme123:user"
    - "user-svc:org:org_acme123:admin"
    
- id: "acme-billing-access"
  permissionId: "billing-svc:access"
  roles:
    - "user-svc:org:org_acme123:admin"
```

## Advanced Permission Patterns

### Hierarchical Role System

#### Service-Specific Roles

```yaml
# payment-service-roles.yaml
- id: "payment-admin-permissions"
  permissionId: "payment-svc:admin"
  roles:
    - "payment-svc:admin"
    
- id: "payment-processor-permissions"
  permissionId: "payment-svc:process"
  roles:
    - "payment-svc:processor"
    - "payment-svc:admin"  # Admins can also process
    
- id: "payment-viewer-permissions"
  permissionId: "payment-svc:view"
  roles:
    - "payment-svc:viewer"
    - "payment-svc:processor"  # Processors can view
    - "payment-svc:admin"      # Admins can view
```

#### Cross-Service Access

```yaml
# cross-service-permissions.yaml
- id: "api-gateway-user-access"
  permissionId: "user-svc:user:view"
  slugs:
    - "api-gateway"
    
- id: "frontend-chat-access"
  permissionId: "chat-svc:message:create"
  slugs:
    - "frontend-app"
    - "mobile-app"
    
- id: "admin-panel-access"
  permissionId: "admin-svc:access"
  roles:
    - "user-svc:admin"
    - "user-svc:org:*:admin"  # Any org admin
```

### Dynamic Role Assignment

#### Role Enrollment for New Users

```yaml
# new-user-auto-enrollments.yaml
- id: "default-user-role"
  role: "user-svc:user"
  contactId: "*"  # Auto-enroll all new users
  
- id: "developer-team-auto"
  role: "dev-team:developer"
  contactId: "*@company.com"  # Auto-enroll company emails
  
- id: "premium-user-enrollment"
  role: "app:premium"
  contactId: "premium@company.com"
```

#### Conditional Role Assignment

```bash
# Enroll users based on conditions
oo enroll save dev-team:lead --contactId lead-dev@company.com
oo enroll save payment-svc:processor --userId usr_experienced_dev
oo enroll save user-svc:org:org_startup:admin --contactId founder@startup.com
```

## Multi-Tenant Applications

### App-Based Isolation

#### Environment Separation

```yaml
# production-app-setup.yaml
app: "production.mycompany.com"
permits:
  - id: "prod-api-access"
    permissionId: "api-svc:access"
    roles: ["api-svc:user"]
    
enrollments:
  - id: "prod-admin"
    role: "user-svc:admin"
    contactId: "admin@mycompany.com"
    app: "production.mycompany.com"
```

```yaml
# staging-app-setup.yaml
app: "staging.mycompany.com"
permits:
  - id: "staging-api-access"
    permissionId: "api-svc:access"
    roles: ["api-svc:user"]
    
enrollments:
  - id: "staging-admin"
    role: "user-svc:admin"
    contactId: "dev@mycompany.com"
    app: "staging.mycompany.com"
```

#### Client Separation

```yaml
# client-a-setup.yaml
app: "client-a.com"
permits:
  - id: "client-a-data-access"
    permissionId: "data-svc:read"
    roles: ["client-a:user"]
    
enrollments:
  - id: "client-a-admin"
    role: "client-a:admin"
    contactId: "admin@client-a.com"
    app: "client-a.com"
```

```yaml
# client-b-setup.yaml
app: "client-b.com"
permits:
  - id: "client-b-data-access"
    permissionId: "data-svc:read"
    roles: ["client-b:user"]
    
enrollments:
  - id: "client-b-admin"
    role: "client-b:admin"
    contactId: "admin@client-b.com"
    app: "client-b.com"
```

### App Context Management

```bash
# Configure different app environments
oo env add client-a https://client-a.1backend.com
oo env add client-b https://client-b.1backend.com
oo env add production https://prod.1backend.com

# Switch between clients
oo env use client-a
oo login admin@client-a.com

oo env use client-b
oo login admin@client-b.com

# Verify app isolation
oo whoami  # Shows different roles per app
```

## Security Patterns

### Token Management

#### Token Verification

```go
// Verify JWT token in service
func verifyToken(tokenString string) (*auth.Claims, error) {
    // Get public key from User Svc
    publicKey, _, err := userClient.GetPublicKey(context.Background()).Execute()
    if err != nil {
        return nil, err
    }
    
    // Parse and verify token
    authorizer := auth.AuthorizerImpl{}
    claims, err := authorizer.ParseJWT(publicKey.PublicKey, tokenString)
    if err != nil {
        return nil, err
    }
    
    return claims, nil
}
```

#### Token Refresh Handling

```javascript
// Handle token refresh in client
async function makeAuthenticatedRequest(url, options = {}) {
    let token = localStorage.getItem('auth_token');
    
    options.headers = {
        ...options.headers,
        'Authorization': `Bearer ${token}`
    };
    
    let response = await fetch(url, options);
    
    // Handle token expiration
    if (response.status === 401) {
        // Refresh token
        const refreshResponse = await fetch('/user-svc/refresh-token', {
            method: 'POST',
            headers: { 'Authorization': `Bearer ${token}` }
        });
        
        if (refreshResponse.ok) {
            const { token: newToken } = await refreshResponse.json();
            localStorage.setItem('auth_token', newToken);
            
            // Retry original request
            options.headers['Authorization'] = `Bearer ${newToken}`;
            response = await fetch(url, options);
        }
    }
    
    return response;
}
```

### Permission Checking

#### Service Permission Validation

```go
// Check permissions in service handler
func requirePermission(permission string) middleware.Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := extractTokenFromHeader(r)
            if token == "" {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            
            // Check permission with User Svc
            hasPermResponse, _, err := userClient.HasPermission(
                r.Context(), 
                permission,
            ).Execute()
            
            if err != nil || !hasPermResponse.Authorized {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            // Add user context to request
            ctx := context.WithValue(r.Context(), "user", hasPermResponse.User)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// Usage in handlers
func setupRoutes() {
    http.Handle("/api/payments", 
        requirePermission("payment-svc:process")(
            http.HandlerFunc(handlePayments)
        )
    )
    
    http.Handle("/api/admin", 
        requirePermission("payment-svc:admin")(
            http.HandlerFunc(handleAdmin)
        )
    )
}
```

#### Frontend Permission Checks

```javascript
// Permission-based UI rendering
class PermissionGuard extends React.Component {
    constructor(props) {
        super(props);
        this.state = { hasPermission: false, loading: true };
    }
    
    async componentDidMount() {
        try {
            const response = await makeAuthenticatedRequest(
                `/user-svc/self/has/${this.props.permission}`,
                { method: 'POST' }
            );
            const result = await response.json();
            this.setState({ 
                hasPermission: result.authorized,
                loading: false 
            });
        } catch (error) {
            this.setState({ hasPermission: false, loading: false });
        }
    }
    
    render() {
        if (this.state.loading) {
            return <div>Loading...</div>;
        }
        
        if (!this.state.hasPermission) {
            return this.props.fallback || null;
        }
        
        return this.props.children;
    }
}

// Usage
function AdminPanel() {
    return (
        <PermissionGuard 
            permission="admin-svc:access"
            fallback={<div>Access Denied</div>}
        >
            <AdminDashboard />
        </PermissionGuard>
    );
}
```

## Production Deployment Patterns

### Service Bootstrap

#### Service Registration Script

```bash
#!/bin/bash
# scripts/bootstrap-services.sh

# Services to register
SERVICES=(
    "api-gateway"
    "user-service"
    "payment-service"
    "notification-service"
    "data-service"
)

echo "Bootstrapping 1Backend services..."

for service in "${SERVICES[@]}"; do
    echo "Setting up $service..."
    
    # Generate secure password
    PASSWORD=$(openssl rand -base64 32)
    
    # Register service
    oo register "$service" "$PASSWORD"
    
    # Store credentials securely
    oo secret save "service-credentials/$service" "$PASSWORD"
    
    echo "$service registered successfully"
done

echo "Applying service permissions..."
oo permit save deploy/permissions/
oo enroll save deploy/enrollments/

echo "Bootstrap complete!"
```

#### Kubernetes Service Setup

```yaml
# k8s/service-auth-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: service-auth-config
data:
  onebackend-url: "https://api.company.com"
  service-slug: "payment-processor"

---
# k8s/service-auth-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: service-auth-secret
type: Opaque
stringData:
  service-password: "secure-service-password"

---
# k8s/payment-service-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: payment-service
spec:
  replicas: 3
  selector:
    matchLabels:
      app: payment-service
  template:
    metadata:
      labels:
        app: payment-service
    spec:
      containers:
      - name: payment-service
        image: company/payment-service:latest
        env:
        - name: ONEBACKEND_URL
          valueFrom:
            configMapKeyRef:
              name: service-auth-config
              key: onebackend-url
        - name: SERVICE_SLUG
          valueFrom:
            configMapKeyRef:
              name: service-auth-config
              key: service-slug
        - name: SERVICE_PASSWORD
          valueFrom:
            secretKeyRef:
              name: service-auth-secret
              key: service-password
```

### Load Balancing & High Availability

#### Service Node Registration

```go
// Shared service account across multiple nodes
func main() {
    // All nodes use same service account
    token, err := boot.RegisterServiceAccount(
        clientFactory.Client().UserSvcAPI,
        "payment-processor",  // Same slug for all nodes
        "Payment Processing Service",
        credentialStore,
    )
    if err != nil {
        log.Fatal("Failed to authenticate service:", err)
    }
    
    // Register node-specific health check
    nodeID := os.Getenv("NODE_ID")
    registerHealthCheck(nodeID, token.Token)
    
    // Start processing with shared authentication
    startProcessor(token.Token)
}
```

#### Database Access Pattern

```go
// Shared database access across service nodes
func connectToDatabase(serviceToken string) *sql.DB {
    // Load database credentials using service authentication
    secrets, _, err := secretClient.ListSecrets(context.Background()).
        Body(openapi.SecretSvcListSecretsRequest{
            Keys: []string{
                "DATABASE_URL",
                "DATABASE_PASSWORD",
                "REDIS_URL",
            },
        }).Execute()
    if err != nil {
        log.Fatal("Failed to load database credentials:", err)
    }
    
    // All service nodes connect with same credentials
    dbURL := findSecret(secrets.Secrets, "DATABASE_URL")
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }
    
    return db
}
```

## Troubleshooting

### Common Authentication Issues

#### Invalid Token Errors

```bash
# Check token validity
oo token
oo whoami

# Refresh expired token
oo login alice
Enter password: [hidden]

# Verify token format
echo $(oo token) | base64 -d | jq .

# Test token with API
curl -H "Authorization: Bearer $(oo token)" \
  https://api.1backend.com/user-svc/self
```

#### Permission Denied Issues

```bash
# Check user permissions
oo whoami

# List user's permits
oo permit list

# Check specific permission
curl -H "Authorization: Bearer $(oo token)" \
  -X POST https://api.1backend.com/user-svc/self/has/target-permission

# Debug permission hierarchy
oo enroll list --userId $(oo whoami | grep id: | cut -d' ' -f2)
```

#### Service Authentication Failures

```bash
# Verify service registration
oo user list --contactId service-name

# Check service permissions
oo login service-name
oo permit list

# Test service-to-service call
curl -H "Authorization: Bearer $(oo token)" \
  https://api.1backend.com/target-service/endpoint
```

### Role and Organization Issues

#### Role Assignment Problems

```bash
# Check role enrollment
oo enroll list --role target-role

# Verify role ownership
oo whoami | grep roles:

# Check organization membership
oo enroll list --role "user-svc:org:*"

# Fix role assignment
oo enroll save correct-role --userId target-user-id
```

#### Organization Access Issues

```bash
# List organization enrollments
oo enroll list --role "user-svc:org:org_id:admin"
oo enroll list --role "user-svc:org:org_id:user"

# Check organization-scoped permissions
oo permit list | grep "org:org_id"

# Debug organization context
oo env current
oo whoami --all
```

### Multi-Tenant Issues

#### App Context Problems

```bash
# Check current app context
oo env current

# Switch app environment
oo env use correct-environment

# Verify app-specific permissions
oo whoami
oo permit list

# Test cross-app isolation
oo env use app-a
oo whoami
oo env use app-b
oo whoami  # Should show different permissions
```

## Integration Examples

### Frontend Authentication

#### React Authentication Hook

```javascript
// hooks/useAuth.js
import { useState, useEffect, createContext, useContext } from 'react';

const AuthContext = createContext();

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    
    useEffect(() => {
        loadUser();
    }, []);
    
    async function loadUser() {
        const token = localStorage.getItem('auth_token');
        if (!token) {
            setLoading(false);
            return;
        }
        
        try {
            const response = await fetch('/user-svc/self', {
                headers: { 'Authorization': `Bearer ${token}` }
            });
            
            if (response.ok) {
                const userData = await response.json();
                setUser(userData);
            } else {
                localStorage.removeItem('auth_token');
            }
        } catch (error) {
            console.error('Failed to load user:', error);
            localStorage.removeItem('auth_token');
        } finally {
            setLoading(false);
        }
    }
    
    async function login(slug, password) {
        const response = await fetch('/user-svc/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ slug, password, device: 'web' })
        });
        
        if (response.ok) {
            const { token } = await response.json();
            localStorage.setItem('auth_token', token.token);
            await loadUser();
            return true;
        }
        
        return false;
    }
    
    function logout() {
        localStorage.removeItem('auth_token');
        setUser(null);
    }
    
    return (
        <AuthContext.Provider value={{
            user,
            loading,
            login,
            logout,
            isAuthenticated: !!user
        }}>
            {children}
        </AuthContext.Provider>
    );
}

export const useAuth = () => useContext(AuthContext);
```

#### Vue.js Authentication Store

```javascript
// stores/auth.js
import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null,
        token: localStorage.getItem('auth_token'),
        loading: false
    }),
    
    getters: {
        isAuthenticated: (state) => !!state.user,
        hasRole: (state) => (role) => {
            return state.user?.roles?.includes(role) || false;
        },
        hasPermission: (state) => async (permission) => {
            if (!state.token) return false;
            
            try {
                const response = await fetch(`/user-svc/self/has/${permission}`, {
                    method: 'POST',
                    headers: { 'Authorization': `Bearer ${state.token}` }
                });
                const result = await response.json();
                return result.authorized;
            } catch {
                return false;
            }
        }
    },
    
    actions: {
        async login(slug, password) {
            this.loading = true;
            
            try {
                const response = await fetch('/user-svc/login', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ slug, password, device: 'web' })
                });
                
                if (response.ok) {
                    const { token } = await response.json();
                    this.token = token.token;
                    localStorage.setItem('auth_token', token.token);
                    await this.loadUser();
                    return true;
                }
            } catch (error) {
                console.error('Login failed:', error);
            } finally {
                this.loading = false;
            }
            
            return false;
        },
        
        async loadUser() {
            if (!this.token) return;
            
            try {
                const response = await fetch('/user-svc/self', {
                    headers: { 'Authorization': `Bearer ${this.token}` }
                });
                
                if (response.ok) {
                    this.user = await response.json();
                } else {
                    this.logout();
                }
            } catch (error) {
                console.error('Failed to load user:', error);
                this.logout();
            }
        },
        
        logout() {
            this.user = null;
            this.token = null;
            localStorage.removeItem('auth_token');
        }
    }
});
```

### Mobile Authentication

#### Flutter Authentication Service

```dart
// lib/services/auth_service.dart
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

class AuthService {
  static const String baseUrl = 'https://api.company.com';
  static const String tokenKey = 'auth_token';
  
  Future<bool> login(String slug, String password) async {
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user-svc/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'slug': slug,
          'password': password,
          'device': 'mobile'
        }),
      );
      
      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        final token = data['token']['token'];
        
        final prefs = await SharedPreferences.getInstance();
        await prefs.setString(tokenKey, token);
        
        return true;
      }
    } catch (e) {
      print('Login error: $e');
    }
    
    return false;
  }
  
  Future<Map<String, dynamic>?> getCurrentUser() async {
    final token = await getToken();
    if (token == null) return null;
    
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user-svc/self'),
        headers: {'Authorization': 'Bearer $token'},
      );
      
      if (response.statusCode == 200) {
        return jsonDecode(response.body);
      }
    } catch (e) {
      print('Failed to get user: $e');
    }
    
    return null;
  }
  
  Future<bool> hasPermission(String permission) async {
    final token = await getToken();
    if (token == null) return false;
    
    try {
      final response = await http.post(
        Uri.parse('$baseUrl/user-svc/self/has/$permission'),
        headers: {'Authorization': 'Bearer $token'},
      );
      
      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        return data['authorized'] ?? false;
      }
    } catch (e) {
      print('Permission check error: $e');
    }
    
    return false;
  }
  
  Future<String?> getToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(tokenKey);
  }
  
  Future<void> logout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(tokenKey);
  }
}
```

## API Reference Summary

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/user-svc/register` | POST | Create new user account |
| `/user-svc/login` | POST | Authenticate user/service |
| `/user-svc/self` | POST | Get current user info |
| `/user-svc/self/has/{permission}` | POST | Check permission |
| `/user-svc/users` | POST | List users (admin) |
| `/user-svc/permits` | PUT | Save permission grants |
| `/user-svc/permits` | POST | List permits |
| `/user-svc/enrolls` | PUT | Assign roles |
| `/user-svc/enrolls` | POST | List role enrollments |
| `/user-svc/organizations` | POST | List organizations |
| `/user-svc/organization` | PUT | Create/update organization |
| `/user-svc/refresh-token` | POST | Refresh authentication token |
| `/user-svc/public-key` | GET | Get JWT verification key |

## Related Documentation

- [User Svc API Reference](/docs/1backend-api/login) - Complete API documentation
- [Secret Svc](/docs/built-in-services/secret-svc) - Secure credential storage
- [Config Svc](/docs/built-in-services/config-svc) - Application configuration
- [Policy Svc](/docs/built-in-services/policy-svc) - Rate limiting and access policies
