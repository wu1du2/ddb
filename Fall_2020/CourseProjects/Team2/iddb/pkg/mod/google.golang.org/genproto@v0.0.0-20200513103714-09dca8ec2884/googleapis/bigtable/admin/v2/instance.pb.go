// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/v2/instance.proto

package admin

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible states of an instance.
type Instance_State int32

const (
	// The state of the instance could not be determined.
	Instance_STATE_NOT_KNOWN Instance_State = 0
	// The instance has been successfully created and can serve requests
	// to its tables.
	Instance_READY Instance_State = 1
	// The instance is currently being created, and may be destroyed
	// if the creation process encounters an error.
	Instance_CREATING Instance_State = 2
)

var Instance_State_name = map[int32]string{
	0: "STATE_NOT_KNOWN",
	1: "READY",
	2: "CREATING",
}

var Instance_State_value = map[string]int32{
	"STATE_NOT_KNOWN": 0,
	"READY":           1,
	"CREATING":        2,
}

func (x Instance_State) String() string {
	return proto.EnumName(Instance_State_name, int32(x))
}

func (Instance_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{0, 0}
}

// The type of the instance.
type Instance_Type int32

const (
	// The type of the instance is unspecified. If set when creating an
	// instance, a `PRODUCTION` instance will be created. If set when updating
	// an instance, the type will be left unchanged.
	Instance_TYPE_UNSPECIFIED Instance_Type = 0
	// An instance meant for production use. `serve_nodes` must be set
	// on the cluster.
	Instance_PRODUCTION Instance_Type = 1
	// The instance is meant for development and testing purposes only; it has
	// no performance or uptime guarantees and is not covered by SLA.
	// After a development instance is created, it can be upgraded by
	// updating the instance to type `PRODUCTION`. An instance created
	// as a production instance cannot be changed to a development instance.
	// When creating a development instance, `serve_nodes` on the cluster must
	// not be set.
	Instance_DEVELOPMENT Instance_Type = 2
)

var Instance_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "PRODUCTION",
	2: "DEVELOPMENT",
}

var Instance_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"PRODUCTION":       1,
	"DEVELOPMENT":      2,
}

func (x Instance_Type) String() string {
	return proto.EnumName(Instance_Type_name, int32(x))
}

func (Instance_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{0, 1}
}

// Possible states of a cluster.
type Cluster_State int32

const (
	// The state of the cluster could not be determined.
	Cluster_STATE_NOT_KNOWN Cluster_State = 0
	// The cluster has been successfully created and is ready to serve requests.
	Cluster_READY Cluster_State = 1
	// The cluster is currently being created, and may be destroyed
	// if the creation process encounters an error.
	// A cluster may not be able to serve requests while being created.
	Cluster_CREATING Cluster_State = 2
	// The cluster is currently being resized, and may revert to its previous
	// node count if the process encounters an error.
	// A cluster is still capable of serving requests while being resized,
	// but may exhibit performance as if its number of allocated nodes is
	// between the starting and requested states.
	Cluster_RESIZING Cluster_State = 3
	// The cluster has no backing nodes. The data (tables) still
	// exist, but no operations can be performed on the cluster.
	Cluster_DISABLED Cluster_State = 4
)

var Cluster_State_name = map[int32]string{
	0: "STATE_NOT_KNOWN",
	1: "READY",
	2: "CREATING",
	3: "RESIZING",
	4: "DISABLED",
}

var Cluster_State_value = map[string]int32{
	"STATE_NOT_KNOWN": 0,
	"READY":           1,
	"CREATING":        2,
	"RESIZING":        3,
	"DISABLED":        4,
}

func (x Cluster_State) String() string {
	return proto.EnumName(Cluster_State_name, int32(x))
}

func (Cluster_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{1, 0}
}

