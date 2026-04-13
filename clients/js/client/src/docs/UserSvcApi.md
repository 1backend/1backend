# UserSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**activateOrganization**](UserSvcApi.md#activateorganization) | **POST** /user-svc/organization/activate | Activate Organization |
| [**changePassword**](UserSvcApi.md#changepassword) | **POST** /user-svc/change-password | Change Password |
| [**createUser**](UserSvcApi.md#createuser) | **POST** /user-svc/user | Create a New User |
| [**deleteMembership**](UserSvcApi.md#deletemembership) | **DELETE** /user-svc/organization/{organizationId}/user/{userId} | Delete Membership |
| [**deleteUser**](UserSvcApi.md#deleteuser) | **DELETE** /user-svc/user/{userId} | Delete a User |
| [**exchangeToken**](UserSvcApi.md#exchangetoken) | **PUT** /user-svc/token/exchange | Exchange Token |
| [**exchangeToken_0**](UserSvcApi.md#exchangetoken_0) | **PUT** /user-svc/token/exchange | Exchange Token |
| [**getPublicKey**](UserSvcApi.md#getpublickey) | **GET** /user-svc/public-key | Get Public Key |
| [**hasPermission**](UserSvcApi.md#haspermission) | **POST** /user-svc/self/has/{permission} | Has Permission |
| [**listApps**](UserSvcApi.md#listapps) | **POST** /user-svc/apps | List Apps |
| [**listEnrolls**](UserSvcApi.md#listenrolls) | **POST** /user-svc/enrolls | List Enrolls |
| [**listOrganizations**](UserSvcApi.md#listorganizations) | **POST** /user-svc/organizations | List Organizations |
| [**listPermissions**](UserSvcApi.md#listpermissions) | **POST** /user-svc/permissions | List Permissions |
| [**listPermits**](UserSvcApi.md#listpermits) | **POST** /user-svc/permits | List Permits |
| [**listUsers**](UserSvcApi.md#listusers) | **POST** /user-svc/users | List Users |
| [**login**](UserSvcApi.md#login) | **POST** /user-svc/login | Login |
| [**readApp**](UserSvcApi.md#readapp) | **POST** /user-svc/app | Read or Create App |
| [**readSelf**](UserSvcApi.md#readself) | **POST** /user-svc/self | Read Self |
| [**refreshToken**](UserSvcApi.md#refreshtoken) | **POST** /user-svc/refresh-token | Refresh Token |
| [**register**](UserSvcApi.md#register) | **POST** /user-svc/register | Register |
| [**resetPassword**](UserSvcApi.md#resetpassword) | **POST** /user-svc/{userId}/reset-password | Reset Password |
| [**revokeTokens**](UserSvcApi.md#revoketokens) | **DELETE** /user-svc/tokens | Revoke Tokens |
| [**saveEnrolls**](UserSvcApi.md#saveenrolls) | **PUT** /user-svc/enrolls | Save Enrolls |
| [**saveMembership**](UserSvcApi.md#savemembership) | **PUT** /user-svc/organization/{organizationId}/user/{userId} | Save Membership |
| [**saveOrganization**](UserSvcApi.md#saveorganization) | **PUT** /user-svc/organization | Save an Organization |
| [**savePermits**](UserSvcApi.md#savepermits) | **PUT** /user-svc/permits | Save Permits |
| [**saveSelf**](UserSvcApi.md#saveself) | **PUT** /user-svc/self | Save User Profile |
| [**saveUser**](UserSvcApi.md#saveuser) | **PUT** /user-svc/user/{userId} | Save User |
| [**sendOtp**](UserSvcApi.md#sendotp) | **POST** /user-svc/otp/send | Send OTP |
| [**updateApp**](UserSvcApi.md#updateapp) | **PUT** /user-svc/app | Update App Host |



## activateOrganization

> UserSvcActivateOrganizationResponse activateOrganization(body)

Activate Organization

Sets the caller user\&#39;s active organization and returns a fresh token reflecting the new active organization.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ActivateOrganizationRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcActivateOrganizationRequest | Activate Organization Request
    body: ...,
  } satisfies ActivateOrganizationRequest;

  try {
    const data = await api.activateOrganization(body);
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
| **body** | [UserSvcActivateOrganizationRequest](UserSvcActivateOrganizationRequest.md) | Activate Organization Request | |

### Return type

[**UserSvcActivateOrganizationResponse**](UserSvcActivateOrganizationResponse.md)

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


## changePassword

> object changePassword(body)

Change Password

Allows an authenticated user to change their own password.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ChangePasswordRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcChangePasswordRequest | Change Password Request
    body: ...,
  } satisfies ChangePasswordRequest;

  try {
    const data = await api.changePassword(body);
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
| **body** | [UserSvcChangePasswordRequest](UserSvcChangePasswordRequest.md) | Change Password Request | |

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
| **200** | Password changed successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## createUser

> object createUser(body)

Create a New User

Allows an authenticated administrator to create a new user with specified details.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { CreateUserRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcCreateUserRequest | Create User Request
    body: ...,
  } satisfies CreateUserRequest;

  try {
    const data = await api.createUser(body);
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
| **body** | [UserSvcCreateUserRequest](UserSvcCreateUserRequest.md) | Create User Request | |

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
| **200** | User created successfully |  -  |
| **400** | Invalid User |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## deleteMembership

> object deleteMembership(organizationId, userId, body)

Delete Membership

Allows an organization admin to remove a user from an organization.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { DeleteMembershipRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | Organization ID
    organizationId: organizationId_example,
    // string | User ID
    userId: userId_example,
    // object | Remove User From Organization Request (optional)
    body: Object,
  } satisfies DeleteMembershipRequest;

  try {
    const data = await api.deleteMembership(body);
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
| **organizationId** | `string` | Organization ID | [Defaults to `undefined`] |
| **userId** | `string` | User ID | [Defaults to `undefined`] |
| **body** | `object` | Remove User From Organization Request | [Optional] |

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
| **200** | User added successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **403** | Forbidden |  -  |
| **404** | Organization/User not found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## deleteUser

> object deleteUser(userId)

Delete a User

Delete a user based on the user ID.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { DeleteUserRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | User ID
    userId: userId_example,
  } satisfies DeleteUserRequest;

  try {
    const data = await api.deleteUser(body);
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
| **userId** | `string` | User ID | [Defaults to `undefined`] |

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


## exchangeToken

> UserSvcExchangeTokenResponse exchangeToken(body)

Exchange Token

Exchange an existing token for a new token scoped to a different app (namespace). The new token represents the same user but contains roles specific to the target app.  The original token remains valid. The minted token is not stored and cannot be refreshed (and will have the same expiration duration as normal tokens), unlike tokens acquired via login.  For now, token exchange is designed to be in situ — the User Svc must be contacted at exchange time. This introduces a stateful dependency on the User Svc, but simplifies things until broader use cases emerge.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ExchangeTokenRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  const body = {
    // UserSvcExchangeTokenRequest | ExchangeToken Request
    body: ...,
  } satisfies ExchangeTokenRequest;

  try {
    const data = await api.exchangeToken(body);
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
| **body** | [UserSvcExchangeTokenRequest](UserSvcExchangeTokenRequest.md) | ExchangeToken Request | |

### Return type

[**UserSvcExchangeTokenResponse**](UserSvcExchangeTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ExchangeToken successful |  -  |
| **400** | Invalid JSON |  -  |
| **404** | User Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## exchangeToken_0

> UserSvcExchangeTokenResponse exchangeToken_0(body)

Exchange Token

Exchange an existing token for a new token scoped to a different app (namespace). The new token represents the same user but contains roles specific to the target app.  The original token remains valid. The minted token is not stored and cannot be refreshed (and will have the same expiration duration as normal tokens), unlike tokens acquired via login.  For now, token exchange is designed to be in situ — the User Svc must be contacted at exchange time. This introduces a stateful dependency on the User Svc, but simplifies things until broader use cases emerge.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ExchangeToken0Request } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  const body = {
    // UserSvcExchangeTokenRequest | ExchangeToken Request
    body: ...,
  } satisfies ExchangeToken0Request;

  try {
    const data = await api.exchangeToken_0(body);
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
| **body** | [UserSvcExchangeTokenRequest](UserSvcExchangeTokenRequest.md) | ExchangeToken Request | |

### Return type

[**UserSvcExchangeTokenResponse**](UserSvcExchangeTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | ExchangeToken successful |  -  |
| **400** | Invalid JSON |  -  |
| **404** | User Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getPublicKey

> UserSvcGetPublicKeyResponse getPublicKey()

Get Public Key

Get the public key to verify the JWT signature.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { GetPublicKeyRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  try {
    const data = await api.getPublicKey();
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

[**UserSvcGetPublicKeyResponse**](UserSvcGetPublicKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid JSON or missing permission id |  -  |
| **401** | Unauthorized |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## hasPermission

> UserSvcHasPermissionResponse hasPermission(permission)

Has Permission

Checks whether the caller has a specific permission. Optimized for caching — only the caller and the permission are required. To assign a permission to a user or role, use the &#x60;Save Permits&#x60; endpoint.  This endpoint does not return 401 Unauthorized if access is denied. Instead, it always returns 200 OK with &#x60;Authorized: false&#x60; if the permission is missing. The response will still include the caller’s user information if not authorized.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { HasPermissionRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | Permission
    permission: permission_example,
  } satisfies HasPermissionRequest;

  try {
    const data = await api.hasPermission(body);
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
| **permission** | `string` | Permission | [Defaults to `undefined`] |

### Return type

[**UserSvcHasPermissionResponse**](UserSvcHasPermissionResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Missing Permission |  -  |
| **422** | No Auth Header |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listApps

> UserSvcListAppsResponse listApps(body)

List Apps

List apps. Role, user ID or contact ID must be specified.  Requires the &#x60;user-svc:app:view&#x60; permission, which by default all users have. Caller can only list apps of roles they own (unless they are an admin).

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListAppsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcListAppsRequest | List Apps Request
    body: ...,
  } satisfies ListAppsRequest;

  try {
    const data = await api.listApps(body);
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
| **body** | [UserSvcListAppsRequest](UserSvcListAppsRequest.md) | List Apps Request | |

### Return type

[**UserSvcListAppsResponse**](UserSvcListAppsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Apps listed successfully |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listEnrolls

> UserSvcListEnrollsResponse listEnrolls(body)

List Enrolls

List enrolls. Role, user ID or contact ID must be specified.  Requires the &#x60;user-svc:enroll:view&#x60; permission, which by default all users have. Caller can only list enrolls of roles they own (unless they are an admin).

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListEnrollsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcListEnrollsRequest | List Enrolls Request
    body: ...,
  } satisfies ListEnrollsRequest;

  try {
    const data = await api.listEnrolls(body);
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
| **body** | [UserSvcListEnrollsRequest](UserSvcListEnrollsRequest.md) | List Enrolls Request | |

### Return type

[**UserSvcListEnrollsResponse**](UserSvcListEnrollsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Enrolls listed successfully |  -  |
| **400** | Role, Contact ID or User ID is Required |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listOrganizations

> UserSvcListOrganizationsResponse listOrganizations(body)

List Organizations

Requires the &#x60;user-svc:organization:view&#x60; permission. With &#x60;all&#x3D;true&#x60;, platform admins see all organizations in the current app. Otherwise users only see organizations they are members of.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListOrganizationsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcListOrganizationsRequest | List Organizations Request (optional)
    body: ...,
  } satisfies ListOrganizationsRequest;

  try {
    const data = await api.listOrganizations(body);
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
| **body** | [UserSvcListOrganizationsRequest](UserSvcListOrganizationsRequest.md) | List Organizations Request | [Optional] |

### Return type

[**UserSvcListOrganizationsResponse**](UserSvcListOrganizationsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Organization listed successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listPermissions

> UserSvcListPermissionsResponse listPermissions(roleId)

List Permissions

List permissions by roles. Caller can only list permissions for roles they have.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListPermissionsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | Role ID
    roleId: roleId_example,
  } satisfies ListPermissionsRequest;

  try {
    const data = await api.listPermissions(body);
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
| **roleId** | `string` | Role ID | [Defaults to `undefined`] |

### Return type

[**UserSvcListPermissionsResponse**](UserSvcListPermissionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listPermits

> UserSvcListPermitsResponse listPermits(body)

List Permits

List permits. Requires the &#x60;user-svc:permit:view&#x60; permission, which only admins have by default. &amp;todo Users should be able to list permits referring to them.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListPermitsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcListPermitsRequest | List Permits Request
    body: ...,
  } satisfies ListPermitsRequest;

  try {
    const data = await api.listPermits(body);
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
| **body** | [UserSvcListPermitsRequest](UserSvcListPermitsRequest.md) | List Permits Request | |

### Return type

[**UserSvcListPermitsResponse**](UserSvcListPermitsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listUsers

> UserSvcListUsersResponse listUsers(body)

List Users

Fetches a list of users with optional query filters and pagination. Requires the &#x60;user-svc:user:view&#x60; permission that only admins have by default.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ListUsersRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcListUsersRequest | List Users Request (optional)
    body: ...,
  } satisfies ListUsersRequest;

  try {
    const data = await api.listUsers(body);
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
| **body** | [UserSvcListUsersRequest](UserSvcListUsersRequest.md) | List Users Request | [Optional] |

### Return type

[**UserSvcListUsersResponse**](UserSvcListUsersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List of users retrieved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## login

> UserSvcLoginResponse login(body)

Login

Authenticates a user and returns a token.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { LoginRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  const body = {
    // UserSvcLoginRequest | Login Request
    body: ...,
  } satisfies LoginRequest;

  try {
    const data = await api.login(body);
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
| **body** | [UserSvcLoginRequest](UserSvcLoginRequest.md) | Login Request | |

### Return type

[**UserSvcLoginResponse**](UserSvcLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Login successful |  -  |
| **400** | Invalid JSON |  -  |
| **404** | User Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## readApp

> UserSvcReadAppResponse readApp(body)

Read or Create App

Get an app by host, or create it if it does not exist.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ReadAppRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcReadAppRequest | Read App Request
    body: ...,
  } satisfies ReadAppRequest;

  try {
    const data = await api.readApp(body);
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
| **body** | [UserSvcReadAppRequest](UserSvcReadAppRequest.md) | Read App Request | |

### Return type

[**UserSvcReadAppResponse**](UserSvcReadAppResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | App retrieved or created successfully |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## readSelf

> UserSvcReadSelfResponse readSelf(body)

Read Self

Retrieves user information based on the authentication token in the request header. Typically called by single-page applications during the initial page load. While some details (such as roles, slug, user ID, and active organization ID) can be extracted from the JWT, this endpoint returns additional data, including the full user object and associated organizations.  ReadSelf intentionally still works after token revocation until the token expires. This is to ensure that the user is not notified of token revocation (though some information is leaked by the count token functionality @todo).

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ReadSelfRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcReadSelfRequest | Read Self Request (optional)
    body: ...,
  } satisfies ReadSelfRequest;

  try {
    const data = await api.readSelf(body);
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
| **body** | [UserSvcReadSelfRequest](UserSvcReadSelfRequest.md) | Read Self Request | [Optional] |

### Return type

[**UserSvcReadSelfResponse**](UserSvcReadSelfResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Token Missing |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## refreshToken

> UserSvcRefreshTokenResponse refreshToken()

Refresh Token

Refreshes an existing token, including inactive ones. The old token becomes inactive (if not already inactive), and a new, active token is issued. This allows continued verification of user roles without requiring a new login. Inactive tokens are refreshable unless explicitly revoked (no mechanism for this yet). Leaked tokens should be handled separately, via a revocation flag or deletion.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { RefreshTokenRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  try {
    const data = await api.refreshToken();
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

[**UserSvcRefreshTokenResponse**](UserSvcRefreshTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Refresh Token successful |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## register

> UserSvcRegisterResponse register(body)

Register

Register a new user with a name, email, and password.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { RegisterRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  const body = {
    // UserSvcRegisterRequest | Register Request
    body: ...,
  } satisfies RegisterRequest;

  try {
    const data = await api.register(body);
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
| **body** | [UserSvcRegisterRequest](UserSvcRegisterRequest.md) | Register Request | |

### Return type

[**UserSvcRegisterResponse**](UserSvcRegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid JSON |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## resetPassword

> object resetPassword(userId, body)

Reset Password

Allows an administrator to change a user\&#39;s password.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { ResetPasswordRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | User ID
    userId: userId_example,
    // UserSvcResetPasswordRequest | Change Password Request
    body: ...,
  } satisfies ResetPasswordRequest;

  try {
    const data = await api.resetPassword(body);
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
| **userId** | `string` | User ID | [Defaults to `undefined`] |
| **body** | [UserSvcResetPasswordRequest](UserSvcResetPasswordRequest.md) | Change Password Request | |

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
| **200** | Password changed successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## revokeTokens

> object revokeTokens(body)

Revoke Tokens

Revoke tokens in one of the following scenarios: - For the current user. - For another user (see &#x60;userId&#x60; field), if permitted (&#x60;user-svc:token:revoke&#x60; permission, typically by admins).

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { RevokeTokensRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcRevokeTokensRequest | Revoke Tokens Request (optional)
    body: ...,
  } satisfies RevokeTokensRequest;

  try {
    const data = await api.revokeTokens(body);
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
| **body** | [UserSvcRevokeTokensRequest](UserSvcRevokeTokensRequest.md) | Revoke Tokens Request | [Optional] |

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
| **400** | Mutually Exclusive Parameters |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveEnrolls

> UserSvcSaveEnrollsResponse saveEnrolls(body)

Save Enrolls

Enroll a list of users by contact or user Id to acquire a role. Works on future or current users.  A user can only enroll an other user to a role if the user \&quot;owns\&quot; that role. A user who owns a role can enroll others in that roll in any app. The same request might contain enrolls for different apps.  A user \&quot;owns\&quot; a role in the following cases: - A static role where the role ID is prefixed with the caller\&#39;s slug. - Any dynamic or static role where the caller is an admin (has &#x60;*:admin&#x60; postfix of that role).  Examples: - A user with the slug &#x60;joe-doe&#x60; owns roles like &#x60;joe-doe:*&#x60; such as &#x60;joe-doe:any-custom-role&#x60;. - A user with any slug who has the role &#x60;my-service:admin&#x60; owns &#x60;my-service:*&#x60; roles such as &#x60;my-service:user&#x60;. - A user with any slug who has the role &#x60;user-svc:org:{%orgId}:admin&#x60; owns &#x60;user-svc:org:{%orgId}:*&#x60; such as &#x60;user-svc:org:{%orgId}:user&#x60;.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SaveEnrollsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcSaveEnrollsRequest | Save Enrolls Request
    body: ...,
  } satisfies SaveEnrollsRequest;

  try {
    const data = await api.saveEnrolls(body);
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
| **body** | [UserSvcSaveEnrollsRequest](UserSvcSaveEnrollsRequest.md) | Save Enrolls Request | |

### Return type

[**UserSvcSaveEnrollsResponse**](UserSvcSaveEnrollsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Enrolls saved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveMembership

> object saveMembership(organizationId, userId, body)

Save Membership

Adds a user to an organization by saving a Membership. Also issues the corresponding Enroll, which grants the user their dynamic organization role (e.g. &#x60;user-svc:org:{org_123}:user&#x60;).

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SaveMembershipRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | Organization ID
    organizationId: organizationId_example,
    // string | User ID
    userId: userId_example,
    // object | Add User to Organization Request (optional)
    body: Object,
  } satisfies SaveMembershipRequest;

  try {
    const data = await api.saveMembership(body);
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
| **organizationId** | `string` | Organization ID | [Defaults to `undefined`] |
| **userId** | `string` | User ID | [Defaults to `undefined`] |
| **body** | `object` | Add User to Organization Request | [Optional] |

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
| **200** | User added successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **403** | Forbidden |  -  |
| **404** | Organization/User not found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveOrganization

> UserSvcSaveOrganizationResponse saveOrganization(body)

Save an Organization

Allows a logged-in user to save an organization. The user initiating the request will be assigned the role of admin for that organization. The initiating user will receive a dynamic role in the format &#x60;user-svc:org:{organizationId}:admin&#x60;, where &#x60;{organizationId}&#x60; is a unique identifier for the saved organization. Dynamic roles are generated based on specific user-resource associations (in this case the resource being the organization), offering more flexible permission management compared to static roles.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SaveOrganizationRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcSaveOrganizationRequest | Save User Request
    body: ...,
  } satisfies SaveOrganizationRequest;

  try {
    const data = await api.saveOrganization(body);
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
| **body** | [UserSvcSaveOrganizationRequest](UserSvcSaveOrganizationRequest.md) | Save User Request | |

### Return type

[**UserSvcSaveOrganizationResponse**](UserSvcSaveOrganizationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | User saved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## savePermits

> object savePermits(body)

Save Permits

Save permits. Permits give access to users with certain slugs and roles to permissions.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SavePermitsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcSavePermitsRequest | Save Permits Request
    body: ...,
  } satisfies SavePermitsRequest;

  try {
    const data = await api.savePermits(body);
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
| **body** | [UserSvcSavePermitsRequest](UserSvcSavePermitsRequest.md) | Save Permits Request | |

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
| **200** | Permit saved successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## saveSelf

> object saveSelf(body)

Save User Profile

Save user\&#39;s own profile information.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SaveSelfRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcSaveSelfRequest | Save Profile Request
    body: ...,
  } satisfies SaveSelfRequest;

  try {
    const data = await api.saveSelf(body);
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
| **body** | [UserSvcSaveSelfRequest](UserSvcSaveSelfRequest.md) | Save Profile Request | |

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


## saveUser

> object saveUser(userId, body)

Save User

Save user information based on the provided user ID. Intended for admins. Requires the &#x60;user-svc:user:edit&#x60; permission. For a user to edit their own profile, see &#x60;saveSelf&#x60;.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SaveUserRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // string | User ID
    userId: userId_example,
    // UserSvcSaveUserRequest | Save Profile Request
    body: ...,
  } satisfies SaveUserRequest;

  try {
    const data = await api.saveUser(body);
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
| **userId** | `string` | User ID | [Defaults to `undefined`] |
| **body** | [UserSvcSaveUserRequest](UserSvcSaveUserRequest.md) | Save Profile Request | |

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


## sendOtp

> UserSvcSendOtpResponse sendOtp(body, acceptLanguage)

Send OTP

Generates and sends a one-time password (OTP) to the specified contact.  The OTP can be used for contact verification or login depending on purpose.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { SendOtpRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new UserSvcApi();

  const body = {
    // UserSvcSendOtpRequest | Send OTP Request
    body: ...,
    // string | Language preference for the email (optional)
    acceptLanguage: acceptLanguage_example,
  } satisfies SendOtpRequest;

  try {
    const data = await api.sendOtp(body);
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
| **body** | [UserSvcSendOtpRequest](UserSvcSendOtpRequest.md) | Send OTP Request | |
| **acceptLanguage** | `string` | Language preference for the email | [Optional] [Defaults to `undefined`] |

### Return type

[**UserSvcSendOtpResponse**](UserSvcSendOtpResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OTP sent successfully |  -  |
| **400** | Invalid request |  -  |
| **404** | Contact not found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## updateApp

> object updateApp(body)

Update App Host

Change the hostname of an existing app. Requires the &#x60;user-svc:app:edit&#x60; permission.

### Example

```ts
import {
  Configuration,
  UserSvcApi,
} from '';
import type { UpdateAppRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new UserSvcApi(config);

  const body = {
    // UserSvcUpdateAppRequest | Update App Request
    body: ...,
  } satisfies UpdateAppRequest;

  try {
    const data = await api.updateApp(body);
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
| **body** | [UserSvcUpdateAppRequest](UserSvcUpdateAppRequest.md) | Update App Request | |

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
| **200** | App updated successfully |  -  |
| **400** | Invalid Request |  -  |
| **404** | App Not Found |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

