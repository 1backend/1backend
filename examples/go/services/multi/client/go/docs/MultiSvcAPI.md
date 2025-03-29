# \MultiSvcAPI

All URIs are relative to *http://localhost:58231*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountPets**](MultiSvcAPI.md#CountPets) | **Get** /multi-svc/pets/count | Count Pets



## CountPets

> MultiSvcCountPetsResponse CountPets(ctx).Body(body).Execute()

Count Pets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/1backend/1backend/examples/go/services/multi/client"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | Count Pets Request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultiSvcAPI.CountPets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultiSvcAPI.CountPets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountPets`: MultiSvcCountPetsResponse
	fmt.Fprintf(os.Stdout, "Response from `MultiSvcAPI.CountPets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Count Pets Request | 

### Return type

[**MultiSvcCountPetsResponse**](MultiSvcCountPetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

