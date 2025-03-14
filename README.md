<p align="center">
  <img width="150px" src="./docs-source/openorch_logo.svg" style="border-radius: 50%;" />
  <div align="center">
    <span>
      <h1 style="border-bottom: none">OpenOrch</h1>
      <a href="https://discord.gg/eRXyzeXEvM" rel="nofollow"><img src="https://camo.githubusercontent.com/66351093b042f69e9698398d33f08a6c36f1b7c56e1494b1e2902950eb24c94f/68747470733a2f2f646362616467652e6c696d65732e70696e6b2f6170692f7365727665722f68747470733a2f2f646973636f72642e67672f655258797a655845764d" alt="" data-canonical-src="https://dcbadge.limes.pink/api/server/https://discord.gg/eRXyzeXEvM" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/openorch/openorch/actions/workflows/backend-tests.yaml/badge.svg"><img src="https://github.com/openorch/openorch/actions/workflows/backend-tests.yaml/badge.svg" alt="backend build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/openorch/openorch/actions/workflows/openorch-frontend-docker-build.yaml/badge.svg"><img src="https://github.com/openorch/openorch/actions/workflows/openorch-frontend-docker-build.yaml/badge.svg" alt="frontend build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/openorch/openorch/actions/workflows/go-client-build.yaml/badge.svg"><img src="https://github.com/openorch/openorch/actions/workflows/go-client-build.yaml/badge.svg" alt="go client build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/openorch/openorch/actions/workflows/js-client-build.yaml/badge.svg"><img src="https://github.com/openorch/openorch/actions/workflows/js-client-build.yaml/badge.svg" alt="js client build" style="max-width: 100%;"></a>
<a target="_blank" rel="noopener noreferrer" href="https://github.com/openorch/openorch/actions/workflows/go-sdk-build.yaml/badge.svg"><img src="https://github.com/openorch/openorch/actions/workflows/go-sdk-build.yaml/badge.svg" alt="go sdk" style="max-width: 100%;"></a>
    </span>
    <div style="margin-top: 2rem">Build AI products faster. A language-agnostic microservices platform for building AI applications.</div>
    <div>
      <a href="https://openorch.org">openorch.org</a>
    </div>
  </div>
</p>

OpenOrch was initially created to solve the challenge of running AI models on private servers, handling many concurrent prompts from both users and services. The goal was to build a ChatGPT-like interface for humans and a network-accessible API for machines.

As the system grew, the authors—despite 10+ years of building both closed and open-source microservices platforms—realized there was still no backend framework that met their needs. So, OpenOrch evolved into the flexible, scalable microservices platform they had been searching for.

## Highlights

- On-premise ChatGPT alternative – Run your AI models locally through a UI, CLI or API.
- A "microservices-first" web framework – Think of it like Angular for the backend, built for large, scalable enterprise codebases.
- Out-of-the-box services – Includes built-in file uploads, downloads, user management, and more.
- Infrastructure simplification – Acts as a container orchestrator, reverse proxy, and more.
- Multi-database support – Comes with its own built-in ORM.
- AI integration – Works with LlamaCpp, StableDiffusion, and other AI platforms.

## Starting

Easiest way to run OpenOrch is with Docker. [Install Docker if you don't have it](https://docs.docker.com/engine/install/).
Step into repo root and:

```sh
docker compose up
```

to run the platform in foreground. It stops running if you Ctrl+C it. If you want to run it in the background:

```sh
docker compose up -d
```

## Using

Now that the OpenOrch is running you have a few options to interact with it.

### UI

You can go to `http://127.0.0.1:3901` and log in with username `openorch` and password `changeme` and start using it just like you would use ChatGPT.

Click on the big "AI" button and download a model first. Don't worry, this model will be persisted across restarts (see volumes in the docker-compose.yaml).

### Clients

For brevity the below example assumes you went to the UI and downloaded a model already. (That could also be done in code with the clients but then the code snippet would be longer).

Let's do a sync prompting in JS. In your project run

```sh
npm init -y && jq '. + { "type": "module" }' package.json > temp.json && mv temp.json package.json
npm i -s @openorch/client
```

Make sure your `package.json` contains `"type": "module"`, put the following snippet into `index.js`

```js
import { UserSvcApi, PromptSvcApi, Configuration } from "@openorch/client";

async function testDrive() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    body: {
      slug: "openorch",
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
openorch-backend-1   | {"time":"2024-11-27T17:27:14.602762664Z","level":"DEBUG","msg":"LLM is streaming","promptId":"prom_e3SA9bJV5u","responsesPerSecond":1,"totalResponses":1}
openorch-backend-1   | {"time":"2024-11-27T17:27:15.602328634Z","level":"DEBUG","msg":"LLM is streaming","promptId":"prom_e3SA9bJV5u","responsesPerSecond":4,"totalResponses":9}
```

## CLI

Install `oo` to get started (at the moment you need Go to install it):

```sh
go install github.com/openorch/openorch/cli/oo@latest
```

```sh
$ oo env add local http://127.0.0.1:58231

$ oo env ls
ENV NAME   SELECTED   URL                           DESCRIPTION
local      *          http://127.0.0.1:58231
```

```sh
$ oo login openorch changeme

$ oo whoami
slug: openorch
id: usr_e9WSQYiJc9
roles:
- user-svc:admin
```

```sh
$ oo post /prompt-svc/prompt --sync=true --prompt="Is a cat an animal? Just answer with yes or no please."
# see example response above...
```

## Context

OpenOrch is a microservices-based AI platform, the seeds of which began taking shape in 2013 while I was at Hailo, an Uber competitor. The idea stuck with me and kept evolving over the years – including during my time at [Micro](https://github.com/micro/micro), a microservices platform company. I assumed someone else would eventually build it, but with the AI boom and the wave of AI apps we’re rolling out, I’ve realized it’s time to build it myself.

## Run On Your Servers

See the [Running the daemon](https://openorch.org/docs/category/running-the-server) page to help you get started.

## Services

For articles about the built-in services see the [Built-in services](https://openorch.org/docs/category/built-in-services) page.
For comprehensive API docs see the [OpenOrch API](https://openorch.org/docs/category/openorch-api) page.

## Run On Your Laptop/PC

We have temporarily discontinued the distribution of the desktop version. Please refer to this page for alternative methods to run the software.

## License

OpenOrch is licensed under AGPL-3.0.
