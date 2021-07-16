// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: config.proto

package sdcoreConfig

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILURE Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type PlmnId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mcc string `protobuf:"bytes,1,opt,name=mcc,proto3" json:"mcc,omitempty"`
	Mnc string `protobuf:"bytes,2,opt,name=mnc,proto3" json:"mnc,omitempty"`
}

func (x *PlmnId) Reset() {
	*x = PlmnId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlmnId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlmnId) ProtoMessage() {}

func (x *PlmnId) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlmnId.ProtoReflect.Descriptor instead.
func (*PlmnId) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *PlmnId) GetMcc() string {
	if x != nil {
		return x.Mcc
	}
	return ""
}

func (x *PlmnId) GetMnc() string {
	if x != nil {
		return x.Mnc
	}
	return ""
}

type NSSAI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sst string `protobuf:"bytes,1,opt,name=Sst,proto3" json:"Sst,omitempty"`
	Sd  string `protobuf:"bytes,2,opt,name=Sd,proto3" json:"Sd,omitempty"`
}

func (x *NSSAI) Reset() {
	*x = NSSAI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSSAI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSSAI) ProtoMessage() {}

func (x *NSSAI) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSSAI.ProtoReflect.Descriptor instead.
func (*NSSAI) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *NSSAI) GetSst() string {
	if x != nil {
		return x.Sst
	}
	return ""
}

func (x *NSSAI) GetSd() string {
	if x != nil {
		return x.Sd
	}
	return ""
}

type QoS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uplink       int32  `protobuf:"varint,1,opt,name=uplink,proto3" json:"uplink,omitempty"`
	Downlink     int32  `protobuf:"varint,2,opt,name=downlink,proto3" json:"downlink,omitempty"`
	TrafficClass string `protobuf:"bytes,3,opt,name=trafficClass,proto3" json:"trafficClass,omitempty"`
}

func (x *QoS) Reset() {
	*x = QoS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QoS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QoS) ProtoMessage() {}

func (x *QoS) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QoS.ProtoReflect.Descriptor instead.
func (*QoS) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *QoS) GetUplink() int32 {
	if x != nil {
		return x.Uplink
	}
	return 0
}

func (x *QoS) GetDownlink() int32 {
	if x != nil {
		return x.Downlink
	}
	return 0
}

func (x *QoS) GetTrafficClass() string {
	if x != nil {
		return x.TrafficClass
	}
	return ""
}

type GNodeB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Tac  int32  `protobuf:"varint,2,opt,name=Tac,proto3" json:"Tac,omitempty"`
}

func (x *GNodeB) Reset() {
	*x = GNodeB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GNodeB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GNodeB) ProtoMessage() {}

func (x *GNodeB) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GNodeB.ProtoReflect.Descriptor instead.
func (*GNodeB) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *GNodeB) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GNodeB) GetTac() int32 {
	if x != nil {
		return x.Tac
	}
	return 0
}

type UpfInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpfName string `protobuf:"bytes,1,opt,name=UpfName,proto3" json:"UpfName,omitempty"`
	UpfPort uint32 `protobuf:"varint,2,opt,name=UpfPort,proto3" json:"UpfPort,omitempty"`
}

func (x *UpfInfo) Reset() {
	*x = UpfInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpfInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpfInfo) ProtoMessage() {}

func (x *UpfInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpfInfo.ProtoReflect.Descriptor instead.
func (*UpfInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *UpfInfo) GetUpfName() string {
	if x != nil {
		return x.UpfName
	}
	return ""
}

func (x *UpfInfo) GetUpfPort() uint32 {
	if x != nil {
		return x.UpfPort
	}
	return 0
}

type SiteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteName string    `protobuf:"bytes,1,opt,name=SiteName,proto3" json:"SiteName,omitempty"`
	Gnb      []*GNodeB `protobuf:"bytes,2,rep,name=Gnb,proto3" json:"Gnb,omitempty"`
	Plmn     *PlmnId   `protobuf:"bytes,3,opt,name=Plmn,proto3" json:"Plmn,omitempty"`
	Upf      *UpfInfo  `protobuf:"bytes,4,opt,name=Upf,proto3" json:"Upf,omitempty"`
}

func (x *SiteInfo) Reset() {
	*x = SiteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteInfo) ProtoMessage() {}

