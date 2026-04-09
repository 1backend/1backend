# FirehoseSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**publishEvent**](FirehoseSvcApi.md#publishevent) | **POST** /firehose-svc/event | Publish an Event |
| [**subscribeToEvents**](FirehoseSvcApi.md#subscribetoevents) | **GET** /firehose-svc/events/subscribe | Subscribe to the Event Stream |



## publishEvent

> publishEvent(event)

Publish an Event

Publishes an event to the firehose service after authorization check

### Example

```ts
import {
  Configuration,
  FirehoseSvcApi,
} from '';
import type { PublishEventRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FirehoseSvcApi(config);

  const body = {
    // FirehoseSvcEventPublishRequest | Event to publish
    event: ...,
  } satisfies PublishEventRequest;

  try {
    const data = await api.publishEvent(body);
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
| **event** | [FirehoseSvcEventPublishRequest](FirehoseSvcEventPublishRequest.md) | Event to publish | |

### Return type

`void` (Empty response body)

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

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## subscribeToEvents

> string subscribeToEvents()

Subscribe to the Event Stream

Establish a subscription to the firehose events and accept a real time stream of them.

### Example

```ts
import {
  Configuration,
  FirehoseSvcApi,
} from '';
import type { SubscribeToEventsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FirehoseSvcApi(config);

  try {
    const data = await api.subscribeToEvents();
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

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `text/event-stream`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Event data |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

