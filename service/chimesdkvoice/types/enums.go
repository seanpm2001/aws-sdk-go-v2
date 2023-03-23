// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AlexaSkillStatus string

// Enum values for AlexaSkillStatus
const (
	AlexaSkillStatusActive   AlexaSkillStatus = "ACTIVE"
	AlexaSkillStatusInactive AlexaSkillStatus = "INACTIVE"
)

// Values returns all known values for AlexaSkillStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlexaSkillStatus) Values() []AlexaSkillStatus {
	return []AlexaSkillStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type CallingNameStatus string

// Enum values for CallingNameStatus
const (
	CallingNameStatusUnassigned       CallingNameStatus = "Unassigned"
	CallingNameStatusUpdateInProgress CallingNameStatus = "UpdateInProgress"
	CallingNameStatusUpdateSucceeded  CallingNameStatus = "UpdateSucceeded"
	CallingNameStatusUpdateFailed     CallingNameStatus = "UpdateFailed"
)

// Values returns all known values for CallingNameStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CallingNameStatus) Values() []CallingNameStatus {
	return []CallingNameStatus{
		"Unassigned",
		"UpdateInProgress",
		"UpdateSucceeded",
		"UpdateFailed",
	}
}

type Capability string

// Enum values for Capability
const (
	CapabilityVoice Capability = "Voice"
	CapabilitySms   Capability = "SMS"
)

// Values returns all known values for Capability. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Capability) Values() []Capability {
	return []Capability{
		"Voice",
		"SMS",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotFound                             ErrorCode = "NotFound"
	ErrorCodePreconditionFailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist         ErrorCode = "PhoneNumberAssociationsExist"
	ErrorCodeGone                                 ErrorCode = "Gone"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
		"Gone",
	}
}

type GeoMatchLevel string

// Enum values for GeoMatchLevel
const (
	GeoMatchLevelCountry  GeoMatchLevel = "Country"
	GeoMatchLevelAreaCode GeoMatchLevel = "AreaCode"
)

// Values returns all known values for GeoMatchLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoMatchLevel) Values() []GeoMatchLevel {
	return []GeoMatchLevel{
		"Country",
		"AreaCode",
	}
}

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeEnUs LanguageCode = "en-US"
)

// Values returns all known values for LanguageCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LanguageCode) Values() []LanguageCode {
	return []LanguageCode{
		"en-US",
	}
}

type NotificationTarget string

// Enum values for NotificationTarget
const (
	NotificationTargetEventBridge NotificationTarget = "EventBridge"
	NotificationTargetSns         NotificationTarget = "SNS"
	NotificationTargetSqs         NotificationTarget = "SQS"
)

// Values returns all known values for NotificationTarget. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationTarget) Values() []NotificationTarget {
	return []NotificationTarget{
		"EventBridge",
		"SNS",
		"SQS",
	}
}

type NumberSelectionBehavior string

// Enum values for NumberSelectionBehavior
const (
	NumberSelectionBehaviorPreferSticky NumberSelectionBehavior = "PreferSticky"
	NumberSelectionBehaviorAvoidSticky  NumberSelectionBehavior = "AvoidSticky"
)

// Values returns all known values for NumberSelectionBehavior. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NumberSelectionBehavior) Values() []NumberSelectionBehavior {
	return []NumberSelectionBehavior{
		"PreferSticky",
		"AvoidSticky",
	}
}

type OrderedPhoneNumberStatus string

// Enum values for OrderedPhoneNumberStatus
const (
	OrderedPhoneNumberStatusProcessing OrderedPhoneNumberStatus = "Processing"
	OrderedPhoneNumberStatusAcquired   OrderedPhoneNumberStatus = "Acquired"
	OrderedPhoneNumberStatusFailed     OrderedPhoneNumberStatus = "Failed"
)

// Values returns all known values for OrderedPhoneNumberStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OrderedPhoneNumberStatus) Values() []OrderedPhoneNumberStatus {
	return []OrderedPhoneNumberStatus{
		"Processing",
		"Acquired",
		"Failed",
	}
}

type OriginationRouteProtocol string

// Enum values for OriginationRouteProtocol
const (
	OriginationRouteProtocolTcp OriginationRouteProtocol = "TCP"
	OriginationRouteProtocolUdp OriginationRouteProtocol = "UDP"
)

