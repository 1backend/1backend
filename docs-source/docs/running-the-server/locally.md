---
sidebar_position: 1
tags:
  - run
  - deploy
  - local
---

# Running the 1Backend server locally

The easiest way to run 1Backend is to use [Docker Compose](https://docs.docker.com/compose/install/).

## Using Docker Compose

The easiest way to run this is to clone [the repo](https://github.com/1backend/1backend), step into the repo root and run:

```sh
git clone git@github.com:1backend/1backend.git
cd 1backend
docker compose up
# or use the -d flag to run it in the background
# docker compose up -d
```

The `docker-compose-yaml` in the root folder is designed to build and run the current code. For a more production ready Docker Compose file see the [Running the 1Backend server with Docker Compose and prebuilt images](./docker-compose/).

### Once it's running in Docker Compose

After the containers successfully start, you can go to [http://127.0.0.1:3901](http://127.0.0.1:3901) and log in with the [Default Credentials](/docs/running-the-server/using#default-credentials).

## Running natively (Go & Angular)

If you have both Go and Angular installed on your computer, the easiest way to dip your feet into 1Backend is to run things locally.

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

By default 1Backend uses the folder `~/.1backend` on your machine for data tables, file downloads, file uploads.
The `~/.1backend/cliConfig.yaml` file is where the [`oo CLI`](/docs/command-line/basics) stores all its data.

#### Download & Uploads

Downloads and uploads are managed by the [File Svc](/docs/built-in-services/file-svc), and by default are stored here:

```bash
~/.1backend/downloads
~/.1backend/uploads
```

#### Data files

By default 1Backend uses local gzipped json files to store database entries. Data access across 1Backend is interface based so the this implementation can be easily swapped out for PostgreSQL and other database backends.

These files are located at

```bash
ls ~/.1backend/data
```

Each file is prefixed by the owner service slug, so the `User Svc` `users` table becomes `userSvcUsers`.

If you want to view the contents of a file:

```bash
cat ~/.1backend/data/userSvcUsers.zip | gzip -dc

# or if you jave jq installed
cat ~/.1backend/data/userSvcUsers.zip | gzip -dc | jq
```
