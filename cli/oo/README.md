For comprehensive documentation, please visit https://1backend.com/.

# CLI

This CLI is first and foremost aimed at administrators (as opposed to both admins and users like the Web UI) who are connected to multiple environments (the web UI is designed to be single env).

## Quick and Basic Overview

### Envs

#### List

```sh
~/1backend/cli$ go run main.go env list
SELECTED   NAME    URL                                DESCRIPTION
*          local   http://127.0.0.1:11337
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
