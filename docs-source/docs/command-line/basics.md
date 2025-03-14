---
sidebar_position: 2
tags:
  - cli
---

# Command line basics

## CLI installation

At the moment you need Go to install the OpenOrch CLI:

```sh
go install github.com/openorch/openorch/cli/oo@latest
```

## CLI usage

Assuming the daemon is running already (see [this section about that](/docs/running-the-server/using#default-credentials)), you can interact with it through the CLI.

Let's add the local environment first:

```sh
$ oo env add local http://127.0.0.1:58231
$ oo env ls
SELECTED   NAME    URL                           DESCRIPTION
*          local   http://127.0.0.1:58231
```

After this you you need to log in:

```sh
$ oo login openorch
Enter password:
```

```sh
$ oo whoami
slug: openorch
id: usr_eH9mXKgmb0
roles:
- user-svc:admin
```

Let's make a GET call:

```sh
$ oo get /config-svc/config
{"config":{"namespace":"","data":{"config-svc":{"directory":"/root"},"download-svc":{"downloadFolder":"/root/downloads"},"model-svc":{"currentModelId":"huggingface/TheBloke/mistral-7b-instruct-v0.2.Q3_K_S.gguf"}}}}
```

Or a POST call:

```sh
$ oo post /user-svc/users
{
  "users": [

    {
      "id": "usr_e9WSYwjRuL",
      "createdAt": "2024-12-06T20:51:38.062985Z",
      "updatedAt": "2024-12-06T20:51:38.062985Z",
      "name": "Proxy Service",
      "slug": "proxy-svc",
      "contacts": [
        {
          "createdAt": "0001-01-01T00:00:00Z",
          "updatedAt": "0001-01-01T00:00:00Z"
        }
      ]
    },
...
```

Or a POST call with some request body parameters:

```sh
$ oo post /secret-svc/encrypt --value=hey
{
  "value": "UsoGq6VCa0+89pzIPhgU49kgoL0p/3jc90IsOR/8ldk="
}
```

Here we should talk a bit about how CLI flags get mapped to request bodies.

## CLI flag to request body mapping

When doing `POST`, `PUT` and `DELETE` queries, CLI flags can be turned into multilevel JSON request bodies, such as this:

```sh
$ oo post /secret-svc/encrypt --value=hey
```

Is roughly equivalent to the pseudocurl

```sh
curl -XPOST -H "Auth..." $ADDR/secret-svc/encrypt -d '{"value": "hey"}'
```

Similarly, dot `.` and dash `-` delimiters get turned into a multidimensional JSON:

```sh
$ oo post /secret-svc/encrypt --value-text=hey
# turns into
{
  "value": {
    "text": "hey"
  }
}
```

```sh
$ oo post /secret-svc/encrypt --value.text=hey
# turns into
{
  "value": {
    "text": "hey"
  }
}
```

```sh
$ oo post /secret-svc/encrypt --valueText=hey
# turns into
{
  "valueText": "hey"
}
```
