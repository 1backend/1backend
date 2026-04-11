# File Service (short overview)

This service stores **uploads** and **downloads** on node-local disk, with optional cloud backing (GCS).

## High-level flow

```text
Client
  |
  v
file-svc API
  |
  +--> metadata store (uploads/downloads records)
  |
  +--> local disk
  |
  '--> optional cloud storage (GCS)
```

## Upload path

```text
PUT /file-svc/upload
      |
      v
Generate FilePath (canonical)
      |
      +--> write local file at FilePath
      +--> if cloud enabled: persist object using same FilePath
      '--> store Upload metadata (includes FilePath)
```

## Download path

```text
PUT /file-svc/download
      |
      v
Generate download FilePath + storage key (sharded)
      |
      +--> fetch from origin URL -> write local file
      +--> if cloud enabled: persist NEW download object to GCS
      '--> store Download metadata (includes FilePath)
```

## Serve downloaded file path

```text
GET /file-svc/serve/download/{url}
      |
      v
Find local replica metadata
      |
      +--> if local file exists: serve local file (no cloud restore)
      '--> if local file missing: delete stale record and redownload from origin
```

## Storage key layout

Downloads are written with sharded keys to avoid huge flat directories:

```text
downloads/<2>/<2>/<full-hash>
```

- `downloads/` is the cloud prefix.
- `<2>/<2>/<full-hash>` comes from `DownloadStorageFilePath(url)`.
