---
sidebar_position: 70
tags:
  - policy-svc
  - rate-limiting
  - blocklist
  - security
  - reliability
  - policies
  - access-control
  - ddos-protection
---

# Policy Svc

The Policy Svc is a comprehensive access control and rate limiting service that provides protection against abuse, ensures fair resource usage, and maintains service reliability through configurable policies.

> This page provides a comprehensive overview of `Policy Svc`. For detailed API information, refer to the [Policy Svc API documentation](/docs/1backend-api/upsert-instance).

## Architecture & Purpose

Policy Svc serves as the **service protection layer** for 1Backend, providing:

- **Rate Limiting**: Prevent resource exhaustion by limiting requests per user/IP/endpoint
- **Access Control**: Block malicious IPs or users from accessing services
- **Abuse Prevention**: Protect against DDoS attacks and automated abuse
- **Fair Usage**: Ensure equitable resource distribution across users
- **Service Reliability**: Maintain system stability under high load

### Key Features

- **Flexible Templates**: Pre-built policies for common protection scenarios
- **Multiple Entities**: Rate limit by user ID, IP address, or custom identifiers
- **Granular Scoping**: Apply policies per-endpoint or globally across services
- **Time Windows**: Configurable periods (seconds, minutes, hours, days)
- **Explicit Integration**: Services opt-in to policy checking for maximum control

## CLI Usage

Policy Svc uses HTTP commands for all operations:

### Creating Policy Instances

```bash
# Rate limit user registrations: 5 per IP per day
oo put /policy-svc/instance/registration-rate-limit \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="1d" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Block malicious IPs from login endpoint
oo put /policy-svc/instance/login-ip-blocklist \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="blocklist" \
  --instance.parameters.blocklist.blockedIPs='["192.168.1.100", "10.0.0.50"]'

# Rate limit API calls: 100 per user per hour across all endpoints
oo put /policy-svc/instance/api-user-rate-limit \
  --instance.endpoint="" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="global"
```

### Checking Access Control

```bash
# Check if request is allowed (call from your service)
oo post /policy-svc/check \
  --endpoint="/user-svc/register" \
  --method="POST" \
  --ip="192.168.1.10" \
  --userId="usr_12345"

# Response: {"allowed": true} or {"allowed": false}
```

### Policy Management

```bash
# Update existing policy (same endpoint, overwrites)
oo put /policy-svc/instance/registration-rate-limit \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="1d" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Multiple policies can apply to the same endpoint
oo put /policy-svc/instance/registration-user-rate-limit \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=3 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

## Policy Templates

### Rate Limit Template

The Rate Limit template provides flexible request throttling:

```json
{
  "templateId": "rate-limit",
  "parameters": {
    "rateLimit": {
      "maxRequests": 10,
      "timeWindow": "1m",
      "entity": "ip",
      "scope": "endpoint"
    }
  }
}
```

**Parameters:**
- `maxRequests`: Maximum number of requests allowed
- `timeWindow`: Time period for the limit (e.g., "30s", "5m", "1h", "24h", "7d")
- `entity`: Who to track (`"ip"`, `"userId"`)
- `scope`: Where to apply (`"endpoint"`, `"global"`)

**Time Window Formats:**
```bash
"30s"    # 30 seconds
"5m"     # 5 minutes  
"1h"     # 1 hour
"24h"    # 24 hours
"7d"     # 7 days
"168h"   # 1 week (alternative)
```

**Entity Types:**
- `"ip"`: Rate limit by client IP address
- `"userId"`: Rate limit by authenticated user ID

**Scope Types:**
- `"endpoint"`: Limit applies only to the specified endpoint
- `"global"`: Limit applies across all endpoints for the entity

### Blocklist Template

The Blocklist template provides IP-based access control:

```json
{
  "templateId": "blocklist", 
  "parameters": {
    "blocklist": {
      "blockedIPs": ["192.168.1.100", "10.0.0.50", "203.0.113.0"]
    }
  }
}
```

**Parameters:**
- `blockedIPs`: Array of IP addresses to block

## Service Integration Patterns

### Explicit Policy Checking

Unlike middleware-based solutions, Policy Svc requires **explicit integration**:

```go
// In your service endpoint
func (s *MyService) CreatePost(w http.ResponseWriter, r *http.Request) {
    // Extract request context
    userID := getUserID(r)
    clientIP := getClientIP(r)
    
    // Check policy
    checkRsp, _, err := s.client.PolicySvcAPI.Check(r.Context()).
        Body(openapi.PolicySvcCheckRequest{
            Endpoint: "/posts-svc/post",
            Method:   "POST", 
            Ip:       clientIP,
            UserId:   userID,
        }).Execute()
    
    if err != nil {
        http.Error(w, "Policy check failed", 500)
        return
    }
    
    if !checkRsp.Allowed {
        http.Error(w, "Rate limit exceeded", 429)
        return
    }
    
    // Proceed with business logic
    s.createPost(r)
}
```

### CLI-Based Integration

```bash
# In a shell script or CLI-based service
USER_ID="usr_12345"
CLIENT_IP="192.168.1.10"

