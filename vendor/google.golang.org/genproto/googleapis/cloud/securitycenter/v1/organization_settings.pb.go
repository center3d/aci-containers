// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/organization_settings.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The mode of inclusion when running Asset Discovery.
// Asset discovery can be limited by explicitly identifying projects to be
// included or excluded. If INCLUDE_ONLY is set, then only those projects
// within the organization and their children are discovered during asset
// discovery. If EXCLUDE is set, then projects that don't match those
// projects are discovered during asset discovery. If neither are set, then
// all projects within the organization are discovered during asset
// discovery.
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode int32

const (
	// Unspecified. Setting the mode with this value will disable
	// inclusion/exclusion filtering for Asset Discovery.
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 0
	// Asset Discovery will capture only the resources within the projects
	// specified. All other resources will be ignored.
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 1
	// Asset Discovery will ignore all resources under the projects specified.
	// All other resources will be retrieved.
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 2
)

var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name = map[int32]string{
	0: "INCLUSION_MODE_UNSPECIFIED",
	1: "INCLUDE_ONLY",
	2: "EXCLUDE",
}
var OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value = map[string]int32{
	"INCLUSION_MODE_UNSPECIFIED": 0,
	"INCLUDE_ONLY":               1,
	"EXCLUDE":                    2,
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) String() string {
	return proto.EnumName(OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, int32(x))
}
func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_b59e0a7b2dd79a5c, []int{0, 0, 0}
}

// User specified settings that are attached to the Cloud Security Command
// Center (Cloud SCC) organization.
type OrganizationSettings struct {
	// The relative resource name of the settings. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/organizationSettings".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	// set to `true`, then discovery of assets will occur. If it is set to `false,
	// all historical assets will remain, but discovery of future assets will not
	// occur.
	EnableAssetDiscovery bool `protobuf:"varint,2,opt,name=enable_asset_discovery,json=enableAssetDiscovery,proto3" json:"enable_asset_discovery,omitempty"`
	// The configuration used for Asset Discovery runs.
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `protobuf:"bytes,3,opt,name=asset_discovery_config,json=assetDiscoveryConfig,proto3" json:"asset_discovery_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *OrganizationSettings) Reset()         { *m = OrganizationSettings{} }
func (m *OrganizationSettings) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings) ProtoMessage()    {}
func (*OrganizationSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_b59e0a7b2dd79a5c, []int{0}
}
func (m *OrganizationSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings.Unmarshal(m, b)
}
func (m *OrganizationSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings.Merge(dst, src)
}
func (m *OrganizationSettings) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings.Size(m)
}
func (m *OrganizationSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings proto.InternalMessageInfo

func (m *OrganizationSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationSettings) GetEnableAssetDiscovery() bool {
	if m != nil {
		return m.EnableAssetDiscovery
	}
	return false
}

func (m *OrganizationSettings) GetAssetDiscoveryConfig() *OrganizationSettings_AssetDiscoveryConfig {
	if m != nil {
		return m.AssetDiscoveryConfig
	}
	return nil
}

// The configuration used for Asset Discovery runs.
type OrganizationSettings_AssetDiscoveryConfig struct {
	// The project ids to use for filtering asset discovery.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	// The mode to use for filtering asset discovery.
	InclusionMode        OrganizationSettings_AssetDiscoveryConfig_InclusionMode `protobuf:"varint,2,opt,name=inclusion_mode,json=inclusionMode,proto3,enum=google.cloud.securitycenter.v1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode" json:"inclusion_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *OrganizationSettings_AssetDiscoveryConfig) Reset() {
	*m = OrganizationSettings_AssetDiscoveryConfig{}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) String() string { return proto.CompactTextString(m) }
func (*OrganizationSettings_AssetDiscoveryConfig) ProtoMessage()    {}
func (*OrganizationSettings_AssetDiscoveryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_organization_settings_b59e0a7b2dd79a5c, []int{0, 0}
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Unmarshal(m, b)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Marshal(b, m, deterministic)
}
func (dst *OrganizationSettings_AssetDiscoveryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Merge(dst, src)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_Size() int {
	return xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.Size(m)
}
func (m *OrganizationSettings_AssetDiscoveryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationSettings_AssetDiscoveryConfig proto.InternalMessageInfo

func (m *OrganizationSettings_AssetDiscoveryConfig) GetProjectIds() []string {
	if m != nil {
		return m.ProjectIds
	}
	return nil
}

func (m *OrganizationSettings_AssetDiscoveryConfig) GetInclusionMode() OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	if m != nil {
		return m.InclusionMode
	}
	return OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*OrganizationSettings)(nil), "google.cloud.securitycenter.v1.OrganizationSettings")
	proto.RegisterType((*OrganizationSettings_AssetDiscoveryConfig)(nil), "google.cloud.securitycenter.v1.OrganizationSettings.AssetDiscoveryConfig")
	proto.RegisterEnum("google.cloud.securitycenter.v1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode", OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name, OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value)
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/organization_settings.proto", fileDescriptor_organization_settings_b59e0a7b2dd79a5c)
}

