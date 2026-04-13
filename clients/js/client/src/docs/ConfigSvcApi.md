# ConfigSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**listConfigVersions**](ConfigSvcApi.md#listconfigversions) | **POST** /config-svc/versions | List Versions |
| [**listConfigs**](ConfigSvcApi.md#listconfigs) | **POST** /config-svc/configs | List Configs |
| [**saveConfig**](ConfigSvcApi.md#saveconfig) | **PUT** /config-svc/config | Save Config |



## listConfigVersions

> ConfigSvcListVersionsResponse listConfigVersions(body)

List Versions

Returns the historical versions of a configuration for a given app. Intended for retrieving the version history of a **single configuration ID**. Supplying multiple IDs is supported but not recommended, since results from different IDs will interleave in the same time-ordered list, making chronological paging ambiguous.

### Example

```ts
import {
  Configuration,
  ConfigSvcApi,
} from '';
import type { ListConfigVersionsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ConfigSvcApi();

  const body = {
    // ConfigSvcListVersionsRequest | List Configs Request
    body: ...,
  } satisfies ListConfigVersionsRequest;

  try {
    const data = await api.listConfigVersions(body);
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
| **body** | [ConfigSvcListVersionsRequest](ConfigSvcListVersionsRequest.md) | List Configs Request | |

### Return type

[**ConfigSvcListVersionsResponse**](ConfigSvcListVersionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Current configuration |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listConfigs

> ConfigSvcListConfigsResponse listConfigs(body, cacheControl)

List Configs

Retrieves the current configurations for a specified app. Since any user can save configurations, it is strongly advised that you supply a list of owners to filter on. If no app is specified, the default \&quot;unnamed\&quot; app is used. This is a public endpoint and does not require authentication. Configuration data is non-sensitive. For sensitive data, refer to the Secret Service.  Configurations are used to control frontend behavior, A/B testing, feature flags, and other non-sensitive settings.

### Example

```ts
import {
  Configuration,
  ConfigSvcApi,
} from '';
import type { ListConfigsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ConfigSvcApi();

  const body = {
    // ConfigSvcListConfigsRequest | List Configs Request
    body: ...,
    // string | Bypass cache (use \'no-cache\') (optional)
    cacheControl: cacheControl_example,
  } satisfies ListConfigsRequest;

  try {
    const data = await api.listConfigs(body);
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
| **body** | [ConfigSvcListConfigsRequest](ConfigSvcListConfigsRequest.md) | List Configs Request | |
| **cacheControl** | `string` | Bypass cache (use \&#39;no-cache\&#39;) | [Optional] [Defaults to `undefined`] |

### Return type

[**ConfigSvcListConfigsResponse**](ConfigSvcListConfigsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Current configuration |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveConfig

> object saveConfig(body)

Save Config

Save the provided configuration to the server. The app from the caller\&#39;s token is used to determine which app the config belongs to. The caller\&#39;s camelCased slug (e.g., \&quot;test-user-slug\&quot; becomes \&quot;testUserSlug\&quot;) is used as the config key automatically, except for users who have the \&quot;config-svc:config:edit-on-behalf\&quot; permission (admins), who can specify any key they want. Admins (users with the \&quot;config-svc:config:edit-on-behalf\&quot; permission) can also provide an \&quot;app\&quot; field in the request body to specify which app the config belongs to, while non-admin users cannot specify the \&quot;app\&quot; field, the app associated with their token will be used.  The save performs a deep merge, that is: - Nested objects are recursively merged rather than replaced. - If a field exists in both the existing and the incoming config and both values are objects, their contents are merged. - If a field exists in both but one or both values are not objects (e.g., string, number, array), the incoming value replaces the existing one. - Fields present only in the incoming config are added. - Fields present only in the existing config are preserved. - Top-level and nested merges follow the same rules.

### Example

```ts
import {
  Configuration,
  ConfigSvcApi,
} from '';
import type { SaveConfigRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new ConfigSvcApi(config);

  const body = {
    // ConfigSvcSaveConfigRequest | Save Config Request
    body: ...,
  } satisfies SaveConfigRequest;

  try {
    const data = await api.saveConfig(body);
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
| **body** | [ConfigSvcSaveConfigRequest](ConfigSvcSaveConfigRequest.md) | Save Config Request | |

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
| **200** | Save Config Response |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

