# 1backend Login Spirit

1backend treats contact ownership as the login primitive.

If a user proves control of an email address or phone number, that verified
contact is enough to enter the matching 1backend account, unless the account
has an additional local requirement such as TOTP. Provider-specific user ids
are implementation details for the provider, not the identity model of
1backend.

OIDC differs from this in a useful but important way: OIDC providers expose a
stable `sub` claim for the provider account, and OIDC client libraries usually
encourage applications to bind local users to that provider subject. 1backend
does not make that the primary rule. For Google and other OIDC providers,
1backend uses OIDC to verify that the provider asserts `email_verified=true`
for an email contact, then logs in by that contact.

Provider credentials are app-host scoped, not process-global. Users are global
in 1backend, but OAuth/OIDC client ids, secrets, redirect branding, and consent
screens belong to the host that starts the login. Store those provider configs
in the Secret Svc for the app host and grant `user-svc` read access.

Providers that do not prove a real contact must not be enabled for contact
login. Relay or noreply addresses are explicitly out of scope because they are
not the user's real email contact in this model.

Provider-specific OAuth2 integrations must follow the same rule: they can only
log in when the provider API proves a real email contact for the configured
application. Provider account ids, usernames, or mutable profile handles are
not enough.
