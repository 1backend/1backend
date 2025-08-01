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
    <div style="margin-top: 2rem">AI-native microservices platform. v0.7.6.</div>
    <div>
      <a href="https://1backend.com">1backend.com</a>
    </div>
  </div>
</p>

## Overview

The story of 1Backend began in 2017, but the current incarnation emerged when a seasoned microservices platform architect set out to build an on-premise ChatGPT alternative. What started as a focused AI project quickly grew in complexity—naturally demanding a microservices approach.

With over a decade of experience and ideas in building scalable systems, the author saw an opportunity to merge two passions: modern AI and battle-tested microservice design.

1Backend is the result—a framework and runtime for building distributed web applications from day one. It includes everything you need to get started—running AI models, user management, config, secrets, request routing, proxying, and more—so you don’t need to duct-tape a dozen tools together. And when your needs evolve, it integrates cleanly with standard tools from the broader ecosystem.

<img src="https://singulatron.com/assets/1b.png" />

## Highlights

- On-premise ChatGPT alternative – Run your AI models locally through a UI, CLI or API.
- A "microservices-first" web framework – Think of it like Angular for the backend, built for large, scalable enterprise codebases.
- Out-of-the-box services – Includes built-in file uploads, downloads, user management, and more.
- Infrastructure simplification – Acts as a container orchestrator, reverse proxy, and more.
- Multi-database support – Comes with its own built-in ORM.
- AI integration – Works with LlamaCpp, StableDiffusion, and other AI platforms.

## Starting

Easiest way to run 1Backend is with Docker. [Install Docker if you don't have it](https://docs.docker.com/engine/install/).
Step into repo root and:

```sh
docker compose up
```

to run the platform in foreground. It stops running if you Ctrl+C it. If you want to run it in the background:

```sh
docker compose up -d
```

Also see [this page](https://1backend.com/docs/category/running-the-server) for other ways to launch 1Backend.

## Building microservices

Check out the [examples](./examples/go/services/) folder or the [relevant documentation](https://1backend.com/docs/writing-custom-services/your-first-service) to learn how to easily build testable, scalable microservices on 1Backend.

## Prompting

Now that the 1Backend is running you have a few options to interact with it.

### UI

You can go to `http://127.0.0.1:3901` and log in with username `1backend` and password `changeme` and start using it just like you would use ChatGPT.

Click on the big "AI" button and download a model first. Don't worry, this model will be persisted across restarts (see volumes in the docker-compose.yaml).

### Clients

For brevity the below example assumes you went to the UI and downloaded a model already. (That could also be done in code with the clients but then the code snippet would be longer).

Let's do a sync prompting in JS. In your project run

```sh
npm init -y && jq '. + { "type": "module" }' package.json > temp.json && mv temp.json package.json
npm i -s @1backend/client
```

Make sure your `package.json` contains `"type": "module"`, put the following snippet into `index.js`

```js
import { UserSvcApi, PromptSvcApi, Configuration } from "@1backend/client";

async function testDrive() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    body: {
      slug: "1backend",
      password: "changeme",
    },
  });

  const promptSvc = new PromptSvcApi(
    new Configuration({
      apiKey: loginResponse.token?.token,
    })
  );

  // Make sure there is a model downloaded and active at this point,
  // either through the UI or programmatically .

  let promptRsp = await promptSvc.prompt({
    body: {
      sync: true,
      prompt: "Is a cat an animal? Just answer with yes or no please.",
    },
  });

  console.log(promptRsp);
}

testDrive();
```

and run

```js
$ node index.js
{
  prompt: {
    createdAt: '2025-02-03T16:53:09.883792389Z',
    id: 'prom_emaAv7SlM2',
    prompt: 'Is a cat an animal? Just answer with yes or no please.',
    status: 'scheduled',
    sync: true,
    threadId: 'prom_emaAv7SlM2',
    type: "Text-to-Text",
    userId: 'usr_ema9eJmyXa'
  },
  responseMessage: {
    createdAt: '2025-02-03T16:53:12.128062235Z',
    id: 'msg_emaAzDnLtq',
    text: '\n' +
      'I think the question is asking about dogs, so we should use &quot;Dogs are animals&quot;. But what about cats?',
    threadId: 'prom_emaAv7SlM2'
  }
}
```

Depending on your system it might take a while for the AI to respond.
In case it takes long check the backend logs if it's processing, you should see something like this:

```sh
1backend-1backend-1   | {"time":"2024-11-27T17:27:14.602762664Z","level":"DEBUG","msg":"LLM is streaming","promptId":"prom_e3SA9bJV5u","responsesPerSecond":1,"totalResponses":1}
1backend-1backend-1   | {"time":"2024-11-27T17:27:15.602328634Z","level":"DEBUG","msg":"LLM is streaming","promptId":"prom_e3SA9bJV5u","responsesPerSecond":4,"totalResponses":9}
```

### CLI

Install `oo` to get started (at the moment you need Go to install it):

```sh
go install github.com/1backend/1backend/cli/oo@latest
```

To make the installed binary available on your system:

```sh
$ go env | grep GOPATH
GOPATH='/home/crufter/go' # this will be different for you

# Open your bashrc file with any editor
vim ~/.bashrc

# And place your go bin folder there
export PATH=$PATH:/home/crufter/go/bin

# Please keep in mind your username won't be crufter probably, change it to yours :)
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
roles:
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
