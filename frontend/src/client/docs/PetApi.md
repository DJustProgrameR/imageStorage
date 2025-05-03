# PetApi

All URIs are relative to *http://localhost:8000/v1*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getPet**](#getpet) | **GET** /pet | |
|[**uploadPet**](#uploadpet) | **PUT** /pet | |

# **getPet**
> GetPet200Response getPet()

Get existing pet

### Example

```typescript
import {
    PetApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new PetApi(configuration);

const { status, data } = await apiInstance.getPet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**GetPet200Response**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**204** | No pet uploaded yet |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **uploadPet**
> uploadPet(petProperties)

Upload new pet

### Example

```typescript
import {
    PetApi,
    Configuration,
    GetPet200Response
} from './api';

const configuration = new Configuration();
const apiInstance = new PetApi(configuration);

let petProperties: GetPet200Response; //Pet properties

const { status, data } = await apiInstance.uploadPet(
    petProperties
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **petProperties** | **GetPet200Response**| Pet properties | |


### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |
|**400** | Bad request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

