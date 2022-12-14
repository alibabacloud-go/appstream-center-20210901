// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApproveOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	OtaType            *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApproveOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskRequest) SetAppInstanceGroupId(v string) *ApproveOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetBizRegionId(v string) *ApproveOtaTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetOtaType(v string) *ApproveOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetStartTime(v string) *ApproveOtaTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetTaskId(v string) *ApproveOtaTaskRequest {
	s.TaskId = &v
	return s
}

type ApproveOtaTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponseBody) SetCode(v string) *ApproveOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetMessage(v string) *ApproveOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveOtaTaskResponseBody) SetRequestId(v string) *ApproveOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type ApproveOtaTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ApproveOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskResponse) SetHeaders(v map[string]*string) *ApproveOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ApproveOtaTaskResponse) SetStatusCode(v int32) *ApproveOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOtaTaskResponse) SetBody(v *ApproveOtaTaskResponseBody) *ApproveOtaTaskResponse {
	s.Body = v
	return s
}

type AuthorizeInstanceGroupRequest struct {
	AppInstanceGroupId *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AuthorizeUserIds   []*string `json:"AuthorizeUserIds,omitempty" xml:"AuthorizeUserIds,omitempty" type:"Repeated"`
	ProductType        *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	UnAuthorizeUserIds []*string `json:"UnAuthorizeUserIds,omitempty" xml:"UnAuthorizeUserIds,omitempty" type:"Repeated"`
}

func (s AuthorizeInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupRequest) SetAppInstanceGroupId(v string) *AuthorizeInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.AuthorizeUserIds = v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetProductType(v string) *AuthorizeInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *AuthorizeInstanceGroupRequest) SetUnAuthorizeUserIds(v []*string) *AuthorizeInstanceGroupRequest {
	s.UnAuthorizeUserIds = v
	return s
}

type AuthorizeInstanceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponseBody) SetRequestId(v string) *AuthorizeInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeInstanceGroupResponse) SetHeaders(v map[string]*string) *AuthorizeInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetStatusCode(v int32) *AuthorizeInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeInstanceGroupResponse) SetBody(v *AuthorizeInstanceGroupResponseBody) *AuthorizeInstanceGroupResponse {
	s.Body = v
	return s
}

type CancelOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskRequest) SetAppInstanceGroupId(v string) *CancelOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CancelOtaTaskRequest) SetTaskId(v string) *CancelOtaTaskRequest {
	s.TaskId = &v
	return s
}

type CancelOtaTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskResponseBody) SetCode(v string) *CancelOtaTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOtaTaskResponseBody) SetMessage(v string) *CancelOtaTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOtaTaskResponseBody) SetRequestId(v string) *CancelOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelOtaTaskResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelOtaTaskResponse) SetHeaders(v map[string]*string) *CancelOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelOtaTaskResponse) SetStatusCode(v int32) *CancelOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelOtaTaskResponse) SetBody(v *CancelOtaTaskResponseBody) *CancelOtaTaskResponse {
	s.Body = v
	return s
}

