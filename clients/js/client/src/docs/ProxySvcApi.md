# ProxySvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteRoutes**](ProxySvcApi.md#deleteroutes) | **DELETE** /proxy-svc/routes | Delete Routes |
| [**listCerts**](ProxySvcApi.md#listcerts) | **POST** /proxy-svc/certs | List Certs |
| [**listRoutes**](ProxySvcApi.md#listroutes) | **POST** /proxy-svc/routes | List Routes |
| [**saveCerts**](ProxySvcApi.md#savecerts) | **PUT** /proxy-svc/certs | Save Certs |
| [**saveRoutes**](ProxySvcApi.md#saveroutes) | **PUT** /proxy-svc/routes | Save Routes |



## deleteRoutes

> object deleteRoutes(body)

Delete Routes

Delete specific routes by their IDs.

### Example

```ts
import {
  Configuration,
  ProxySvcApi,
} from '';
import type { DeleteRoutesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ProxySvcApi(config);

  const body = {
    // ProxySvcDeleteRoutesRequest | Delete Routes Request
    body: ...,
  } satisfies DeleteRoutesRequest;

  try {
    const data = await api.deleteRoutes(body);
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
| **body** | [ProxySvcDeleteRoutesRequest](ProxySvcDeleteRoutesRequest.md) | Delete Routes Request | |

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
| **200** | Routes deleted successfully |  -  |
| **400** | Invalid JSON or missing IDs |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listCerts

> ProxySvcListCertsResponse listCerts(body)

List Certs

List certs that the edge proxy will use to cert requests.

### Example

```ts
import {
  Configuration,
  ProxySvcApi,
} from '';
import type { ListCertsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ProxySvcApi(config);

  const body = {
    // ProxySvcListCertsRequest | List Certs Request (optional)
    body: ...,
  } satisfies ListCertsRequest;

  try {
    const data = await api.listCerts(body);
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
| **body** | [ProxySvcListCertsRequest](ProxySvcListCertsRequest.md) | List Certs Request | [Optional] |

### Return type

[**ProxySvcListCertsResponse**](ProxySvcListCertsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Certs listed successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listRoutes

> ProxySvcListRoutesResponse listRoutes(body)

List Routes

List routes that the edge proxy will use to route requests.

### Example

```ts
import {
  Configuration,
  ProxySvcApi,
} from '';
import type { ListRoutesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ProxySvcApi(config);

  const body = {
    // ProxySvcListRoutesRequest | List Routes Request (optional)
    body: ...,
  } satisfies ListRoutesRequest;

  try {
    const data = await api.listRoutes(body);
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
| **body** | [ProxySvcListRoutesRequest](ProxySvcListRoutesRequest.md) | List Routes Request | [Optional] |

### Return type

[**ProxySvcListRoutesResponse**](ProxySvcListRoutesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Routes listd successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveCerts

> object saveCerts(body)

Save Certs

This endpoint only exist for testing purposes. Only callable by admins Certs should be saved by the Proxy Svc and its edge proxying functionality internally, not through this endpoint.

### Example

```ts
import {
  Configuration,
  ProxySvcApi,
} from '';
import type { SaveCertsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ProxySvcApi(config);

  const body = {
    // ProxySvcSaveCertsRequest | Save Certs Request
    body: ...,
  } satisfies SaveCertsRequest;

  try {
    const data = await api.saveCerts(body);
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
| **body** | [ProxySvcSaveCertsRequest](ProxySvcSaveCertsRequest.md) | Save Certs Request | |

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
| **200** | Certs saved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveRoutes

> ProxySvcSaveRoutesResponse saveRoutes(body)

Save Routes

Save routes that the edge proxy will use to route requests.

### Example

```ts
import {
  Configuration,
  ProxySvcApi,
} from '';
import type { SaveRoutesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ProxySvcApi(config);

  const body = {
    // ProxySvcSaveRoutesRequest | Save Routes Request
    body: ...,
  } satisfies SaveRoutesRequest;

  try {
    const data = await api.saveRoutes(body);
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
| **body** | [ProxySvcSaveRoutesRequest](ProxySvcSaveRoutesRequest.md) | Save Routes Request | |

### Return type

[**ProxySvcSaveRoutesResponse**](ProxySvcSaveRoutesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Routes saved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