// A collection of Bigtable [Tables][google.bigtable.admin.v2.Table] and
// the resources that serve them.
// All tables in an instance are served from all
// [Clusters][google.bigtable.admin.v2.Cluster] in the instance.
type Instance struct {
	// The unique name of the instance. Values are of the form
	// `projects/{project}/instances/[a-z][a-z0-9\\-]+[a-z0-9]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The descriptive name for this instance as it appears in UIs.
	// Can be changed at any time, but should be kept globally unique
	// to avoid confusion.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// (`OutputOnly`)
	// The current state of the instance.
	State Instance_State `protobuf:"varint,3,opt,name=state,proto3,enum=google.bigtable.admin.v2.Instance_State" json:"state,omitempty"`
	// The type of the instance. Defaults to `PRODUCTION`.
	Type Instance_Type `protobuf:"varint,4,opt,name=type,proto3,enum=google.bigtable.admin.v2.Instance_Type" json:"type,omitempty"`
	// Labels are a flexible and lightweight mechanism for organizing cloud
	// resources into groups that reflect a customer's organizational needs and
	// deployment strategies. They can be used to filter resources and aggregate
	// metrics.
	//
	// * Label keys must be between 1 and 63 characters long and must conform to
	//   the regular expression: `[\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}`.
	// * Label values must be between 0 and 63 characters long and must conform to
	//   the regular expression: `[\p{Ll}\p{Lo}\p{N}_-]{0,63}`.
	// * No more than 64 labels can be associated with a given resource.
	// * Keys and values must both be under 128 bytes.
	Labels               map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{0}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Instance) GetState() Instance_State {
	if m != nil {
		return m.State
	}
	return Instance_STATE_NOT_KNOWN
}

func (m *Instance) GetType() Instance_Type {
	if m != nil {
		return m.Type
	}
	return Instance_TYPE_UNSPECIFIED
}

func (m *Instance) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// A resizable group of nodes in a particular cloud location, capable
// of serving all [Tables][google.bigtable.admin.v2.Table] in the parent
// [Instance][google.bigtable.admin.v2.Instance].
type Cluster struct {
	// The unique name of the cluster. Values are of the form
	// `projects/{project}/instances/{instance}/clusters/[a-z][-a-z0-9]*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (`CreationOnly`)
	// The location where this cluster's nodes and storage reside. For best
	// performance, clients should be located as close as possible to this
	// cluster. Currently only zones are supported, so values should be of the
	// form `projects/{project}/locations/{zone}`.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// The current state of the cluster.
	State Cluster_State `protobuf:"varint,3,opt,name=state,proto3,enum=google.bigtable.admin.v2.Cluster_State" json:"state,omitempty"`
	// Required. The number of nodes allocated to this cluster. More nodes enable
	// higher throughput and more consistent performance.
	ServeNodes int32 `protobuf:"varint,4,opt,name=serve_nodes,json=serveNodes,proto3" json:"serve_nodes,omitempty"`
	// (`CreationOnly`)
	// The type of storage used by this cluster to serve its
	// parent instance's tables, unless explicitly overridden.
	DefaultStorageType   StorageType `protobuf:"varint,5,opt,name=default_storage_type,json=defaultStorageType,proto3,enum=google.bigtable.admin.v2.StorageType" json:"default_storage_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{1}
}

func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (m *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(m, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Cluster) GetState() Cluster_State {
	if m != nil {
		return m.State
	}
	return Cluster_STATE_NOT_KNOWN
}

func (m *Cluster) GetServeNodes() int32 {
	if m != nil {
		return m.ServeNodes
	}
	return 0
}

func (m *Cluster) GetDefaultStorageType() StorageType {
	if m != nil {
		return m.DefaultStorageType
	}
	return StorageType_STORAGE_TYPE_UNSPECIFIED
}

// A configuration object describing how Cloud Bigtable should treat traffic
// from a particular end user application.
type AppProfile struct {
	// (`OutputOnly`)
	// The unique name of the app profile. Values are of the form
	// `projects/<project>/instances/<instance>/appProfiles/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Strongly validated etag for optimistic concurrency control. Preserve the
	// value returned from `GetAppProfile` when calling `UpdateAppProfile` to
	// fail the request if there has been a modification in the mean time. The
	// `update_mask` of the request need not include `etag` for this protection
	// to apply.
	// See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and
	// [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more
	// details.
	Etag string `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional long form description of the use case for this AppProfile.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The routing policy for all read/write requests that use this app profile.
	// A value must be explicitly set.
	//
	// Types that are valid to be assigned to RoutingPolicy:
	//	*AppProfile_MultiClusterRoutingUseAny_
	//	*AppProfile_SingleClusterRouting_
	RoutingPolicy        isAppProfile_RoutingPolicy `protobuf_oneof:"routing_policy"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AppProfile) Reset()         { *m = AppProfile{} }
