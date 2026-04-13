# EmailSvcApi

All URIs are relative to *http://localhost:11337*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**sendEmail**](EmailSvcApi.md#sendemail) | **POST** /email-svc/email | Send an Email |



## sendEmail

> EmailSvcSendEmailResponse sendEmail(body)

Send an Email

Sends an email with optional attachments via a supported email provider.  Currently, only SendGrid is supported. Additional providers may be added in the future.  Required secrets from the Secret Svc for SendGrid: - &#x60;sender-email&#x60;: Sender\&#39;s email address. - &#x60;sender-name&#x60;: Sender\&#39;s display name. - &#x60;sendgrid-api-key&#x60;: API key for SendGrid.

### Example

```ts
import {
  Configuration,
  EmailSvcApi,
} from '';
import type { SendEmailRequest } from '';

async function example() {
  console.log("🚀 Testing  SDK...");
  const config = new Configuration({ 
    // To configure API key authorization: BearerAuth
    apiKey: "YOUR API KEY",
  });
  const api = new EmailSvcApi(config);

  const body = {
    // EmailSvcSendEmailRequest | Send Email Request
    body: ...,
  } satisfies SendEmailRequest;

  try {
    const data = await api.sendEmail(body);
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
| **body** | [EmailSvcSendEmailRequest](EmailSvcSendEmailRequest.md) | Send Email Request | |

### Return type

[**EmailSvcSendEmailResponse**](EmailSvcSendEmailResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Successfully sent the email |  -  |
| **400** | Invalid JSON |  -  |
| **401** | Unauthorized |  -  |
| **500** | Internal Server Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

