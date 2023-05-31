# \DefinedInterfaceBehaviorsApi

All URIs are relative to *https://localhost/cloudapi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDefinedEntityTypeAccess**](DefinedInterfaceBehaviorsApi.md#AddDefinedEntityTypeAccess) | **Post** /1.0.0/entityTypes/{id}/behaviorAccessControls | Adds an access control configuration of an entity type&#39;s behavior
[**AddInterfaceBehavior**](DefinedInterfaceBehaviorsApi.md#AddInterfaceBehavior) | **Post** /1.0.0/interfaces/{id}/behaviors | Add a new interface Behavior if the Interface is not in use
[**DeleteInterfaceBehavior**](DefinedInterfaceBehaviorsApi.md#DeleteInterfaceBehavior) | **Delete** /1.0.0/interfaces/{id}/behaviors/{behaviorId} | Remove a Behavior from the Defined Interface.
[**GetBehaviorExecutionLog**](DefinedInterfaceBehaviorsApi.md#GetBehaviorExecutionLog) | **Get** /1.0.0/entities/{id}/behaviors/{behaviorId}/invocations/{invocationId}/log | Download the execution log from a behavior invocation
[**GetDefinedEntityTypeAccess**](DefinedInterfaceBehaviorsApi.md#GetDefinedEntityTypeAccess) | **Get** /1.0.0/entityTypes/{id}/behaviorAccessControls | Gets the access control configuration of the entity type&#39;s behaviors
[**GetDefinedEntityTypeBehaviors**](DefinedInterfaceBehaviorsApi.md#GetDefinedEntityTypeBehaviors) | **Get** /1.0.0/entityTypes/{id}/behaviors | Get the Behaviors of the Defined Entity Type.
[**GetInterfaceBehavior**](DefinedInterfaceBehaviorsApi.md#GetInterfaceBehavior) | **Get** /1.0.0/interfaces/{id}/behaviors/{behaviorId} | Get a Behavior in the Defined Interface.
[**GetInterfaceBehaviors**](DefinedInterfaceBehaviorsApi.md#GetInterfaceBehaviors) | **Get** /1.0.0/interfaces/{id}/behaviors | Get the Behaviors of the Defined Interface.
[**GetTypeBehavior**](DefinedInterfaceBehaviorsApi.md#GetTypeBehavior) | **Get** /1.0.0/entityTypes/{id}/behaviors/{behaviorId} | Get a Behavior in the Defined Type
[**InvokeDefinedEntityBehavior**](DefinedInterfaceBehaviorsApi.md#InvokeDefinedEntityBehavior) | **Post** /1.0.0/entities/{id}/behaviors/{behaviorId}/invocations | Invokes a behavior on a defined entity
[**RemoveBehaviorOverride**](DefinedInterfaceBehaviorsApi.md#RemoveBehaviorOverride) | **Delete** /1.0.0/entityTypes/{id}/behaviors/{behaviorId} | Remove a Behavior override from the Defined Entity Type.
[**SetDefinedEntityTypeAccess**](DefinedInterfaceBehaviorsApi.md#SetDefinedEntityTypeAccess) | **Put** /1.0.0/entityTypes/{id}/behaviorAccessControls | Sets the access control configuration of the entity type&#39;s behaviors
[**UpdateInterfaceBehavior**](DefinedInterfaceBehaviorsApi.md#UpdateInterfaceBehavior) | **Put** /1.0.0/interfaces/{id}/behaviors/{behaviorId} | Update the execution of the specified Behavior in the Defined Interface
[**UpdateInterfaceBehaviors**](DefinedInterfaceBehaviorsApi.md#UpdateInterfaceBehaviors) | **Put** /1.0.0/interfaces/{id}/behaviors | Update all the executions of the specified Behaviors (possibly removing or adding some)
[**UpdateTypeBehavior**](DefinedInterfaceBehaviorsApi.md#UpdateTypeBehavior) | **Put** /1.0.0/entityTypes/{id}/behaviors/{behaviorId} | Update the execution of the specified Behavior in the Defined Entity Type


# **AddDefinedEntityTypeAccess**
> BehaviorAccess AddDefinedEntityTypeAccess(ctx, definition, id)
Adds an access control configuration of an entity type's behavior

Adds an access control configuration of an entity type's behavior

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definition** | [**BehaviorAccess**](BehaviorAccess.md)|  | 
  **id** | **string**|  | 

### Return type

[**BehaviorAccess**](BehaviorAccess.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddInterfaceBehavior**
> Behavior AddInterfaceBehavior(ctx, behavior, id)
Add a new interface Behavior if the Interface is not in use

Add a new Behavior to the Interface. Only allowed if the Interface is not in use. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **behavior** | [**Behavior**](Behavior.md)|  | 
  **id** | **string**|  | 

### Return type

[**Behavior**](Behavior.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInterfaceBehavior**
> DeleteInterfaceBehavior(ctx, id, behaviorId)
Remove a Behavior from the Defined Interface.

Remove a Behavior from the Defined Interface. The Behaviors can be specified by ID or by name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBehaviorExecutionLog**
> string GetBehaviorExecutionLog(ctx, id, behaviorId, invocationId)
Download the execution log from a behavior invocation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 
  **invocationId** | **string**|  | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntityTypeAccess**
> BehaviorAccesses GetDefinedEntityTypeAccess(ctx, page, pageSize, id)
Gets the access control configuration of the entity type's behaviors

Gets the access control configuration of the entity type's behaviors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 

### Return type

[**BehaviorAccesses**](BehaviorAccesses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDefinedEntityTypeBehaviors**
> Behaviors GetDefinedEntityTypeBehaviors(ctx, page, pageSize, id)
Get the Behaviors of the Defined Entity Type.

Retrieve the Behaviors of the specified Defined Entity Type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 

### Return type

[**Behaviors**](Behaviors.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterfaceBehavior**
> Behavior GetInterfaceBehavior(ctx, id, behaviorId)
Get a Behavior in the Defined Interface.

Retrieve a specific Behavior in the specified Defined Interface. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

[**Behavior**](Behavior.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterfaceBehaviors**
> Behaviors GetInterfaceBehaviors(ctx, page, pageSize, id)
Get the Behaviors of the Defined Interface.

Retrieve the Behaviors of the specified Defined Interface. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Page to fetch, zero offset. | [default to 1]
  **pageSize** | **int32**| Results per page to fetch. | [default to 25]
  **id** | **string**|  | 

### Return type

[**Behaviors**](Behaviors.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTypeBehavior**
> Behavior GetTypeBehavior(ctx, id, behaviorId)
Get a Behavior in the Defined Type

Retrieve a specific Behavior in the Defined Type. The Behavior must be specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

[**Behavior**](Behavior.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeDefinedEntityBehavior**
> InvokeDefinedEntityBehavior(ctx, id, behaviorId, optional)
Invokes a behavior on a defined entity

Invokes a behavior on a defined entity. The contract of the behavior is specified in the behavior description. If an Activity behavior is invoked with an 'operationId' in the invocation metadata, then another invocation of the behavior with the same 'operationId' will be ignored within the next 1 hour. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 
 **optional** | ***DefinedInterfaceBehaviorsApiInvokeDefinedEntityBehaviorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefinedInterfaceBehaviorsApiInvokeDefinedEntityBehaviorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invocation** | [**optional.Interface of BehaviorInvocation**](BehaviorInvocation.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveBehaviorOverride**
> RemoveBehaviorOverride(ctx, id, behaviorId)
Remove a Behavior override from the Defined Entity Type.

Remove a Behavior override in the Defined Entity Type. The Behavior must be specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: *_/_*;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetDefinedEntityTypeAccess**
> BehaviorAccesses SetDefinedEntityTypeAccess(ctx, definition, id)
Sets the access control configuration of the entity type's behaviors

Sets the access control configuration of the entity type's behaviors

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **definition** | [**BehaviorAccesses**](BehaviorAccesses.md)|  | 
  **id** | **string**|  | 

### Return type

[**BehaviorAccesses**](BehaviorAccesses.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInterfaceBehavior**
> Behavior UpdateInterfaceBehavior(ctx, behavior, id, behaviorId)
Update the execution of the specified Behavior in the Defined Interface

Update the execution of the specified Behavior in the Defined Interface. The Behaviors can be specified by ID or by name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **behavior** | [**Behavior**](Behavior.md)|  | 
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

[**Behavior**](Behavior.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInterfaceBehaviors**
> Behaviors UpdateInterfaceBehaviors(ctx, behaviors, id)
Update all the executions of the specified Behaviors (possibly removing or adding some)

Update all Behaviors, possibly adding or removing some if the Interface is not in use. If the Interface is in use, then only the executions of the existing Behaviors can be updated. The Behaviors can be specified by ID or by name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **behaviors** | [**Behaviors**](Behaviors.md)|  | 
  **id** | **string**|  | 

### Return type

[**Behaviors**](Behaviors.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTypeBehavior**
> Behavior UpdateTypeBehavior(ctx, behavior, id, behaviorId)
Update the execution of the specified Behavior in the Defined Entity Type

Override the execution of the specified Behavior in the Defined Entity Type. The Behavior must be specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **behavior** | [**Behavior**](Behavior.md)|  | 
  **id** | **string**|  | 
  **behaviorId** | **string**|  | 

### Return type

[**Behavior**](Behavior.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;version=37.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