func (m *AppProfile) String() string { return proto.CompactTextString(m) }
func (*AppProfile) ProtoMessage()    {}
func (*AppProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{2}
}

func (m *AppProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppProfile.Unmarshal(m, b)
}
func (m *AppProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppProfile.Marshal(b, m, deterministic)
}
func (m *AppProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProfile.Merge(m, src)
}
func (m *AppProfile) XXX_Size() int {
	return xxx_messageInfo_AppProfile.Size(m)
}
func (m *AppProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProfile.DiscardUnknown(m)
}

var xxx_messageInfo_AppProfile proto.InternalMessageInfo

func (m *AppProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppProfile) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *AppProfile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type isAppProfile_RoutingPolicy interface {
	isAppProfile_RoutingPolicy()
}

type AppProfile_MultiClusterRoutingUseAny_ struct {
	MultiClusterRoutingUseAny *AppProfile_MultiClusterRoutingUseAny `protobuf:"bytes,5,opt,name=multi_cluster_routing_use_any,json=multiClusterRoutingUseAny,proto3,oneof"`
}

type AppProfile_SingleClusterRouting_ struct {
	SingleClusterRouting *AppProfile_SingleClusterRouting `protobuf:"bytes,6,opt,name=single_cluster_routing,json=singleClusterRouting,proto3,oneof"`
}

func (*AppProfile_MultiClusterRoutingUseAny_) isAppProfile_RoutingPolicy() {}

func (*AppProfile_SingleClusterRouting_) isAppProfile_RoutingPolicy() {}

func (m *AppProfile) GetRoutingPolicy() isAppProfile_RoutingPolicy {
	if m != nil {
		return m.RoutingPolicy
	}
	return nil
}

func (m *AppProfile) GetMultiClusterRoutingUseAny() *AppProfile_MultiClusterRoutingUseAny {
	if x, ok := m.GetRoutingPolicy().(*AppProfile_MultiClusterRoutingUseAny_); ok {
		return x.MultiClusterRoutingUseAny
	}
	return nil
}

func (m *AppProfile) GetSingleClusterRouting() *AppProfile_SingleClusterRouting {
	if x, ok := m.GetRoutingPolicy().(*AppProfile_SingleClusterRouting_); ok {
		return x.SingleClusterRouting
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AppProfile) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AppProfile_MultiClusterRoutingUseAny_)(nil),
		(*AppProfile_SingleClusterRouting_)(nil),
	}
}

// Read/write requests are routed to the nearest cluster in the instance, and
// will fail over to the nearest cluster that is available in the event of
// transient errors or delays. Clusters in a region are considered
// equidistant. Choosing this option sacrifices read-your-writes consistency
// to improve availability.
type AppProfile_MultiClusterRoutingUseAny struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppProfile_MultiClusterRoutingUseAny) Reset()         { *m = AppProfile_MultiClusterRoutingUseAny{} }
func (m *AppProfile_MultiClusterRoutingUseAny) String() string { return proto.CompactTextString(m) }
func (*AppProfile_MultiClusterRoutingUseAny) ProtoMessage()    {}
func (*AppProfile_MultiClusterRoutingUseAny) Descriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{2, 0}
}

