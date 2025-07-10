---
sidebar_position: 2
tags:
  - cli
---

# Command line basics

## CLI installation

Currently you need Go to install the 1Backend CLI:

```sh
go install github.com/1backend/1backend/cli/oo@latest
```

## CLI usage

Assuming the daemon is running already (see [this section about that](/docs/running-the-server/using#default-credentials)), you can interact with it through the CLI.

Let's add the local environment first:

```sh
$ oo env add local http://127.0.0.1:11337
$ oo env ls
SELECTED   NAME    URL                           DESCRIPTION
*          local   http://127.0.0.1:11337
```

After this you you need to log in:

```sh
$ oo login 1backend
Enter password:
```

```sh
$ oo whoami
slug: 1backend
id: usr_eH9mXKgmb0
roles:
- user-svc:admin
```

Let's make a GET call:

```sh
$ oo get /secret-svc/is-secure
{
  "isSecure":false
}
```

Or a POST call:

```sh
$ oo post /user-svc/users
{
  "users": [
   {
      "id": "usr_g5WbJXmxuQ",
      "createdAt": "2025-06-05T16:58:15.674576217+02:00",
      "updatedAt": "2025-06-05T16:58:15.674576217+02:00",
      "name": "Admin",
      "slug": "1backend"
    }
  ],
  # Some other fields might be included such as for
  # pagination or count.
}
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
$ oo post /registry-svc/echo --value=hey
```

Is roughly equivalent to the pseudocurl

```sh
curl -XPOST -H "Auth..." $ADDR/registry-svc/echo -d '{"value": "hey"}'
```

Similarly, dot `.` and dash `-` delimiters get turned into a multidimensional JSON:

```sh
$ oo post /registry-svc/echo --value-text=hey
# turns into
{
  "value": {
    "text": "hey"
  }
}
```

```sh
$ oo post /registry-svc/echo --value.text=hey
# turns into
{
  "value": {
    "text": "hey"
  }
}
```

```sh
$ oo post /registry-svc/echo --valueText=hey
# turns into
{
  "valueText": "hey"
}
```