type CreateAppInstanceGroupRequest struct {
	AppCenterImageId     *string                                `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string                                `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AutoPay              *bool                                  `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool                                  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BizRegionId          *string                                `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeResourceMode   *string                                `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType           *string                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodePool             *CreateAppInstanceGroupRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	Period               *int32                                 `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string                                `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType          *string                                `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	PromotionId          *string                                `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	SessionTimeout       *int32                                 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	UserInfo             *CreateAppInstanceGroupRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	Users                []*string                              `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoPay(v bool) *CreateAppInstanceGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetBizRegionId(v string) *CreateAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetChargeType(v string) *CreateAppInstanceGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetNodePool(v *CreateAppInstanceGroupRequestNodePool) *CreateAppInstanceGroupRequest {
	s.NodePool = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriod(v int32) *CreateAppInstanceGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetProductType(v string) *CreateAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetPromotionId(v string) *CreateAppInstanceGroupRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUserInfo(v *CreateAppInstanceGroupRequestUserInfo) *CreateAppInstanceGroupRequest {
	s.UserInfo = v
	return s
}

func (s *CreateAppInstanceGroupRequest) SetUsers(v []*string) *CreateAppInstanceGroupRequest {
	s.Users = v
	return s
}

type CreateAppInstanceGroupRequestNodePool struct {
	MaxScalingAmount            *int32  `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	NodeAmount                  *int32  `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	NodeCapacity                *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType            *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	ScalingDownAfterIdleMinutes *int32  `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32  `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyType                *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateAppInstanceGroupRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestNodePool) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestNodePool) SetMaxScalingAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeAmount(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeAmount = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeCapacity(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetNodeInstanceType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingDownAfterIdleMinutes(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingStep(v int32) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingStep = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetScalingUsageThreshold(v string) *CreateAppInstanceGroupRequestNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *CreateAppInstanceGroupRequestNodePool) SetStrategyType(v string) *CreateAppInstanceGroupRequestNodePool {
	s.StrategyType = &v
	return s
}

type CreateAppInstanceGroupRequestUserInfo struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppInstanceGroupRequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupRequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupRequestUserInfo) SetType(v string) *CreateAppInstanceGroupRequestUserInfo {
	s.Type = &v
	return s
}

type CreateAppInstanceGroupShrinkRequest struct {
	AppCenterImageId     *string   `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupName *string   `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AutoPay              *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	BizRegionId          *string   `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeResourceMode   *string   `json:"ChargeResourceMode,omitempty" xml:"ChargeResourceMode,omitempty"`
	ChargeType           *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodePoolShrink       *string   `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	Period               *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType          *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	PromotionId          *string   `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	SessionTimeout       *int32    `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	UserInfoShrink       *string   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	Users                []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateAppInstanceGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppCenterImageId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAppInstanceGroupName(v string) *CreateAppInstanceGroupShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoPay(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetAutoRenew(v bool) *CreateAppInstanceGroupShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetBizRegionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeResourceMode(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeResourceMode = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetChargeType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetNodePoolShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriod(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPeriodUnit(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetProductType(v string) *CreateAppInstanceGroupShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetPromotionId(v string) *CreateAppInstanceGroupShrinkRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetSessionTimeout(v int32) *CreateAppInstanceGroupShrinkRequest {
	s.SessionTimeout = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUserInfoShrink(v string) *CreateAppInstanceGroupShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *CreateAppInstanceGroupShrinkRequest) SetUsers(v []*string) *CreateAppInstanceGroupShrinkRequest {
	s.Users = v
	return s
}

type CreateAppInstanceGroupResponseBody struct {
	AppInstanceGroupModel *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel `json:"AppInstanceGroupModel,omitempty" xml:"AppInstanceGroupModel,omitempty" type:"Struct"`
	RequestId             *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBody) SetAppInstanceGroupModel(v *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) *CreateAppInstanceGroupResponseBody {
	s.AppInstanceGroupModel = v
	return s
}

func (s *CreateAppInstanceGroupResponseBody) SetRequestId(v string) *CreateAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppInstanceGroupResponseBodyAppInstanceGroupModel struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	NodePoolId         *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetAppInstanceGroupId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.AppInstanceGroupId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetNodePoolId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.NodePoolId = &v
	return s
}

func (s *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel) SetOrderId(v string) *CreateAppInstanceGroupResponseBodyAppInstanceGroupModel {
	s.OrderId = &v
	return s
}

type CreateAppInstanceGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppInstanceGroupResponse) SetHeaders(v map[string]*string) *CreateAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetStatusCode(v int32) *CreateAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppInstanceGroupResponse) SetBody(v *CreateAppInstanceGroupResponseBody) *CreateAppInstanceGroupResponse {
	s.Body = v
	return s
}

type GetOtaTaskByTaskIdRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetOtaTaskByTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdRequest) SetTaskId(v string) *GetOtaTaskByTaskIdRequest {
	s.TaskId = &v
	return s
}

type GetOtaTaskByTaskIdResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OtaVersion    *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	ReleaseNote   *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStartTime *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s GetOtaTaskByTaskIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponseBody) SetCode(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetMessage(v string) *GetOtaTaskByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetOtaVersion(v string) *GetOtaTaskByTaskIdResponseBody {
	s.OtaVersion = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetReleaseNote(v string) *GetOtaTaskByTaskIdResponseBody {
	s.ReleaseNote = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetRequestId(v string) *GetOtaTaskByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponseBody) SetTaskStartTime(v string) *GetOtaTaskByTaskIdResponseBody {
	s.TaskStartTime = &v
	return s
}

type GetOtaTaskByTaskIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOtaTaskByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOtaTaskByTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOtaTaskByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *GetOtaTaskByTaskIdResponse) SetHeaders(v map[string]*string) *GetOtaTaskByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetStatusCode(v int32) *GetOtaTaskByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOtaTaskByTaskIdResponse) SetBody(v *GetOtaTaskByTaskIdResponseBody) *GetOtaTaskByTaskIdResponse {
	s.Body = v
	return s
}

type GetResourcePriceRequest struct {
	Amount           *int64  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	BizRegionId      *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ChargeType       *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	Period           *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ProductType      *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourcePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePriceRequest) SetAmount(v int64) *GetResourcePriceRequest {
	s.Amount = &v
	return s
}

func (s *GetResourcePriceRequest) SetBizRegionId(v string) *GetResourcePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetResourcePriceRequest) SetChargeType(v string) *GetResourcePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *GetResourcePriceRequest) SetNodeInstanceType(v string) *GetResourcePriceRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriod(v int64) *GetResourcePriceRequest {
	s.Period = &v
	return s
}

