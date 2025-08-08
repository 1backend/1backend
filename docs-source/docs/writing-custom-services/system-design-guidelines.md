---
sidebar_position: 4
tags:
  - test
---

# System design guidelines

1Backend doesn't force you to use any particular patterns, languages, stacks, databases, or even conventions. 1Backend services themselves however follow some light conventions.

## Pagination

### Offset vs. cursor pagination in distributed systems

In distributed systems—especially when dealing with large datasets or high write volumes—**OFFSET-based pagination** is inefficient and unstable because:

- **Performance:** The database still has to scan through all skipped rows before returning results.
- **Consistency:** New inserts or deletes between requests can cause missing or duplicated results.

---

### Cursor-Based Pagination

Cursor-based pagination is the preferred strategy in distributed systems because it avoids the problems of offset-based pagination — such as inconsistent results when data changes between pages, and performance issues with large offsets.

Instead of skipping rows, cursor-based pagination uses a **cursor value** to mark the position in the dataset and fetch the next page relative to that position.

---

#### How It Works

1. The client receives a **cursor** from the API in each paginated response.
2. The client includes that cursor in the next request to fetch the next page.
3. The server uses the cursor to determine where to resume from in the dataset.

---

#### Cursor Format

The cursor is typically represented as a slice of values (`[]any`), such as:

```json
["2023-08-01T12:00:00Z", "user-123"]
```

This allows precise pagination even when multiple records share the same timestamp (e.g., `createdAt`). The second element (like a unique ID) ensures deterministic ordering and avoids skipping or duplicating records.

#### Example

**Request 1**:

```http
GET /items?limit=10
```

**Response 1**:

```json
{
  "items": [...],
  // This is actually optional, one can easily take the fields from the last item in the list.
  "nextCursor": ["2023-08-01T12:00:00Z", "user-123"]
}
```

```http
GET /items?limit=10&after=["2023-08-01T12:00:00Z","user-123"]
```

### Field Name Conventions

Pagination parameters should be top-level fields in a `List` request. For example:

```json
{
  "limit": 10,
  "after": ["2023-08-01T12:00:00Z", "user-123"],
  "order": "asc",
  "orderBy": "createdAt"
}
```

### Common Ordering Strategies

- **`createdAt` (ascending)**:  
  Use when exporting or scanning all records from the beginning. This gives a complete, chronological view of the dataset.

- **`updatedAt` (descending)**:  
  Use when syncing recent changes (e.g., polling for updates). Returns the newest records first.

---

### Why Timestamps Alone Are Not Enough

Timestamps like `createdAt` or `updatedAt` are commonly used for ordering, but they are not unique. Multiple records can share the same timestamp, leading to unstable or overlapping pagination.

To prevent this, always pair the timestamp with a **unique, stable identifier** (like an ID). This ensures:

- Deterministic ordering
- No records are skipped or repeated
- Proper continuation across pages, even when timestamps collide

No special parsing is required — the cursor is passed as a JSON array and can be unmarshaled directly as `[]any` in Go.

## API naming guidelines

### List\* endpoints

It is advised that endpoints returning entities are called `ListEntities`, eg. `ListMessages`. Do not create read/query endpoints that return a single entity unless you have good reasons to do so.

The list endpoints should have a few standard filters ideally:

#### `ids` field

The field `ids` should enable single and multiread by ID queries in list endpoints.

### Save\* endpoints

For endpoints that create or update entities, it’s recommended to use a unified SaveEntities naming convention (e.g., SaveMessages).

While separating Create and Update operations allows for more precise request schemas (since the required fields may differ), 1Backend favors the following principles:

- **Minimize the number of endpoints** — simpler APIs are easier to maintain and evolve.

- **Prioritize idempotency** — Create endpoints typically aren't idempotent, while Save can be designed to be.

- **Optimize for batch operations** — network calls are costly. Design endpoints to handle multiple entities at once. This not only improves efficiency but also makes it easier to scale or create specialized multi-\* versions of your endpoint when performance demands arise.

### IDs

Ids are by convention prefixed by a shorthand inspired by the entity name, think thread IDs being prefixed by `thr_`.
