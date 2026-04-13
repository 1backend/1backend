# DataSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createObject**](DataSvcApi.md#createobject) | **POST** /data-svc/object | Create a Generic Object |
| [**deleteObjects**](DataSvcApi.md#deleteobjects) | **POST** /data-svc/objects/delete | Delete Objects |
| [**queryObjects**](DataSvcApi.md#queryobjects) | **POST** /data-svc/objects | Query Objects |
| [**updateObjects**](DataSvcApi.md#updateobjects) | **POST** /data-svc/objects/update | Update Objects |
| [**upsertObject**](DataSvcApi.md#upsertobject) | **PUT** /data-svc/object/{objectId} | Upsert a Generic Object |
| [**upsertObjects**](DataSvcApi.md#upsertobjects) | **PUT** /data-svc/objects/upsert | Upsert Objects |



## createObject

> DataSvcCreateObjectResponse createObject(body)

Create a Generic Object

Creates a new object with the provided details. Requires authorization and user authentication.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { CreateObjectRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // DataSvcCreateObjectRequest | Create request payload
    body: ...,
  } satisfies CreateObjectRequest;

  try {
    const data = await api.createObject(body);
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
| **body** | [DataSvcCreateObjectRequest](DataSvcCreateObjectRequest.md) | Create request payload | |

### Return type

[**DataSvcCreateObjectResponse**](DataSvcCreateObjectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## deleteObjects

> object deleteObjects(body)

Delete Objects

Deletes all objects matchin the provided filters.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { DeleteObjectsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // DataSvcDeleteObjectRequest | Delete request payload
    body: ...,
  } satisfies DeleteObjectsRequest;

  try {
    const data = await api.deleteObjects(body);
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
| **body** | [DataSvcDeleteObjectRequest](DataSvcDeleteObjectRequest.md) | Delete request payload | |

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
| **200** | Successful deletion of object |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## queryObjects

> DataSvcQueryResponse queryObjects(body)

Query Objects

Retrieves objects from a specified table based on search criteria. Requires authorization and user authentication.   Use helper functions in your respective client library such as condition constructors (&#x60;equal&#x60;, &#x60;contains&#x60;, &#x60;startsWith&#x60;) and field selectors (&#x60;field&#x60;, &#x60;fields&#x60;, &#x60;id&#x60;) for easier access.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { QueryObjectsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // DataSvcQueryRequest | Query Request (optional)
    body: ...,
  } satisfies QueryObjectsRequest;

  try {
    const data = await api.queryObjects(body);
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
| **body** | [DataSvcQueryRequest](DataSvcQueryRequest.md) | Query Request | [Optional] |

### Return type

[**DataSvcQueryResponse**](DataSvcQueryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful retrieval of objects |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## updateObjects

> object updateObjects(body)

Update Objects

Update fields of objects that match the given filters using the provided object. Any fields not included in the incoming object will remain unchanged.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { UpdateObjectsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // DataSvcUpdateObjectsRequest | Update request payload
    body: ...,
  } satisfies UpdateObjectsRequest;

  try {
    const data = await api.updateObjects(body);
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
| **body** | [DataSvcUpdateObjectsRequest](DataSvcUpdateObjectsRequest.md) | Update request payload | |

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
| **200** | Successful update of objects |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## upsertObject

> DataSvcUpsertObjectResponse upsertObject(objectId, body)

Upsert a Generic Object

Creates a new dynamic object or updates an existing one based on the provided data. Requires authorization and user authentication.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { UpsertObjectRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // string | Object ID
    objectId: objectId_example,
    // DataSvcUpsertObjectRequest | Upsert request payload
    body: ...,
  } satisfies UpsertObjectRequest;

  try {
    const data = await api.upsertObject(body);
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
| **objectId** | `string` | Object ID | [Defaults to `undefined`] |
| **body** | [DataSvcUpsertObjectRequest](DataSvcUpsertObjectRequest.md) | Upsert request payload | |

### Return type

[**DataSvcUpsertObjectResponse**](DataSvcUpsertObjectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful creation or update of object |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## upsertObjects

> DataSvcUpsertObjectResponse upsertObjects(body)

Upsert Objects

Upserts objects by ids.

### Example

```ts
import {
  Configuration,
  DataSvcApi,
} from '';
import type { UpsertObjectsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new DataSvcApi(config);

  const body = {
    // DataSvcUpsertObjectRequest | Upsert request payload
    body: ...,
  } satisfies UpsertObjectsRequest;

  try {
    const data = await api.upsertObjects(body);
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
| **body** | [DataSvcUpsertObjectRequest](DataSvcUpsertObjectRequest.md) | Upsert request payload | |

### Return type

[**DataSvcUpsertObjectResponse**](DataSvcUpsertObjectResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successful upsert of objects |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

