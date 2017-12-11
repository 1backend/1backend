### [&#8592; Back](README.md)

# Quotas

Quotas are intended to simplify the cost calculation of platform usage.
When quotas are enabled, each user has a quota he can distribute amongst his tokens.

Every service call reduces this quota - for now only by 1, but this will likely change later when we introduce long running apps (right now the max timeout for a service call is 10 seconds), and more expensive operations.
