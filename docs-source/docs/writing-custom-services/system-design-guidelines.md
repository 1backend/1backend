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

### Cursor-based pagination

Cursor-based pagination solves these problems by:

1. Using a **cursor** (often an encoded value, like a timestamp or an ID) that marks your position in the result set.
2. Querying for results **"after"** that cursor instead of skipping rows.
3. Returning a new cursor in each response for the client to use for the next page.

---

#### Example

**Request 1:**

```http
GET /items?limit=10
```

**Response 1:**

```json
{
  "items": [...],
  "next_cursor": "2023-08-01T12:00:00Z"
}
```

**Request 2:**

```http
GET /items?limit=10&after=2023-08-01T12:00:00Z
```

**Common names for cursor-bbased pagination**:

- Cursor pagination (common in APIs like GraphQL and Twitter API)
- Keyset pagination (common in SQL performance discussions)
- Sometimes just "after-based pagination" (naming it after the query parameter)

### Field name conventions

Pagination should happen with these fields in the top level of a `List` request:

```json
{
  "limit": 10,
  "__comment_limit": "Limit is the maximum number of users to return.",

  "afterTime": "2023-01-01T00:00:00Z",
  "__comment_afterTime": "AfterTime is a time in RFC3339 format. It is used to paginate the results when the `orderBy` is set to `createdAt` or `updatedAt`. The results will be returned after this time.",

  "order": "desc",
  "orderBy": "createdAt",

  "count": false,
  "__comment_count": "Count is a flag that indicates if the count of the users should be returned."
}
```

### `createdAt` & `updatedAt` ordering

When exporting data from microservices into external systems (such as BigQuery), it’s important to respect **service boundaries**.  
Each service owns its own data, and the only supported way to access it is through that service’s API.  
This makes **efficient pagination** critical for both performance and reliability.

Offset-based pagination (e.g., `offset=100&limit=50`) can be slow and unreliable in distributed systems, especially as datasets grow or records change during retrieval.  
Instead, 1Backend services use **cursor-based pagination**, where the client provides a “bookmark” that tells the API where to resume.  
For time-based pagination, this bookmark is expressed through the `afterTime` field.

### Common ordering strategies

- **Pagination by `createdAt`**  
  Use when you need to retrieve _all_ records from the beginning, such as during a full export.  
  Typically ordered **ascending**, starting with the oldest entries and moving forward in time.

- **Pagination by `updatedAt`**  
  Use when you need to retrieve _recently changed_ records, such as for an incremental sync.  
  Typically ordered **descending**, returning the most recent changes first.

---

### `after` and `afterTime` Fields

In theory, a pagination cursor could be named simply `after` and accept any type of value (or an array of values).  
However, when the cursor represents a timestamp, using a generic `after` field would require extra parsing logic to convert it into a `time.Time` object in Go.

To avoid this and allow **automatic unmarshaling** into `time.Time`, the convention for time-based pagination is to name the field **`afterTime`**.  
This makes the expected format explicit (RFC3339 timestamp) and removes the need for custom parsing.

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