ALLOWED=$(oo post /policy-svc/check \
  --endpoint="/api/data/upload" \
  --method="POST" \
  --ip="$CLIENT_IP" \
  --userId="$USER_ID" | jq -r '.allowed')

if [ "$ALLOWED" = "true" ]; then
  echo "Processing request..."
  # Handle the request
else
  echo "Request blocked by policy"
  exit 1
fi
```

## Real-World Usage Examples

### 1. User Registration Protection

```bash
# Prevent registration abuse: 3 registrations per IP per day
oo put /policy-svc/instance/registration-ip-limit \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=3 \
  --instance.parameters.rateLimit.timeWindow="24h" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Also limit per user: 1 registration per user per hour (prevent multiple accounts)
oo put /policy-svc/instance/registration-user-limit \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Test the policy
oo post /policy-svc/check \
  --endpoint="/user-svc/register" \
  --method="POST" \
  --ip="192.168.1.10" \
  --userId="usr_test"
```

### 2. API Rate Limiting

```bash
# General API rate limiting: 1000 requests per user per hour
oo put /policy-svc/instance/api-rate-limit \
  --instance.endpoint="" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1000 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="global"

# Expensive operations: 10 AI prompts per user per hour
oo put /policy-svc/instance/prompt-rate-limit \
  --instance.endpoint="/prompt-svc/prompt" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# File uploads: 5 uploads per user per 10 minutes
oo put /policy-svc/instance/upload-rate-limit \
  --instance.endpoint="/file-svc/upload" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="10m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

### 3. DDoS Protection

```bash
# Aggressive rate limiting for login attempts
oo put /policy-svc/instance/login-ip-protection \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="5m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Block known malicious IPs
oo put /policy-svc/instance/malicious-ip-blocklist \
  --instance.endpoint="" \
  --instance.templateId="blocklist" \
  --instance.parameters.blocklist.blockedIPs='["185.220.101.0", "185.220.102.0", "tor-exit-node-1.com"]'

# Emergency rate limiting during attack
oo put /policy-svc/instance/emergency-rate-limit \
  --instance.endpoint="" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="global"
```

### 4. Fair Resource Usage

```bash
# Chat service: 50 messages per user per hour
oo put /policy-svc/instance/chat-rate-limit \
  --instance.endpoint="/chat-svc/thread/*/message" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=50 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Data service: 100 object creations per user per day
oo put /policy-svc/instance/data-creation-limit \
  --instance.endpoint="/data-svc/object" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="24h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Model usage: 20 model starts per user per day (expensive operation)
oo put /policy-svc/instance/model-start-limit \
  --instance.endpoint="/model-svc/model/*/start" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=20 \
  --instance.parameters.rateLimit.timeWindow="24h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

### 5. Abuse Prevention

```bash
# Prevent password reset abuse: 3 attempts per IP per hour
oo put /policy-svc/instance/password-reset-limit \
  --instance.endpoint="/user-svc/password-reset" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=3 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Email verification: 5 requests per user per day
oo put /policy-svc/instance/email-verification-limit \
  --instance.endpoint="/user-svc/verify-email" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="24h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Account creation from specific regions (compliance)
oo put /policy-svc/instance/region-blocklist \
  --instance.endpoint="/user-svc/register" \
  --instance.templateId="blocklist" \
  --instance.parameters.blocklist.blockedIPs='["192.0.2.0", "203.0.113.0"]'
```

### 6. Service-Specific Policies

```bash
# Container service: Limit container launches
oo put /policy-svc/instance/container-launch-limit \
  --instance.endpoint="/container-svc/container" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Image generation: Expensive AI operations
