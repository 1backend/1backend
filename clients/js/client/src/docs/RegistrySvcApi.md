# RegistrySvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteDefinition**](RegistrySvcApi.md#deletedefinition) | **DELETE** /registry-svc/definition/{id} | Delete Definition |
| [**deleteNode**](RegistrySvcApi.md#deletenode) | **DELETE** /registry-svc/node/{url} | Delete Node |
| [**echoGet**](RegistrySvcApi.md#echoget) | **GET** /registry-svc/echo | Echo the query parameters in the response body. |
| [**echoPost**](RegistrySvcApi.md#echopost) | **POST** /registry-svc/echo | Echo the request body in the response body. |
| [**echoPut**](RegistrySvcApi.md#echoput) | **PUT** /registry-svc/echo | Echo the request body in the response body. |
| [**listDefinitions**](RegistrySvcApi.md#listdefinitions) | **GET** /registry-svc/definitions | List Definitions |
| [**listInstances**](RegistrySvcApi.md#listinstances) | **GET** /registry-svc/instances | List Service Instances |
| [**listNodes**](RegistrySvcApi.md#listnodes) | **POST** /registry-svc/nodes | List Nodes |
| [**registerInstance**](RegistrySvcApi.md#registerinstance) | **PUT** /registry-svc/instance | Register Instance |
| [**removeInstance**](RegistrySvcApi.md#removeinstance) | **DELETE** /registry-svc/instance/{id} | Remove Instance |
| [**saveDefinition**](RegistrySvcApi.md#savedefinition) | **PUT** /registry-svc/definition | Register a Definition |
| [**selfNode**](RegistrySvcApi.md#selfnode) | **GET** /registry-svc/node/self | View Self Node |



## deleteDefinition

> deleteDefinition(id)

Delete Definition

Deletes a registered definition by ID.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { DeleteDefinitionRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // string | Definition ID
    id: id_example,
  } satisfies DeleteDefinitionRequest;

  try {
    const data = await api.deleteDefinition(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | `string` | Definition ID | [Defaults to `undefined`] |

### Return type

`void` (Empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | No Content |  -  |
| **400** | Invalid ID |  -  |
| **401** | Unauthorized |  -  |
| **404** | Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## deleteNode

> deleteNode(url)

Delete Node

Deletes a registered node by node URL. This endpoint is useful when a node is no longer available but it\&#39;s still present in the database.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { DeleteNodeRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // string | Node URL
    url: url_example,
  } satisfies DeleteNodeRequest;

  try {
    const data = await api.deleteNode(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **url** | `string` | Node URL | [Defaults to `undefined`] |

### Return type

`void` (Empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | No Content |  -  |
| **400** | Invalid ID |  -  |
| **401** | Unauthorized |  -  |
| **404** | Service not found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## echoGet

> { [key: string]: any; } echoGet()

Echo the query parameters in the response body.

This endpoint is used to test the server\&#39;s response to a GET request. It echoes back the query parameters as a JSON object.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { EchoGetRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  try {
    const data = await api.echoGet();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

**{ [key: string]: any; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Echoed query parameters |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## echoPost

> { [key: string]: any; } echoPost()

Echo the request body in the response body.

This endpoint is used to test the server\&#39;s response to a request. It simply echoes back the request body as a JSON response.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { EchoPostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  try {
    const data = await api.echoPost();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

**{ [key: string]: any; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Echoed response |  -  |
| **400** | Invalid JSON |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## echoPut

> { [key: string]: any; } echoPut()

Echo the request body in the response body.

This endpoint is used to test the server\&#39;s response to a request. It simply echoes back the request body as a JSON response.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { EchoPutRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  try {
    const data = await api.echoPut();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

**{ [key: string]: any; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Echoed response |  -  |
| **400** | Invalid JSON |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listDefinitions

> RegistrySvcListDefinitionsResponse listDefinitions()

List Definitions

Retrieves a list of all definitions or filters them by specific criteria.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { ListDefinitionsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  try {
    const data = await api.listDefinitions();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**RegistrySvcListDefinitionsResponse**](RegistrySvcListDefinitionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid filters |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listInstances

> RegistrySvcListInstancesResponse listInstances(scheme, ip, deploymentId, host, ip2, id, slug)

List Service Instances

Retrieves a list of all instances or filters them by specific criteria (e.g., host, IP).

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { ListInstancesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // string | Scheme to filter by (optional)
    scheme: scheme_example,
    // string | IP to filter by (optional)
    ip: ip_example,
    // string | Deployment ID to filter by (optional)
    deploymentId: deploymentId_example,
    // string | Host to filter by (optional)
    host: host_example,
    // string | IP to filter by (optional)
    ip2: ip_example,
    // string | Id to filter by (optional)
    id: id_example,
    // string | Slug to filter by (optional)
    slug: slug_example,
  } satisfies ListInstancesRequest;

  try {
    const data = await api.listInstances(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **scheme** | `string` | Scheme to filter by | [Optional] [Defaults to `undefined`] |
| **ip** | `string` | IP to filter by | [Optional] [Defaults to `undefined`] |
| **deploymentId** | `string` | Deployment ID to filter by | [Optional] [Defaults to `undefined`] |
| **host** | `string` | Host to filter by | [Optional] [Defaults to `undefined`] |
| **ip2** | `string` | IP to filter by | [Optional] [Defaults to `undefined`] |
| **id** | `string` | Id to filter by | [Optional] [Defaults to `undefined`] |
| **slug** | `string` | Slug to filter by | [Optional] [Defaults to `undefined`] |

### Return type

[**RegistrySvcListInstancesResponse**](RegistrySvcListInstancesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid filters |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listNodes

> RegistrySvcListNodesResponse listNodes(body)

List Nodes

Retrieve a list of nodes.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { ListNodesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // RegistrySvcListNodesRequest | List Nodes Request (optional)
    body: ...,
  } satisfies ListNodesRequest;

  try {
    const data = await api.listNodes(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [RegistrySvcListNodesRequest](RegistrySvcListNodesRequest.md) | List Nodes Request | [Optional] |

### Return type

[**RegistrySvcListNodesResponse**](RegistrySvcListNodesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## registerInstance

> object registerInstance(body)

Register Instance

Registers an instance. Idempotent.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { RegisterInstanceRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // RegistrySvcRegisterInstanceRequest | Register Instance Request
    body: ...,
  } satisfies RegisterInstanceRequest;

  try {
    const data = await api.registerInstance(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [RegistrySvcRegisterInstanceRequest](RegistrySvcRegisterInstanceRequest.md) | Register Instance Request | |

### Return type

**object**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Created |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## removeInstance

> removeInstance(id)

Remove Instance

Removes a registered instance by ID.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { RemoveInstanceRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // string | Instance ID
    id: id_example,
  } satisfies RemoveInstanceRequest;

  try {
    const data = await api.removeInstance(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | `string` | Instance ID | [Defaults to `undefined`] |

### Return type

`void` (Empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **204** | No Content |  -  |
| **400** | Invalid ID |  -  |
| **401** | Unauthorized |  -  |
| **404** | Service Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveDefinition

> object saveDefinition(body)

Register a Definition

Registers a new definition, associating an definition address with a slug acquired from the bearer token.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { SaveDefinitionRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // RegistrySvcSaveDefinitionRequest | Register Service Definition Request
    body: ...,
  } satisfies SaveDefinitionRequest;

  try {
    const data = await api.saveDefinition(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [RegistrySvcSaveDefinitionRequest](RegistrySvcSaveDefinitionRequest.md) | Register Service Definition Request | |

### Return type

**object**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **201** | Created |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## selfNode

> RegistrySvcNodeSelfResponse selfNode(body)

View Self Node

Show the local node.

### Example

```ts
import {
  Configuration,
  RegistrySvcApi,
} from '';
import type { SelfNodeRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new RegistrySvcApi(config);

  const body = {
    // object | List Registrys Request (optional)
    body: Object,
  } satisfies SelfNodeRequest;

  try {
    const data = await api.selfNode(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | `object` | List Registrys Request | [Optional] |

### Return type

[**RegistrySvcNodeSelfResponse**](RegistrySvcNodeSelfResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

