---
sidebar_position: 30
tags:
  - config-svc
  - configuration
  - services
  - multitenant
---

# Config Svc

The Config Svc stores public, non-sensitive and end-user-facing data.

> This page provides a high-level overview of `Config Svc`. For detailed information, refer to the [Config Svc API documentation](/docs/1backend-api/list-configs).

The Config Svc is less critical than it might seem—most configuration happens internally through the [Secret Svc](/docs/built-in-services/secret-svc).

At the moment, it functions more like a minimal feature-flag service.

## CLI Usage

### Saving Configs

Use the `config save` command to save configuration data from YAML files:

```bash
# Save a single config file
oo config save ./my-config.yaml

# Save multiple config files from a directory
oo config save ./configs/

# Using aliases
oo co s ./my-config.yaml
oo configs save ./my-config.yaml
```

#### Config File Structure

Config files should be in YAML format with the following structure:

```yaml
# Optional: specify target app.
# Only users with the permission "config-svc:config:edit-on-behalf" (typically admins) can specify this.
# For other users the app from the token will be used.
app: "my-app"
# Optional: specify target app.
# Only users with the permission "config-svc:config:edit-on-behalf" (typically admins) can specify this.
# For other users the slug of the user will be used.
key: "user-slug"
data:
  # Your configuration data here
  featureFlags:
    enableNewUI: true
    maxUsers: 100
  settings:
    theme: "dark"
    language: "en"
    notifications:
      email: true
      push: false
```

#### Multiple Configs in One File

You can also define multiple configs in a single YAML file as an array:

```yaml
- app: "my-app"
  key: "user1"
  data:
    preferences:
      theme: "light"
- app: "my-app"
  key: "user2"
  data:
    preferences:
      theme: "dark"
```

### Querying Configs

Use the `config list` command to retrieve and view configurations:

```bash
# List all configs for the current app
oo config list

# Using aliases
oo co ls
oo configs list
```

The output shows configs in a tabular format:

```
CONFIG KEY    APP      JSON
user1         my-app   {"preferences":{"theme":"light"}}
user2         my-app   {"preferences":{"theme":"dark"}}
```

## Data Structure & Behavior

### Deep Merge

The Config Svc performs **deep merging** when saving configurations:

- **Nested objects** are recursively merged rather than replaced
- **Conflicting primitive values** (strings, numbers, booleans) are replaced by incoming values
- **New fields** are added to existing configs
- **Existing fields** not present in the incoming config are preserved

#### Deep Merge Example

Existing config:

```json
{
  "ui": {
    "theme": "light",
    "sidebar": {
      "collapsed": false,
      "width": 250
    }
  },
  "features": {
    "notifications": true
  }
}
```

Incoming config:

```yaml
data:
  ui:
    theme: "dark" # This will replace "light"
    sidebar:
      width: 300 # This will replace 250, collapsed: false is preserved
  newFeature: true # This will be added
```

Result after merge:

```json
{
  "ui": {
    "theme": "dark",        # Updated
    "sidebar": {
      "collapsed": false,   # Preserved
      "width": 300          # Updated
    }
  },
  "features": {
    "notifications": true   # Preserved
  },
  "newFeature": true        # Added
}
```

### Config Identification

Configs are uniquely identified by:

- **App**: The application name. Automatically determined from user's token unless the caller has the `"config-svc:config:edit-on-behalf"` permission - in that case the caller can specify this field.
- **Key**: The camelCased user slug (e.g., `jane-doe` becomes `janeDoе`). Automatically determined from the user's slug unless the caller has the `"config-svc:config:edit-on-behalf"` permission - in that case the caller can specify this field.

The internal ID format is: `{app}_{camelCasedSlug}`

## Access Rules

### Read Access

All configs are **publicly readable** even without a user account.

### Write Access

#### Regular Users

- Can only write to configs under their own slug
- The `key` field is automatically set to their camelCased slug
- The `app` field is automatically set from their token
- Cannot specify custom `app` or `key` values

#### Administrators

- Have the `config-svc:config:edit-on-behalf` permission
- Can specify any `key` value to edit configs for other users
- Can specify any `app` value to target different applications
- **Must** provide a `key` when editing on behalf of others

#### Permission-Based Access Example

Regular user with slug `jane-doe`:

```yaml
# This file will automatically save under key "janeDoe"
data:
  preferences:
    theme: "dark"
```

User with the `config-svc:config:edit-on-behalf` permission (e.g., admin):

```yaml
app: "production-app" # Can specify target app
key: "john-smith" # Can specify target user
data:
  adminSettings:
    maxFileSize: "10MB"
```

## Related

- [Secret Svc](/docs/built-in-services/secret-svc) to store sensitive data like internal configuration and secrets
- [Config Svc API documentation](/docs/1backend-api/list-configs) for programmatic access