oo put /policy-svc/instance/image-generation-limit \
  --instance.endpoint="/prompt-svc/prompt" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Secret management: Sensitive operations
oo put /policy-svc/instance/secret-access-limit \
  --instance.endpoint="/secret-svc/secret" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=50 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

## Policy Patterns & Best Practices

### Layered Protection Strategy

```bash
# Layer 1: Global IP-based protection (very permissive)
oo put /policy-svc/instance/global-ip-protection \
  --instance.endpoint="" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1000 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="global"

# Layer 2: Endpoint-specific limits (moderate)
oo put /policy-svc/instance/api-endpoint-limit \
  --instance.endpoint="/api/v1/data" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Layer 3: User-based fair usage (strict)
oo put /policy-svc/instance/user-fair-usage \
  --instance.endpoint="/api/v1/data" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=50 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

### Time Window Strategies

```bash
# Burst protection: Short-term limits
oo put /policy-svc/instance/burst-protection \
  --instance.endpoint="/api/search" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="30s" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Sustained usage: Medium-term limits  
oo put /policy-svc/instance/sustained-usage \
  --instance.endpoint="/api/search" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="10m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# Daily quotas: Long-term limits
oo put /policy-svc/instance/daily-quota \
  --instance.endpoint="/api/search" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1000 \
  --instance.parameters.rateLimit.timeWindow="24h" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

## Monitoring & Observability

### Policy Effectiveness Testing

```bash
# Test rate limit behavior
test_rate_limit() {
  local endpoint=$1
  local max_requests=$2
  
  echo "Testing rate limit for $endpoint (max: $max_requests)"
  
  for i in $(seq 1 $((max_requests + 2))); do
    result=$(oo post /policy-svc/check \
      --endpoint="$endpoint" \
      --method="POST" \
      --ip="192.168.1.10" \
      --userId="test_user" | jq -r '.allowed')
    
    echo "Request $i: $result"
  done
}

# Test registration endpoint
test_rate_limit "/user-svc/register" 5
```

### Policy Audit Commands

```bash
# Check policy coverage for critical endpoints
critical_endpoints=(
  "/user-svc/register"
  "/user-svc/login" 
  "/prompt-svc/prompt"
  "/file-svc/upload"
  "/container-svc/container"
)

for endpoint in "${critical_endpoints[@]}"; do
  echo "Testing policy for $endpoint"
  oo post /policy-svc/check \
    --endpoint="$endpoint" \
    --method="POST" \
    --ip="192.168.1.10" \
    --userId="audit_user"
done
```

### Load Testing Integration

```bash
# Use with load testing tools
load_test_with_policy() {
  local endpoint=$1
  local concurrent_users=$2
  
  # Pre-check policy status
  oo post /policy-svc/check \
    --endpoint="$endpoint" \
    --method="POST" \
    --ip="192.168.1.10" \
    --userId="load_test_user"
  
  # Run load test (example with curl)
  for i in $(seq 1 $concurrent_users); do
    curl -X POST "$endpoint" \
      -H "Authorization: Bearer $TOKEN" \
      -H "X-Forwarded-For: 192.168.1.$i" &
  done
  wait
}
```

## Policy Response Handling

### Service Implementation Patterns

```bash
# Standard response handling in services
handle_policy_check() {
  local endpoint=$1
  local user_id=$2
  local client_ip=$3
  
  response=$(oo post /policy-svc/check \
    --endpoint="$endpoint" \
    --method="POST" \
    --ip="$client_ip" \
    --userId="$user_id")
  
  allowed=$(echo "$response" | jq -r '.allowed')
  
  if [ "$allowed" = "true" ]; then
    return 0  # Allow request
  else
    echo "HTTP/1.1 429 Too Many Requests"
    echo "Content-Type: application/json"
    echo ""
    echo '{"error":"Rate limit exceeded","retryAfter":"60"}'
    return 1  # Block request
  fi
}

# Usage in endpoint handler
if handle_policy_check "/api/data/create" "$USER_ID" "$CLIENT_IP"; then
  # Process the request
  echo "Processing request..."
else
  # Request was blocked
  exit 1
fi
```

### HTTP Status Code Standards

