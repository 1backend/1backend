---
sidebar_position: 4
tags:
  - test
---

# System design guidelines

1Backend doesn't force you to use any particular patterns, languages, stacks, databases, or even conventions. 1Backend services themselves however follow some light conventions.

## API naming guidelines

### List\* endpoints

It is advised that endpoints returning entities are called `ListEntities`, eg. `ListMessages`. Do not create read/query endpoints that return a single entity unless you have good reasons to do so.

The list endpoints should have a few standard filters ideally:

#### `ids` field

The field `ids` should enable single and multiread by ID queries in list endpoints.

#### Pagination

Pagination should happen with these fields in the top level of the List request:

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

#### createdAt & updatedAt ordering

Extracting data from microservices into external systems (like BigQuery) must respect service boundaries. Since each service owns its data, the only supported way to access it is through its API—making efficient pagination essential.

Traditional offset-based pagination can be inefficient in distributed systems, so 1Backend services use cursor-based pagination via the afterTime field.

- Pagination using createdAt allows clients to perform a full scan of all records. This is typically done in ascending order, starting from the earliest entry.

- Pagination using updatedAt is ideal for retrieving recently modified records. When sorted in descending order, it helps clients fetch the latest updates efficiently.

## Save\* endpoints

For endpoints that create or update entities, it’s recommended to use a unified SaveEntities naming convention (e.g., SaveMessages).

While separating Create and Update operations allows for more precise request schemas (since the required fields may differ), 1Backend favors the following principles:

- **Minimize the number of endpoints** — simpler APIs are easier to maintain and evolve.

- **Prioritize idempotency** — Create endpoints typically aren't idempotent, while Save can be designed to be.

- **Optimize for batch operations** — network calls are costly. Design endpoints to handle multiple entities at once. This not only improves efficiency but also makes it easier to scale or create specialized multi-\* versions of your endpoint when performance demands arise.

### IDs

Ids are by convention prefixed by a shorthand inspired by the entity name, think thread IDs being prefixed by `thr_`.