func (x *SiteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteInfo.ProtoReflect.Descriptor instead.
func (*SiteInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *SiteInfo) GetSiteName() string {
	if x != nil {
		return x.SiteName
	}
	return ""
}

func (x *SiteInfo) GetGnb() []*GNodeB {
	if x != nil {
		return x.Gnb
	}
	return nil
}

func (x *SiteInfo) GetPlmn() *PlmnId {
	if x != nil {
		return x.Plmn
	}
	return nil
}

func (x *SiteInfo) GetUpf() *UpfInfo {
	if x != nil {
		return x.Upf
	}
	return nil
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string `protobuf:"bytes,1,opt,name=AppName,proto3" json:"AppName,omitempty"`
	Startport uint32 `protobuf:"varint,2,opt,name=Startport,proto3" json:"Startport,omitempty"`
	Endport   uint32 `protobuf:"varint,3,opt,name=Endport,proto3" json:"Endport,omitempty"`
	Protocol  uint32 `protobuf:"varint,4,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	EndPoint  string `protobuf:"bytes,5,opt,name=EndPoint,proto3" json:"EndPoint,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

func (x *AppInfo) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppInfo) GetStartport() uint32 {
	if x != nil {
		return x.Startport
	}
	return 0
}

func (x *AppInfo) GetEndport() uint32 {
	if x != nil {
		return x.Endport
	}
	return 0
}

func (x *AppInfo) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *AppInfo) GetEndPoint() string {
	if x != nil {
		return x.EndPoint
	}
	return ""
}

type NetworkSlice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Nssai       *NSSAI         `protobuf:"bytes,2,opt,name=Nssai,proto3" json:"Nssai,omitempty"`
	Qos         *QoS           `protobuf:"bytes,3,opt,name=Qos,proto3" json:"Qos,omitempty"`
	DeviceGroup []*DeviceGroup `protobuf:"bytes,4,rep,name=DeviceGroup,proto3" json:"DeviceGroup,omitempty"`
	Site        *SiteInfo      `protobuf:"bytes,5,opt,name=Site,proto3" json:"Site,omitempty"`
	DenyApps    []string       `protobuf:"bytes,6,rep,name=DenyApps,proto3" json:"DenyApps,omitempty"`
	PermitApps  []string       `protobuf:"bytes,7,rep,name=PermitApps,proto3" json:"PermitApps,omitempty"`
	AppInfo     []*AppInfo     `protobuf:"bytes,8,rep,name=AppInfo,proto3" json:"AppInfo,omitempty"`
}

func (x *NetworkSlice) Reset() {
	*x = NetworkSlice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkSlice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSlice) ProtoMessage() {}

func (x *NetworkSlice) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSlice.ProtoReflect.Descriptor instead.
func (*NetworkSlice) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkSlice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NetworkSlice) GetNssai() *NSSAI {
	if x != nil {
		return x.Nssai
	}
	return nil
}

func (x *NetworkSlice) GetQos() *QoS {
	if x != nil {
		return x.Qos
	}
	return nil
}

func (x *NetworkSlice) GetDeviceGroup() []*DeviceGroup {
	if x != nil {
		return x.DeviceGroup
	}
	return nil
}

func (x *NetworkSlice) GetSite() *SiteInfo {
	if x != nil {
		return x.Site
	}
	return nil
}

func (x *NetworkSlice) GetDenyApps() []string {
	if x != nil {
		return x.DenyApps
	}
	return nil
}

func (x *NetworkSlice) GetPermitApps() []string {
	if x != nil {
		return x.PermitApps
	}
	return nil
}

func (x *NetworkSlice) GetAppInfo() []*AppInfo {
	if x != nil {
		return x.AppInfo
	}
	return nil
}

type DeviceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string    `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	IpDomainDetails *IpDomain `protobuf:"bytes,2,opt,name=IpDomainDetails,proto3" json:"IpDomainDetails,omitempty"`
	Imsi            []string  `protobuf:"bytes,3,rep,name=Imsi,proto3" json:"Imsi,omitempty"`
}

func (x *DeviceGroup) Reset() {
	*x = DeviceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceGroup) ProtoMessage() {}

func (x *DeviceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceGroup.ProtoReflect.Descriptor instead.
func (*DeviceGroup) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{8}
}

func (x *DeviceGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceGroup) GetIpDomainDetails() *IpDomain {
	if x != nil {
		return x.IpDomainDetails
	}
	return nil
}

func (x *DeviceGroup) GetImsi() []string {
	if x != nil {
		return x.Imsi
	}
	return nil
}

type IpDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	DnnName    string `protobuf:"bytes,2,opt,name=DnnName,proto3" json:"DnnName,omitempty"`
	UePool     string `protobuf:"bytes,3,opt,name=UePool,proto3" json:"UePool,omitempty"`
	DnsPrimary string `protobuf:"bytes,4,opt,name=DnsPrimary,proto3" json:"DnsPrimary,omitempty"`
	Mtu        int32  `protobuf:"varint,5,opt,name=Mtu,proto3" json:"Mtu,omitempty"`
}

func (x *IpDomain) Reset() {
	*x = IpDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpDomain) ProtoMessage() {}

func (x *IpDomain) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpDomain.ProtoReflect.Descriptor instead.
func (*IpDomain) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{9}
}

func (x *IpDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IpDomain) GetDnnName() string {
	if x != nil {
		return x.DnnName
	}
	return ""
}

func (x *IpDomain) GetUePool() string {
	if x != nil {
		return x.UePool
	}
	return ""
}

func (x *IpDomain) GetDnsPrimary() string {
	if x != nil {
		return x.DnsPrimary
	}
	return ""
}

func (x *IpDomain) GetMtu() int32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

type NetworkSliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestartCounter uint32 `protobuf:"varint,1,opt,name=RestartCounter,proto3" json:"RestartCounter,omitempty"`
	ClientId       string `protobuf:"bytes,2,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	ImsiRequested  bool   `protobuf:"varint,3,opt,name=ImsiRequested,proto3" json:"ImsiRequested,omitempty"`
}

func (x *NetworkSliceRequest) Reset() {
	*x = NetworkSliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkSliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSliceRequest) ProtoMessage() {}

func (x *NetworkSliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSliceRequest.ProtoReflect.Descriptor instead.
func (*NetworkSliceRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{10}
}

func (x *NetworkSliceRequest) GetRestartCounter() uint32 {
	if x != nil {
		return x.RestartCounter
	}
	return 0
}

func (x *NetworkSliceRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *NetworkSliceRequest) GetImsiRequested() bool {
	if x != nil {
		return x.ImsiRequested
	}
	return false
}

type NetworkSliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestartCounter uint32          `protobuf:"varint,1,opt,name=RestartCounter,proto3" json:"RestartCounter,omitempty"`
	NetworkSlice   []*NetworkSlice `protobuf:"bytes,2,rep,name=NetworkSlice,proto3" json:"NetworkSlice,omitempty"`
	ConfigUpdated  uint32          `protobuf:"varint,3,opt,name=ConfigUpdated,proto3" json:"ConfigUpdated,omitempty"`
}

func (x *NetworkSliceResponse) Reset() {
	*x = NetworkSliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkSliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkSliceResponse) ProtoMessage() {}

