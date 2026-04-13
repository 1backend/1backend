# SourceSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**checkoutRepo**](SourceSvcApi.md#checkoutrepo) | **POST** /source-svc/repo/checkout | Checkout a git repository |



## checkoutRepo

> SourceSvcCheckoutRepoResponse checkoutRepo(body)

Checkout a git repository

Checkout a git repository over https or ssh at a specific version into a temporary directory. Performs a shallow clone with minimal history for faster checkout.

### Example

```ts
import {
  Configuration,
  SourceSvcApi,
} from '';
import type { CheckoutRepoRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SourceSvcApi(config);

  const body = {
    // SourceSvcCheckoutRepoRequest | Checkout Repo Request
    body: ...,
  } satisfies CheckoutRepoRequest;

  try {
    const data = await api.checkoutRepo(body);
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
| **body** | [SourceSvcCheckoutRepoRequest](SourceSvcCheckoutRepoRequest.md) | Checkout Repo Request | |

### Return type

[**SourceSvcCheckoutRepoResponse**](SourceSvcCheckoutRepoResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successfully checked out the repository |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

