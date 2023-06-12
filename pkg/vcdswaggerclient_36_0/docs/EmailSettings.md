# EmailSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSmtpServer** | **bool** | Flag indicating if the smtp server settings are system default (true) or a particular organization (false)  | [optional] [default to null]
**DefaultOrgEmail** | **bool** | Flag indicating if the email settings are system default (true) or for a particular organization (false)  | [optional] [default to null]
**SmtpServer** | [***SmtpServerSpec**](SmtpServerSpec.md) |  | [optional] [default to null]
**SenderEmailAddress** | **string** | Sender email address in an email notification or alert  | [optional] [default to null]
**DefaultEmailSubjectPrefix** | **string** | The prefix used in the email subject line for all email notifications and alerts from the system  | [optional] [default to null]
**AlertEmailToAllAdmins** | **bool** | A flag to indicate the choice between sending alert emails to all system administrators in the system and designated list of email recipients  | [optional] [default to null]
**AlertEmailTo** | **string** | A comma separated email addresses to send all alert messages to  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


