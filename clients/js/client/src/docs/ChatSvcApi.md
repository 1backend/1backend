# ChatSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteMessage**](ChatSvcApi.md#deletemessage) | **DELETE** /chat-svc/message/{messageId} | Delete a Message |
| [**deleteThread**](ChatSvcApi.md#deletethread) | **DELETE** /chat-svc/thread/{threadId} | Delete a Thread |
| [**events**](ChatSvcApi.md#events) | **GET** /chat-svc/events | Events |
| [**listMessages**](ChatSvcApi.md#listmessages) | **POST** /chat-svc/messages | List Messages |
| [**listThreads**](ChatSvcApi.md#listthreads) | **POST** /chat-svc/threads | List Threads |
| [**saveMessage**](ChatSvcApi.md#savemessage) | **POST** /chat-svc/thread/{threadId}/message | Save Message |
| [**saveThread**](ChatSvcApi.md#savethread) | **POST** /chat-svc/thread | Save Thread |



## deleteMessage

> { [key: string]: any; } deleteMessage(messageId)

Delete a Message

Delete a specific message from a chat thread by its ID

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { DeleteMessageRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // string | Message ID
    messageId: messageId_example,
  } satisfies DeleteMessageRequest;

  try {
    const data = await api.deleteMessage(body);
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
| **messageId** | `string` | Message ID | [Defaults to `undefined`] |

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
| **200** | Message successfully deleted |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## deleteThread

> { [key: string]: any; } deleteThread(threadId)

Delete a Thread

Delete a specific chat thread by its ID

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { DeleteThreadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // string | Thread ID
    threadId: threadId_example,
  } satisfies DeleteThreadRequest;

  try {
    const data = await api.deleteThread(body);
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

**{ [key: string]: any; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Thread successfully deleted |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## events

> ChatSvcEventThreadUpdate events()

Events

Events is a dummy endpoint to display documentation about the events that this service emits.

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { EventsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ChatSvcApi();

  try {
    const data = await api.events();
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

[**ChatSvcEventThreadUpdate**](ChatSvcEventThreadUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listMessages

> ChatSvcListMessagesResponse listMessages(body)

List Messages

Fetch messages (and associated assets) for a specific chat thread.

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { ListMessagesRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // ChatSvcListMessagesRequest | List Messages Request
    body: ...,
  } satisfies ListMessagesRequest;

  try {
    const data = await api.listMessages(body);
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
| **body** | [ChatSvcListMessagesRequest](ChatSvcListMessagesRequest.md) | List Messages Request | |

### Return type

[**ChatSvcListMessagesResponse**](ChatSvcListMessagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Messages and assets successfully retrieved |  -  |
| **400** | Invalid JSON |  -  |
| **401** | unauthorized |  -  |
| **500** | internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listThreads

> ChatSvcListThreadsResponse listThreads(body)

List Threads

Fetch all chat threads associated with a specific user

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { ListThreadsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // ChatSvcListThreadsRequest | List Threads Request
    body: ...,
  } satisfies ListThreadsRequest;

  try {
    const data = await api.listThreads(body);
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
| **body** | [ChatSvcListThreadsRequest](ChatSvcListThreadsRequest.md) | List Threads Request | |

### Return type

[**ChatSvcListThreadsResponse**](ChatSvcListThreadsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Threads successfully retrieved |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveMessage

> { [key: string]: any; } saveMessage(threadId, body)

Save Message

Save a new message to a specific thread.

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { SaveMessageRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // string | Thread ID
    threadId: threadId_example,
    // ChatSvcSaveMessageRequest | Save Message Request
    body: ...,
  } satisfies SaveMessageRequest;

  try {
    const data = await api.saveMessage(body);
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
| **body** | [ChatSvcSaveMessageRequest](ChatSvcSaveMessageRequest.md) | Save Message Request | |

### Return type

**{ [key: string]: any; }**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Message successfully added |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveThread

> ChatSvcSaveThreadResponse saveThread(body)

Save Thread

Create or update a chat thread. Requires the &#x60;chat-svc:thread:edit&#x60; permission.

### Example

```ts
import {
  Configuration,
  ChatSvcApi,
} from '';
import type { SaveThreadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ChatSvcApi(config);

  const body = {
    // ChatSvcSaveThreadRequest | Save Thread Request
    body: ...,
  } satisfies SaveThreadRequest;

  try {
    const data = await api.saveThread(body);
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
| **body** | [ChatSvcSaveThreadRequest](ChatSvcSaveThreadRequest.md) | Save Thread Request | |

### Return type

[**ChatSvcSaveThreadResponse**](ChatSvcSaveThreadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Thread successfully created |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