// Values returns all known values for OriginationRouteProtocol. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginationRouteProtocol) Values() []OriginationRouteProtocol {
	return []OriginationRouteProtocol{
		"TCP",
		"UDP",
	}
}

type PhoneNumberAssociationName string

// Enum values for PhoneNumberAssociationName
const (
	PhoneNumberAssociationNameVoiceConnectorId      PhoneNumberAssociationName = "VoiceConnectorId"
	PhoneNumberAssociationNameVoiceConnectorGroupId PhoneNumberAssociationName = "VoiceConnectorGroupId"
	PhoneNumberAssociationNameSipRuleId             PhoneNumberAssociationName = "SipRuleId"
)

// Values returns all known values for PhoneNumberAssociationName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberAssociationName) Values() []PhoneNumberAssociationName {
	return []PhoneNumberAssociationName{
		"VoiceConnectorId",
		"VoiceConnectorGroupId",
		"SipRuleId",
	}
}

type PhoneNumberOrderStatus string

// Enum values for PhoneNumberOrderStatus
const (
	PhoneNumberOrderStatusProcessing       PhoneNumberOrderStatus = "Processing"
	PhoneNumberOrderStatusSuccessful       PhoneNumberOrderStatus = "Successful"
	PhoneNumberOrderStatusFailed           PhoneNumberOrderStatus = "Failed"
	PhoneNumberOrderStatusPartial          PhoneNumberOrderStatus = "Partial"
	PhoneNumberOrderStatusPendingDocuments PhoneNumberOrderStatus = "PendingDocuments"
	PhoneNumberOrderStatusSubmitted        PhoneNumberOrderStatus = "Submitted"
	PhoneNumberOrderStatusFoc              PhoneNumberOrderStatus = "FOC"
	PhoneNumberOrderStatusChangeRequested  PhoneNumberOrderStatus = "ChangeRequested"
	PhoneNumberOrderStatusException        PhoneNumberOrderStatus = "Exception"
	PhoneNumberOrderStatusCancelRequested  PhoneNumberOrderStatus = "CancelRequested"
	PhoneNumberOrderStatusCancelled        PhoneNumberOrderStatus = "Cancelled"
)

// Values returns all known values for PhoneNumberOrderStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberOrderStatus) Values() []PhoneNumberOrderStatus {
	return []PhoneNumberOrderStatus{
		"Processing",
		"Successful",
		"Failed",
		"Partial",
		"PendingDocuments",
		"Submitted",
		"FOC",
		"ChangeRequested",
		"Exception",
		"CancelRequested",
		"Cancelled",
	}
}

type PhoneNumberOrderType string

// Enum values for PhoneNumberOrderType
const (
	PhoneNumberOrderTypeNew     PhoneNumberOrderType = "New"
	PhoneNumberOrderTypePorting PhoneNumberOrderType = "Porting"
)

// Values returns all known values for PhoneNumberOrderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberOrderType) Values() []PhoneNumberOrderType {
	return []PhoneNumberOrderType{
		"New",
		"Porting",
	}
}

type PhoneNumberProductType string

// Enum values for PhoneNumberProductType
const (
	PhoneNumberProductTypeVoiceConnector            PhoneNumberProductType = "VoiceConnector"
	PhoneNumberProductTypeSipMediaApplicationDialIn PhoneNumberProductType = "SipMediaApplicationDialIn"
)

// Values returns all known values for PhoneNumberProductType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberProductType) Values() []PhoneNumberProductType {
	return []PhoneNumberProductType{
		"VoiceConnector",
		"SipMediaApplicationDialIn",
	}
}

type PhoneNumberStatus string

