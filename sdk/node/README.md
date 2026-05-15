# @1backend/sdk-node

Node.js SDK for writing and launching custom 1Backend services.

This package is intentionally different from `@1backend/node`:

- `@1backend/node` is the generated API client.
- `@1backend/sdk-node` is the service SDK (boot options, account registration, auth middlewares, token refresh, and permission checks).

## Install

```bash
npm i @1backend/sdk-node @1backend/client
```

## Quick start

```ts
import {
  createDefaultOptions,
  registerServiceAccount,
  registerServiceInstance,
  listenAddress,
  createCredentialFileStore,
} from "@1backend/sdk-node";

const options = await createDefaultOptions({
  serverUrl:
    process.env.OB_INTERNAL_SERVER_URL ??
    process.env.OB_SERVER_URL ??
    "http://127.0.0.1:11337",
  selfUrl:
    process.env.OB_PUBLIC_URL ??
    process.env.OB_SELF_URL ??
    "http://127.0.0.1:9111",
});

const store = createCredentialFileStore("./.service-credential.json");
const token = await registerServiceAccount(
  options.api.userSvc,
  "my-svc",
  "My Service",
  store,
);

await registerServiceInstance(options.clientFactory.create(token.token).registrySvc, options.selfUrl);

console.log(`Listen on ${listenAddress(options.selfUrl)}`);
```
