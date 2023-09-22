//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/events/sandbox.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SandboxCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
}

func (x *SandboxCreate) Reset() {
	*x = SandboxCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreate) ProtoMessage() {}

func (x *SandboxCreate) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreate.ProtoReflect.Descriptor instead.
func (*SandboxCreate) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{0}
}

func (x *SandboxCreate) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

type SandboxStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
}

func (x *SandboxStart) Reset() {
	*x = SandboxStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxStart) ProtoMessage() {}

func (x *SandboxStart) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxStart.ProtoReflect.Descriptor instead.
func (*SandboxStart) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{1}
}

func (x *SandboxStart) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

type SandboxExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId  string                 `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
	ExitStatus uint32                 `protobuf:"varint,2,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
}

func (x *SandboxExit) Reset() {
	*x = SandboxExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxExit) ProtoMessage() {}

func (x *SandboxExit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxExit.ProtoReflect.Descriptor instead.
func (*SandboxExit) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{2}
}

func (x *SandboxExit) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

func (x *SandboxExit) GetExitStatus() uint32 {
	if x != nil {
		return x.ExitStatus
	}
	return 0
}

func (x *SandboxExit) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

var File_github_com_containerd_containerd_api_events_sandbox_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2e, 0x0a, 0x0d, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x22,
	0x2d, 0x0a, 0x0c, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x22, 0x86,
	0x01, 0x0a, 0x0b, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x45, 0x78, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x78, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData = file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc
)

func file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData
}

var file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes = []interface{}{
	(*SandboxCreate)(nil),         // 0: containerd.events.SandboxCreate
	(*SandboxStart)(nil),          // 1: containerd.events.SandboxStart
	(*SandboxExit)(nil),           // 2: containerd.events.SandboxExit
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs = []int32{
	3, // 0: containerd.events.SandboxExit.exited_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_events_sandbox_proto_init() }
func file_github_com_containerd_containerd_api_events_sandbox_proto_init() {
	if File_github_com_containerd_containerd_api_events_sandbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxCreate); i {
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
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxStart); i {
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
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxExit); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_events_sandbox_proto = out.File
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes = nil
	file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs = nil
}
//
//Copyright The containerd Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: github.com/containerd/containerd/api/events/sandbox.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SandboxCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxID string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
}

func (x *SandboxCreate) Reset() {
	*x = SandboxCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreate) ProtoMessage() {}

func (x *SandboxCreate) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreate.ProtoReflect.Descriptor instead.
func (*SandboxCreate) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{0}
}

func (x *SandboxCreate) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

type SandboxStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxID string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
}

func (x *SandboxStart) Reset() {
	*x = SandboxStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxStart) ProtoMessage() {}

func (x *SandboxStart) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxStart.ProtoReflect.Descriptor instead.
func (*SandboxStart) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{1}
}

func (x *SandboxStart) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

type SandboxExit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxID  string                 `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
	ExitStatus uint32                 `protobuf:"varint,2,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
	ExitedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=exited_at,json=exitedAt,proto3" json:"exited_at,omitempty"`
}

func (x *SandboxExit) Reset() {
	*x = SandboxExit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxExit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxExit) ProtoMessage() {}

func (x *SandboxExit) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxExit.ProtoReflect.Descriptor instead.
func (*SandboxExit) Descriptor() ([]byte, []int) {
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP(), []int{2}
}

func (x *SandboxExit) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

func (x *SandboxExit) GetExitStatus() uint32 {
	if x != nil {
		return x.ExitStatus
	}
	return 0
}

func (x *SandboxExit) GetExitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExitedAt
	}
	return nil
}

var File_github_com_containerd_containerd_api_events_sandbox_proto protoreflect.FileDescriptor

var file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x2e, 0x0a, 0x0d, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x22,
	0x2d, 0x0a, 0x0c, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x22, 0x86,
	0x01, 0x0a, 0x0b, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x45, 0x78, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65,
	0x78, 0x69, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescOnce sync.Once
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData = file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc
)

func file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescGZIP() []byte {
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescOnce.Do(func() {
		file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData)
	})
	return file_github_com_containerd_containerd_api_events_sandbox_proto_rawDescData
}

var file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes = []interface{}{
	(*SandboxCreate)(nil),         // 0: containerd.events.SandboxCreate
	(*SandboxStart)(nil),          // 1: containerd.events.SandboxStart
	(*SandboxExit)(nil),           // 2: containerd.events.SandboxExit
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs = []int32{
	3, // 0: containerd.events.SandboxExit.exited_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_containerd_containerd_api_events_sandbox_proto_init() }
func file_github_com_containerd_containerd_api_events_sandbox_proto_init() {
	if File_github_com_containerd_containerd_api_events_sandbox_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxCreate); i {
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
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxStart); i {
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
		file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxExit); i {
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
			RawDescriptor: file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes,
		DependencyIndexes: file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs,
		MessageInfos:      file_github_com_containerd_containerd_api_events_sandbox_proto_msgTypes,
	}.Build()
	File_github_com_containerd_containerd_api_events_sandbox_proto = out.File
	file_github_com_containerd_containerd_api_events_sandbox_proto_rawDesc = nil
	file_github_com_containerd_containerd_api_events_sandbox_proto_goTypes = nil
	file_github_com_containerd_containerd_api_events_sandbox_proto_depIdxs = nil
}
