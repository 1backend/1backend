# desktop

## Install

### windows

You need Visual Studio installed with "Desktop Development for C++" for nodegyp : /.

Alternatively you can try these or similar maybe:

```sh
choco install visualstudio2022buildtools visualstudio2022-workload-vctools
```

But the best bet is probably downloading VS.

If you get complaints about "distutils":
(https://github.com/nodejs/node-gyp/issues/2869)

```sh
pip install setuptools
```

## Quickstart in the codebase

Backend entrypoint:
```sh
desktop/workspaces/electron-app/main/components/app.ts
```

Frontend entrypoint:
```sh
desktop/worker/workspaces/angular-app/src/app/app.module.ts
```