```bash
# Proper HTTP responses for policy violations

# Rate limiting
echo "HTTP/1.1 429 Too Many Requests"
echo "Retry-After: 60"
echo "X-RateLimit-Limit: 100"
echo "X-RateLimit-Remaining: 0"
echo "X-RateLimit-Reset: 1640995200"

# IP blocking  
echo "HTTP/1.1 403 Forbidden"
echo "Content-Type: application/json"
echo '{"error":"Access denied from this IP address"}'

# General policy violation
echo "HTTP/1.1 403 Forbidden" 
echo "Content-Type: application/json"
echo '{"error":"Request blocked by security policy"}'
```

## Advanced Configuration

### Dynamic Policy Updates

```bash
# Update policy during high load (emergency response)
oo put /policy-svc/instance/emergency-login-limit \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1 \
  --instance.parameters.rateLimit.timeWindow="5m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Relax policies during maintenance
oo put /policy-svc/instance/maintenance-mode \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=50 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Remove emergency restrictions
oo put /policy-svc/instance/normal-login-limit \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=10 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"
```

### Conditional Policies

```bash
# Different limits for different user types (implemented in service logic)
check_policy_with_user_tier() {
  local user_id=$1
  local endpoint=$2
  local client_ip=$3
  
  # Get user tier from user service
  user_tier=$(oo get /user-svc/user/$user_id | jq -r '.user.tier')
  
  case $user_tier in
    "premium")
      instance_id="premium-user-limit"
      ;;
    "basic")
      instance_id="basic-user-limit"
      ;;
    *)
      instance_id="default-user-limit"
      ;;
  esac
  
  # Check appropriate policy
  oo post /policy-svc/check \
    --endpoint="$endpoint" \
    --method="POST" \
    --ip="$client_ip" \
    --userId="$user_id"
}
```

## Troubleshooting

### Common Issues

#### **Policy Not Applying**
```bash
# Check if policy instance exists
oo put /policy-svc/instance/test-policy \
  --instance.endpoint="/test-endpoint" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=1 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# Test the policy immediately
oo post /policy-svc/check \
  --endpoint="/test-endpoint" \
  --method="POST" \
  --ip="192.168.1.10" \
  --userId="test_user"
```

#### **Rate Limits Too Restrictive**
```bash
# Check current rate limit status
oo post /policy-svc/check \
  --endpoint="/api/problematic-endpoint" \
  --method="POST" \
  --ip="192.168.1.10" \
  --userId="affected_user"

# Temporarily increase limits
oo put /policy-svc/instance/temporary-increase \
  --instance.endpoint="/api/problematic-endpoint" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

#### **Blocked IPs**
```bash
# Check if IP is in blocklist
oo post /policy-svc/check \
  --endpoint="/any-endpoint" \
  --method="GET" \
  --ip="SUSPECTED_IP" \
  --userId="any_user"

# Update blocklist to remove IP
oo put /policy-svc/instance/ip-blocklist \
  --instance.endpoint="" \
  --instance.templateId="blocklist" \
  --instance.parameters.blocklist.blockedIPs='["192.168.1.100"]'  # Removed the IP
```

#### **Policy Conflicts**
```bash
# Multiple policies can apply to the same endpoint
# They are evaluated in sequence - if ANY policy blocks, request is denied

# Example: Both IP and user rate limits apply
oo put /policy-svc/instance/endpoint-ip-limit \
  --instance.endpoint="/api/data" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=100 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

oo put /policy-svc/instance/endpoint-user-limit \
  --instance.endpoint="/api/data" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=50 \
  --instance.parameters.rateLimit.timeWindow="1m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"
```

### Debug Commands

```bash
# Test policy behavior with different parameters
debug_policy() {
  local endpoint=$1
  local user_id=$2
  local ip=$3
  
  echo "=== Policy Debug for $endpoint ==="
  echo "User: $user_id, IP: $ip"
  
  response=$(oo post /policy-svc/check \
    --endpoint="$endpoint" \
    --method="POST" \
    --ip="$ip" \
    --userId="$user_id")
  
  echo "Response: $response"
  
  allowed=$(echo "$response" | jq -r '.allowed')
  if [ "$allowed" = "true" ]; then
    echo "✅ Request ALLOWED"
  else
    echo "❌ Request BLOCKED"
  fi
}

# Usage
debug_policy "/user-svc/login" "usr_test" "192.168.1.10"
```

## Security Considerations

### Policy Security Best Practices

```bash
# 1. Protect admin endpoints aggressively
oo put /policy-svc/instance/admin-protection \
  --instance.endpoint="/admin/*" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="1h" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

