---
sidebar_position: 100
tags:
  - data-svc
  - data
  - permissions
  - roles
  - authentication
  - authorization
  - services
  - firebase
  - backendless
---

# Data Svc

The Data Svc is a Firebase-like backendless database service that allows direct data operations from frontends, eliminating the need for basic CRUD microservices.

> This page provides a comprehensive overview of `Data Svc`. For detailed API information, refer to the [Data Svc API documentation](/docs/1backend-api/query-objects).

## Architecture & Purpose

Data Svc serves as a **lightweight database abstraction** designed for rapid prototyping and backendless applications. It provides:

- **Direct frontend access** to database operations
- **Multi-tenant data isolation** through permissions
- **Firebase-like flexibility** with schemaless JSON objects
- **Comprehensive permission system** for fine-grained access control

### Key Features

- **Tables & Objects**: Store schemaless JSON data in named tables
- **Permission-Based Access**: Control who can read, write, and delete data
- **Advanced Querying**: Filter, sort, and paginate results
- **Real-time Capabilities**: Integration with Firehose Svc for live updates
- **CLI & API Access**: Use both command-line and programmatic interfaces

## CLI Usage

Data Svc doesn't have dedicated CLI commands like Config Svc. Instead, use the generic HTTP commands to interact with the API endpoints:

### Creating Objects

```bash
# Create a single object
oo post /data-svc/object \
  --object.table="user_profiles" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.deleters="_self" \
  --object.data.name="John Doe" \
  --object.data.email="john@example.com" \
  --object.data.preferences.theme="dark"

# Create a pet object with specific ID
oo post /data-svc/object \
  --object.id="pet_fluffy123" \
  --object.table="pet" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.deleters="_self" \
  --object.data.name="Fluffy" \
  --object.data.species="cat" \
  --object.data.age=3
```

### Querying Objects

```bash
# Query all objects in a table
oo post /data-svc/objects \
  --table="user_profiles" \
  --readers="_self"

# Query with filtering
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.filters='[{"fields":["data.species"],"op":"equals","jsonValues":"[\"cat\"]"}]'

# Query with sorting and pagination
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","desc":true,"sortingType":"numeric"}]' \
  --query.limit=10

# Query with pagination using after
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","desc":true,"sortingType":"numeric"}]' \
  --query.limit=10 \
  --query.jsonAfter='[5]'
```

### Updating Objects

```bash
# Update objects by filter
oo post /data-svc/objects/update \
  --table="pet" \
  --filters='[{"fields":["data.species"],"op":"equals","jsonValues":"[\"cat\"]"}]' \
  --object.data.vaccinated=true

# Upsert specific object by ID
oo put /data-svc/object/pet_fluffy123 \
  --object.data.name="Fluffy Updated" \
  --object.data.age=4
```

### Deleting Objects

```bash
# Delete objects by filter
oo post /data-svc/objects/delete \
  --table="pet" \
  --filters='[{"fields":["data.species"],"op":"equals","jsonValues":"[\"dog\"]"}]'
```

## Object Structure & Data Model

### Basic Object Format

```json
{
  "id": "pet_fluffy123",
  "table": "pet",
  "readers": ["usr_12345", "org_67890"],
  "writers": ["usr_12345"],
  "deleters": ["usr_12345"],
  "authors": ["usr_12345"],
  "data": {
    "name": "Fluffy",
    "species": "cat",
    "age": 3,
    "preferences": {
      "food": "salmon",
      "toys": ["ball", "mouse"]
    }
  },
  "createdAt": "2023-01-01T10:00:00Z",
  "updatedAt": "2023-01-01T10:00:00Z"
}
```

### Table Naming & Object IDs

- **Table names**: Use singular form (e.g., `pet`, `user_profile`)
- **Object IDs**: Must be prefixed with table name (e.g., `pet_fluffy123`)
- **Auto-generated IDs**: Leave empty for automatic generation

```bash
# Auto-generated ID
oo post /data-svc/object \
  --object.table="pet" \
  --object.data.name="Auto Pet"

# Custom ID (must match table prefix)
oo post /data-svc/object \
  --object.id="pet_custom123" \
  --object.table="pet" \
  --object.data.name="Custom Pet"
```

## Permission System

The Data Svc permission model provides **fine-grained access control** with four permission types:

### Permission Types

#### **Readers**
Users/roles who can **view** objects:
```json
{
  "readers": ["usr_12345", "org_67890", "role_managers"]
}
```

#### **Writers** 
Users/roles who can **modify** objects:
```json
{
  "writers": ["usr_12345", "org_67890"]
}
```

#### **Deleters**
Users/roles who can **delete** objects:
```json
{
  "deleters": ["usr_12345"]
}
```

#### **Authors**
**System-assigned** creators (cannot be spoofed):
```json
{
  "authors": ["usr_12345", "org_67890"]
}
```

### Permission Shortcuts

- **`_self`**: Refers to the current user's ID
- **`_org`**: Refers to the current user's active organization

```bash
# Using permission shortcuts
oo post /data-svc/object \
  --object.table="private_note" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.deleters="_self" \
  --object.data.content="My private note"
```

