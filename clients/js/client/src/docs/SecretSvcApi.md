# SecretSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**decryptValue**](SecretSvcApi.md#decryptvalue) | **POST** /secret-svc/decrypt | Decrypt a Value |
| [**encryptValue**](SecretSvcApi.md#encryptvalue) | **POST** /secret-svc/encrypt | Encrypt a Value |
| [**isSecure**](SecretSvcApi.md#issecure) | **GET** /secret-svc/is-secure | Check Security Status |
| [**listSecrets**](SecretSvcApi.md#listsecrets) | **POST** /secret-svc/secrets | List Secrets |
| [**removeSecrets**](SecretSvcApi.md#removesecrets) | **DELETE** /secret-svc/secrets | Remove Secrets |
| [**saveSecrets**](SecretSvcApi.md#savesecrets) | **PUT** /secret-svc/secrets | Save Secrets |



## decryptValue

> SecretSvcDecryptValueResponse decryptValue(body)

Decrypt a Value

Decrypt a value and return the encrypted result

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { DecryptValueRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  const body = {
    // SecretSvcDecryptValueRequest | Decrypt Value Request
    body: ...,
  } satisfies DecryptValueRequest;

  try {
    const data = await api.decryptValue(body);
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
| **body** | [SecretSvcDecryptValueRequest](SecretSvcDecryptValueRequest.md) | Decrypt Value Request | |

### Return type

[**SecretSvcDecryptValueResponse**](SecretSvcDecryptValueResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Decrypt Value Response |  -  |
| **400** | Bad Request |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## encryptValue

> SecretSvcEncryptValueResponse encryptValue(body)

Encrypt a Value

Encrypt a value and return the encrypted result

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { EncryptValueRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  const body = {
    // SecretSvcEncryptValueRequest | Encrypt Value Request
    body: ...,
  } satisfies EncryptValueRequest;

  try {
    const data = await api.encryptValue(body);
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
| **body** | [SecretSvcEncryptValueRequest](SecretSvcEncryptValueRequest.md) | Encrypt Value Request | |

### Return type

[**SecretSvcEncryptValueResponse**](SecretSvcEncryptValueResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Encrypt Value Response |  -  |
| **400** | Missing Data |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## isSecure

> SecretSvcIsSecureResponse isSecure()

Check Security Status

Returns true if the encryption key is sufficiently secure.

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { IsSecureRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  try {
    const data = await api.isSecure();
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

[**SecretSvcIsSecureResponse**](SecretSvcIsSecureResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Encrypt Value Response |  -  |
| **400** | Bad Request |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listSecrets

> SecretSvcListSecretsResponse listSecrets(body)

List Secrets

List secrets by key(s) if authorized.

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { ListSecretsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  const body = {
    // SecretSvcListSecretsRequest | List Secret Request (optional)
    body: ...,
  } satisfies ListSecretsRequest;

  try {
    const data = await api.listSecrets(body);
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
| **body** | [SecretSvcListSecretsRequest](SecretSvcListSecretsRequest.md) | List Secret Request | [Optional] |

### Return type

[**SecretSvcListSecretsResponse**](SecretSvcListSecretsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List Secret Response |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## removeSecrets

> object removeSecrets(body)

Remove Secrets

Remove secrets if authorized to do so

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { RemoveSecretsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  const body = {
    // SecretSvcRemoveSecretsRequest | Remove Secret Request
    body: ...,
  } satisfies RemoveSecretsRequest;

  try {
    const data = await api.removeSecrets(body);
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
| **body** | [SecretSvcRemoveSecretsRequest](SecretSvcRemoveSecretsRequest.md) | Remove Secret Request | |

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
| **200** | Remove Secret Response |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveSecrets

> object saveSecrets(body)

Save Secrets

Save secrets if authorized to do so. Requires the &#x60;secret-svc:secret:save&#x60; permission. Users can only save secrets prefixed with their user slug unless they also have the &#x60;secret-svc:secret:save-unprefixed&#x60; permission, which allows them to save a secret without a slug prefix. &#x60;secret-svc:secret:save:$secretId&#x60; (eg. &#x60;secret-svc:secret:save:sendgrid-api-key&#x60;) permission allows callers to save secrets otherwise they don\&#39;t have access to. This permission also supports tail-wildcards by splitting the ID with hyphens (e.g., &#x60;secret-svc:secret:save:otp-*&#x60; grants access to &#x60;otp-body-en&#x60; and &#x60;otp-subject-hu&#x60;).

### Example

```ts
import {
  Configuration,
  SecretSvcApi,
} from '';
import type { SaveSecretsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new SecretSvcApi(config);

  const body = {
    // SecretSvcSaveSecretsRequest | Save Secret Request
    body: ...,
  } satisfies SaveSecretsRequest;

  try {
    const data = await api.saveSecrets(body);
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
| **body** | [SecretSvcSaveSecretsRequest](SecretSvcSaveSecretsRequest.md) | Save Secret Request | |

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
| **200** | Save Secret Response |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

