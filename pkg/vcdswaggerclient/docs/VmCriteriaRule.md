# VmCriteriaRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeType** | **string** | The attribute type of a VM used for VM matching.  Below are supported types: &lt;ul&gt; &lt;li&gt; VM_TAG - Match the VM based on the tags associated with that VM. &lt;li&gt; VM_NAME - Match the VM based on the name of the VM. Only CONTAINS and STARTS_WITH operators are supported for this type. &lt;/ul&gt;  | [default to null]
**AttributeValue** | **string** | The attribute value that is used to determine if a VM&#39;s attribute value matches the rule. Example: if the attribute type is VM_NAME, user should set this value to the name of the VM to match with.  | [default to null]
**Operator** | **string** | The operator to perform to determine whether the rule&#39;s attribute value matches a VM&#39;s attribute value.  Example: if the attribute type is VM_NAME, user can set this operator to determine whether a VM&#39;s name must be an exact match or starts with that name. Below are supported types: &lt;ul&gt; &lt;li&gt; EQUALS - Match occurs if the VM&#39;s attribute value is exactly the same as the rule&#39;s attribute value. &lt;li&gt; CONTAINS - Match occurs if the VM&#39;s attribute value is contains the rule&#39;s attribute value. &lt;li&gt; STARTS_WITH - Match occurs if the VM&#39;s attribute value starts with the rule&#39;s attribute value. &lt;li&gt; ENDS_WITH - Match occurs if the VM&#39;s attribute value ends with the rule&#39;s attribute value. &lt;/ul&gt;  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