func (m *AppProfile_MultiClusterRoutingUseAny) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny.Unmarshal(m, b)
}
func (m *AppProfile_MultiClusterRoutingUseAny) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny.Marshal(b, m, deterministic)
}
func (m *AppProfile_MultiClusterRoutingUseAny) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny.Merge(m, src)
}
func (m *AppProfile_MultiClusterRoutingUseAny) XXX_Size() int {
	return xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny.Size(m)
}
func (m *AppProfile_MultiClusterRoutingUseAny) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny.DiscardUnknown(m)
}

var xxx_messageInfo_AppProfile_MultiClusterRoutingUseAny proto.InternalMessageInfo

// Unconditionally routes all read/write requests to a specific cluster.
// This option preserves read-your-writes consistency but does not improve
// availability.
type AppProfile_SingleClusterRouting struct {
	// The cluster to which read/write requests should be routed.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Whether or not `CheckAndMutateRow` and `ReadModifyWriteRow` requests are
	// allowed by this app profile. It is unsafe to send these requests to
	// the same table/row/column in multiple clusters.
	AllowTransactionalWrites bool     `protobuf:"varint,2,opt,name=allow_transactional_writes,json=allowTransactionalWrites,proto3" json:"allow_transactional_writes,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *AppProfile_SingleClusterRouting) Reset()         { *m = AppProfile_SingleClusterRouting{} }
func (m *AppProfile_SingleClusterRouting) String() string { return proto.CompactTextString(m) }
func (*AppProfile_SingleClusterRouting) ProtoMessage()    {}
func (*AppProfile_SingleClusterRouting) Descriptor() ([]byte, []int) {
	return fileDescriptor_712127d2a900984d, []int{2, 1}
}

func (m *AppProfile_SingleClusterRouting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppProfile_SingleClusterRouting.Unmarshal(m, b)
}
func (m *AppProfile_SingleClusterRouting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppProfile_SingleClusterRouting.Marshal(b, m, deterministic)
}
func (m *AppProfile_SingleClusterRouting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppProfile_SingleClusterRouting.Merge(m, src)
}
func (m *AppProfile_SingleClusterRouting) XXX_Size() int {
	return xxx_messageInfo_AppProfile_SingleClusterRouting.Size(m)
}
func (m *AppProfile_SingleClusterRouting) XXX_DiscardUnknown() {
	xxx_messageInfo_AppProfile_SingleClusterRouting.DiscardUnknown(m)
}

var xxx_messageInfo_AppProfile_SingleClusterRouting proto.InternalMessageInfo

func (m *AppProfile_SingleClusterRouting) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *AppProfile_SingleClusterRouting) GetAllowTransactionalWrites() bool {
	if m != nil {
		return m.AllowTransactionalWrites
	}
	return false
}

func init() {
	proto.RegisterEnum("google.bigtable.admin.v2.Instance_State", Instance_State_name, Instance_State_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Instance_Type", Instance_Type_name, Instance_Type_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Cluster_State", Cluster_State_name, Cluster_State_value)
	proto.RegisterType((*Instance)(nil), "google.bigtable.admin.v2.Instance")
	proto.RegisterMapType((map[string]string)(nil), "google.bigtable.admin.v2.Instance.LabelsEntry")
	proto.RegisterType((*Cluster)(nil), "google.bigtable.admin.v2.Cluster")
	proto.RegisterType((*AppProfile)(nil), "google.bigtable.admin.v2.AppProfile")
	proto.RegisterType((*AppProfile_MultiClusterRoutingUseAny)(nil), "google.bigtable.admin.v2.AppProfile.MultiClusterRoutingUseAny")
	proto.RegisterType((*AppProfile_SingleClusterRouting)(nil), "google.bigtable.admin.v2.AppProfile.SingleClusterRouting")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/v2/instance.proto", fileDescriptor_712127d2a900984d)
}

var fileDescriptor_712127d2a900984d = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x18, 0xdd, 0xfc, 0x2d, 0xbb, 0x5f, 0xca, 0xd6, 0x1a, 0x56, 0x90, 0x0d, 0x94, 0x86, 0x40, 0xbb,
	0x7b, 0x65, 0x4b, 0x41, 0x48, 0x34, 0xa5, 0x95, 0x9c, 0xc4, 0x6d, 0xa3, 0x6e, 0x9d, 0xe0, 0x78,
	0x77, 0xd5, 0x6a, 0x91, 0x99, 0xd8, 0xb3, 0xc6, 0x65, 0xe2, 0x31, 0x1e, 0x27, 0x55, 0xb4, 0xda,
	0x1b, 0xde, 0x80, 0xd7, 0xe0, 0x8a, 0xe7, 0xe0, 0x21, 0x2a, 0xae, 0xf7, 0x11, 0xb8, 0x42, 0x1e,
	0x8f, 0xd3, 0xb4, 0xd4, 0x28, 0xe2, 0x6e, 0xe6, 0x9b, 0x73, 0xbe, 0xbf, 0x73, 0x64, 0xc3, 0xa1,
	0xcf, 0x98, 0x4f, 0x89, 0x36, 0x0d, 0xfc, 0x04, 0x4f, 0x29, 0xd1, 0xb0, 0x37, 0x0b, 0x42, 0x6d,
	0xd1, 0xd1, 0x82, 0x90, 0x27, 0x38, 0x74, 0x89, 0x1a, 0xc5, 0x2c, 0x61, 0xa8, 0x91, 0x01, 0xd5,
	0x1c, 0xa8, 0x0a, 0xa0, 0xba, 0xe8, 0x34, 0x6f, 0xcb, 0x14, 0x38, 0x0a, 0xb4, 0x8b, 0x80, 0x50,
	0xcf, 0x99, 0x92, 0x9f, 0xf0, 0x22, 0x60, 0x71, 0x46, 0x6d, 0x1e, 0xac, 0x01, 0x62, 0xc2, 0xd9,
	0x3c, 0xce, 0xb3, 0x36, 0xef, 0x14, 0x96, 0x77, 0xd9, 0x6c, 0xc6, 0xc2, 0x0c, 0xd6, 0xfe, 0xad,
	0x0a, 0x3b, 0x43, 0xd9, 0x0f, 0xfa, 0x04, 0xaa, 0x21, 0x9e, 0x91, 0x46, 0xa9, 0x55, 0x3a, 0xda,
	0xed, 0x55, 0xfe, 0xd2, 0x2b, 0x96, 0x08, 0xa0, 0xbb, 0x70, 0xc3, 0x0b, 0x78, 0x44, 0xf1, 0xd2,
	0x11, 0x80, 0x72, 0x0e, 0x28, 0x5b, 0x75, 0xf9, 0x60, 0xa6, 0xb8, 0x87, 0x50, 0xe3, 0x09, 0x4e,
	0x48, 0xa3, 0xd2, 0x2a, 0x1d, 0xed, 0x75, 0x8e, 0xd4, 0xa2, 0xd1, 0xd4, 0xbc, 0xa6, 0x3a, 0x49,
	0xf1, 0x56, 0x46, 0x43, 0xf7, 0xa1, 0x9a, 0x2c, 0x23, 0xd2, 0xa8, 0x0a, 0xfa, 0xe1, 0x06, 0x74,
	0x7b, 0x19, 0x11, 0x4b, 0x90, 0xd0, 0x23, 0xd8, 0xa6, 0x78, 0x4a, 0x28, 0x6f, 0xd4, 0x5a, 0x95,
	0xa3, 0x7a, 0x47, 0xdd, 0x80, 0x7e, 0x2c, 0x08, 0x46, 0x98, 0xc4, 0x4b, 0x4b, 0xb2, 0x9b, 0xf7,
	0xa0, 0xbe, 0x16, 0x46, 0x0a, 0x54, 0x7e, 0x26, 0xcb, 0x6c, 0x27, 0x56, 0x7a, 0x44, 0xfb, 0x50,
	0x5b, 0x60, 0x3a, 0x97, 0x6b, 0xb0, 0xb2, 0x4b, 0xb7, 0xfc, 0x6d, 0xa9, 0xfd, 0x0d, 0xd4, 0xc4,
	0x3c, 0xe8, 0x23, 0xb8, 0x39, 0xb1, 0x75, 0xdb, 0x70, 0xcc, 0x91, 0xed, 0x3c, 0x35, 0x47, 0x67,
	0xa6, 0xb2, 0x85, 0x76, 0xa1, 0x66, 0x19, 0xfa, 0xe0, 0xb9, 0x52, 0x42, 0x37, 0x60, 0xa7, 0x6f,
	0x19, 0xba, 0x3d, 0x34, 0x1f, 0x2b, 0xe5, 0xf6, 0x03, 0xa8, 0xa6, 0x73, 0xa0, 0x7d, 0x50, 0xec,
	0xe7, 0x63, 0xc3, 0x39, 0x31, 0x27, 0x63, 0xa3, 0x3f, 0x7c, 0x34, 0x34, 0x06, 0xca, 0x16, 0xda,
	0x03, 0x18, 0x5b, 0xa3, 0xc1, 0x49, 0xdf, 0x1e, 0x8e, 0x4c, 0xa5, 0x84, 0x6e, 0x42, 0x7d, 0x60,
	0x9c, 0x1a, 0xc7, 0xa3, 0xf1, 0x33, 0xc3, 0xb4, 0x95, 0x72, 0xd7, 0xbc, 0xd6, 0x9f, 0x42, 0x6b,
	0x35, 0x65, 0x36, 0x35, 0x8e, 0x02, 0xae, 0xba, 0x6c, 0xa6, 0xad, 0xd4, 0x3d, 0x8c, 0x62, 0xf6,
	0x92, 0xb8, 0x09, 0xd7, 0x2e, 0xe5, 0xe9, 0x6a, 0x65, 0x45, 0xae, 0x5d, 0xe6, 0xc7, 0xab, 0xf6,
	0xeb, 0x0a, 0x7c, 0xd0, 0xa7, 0x73, 0x9e, 0x90, 0xb8, 0xd8, 0x12, 0x3d, 0xd8, 0xa1, 0xcc, 0xc5,
	0x49, 0xc0, 0x42, 0x69, 0x87, 0xbb, 0x7f, 0xeb, 0x5f, 0xc2, 0x17, 0x79, 0x90, 0xbf, 0xdb, 0xc7,
	0xb1, 0x7c, 0xb0, 0x56, 0x3c, 0xd4, 0x7b, 0xdb, 0x2e, 0xff, 0xa1, 0xb7, 0x6c, 0x27, 0x73, 0x4b,
	0xd6, 0x86, 0xb4, 0xcc, 0x57, 0x50, 0xe7, 0x24, 0x5e, 0x10, 0x27, 0x64, 0x1e, 0xe1, 0xc2, 0x39,
	0xb5, 0xcc, 0x99, 0x20, 0xe2, 0x66, 0x1a, 0x46, 0x67, 0xb0, 0xef, 0x91, 0x0b, 0x3c, 0xa7, 0x89,
	0xc3, 0x13, 0x16, 0x63, 0x9f, 0x38, 0xc2, 0x68, 0x35, 0x51, 0xf8, 0x4e, 0x71, 0xe1, 0x49, 0x86,
	0x16, 0x36, 0x43, 0x32, 0xc5, 0x5a, 0xac, 0xfd, 0xfd, 0xff, 0x52, 0x3c, 0xbd, 0x59, 0xc6, 0x64,
	0xf8, 0x22, 0xbd, 0x55, 0xd2, 0xdb, 0x60, 0x38, 0xd1, 0x7b, 0xc7, 0xc6, 0x40, 0xa9, 0x76, 0x7f,
	0xbc, 0xd6, 0x7f, 0x80, 0xdb, 0x45, 0x72, 0xe6, 0xc2, 0x74, 0x37, 0x54, 0x53, 0x73, 0x33, 0x02,
	0xd7, 0x2e, 0xe5, 0xe9, 0xaa, 0xfd, 0xba, 0x0a, 0xa0, 0x47, 0xd1, 0x38, 0x66, 0x17, 0x01, 0x25,
	0x08, 0xad, 0x6b, 0x2c, 0xe5, 0x45, 0x50, 0x25, 0x09, 0xf6, 0xa5, 0xc5, 0xc5, 0x19, 0xb5, 0xa0,
	0xee, 0x11, 0xee, 0xc6, 0x41, 0x24, 0x54, 0xaf, 0x88, 0xa7, 0xf5, 0x10, 0xfa, 0xb5, 0x04, 0xb7,
	0x66, 0x73, 0x9a, 0x04, 0x8e, 0xac, 0xe5, 0xc4, 0x6c, 0x9e, 0x04, 0xa1, 0xef, 0xcc, 0x39, 0x71,
	0x70, 0xb8, 0x14, 0x0b, 0xaf, 0x77, 0x1e, 0x16, 0x2f, 0xfc, 0x4d, 0x5f, 0xea, 0xb3, 0x34, 0x93,
	0x9c, 0xd7, 0xca, 0xf2, 0x9c, 0x70, 0xa2, 0x87, 0xcb, 0x27, 0x5b, 0xd6, 0xc1, 0xac, 0xe8, 0x11,
	0xfd, 0x02, 0x1f, 0xf3, 0x20, 0xf4, 0x29, 0x79, 0xb7, 0x89, 0xc6, 0xb6, 0x28, 0x7e, 0x6f, 0xa3,
	0xe2, 0x13, 0x91, 0xe2, 0xed, 0x02, 0x4f, 0xb6, 0xac, 0x7d, 0xfe, 0x9e, 0x78, 0xf3, 0x53, 0x38,
	0x28, 0x6c, 0xb6, 0xc9, 0x61, 0xff, 0x7d, 0xc9, 0xd0, 0x2d, 0x80, 0xbc, 0xc1, 0xc0, 0x93, 0xcb,
	0xdf, 0x95, 0x91, 0xa1, 0x87, 0xbe, 0x83, 0x26, 0xa6, 0x94, 0xbd, 0x72, 0x92, 0x18, 0x87, 0x1c,
	0xbb, 0xe9, 0x82, 0x31, 0x75, 0x5e, 0xc5, 0x41, 0x42, 0xb8, 0xd0, 0x65, 0xc7, 0x6a, 0x08, 0x84,
	0xbd, 0x0e, 0x38, 0x13, 0xef, 0xdd, 0x97, 0xd7, 0xba, 0x0f, 0xed, 0x22, 0x13, 0xad, 0x89, 0xaf,
	0x6f, 0xea, 0x23, 0xbc, 0xe2, 0x70, 0xed, 0x12, 0x47, 0x91, 0x13, 0x65, 0xb7, 0xab, 0x9e, 0x02,
	0x7b, 0xb9, 0xcc, 0x11, 0xa3, 0x81, 0xbb, 0xec, 0xfd, 0x51, 0x82, 0xcf, 0x5c, 0x36, 0x2b, 0x5c,
	0x74, 0xef, 0xc3, 0xfc, 0xab, 0x34, 0x4e, 0xff, 0x42, 0xe3, 0xd2, 0x8b, 0x07, 0x12, 0xea, 0x33,
	0x8a, 0x43, 0x5f, 0x65, 0xb1, 0xaf, 0xf9, 0x24, 0x14, 0xff, 0x28, 0xed, 0x4d, 0xf3, 0xff, 0xfe,
	0x9b, 0xdd, 0x17, 0x87, 0xdf, 0xcb, 0x9f, 0x3f, 0xce, 0xf8, 0x7d, 0xca, 0xe6, 0x9e, 0xda, 0xcb,
	0x0b, 0xea, 0xa2, 0xe0, 0x69, 0xe7, 0xcf, 0x1c, 0x70, 0x2e, 0x00, 0xe7, 0x39, 0xe0, 0x5c, 0x00,
	0xce, 0x4f, 0x3b, 0xd3, 0x6d, 0x51, 0xeb, 0xeb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x78,
	0x6f, 0xa9, 0xb7, 0x07, 0x00, 0x00,
}
