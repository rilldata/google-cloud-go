// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.3
// source: google/chat/v1/group.proto

package chatpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Google Group in Google Chat.
type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource name for a Google Group.
	//
	// Represents a
	// [group](https://cloud.google.com/identity/docs/reference/rest/v1/groups) in
	// Cloud Identity Groups API.
	//
	// Format: groups/{group}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_google_chat_v1_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_google_chat_v1_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_google_chat_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_chat_v1_group_proto protoreflect.FileDescriptor

var file_google_chat_v1_group_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x05,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xa3, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x70, 0x62, 0x3b, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x0b, 0x44,
	0x59, 0x4e, 0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x13, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x43,
	0x68, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_chat_v1_group_proto_rawDescOnce sync.Once
	file_google_chat_v1_group_proto_rawDescData = file_google_chat_v1_group_proto_rawDesc
)

func file_google_chat_v1_group_proto_rawDescGZIP() []byte {
	file_google_chat_v1_group_proto_rawDescOnce.Do(func() {
		file_google_chat_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_chat_v1_group_proto_rawDescData)
	})
	return file_google_chat_v1_group_proto_rawDescData
}

var file_google_chat_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_chat_v1_group_proto_goTypes = []any{
	(*Group)(nil), // 0: google.chat.v1.Group
}
var file_google_chat_v1_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_chat_v1_group_proto_init() }
func file_google_chat_v1_group_proto_init() {
	if File_google_chat_v1_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_chat_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_chat_v1_group_proto_goTypes,
		DependencyIndexes: file_google_chat_v1_group_proto_depIdxs,
		MessageInfos:      file_google_chat_v1_group_proto_msgTypes,
	}.Build()
	File_google_chat_v1_group_proto = out.File
	file_google_chat_v1_group_proto_rawDesc = nil
	file_google_chat_v1_group_proto_goTypes = nil
	file_google_chat_v1_group_proto_depIdxs = nil
}
