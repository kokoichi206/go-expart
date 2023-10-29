# \PetAPI

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPet**](PetAPI.md#AddPet) | **Post** /pet | Add a new pet to the store
[**DeletePet**](PetAPI.md#DeletePet) | **Delete** /pet/{petId} | Deletes a pet
[**GetPetById**](PetAPI.md#GetPetById) | **Get** /pet/{petId} | Find pet by ID
[**UpdatePet**](PetAPI.md#UpdatePet) | **Post** /pet/{petId} | Updates a pet in the store



## AddPet

> Pet AddPet(ctx).Pet(pet).Execute()

Add a new pet to the store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pet := *openapiclient.NewPet("doggie") // Pet | Create a new pet in the store

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PetAPI.AddPet(context.Background()).Pet(pet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetAPI.AddPet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPet`: Pet
    fmt.Fprintf(os.Stdout, "Response from `PetAPI.AddPet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pet** | [**Pet**](Pet.md) | Create a new pet in the store | 

### Return type

[**Pet**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePet

> DeletePet(ctx, petId).Execute()

Deletes a pet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    petId := int64(789) // int64 | Pet id to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PetAPI.DeletePet(context.Background(), petId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetAPI.DeletePet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | Pet id to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPetById

> Pet GetPetById(ctx, petId).Execute()

Find pet by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    petId := int64(789) // int64 | ID of pet to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PetAPI.GetPetById(context.Background(), petId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetAPI.GetPetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPetById`: Pet
    fmt.Fprintf(os.Stdout, "Response from `PetAPI.GetPetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | ID of pet to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Pet**](Pet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePet

> UpdatePet(ctx, petId).Name(name).Status(status).Execute()

Updates a pet in the store



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    petId := int64(789) // int64 | ID of pet that needs to be updated
    name := "name_example" // string | Name of pet that needs to be updated (optional)
    status := openapiclient.PetStatus("available") // PetStatus | Status of pet that needs to be updated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PetAPI.UpdatePet(context.Background(), petId).Name(name).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PetAPI.UpdatePet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**petId** | **int64** | ID of pet that needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of pet that needs to be updated | 
 **status** | [**PetStatus**](PetStatus.md) | Status of pet that needs to be updated | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