// Enum values for PhoneNumberStatus
const (
	PhoneNumberStatusCancelled             PhoneNumberStatus = "Cancelled"
	PhoneNumberStatusPortinCancelRequested PhoneNumberStatus = "PortinCancelRequested"
	PhoneNumberStatusPortinInProgress      PhoneNumberStatus = "PortinInProgress"
	PhoneNumberStatusAcquireInProgress     PhoneNumberStatus = "AcquireInProgress"
	PhoneNumberStatusAcquireFailed         PhoneNumberStatus = "AcquireFailed"
	PhoneNumberStatusUnassigned            PhoneNumberStatus = "Unassigned"
	PhoneNumberStatusAssigned              PhoneNumberStatus = "Assigned"
	PhoneNumberStatusReleaseInProgress     PhoneNumberStatus = "ReleaseInProgress"
	PhoneNumberStatusDeleteInProgress      PhoneNumberStatus = "DeleteInProgress"
	PhoneNumberStatusReleaseFailed         PhoneNumberStatus = "ReleaseFailed"
	PhoneNumberStatusDeleteFailed          PhoneNumberStatus = "DeleteFailed"
)

// Values returns all known values for PhoneNumberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberStatus) Values() []PhoneNumberStatus {
	return []PhoneNumberStatus{
		"Cancelled",
		"PortinCancelRequested",
		"PortinInProgress",
		"AcquireInProgress",
		"AcquireFailed",
		"Unassigned",
		"Assigned",
		"ReleaseInProgress",
		"DeleteInProgress",
		"ReleaseFailed",
		"DeleteFailed",
	}
}

type PhoneNumberType string

// Enum values for PhoneNumberType
const (
	PhoneNumberTypeLocal    PhoneNumberType = "Local"
	PhoneNumberTypeTollFree PhoneNumberType = "TollFree"
)

// Values returns all known values for PhoneNumberType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberType) Values() []PhoneNumberType {
	return []PhoneNumberType{
		"Local",
		"TollFree",
	}
}

type ProxySessionStatus string

// Enum values for ProxySessionStatus
const (
	ProxySessionStatusOpen       ProxySessionStatus = "Open"
	ProxySessionStatusInProgress ProxySessionStatus = "InProgress"
	ProxySessionStatusClosed     ProxySessionStatus = "Closed"
)

// Values returns all known values for ProxySessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProxySessionStatus) Values() []ProxySessionStatus {
	return []ProxySessionStatus{
		"Open",
		"InProgress",
		"Closed",
	}
}

type SipRuleTriggerType string

// Enum values for SipRuleTriggerType
const (
	SipRuleTriggerTypeToPhoneNumber      SipRuleTriggerType = "ToPhoneNumber"
	SipRuleTriggerTypeRequestUriHostname SipRuleTriggerType = "RequestUriHostname"
)

// Values returns all known values for SipRuleTriggerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SipRuleTriggerType) Values() []SipRuleTriggerType {
	return []SipRuleTriggerType{
		"ToPhoneNumber",
		"RequestUriHostname",
	}
}

type VoiceConnectorAwsRegion string

// Enum values for VoiceConnectorAwsRegion
const (
	VoiceConnectorAwsRegionUsEast1      VoiceConnectorAwsRegion = "us-east-1"
	VoiceConnectorAwsRegionUsWest2      VoiceConnectorAwsRegion = "us-west-2"
	VoiceConnectorAwsRegionCaCentral1   VoiceConnectorAwsRegion = "ca-central-1"
	VoiceConnectorAwsRegionEuCentral1   VoiceConnectorAwsRegion = "eu-central-1"
	VoiceConnectorAwsRegionEuWest1      VoiceConnectorAwsRegion = "eu-west-1"
	VoiceConnectorAwsRegionEuWest2      VoiceConnectorAwsRegion = "eu-west-2"
	VoiceConnectorAwsRegionApNortheast2 VoiceConnectorAwsRegion = "ap-northeast-2"
	VoiceConnectorAwsRegionApNortheast1 VoiceConnectorAwsRegion = "ap-northeast-1"
	VoiceConnectorAwsRegionApSoutheast1 VoiceConnectorAwsRegion = "ap-southeast-1"
	VoiceConnectorAwsRegionApSoutheast2 VoiceConnectorAwsRegion = "ap-southeast-2"
)

// Values returns all known values for VoiceConnectorAwsRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VoiceConnectorAwsRegion) Values() []VoiceConnectorAwsRegion {
	return []VoiceConnectorAwsRegion{
		"us-east-1",
		"us-west-2",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"ap-northeast-2",
		"ap-northeast-1",
		"ap-southeast-1",
		"ap-southeast-2",
	}
}
