### [&#8592; Back](README.md)

# Types

1Backend places a lot of focus on enabling the interoperability of different
languages and tech stacks, since one of the biggest benefits of microservices is
that you can use the best tool for a given job.

To achieve this, there is strong support for API definitions:

Every [service](services.md) has a "type" section where you can define the types
living in your service. An example would be:

```json
{
  "user": [{ "id": "string" }, { "name": "string" }],
  "product": [
    { "id": "string" },
    { "price": "float" },
    { "onSale": "bool" },
    { "inStock": "int" }
  ]
}
```

And each endpoint can refer to these types in their "input" and "output"
sections.

An input section of a fictional "/products/on-sale" might look like:

```json
[{ "priceMin": "float" }, { "priceMax": "float" }]
```

With the following output definition:

```
product[]
```

This will generate an Angular service method with the signature:

```typescript
getProductsOnSale(priceMin: number, priceMax: number): Product[]
```

### Unnamed types

The following built-in types are supported:

```
bool;
bool[];
string;
string[];
int;
int[];
float;
float[];
```

### Named types

Named types are a list of field name -> type pairs. The type of a field can be another
named type. See the "owner" field in the following example:

```json
{
  "user": [{ "id": "string" }, { "name": "string" }],
  "product": [
    { "id": "string" },
    { "price": "float" },
    { "onSale": "bool" },
    { "inStock": "int" },
    { "owner": "User" }
  ]
}
```

### Referring to named types from other services

Since services are designed to talk to each other, the picture would not be
complete without a way to refer to types in other services:

```json
{
  "product": [
    { "id": "string" },
    { "price": "float" },
    { "onSale": "bool" },
    { "inStock": "int" },
    { "owner": "crufter/user-service.User" }
  ]
}
```

As you can see, the syntax is "[User]/[ServiceName].[Type]".

### List types

All types can be made into lists by appending `[]` to their names:

```
 { "owners": "crufter/user-service.User[]" }
```

### Codebase pointers

You can find the code related to client generation in the [api-pack](../backend/api-pack) directory.
