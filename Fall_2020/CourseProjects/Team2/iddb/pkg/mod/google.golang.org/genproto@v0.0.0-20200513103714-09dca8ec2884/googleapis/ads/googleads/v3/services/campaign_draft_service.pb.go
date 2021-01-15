// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/campaign_draft_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [CampaignDraftService.GetCampaignDraft][google.ads.googleads.v3.services.CampaignDraftService.GetCampaignDraft].
type GetCampaignDraftRequest struct {
	// Required. The resource name of the campaign draft to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignDraftRequest) Reset()         { *m = GetCampaignDraftRequest{} }
func (m *GetCampaignDraftRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignDraftRequest) ProtoMessage()    {}
func (*GetCampaignDraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{0}
}

func (m *GetCampaignDraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignDraftRequest.Unmarshal(m, b)
}
func (m *GetCampaignDraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignDraftRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignDraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignDraftRequest.Merge(m, src)
}
func (m *GetCampaignDraftRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignDraftRequest.Size(m)
}
func (m *GetCampaignDraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignDraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignDraftRequest proto.InternalMessageInfo

func (m *GetCampaignDraftRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignDraftService.MutateCampaignDrafts][google.ads.googleads.v3.services.CampaignDraftService.MutateCampaignDrafts].
type MutateCampaignDraftsRequest struct {
	// Required. The ID of the customer whose campaign drafts are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual campaign drafts.
	Operations []*CampaignDraftOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignDraftsRequest) Reset()         { *m = MutateCampaignDraftsRequest{} }
func (m *MutateCampaignDraftsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftsRequest) ProtoMessage()    {}
func (*MutateCampaignDraftsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{1}
}

func (m *MutateCampaignDraftsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Unmarshal(m, b)
}
func (m *MutateCampaignDraftsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftsRequest.Merge(m, src)
}
func (m *MutateCampaignDraftsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftsRequest.Size(m)
}
func (m *MutateCampaignDraftsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftsRequest proto.InternalMessageInfo

func (m *MutateCampaignDraftsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateCampaignDraftsRequest) GetOperations() []*CampaignDraftOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateCampaignDraftsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateCampaignDraftsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// Request message for [CampaignDraftService.PromoteCampaignDraft][google.ads.googleads.v3.services.CampaignDraftService.PromoteCampaignDraft].
type PromoteCampaignDraftRequest struct {
	// Required. The resource name of the campaign draft to promote.
	CampaignDraft        string   `protobuf:"bytes,1,opt,name=campaign_draft,json=campaignDraft,proto3" json:"campaign_draft,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PromoteCampaignDraftRequest) Reset()         { *m = PromoteCampaignDraftRequest{} }
func (m *PromoteCampaignDraftRequest) String() string { return proto.CompactTextString(m) }
func (*PromoteCampaignDraftRequest) ProtoMessage()    {}
func (*PromoteCampaignDraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{2}
}

func (m *PromoteCampaignDraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Unmarshal(m, b)
}
func (m *PromoteCampaignDraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Marshal(b, m, deterministic)
}
func (m *PromoteCampaignDraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromoteCampaignDraftRequest.Merge(m, src)
}
func (m *PromoteCampaignDraftRequest) XXX_Size() int {
	return xxx_messageInfo_PromoteCampaignDraftRequest.Size(m)
}
func (m *PromoteCampaignDraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PromoteCampaignDraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PromoteCampaignDraftRequest proto.InternalMessageInfo

func (m *PromoteCampaignDraftRequest) GetCampaignDraft() string {
	if m != nil {
		return m.CampaignDraft
	}
	return ""
}

// A single operation (create, update, remove) on a campaign draft.
type CampaignDraftOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*CampaignDraftOperation_Create
	//	*CampaignDraftOperation_Update
	//	*CampaignDraftOperation_Remove
	Operation            isCampaignDraftOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CampaignDraftOperation) Reset()         { *m = CampaignDraftOperation{} }
func (m *CampaignDraftOperation) String() string { return proto.CompactTextString(m) }
func (*CampaignDraftOperation) ProtoMessage()    {}
func (*CampaignDraftOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{3}
}

func (m *CampaignDraftOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignDraftOperation.Unmarshal(m, b)
}
func (m *CampaignDraftOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignDraftOperation.Marshal(b, m, deterministic)
}
func (m *CampaignDraftOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignDraftOperation.Merge(m, src)
}
func (m *CampaignDraftOperation) XXX_Size() int {
	return xxx_messageInfo_CampaignDraftOperation.Size(m)
}
func (m *CampaignDraftOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignDraftOperation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignDraftOperation proto.InternalMessageInfo

func (m *CampaignDraftOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isCampaignDraftOperation_Operation interface {
	isCampaignDraftOperation_Operation()
}

type CampaignDraftOperation_Create struct {
	Create *resources.CampaignDraft `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type CampaignDraftOperation_Update struct {
	Update *resources.CampaignDraft `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type CampaignDraftOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*CampaignDraftOperation_Create) isCampaignDraftOperation_Operation() {}

func (*CampaignDraftOperation_Update) isCampaignDraftOperation_Operation() {}

func (*CampaignDraftOperation_Remove) isCampaignDraftOperation_Operation() {}

func (m *CampaignDraftOperation) GetOperation() isCampaignDraftOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *CampaignDraftOperation) GetCreate() *resources.CampaignDraft {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *CampaignDraftOperation) GetUpdate() *resources.CampaignDraft {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *CampaignDraftOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*CampaignDraftOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignDraftOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignDraftOperation_Create)(nil),
		(*CampaignDraftOperation_Update)(nil),
		(*CampaignDraftOperation_Remove)(nil),
	}
}

// Response message for campaign draft mutate.
type MutateCampaignDraftsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateCampaignDraftResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *MutateCampaignDraftsResponse) Reset()         { *m = MutateCampaignDraftsResponse{} }
func (m *MutateCampaignDraftsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftsResponse) ProtoMessage()    {}
func (*MutateCampaignDraftsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{4}
}

func (m *MutateCampaignDraftsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Unmarshal(m, b)
}
func (m *MutateCampaignDraftsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftsResponse.Merge(m, src)
}
func (m *MutateCampaignDraftsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftsResponse.Size(m)
}
func (m *MutateCampaignDraftsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftsResponse proto.InternalMessageInfo

func (m *MutateCampaignDraftsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateCampaignDraftsResponse) GetResults() []*MutateCampaignDraftResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the campaign draft mutate.
type MutateCampaignDraftResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateCampaignDraftResult) Reset()         { *m = MutateCampaignDraftResult{} }
func (m *MutateCampaignDraftResult) String() string { return proto.CompactTextString(m) }
func (*MutateCampaignDraftResult) ProtoMessage()    {}
func (*MutateCampaignDraftResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{5}
}

func (m *MutateCampaignDraftResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateCampaignDraftResult.Unmarshal(m, b)
}
func (m *MutateCampaignDraftResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateCampaignDraftResult.Marshal(b, m, deterministic)
}
func (m *MutateCampaignDraftResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateCampaignDraftResult.Merge(m, src)
}
func (m *MutateCampaignDraftResult) XXX_Size() int {
	return xxx_messageInfo_MutateCampaignDraftResult.Size(m)
}
func (m *MutateCampaignDraftResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateCampaignDraftResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateCampaignDraftResult proto.InternalMessageInfo

func (m *MutateCampaignDraftResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v3.services.CampaignDraftService.ListCampaignDraftAsyncErrors].
type ListCampaignDraftAsyncErrorsRequest struct {
	// Required. The name of the campaign draft from which to retrieve the async errors.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Token of the page to retrieve. If not specified, the first
	// page of results will be returned. Use the value obtained from
	// `next_page_token` in the previous response in order to request
	// the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Number of elements to retrieve in a single page.
	// When a page request is too large, the server may decide to
	// further limit the number of returned resources.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCampaignDraftAsyncErrorsRequest) Reset()         { *m = ListCampaignDraftAsyncErrorsRequest{} }
func (m *ListCampaignDraftAsyncErrorsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCampaignDraftAsyncErrorsRequest) ProtoMessage()    {}
func (*ListCampaignDraftAsyncErrorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{6}
}

func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Unmarshal(m, b)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Marshal(b, m, deterministic)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Merge(m, src)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.Size(m)
}
func (m *ListCampaignDraftAsyncErrorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCampaignDraftAsyncErrorsRequest proto.InternalMessageInfo

func (m *ListCampaignDraftAsyncErrorsRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ListCampaignDraftAsyncErrorsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCampaignDraftAsyncErrorsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v3.services.CampaignDraftService.ListCampaignDraftAsyncErrors].
type ListCampaignDraftAsyncErrorsResponse struct {
	// Details of the errors when performing the asynchronous operation.
	Errors []*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	// Pagination token used to retrieve the next page of results.
	// Pass the content of this string as the `page_token` attribute of
	// the next request. `next_page_token` is not returned for the last
	// page.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCampaignDraftAsyncErrorsResponse) Reset()         { *m = ListCampaignDraftAsyncErrorsResponse{} }
func (m *ListCampaignDraftAsyncErrorsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCampaignDraftAsyncErrorsResponse) ProtoMessage()    {}
func (*ListCampaignDraftAsyncErrorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01b2400b8d9061a5, []int{7}
}

func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Unmarshal(m, b)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Marshal(b, m, deterministic)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Merge(m, src)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.Size(m)
}
func (m *ListCampaignDraftAsyncErrorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCampaignDraftAsyncErrorsResponse proto.InternalMessageInfo

func (m *ListCampaignDraftAsyncErrorsResponse) GetErrors() []*status.Status {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *ListCampaignDraftAsyncErrorsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignDraftRequest)(nil), "google.ads.googleads.v3.services.GetCampaignDraftRequest")
	proto.RegisterType((*MutateCampaignDraftsRequest)(nil), "google.ads.googleads.v3.services.MutateCampaignDraftsRequest")
	proto.RegisterType((*PromoteCampaignDraftRequest)(nil), "google.ads.googleads.v3.services.PromoteCampaignDraftRequest")
	proto.RegisterType((*CampaignDraftOperation)(nil), "google.ads.googleads.v3.services.CampaignDraftOperation")
	proto.RegisterType((*MutateCampaignDraftsResponse)(nil), "google.ads.googleads.v3.services.MutateCampaignDraftsResponse")
	proto.RegisterType((*MutateCampaignDraftResult)(nil), "google.ads.googleads.v3.services.MutateCampaignDraftResult")
	proto.RegisterType((*ListCampaignDraftAsyncErrorsRequest)(nil), "google.ads.googleads.v3.services.ListCampaignDraftAsyncErrorsRequest")
	proto.RegisterType((*ListCampaignDraftAsyncErrorsResponse)(nil), "google.ads.googleads.v3.services.ListCampaignDraftAsyncErrorsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/campaign_draft_service.proto", fileDescriptor_01b2400b8d9061a5)
}

var fileDescriptor_01b2400b8d9061a5 = []byte{
	// 1023 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x6f, 0xdc, 0xc4,
	0x1b, 0xff, 0xdb, 0xf9, 0x93, 0x76, 0x67, 0x9b, 0x16, 0x0d, 0x69, 0xbb, 0xd9, 0xa4, 0x62, 0xe5,
	0x44, 0x25, 0x5a, 0x21, 0xbb, 0xda, 0x95, 0x50, 0xeb, 0x28, 0x55, 0xbd, 0x90, 0x97, 0x22, 0x4a,
	0x23, 0x07, 0x02, 0x42, 0x41, 0xd6, 0xc4, 0x9e, 0x18, 0x2b, 0xf6, 0x8c, 0x99, 0xb1, 0x57, 0x24,
	0x55, 0x25, 0xc4, 0x01, 0x3e, 0x00, 0x7c, 0x02, 0x8e, 0x7c, 0x00, 0x6e, 0x7c, 0x00, 0x72, 0xe5,
	0x96, 0x53, 0x0f, 0x9c, 0x38, 0x20, 0xce, 0x95, 0x90, 0x90, 0xed, 0x99, 0xdd, 0xf5, 0xd6, 0x9b,
	0x6d, 0x83, 0xb8, 0x8d, 0x9f, 0x97, 0xdf, 0xf3, 0xf2, 0x7b, 0x9e, 0x67, 0x17, 0xac, 0xfb, 0x94,
	0xfa, 0x21, 0x36, 0x90, 0xc7, 0x8d, 0xe2, 0x99, 0xbd, 0xfa, 0x5d, 0x83, 0x63, 0xd6, 0x0f, 0x5c,
	0xcc, 0x0d, 0x17, 0x45, 0x31, 0x0a, 0x7c, 0xe2, 0x78, 0x0c, 0x1d, 0x26, 0x8e, 0x90, 0xeb, 0x31,
	0xa3, 0x09, 0x85, 0xad, 0xc2, 0x47, 0x47, 0x1e, 0xd7, 0x07, 0xee, 0x7a, 0xbf, 0xab, 0x4b, 0xf7,
	0xe6, 0x3b, 0x93, 0x02, 0x30, 0xcc, 0x69, 0xca, 0x5e, 0x8c, 0x50, 0x20, 0x37, 0x97, 0xa4, 0x5f,
	0x1c, 0x18, 0x88, 0x10, 0x9a, 0xa0, 0x24, 0xa0, 0x84, 0x0b, 0xed, 0xcd, 0x11, 0xad, 0x1b, 0x06,
	0x98, 0x48, 0xb7, 0x37, 0x47, 0x14, 0x87, 0x01, 0x0e, 0x3d, 0xe7, 0x00, 0x7f, 0x81, 0xfa, 0x01,
	0x65, 0xc2, 0x60, 0x61, 0xc4, 0x40, 0xa6, 0x20, 0x54, 0xcb, 0x42, 0x15, 0x52, 0xe2, 0xb3, 0x94,
	0x90, 0x80, 0xf8, 0x06, 0x8d, 0x31, 0x2b, 0x45, 0x16, 0x15, 0x1b, 0xf9, 0xd7, 0x41, 0x7a, 0x28,
	0xa2, 0x44, 0x88, 0x1f, 0x8d, 0xe5, 0xc6, 0x62, 0xd7, 0xe0, 0x09, 0x4a, 0x52, 0xe1, 0xaa, 0x11,
	0x70, 0x73, 0x0b, 0x27, 0xef, 0x8a, 0x6a, 0xdf, 0xcb, 0x8a, 0xb5, 0xf1, 0x97, 0x29, 0xe6, 0x09,
	0xdc, 0x05, 0x73, 0x32, 0x19, 0x87, 0xa0, 0x08, 0x37, 0x94, 0x96, 0xb2, 0x5a, 0xeb, 0xe9, 0xcf,
	0x2c, 0xf5, 0xb9, 0xb5, 0x0a, 0x6e, 0x0f, 0x7b, 0x2b, 0x5e, 0x71, 0xc0, 0x75, 0x97, 0x46, 0x46,
	0x19, 0xed, 0x8a, 0x04, 0xf9, 0x10, 0x45, 0x58, 0xfb, 0x4b, 0x01, 0x8b, 0x8f, 0xd2, 0x04, 0x25,
	0xb8, 0x64, 0xc5, 0x65, 0xd0, 0x15, 0x50, 0x77, 0x53, 0x9e, 0xd0, 0x08, 0x33, 0x27, 0xf0, 0x44,
	0xc8, 0x99, 0x67, 0x96, 0x6a, 0x03, 0x29, 0x7f, 0xe8, 0xc1, 0xcf, 0x01, 0x18, 0x36, 0xa1, 0xa1,
	0xb6, 0x66, 0x56, 0xeb, 0x9d, 0xbb, 0xfa, 0x34, 0xde, 0xf5, 0x52, 0xc8, 0xc7, 0x12, 0x40, 0xc0,
	0x0f, 0x01, 0xe1, 0x5b, 0xe0, 0x5a, 0x8c, 0x58, 0x12, 0xa0, 0xd0, 0x39, 0x44, 0x41, 0x98, 0x32,
	0xdc, 0x98, 0x69, 0x29, 0xab, 0x97, 0xed, 0xab, 0x42, 0xbc, 0x59, 0x48, 0xe1, 0x32, 0x98, 0xeb,
	0xa3, 0x30, 0xf0, 0x50, 0x82, 0x1d, 0x4a, 0xc2, 0xe3, 0xc6, 0xff, 0x73, 0xb3, 0x2b, 0x52, 0xf8,
	0x98, 0x84, 0xc7, 0xda, 0x43, 0xb0, 0xb8, 0xc3, 0x68, 0x44, 0xc7, 0x4a, 0x96, 0x15, 0xb7, 0xc1,
	0xd5, 0xf2, 0xb0, 0x8d, 0x16, 0x3d, 0xe7, 0x8e, 0xba, 0x68, 0x3f, 0xa8, 0xe0, 0x46, 0x75, 0x11,
	0x70, 0x0d, 0xd4, 0xd3, 0x38, 0x4f, 0x24, 0xa3, 0x3d, 0x4f, 0xa4, 0xde, 0x69, 0xca, 0x9e, 0xc8,
	0xc9, 0xd0, 0x37, 0xb3, 0xc9, 0x78, 0x84, 0xf8, 0x91, 0x0d, 0x0a, 0xf3, 0xec, 0x0d, 0xdf, 0x07,
	0xb3, 0x2e, 0xc3, 0x28, 0x29, 0x38, 0xae, 0x77, 0xee, 0x4c, 0xec, 0xe5, 0x60, 0x43, 0xca, 0xcd,
	0xdc, 0xfe, 0x9f, 0x2d, 0x10, 0x32, 0xac, 0x02, 0xb9, 0xa1, 0x5e, 0x1c, 0xab, 0x40, 0x80, 0x0d,
	0x30, 0xcb, 0x70, 0x44, 0xfb, 0x45, 0xff, 0x6b, 0x99, 0xa6, 0xf8, 0xee, 0xd5, 0x41, 0x6d, 0x40,
	0x98, 0xf6, 0x8b, 0x02, 0x96, 0xaa, 0x87, 0x8a, 0xc7, 0x94, 0x70, 0x0c, 0x37, 0xc1, 0xf5, 0x31,
	0x42, 0x1d, 0xcc, 0x18, 0x65, 0x39, 0x6c, 0xbd, 0x03, 0x65, 0x8a, 0x2c, 0x76, 0xf5, 0xdd, 0x7c,
	0x3d, 0xec, 0x37, 0xca, 0x54, 0x6f, 0x64, 0xe6, 0xf0, 0x63, 0x70, 0x89, 0x61, 0x9e, 0x86, 0x89,
	0x1c, 0xba, 0xb5, 0xe9, 0x43, 0x57, 0x91, 0x98, 0x9d, 0x63, 0xd8, 0x12, 0x4b, 0x7b, 0x00, 0x16,
	0x26, 0x5a, 0x65, 0x33, 0x56, 0xb1, 0x86, 0x63, 0x6b, 0xf5, 0xb3, 0x02, 0x96, 0x3f, 0x08, 0x78,
	0x79, 0x91, 0x2d, 0x7e, 0x4c, 0xdc, 0x3c, 0x71, 0xfe, 0x5f, 0xee, 0x34, 0xbc, 0x05, 0x40, 0x8c,
	0x7c, 0xec, 0x24, 0xf4, 0x08, 0x93, 0x9c, 0xf5, 0x9a, 0x5d, 0xcb, 0x24, 0x1f, 0x65, 0x02, 0xb8,
	0x08, 0xf2, 0x0f, 0x87, 0x07, 0x27, 0x05, 0x8f, 0xaf, 0xd9, 0x97, 0x33, 0xc1, 0x6e, 0x70, 0x82,
	0xb5, 0x13, 0xb0, 0x72, 0x7e, 0xde, 0x82, 0xc1, 0x36, 0x98, 0xcd, 0x19, 0xe3, 0x0d, 0x25, 0x6f,
	0x7c, 0x15, 0x65, 0xc2, 0x02, 0xde, 0x06, 0xd7, 0x08, 0xfe, 0x2a, 0x71, 0x5e, 0x48, 0x6a, 0x2e,
	0x13, 0xef, 0xc8, 0xc4, 0x3a, 0xcf, 0x2f, 0x81, 0xf9, 0x52, 0xe0, 0xdd, 0x82, 0x33, 0xf8, 0xab,
	0x02, 0x5e, 0x1f, 0xbf, 0x8a, 0xf0, 0xde, 0x74, 0xaa, 0x27, 0x5c, 0xd2, 0xe6, 0x2b, 0xaf, 0x80,
	0xb6, 0x7d, 0x66, 0x95, 0x89, 0xfa, 0xe6, 0xb7, 0xdf, 0xbf, 0x57, 0x3b, 0xf0, 0x4e, 0xf6, 0x2b,
	0xf5, 0xa4, 0xa4, 0x59, 0x97, 0x77, 0x91, 0x1b, 0x6d, 0xa3, 0x74, 0x2e, 0xb8, 0xd1, 0x7e, 0x0a,
	0xff, 0x54, 0xc0, 0x7c, 0xd5, 0x6e, 0xc0, 0xf5, 0x0b, 0x8d, 0xae, 0x9c, 0xa4, 0xe6, 0xfd, 0x8b,
	0xba, 0x17, 0x84, 0x6a, 0x9f, 0x9c, 0x59, 0x37, 0x46, 0x2e, 0xfd, 0xdb, 0xc3, 0xf3, 0x9b, 0x97,
	0x7a, 0x57, 0xeb, 0x66, 0xa5, 0x0e, 0x6b, 0x7b, 0x32, 0x62, 0xbc, 0xde, 0x7e, 0x3a, 0x56, 0xa9,
	0x19, 0xe5, 0xb1, 0x4c, 0xa5, 0x0d, 0xff, 0x56, 0xc0, 0x7c, 0xd5, 0xbd, 0x7d, 0x99, 0x82, 0xcf,
	0xb9, 0xd3, 0xcd, 0x5b, 0xd2, 0x7d, 0xe4, 0xa7, 0x58, 0x1f, 0xdc, 0x5f, 0xed, 0x5b, 0xe5, 0xd4,
	0xd2, 0xc1, 0xf5, 0xf1, 0x7b, 0xbb, 0x11, 0xc5, 0xc9, 0x31, 0xac, 0x16, 0x9f, 0x59, 0x63, 0x97,
	0x3f, 0xaf, 0xfc, 0xbe, 0x76, 0x2f, 0x27, 0xb9, 0xac, 0x3a, 0x9f, 0x65, 0x33, 0x2e, 0x12, 0xcf,
	0xea, 0xff, 0x4e, 0x05, 0x4b, 0xe7, 0xad, 0x14, 0xdc, 0x98, 0xde, 0x87, 0x97, 0x38, 0x25, 0xcd,
	0xcd, 0x7f, 0x0b, 0x23, 0x06, 0xe1, 0xd3, 0xca, 0x51, 0xef, 0xc1, 0x07, 0xaf, 0x3a, 0xea, 0x66,
	0x18, 0xf0, 0xd1, 0x08, 0xcd, 0xc5, 0x53, 0xab, 0x31, 0xe9, 0xa0, 0xf5, 0xbe, 0x56, 0xc1, 0x8a,
	0x4b, 0xa3, 0xa9, 0x45, 0xf4, 0x16, 0xaa, 0x4e, 0xc4, 0x4e, 0x46, 0xe4, 0x8e, 0xf2, 0xd9, 0xb6,
	0x70, 0xf7, 0x69, 0x88, 0x88, 0xaf, 0x53, 0xe6, 0x1b, 0x3e, 0x26, 0x39, 0xcd, 0xc6, 0x30, 0xe0,
	0xe4, 0x7f, 0xb2, 0x6b, 0xf2, 0xf1, 0xa3, 0x3a, 0xb3, 0x65, 0x59, 0x3f, 0xa9, 0xad, 0xad, 0x02,
	0xd0, 0xf2, 0xb8, 0x5e, 0x3c, 0xb3, 0xd7, 0x5e, 0x57, 0x17, 0x81, 0xf9, 0xa9, 0x34, 0xd9, 0xb7,
	0x3c, 0xbe, 0x3f, 0x30, 0xd9, 0xdf, 0xeb, 0xee, 0x4b, 0x93, 0x3f, 0xd4, 0x95, 0x42, 0x6e, 0x9a,
	0x96, 0xc7, 0x4d, 0x73, 0x60, 0x64, 0x9a, 0x7b, 0x5d, 0xd3, 0x94, 0x66, 0x07, 0xb3, 0x79, 0x9e,
	0xdd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xdd, 0xee, 0x9b, 0x70, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignDraftServiceClient is the client API for CampaignDraftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignDraftServiceClient interface {
	// Returns the requested campaign draft in full detail.
	GetCampaignDraft(ctx context.Context, in *GetCampaignDraftRequest, opts ...grpc.CallOption) (*resources.CampaignDraft, error)
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v3.services.CampaignDraftService.ListCampaignDraftAsyncErrors] to view the list of
	// error reasons.
	PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error)
}

type campaignDraftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignDraftServiceClient(cc grpc.ClientConnInterface) CampaignDraftServiceClient {
	return &campaignDraftServiceClient{cc}
}

func (c *campaignDraftServiceClient) GetCampaignDraft(ctx context.Context, in *GetCampaignDraftRequest, opts ...grpc.CallOption) (*resources.CampaignDraft, error) {
	out := new(resources.CampaignDraft)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignDraftService/GetCampaignDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) MutateCampaignDrafts(ctx context.Context, in *MutateCampaignDraftsRequest, opts ...grpc.CallOption) (*MutateCampaignDraftsResponse, error) {
	out := new(MutateCampaignDraftsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignDraftService/MutateCampaignDrafts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) PromoteCampaignDraft(ctx context.Context, in *PromoteCampaignDraftRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignDraftService/PromoteCampaignDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *campaignDraftServiceClient) ListCampaignDraftAsyncErrors(ctx context.Context, in *ListCampaignDraftAsyncErrorsRequest, opts ...grpc.CallOption) (*ListCampaignDraftAsyncErrorsResponse, error) {
	out := new(ListCampaignDraftAsyncErrorsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.CampaignDraftService/ListCampaignDraftAsyncErrors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignDraftServiceServer is the server API for CampaignDraftService service.
type CampaignDraftServiceServer interface {
	// Returns the requested campaign draft in full detail.
	GetCampaignDraft(context.Context, *GetCampaignDraftRequest) (*resources.CampaignDraft, error)
	// Creates, updates, or removes campaign drafts. Operation statuses are
	// returned.
	MutateCampaignDrafts(context.Context, *MutateCampaignDraftsRequest) (*MutateCampaignDraftsResponse, error)
	// Promotes the changes in a draft back to the base campaign.
	//
	// This method returns a Long Running Operation (LRO) indicating if the
	// Promote is done. Use [Operations.GetOperation] to poll the LRO until it
	// is done. Only a done status is returned in the response. See the status
	// in the Campaign Draft resource to determine if the promotion was
	// successful. If the LRO failed, use
	// [CampaignDraftService.ListCampaignDraftAsyncErrors][google.ads.googleads.v3.services.CampaignDraftService.ListCampaignDraftAsyncErrors] to view the list of
	// error reasons.
	PromoteCampaignDraft(context.Context, *PromoteCampaignDraftRequest) (*longrunning.Operation, error)
	// Returns all errors that occurred during CampaignDraft promote. Throws an
	// error if called before campaign draft is promoted.
	// Supports standard list paging.
	ListCampaignDraftAsyncErrors(context.Context, *ListCampaignDraftAsyncErrorsRequest) (*ListCampaignDraftAsyncErrorsResponse, error)
}

// UnimplementedCampaignDraftServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignDraftServiceServer struct {
}

func (*UnimplementedCampaignDraftServiceServer) GetCampaignDraft(ctx context.Context, req *GetCampaignDraftRequest) (*resources.CampaignDraft, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetCampaignDraft not implemented")
}
func (*UnimplementedCampaignDraftServiceServer) MutateCampaignDrafts(ctx context.Context, req *MutateCampaignDraftsRequest) (*MutateCampaignDraftsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateCampaignDrafts not implemented")
}
func (*UnimplementedCampaignDraftServiceServer) PromoteCampaignDraft(ctx context.Context, req *PromoteCampaignDraftRequest) (*longrunning.Operation, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method PromoteCampaignDraft not implemented")
}
func (*UnimplementedCampaignDraftServiceServer) ListCampaignDraftAsyncErrors(ctx context.Context, req *ListCampaignDraftAsyncErrorsRequest) (*ListCampaignDraftAsyncErrorsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListCampaignDraftAsyncErrors not implemented")
}

func RegisterCampaignDraftServiceServer(s *grpc.Server, srv CampaignDraftServiceServer) {
	s.RegisterService(&_CampaignDraftService_serviceDesc, srv)
}

func _CampaignDraftService_GetCampaignDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).GetCampaignDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignDraftService/GetCampaignDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).GetCampaignDraft(ctx, req.(*GetCampaignDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_MutateCampaignDrafts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCampaignDraftsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignDraftService/MutateCampaignDrafts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).MutateCampaignDrafts(ctx, req.(*MutateCampaignDraftsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_PromoteCampaignDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoteCampaignDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignDraftService/PromoteCampaignDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).PromoteCampaignDraft(ctx, req.(*PromoteCampaignDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCampaignDraftAsyncErrorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.CampaignDraftService/ListCampaignDraftAsyncErrors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignDraftServiceServer).ListCampaignDraftAsyncErrors(ctx, req.(*ListCampaignDraftAsyncErrorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignDraftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.CampaignDraftService",
	HandlerType: (*CampaignDraftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignDraft",
			Handler:    _CampaignDraftService_GetCampaignDraft_Handler,
		},
		{
			MethodName: "MutateCampaignDrafts",
			Handler:    _CampaignDraftService_MutateCampaignDrafts_Handler,
		},
		{
			MethodName: "PromoteCampaignDraft",
			Handler:    _CampaignDraftService_PromoteCampaignDraft_Handler,
		},
		{
			MethodName: "ListCampaignDraftAsyncErrors",
			Handler:    _CampaignDraftService_ListCampaignDraftAsyncErrors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/campaign_draft_service.proto",
}
