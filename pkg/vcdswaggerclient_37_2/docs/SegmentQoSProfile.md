# SegmentQoSProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of the segment profile. | [optional] [default to null]
**DisplayName** | **string** | Name of the segment profile. This corresponds to the name used in NSX-T manager&#39;s logs or GUI. | [optional] [default to null]
**Description** | **string** | The description of the segment profile. | [optional] [default to null]
**NsxTManagerRef** | [***EntityReference**](EntityReference.md) | The NSX-T manager where this segment profile is configured. | [optional] [default to null]
**ClassOfService** | **int32** | Class of service groups similar types of traffic in the network and each type of traffic is treated as a class with its own level of service priority. The lower priority traffic is slowed down or in some cases dropped to provide better throughput for higher priority traffic.  | [optional] [default to null]
**DscpConfig** | [***QoSDscpConfig**](QoSDscpConfig.md) | A Differentiated Services Code Point (DSCP) Configuration for this Segment QoS Profile. | [optional] [default to null]
**EgressRateLimiter** | [***QoSRateLimiterShaper**](QoSRateLimiterShaper.md) | Egress rate properties in Mb/s. | [optional] [default to null]
**IngressBroadcastRateLimiter** | [***QoSRateLimiterShaper**](QoSRateLimiterShaper.md) | Ingress broadcast rate properties in Mb/s. | [optional] [default to null]
**IngressRateLimiter** | [***QoSRateLimiterShaper**](QoSRateLimiterShaper.md) | Ingress rate properties in Mb/s. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