func (s *GetResourcePriceRequest) SetPeriodUnit(v string) *GetResourcePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetResourcePriceRequest) SetProductType(v string) *GetResourcePriceRequest {
	s.ProductType = &v
	return s
}

type GetResourcePriceResponseBody struct {
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PriceModel *GetResourcePriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourcePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBody) SetCode(v string) *GetResourcePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetMessage(v string) *GetResourcePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceModel(v *GetResourcePriceResponseBodyPriceModel) *GetResourcePriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *GetResourcePriceResponseBody) SetRequestId(v string) *GetResourcePriceResponseBody {
	s.RequestId = &v
	return s
}

type GetResourcePriceResponseBodyPriceModel struct {
	Price *GetResourcePriceResponseBodyPriceModelPrice   `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	Rules []*GetResourcePriceResponseBodyPriceModelRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceModel) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModel) SetPrice(v *GetResourcePriceResponseBodyPriceModelPrice) *GetResourcePriceResponseBodyPriceModel {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModel) SetRules(v []*GetResourcePriceResponseBodyPriceModelRules) *GetResourcePriceResponseBodyPriceModel {
	s.Rules = v
	return s
}

type GetResourcePriceResponseBodyPriceModelPrice struct {
	Currency      *string                                                  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice *string                                                  `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice *string                                                  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*GetResourcePriceResponseBodyPriceModelPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	TradePrice    *string                                                  `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPrice) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceModelPricePromotions) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.TradePrice = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelPricePromotions struct {
	OptionCode    *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	PromotionId   *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	Selected      *bool   `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.Selected = &v
	return s
}

type GetResourcePriceResponseBodyPriceModelRules struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RuleId      *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceModelRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceModelRules {
	s.RuleId = &v
	return s
}

type GetResourcePriceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourcePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourcePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponse) SetHeaders(v map[string]*string) *GetResourcePriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePriceResponse) SetStatusCode(v int32) *GetResourcePriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePriceResponse) SetBody(v *GetResourcePriceResponseBody) *GetResourcePriceResponse {
	s.Body = v
	return s
}

type ListAppInstanceGroupRequest struct {
	AppCenterImageId     *string   `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId   *string   `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string   `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	PageNumber           *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType          *string   `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupRequest) SetAppCenterImageId(v string) *ListAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageNumber(v int32) *ListAppInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageSize(v int32) *ListAppInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetProductType(v string) *ListAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetRegionId(v string) *ListAppInstanceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetStatus(v []*string) *ListAppInstanceGroupRequest {
	s.Status = v
	return s
}

