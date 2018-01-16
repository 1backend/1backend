### [&#8592; Back](README.md)

# Services

Services are collections of lambda functions exposed to the outside world as
HTTP endpoints. They are nothing fancy - it's just backend code running somewhere.

That awesome monolithic app you are working on at your company? Its backend can be a service on 1Backend, or there can be many - if you prefer to break it down microservices-style.

Service are namespaced by their authors - the user who creates them - just like GitHub repos are.

### Calling services

Every service's every endpoint (function) is a HTTP endpoint. HTTP is used for both `client -> service` and `service -> service` communication.

#### Calling a service from the outside world

```
            internet                              1Backend network
/------------------------------\ /-------------------------------------------------\

              client request                        client request
              to service A                          to service A
(           ) -----------------> |----------------| -----------------> |-----------|
(  client   )                    | 1Backend proxy |                    | service A |
(           ) <----------------- |----------------| <----------------- |-----------|
              service A response                    service A response
              to client                             to client
```

#### Calling a service from another service

Services do not call each other directly. Each request goes through the same 1Backend proxy.
This is because the proxy is a perfect place to implement a lot of things we would otherwise have to do in clients. That would mean reimplementing things in a lot of languages, and also - you can't really trust clients.

```
                                 1Backend network
/-----------------------------------------------------------------------------------\

              service A request                     service A request
              to service B                          to service B
|-----------| -----------------> |----------------| -----------------> |-----------|
| service A |                    | 1Backend proxy |                    | service B |
|-----------| <----------------- |----------------| <----------------- |-----------|
              service B response                    service B response
              to service A                          to service A
                                         ^
                                         |
                                         |---- place of instrumentation and other magic
```

You can call service endpoints with any HTTP tool (e.g. cURL), and 1Backend also generates type-safe clients in different languages if you specify the [types](types.md) of your service and endpoints.

To read more about the generated type-safe clients, go [here](types.md).

If you prefer to do things the hard/low-level way, here is how you can make a curl request to a fictional service called "service" of the user called "user". The endpoint we are going to call is called "endpoint":

```sh
curl 'https://1backend.com:9993/app/user/service/endpoint' -H 'token: 320b5933-e8f6-4daf-ad16-fb36ede68233'
```

(What is that token thingie? Read more [here](tokens.md))

Currently, both the input and the output are expected to be JSON on 1Backend. This is not enforced in any way, but the generated clients expect it. So we advise you to make your services follow this schema:

#### Inputs

To pass in inputs to services that expect them, just do what you usually do with JSON-speaking services:

* Send form parameters to `GET` and `DELETE` requests (i.e. no JSON here).
* Send a marshalled JSON body to `POST` and `PUT` requests.

#### Output

Outputs are uniformly just marshalled JSON in the response body.
(Support for XML and other formats might come later.)

### Service privacy

#### Private vs public

Private services are unlisted; only you can see them. Public services have three levels of openness:

#### Closed source, open source, open data

* Closed source services only publish their types and API, so others can consume
  them.
* Open source services publish their source code.
* Open data services publish their source and their data.

By default, every service is open source, but not open data.

### [Continue with reading about namespaces &#8594;](README.md)
