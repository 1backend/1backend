# PromptSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**listPrompts**](PromptSvcApi.md#listprompts) | **POST** /prompt-svc/prompts | List Prompts |
| [**prompt**](PromptSvcApi.md#prompt) | **POST** /prompt-svc/prompt | Prompt an AI |
| [**promptTypes**](PromptSvcApi.md#prompttypes) | **POST** /prompt-svc/types | Prompt Types |
| [**removePrompt**](PromptSvcApi.md#removeprompt) | **POST** /prompt-svc/remove | Remove Prompt |
| [**subscribeToPromptResponses**](PromptSvcApi.md#subscribetopromptresponses) | **GET** /prompt-svc/prompts/{threadId}/responses/subscribe | Subscribe to Prompt Responses by Thread |



## listPrompts

> PromptSvcListPromptsResponse listPrompts(body)

List Prompts

List prompts that satisfy a query.

### Example

```ts
import {
  Configuration,
  PromptSvcApi,
} from '';
import type { ListPromptsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PromptSvcApi(config);

  const body = {
    // PromptSvcListPromptsRequest | List Prompts Request (optional)
    body: ...,
  } satisfies ListPromptsRequest;

  try {
    const data = await api.listPrompts(body);
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
| **body** | [PromptSvcListPromptsRequest](PromptSvcListPromptsRequest.md) | List Prompts Request | [Optional] |

### Return type

[**PromptSvcListPromptsResponse**](PromptSvcListPromptsResponse.md)

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


## prompt

> PromptSvcPromptResponse prompt(body)

Prompt an AI

Sends a prompt and waits for a response if sync is true. If sync is false, adds the prompt to the queue and returns immediately.  Prompts can be used for &#x60;text-to-text&#x60;, &#x60;text-to-image&#x60;, &#x60;image-to-image&#x60;, and other types of generation. If no model ID is specified, the default model will be used (see &#x60;Model Svc&#x60; for details). The default model may or may not support the requested generation type.  **Prompting Modes** - **High-Level Parameters**: Uses predefined parameters relevant to &#x60;text-to-image&#x60;, &#x60;image-to-image&#x60;, etc. This mode abstracts away the underlying engine (e.g., LLaMA, Stable Diffusion) and focuses on functionality. - **Engine-Specific Parameters**: Uses &#x60;engineParameters&#x60; to directly specify an AI engine, exposing all available parameters for fine-tuned control.  **Permissions Required:** &#x60;prompt-svc:prompt:create&#x60;

### Example

```ts
import {
  Configuration,
  PromptSvcApi,
} from '';
import type { PromptRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PromptSvcApi(config);

  const body = {
    // PromptSvcPromptRequest | Add Prompt Request
    body: ...,
  } satisfies PromptRequest;

  try {
    const data = await api.prompt(body);
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
| **body** | [PromptSvcPromptRequest](PromptSvcPromptRequest.md) | Add Prompt Request | |

### Return type

[**PromptSvcPromptResponse**](PromptSvcPromptResponse.md)

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


## promptTypes

> PromptSvcTypesResponse promptTypes(body)

Prompt Types

The only purpose of this \&quot;endpoint\&quot; is to export types otherwise not appearing in the API docs. This endpoint otherwise does nothing. Do not depend on this endpoint, only its types.

### Example

```ts
import {
  Configuration,
  PromptSvcApi,
} from '';
import type { PromptTypesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PromptSvcApi(config);

  const body = {
    // object | Types Request
    body: Object,
  } satisfies PromptTypesRequest;

  try {
    const data = await api.promptTypes(body);
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
| **body** | `object` | Types Request | |

### Return type

[**PromptSvcTypesResponse**](PromptSvcTypesResponse.md)

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


## removePrompt

> object removePrompt(body)

Remove Prompt

Remove a prompt by ID.

### Example

```ts
import {
  Configuration,
  PromptSvcApi,
} from '';
import type { RemovePromptRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PromptSvcApi(config);

  const body = {
    // PromptSvcRemovePromptRequest | Remove Prompt Request
    body: ...,
  } satisfies RemovePromptRequest;

  try {
    const data = await api.removePrompt(body);
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
| **body** | [PromptSvcRemovePromptRequest](PromptSvcRemovePromptRequest.md) | Remove Prompt Request | |

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
| **200** | {} |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## subscribeToPromptResponses

> string subscribeToPromptResponses(threadId)

Subscribe to Prompt Responses by Thread

Subscribe to prompt responses by thread via Server-Sent Events (SSE). You can subscribe to threads before they are created. The streamed strings are of type &#x60;StreamChunk&#x60;, see the PromptTypes endpoint for more details.

### Example

```ts
import {
  Configuration,
  PromptSvcApi,
} from '';
import type { SubscribeToPromptResponsesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PromptSvcApi(config);

  const body = {
    // string | Thread ID
    threadId: threadId_example,
  } satisfies SubscribeToPromptResponsesRequest;

  try {
    const data = await api.subscribeToPromptResponses(body);
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
| **threadId** | `string` | Thread ID | [Defaults to `undefined`] |

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `*/*`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Streaming response |  -  |
| **400** | Missing Parameter |  -  |
| **401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

