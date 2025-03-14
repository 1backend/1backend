---
sidebar_position: 60
tags:
  - file-svc
  - file
  - containers
  - services
  - uploads
  - downloads
  - file serve
  - file serving
  - file proxy
  - non-distributed
---

# File Svc

The File Service handles file-related operations, including downloading files from the internet (to cache them for faster access), accepting file uploads, and serving both downloaded and uploaded files.

The File Svc is distributed. Downloads and uploads may reside on any node in the system, any File Svc node will be able to proxy them to you.

> This page provides a high-level overview of `File Svc`. For detailed information, refer to the [File Svc API documentation](/docs/openorch/download-file).

## Responsibilities

- Download and cache files from the internet for faster access.
- Accept and store file uploads.
- Serve files from both cached downloads and uploaded sources.

## Use cases

### Internal file transfer

Upload a file from a local node and retrieve it later using the "serve upload" endpoint.

### Application file uploads

Enable users or services to upload files, such as profile pictures or media attachments.

### Component dependencies

Download external files (e.g., AI models) via URL and provide them to system components as needed.