func (x *NetworkSliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkSliceResponse.ProtoReflect.Descriptor instead.
func (*NetworkSliceResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{11}
}

func (x *NetworkSliceResponse) GetRestartCounter() uint32 {
	if x != nil {
		return x.RestartCounter
	}
	return 0
}

func (x *NetworkSliceResponse) GetNetworkSlice() []*NetworkSlice {
	if x != nil {
		return x.NetworkSlice
	}
	return nil
}

func (x *NetworkSliceResponse) GetConfigUpdated() uint32 {
	if x != nil {
		return x.ConfigUpdated
	}
	return 0
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2c, 0x0a, 0x06,
	0x50, 0x6c, 0x6d, 0x6e, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x63, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x63, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x6e, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x6e, 0x63, 0x22, 0x29, 0x0a, 0x05, 0x4e, 0x53,
	0x53, 0x41, 0x49, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x53, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x53, 0x64, 0x22, 0x5d, 0x0a, 0x03, 0x51, 0x6f, 0x53, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x22, 0x2e, 0x0a, 0x06, 0x47, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x54, 0x61, 0x63, 0x22, 0x3d, 0x0a, 0x07, 0x55, 0x70, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x55, 0x70, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x55, 0x70, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x70, 0x66,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x55, 0x70, 0x66, 0x50,
	0x6f, 0x72, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x08, 0x53, 0x69, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x69, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x69, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x03,
	0x47, 0x6e, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x64, 0x63, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x52,
	0x03, 0x47, 0x6e, 0x62, 0x12, 0x28, 0x0a, 0x04, 0x50, 0x6c, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x6c, 0x6d, 0x6e, 0x49, 0x64, 0x52, 0x04, 0x50, 0x6c, 0x6d, 0x6e, 0x12, 0x27,
	0x0a, 0x03, 0x55, 0x70, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x64,
	0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x70, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x55, 0x70, 0x66, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0xc8, 0x02,
	0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x4e, 0x73, 0x73, 0x61, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4e, 0x53, 0x53, 0x41, 0x49, 0x52, 0x05, 0x4e, 0x73, 0x73, 0x61, 0x69, 0x12, 0x23, 0x0a,
	0x03, 0x51, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x64, 0x63,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x51, 0x6f, 0x53, 0x52, 0x03, 0x51,
	0x6f, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x2a, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x69, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x53, 0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44,
	0x65, 0x6e, 0x79, 0x41, 0x70, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x44,
	0x65, 0x6e, 0x79, 0x41, 0x70, 0x70, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x74, 0x41, 0x70, 0x70, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x74, 0x41, 0x70, 0x70, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x77, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x49,
	0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x49, 0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x0f, 0x49, 0x70,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x49, 0x6d, 0x73, 0x69, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6d, 0x73,
	0x69, 0x22, 0x82, 0x01, 0x0a, 0x08, 0x49, 0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x6e, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x6e, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x6e, 0x73, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x74, 0x75, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x4d, 0x74, 0x75, 0x22, 0x7f, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6d, 0x73, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x49, 0x6d, 0x73, 0x69, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x0c, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x2a, 0x22,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45,
	0x10, 0x01, 0x32, 0x6b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x64, 0x63, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x64, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_config_proto_goTypes = []interface{}{
	(Status)(0),                  // 0: sdcoreConfig.Status
	(*PlmnId)(nil),               // 1: sdcoreConfig.PlmnId
	(*NSSAI)(nil),                // 2: sdcoreConfig.NSSAI
	(*QoS)(nil),                  // 3: sdcoreConfig.QoS
	(*GNodeB)(nil),               // 4: sdcoreConfig.GNodeB
	(*UpfInfo)(nil),              // 5: sdcoreConfig.UpfInfo
	(*SiteInfo)(nil),             // 6: sdcoreConfig.SiteInfo
	(*AppInfo)(nil),              // 7: sdcoreConfig.AppInfo
	(*NetworkSlice)(nil),         // 8: sdcoreConfig.NetworkSlice
	(*DeviceGroup)(nil),          // 9: sdcoreConfig.DeviceGroup
	(*IpDomain)(nil),             // 10: sdcoreConfig.IpDomain
	(*NetworkSliceRequest)(nil),  // 11: sdcoreConfig.NetworkSliceRequest
	(*NetworkSliceResponse)(nil), // 12: sdcoreConfig.NetworkSliceResponse
}
var file_config_proto_depIdxs = []int32{
	4,  // 0: sdcoreConfig.SiteInfo.Gnb:type_name -> sdcoreConfig.GNodeB
	1,  // 1: sdcoreConfig.SiteInfo.Plmn:type_name -> sdcoreConfig.PlmnId
	5,  // 2: sdcoreConfig.SiteInfo.Upf:type_name -> sdcoreConfig.UpfInfo
	2,  // 3: sdcoreConfig.NetworkSlice.Nssai:type_name -> sdcoreConfig.NSSAI
	3,  // 4: sdcoreConfig.NetworkSlice.Qos:type_name -> sdcoreConfig.QoS
	9,  // 5: sdcoreConfig.NetworkSlice.DeviceGroup:type_name -> sdcoreConfig.DeviceGroup
	6,  // 6: sdcoreConfig.NetworkSlice.Site:type_name -> sdcoreConfig.SiteInfo
	7,  // 7: sdcoreConfig.NetworkSlice.AppInfo:type_name -> sdcoreConfig.AppInfo
	10, // 8: sdcoreConfig.DeviceGroup.IpDomainDetails:type_name -> sdcoreConfig.IpDomain
	8,  // 9: sdcoreConfig.NetworkSliceResponse.NetworkSlice:type_name -> sdcoreConfig.NetworkSlice
	11, // 10: sdcoreConfig.ConfigService.GetNetworkSlice:input_type -> sdcoreConfig.NetworkSliceRequest
	12, // 11: sdcoreConfig.ConfigService.GetNetworkSlice:output_type -> sdcoreConfig.NetworkSliceResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlmnId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSSAI); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QoS); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GNodeB); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpfInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkSlice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpDomain); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkSliceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkSliceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
