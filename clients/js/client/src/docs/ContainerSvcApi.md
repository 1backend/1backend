# ContainerSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**buildImage**](ContainerSvcApi.md#buildimage) | **PUT** /container-svc/image | Build an Image |
| [**containerDaemonInfo**](ContainerSvcApi.md#containerdaemoninfo) | **GET** /container-svc/daemon/info | Get Container Daemon Information |
| [**containerIsRunning**](ContainerSvcApi.md#containerisrunning) | **GET** /container-svc/container/is-running | Check If a Container Is Running |
| [**containerSummary**](ContainerSvcApi.md#containersummary) | **GET** /container-svc/container/summary | Get Container Summary |
| [**getHost**](ContainerSvcApi.md#gethost) | **GET** /container-svc/host | Get Container Host |
| [**imagePullable**](ContainerSvcApi.md#imagepullable) | **GET** /container-svc/image/{imageName}/pullable | Check if Container Image is Pullable |
| [**listContainerLogs**](ContainerSvcApi.md#listcontainerlogs) | **POST** /container-svc/logs | List Logs |
| [**listContainers**](ContainerSvcApi.md#listcontainers) | **POST** /container-svc/containers | List Containers |
| [**runContainer**](ContainerSvcApi.md#runcontainer) | **PUT** /container-svc/container | Run a Container |
| [**stopContainer**](ContainerSvcApi.md#stopcontainer) | **PUT** /container-svc/container/stop | Stop a Container |



## buildImage

> object buildImage(body)

Build an Image

Builds a Docker image with the specified parameters.  Requires the &#x60;container-svc:image:build&#x60; permission.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { BuildImageRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // ContainerSvcBuildImageRequest | Build Image Request
    body: ...,
  } satisfies BuildImageRequest;

  try {
    const data = await api.buildImage(body);
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
| **body** | [ContainerSvcBuildImageRequest](ContainerSvcBuildImageRequest.md) | Build Image Request | |

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
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## containerDaemonInfo

> ContainerSvcDaemonInfoResponse containerDaemonInfo()

Get Container Daemon Information

Retrieve detailed information about the availability and status of container daemons on the node.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ContainerDaemonInfoRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  try {
    const data = await api.containerDaemonInfo();
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

[**ContainerSvcDaemonInfoResponse**](ContainerSvcDaemonInfoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Service Information |  -  |
| **401** | unauthorized |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## containerIsRunning

> ContainerSvcContainerIsRunningResponse containerIsRunning(hash, name)

Check If a Container Is Running

Check if a Docker container is running, identified by hash or name.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ContainerIsRunningRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // string | Container Hash (optional)
    hash: hash_example,
    // string | Container Name (optional)
    name: name_example,
  } satisfies ContainerIsRunningRequest;

  try {
    const data = await api.containerIsRunning(body);
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
| **hash** | `string` | Container Hash | [Optional] [Defaults to `undefined`] |
| **name** | `string` | Container Name | [Optional] [Defaults to `undefined`] |

### Return type

[**ContainerSvcContainerIsRunningResponse**](ContainerSvcContainerIsRunningResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | missing parameters |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## containerSummary

> ContainerSvcGetContainerSummaryResponse containerSummary(hash, name, lines)

Get Container Summary

Get a summary of the Docker container identified by hash or name, limited to a specified number of lines.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ContainerSummaryRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // string | Container Hash (optional)
    hash: hash_example,
    // string | Container Name (optional)
    name: name_example,
    // number | Number of Lines (optional)
    lines: 56,
  } satisfies ContainerSummaryRequest;

  try {
    const data = await api.containerSummary(body);
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
| **hash** | `string` | Container Hash | [Optional] [Defaults to `undefined`] |
| **name** | `string` | Container Name | [Optional] [Defaults to `undefined`] |
| **lines** | `number` | Number of Lines | [Optional] [Defaults to `undefined`] |

### Return type

[**ContainerSvcGetContainerSummaryResponse**](ContainerSvcGetContainerSummaryResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | missing parameters |  -  |
| **401** | unauthorized |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getHost

> ContainerSvcGetHostResponse getHost()

Get Container Host

Retrieve information about the Container host

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { GetHostRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  try {
    const data = await api.getHost();
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

[**ContainerSvcGetHostResponse**](ContainerSvcGetHostResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | unauthorized |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## imagePullable

> ContainerSvcImagePullableResponse imagePullable(imageName)

Check if Container Image is Pullable

Check if an image exists on in the container registry and is pullable.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ImagePullableRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // string | Image name
    imageName: imageName_example,
  } satisfies ImagePullableRequest;

  try {
    const data = await api.imagePullable(body);
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
| **imageName** | `string` | Image name | [Defaults to `undefined`] |

### Return type

[**ContainerSvcImagePullableResponse**](ContainerSvcImagePullableResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | model ID in path is not URL encoded |  -  |
| **401** | unauthorized |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listContainerLogs

> ContainerSvcListLogsResponse listContainerLogs(body)

List Logs

List Container logs.  Requires the &#x60;container-svc:log:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ListContainerLogsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // ContainerSvcListLogsRequest | List Logs Request
    body: ...,
  } satisfies ListContainerLogsRequest;

  try {
    const data = await api.listContainerLogs(body);
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
| **body** | [ContainerSvcListLogsRequest](ContainerSvcListLogsRequest.md) | List Logs Request | |

### Return type

[**ContainerSvcListLogsResponse**](ContainerSvcListLogsResponse.md)

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


## listContainers

> ContainerSvcListContainersResponse listContainers(body)

List Containers

List containers.  Requires the &#x60;container-svc:container:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { ListContainersRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // ContainerSvcListContainersRequest | List Containers Request
    body: ...,
  } satisfies ListContainersRequest;

  try {
    const data = await api.listContainers(body);
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
| **body** | [ContainerSvcListContainersRequest](ContainerSvcListContainersRequest.md) | List Containers Request | |

### Return type

[**ContainerSvcListContainersResponse**](ContainerSvcListContainersResponse.md)

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


## runContainer

> ContainerSvcRunContainerResponse runContainer(body)

Run a Container

Runs a Docker container with the specified parameters.  Requires the &#x60;container-svc:container:run&#x60; permission.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { RunContainerRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // ContainerSvcRunContainerRequest | Run Container Request
    body: ...,
  } satisfies RunContainerRequest;

  try {
    const data = await api.runContainer(body);
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
| **body** | [ContainerSvcRunContainerRequest](ContainerSvcRunContainerRequest.md) | Run Container Request | |

### Return type

[**ContainerSvcRunContainerResponse**](ContainerSvcRunContainerResponse.md)

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


## stopContainer

> object stopContainer(body)

Stop a Container

Stops a Docker container with the specified parameters.  Requires the &#x60;container-svc:container:stop&#x60; permission.

### Example

```ts
import {
  Configuration,
  ContainerSvcApi,
} from '';
import type { StopContainerRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ContainerSvcApi(config);

  const body = {
    // ContainerSvcStopContainerRequest | Stop Container Request
    body: ...,
  } satisfies StopContainerRequest;

  try {
    const data = await api.stopContainer(body);
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
| **body** | [ContainerSvcStopContainerRequest](ContainerSvcStopContainerRequest.md) | Stop Container Request | |

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
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

