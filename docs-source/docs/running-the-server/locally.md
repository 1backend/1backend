---
sidebar_position: 1
tags:
  - run
  - deploy
  - local
---

# Running the OpenOrch server locally

The easiest way to run OpenOrch is to use [Docker Compose](https://docs.docker.com/compose/install/).

## Using Docker Compose

The easiest way to run this is to clone [the repo](https://github.com/openorch/openorch), step into the repo root and run:

```sh
git clone git@github.com:openorch/openorch.git
cd openorch
docker compose up
# or use the -d flag to run it in the background
# docker compose up -d
```

The `docker-compose-yaml` in the root folder is designed to build and run the current code. For a more production ready Docker Compose file see the [Running the OpenOrch server with Docker Compose and prebuilt images](./docker-compose/).

### Once it's running in Docker Compose

After the containers successfully start, you can go to [http://127.0.0.1:3901](http://127.0.0.1:3901) and log in with the [Default Credentials](/docs/running-the-server/using#default-credentials).

## Running natively (Go & Angular)

If you have both Go and Angular installed on your computer, the easiest way to dip your feet into OpenOrch is to run things locally.

## Running the backend natively (with Go)

```bash
cd server;
go run main.go
```

## Running the frontend natively (with Angular)

```bash
cd desktop/workspaces/angular-app/;
npm run start
```

### Once it's running on the host

After the both the backend and frontend starts, you can go to [http://127.0.0.1:4200](http://127.0.0.1:4200) and log in with the [Default Credentials](/docs/running-the-server/using#default-credentials).

## Administration

### Local files

By default OpenOrch uses the folder `~/.openorch` on your machine for data tables, file downloads, file uploads.
The `~/.openorch/cliConfig.yaml` file is where the [`oo CLI`](/docs/command-line/basics) stores all its data.

#### Download & Uploads

Downloads and uploads are managed by the [File Svc](/docs/built-in-services/file-svc), and by default are stored here:

```bash
~/.openorch/downloads
~/.openorch/uploads
```

#### Data files

By default OpenOrch uses local gzipped json files to store database entries. Data access across OpenOrch is interface based so the this implementation can be easily swapped out for PostgreSQL and other database backends.

These files are located at

```bash
ls ~/.openorch/data
```

Each file is prefixed by the owner service slug, so the `User Svc` `users` table becomes `userSvcUsers`.

If you want to view the contents of a file:

```bash
cat ~/.openorch/data/userSvcUsers.zip | gzip -dc

# or if you jave jq installed
cat ~/.openorch/data/userSvcUsers.zip | gzip -dc | jq
```
