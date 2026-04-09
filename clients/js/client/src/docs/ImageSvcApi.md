# ImageSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**serveDownloadedImage**](ImageSvcApi.md#servedownloadedimage) | **GET** /image-svc/serve/download/{url} | Serve Downloaded Image |
| [**serveUploadedImage**](ImageSvcApi.md#serveuploadedimage) | **GET** /image-svc/serve/upload/{fileId} | Serve Uploaded Image |



## serveDownloadedImage

> Blob serveDownloadedImage(url, width, height, quality, format, fit, position)

Serve Downloaded Image

Retrieves, caches, resizes, and serves an image referenced by its original URL.

### Example

```ts
import {
  Configuration,
  ImageSvcApi,
} from '';
import type { ServeDownloadedImageRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ImageSvcApi();

  const body = {
    // string | Original URL of the downloaded file (path-escaped)
    url: url_example,
    // number | Optional width to resize the image to (optional)
    width: 56,
    // number | Optional height to resize the image to (optional)
    height: 56,
    // number | Optional quality for lossy output formats (default 85) (optional)
    quality: 56,
    // string | Optional output format: webp, jpeg, png, gif, avif (optional)
    format: format_example,
    // string | Resize strategy: contain|cover (default contain) (optional)
    fit: fit_example,
    // string | Crop anchor when fit=cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right (optional)
    position: position_example,
  } satisfies ServeDownloadedImageRequest;

  try {
    const data = await api.serveDownloadedImage(body);
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
| **url** | `string` | Original URL of the downloaded file (path-escaped) | [Defaults to `undefined`] |
| **width** | `number` | Optional width to resize the image to | [Optional] [Defaults to `undefined`] |
| **height** | `number` | Optional height to resize the image to | [Optional] [Defaults to `undefined`] |
| **quality** | `number` | Optional quality for lossy output formats (default 85) | [Optional] [Defaults to `undefined`] |
| **format** | `string` | Optional output format: webp, jpeg, png, gif, avif | [Optional] [Defaults to `undefined`] |
| **fit** | `string` | Resize strategy: contain|cover (default contain) | [Optional] [Defaults to `undefined`] |
| **position** | `string` | Crop anchor when fit&#x3D;cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right | [Optional] [Defaults to `undefined`] |

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
| **200** | Image served successfully |  -  |
| **400** | Invalid URL |  -  |
| **404** | File Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## serveUploadedImage

> Blob serveUploadedImage(fileId, width, height, fit, position)

Serve Uploaded Image

Retrieves and serves a previously uploaded image file using its File ID.

### Example

```ts
import {
  Configuration,
  ImageSvcApi,
} from '';
import type { ServeUploadedImageRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const api = new ImageSvcApi();

  const body = {
    // string | FileID uniquely identifies the file itself (not an ID, which represents a specific replica)
    fileId: fileId_example,
    // number | Optional width to resize the image to (optional)
    width: 56,
    // number | Optional height to resize the image to (optional)
    height: 56,
    // string | Resize strategy: contain|cover (default contain) (optional)
    fit: fit_example,
    // string | Crop anchor when fit=cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right (optional)
    position: position_example,
  } satisfies ServeUploadedImageRequest;

  try {
    const data = await api.serveUploadedImage(body);
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
| **width** | `number` | Optional width to resize the image to | [Optional] [Defaults to `undefined`] |
| **height** | `number` | Optional height to resize the image to | [Optional] [Defaults to `undefined`] |
| **fit** | `string` | Resize strategy: contain|cover (default contain) | [Optional] [Defaults to `undefined`] |
| **position** | `string` | Crop anchor when fit&#x3D;cover: center|top|bottom|left|right|top-left|top-right|bottom-left|bottom-right | [Optional] [Defaults to `undefined`] |

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
| **400** | Missing File ID |  -  |
| **404** | File Not Found |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