var fileDescriptor_organization_settings_b59e0a7b2dd79a5c = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x27, 0x6d, 0x37, 0x3a, 0x97, 0x4d, 0x91, 0x55, 0x4d, 0x55, 0x85, 0x4a, 0xd5, 0x53, 0x4f,
	0x8e, 0x3a, 0x38, 0xc1, 0x09, 0xd2, 0x80, 0x32, 0x75, 0xe9, 0x94, 0x6a, 0xe3, 0x8f, 0x2a, 0x45,
	0x9e, 0x63, 0x2c, 0xa3, 0xd4, 0x5f, 0x14, 0xbb, 0x95, 0xc6, 0x01, 0x8e, 0xbc, 0x0b, 0x0f, 0xc0,
	0x43, 0xf0, 0x42, 0x5c, 0xd1, 0x9c, 0x54, 0x6a, 0x50, 0x81, 0xcb, 0x6e, 0xc9, 0xef, 0xef, 0xf7,
	0x59, 0x1f, 0x7a, 0x2e, 0x00, 0x44, 0xc6, 0x3d, 0x96, 0xc1, 0x3a, 0xf5, 0x34, 0x67, 0xeb, 0x42,
	0x9a, 0x5b, 0xc6, 0x95, 0xe1, 0x85, 0xb7, 0x99, 0x78, 0x50, 0x08, 0xaa, 0xe4, 0x67, 0x6a, 0x24,
	0xa8, 0x44, 0x73, 0x63, 0xa4, 0x12, 0x9a, 0xe4, 0x05, 0x18, 0xc0, 0x83, 0xd2, 0x4b, 0xac, 0x97,
	0xd4, 0xbd, 0x64, 0x33, 0xe9, 0x3f, 0xae, 0xb2, 0x69, 0x2e, 0x3d, 0xaa, 0x14, 0x18, 0x9b, 0x52,
	0xb9, 0x47, 0xbf, 0x9a, 0xa8, 0x3b, 0xdf, 0x49, 0x5f, 0x54, 0xe1, 0x18, 0xa3, 0x96, 0xa2, 0x2b,
	0xde, 0x73, 0x86, 0xce, 0xf8, 0x28, 0xb6, 0xdf, 0xf8, 0x19, 0x3a, 0xe5, 0x8a, 0xde, 0x64, 0x3c,
	0xa1, 0x5a, 0x73, 0x93, 0xa4, 0x52, 0x33, 0xd8, 0xf0, 0xe2, 0xb6, 0xd7, 0x18, 0x3a, 0xe3, 0x76,
	0xdc, 0x2d, 0xd9, 0x97, 0x77, 0xe4, 0x74, 0xcb, 0xe1, 0xaf, 0xe8, 0xf4, 0x0f, 0x79, 0xc2, 0x40,
	0x7d, 0x94, 0xa2, 0xd7, 0x1c, 0x3a, 0xe3, 0xce, 0x59, 0x48, 0xfe, 0xbd, 0x01, 0xd9, 0x37, 0x1f,
	0xa9, 0x97, 0xf8, 0x36, 0x30, 0xee, 0xd2, 0x3d, 0x68, 0xff, 0x5b, 0x03, 0x75, 0xf7, 0xc9, 0xf1,
	0x13, 0xd4, 0xc9, 0x0b, 0xf8, 0xc4, 0x99, 0x49, 0x64, 0xaa, 0x7b, 0xce, 0xb0, 0x39, 0x3e, 0x8a,
	0x51, 0x05, 0x85, 0xa9, 0xc6, 0x5f, 0xd0, 0x89, 0x54, 0x2c, 0x5b, 0xeb, 0xbb, 0x77, 0x5f, 0x41,
	0xca, 0xed, 0xa2, 0x27, 0x67, 0x6f, 0xef, 0x6d, 0x64, 0x12, 0x6e, 0xf3, 0x2f, 0x20, 0xe5, 0xf1,
	0xb1, 0xdc, 0xfd, 0x1d, 0x45, 0xe8, 0xb8, 0xc6, 0xe3, 0x01, 0xea, 0x87, 0x91, 0x3f, 0xbb, 0x5a,
	0x84, 0xf3, 0x28, 0xb9, 0x98, 0x4f, 0x83, 0xe4, 0x2a, 0x5a, 0x5c, 0x06, 0x7e, 0xf8, 0x3a, 0x0c,
	0xa6, 0xee, 0x03, 0xec, 0xa2, 0x47, 0x96, 0x9f, 0x06, 0xc9, 0x3c, 0x9a, 0xbd, 0x77, 0x1d, 0xdc,
	0x41, 0x0f, 0x83, 0x77, 0x16, 0x71, 0x1b, 0xe7, 0xad, 0x76, 0xcb, 0x3d, 0x38, 0x6f, 0xb5, 0x0f,
	0xdc, 0xc3, 0x57, 0x3f, 0x1c, 0x34, 0x62, 0xb0, 0xfa, 0xcf, 0x26, 0x97, 0xce, 0x87, 0x59, 0xa5,
	0x10, 0x90, 0x51, 0x25, 0x08, 0x14, 0xc2, 0x13, 0x5c, 0xd9, 0xf3, 0xf1, 0x4a, 0x8a, 0xe6, 0x52,
	0xff, 0xed, 0x76, 0x5f, 0xd4, 0x91, 0xef, 0x8d, 0xc1, 0x9b, 0x32, 0xce, 0xb7, 0x85, 0x8b, 0x8a,
	0xf5, 0xcb, 0xc2, 0xeb, 0xc9, 0xcf, 0xad, 0x60, 0x69, 0x05, 0xcb, 0xba, 0x60, 0x79, 0x3d, 0xb9,
	0x39, 0xb4, 0xd5, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x01, 0xcb, 0xb6, 0x35, 0x03,
	0x00, 0x00,
}
