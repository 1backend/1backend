<p align="center">
  <img width="150px" src="./docs-source/1b_logo.svg" style="border-radius: 50%;" />
  <div align="center">
    <span>
      <h1 style="border-bottom: none">1Backend</h1>
      <a href="https://discord.gg/eRXyzeXEvM" rel="nofollow"><img src="https://camo.githubusercontent.com/66351093b042f69e9698398d33f08a6c36f1b7c56e1494b1e2902950eb24c94f/68747470733a2f2f646362616467652e6c696d65732e70696e6b2f6170692f7365727665722f68747470733a2f2f646973636f72642e67672f655258797a655845764d" alt="" data-canonical-src="https://dcbadge.limes.pink/api/server/https://discord.gg/eRXyzeXEvM" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/1backend/1backend/actions/workflows/backend-tests.yaml/badge.svg"><img src="https://github.com/1backend/1backend/actions/workflows/backend-tests.yaml/badge.svg" alt="backend build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/1backend/1backend/actions/workflows/1backend-ui-docker-build.yaml/badge.svg"><img src="https://github.com/1backend/1backend/actions/workflows/1backend-ui-docker-build.yaml/badge.svg" alt="frontend build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/1backend/1backend/actions/workflows/go-client-build.yaml/badge.svg"><img src="https://github.com/1backend/1backend/actions/workflows/go-client-build.yaml/badge.svg" alt="go client build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/1backend/1backend/actions/workflows/js-client-build.yaml/badge.svg"><img src="https://github.com/1backend/1backend/actions/workflows/js-client-build.yaml/badge.svg" alt="js client build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/1backend/1backend/actions/workflows/go-sdk-build.yaml/badge.svg"><img src="https://github.com/1backend/1backend/actions/workflows/go-sdk-build.yaml/badge.svg" alt="go sdk" style="max-width: 100%;"></a>
    </span>
    <div style="margin-top: 2rem">AI-native microservices platform. v0.8.2.</div>
    <div>
      <a href="https://1backend.com">1backend.com</a>
    </div>
  </div>
</p>

> "Works of art make rules; rules do not make works of art."  
> — Claude Debussy

## Overview

1Backend rethinks software development for the AI age: it lets you build modern distributed apps quickly without infrastructure components.
It's a framework, it's a proxy (edge proxy, reverse proxy and more), handles authorization, microservice calls, microfrontend routing, emails.

It lets you run LLMs through containers and program them. It's zero trust, zero config, comes with its own ORM that can let you test-drive it even without a DB.

## Launching the server

Easiest way to run 1Backend is with Docker. [Install Docker if you don't have it](https://docs.docker.com/engine/install/).
Step into repo root and:

```sh
docker compose up
```

Also see [this page](https://1backend.com/docs/category/running-the-server) for other ways to launch 1Backend.

## Build and use your first service in 3 steps

Check out the [examples](./examples/go/services/) folder or the [relevant documentation](https://1backend.com/docs/writing-custom-services/your-first-service) to learn how to easily build testable, scalable microservices on 1Backend.

Here are the conceptual steps though to illustrate how simple is a 1Backend service:

### Step 1: Register your service's user account

Each service - just like humans - must have their account and manage their own credentials.
Just like humans, they must remember their passwords :)).

#### Go SDK example

```go
import (
  "github.com/1backend/1backend/sdk/go/boot"
  "github.com/1backend/1backend/sdk/go/client"
)

oneBackendClient  := client.NewApiClientFactory(service.Options.ServerUrl).Client()
token, err := boot.RegisterServiceAccount(
		oneBackendClient .UserSvcAPI,
		"your-svc",
		"Your Service",
		credentialStore, // This is just a DB handle
	)
```

### Step 2: Register your service

#### Language-agnostic

Call [`RegisterInstance` endpoint](https://1backend.com/docs/1backend-api/register-instance) directly from any programming language.

#### Go SDK example

```go
import (
  "github.com/1backend/1backend/sdk/go/client"
)

oneBackendClient := client.NewApiClientFactory(service.Options.ServerUrl).Client()

oneBackendClient.RegistrySvcAPI.
		RegisterInstance(context.Background()).
		Body(openapi.RegistrySvcRegisterInstanceRequest{
			Url: "https://your-service-url",
		}).Execute()
```

### Step 3: Call your service through 1Backend

All services registered will be available under their slug at `http(s)://your-1backend-host/{service-slug}/...`.
For more details see doc about the [Proxy Svc](https://1backend.com/docs/built-in-services/proxy-svc).

### CLI

Install `oo` to get started (at the moment you need Go to install it):

```sh
go install github.com/1backend/1backend/cli/oo@latest
```

```sh
$ oo env ls
ENV NAME   SELECTED   URL                           DESCRIPTION   REACHABLE
local      *          http://127.0.0.1:11337                      true
```

```sh
$ oo login 1backend changeme

$ oo whoami
slug: 1backend
id: usr_e9WSQYiJc9
app:
  id: app_hRKWXZzK6P
  host: unnamed
roles:
- user-svc:user
- user-svc:admin
```

```sh
$ oo post /prompt-svc/prompt --sync=true --prompt="Is a cat an animal? Just answer with yes or no please."
# see example response above...
```

## Run On Your Servers

See the [Running the daemon](https://1backend.com/docs/category/running-the-server) page to help you get started.

## Services

For articles about the built-in services see the [Built-in services](https://1backend.com/docs/category/built-in-services) page.
For comprehensive API docs see the [1Backend API](https://1backend.com/docs/category/1backend-api) page.

## Run On Your Laptop/PC

We have temporarily discontinued the distribution of the desktop version. Please refer to this page for alternative methods to run the software.

## License

1Backend is licensed under AGPL-3.0.