### Multi-Tenant Access Control

Data Svc supports **shared tables** where multiple users store data with isolation:

```bash
# User A creates a profile
oo post /data-svc/object \
  --object.table="user_profile" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.data.name="Alice"

# User B creates a separate profile in same table
oo post /data-svc/object \
  --object.table="user_profile" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.data.name="Bob"
```

### Cross-User Permissions

Grant access to specific users or organizations:

```bash
# Share data with specific user
oo post /data-svc/object \
  --object.table="shared_document" \
  --object.readers='["usr_alice123", "usr_bob456"]' \
  --object.writers='["usr_alice123"]' \
  --object.deleters='["usr_alice123"]' \
  --object.data.title="Shared Project Plan"

# Share with organization
oo post /data-svc/object \
  --object.table="team_resource" \
  --object.readers='["org_company123"]' \
  --object.writers='["org_company123"]' \
  --object.data.resource="Team Guidelines"
```

## Advanced Querying

### Filter Operations

Data Svc supports comprehensive filtering with various operators:

```bash
# Equals filter
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.filters='[{"fields":["data.species"],"op":"equals","jsonValues":"[\"cat\"]"}]'

# Contains filter (for strings)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.filters='[{"fields":["data.name"],"op":"contains","jsonValues":"[\"flu\"]"}]'

# Greater than filter (for numbers)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.filters='[{"fields":["data.age"],"op":"greaterThan","jsonValues":"[5]"}]'

# Multiple filters (AND logic)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.filters='[
    {"fields":["data.species"],"op":"equals","jsonValues":"[\"cat\"]"},
    {"fields":["data.age"],"op":"greaterThan","jsonValues":"[2]"}
  ]'
```

### Sorting & Ordering

```bash
# Sort by age (ascending)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","sortingType":"numeric"}]'

# Sort by age (descending)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","desc":true,"sortingType":"numeric"}]'

# Sort by name (string)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.name","sortingType":"string"}]'
```

### Pagination

```bash
# First page (limit 5)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","sortingType":"numeric"}]' \
  --query.limit=5

# Next page (after age 3)
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  --query.orderBys='[{"field":"data.age","sortingType":"numeric"}]' \
  --query.limit=5 \
  --query.jsonAfter='[3]'
```

## Real-World Usage Examples

### 1. User Profile Management

```bash
# Create user profile
oo post /data-svc/object \
  --object.table="user_profile" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.deleters="_self" \
  --object.data.name="John Doe" \
  --object.data.email="john@example.com" \
  --object.data.preferences.theme="dark" \
  --object.data.preferences.notifications=true

# Update profile preferences
oo post /data-svc/objects/update \
  --table="user_profile" \
  --filters='[{"fields":["authors"],"op":"intersects","jsonValues":"[\"usr_currentuser\"]"}]' \
  --object.data.preferences.theme="light"

# Query own profile
oo post /data-svc/objects \
  --table="user_profile" \
  --readers="_self"
```

### 2. Todo List Application

```bash
# Create todo item
oo post /data-svc/object \
  --object.table="todo" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.deleters="_self" \
  --object.data.title="Learn Data Svc" \
  --object.data.completed=false \
  --object.data.priority="high" \
  --object.data.dueDate="2023-12-31"

# Mark todo as completed
oo post /data-svc/objects/update \
  --table="todo" \
  --filters='[{"fields":["data.title"],"op":"equals","jsonValues":"[\"Learn Data Svc\"]"}]' \
  --object.data.completed=true

# Get pending todos
oo post /data-svc/objects \
  --table="todo" \
  --readers="_self" \
  --query.filters='[{"fields":["data.completed"],"op":"equals","jsonValues":"[false]"}]' \
  --query.orderBys='[{"field":"data.priority","sortingType":"string"}]'
```

### 3. Team Collaboration

```bash
# Create shared team document
oo post /data-svc/object \
  --object.table="team_document" \
  --object.readers='["org_team123"]' \
  --object.writers='["org_team123"]' \
  --object.deleters='["usr_teamlead456"]' \
  --object.data.title="Project Specs" \
  --object.data.content="Initial project specifications..." \
  --object.data.status="draft"

# Add team member comments
oo post /data-svc/object \
  --object.table="document_comment" \
  --object.readers='["org_team123"]' \
  --object.writers='["org_team123"]' \
  --object.data.documentId="team_document_proj123" \
  --object.data.author="Alice Smith" \
  --object.data.comment="Looks good, needs timeline"

# Query team documents
oo post /data-svc/objects \
  --table="team_document" \
  --readers='["org_team123"]' \
  --query.filters='[{"fields":["data.status"],"op":"equals","jsonValues":"[\"draft\"]"}]'
```

### 4. E-commerce Product Catalog