type ListAppInstanceGroupResponseBody struct {
	AppInstanceGroupModels []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels `json:"AppInstanceGroupModels,omitempty" xml:"AppInstanceGroupModels,omitempty" type:"Repeated"`
	PageNumber             *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppInstanceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBody) SetAppInstanceGroupModels(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModels) *ListAppInstanceGroupResponseBody {
	s.AppInstanceGroupModels = v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageNumber(v int32) *ListAppInstanceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetPageSize(v int32) *ListAppInstanceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetRequestId(v string) *ListAppInstanceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBody) SetTotalCount(v int32) *ListAppInstanceGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModels struct {
	Amount               *int32                                                            `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AppCenterImageId     *string                                                           `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId   *string                                                           `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string                                                           `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	AppInstanceType      *string                                                           `json:"AppInstanceType,omitempty" xml:"AppInstanceType,omitempty"`
	Apps                 []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps     `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	ChargeType           *string                                                           `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExpiredTime          *string                                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	GmtCreate            *string                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	NodePool             []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Repeated"`
	OsType               *string                                                           `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OtaInfo              *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo    `json:"OtaInfo,omitempty" xml:"OtaInfo,omitempty" type:"Struct"`
	ProductType          *string                                                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId             *string                                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionTimeout       *string                                                           `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	SpecId               *string                                                           `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status               *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModels) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppCenterImageId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceGroupName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetAppInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.AppInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetApps(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Apps = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetChargeType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ChargeType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetExpiredTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ExpiredTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetGmtCreate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.GmtCreate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetNodePool(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.NodePool = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOsType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OsType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetOtaInfo(v *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.OtaInfo = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetProductType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetRegionId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSessionTimeout(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SessionTimeout = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetSpecId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.SpecId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModels) SetStatus(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModels {
	s.Status = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps) SetAppName(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsApps {
	s.AppName = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool struct {
	Amount                      *int32                                                                               `json:"Amount,omitempty" xml:"Amount,omitempty"`
	MaxScalingAmount            *int32                                                                               `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	NodeAmount                  *int32                                                                               `json:"NodeAmount,omitempty" xml:"NodeAmount,omitempty"`
	NodeCapacity                *int32                                                                               `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodeInstanceType            *string                                                                              `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodePoolId                  *string                                                                              `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
	NodeUsed                    *int32                                                                               `json:"NodeUsed,omitempty" xml:"NodeUsed,omitempty"`
	RecurrenceSchedules         []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules `json:"RecurrenceSchedules,omitempty" xml:"RecurrenceSchedules,omitempty" type:"Repeated"`
	ScalingDownAfterIdleMinutes *int32                                                                               `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingNodeAmount           *int32                                                                               `json:"ScalingNodeAmount,omitempty" xml:"ScalingNodeAmount,omitempty"`
	ScalingNodeUsed             *int32                                                                               `json:"ScalingNodeUsed,omitempty" xml:"ScalingNodeUsed,omitempty"`
	ScalingStep                 *int32                                                                               `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string                                                                              `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyDisableDate         *string                                                                              `json:"StrategyDisableDate,omitempty" xml:"StrategyDisableDate,omitempty"`
	StrategyEnableDate          *string                                                                              `json:"StrategyEnableDate,omitempty" xml:"StrategyEnableDate,omitempty"`
	StrategyType                *string                                                                              `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	WarmUp                      *bool                                                                                `json:"WarmUp,omitempty" xml:"WarmUp,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetMaxScalingAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.MaxScalingAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeCapacity(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeInstanceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodePoolId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodePoolId = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.NodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetRecurrenceSchedules(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.RecurrenceSchedules = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingDownAfterIdleMinutes(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeAmount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingNodeUsed(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingNodeUsed = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingStep(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingStep = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetScalingUsageThreshold(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyDisableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyDisableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyEnableDate(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyEnableDate = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetStrategyType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.StrategyType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool) SetWarmUp(v bool) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePool {
	s.WarmUp = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules struct {
	RecurrenceType   *string                                                                                          `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValues []*int32                                                                                         `json:"RecurrenceValues,omitempty" xml:"RecurrenceValues,omitempty" type:"Repeated"`
	TimerPeriods     []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods `json:"TimerPeriods,omitempty" xml:"TimerPeriods,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceType(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceType = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetRecurrenceValues(v []*int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.RecurrenceValues = v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules) SetTimerPeriods(v []*ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedules {
	s.TimerPeriods = v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods struct {
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetAmount(v int32) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.Amount = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetEndTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.EndTime = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods) SetStartTime(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsNodePoolRecurrenceSchedulesTimerPeriods {
	s.StartTime = &v
	return s
}

type ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo struct {
	NewOtaVersion *string `json:"NewOtaVersion,omitempty" xml:"NewOtaVersion,omitempty"`
	OtaVersion    *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	TaskId        *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetNewOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.NewOtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetOtaVersion(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.OtaVersion = &v
	return s
}

func (s *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo) SetTaskId(v string) *ListAppInstanceGroupResponseBodyAppInstanceGroupModelsOtaInfo {
	s.TaskId = &v
	return s
}

type ListAppInstanceGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppInstanceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppInstanceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppInstanceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupResponse) SetHeaders(v map[string]*string) *ListAppInstanceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListAppInstanceGroupResponse) SetStatusCode(v int32) *ListAppInstanceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInstanceGroupResponse) SetBody(v *ListAppInstanceGroupResponseBody) *ListAppInstanceGroupResponse {
	s.Body = v
	return s
}

type ListNodeInstanceTypeRequest struct {
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	Language    *string `json:"Language,omitempty" xml:"Language,omitempty"`
	OsType      *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListNodeInstanceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeRequest) SetBizRegionId(v string) *ListNodeInstanceTypeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetLanguage(v string) *ListNodeInstanceTypeRequest {
	s.Language = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetOsType(v string) *ListNodeInstanceTypeRequest {
	s.OsType = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageNumber(v int32) *ListNodeInstanceTypeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetPageSize(v int32) *ListNodeInstanceTypeRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeRequest) SetProductType(v string) *ListNodeInstanceTypeRequest {
	s.ProductType = &v
	return s
}

type ListNodeInstanceTypeResponseBody struct {
	NodeInstanceTypeModels []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels `json:"NodeInstanceTypeModels,omitempty" xml:"NodeInstanceTypeModels,omitempty" type:"Repeated"`
	PageNumber             *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeInstanceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBody) SetNodeInstanceTypeModels(v []*ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) *ListNodeInstanceTypeResponseBody {
	s.NodeInstanceTypeModels = v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageNumber(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetPageSize(v int32) *ListNodeInstanceTypeResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetRequestId(v string) *ListNodeInstanceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBody) SetTotalCount(v int32) *ListNodeInstanceTypeResponseBody {
	s.TotalCount = &v
	return s
}

type ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels struct {
	Cpu                    *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu                    *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	GpuMemory              *int64  `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	MaxCapacity            *int32  `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	Memory                 *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NodeInstanceType       *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	NodeInstanceTypeFamily *string `json:"NodeInstanceTypeFamily,omitempty" xml:"NodeInstanceTypeFamily,omitempty"`
	NodeTypeName           *string `json:"NodeTypeName,omitempty" xml:"NodeTypeName,omitempty"`
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetCpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Cpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpu(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Gpu = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetGpuMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.GpuMemory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMaxCapacity(v int32) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.MaxCapacity = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetMemory(v int64) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.Memory = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceType(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceType = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeInstanceTypeFamily(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeInstanceTypeFamily = &v
	return s
}

func (s *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels) SetNodeTypeName(v string) *ListNodeInstanceTypeResponseBodyNodeInstanceTypeModels {
	s.NodeTypeName = &v
	return s
}

type ListNodeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListNodeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodeInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListNodeInstanceTypeResponse) SetHeaders(v map[string]*string) *ListNodeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetStatusCode(v int32) *ListNodeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeInstanceTypeResponse) SetBody(v *ListNodeInstanceTypeResponseBody) *ListNodeInstanceTypeResponse {
	s.Body = v
	return s
}

type ListOtaTaskRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	OtaType            *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOtaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOtaTaskRequest) SetAppInstanceGroupId(v string) *ListOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListOtaTaskRequest) SetOtaType(v string) *ListOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageNumber(v int32) *ListOtaTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskRequest) SetPageSize(v int32) *ListOtaTaskRequest {
	s.PageSize = &v
	return s
}

type ListOtaTaskResponseBody struct {
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList   []*ListOtaTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOtaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBody) SetPageNumber(v int32) *ListOtaTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetPageSize(v int32) *ListOtaTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetRequestId(v string) *ListOtaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOtaTaskResponseBody) SetTaskList(v []*ListOtaTaskResponseBodyTaskList) *ListOtaTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *ListOtaTaskResponseBody) SetTotalCount(v int32) *ListOtaTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListOtaTaskResponseBodyTaskList struct {
	OtaVersion        *string `json:"OtaVersion,omitempty" xml:"OtaVersion,omitempty"`
	TaskDisplayStatus *string `json:"TaskDisplayStatus,omitempty" xml:"TaskDisplayStatus,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStartTime     *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
}

func (s ListOtaTaskResponseBodyTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponseBodyTaskList) SetOtaVersion(v string) *ListOtaTaskResponseBodyTaskList {
	s.OtaVersion = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskDisplayStatus(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskDisplayStatus = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskId(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *ListOtaTaskResponseBodyTaskList) SetTaskStartTime(v string) *ListOtaTaskResponseBodyTaskList {
	s.TaskStartTime = &v
	return s
}

type ListOtaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOtaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOtaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOtaTaskResponse) GoString() string {
	return s.String()
}

func (s *ListOtaTaskResponse) SetHeaders(v map[string]*string) *ListOtaTaskResponse {
	s.Headers = v
	return s
}

func (s *ListOtaTaskResponse) SetStatusCode(v int32) *ListOtaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOtaTaskResponse) SetBody(v *ListOtaTaskResponseBody) *ListOtaTaskResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	RegionModels []*ListRegionsResponseBodyRegionModels `json:"RegionModels,omitempty" xml:"RegionModels,omitempty" type:"Repeated"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegionModels(v []*ListRegionsResponseBodyRegionModels) *ListRegionsResponseBody {
	s.RegionModels = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegionModels struct {
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionModels) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionModels) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionModels) SetRegionId(v string) *ListRegionsResponseBodyRegionModels {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ModifyAppInstanceGroupAttributeRequest struct {
	AppInstanceGroupId   *string                                         `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string                                         `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	NodePool             *ModifyAppInstanceGroupAttributeRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	ProductType          *string                                         `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SessionTimeout       *int32                                          `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNodePool(v *ModifyAppInstanceGroupAttributeRequestNodePool) *ModifyAppInstanceGroupAttributeRequest {
	s.NodePool = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeRequest {
	s.SessionTimeout = &v
	return s
}

type ModifyAppInstanceGroupAttributeRequestNodePool struct {
	NodeCapacity *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolId   *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodeCapacity(v int32) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodePoolId(v string) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodePoolId = &v
	return s
}

type ModifyAppInstanceGroupAttributeShrinkRequest struct {
	AppInstanceGroupId   *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	NodePoolShrink       *string `json:"NodePool,omitempty" xml:"NodePool,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetNodePoolShrink(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.NodePoolShrink = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeShrinkRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeShrinkRequest {
	s.SessionTimeout = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetCode(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetMessage(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponseBody) SetRequestId(v string) *ModifyAppInstanceGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppInstanceGroupAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppInstanceGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppInstanceGroupAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetHeaders(v map[string]*string) *ModifyAppInstanceGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetStatusCode(v int32) *ModifyAppInstanceGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeResponse) SetBody(v *ModifyAppInstanceGroupAttributeResponseBody) *ModifyAppInstanceGroupAttributeResponse {
	s.Body = v
	return s
}

type ModifyNodePoolAttributeRequest struct {
	BizRegionId      *string                                         `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeCapacity     *int32                                          `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategy *ModifyNodePoolAttributeRequestNodePoolStrategy `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty" type:"Struct"`
	PoolId           *string                                         `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	ProductType      *string                                         `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetNodePoolStrategy(v *ModifyNodePoolAttributeRequestNodePoolStrategy) *ModifyNodePoolAttributeRequest {
	s.NodePoolStrategy = v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetPoolId(v string) *ModifyNodePoolAttributeRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeRequest) SetProductType(v string) *ModifyNodePoolAttributeRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeRequestNodePoolStrategy struct {
	MaxScalingAmount            *int32  `json:"MaxScalingAmount,omitempty" xml:"MaxScalingAmount,omitempty"`
	ScalingDownAfterIdleMinutes *int32  `json:"ScalingDownAfterIdleMinutes,omitempty" xml:"ScalingDownAfterIdleMinutes,omitempty"`
	ScalingStep                 *int32  `json:"ScalingStep,omitempty" xml:"ScalingStep,omitempty"`
	ScalingUsageThreshold       *string `json:"ScalingUsageThreshold,omitempty" xml:"ScalingUsageThreshold,omitempty"`
	StrategyType                *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeRequestNodePoolStrategy) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetMaxScalingAmount(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.MaxScalingAmount = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingDownAfterIdleMinutes(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingDownAfterIdleMinutes = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingStep(v int32) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingStep = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetScalingUsageThreshold(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.ScalingUsageThreshold = &v
	return s
}

func (s *ModifyNodePoolAttributeRequestNodePoolStrategy) SetStrategyType(v string) *ModifyNodePoolAttributeRequestNodePoolStrategy {
	s.StrategyType = &v
	return s
}

type ModifyNodePoolAttributeShrinkRequest struct {
	BizRegionId            *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	NodeCapacity           *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategyShrink *string `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty"`
	PoolId                 *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	ProductType            *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeShrinkRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodePoolStrategyShrink(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.NodePoolStrategyShrink = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetPoolId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetProductType(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

type ModifyNodePoolAttributeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNodePoolAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponseBody) SetCode(v string) *ModifyNodePoolAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetMessage(v string) *ModifyNodePoolAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyNodePoolAttributeResponseBody) SetRequestId(v string) *ModifyNodePoolAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNodePoolAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyNodePoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyNodePoolAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNodePoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeResponse) SetHeaders(v map[string]*string) *ModifyNodePoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetStatusCode(v int32) *ModifyNodePoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNodePoolAttributeResponse) SetBody(v *ModifyNodePoolAttributeResponseBody) *ModifyNodePoolAttributeResponse {
	s.Body = v
	return s
}

type PageListAppInstanceGroupUserRequest struct {
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s PageListAppInstanceGroupUserRequest) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserRequest) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserRequest) SetAppInstanceGroupId(v string) *PageListAppInstanceGroupUserRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageNumber(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageNumber = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetPageSize(v int32) *PageListAppInstanceGroupUserRequest {
	s.PageSize = &v
	return s
}

func (s *PageListAppInstanceGroupUserRequest) SetProductType(v string) *PageListAppInstanceGroupUserRequest {
	s.ProductType = &v
	return s
}

type PageListAppInstanceGroupUserResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s PageListAppInstanceGroupUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponseBody) SetRequestId(v string) *PageListAppInstanceGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponseBody) SetUsers(v []*string) *PageListAppInstanceGroupUserResponseBody {
	s.Users = v
	return s
}

type PageListAppInstanceGroupUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageListAppInstanceGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageListAppInstanceGroupUserResponse) String() string {
	return tea.Prettify(s)
}

func (s PageListAppInstanceGroupUserResponse) GoString() string {
	return s.String()
}

func (s *PageListAppInstanceGroupUserResponse) SetHeaders(v map[string]*string) *PageListAppInstanceGroupUserResponse {
	s.Headers = v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetStatusCode(v int32) *PageListAppInstanceGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *PageListAppInstanceGroupUserResponse) SetBody(v *PageListAppInstanceGroupUserResponseBody) *PageListAppInstanceGroupUserResponse {
	s.Body = v
	return s
}

type UpdateAppInstanceGroupImageRequest struct {
	AppCenterImageId   *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s UpdateAppInstanceGroupImageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppCenterImageId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetAppInstanceGroupId(v string) *UpdateAppInstanceGroupImageRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetBizRegionId(v string) *UpdateAppInstanceGroupImageRequest {
	s.BizRegionId = &v
	return s
}

func (s *UpdateAppInstanceGroupImageRequest) SetProductType(v string) *UpdateAppInstanceGroupImageRequest {
	s.ProductType = &v
	return s
}

type UpdateAppInstanceGroupImageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAppInstanceGroupImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetCode(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetMessage(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponseBody) SetRequestId(v string) *UpdateAppInstanceGroupImageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAppInstanceGroupImageResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAppInstanceGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppInstanceGroupImageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppInstanceGroupImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppInstanceGroupImageResponse) SetHeaders(v map[string]*string) *UpdateAppInstanceGroupImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetStatusCode(v int32) *UpdateAppInstanceGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppInstanceGroupImageResponse) SetBody(v *UpdateAppInstanceGroupImageResponseBody) *UpdateAppInstanceGroupImageResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveOtaTaskWithOptions(request *ApproveOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ApproveOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveOtaTask(request *ApproveOtaTaskRequest) (_result *ApproveOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOtaTaskResponse{}
	_body, _err := client.ApproveOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeInstanceGroupWithOptions(request *AuthorizeInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *AuthorizeInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizeUserIds)) {
		body["AuthorizeUserIds"] = request.AuthorizeUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.UnAuthorizeUserIds)) {
		body["UnAuthorizeUserIds"] = request.UnAuthorizeUserIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeInstanceGroup(request *AuthorizeInstanceGroupRequest) (_result *AuthorizeInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeInstanceGroupResponse{}
	_body, _err := client.AuthorizeInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOtaTaskWithOptions(request *CancelOtaTaskRequest, runtime *util.RuntimeOptions) (_result *CancelOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOtaTask(request *CancelOtaTaskRequest) (_result *CancelOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOtaTaskResponse{}
	_body, _err := client.CancelOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppInstanceGroupWithOptions(tmpReq *CreateAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppInstanceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.NodePool))) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.NodePool), tea.String("NodePool"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.UserInfo))) {
		request.UserInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.UserInfo), tea.String("UserInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		body["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		body["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeResourceMode)) {
		body["ChargeResourceMode"] = request.ChargeResourceMode
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		body["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.PromotionId)) {
		body["PromotionId"] = request.PromotionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		body["SessionTimeout"] = request.SessionTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfoShrink)) {
		body["UserInfo"] = request.UserInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppInstanceGroup(request *CreateAppInstanceGroupRequest) (_result *CreateAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppInstanceGroupResponse{}
	_body, _err := client.CreateAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOtaTaskByTaskIdWithOptions(request *GetOtaTaskByTaskIdRequest, runtime *util.RuntimeOptions) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOtaTaskByTaskId"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOtaTaskByTaskId(request *GetOtaTaskByTaskIdRequest) (_result *GetOtaTaskByTaskIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOtaTaskByTaskIdResponse{}
	_body, _err := client.GetOtaTaskByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourcePriceWithOptions(request *GetResourcePriceRequest, runtime *util.RuntimeOptions) (_result *GetResourcePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeInstanceType)) {
		query["NodeInstanceType"] = request.NodeInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourcePrice"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourcePrice(request *GetResourcePriceRequest) (_result *GetResourcePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePriceResponse{}
	_body, _err := client.GetResourcePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppInstanceGroupWithOptions(request *ListAppInstanceGroupRequest, runtime *util.RuntimeOptions) (_result *ListAppInstanceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppInstanceGroup"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppInstanceGroup(request *ListAppInstanceGroupRequest) (_result *ListAppInstanceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppInstanceGroupResponse{}
	_body, _err := client.ListAppInstanceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodeInstanceTypeWithOptions(request *ListNodeInstanceTypeRequest, runtime *util.RuntimeOptions) (_result *ListNodeInstanceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeInstanceType"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodeInstanceType(request *ListNodeInstanceTypeRequest) (_result *ListNodeInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeInstanceTypeResponse{}
	_body, _err := client.ListNodeInstanceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOtaTaskWithOptions(request *ListOtaTaskRequest, runtime *util.RuntimeOptions) (_result *ListOtaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OtaType)) {
		body["OtaType"] = request.OtaType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOtaTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOtaTask(request *ListOtaTaskRequest) (_result *ListOtaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOtaTaskResponse{}
	_body, _err := client.ListOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppInstanceGroupAttributeWithOptions(tmpReq *ModifyAppInstanceGroupAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppInstanceGroupAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.NodePool))) {
		request.NodePoolShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.NodePool), tea.String("NodePool"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupName)) {
		query["AppInstanceGroupName"] = request.AppInstanceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolShrink)) {
		query["NodePool"] = request.NodePoolShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppInstanceGroupAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppInstanceGroupAttribute(request *ModifyAppInstanceGroupAttributeRequest) (_result *ModifyAppInstanceGroupAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppInstanceGroupAttributeResponse{}
	_body, _err := client.ModifyAppInstanceGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNodePoolAttributeWithOptions(tmpReq *ModifyNodePoolAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyNodePoolAttributeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyNodePoolAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.NodePoolStrategy))) {
		request.NodePoolStrategyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.NodePoolStrategy), tea.String("NodePoolStrategy"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeCapacity)) {
		body["NodeCapacity"] = request.NodeCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.NodePoolStrategyShrink)) {
		body["NodePoolStrategy"] = request.NodePoolStrategyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PoolId)) {
		body["PoolId"] = request.PoolId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNodePoolAttribute"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNodePoolAttribute(request *ModifyNodePoolAttributeRequest) (_result *ModifyNodePoolAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNodePoolAttributeResponse{}
	_body, _err := client.ModifyNodePoolAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageListAppInstanceGroupUserWithOptions(request *PageListAppInstanceGroupUserRequest, runtime *util.RuntimeOptions) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageListAppInstanceGroupUser"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageListAppInstanceGroupUser(request *PageListAppInstanceGroupUserRequest) (_result *PageListAppInstanceGroupUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageListAppInstanceGroupUserResponse{}
	_body, _err := client.PageListAppInstanceGroupUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppInstanceGroupImageWithOptions(request *UpdateAppInstanceGroupImageRequest, runtime *util.RuntimeOptions) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCenterImageId)) {
		query["AppCenterImageId"] = request.AppCenterImageId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		query["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAppInstanceGroupImage"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAppInstanceGroupImage(request *UpdateAppInstanceGroupImageRequest) (_result *UpdateAppInstanceGroupImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppInstanceGroupImageResponse{}
	_body, _err := client.UpdateAppInstanceGroupImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
