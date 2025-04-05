# \BasicSvcAPI

All URIs are relative to *http://localhost:11337*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPets**](BasicSvcAPI.md#ListPets) | **Post** /basic-svc/pets | List Pets
[**SavePet**](BasicSvcAPI.md#SavePet) | **Put** /basic-svc/pet | Save Pet



## ListPets

> BasicSvcListPetsResponse ListPets(ctx).Body(body).Execute()

List Pets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/1backend/1backend/examples/go/services/basic/client"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | Registration Tracking Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicSvcAPI.ListPets(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicSvcAPI.ListPets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPets`: BasicSvcListPetsResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicSvcAPI.ListPets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Registration Tracking Request | 

### Return type

[**BasicSvcListPetsResponse**](BasicSvcListPetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavePet

> map[string]interface{} SavePet(ctx).Body(body).Execute()

Save Pet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/1backend/1backend/examples/go/services/basic/client"
)

func main() {
	body := *openapiclient.NewBasicSvcSavePetRequest("Name_example") // BasicSvcSavePetRequest | Registration Tracking Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicSvcAPI.SavePet(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicSvcAPI.SavePet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SavePet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BasicSvcAPI.SavePet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BasicSvcSavePetRequest**](BasicSvcSavePetRequest.md) | Registration Tracking Request | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

