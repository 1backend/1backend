For comprehensive documentation, please visit https://1backend.com/.

# CLI

This CLI is first and foremost aimed at administrators (as opposed to both admins and users like the Web UI) who are connected to multiple environments (the web UI is designed to be single env).

## Quick and Basic Overview

### Envs

#### List

```sh
~/1backend/cli$ go run main.go env list
SELECTED   NAME    URL                                DESCRIPTION
*          local   http://127.0.0.1:58231
           prod    https://api.myprodserver.com
```

### Account

#### Login

```sh
~/1backend/cli$ go run main.go login 1backend changeme
```

#### Whoami

```sh
$ ~/1backend/cli$ go run main.go whoami
1backend
```

### Service Definitions

```sh
~/1backend/cli$ cat fixtures/definitionA.yaml
id: test-a
image:
  name: hashicorp/http-echo
  port: 8080
hostPort: 8887

~/1backend/cli$ cat fixtures/definitionB.yaml
id: test-b
repository:
  url: https://github.com/1backend/1backend.git
  containerFile: server/docker/Dockerfile
  port: 58231
hostPort: 9998

~/1backend/cli$ go run main.go definition save fixtures/definitionA.yaml
~/1backend/cli$ go run main.go definition save fixtures/definitionB.yaml
```

### Deployments

#### List

```sh
~/1backend/cli$ go run main.go deployment list
ID                DEFINITION ID   STATUS   DETAILS
depl_dy2PDIkzqf   test-b          Error    build failed: COPY failed: file not found in build context or excluded by…
depl_dbOdi5eLQK   test-a          OK
```