```bash
# Create product
oo post /data-svc/object \
  --object.table="product" \
  --object.readers='["any"]' \
  --object.writers='["role_admin"]' \
  --object.deleters='["role_admin"]' \
  --object.data.name="Premium Headphones" \
  --object.data.price=199.99 \
  --object.data.category="electronics" \
  --object.data.inStock=true \
  --object.data.tags='["wireless", "premium", "audio"]'

# Search products by category
oo post /data-svc/objects \
  --table="product" \
  --readers='["any"]' \
  --query.filters='[{"fields":["data.category"],"op":"equals","jsonValues":"[\"electronics\"]"}]' \
  --query.orderBys='[{"field":"data.price","sortingType":"numeric"}]'

# Filter by price range
oo post /data-svc/objects \
  --table="product" \
  --readers='["any"]' \
  --query.filters='[
    {"fields":["data.price"],"op":"greaterThanOrEqual","jsonValues":"[100]"},
    {"fields":["data.price"],"op":"lessThanOrEqual","jsonValues":"[300]"}
  ]'
```

## Integration Patterns

### With Firehose Svc (Real-time Updates)

Data Svc automatically publishes events to Firehose Svc for real-time capabilities:

```bash
# Create object (triggers real-time event)
oo post /data-svc/object \
  --object.table="chat_message" \
  --object.readers='["org_chatroom123"]' \
  --object.data.message="Hello everyone!" \
  --object.data.sender="Alice"

# Subscribe to real-time updates via Firehose
oo get /firehose-svc/events \
  --topic="data-svc:chat_message"
```

### With File Svc (Attachments)

Store file references in Data Svc objects:

```bash
# Create object with file references
oo post /data-svc/object \
  --object.table="document" \
  --object.readers="_self" \
  --object.writers="_self" \
  --object.data.title="Report with Attachments" \
  --object.data.fileIds='["file_document123", "file_image456"]'
```

## Access Control & Security

### Default Permissions

- **Public Read**: Use `"any"` in readers for public access
- **Organization Access**: Use `"org_id"` for organization-wide access
- **Role-Based**: Use `"role_name"` for role-based permissions

```bash
# Public read, admin write
oo post /data-svc/object \
  --object.table="public_announcement" \
  --object.readers='["any"]' \
  --object.writers='["role_admin"]' \
  --object.deleters='["role_admin"]' \
  --object.data.title="System Maintenance Notice"

# Organization-wide access
oo post /data-svc/object \
  --object.table="company_policy" \
  --object.readers='["org_company123"]' \
  --object.writers='["role_hr"]' \
  --object.data.policy="Remote Work Guidelines"
```

### Security Best Practices

1. **Principle of Least Privilege**: Grant minimal required permissions
2. **Validate Authors**: Use authors field to verify data origin
3. **Audit Trails**: Track who created/modified objects
4. **Sensitive Data**: Use [Secret Svc](/docs/built-in-services/secret-svc) for secrets

## Error Handling & Troubleshooting

### Common Issues

#### **Permission Denied**
```bash
# Error: User cannot access object
# Solution: Check readers/writers/deleters permissions
oo post /data-svc/objects \
  --table="restricted_data" \
  --readers='["usr_correctuser"]'
```

#### **Invalid Object ID**
```bash
# Error: ID doesn't match table prefix
# Solution: Ensure ID starts with table name
oo post /data-svc/object \
  --object.id="pet_fluffy123" \
  --object.table="pet"
```

#### **Empty Query Results**
```bash
# Check permissions and filters
oo post /data-svc/objects \
  --table="pet" \
  --readers="_self" \
  # Verify you have read access to objects
```

## Performance Considerations

### Indexing Strategy

- **Filter Fields**: Common filter fields are automatically indexed
- **Sort Fields**: Specify sortingType for optimal performance
- **Limit Results**: Use pagination for large datasets

### Query Optimization

```bash
# Good: Specific filters with limits
oo post /data-svc/objects \
  --table="large_dataset" \
  --readers="_self" \
  --query.filters='[{"fields":["data.status"],"op":"equals","jsonValues":"[\"active\"]"}]' \
  --query.limit=50

# Avoid: Querying all objects without filters
oo post /data-svc/objects \
  --table="large_dataset" \
  --readers="_self"
```

## API Reference Summary

| Endpoint | Method | Purpose |
|----------|---------|---------|
| `/data-svc/object` | POST | Create single object |
| `/data-svc/objects` | POST | Query objects |
| `/data-svc/objects/update` | POST | Update objects by filter |
| `/data-svc/objects/delete` | POST | Delete objects by filter |
| `/data-svc/object/{id}` | PUT | Upsert object by ID |

## Related Services

- **[Secret Svc](/docs/built-in-services/secret-svc)**: Store sensitive data securely
- **[Firehose Svc](/docs/built-in-services/firehose-svc)**: Real-time event streaming
- **[File Svc](/docs/built-in-services/file-svc)**: File storage and management
- **[User Svc](/docs/built-in-services/user-svc)**: User authentication and roles

## Limitations

- **Schema Flexibility**: Objects are schemaless but lack schema validation
- **Complex Queries**: Limited support for complex joins or aggregations
- **Transaction Support**: No multi-object transactions
- **Storage Limits**: Consider using File Svc for large binary data

For production applications requiring complex data relationships, consider using dedicated databases alongside Data Svc for specific use cases.
