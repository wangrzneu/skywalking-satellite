//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: language-agent/Meter.proto

package v3

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	v3 "skywalking/network/common/v3"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Label of the meter
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// The histogram element definition. It includes the bucket lower boundary and the count in the bucket.
type MeterBucketValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The value represents the min value of the bucket,
	// the  upper boundary is determined by next MeterBucketValue$bucket,
	// if it doesn't exist, the upper boundary is positive infinity.
	// Also, could use Int32.MIN_VALUE to represent negative infinity.
	Bucket float64 `protobuf:"fixed64,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Count  int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *MeterBucketValue) Reset() {
	*x = MeterBucketValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterBucketValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterBucketValue) ProtoMessage() {}

func (x *MeterBucketValue) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterBucketValue.ProtoReflect.Descriptor instead.
func (*MeterBucketValue) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{1}
}

func (x *MeterBucketValue) GetBucket() float64 {
	if x != nil {
		return x.Bucket
	}
	return 0
}

func (x *MeterBucketValue) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Meter single value
type MeterSingleValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Meter name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Labels
	Labels []*Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// Single value
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MeterSingleValue) Reset() {
	*x = MeterSingleValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterSingleValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterSingleValue) ProtoMessage() {}

func (x *MeterSingleValue) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterSingleValue.ProtoReflect.Descriptor instead.
func (*MeterSingleValue) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{2}
}

func (x *MeterSingleValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeterSingleValue) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MeterSingleValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Histogram
type MeterHistogram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Meter name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Labels
	Labels []*Label `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	// Customize the buckets
	Values []*MeterBucketValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MeterHistogram) Reset() {
	*x = MeterHistogram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterHistogram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterHistogram) ProtoMessage() {}

func (x *MeterHistogram) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterHistogram.ProtoReflect.Descriptor instead.
func (*MeterHistogram) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{3}
}

func (x *MeterHistogram) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeterHistogram) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *MeterHistogram) GetValues() []*MeterBucketValue {
	if x != nil {
		return x.Values
	}
	return nil
}

// Single meter data, if the same metrics have a different label, they will separate.
type MeterData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Meter data could be a single value or histogram.
	//
	// Types that are assignable to Metric:
	//	*MeterData_SingleValue
	//	*MeterData_Histogram
	Metric isMeterData_Metric `protobuf_oneof:"metric"`
	// Service name, be set value in the first element in the stream-call.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	// Service instance name, be set value in the first element in the stream-call.
	ServiceInstance string `protobuf:"bytes,4,opt,name=serviceInstance,proto3" json:"serviceInstance,omitempty"`
	// Meter data report time, be set value in the first element in the stream-call.
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MeterData) Reset() {
	*x = MeterData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterData) ProtoMessage() {}

func (x *MeterData) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterData.ProtoReflect.Descriptor instead.
func (*MeterData) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{4}
}

func (m *MeterData) GetMetric() isMeterData_Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func (x *MeterData) GetSingleValue() *MeterSingleValue {
	if x, ok := x.GetMetric().(*MeterData_SingleValue); ok {
		return x.SingleValue
	}
	return nil
}

func (x *MeterData) GetHistogram() *MeterHistogram {
	if x, ok := x.GetMetric().(*MeterData_Histogram); ok {
		return x.Histogram
	}
	return nil
}

func (x *MeterData) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *MeterData) GetServiceInstance() string {
	if x != nil {
		return x.ServiceInstance
	}
	return ""
}

func (x *MeterData) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type isMeterData_Metric interface {
	isMeterData_Metric()
}

type MeterData_SingleValue struct {
	SingleValue *MeterSingleValue `protobuf:"bytes,1,opt,name=singleValue,proto3,oneof"`
}

type MeterData_Histogram struct {
	Histogram *MeterHistogram `protobuf:"bytes,2,opt,name=histogram,proto3,oneof"`
}

func (*MeterData_SingleValue) isMeterData_Metric() {}

func (*MeterData_Histogram) isMeterData_Metric() {}

type MeterDataCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MeterData []*MeterData `protobuf:"bytes,1,rep,name=meterData,proto3" json:"meterData,omitempty"`
}

func (x *MeterDataCollection) Reset() {
	*x = MeterDataCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_language_agent_Meter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeterDataCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeterDataCollection) ProtoMessage() {}

