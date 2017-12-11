### [&#8592; Back](README.md)

# Services

Services are a collection of lambda functions exposed to the outside world as
HTTP endpoints. Think of services as backend code running somewhere.

That awesome monolithic app you are working on at yor company? The backend of it can be a service on 1Backend, or it can be many - if you prefer to break it down microservices style.

Service are namespaced by their authors - the user who creates them - just like GitHub repos are.

### How to call a service

Every service's every endpoint is a HTTP endpoint.
You can call them with any HTTP tool (eg. cURL), and 1Backend also generates type safe clients in different languages if you specify the [types](types.md) of your service and endpoints.

To read more about the generate type safe clients, go [here](types.md).

If you prefer to do things the hard/low level way, here is how you can make a curl to a fictional service called "service" of the user called "user". The endpoint we are going to call is called "endpoint":

```sh
curl -XGET 'https://1backend.com:9993/app/user/service/endpoint' -H 'token: 320b5933-e8f6-4daf-ad16-fb36ede68233'
```

(What is that token thingie? Read more [here](tokens.md))

Currently both the input and the output is expected to be JSON on 1Backend. This is not enforced in any way, but the generated clients expect it. So we advise you to make your services follow this schema:

#### Inputs

To pass in inputs to services that expect them, just do what you usually do with json speaking services:

* Send form parameters to `GET` and `DELETE` requests (ie. no JSON here).
* Send a marshalled json body to `POST` and `PUT` requests.

#### Output

Outputs are uniformly just marshalled JSON in the response body.
(Support for XML and other formats might come later.)

### Service privacy

#### Private vs public

Private services are unlisted, only you can see them. Public services have three levels of openness:

#### Closed source, open source, open data

* Closed source services only publish their types and API, so others can consume
  them.
* Open source services publish their source code.
* Open data services publish their source and their data.

By default every service is open source, but not open data (for security reasons).
