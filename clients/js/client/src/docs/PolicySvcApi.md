# PolicySvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**check**](PolicySvcApi.md#check) | **POST** /policy-svc/check | Check |
| [**upsertInstance**](PolicySvcApi.md#upsertinstance) | **PUT** /policy-svc/instance/{instanceId} | Upsert an Instance |



## check

> PolicySvcCheckResponse check(body)

Check

Check records a resource access and returns if the access is allowed.

### Example

```ts
import {
  Configuration,
  PolicySvcApi,
} from '';
import type { CheckRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PolicySvcApi(config);

  const body = {
    // PolicySvcCheckRequest | Check Request
    body: ...,
  } satisfies CheckRequest;

  try {
    const data = await api.check(body);
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
| **body** | [PolicySvcCheckRequest](PolicySvcCheckRequest.md) | Check Request | |

### Return type

[**PolicySvcCheckResponse**](PolicySvcCheckResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Checked successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## upsertInstance

> object upsertInstance(instanceId, body)

Upsert an Instance

Allows user to upsert a new policy instance based on a template.

### Example

```ts
import {
  Configuration,
  PolicySvcApi,
} from '';
import type { UpsertInstanceRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new PolicySvcApi(config);

  const body = {
    // string | Instance ID
    instanceId: instanceId_example,
    // PolicySvcUpsertInstanceRequest | Upsert Instance Request
    body: ...,
  } satisfies UpsertInstanceRequest;

  try {
    const data = await api.upsertInstance(body);
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
| **instanceId** | `string` | Instance ID | [Defaults to `undefined`] |
| **body** | [PolicySvcUpsertInstanceRequest](PolicySvcUpsertInstanceRequest.md) | Upsert Instance Request | |

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
| **200** | Instance upserted successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

