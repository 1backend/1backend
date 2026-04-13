# FileSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteUpload**](FileSvcApi.md#deleteupload) | **DELETE** /file-svc/upload/{fileId} | Delete an Uploaded File |
| [**downloadFile**](FileSvcApi.md#downloadfile) | **PUT** /file-svc/download | Download a File |
| [**getDownload**](FileSvcApi.md#getdownload) | **GET** /file-svc/download/{url} | Get a Download |
| [**listDownloads**](FileSvcApi.md#listdownloads) | **POST** /file-svc/downloads | List Downloads |
| [**listUploads**](FileSvcApi.md#listuploads) | **POST** /file-svc/uploads | List Uploads |
| [**pauseDownload**](FileSvcApi.md#pausedownload) | **PUT** /file-svc/download/{url}/pause | Pause a Download |
| [**serveDownload**](FileSvcApi.md#servedownload) | **GET** /file-svc/serve/download/{url} | Serve a Downloaded file |
| [**serveUpload**](FileSvcApi.md#serveupload) | **GET** /file-svc/serve/upload/{fileId} | Serve an Uploaded File |
| [**uploadFile**](FileSvcApi.md#uploadfile) | **PUT** /file-svc/upload | Upload a File |



## deleteUpload

> { [key: string]: any; } deleteUpload(fileId)

Delete an Uploaded File

Deletes an uploaded file and its metadata by &#x60;fileId&#x60;.  Requires the &#x60;file-svc:upload:delete&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { DeleteUploadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // string | File ID
    fileId: fileId_example,
  } satisfies DeleteUploadRequest;

  try {
    const data = await api.deleteUpload(body);
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
| **fileId** | `string` | File ID | [Defaults to `undefined`] |

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
| **200** | File deleted successfully |  -  |
| **400** | invalid request |  -  |
| **401** | Unauthorized |  -  |
| **404** | File not found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## downloadFile

> { [key: string]: any; } downloadFile(body)

Download a File

Start or resume the download for a specified URL.  Requires the &#x60;file-svc:download:create&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { DownloadFileRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // FileSvcDownloadFileRequest | Download Request
    body: ...,
  } satisfies DownloadFileRequest;

  try {
    const data = await api.downloadFile(body);
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
| **body** | [FileSvcDownloadFileRequest](FileSvcDownloadFileRequest.md) | Download Request | |

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
| **200** | Download initiated successfully |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Failed to download file |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getDownload

> FileSvcGetDownloadResponse getDownload(url)

Get a Download

Get a download by URL.  Requires the &#x60;file-svc:download:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { GetDownloadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // string | url
    url: url_example,
  } satisfies GetDownloadRequest;

  try {
    const data = await api.getDownload(body);
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
| **url** | `string` | url | [Defaults to `undefined`] |

### Return type

[**FileSvcGetDownloadResponse**](FileSvcGetDownloadResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |
| **400** | Invalid URL |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listDownloads

> FileSvcDownloadsResponse listDownloads()

List Downloads

List download details.  Requires the &#x60;file-svc:download:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { ListDownloadsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  try {
    const data = await api.listDownloads();
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

[**FileSvcDownloadsResponse**](FileSvcDownloadsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List of downloads |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listUploads

> FileSvcListUploadsResponse listUploads(body)

List Uploads

Lists uploaded files, returning only metadata about each upload. To retrieve file content, use the &#x60;Serve an Uploaded File&#x60; endpoint, which serves a single file per request. Note: Retrieving the contents of multiple files in a single request is not supported currently.  Requires the &#x60;file-svc:upload:view&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { ListUploadsRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // FileSvcListUploadsRequest | List Uploads Request (optional)
    body: ...,
  } satisfies ListUploadsRequest;

  try {
    const data = await api.listUploads(body);
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
| **body** | [FileSvcListUploadsRequest](FileSvcListUploadsRequest.md) | List Uploads Request | [Optional] |

### Return type

[**FileSvcListUploadsResponse**](FileSvcListUploadsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | List of uploads |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## pauseDownload

> { [key: string]: any; } pauseDownload(url)

Pause a Download

Pause a download that is currently in progress.  Requires the &#x60;file-svc:download:edit&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { PauseDownloadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // string | Download URL
    url: url_example,
  } satisfies PauseDownloadRequest;

  try {
    const data = await api.pauseDownload(body);
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
| **url** | `string` | Download URL | [Defaults to `undefined`] |

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
| **200** | Success response |  -  |
| **400** | Download ID in path is not URL encoded |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## serveDownload

> Blob serveDownload(url)

Serve a Downloaded file

Serves a previously downloaded file based on its URL.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { ServeDownloadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new FileSvcApi();

  const body = {
    // string | URL of the file. Even after downloading, the file is still referenced by its original internet URL.
    url: url_example,
  } satisfies ServeDownloadRequest;

  try {
    const data = await api.serveDownload(body);
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
| **url** | `string` | URL of the file. Even after downloading, the file is still referenced by its original internet URL. | [Defaults to `undefined`] |

### Return type

**Blob**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/octet-stream`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | File served successfully |  -  |
| **400** | Error Parsing Download URL |  -  |
| **404** | File Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## serveUpload

> Blob serveUpload(fileId)

Serve an Uploaded File

Retrieves and serves a previously uploaded file using its File ID. Note: The &#x60;ID&#x60; and &#x60;FileID&#x60; fields of an upload are different. - &#x60;FileID&#x60; is a unique identifier for the file itself. - &#x60;ID&#x60; is a unique identifier for a specific replica of the file. Since 1Backend is a distributed system, files can be replicated across multiple nodes. This means each uploaded file may have multiple records with the same &#x60;FileID&#x60; but different &#x60;ID&#x60;s.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { ServeUploadRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new FileSvcApi();

  const body = {
    // string | FileID uniquely identifies the file itself (not an ID, which represents a specific replica)
    fileId: fileId_example,
  } satisfies ServeUploadRequest;

  try {
    const data = await api.serveUpload(body);
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
| **fileId** | `string` | FileID uniquely identifies the file itself (not an ID, which represents a specific replica) | [Defaults to `undefined`] |

### Return type

**Blob**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/octet-stream`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | File served successfully |  -  |
| **400** | Missing Upload ID |  -  |
| **404** | File Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## uploadFile

> FileSvcUploadFileResponse uploadFile(file)

Upload a File

Uploads a file to the server. Currently if using the clients only one file can be uploaded at a time due to this bug https://github.com/OpenAPITools/openapi-generator/issues/11341 Once that is fixed we should have an &#x60;PUT /file-svc/uploads&#x60;/uploadFiles (note the plural) endpoints. In reality the endpoint \&quot;unofficially\&quot; supports multiple files. YMMV.  Requires the &#x60;file-svc:upload:create&#x60; permission.

### Example

```ts
import {
  Configuration,
  FileSvcApi,
} from '';
import type { UploadFileRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new FileSvcApi(config);

  const body = {
    // Blob | File to upload
    file: BINARY_DATA_HERE,
  } satisfies UploadFileRequest;

  try {
    const data = await api.uploadFile(body);
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
| **file** | `Blob` | File to upload | [Defaults to `undefined`] |

### Return type

[**FileSvcUploadFileResponse**](FileSvcUploadFileResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `multipart/form-data`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | File uploaded successfully |  -  |
| **400** | invalid request |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

