# ModelSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getDefaultModelStatus**](ModelSvcApi.md#getdefaultmodelstatus) | **GET** /model-svc/default-model/status | Get Default Model Status |
| [**getModel**](ModelSvcApi.md#getmodel) | **GET** /model-svc/model/{modelId} | Get a Model |
| [**getModelStatus**](ModelSvcApi.md#getmodelstatus) | **GET** /model-svc/model/{modelId}/status | Get Model Status |
| [**listModels**](ModelSvcApi.md#listmodels) | **POST** /model-svc/models | List Models |
| [**listPlatforms**](ModelSvcApi.md#listplatforms) | **POST** /model-svc/platforms | List Platforms |
| [**makeDefault**](ModelSvcApi.md#makedefault) | **PUT** /model-svc/model/{modelId}/make-default | Make a Model Default |
| [**startDefaultModel**](ModelSvcApi.md#startdefaultmodel) | **PUT** /model-svc/default-model/start | Start the Default Model |
| [**startModel**](ModelSvcApi.md#startmodel) | **PUT** /model-svc/model/{modelId}/start | Start a Model |



## getDefaultModelStatus

> ModelSvcStatusResponse getDefaultModelStatus()

Get Default Model Status

Retrieves the status of the default model.  Requires the &#x60;model-svc:model:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { GetDefaultModelStatusRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  try {
    const data = await api.getDefaultModelStatus();
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

[**ModelSvcStatusResponse**](ModelSvcStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Model status retrieved successfully |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getModel

> ModelSvcGetModelResponse getModel(modelId)

Get a Model

Retrieves the details of a model by its ID.  the Requires &#x60;model.view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { GetModelRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  const body = {
    // string | Model ID
    modelId: modelId_example,
  } satisfies GetModelRequest;

  try {
    const data = await api.getModel(body);
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
| **modelId** | `string` | Model ID | [Defaults to `undefined`] |

### Return type

[**ModelSvcGetModelResponse**](ModelSvcGetModelResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid Model ID |  -  |
| **401** | Unauthorized |  -  |
| **404** | Model Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getModelStatus

> ModelSvcStatusResponse getModelStatus(modelId)

Get Model Status

Retrieves the status of a model by ID.  Requires the &#x60;model-svc:model:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { GetModelStatusRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  const body = {
    // string | Model ID
    modelId: modelId_example,
  } satisfies GetModelStatusRequest;

  try {
    const data = await api.getModelStatus(body);
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
| **modelId** | `string` | Model ID | [Defaults to `undefined`] |

### Return type

[**ModelSvcStatusResponse**](ModelSvcStatusResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Model status retrieved successfully |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listModels

> ModelSvcListModelsResponse listModels()

List Models

Retrieves a list of models.  Requires &#x60;model-svc:model:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { ListModelsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  try {
    const data = await api.listModels();
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

[**ModelSvcListModelsResponse**](ModelSvcListModelsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listPlatforms

> ModelSvcListPlatformsResponse listPlatforms()

List Platforms

Retrieves a list of AI platforms. Eg. LlamaCpp, StableDiffusion etc.  Requires &#x60;model-svc:platform:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { ListPlatformsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  try {
    const data = await api.listPlatforms();
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

[**ModelSvcListPlatformsResponse**](ModelSvcListPlatformsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## makeDefault

> object makeDefault(modelId)

Make a Model Default

Sets a model as the default model — when prompts are sent without a Model ID, the default model is used.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { MakeDefaultRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  const body = {
    // string | Model ID
    modelId: modelId_example,
  } satisfies MakeDefaultRequest;

  try {
    const data = await api.makeDefault(body);
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
| **modelId** | `string` | Model ID | [Defaults to `undefined`] |

### Return type

**object**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid Model ID |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## startDefaultModel

> object startDefaultModel()

Start the Default Model

Starts The Default Model.  Requires the &#x60;model-svc:model:create&#x60; permission.

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { StartDefaultModelRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  try {
    const data = await api.startDefaultModel();
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

**object**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## startModel

> object startModel(modelId)

Start a Model

Starts a model by ID

### Example

```ts
import {
  Configuration,
  ModelSvcApi,
} from '';
import type { StartModelRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ModelSvcApi(config);

  const body = {
    // string | Model ID
    modelId: modelId_example,
  } satisfies StartModelRequest;

  try {
    const data = await api.startModel(body);
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
| **modelId** | `string` | Model ID | [Defaults to `undefined`] |

### Return type

**object**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