func (x *MeterDataCollection) ProtoReflect() protoreflect.Message {
	mi := &file_language_agent_Meter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeterDataCollection.ProtoReflect.Descriptor instead.
func (*MeterDataCollection) Descriptor() ([]byte, []int) {
	return file_language_agent_Meter_proto_rawDescGZIP(), []int{5}
}

func (x *MeterDataCollection) GetMeterData() []*MeterData {
	if x != nil {
		return x.MeterData
	}
	return nil
}

var File_language_agent_Meter_proto protoreflect.FileDescriptor

var file_language_agent_Meter_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6b,
	0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x6a, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x37, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0xfb, 0x01, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a,
	0x0b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x4d, 0x0a,
	0x13, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x32, 0x56, 0x0a, 0x12,
	0x4d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x18, 0x2e,
	0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65,
	0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x17, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x5d, 0x0a, 0x33, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x6d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x50, 0x01, 0x5a, 0x24, 0x73,
	0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_language_agent_Meter_proto_rawDescOnce sync.Once
	file_language_agent_Meter_proto_rawDescData = file_language_agent_Meter_proto_rawDesc
)

func file_language_agent_Meter_proto_rawDescGZIP() []byte {
	file_language_agent_Meter_proto_rawDescOnce.Do(func() {
		file_language_agent_Meter_proto_rawDescData = protoimpl.X.CompressGZIP(file_language_agent_Meter_proto_rawDescData)
	})
	return file_language_agent_Meter_proto_rawDescData
}

var file_language_agent_Meter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_language_agent_Meter_proto_goTypes = []interface{}{
	(*Label)(nil),               // 0: skywalking.v3.Label
	(*MeterBucketValue)(nil),    // 1: skywalking.v3.MeterBucketValue
	(*MeterSingleValue)(nil),    // 2: skywalking.v3.MeterSingleValue
	(*MeterHistogram)(nil),      // 3: skywalking.v3.MeterHistogram
	(*MeterData)(nil),           // 4: skywalking.v3.MeterData
	(*MeterDataCollection)(nil), // 5: skywalking.v3.MeterDataCollection
	(*v3.Commands)(nil),         // 6: skywalking.v3.Commands
}
var file_language_agent_Meter_proto_depIdxs = []int32{
	0, // 0: skywalking.v3.MeterSingleValue.labels:type_name -> skywalking.v3.Label
	0, // 1: skywalking.v3.MeterHistogram.labels:type_name -> skywalking.v3.Label
	1, // 2: skywalking.v3.MeterHistogram.values:type_name -> skywalking.v3.MeterBucketValue
	2, // 3: skywalking.v3.MeterData.singleValue:type_name -> skywalking.v3.MeterSingleValue
	3, // 4: skywalking.v3.MeterData.histogram:type_name -> skywalking.v3.MeterHistogram
	4, // 5: skywalking.v3.MeterDataCollection.meterData:type_name -> skywalking.v3.MeterData
	4, // 6: skywalking.v3.MeterReportService.collect:input_type -> skywalking.v3.MeterData
	6, // 7: skywalking.v3.MeterReportService.collect:output_type -> skywalking.v3.Commands
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_language_agent_Meter_proto_init() }
func file_language_agent_Meter_proto_init() {
	if File_language_agent_Meter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_language_agent_Meter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_language_agent_Meter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterBucketValue); i {
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
		file_language_agent_Meter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterSingleValue); i {
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
		file_language_agent_Meter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterHistogram); i {
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
		file_language_agent_Meter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterData); i {
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
		file_language_agent_Meter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeterDataCollection); i {
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
	file_language_agent_Meter_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*MeterData_SingleValue)(nil),
		(*MeterData_Histogram)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_language_agent_Meter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_language_agent_Meter_proto_goTypes,
		DependencyIndexes: file_language_agent_Meter_proto_depIdxs,
		MessageInfos:      file_language_agent_Meter_proto_msgTypes,
	}.Build()
	File_language_agent_Meter_proto = out.File
	file_language_agent_Meter_proto_rawDesc = nil
	file_language_agent_Meter_proto_goTypes = nil
	file_language_agent_Meter_proto_depIdxs = nil
}