# 2. Layer multiple protection mechanisms
oo put /policy-svc/instance/login-ip-limit \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=5 \
  --instance.parameters.rateLimit.timeWindow="5m" \
  --instance.parameters.rateLimit.entity="ip" \
  --instance.parameters.rateLimit.scope="endpoint"

oo put /policy-svc/instance/login-user-limit \
  --instance.endpoint="/user-svc/login" \
  --instance.templateId="rate-limit" \
  --instance.parameters.rateLimit.maxRequests=3 \
  --instance.parameters.rateLimit.timeWindow="5m" \
  --instance.parameters.rateLimit.entity="userId" \
  --instance.parameters.rateLimit.scope="endpoint"

# 3. Block known malicious IP ranges
oo put /policy-svc/instance/tor-blocklist \
  --instance.endpoint="" \
  --instance.templateId="blocklist" \
  --instance.parameters.blocklist.blockedIPs='["185.220.101.0", "185.220.102.0"]'
```

### Incident Response

```bash
# Emergency lockdown procedure
emergency_lockdown() {
  echo "🚨 Implementing emergency lockdown..."
  
  # Severely restrict all endpoints
  oo put /policy-svc/instance/emergency-global-limit \
    --instance.endpoint="" \
    --instance.templateId="rate-limit" \
    --instance.parameters.rateLimit.maxRequests=1 \
    --instance.parameters.rateLimit.timeWindow="1m" \
    --instance.parameters.rateLimit.entity="ip" \
    --instance.parameters.rateLimit.scope="global"
  
  echo "✅ Emergency lockdown active"
}

# Gradual recovery procedure
gradual_recovery() {
  echo "🔄 Starting gradual recovery..."
  
  # Increase limits gradually
  oo put /policy-svc/instance/recovery-phase-1 \
    --instance.endpoint="" \
    --instance.templateId="rate-limit" \
    --instance.parameters.rateLimit.maxRequests=10 \
    --instance.parameters.rateLimit.timeWindow="1m" \
    --instance.parameters.rateLimit.entity="ip" \
    --instance.parameters.rateLimit.scope="global"
  
  sleep 300  # Wait 5 minutes
  
  oo put /policy-svc/instance/recovery-phase-2 \
    --instance.endpoint="" \
    --instance.templateId="rate-limit" \
    --instance.parameters.rateLimit.maxRequests=50 \
    --instance.parameters.rateLimit.timeWindow="1m" \
    --instance.parameters.rateLimit.entity="ip" \
    --instance.parameters.rateLimit.scope="global"
  
  echo "✅ Recovery phase 2 complete"
}
```

## API Reference Summary

| Endpoint | Method | Purpose |
|----------|---------|---------|
| `/policy-svc/instance/{instanceId}` | PUT | Create or update policy instance |
| `/policy-svc/check` | POST | Check if request is allowed by policies |

## Permissions & Security

```bash
# Admin permissions (required for policy management)
policy-svc:instance:create    # Create policy instances
policy-svc:instance:edit      # Modify policy instances
policy-svc:instance:delete    # Remove policy instances
policy-svc:template:view      # View policy templates

# User permissions (for policy checking)
policy-svc:check:view         # Check policies (typically granted to services)
```

## Related Services

- **[User Svc](/docs/built-in-services/user-svc)**: User authentication and registration protection
- **[File Svc](/docs/built-in-services/file-svc)**: Upload rate limiting and abuse prevention
- **[Prompt Svc](/docs/built-in-services/prompt-svc)**: AI usage rate limiting and fair access
- **[Container Svc](/docs/built-in-services/container-svc)**: Resource-intensive operation limiting

## Future Enhancements

### Planned Features
- **Geographic Blocking**: Block requests from specific countries/regions
- **User Behavior Analysis**: Dynamic rate limiting based on user patterns
- **Custom Templates**: User-defined policy templates
- **Policy Analytics**: Detailed reporting on policy effectiveness
- **Whitelist Support**: Allow-lists for trusted IPs/users

### Integration Roadmap
- **Automatic Detection**: AI-powered abuse detection and policy adjustment
- **External Feeds**: Integration with threat intelligence feeds
- **Circuit Breakers**: Automatic service protection during overload
- **Policy Scheduling**: Time-based policy activation/deactivation
- **Multi-Factor Policies**: Complex policies combining multiple conditions

Policy Svc provides the essential protection layer for 1Backend services, ensuring system reliability, fair resource usage, and protection against abuse through flexible, configurable policies.
