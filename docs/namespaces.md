### [&#8592; Back](README.md)

# Namespaces

Namespaces enable a service to shard their database to support multiple services building on top of it.

Think about the use case of writing a user service that enables others to register their users into it.
Let's say the service is in question is `johnsmith/users`, and your database table looks like this:

```sql
email               | password
------------------------------
joe@gmail.com       | pwhash1
kate@gmail.com      | pwhash2
ann@gmail.com       | pwhash3
```

It's all cool and nice until you realise you can't differentiate between the users of `news.ycombinator.com` and `rt.com`. Ideally you would have something like this:

```sql
namespace  | email               | password
-------------------------------------------
ycomb      | joe@gmail.com       | pwhash1
ycomb      | kate@gmail.com      | pwhash2
ycomb      | ann@gmail.com       | pwhash3
rt         | vlad@mail.ru        | pwhash4
```

And this is the reason we need the concept of `namespaces`.

### Namespace resolution

The only problem remains is how do we assign namespaces to service calls - after all we can't let the caller tell simply which namspace it wants to access, because it could easily lie and get, for example, the users of an other application.

The concept of secret `caller ids` were created to prevent this scenario of one service impersonating an other one:

Each service has a unique and secret caller id, and when it makes a request to an other service it passes its caller id in the request header. The proxy then translates this id to the service name and passes the translater caller name to the called service (also in the http header):

```
                                                    caller-id: x6s70df
|-----------| caller-id: x6s70df |----------------| caller-namespace: rt:default   |-----------------|
| rt/web    | -----------------> | 1Backend proxy | -----------------------------> | johnsmith/users |
|-----------|                    |----------------|                                |-----------------|
```

The respective header names are `caller-id` and `caller-namespace`.

As you can see, each project has a namespace assigned to it. By default it is $AUTHOR:default.

#### A chain of calls

You might wonder what happens when the call chain is longer than 1 (is service A calls service B).
It is entirely up to a service to pass on either its own caller id or the caller id of its caller.

The services you are calling can potentially mishandle the caller id - just like they can mishandle any of your other data. As a general rule of thumb, any service you is assumed trustworthy.

####

If you are really worried about a service seeing your data we advise you to use forking (not implemented yet).

#### The source of the caller id

As we mentioned, the secret caller id is unique to each service. Every service container has access to it its own caller id - it is in the environment variable named `CALLER_ID`.

### [Continue with reading about types &#8594;](types.md)
