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
// distributed //under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/broker/broker.proto

/*
Package broker is a generated protocol buffer package.

It is generated from these files:
	server/broker/broker.proto

It has these top-level messages:
	Participant
	Version
	Pact
	PactVersion
	Tag
	PublishPactRequest
	PublishPactResponse
	GetAllProviderPactsRequest
	ConsumerInfo
	Links
	GetAllProviderPactsResponse
	GetProviderConsumerVersionPactRequest
	GetProviderConsumerVersionPactResponse
	Verification
	VerificationSummary
	VerificationDetail
	VerificationDetails
	VerificationResult
	PublishVerificationRequest
	PublishVerificationResponse
	RetrieveVerificationRequest
	RetrieveVerificationResponse
	BaseBrokerRequest
	BrokerAPIInfoEntry
	BrokerHomeResponse
*/
package brokerpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import services "github.com/apache/incubator-servicecomb-service-center/server/core/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Participant struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AppId       string `protobuf:"bytes,2,opt,name=appId" json:"appId,omitempty"`
	ServiceName string `protobuf:"bytes,3,opt,name=serviceName" json:"serviceName,omitempty"`
}

func (m *Participant) Reset()                    { *m = Participant{} }
func (m *Participant) String() string            { return proto.CompactTextString(m) }
func (*Participant) ProtoMessage()               {}
func (*Participant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Participant) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Participant) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *Participant) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type Version struct {
	Id            int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Number        string `protobuf:"bytes,2,opt,name=number" json:"number,omitempty"`
	ParticipantId int32  `protobuf:"varint,3,opt,name=participantId" json:"participantId,omitempty"`
	Order         int32  `protobuf:"varint,4,opt,name=order" json:"order,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Version) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Version) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Version) GetParticipantId() int32 {
	if m != nil {
		return m.ParticipantId
	}
	return 0
}

func (m *Version) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

type Pact struct {
	Id                    int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ConsumerParticipantId int32  `protobuf:"varint,2,opt,name=consumerParticipantId" json:"consumerParticipantId,omitempty"`
	ProviderParticipantId int32  `protobuf:"varint,3,opt,name=providerParticipantId" json:"providerParticipantId,omitempty"`
	Sha                   []byte `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty"`
	Content               []byte `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Pact) Reset()                    { *m = Pact{} }
func (m *Pact) String() string            { return proto.CompactTextString(m) }
func (*Pact) ProtoMessage()               {}
func (*Pact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Pact) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pact) GetConsumerParticipantId() int32 {
	if m != nil {
		return m.ConsumerParticipantId
	}
	return 0
}

func (m *Pact) GetProviderParticipantId() int32 {
	if m != nil {
		return m.ProviderParticipantId
	}
	return 0
}

func (m *Pact) GetSha() []byte {
	if m != nil {
		return m.Sha
	}
	return nil
}

func (m *Pact) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PactVersion struct {
	Id                    int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	VersionId             int32 `protobuf:"varint,2,opt,name=versionId" json:"versionId,omitempty"`
	PactId                int32 `protobuf:"varint,3,opt,name=pactId" json:"pactId,omitempty"`
	ProviderParticipantId int32 `protobuf:"varint,4,opt,name=providerParticipantId" json:"providerParticipantId,omitempty"`
}

func (m *PactVersion) Reset()                    { *m = PactVersion{} }
func (m *PactVersion) String() string            { return proto.CompactTextString(m) }
func (*PactVersion) ProtoMessage()               {}
func (*PactVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PactVersion) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PactVersion) GetVersionId() int32 {
	if m != nil {
		return m.VersionId
	}
	return 0
}

func (m *PactVersion) GetPactId() int32 {
	if m != nil {
		return m.PactId
	}
	return 0
}

func (m *PactVersion) GetProviderParticipantId() int32 {
	if m != nil {
		return m.ProviderParticipantId
	}
	return 0
}

type Tag struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	VersionId int32  `protobuf:"varint,2,opt,name=versionId" json:"versionId,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetVersionId() int32 {
	if m != nil {
		return m.VersionId
	}
	return 0
}

type PublishPactRequest struct {
	ProviderId string `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	ConsumerId string `protobuf:"bytes,2,opt,name=consumerId" json:"consumerId,omitempty"`
	Version    string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Pact       []byte `protobuf:"bytes,4,opt,name=pact,proto3" json:"pact,omitempty"`
}

func (m *PublishPactRequest) Reset()                    { *m = PublishPactRequest{} }
func (m *PublishPactRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishPactRequest) ProtoMessage()               {}
func (*PublishPactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PublishPactRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *PublishPactRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *PublishPactRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PublishPactRequest) GetPact() []byte {
	if m != nil {
		return m.Pact
	}
	return nil
}

type PublishPactResponse struct {
	Response *services.Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *PublishPactResponse) Reset()                    { *m = PublishPactResponse{} }
func (m *PublishPactResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishPactResponse) ProtoMessage()               {}
func (*PublishPactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PublishPactResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetAllProviderPactsRequest struct {
	ProviderId string             `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	BaseUrl    *BaseBrokerRequest `protobuf:"bytes,2,opt,name=baseUrl" json:"baseUrl,omitempty"`
}

func (m *GetAllProviderPactsRequest) Reset()                    { *m = GetAllProviderPactsRequest{} }
func (m *GetAllProviderPactsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllProviderPactsRequest) ProtoMessage()               {}
func (*GetAllProviderPactsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetAllProviderPactsRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *GetAllProviderPactsRequest) GetBaseUrl() *BaseBrokerRequest {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

type ConsumerInfo struct {
	Href string `protobuf:"bytes,1,opt,name=href" json:"href,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ConsumerInfo) Reset()                    { *m = ConsumerInfo{} }
func (m *ConsumerInfo) String() string            { return proto.CompactTextString(m) }
func (*ConsumerInfo) ProtoMessage()               {}
func (*ConsumerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ConsumerInfo) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *ConsumerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Links struct {
	Pacts []*ConsumerInfo `protobuf:"bytes,1,rep,name=pacts" json:"pacts,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Links) GetPacts() []*ConsumerInfo {
	if m != nil {
		return m.Pacts
	}
	return nil
}

type GetAllProviderPactsResponse struct {
	Response *services.Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XLinks   *Links             `protobuf:"bytes,2,opt,name=_links,json=Links" json:"_links,omitempty"`
}

func (m *GetAllProviderPactsResponse) Reset()                    { *m = GetAllProviderPactsResponse{} }
func (m *GetAllProviderPactsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllProviderPactsResponse) ProtoMessage()               {}
func (*GetAllProviderPactsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GetAllProviderPactsResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetAllProviderPactsResponse) GetXLinks() *Links {
	if m != nil {
		return m.XLinks
	}
	return nil
}

type GetProviderConsumerVersionPactRequest struct {
	ProviderId string             `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	ConsumerId string             `protobuf:"bytes,2,opt,name=consumerId" json:"consumerId,omitempty"`
	Version    string             `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	BaseUrl    *BaseBrokerRequest `protobuf:"bytes,4,opt,name=baseUrl" json:"baseUrl,omitempty"`
}

func (m *GetProviderConsumerVersionPactRequest) Reset()         { *m = GetProviderConsumerVersionPactRequest{} }
func (m *GetProviderConsumerVersionPactRequest) String() string { return proto.CompactTextString(m) }
func (*GetProviderConsumerVersionPactRequest) ProtoMessage()    {}
func (*GetProviderConsumerVersionPactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11}
}

func (m *GetProviderConsumerVersionPactRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *GetProviderConsumerVersionPactRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *GetProviderConsumerVersionPactRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetProviderConsumerVersionPactRequest) GetBaseUrl() *BaseBrokerRequest {
	if m != nil {
		return m.BaseUrl
	}
	return nil
}

type GetProviderConsumerVersionPactResponse struct {
	Response *services.Response `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Pact     []byte             `protobuf:"bytes,2,opt,name=pact,proto3" json:"pact,omitempty"`
}

func (m *GetProviderConsumerVersionPactResponse) Reset() {
	*m = GetProviderConsumerVersionPactResponse{}
}
func (m *GetProviderConsumerVersionPactResponse) String() string { return proto.CompactTextString(m) }
func (*GetProviderConsumerVersionPactResponse) ProtoMessage()    {}
func (*GetProviderConsumerVersionPactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12}
}

func (m *GetProviderConsumerVersionPactResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *GetProviderConsumerVersionPactResponse) GetPact() []byte {
	if m != nil {
		return m.Pact
	}
	return nil
}

type Verification struct {
	Id               int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Number           int32  `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	PactVersionId    int32  `protobuf:"varint,3,opt,name=pactVersionId" json:"pactVersionId,omitempty"`
	Success          bool   `protobuf:"varint,4,opt,name=success" json:"success,omitempty"`
	ProviderVersion  string `protobuf:"bytes,5,opt,name=providerVersion" json:"providerVersion,omitempty"`
	BuildUrl         string `protobuf:"bytes,6,opt,name=buildUrl" json:"buildUrl,omitempty"`
	VerificationDate string `protobuf:"bytes,7,opt,name=verificationDate" json:"verificationDate,omitempty"`
}

func (m *Verification) Reset()                    { *m = Verification{} }
func (m *Verification) String() string            { return proto.CompactTextString(m) }
func (*Verification) ProtoMessage()               {}
func (*Verification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Verification) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Verification) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Verification) GetPactVersionId() int32 {
	if m != nil {
		return m.PactVersionId
	}
	return 0
}

func (m *Verification) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Verification) GetProviderVersion() string {
	if m != nil {
		return m.ProviderVersion
	}
	return ""
}

func (m *Verification) GetBuildUrl() string {
	if m != nil {
		return m.BuildUrl
	}
	return ""
}

func (m *Verification) GetVerificationDate() string {
	if m != nil {
		return m.VerificationDate
	}
	return ""
}

type VerificationSummary struct {
	Successful []string `protobuf:"bytes,1,rep,name=successful" json:"successful,omitempty"`
	Failed     []string `protobuf:"bytes,2,rep,name=failed" json:"failed,omitempty"`
	Unknown    []string `protobuf:"bytes,3,rep,name=unknown" json:"unknown,omitempty"`
}

func (m *VerificationSummary) Reset()                    { *m = VerificationSummary{} }
func (m *VerificationSummary) String() string            { return proto.CompactTextString(m) }
func (*VerificationSummary) ProtoMessage()               {}
func (*VerificationSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *VerificationSummary) GetSuccessful() []string {
	if m != nil {
		return m.Successful
	}
	return nil
}

func (m *VerificationSummary) GetFailed() []string {
	if m != nil {
		return m.Failed
	}
	return nil
}

func (m *VerificationSummary) GetUnknown() []string {
	if m != nil {
		return m.Unknown
	}
	return nil
}

type VerificationDetail struct {
	ProviderName               string `protobuf:"bytes,1,opt,name=providerName" json:"providerName,omitempty"`
	ProviderApplicationVersion string `protobuf:"bytes,2,opt,name=providerApplicationVersion" json:"providerApplicationVersion,omitempty"`
	Success                    bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	VerificationDate           string `protobuf:"bytes,4,opt,name=verificationDate" json:"verificationDate,omitempty"`
}

func (m *VerificationDetail) Reset()                    { *m = VerificationDetail{} }
func (m *VerificationDetail) String() string            { return proto.CompactTextString(m) }
func (*VerificationDetail) ProtoMessage()               {}
func (*VerificationDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *VerificationDetail) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *VerificationDetail) GetProviderApplicationVersion() string {
	if m != nil {
		return m.ProviderApplicationVersion
	}
	return ""
}

func (m *VerificationDetail) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *VerificationDetail) GetVerificationDate() string {
	if m != nil {
		return m.VerificationDate
	}
	return ""
}

type VerificationDetails struct {
	VerificationResults []*VerificationDetail `protobuf:"bytes,1,rep,name=verificationResults" json:"verificationResults,omitempty"`
}

func (m *VerificationDetails) Reset()                    { *m = VerificationDetails{} }
func (m *VerificationDetails) String() string            { return proto.CompactTextString(m) }
func (*VerificationDetails) ProtoMessage()               {}
func (*VerificationDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *VerificationDetails) GetVerificationResults() []*VerificationDetail {
	if m != nil {
		return m.VerificationResults
	}
	return nil
}

type VerificationResult struct {
	Success         bool                 `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	ProviderSummary *VerificationSummary `protobuf:"bytes,2,opt,name=providerSummary" json:"providerSummary,omitempty"`
	XEmbedded       *VerificationDetails `protobuf:"bytes,3,opt,name=_embedded,json=Embedded" json:"_embedded,omitempty"`
}

func (m *VerificationResult) Reset()                    { *m = VerificationResult{} }
func (m *VerificationResult) String() string            { return proto.CompactTextString(m) }
func (*VerificationResult) ProtoMessage()               {}
func (*VerificationResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *VerificationResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *VerificationResult) GetProviderSummary() *VerificationSummary {
	if m != nil {
		return m.ProviderSummary
	}
	return nil
}

func (m *VerificationResult) GetXEmbedded() *VerificationDetails {
	if m != nil {
		return m.XEmbedded
	}
	return nil
}

type PublishVerificationRequest struct {
	ProviderId                 string `protobuf:"bytes,1,opt,name=providerId" json:"providerId,omitempty"`
	ConsumerId                 string `protobuf:"bytes,2,opt,name=consumerId" json:"consumerId,omitempty"`
	PactId                     int32  `protobuf:"varint,3,opt,name=pactId" json:"pactId,omitempty"`
	Success                    bool   `protobuf:"varint,4,opt,name=success" json:"success,omitempty"`
	ProviderApplicationVersion string `protobuf:"bytes,5,opt,name=providerApplicationVersion" json:"providerApplicationVersion,omitempty"`
}

func (m *PublishVerificationRequest) Reset()                    { *m = PublishVerificationRequest{} }
func (m *PublishVerificationRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishVerificationRequest) ProtoMessage()               {}
func (*PublishVerificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *PublishVerificationRequest) GetProviderId() string {
	if m != nil {
		return m.ProviderId
	}
	return ""
}

func (m *PublishVerificationRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *PublishVerificationRequest) GetPactId() int32 {
	if m != nil {
		return m.PactId
	}
	return 0
}

func (m *PublishVerificationRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PublishVerificationRequest) GetProviderApplicationVersion() string {
	if m != nil {
		return m.ProviderApplicationVersion
	}
	return ""
}

type PublishVerificationResponse struct {
	Response     *services.Response  `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Confirmation *VerificationDetail `protobuf:"bytes,2,opt,name=confirmation" json:"confirmation,omitempty"`
}

func (m *PublishVerificationResponse) Reset()                    { *m = PublishVerificationResponse{} }
func (m *PublishVerificationResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishVerificationResponse) ProtoMessage()               {}
func (*PublishVerificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PublishVerificationResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *PublishVerificationResponse) GetConfirmation() *VerificationDetail {
	if m != nil {
		return m.Confirmation
	}
	return nil
}

type RetrieveVerificationRequest struct {
	ConsumerId      string `protobuf:"bytes,1,opt,name=consumerId" json:"consumerId,omitempty"`
	ConsumerVersion string `protobuf:"bytes,2,opt,name=consumerVersion" json:"consumerVersion,omitempty"`
}

func (m *RetrieveVerificationRequest) Reset()                    { *m = RetrieveVerificationRequest{} }
func (m *RetrieveVerificationRequest) String() string            { return proto.CompactTextString(m) }
func (*RetrieveVerificationRequest) ProtoMessage()               {}
func (*RetrieveVerificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RetrieveVerificationRequest) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *RetrieveVerificationRequest) GetConsumerVersion() string {
	if m != nil {
		return m.ConsumerVersion
	}
	return ""
}

type RetrieveVerificationResponse struct {
	Response *services.Response  `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Result   *VerificationResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *RetrieveVerificationResponse) Reset()                    { *m = RetrieveVerificationResponse{} }
func (m *RetrieveVerificationResponse) String() string            { return proto.CompactTextString(m) }
func (*RetrieveVerificationResponse) ProtoMessage()               {}
func (*RetrieveVerificationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *RetrieveVerificationResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *RetrieveVerificationResponse) GetResult() *VerificationResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type BaseBrokerRequest struct {
	HostAddress string `protobuf:"bytes,1,opt,name=hostAddress" json:"hostAddress,omitempty"`
	Scheme      string `protobuf:"bytes,2,opt,name=scheme" json:"scheme,omitempty"`
}

func (m *BaseBrokerRequest) Reset()                    { *m = BaseBrokerRequest{} }
func (m *BaseBrokerRequest) String() string            { return proto.CompactTextString(m) }
func (*BaseBrokerRequest) ProtoMessage()               {}
func (*BaseBrokerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *BaseBrokerRequest) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

func (m *BaseBrokerRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

type BrokerAPIInfoEntry struct {
	Href      string `protobuf:"bytes,1,opt,name=href" json:"href,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Templated bool   `protobuf:"varint,4,opt,name=templated" json:"templated,omitempty"`
}

func (m *BrokerAPIInfoEntry) Reset()                    { *m = BrokerAPIInfoEntry{} }
func (m *BrokerAPIInfoEntry) String() string            { return proto.CompactTextString(m) }
func (*BrokerAPIInfoEntry) ProtoMessage()               {}
func (*BrokerAPIInfoEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *BrokerAPIInfoEntry) GetHref() string {
	if m != nil {
		return m.Href
	}
	return ""
}

func (m *BrokerAPIInfoEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BrokerAPIInfoEntry) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BrokerAPIInfoEntry) GetTemplated() bool {
	if m != nil {
		return m.Templated
	}
	return false
}

type BrokerHomeResponse struct {
	Response *services.Response             `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XLinks   map[string]*BrokerAPIInfoEntry `protobuf:"bytes,2,rep,name=_links,json=Links" json:"_links,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Curies   []*BrokerAPIInfoEntry          `protobuf:"bytes,3,rep,name=curies" json:"curies,omitempty"`
}

func (m *BrokerHomeResponse) Reset()                    { *m = BrokerHomeResponse{} }
func (m *BrokerHomeResponse) String() string            { return proto.CompactTextString(m) }
func (*BrokerHomeResponse) ProtoMessage()               {}
func (*BrokerHomeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *BrokerHomeResponse) GetResponse() *services.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *BrokerHomeResponse) GetXLinks() map[string]*BrokerAPIInfoEntry {
	if m != nil {
		return m.XLinks
	}
	return nil
}

func (m *BrokerHomeResponse) GetCuries() []*BrokerAPIInfoEntry {
	if m != nil {
		return m.Curies
	}
	return nil
}

func init() {
	proto.RegisterType((*Participant)(nil), "Participant")
	proto.RegisterType((*Version)(nil), "Version")
	proto.RegisterType((*Pact)(nil), "Pact")
	proto.RegisterType((*PactVersion)(nil), "PactVersion")
	proto.RegisterType((*Tag)(nil), "Tag")
	proto.RegisterType((*PublishPactRequest)(nil), "PublishPactRequest")
	proto.RegisterType((*PublishPactResponse)(nil), "PublishPactResponse")
	proto.RegisterType((*GetAllProviderPactsRequest)(nil), "GetAllProviderPactsRequest")
	proto.RegisterType((*ConsumerInfo)(nil), "ConsumerInfo")
	proto.RegisterType((*Links)(nil), "Links")
	proto.RegisterType((*GetAllProviderPactsResponse)(nil), "GetAllProviderPactsResponse")
	proto.RegisterType((*GetProviderConsumerVersionPactRequest)(nil), "GetProviderConsumerVersionPactRequest")
	proto.RegisterType((*GetProviderConsumerVersionPactResponse)(nil), "GetProviderConsumerVersionPactResponse")
	proto.RegisterType((*Verification)(nil), "Verification")
	proto.RegisterType((*VerificationSummary)(nil), "VerificationSummary")
	proto.RegisterType((*VerificationDetail)(nil), "VerificationDetail")
	proto.RegisterType((*VerificationDetails)(nil), "VerificationDetails")
	proto.RegisterType((*VerificationResult)(nil), "VerificationResult")
	proto.RegisterType((*PublishVerificationRequest)(nil), "PublishVerificationRequest")
	proto.RegisterType((*PublishVerificationResponse)(nil), "PublishVerificationResponse")
	proto.RegisterType((*RetrieveVerificationRequest)(nil), "RetrieveVerificationRequest")
	proto.RegisterType((*RetrieveVerificationResponse)(nil), "RetrieveVerificationResponse")
	proto.RegisterType((*BaseBrokerRequest)(nil), "BaseBrokerRequest")
	proto.RegisterType((*BrokerAPIInfoEntry)(nil), "BrokerAPIInfoEntry")
	proto.RegisterType((*BrokerHomeResponse)(nil), "BrokerHomeResponse")
}

func init() { proto.RegisterFile("server/broker/broker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdd, 0x8e, 0x1b, 0x35,
	0x14, 0xd6, 0x64, 0x93, 0x6c, 0x72, 0x92, 0xd2, 0xe2, 0x2d, 0x28, 0x4a, 0x4b, 0x15, 0x19, 0x8a,
	0x02, 0x54, 0x59, 0xb1, 0xfc, 0x14, 0x21, 0x54, 0x69, 0x97, 0xae, 0xca, 0x4a, 0xb4, 0x8a, 0x0c,
	0xed, 0x05, 0x42, 0xaa, 0x26, 0x33, 0x27, 0x9b, 0x61, 0xe7, 0x0f, 0xdb, 0x13, 0xb4, 0x17, 0xdc,
	0x70, 0xc7, 0x63, 0xf0, 0x02, 0xf0, 0x06, 0x3c, 0x02, 0xcf, 0xc1, 0x15, 0xcf, 0x80, 0xec, 0xb1,
	0x27, 0x33, 0xc9, 0x64, 0xb7, 0x91, 0x10, 0x57, 0xf1, 0xf9, 0xf1, 0xf1, 0x77, 0xbe, 0x73, 0xe6,
	0xd8, 0x81, 0xa1, 0x40, 0xbe, 0x44, 0x7e, 0x38, 0xe3, 0xc9, 0x45, 0xf1, 0x33, 0x49, 0x79, 0x22,
	0x93, 0xe1, 0xc8, 0xd8, 0xbc, 0x84, 0xe3, 0xa1, 0x56, 0x1d, 0x2a, 0x4d, 0xe0, 0xa1, 0xc8, 0x3d,
	0xe8, 0x73, 0xe8, 0x4d, 0x5d, 0x2e, 0x03, 0x2f, 0x48, 0xdd, 0x58, 0x92, 0xd7, 0xa0, 0x11, 0xf8,
	0x03, 0x67, 0xe4, 0x8c, 0x5b, 0xac, 0x11, 0xf8, 0xe4, 0x36, 0xb4, 0xdc, 0x34, 0x3d, 0xf3, 0x07,
	0x8d, 0x91, 0x33, 0xee, 0xb2, 0x5c, 0x20, 0x23, 0xe8, 0x99, 0x30, 0xcf, 0xdc, 0x08, 0x07, 0x7b,
	0xda, 0x56, 0x56, 0xd1, 0x08, 0xf6, 0x5f, 0x20, 0x17, 0x41, 0x12, 0x6f, 0x84, 0x7c, 0x13, 0xda,
	0x71, 0x16, 0xcd, 0x90, 0x9b, 0x98, 0x46, 0x22, 0xef, 0xc0, 0x8d, 0x74, 0x85, 0xe4, 0xcc, 0xd7,
	0x61, 0x5b, 0xac, 0xaa, 0x54, 0x80, 0x12, 0xee, 0x23, 0x1f, 0x34, 0xb5, 0x35, 0x17, 0xe8, 0xef,
	0x0e, 0x34, 0xa7, 0xae, 0xb7, 0x89, 0xff, 0x63, 0x78, 0xc3, 0x4b, 0x62, 0x91, 0x45, 0xc8, 0xa7,
	0x95, 0xe0, 0x0d, 0xed, 0x52, 0x6f, 0x54, 0xbb, 0x52, 0x9e, 0x2c, 0x03, 0x7f, 0x7d, 0x57, 0x0e,
	0xa9, 0xde, 0x48, 0x6e, 0xc1, 0x9e, 0x58, 0xb8, 0x1a, 0x58, 0x9f, 0xa9, 0x25, 0x19, 0xc0, 0xbe,
	0x97, 0xc4, 0x12, 0x63, 0x39, 0x68, 0x69, 0xad, 0x15, 0xe9, 0xaf, 0x8e, 0xe2, 0xdd, 0x93, 0xdb,
	0x48, 0xba, 0x0b, 0xdd, 0x65, 0x6e, 0x2a, 0xb0, 0xae, 0x14, 0x8a, 0xc2, 0xd4, 0xf5, 0x56, 0x80,
	0x8c, 0xb4, 0x1d, 0x77, 0xf3, 0x0a, 0xdc, 0xf4, 0x21, 0xec, 0x7d, 0xeb, 0x9e, 0x13, 0x02, 0xcd,
	0x58, 0x55, 0xd3, 0xd1, 0x55, 0xd1, 0xeb, 0xab, 0x61, 0xd0, 0x5f, 0x1c, 0x20, 0xd3, 0x6c, 0x16,
	0x06, 0x62, 0xa1, 0x72, 0x61, 0xf8, 0x63, 0x86, 0x42, 0x92, 0x7b, 0x00, 0xf6, 0xa0, 0x33, 0xdf,
	0x84, 0x2b, 0x69, 0x94, 0xdd, 0xd2, 0x5e, 0x34, 0x56, 0x49, 0xa3, 0x58, 0x33, 0x67, 0x98, 0xce,
	0xb2, 0xa2, 0x82, 0xa8, 0x32, 0x35, 0x14, 0xeb, 0x35, 0xfd, 0x02, 0x0e, 0x2a, 0x18, 0x44, 0x9a,
	0xc4, 0x02, 0xc9, 0x7d, 0xe8, 0x70, 0xb3, 0xd6, 0x10, 0x7a, 0x47, 0xdd, 0x89, 0x35, 0xb2, 0xc2,
	0x44, 0x7f, 0x80, 0xe1, 0x13, 0x94, 0xc7, 0x61, 0x38, 0x2d, 0xa8, 0xf1, 0xa4, 0x78, 0xd5, 0x4c,
	0x1e, 0xc0, 0xfe, 0xcc, 0x15, 0xf8, 0x9c, 0x87, 0x3a, 0x8d, 0xde, 0x11, 0x99, 0x9c, 0xb8, 0x02,
	0x4f, 0xf4, 0x27, 0x68, 0x82, 0x30, 0xeb, 0x42, 0x3f, 0x85, 0xfe, 0x97, 0x36, 0xcb, 0x78, 0x9e,
	0xa8, 0x6c, 0x16, 0x1c, 0xe7, 0x96, 0x70, 0xb5, 0x2e, 0x8a, 0xd0, 0x58, 0x15, 0x81, 0x3e, 0x80,
	0xd6, 0xd7, 0x41, 0x7c, 0x21, 0xc8, 0xdb, 0xd0, 0x52, 0x29, 0x8b, 0x81, 0x33, 0xda, 0x1b, 0xf7,
	0x8e, 0x6e, 0x4c, 0xca, 0xe1, 0x58, 0x6e, 0xa3, 0x1e, 0xdc, 0xa9, 0xcd, 0x68, 0x27, 0x5e, 0xc8,
	0x5b, 0xd0, 0x7e, 0x19, 0xaa, 0x43, 0x4d, 0x62, 0xed, 0x89, 0x86, 0xc0, 0x72, 0x24, 0xf4, 0x0f,
	0x07, 0xee, 0x3f, 0x41, 0x69, 0x8f, 0xb0, 0x38, 0x4c, 0x37, 0xff, 0x3f, 0xcd, 0x50, 0x22, 0xbf,
	0x79, 0x3d, 0xf9, 0x1e, 0xbc, 0x7b, 0x1d, 0xe0, 0xdd, 0x18, 0xb2, 0xbd, 0xd8, 0x28, 0xf5, 0xe2,
	0xdf, 0x0e, 0xf4, 0x5f, 0x20, 0x0f, 0xe6, 0x81, 0xe7, 0xca, 0xeb, 0x67, 0x5f, 0xab, 0x3a, 0xfb,
	0x8a, 0x69, 0x50, 0x9e, 0x7d, 0x25, 0xa5, 0xe2, 0x42, 0x64, 0x9e, 0x87, 0x42, 0xe8, 0x8c, 0x3b,
	0xcc, 0x8a, 0x64, 0x0c, 0x37, 0x2d, 0xa7, 0xc6, 0x5d, 0x0f, 0x9c, 0x2e, 0x5b, 0x57, 0x93, 0x21,
	0x74, 0x66, 0x59, 0x10, 0xfa, 0x8a, 0xb6, 0xb6, 0x76, 0x29, 0x64, 0xf2, 0x3e, 0xdc, 0x5a, 0x96,
	0xd0, 0x3f, 0x76, 0x25, 0x0e, 0xf6, 0xb5, 0xcf, 0x86, 0x9e, 0x9e, 0xc3, 0x41, 0x39, 0xd3, 0x6f,
	0xb2, 0x28, 0x72, 0xf9, 0xa5, 0x2a, 0xa7, 0xc1, 0x34, 0xcf, 0x42, 0xdd, 0xa7, 0x5d, 0x56, 0xd2,
	0x28, 0x02, 0xe6, 0x6e, 0x10, 0xa2, 0x2a, 0xb5, 0xb2, 0x19, 0x49, 0xa5, 0x96, 0xc5, 0x17, 0x71,
	0xf2, 0x93, 0x2a, 0xb3, 0x32, 0x58, 0x91, 0xfe, 0xe9, 0x00, 0x29, 0x9f, 0xf4, 0x18, 0xa5, 0x1b,
	0x84, 0x84, 0x42, 0xdf, 0xa6, 0xf6, 0x6c, 0x35, 0xb5, 0x2a, 0x3a, 0xf2, 0x08, 0x86, 0x56, 0x3e,
	0x4e, 0xd3, 0xd0, 0x04, 0xb0, 0x04, 0xe5, 0xbd, 0x76, 0x85, 0x47, 0x99, 0xef, 0xbd, 0x2a, 0xdf,
	0x75, 0x4c, 0x35, 0xb7, 0x30, 0xf5, 0x7d, 0x95, 0xa9, 0x1c, 0xbf, 0x20, 0xa7, 0x70, 0x50, 0x76,
	0x65, 0x28, 0xb2, 0xb0, 0xf8, 0xb4, 0x0f, 0x26, 0x9b, 0x5b, 0x58, 0x9d, 0x3f, 0xfd, 0x6d, 0x8d,
	0x9e, 0x5c, 0x5f, 0x86, 0xee, 0x54, 0xa1, 0x3f, 0x5a, 0xb5, 0x8a, 0x29, 0x9a, 0xf9, 0xc4, 0x6f,
	0x4f, 0x6a, 0x0a, 0xca, 0xd6, 0x9d, 0xc9, 0x87, 0xd0, 0x7d, 0x89, 0xd1, 0x0c, 0x7d, 0x1f, 0xf3,
	0x36, 0x5d, 0xdf, 0x69, 0x12, 0x64, 0x9d, 0x53, 0xe3, 0x45, 0xff, 0x72, 0x60, 0x68, 0x66, 0x74,
	0x15, 0xea, 0x7f, 0x33, 0x22, 0xb6, 0xdd, 0x86, 0xdb, 0x3f, 0x97, 0xab, 0x1b, 0xa3, 0x75, 0x5d,
	0x63, 0xd0, 0x9f, 0xe1, 0x4e, 0x6d, 0x3e, 0xbb, 0x4d, 0x90, 0x87, 0xd0, 0xf7, 0x92, 0x78, 0x1e,
	0xf0, 0x48, 0x6f, 0x37, 0x65, 0xa8, 0x2d, 0x7d, 0xc5, 0x91, 0x9e, 0xc3, 0x1d, 0x86, 0x92, 0x07,
	0xb8, 0xc4, 0x2d, 0x7c, 0x96, 0xf8, 0x72, 0x36, 0xf8, 0x1a, 0xc3, 0x4d, 0xaf, 0x3a, 0xff, 0x0c,
	0xa9, 0xeb, 0x6a, 0xca, 0xe1, 0x6e, 0xfd, 0x41, 0xbb, 0x25, 0xfa, 0x01, 0xb4, 0xb9, 0x6e, 0xcb,
	0xda, 0x14, 0xf3, 0x8e, 0x65, 0xc6, 0x85, 0x3e, 0x85, 0xd7, 0x37, 0xc6, 0xb8, 0x7a, 0x70, 0x2e,
	0x12, 0x21, 0x8f, 0x7d, 0x9f, 0xdb, 0x96, 0xee, 0xb2, 0xb2, 0x4a, 0x35, 0x81, 0xf0, 0x16, 0x58,
	0x5c, 0x9d, 0x46, 0xa2, 0x29, 0x90, 0x3c, 0xd4, 0xf1, 0xf4, 0x4c, 0x5d, 0x93, 0xa7, 0xb1, 0xe4,
	0x97, 0xaf, 0x7a, 0xf5, 0xaa, 0xd7, 0xa6, 0x0c, 0x64, 0x68, 0x9f, 0xb8, 0xb9, 0xa0, 0x5e, 0x45,
	0x12, 0xa3, 0x34, 0x74, 0x25, 0xfa, 0xa6, 0xb5, 0x56, 0x0a, 0xfa, 0x8f, 0x63, 0x8f, 0xfc, 0x2a,
	0x89, 0x70, 0x57, 0xae, 0x3e, 0x29, 0x5d, 0xbc, 0x6a, 0x12, 0xdc, 0x9b, 0x6c, 0xc6, 0xca, 0xef,
	0x62, 0x9d, 0x89, 0xb9, 0x90, 0x15, 0xc5, 0x5e, 0xc6, 0x03, 0x14, 0x7a, 0x7c, 0x2a, 0x8a, 0x37,
	0xb3, 0x66, 0xc6, 0x65, 0xf8, 0x14, 0x60, 0x15, 0x41, 0x3d, 0x5b, 0x2f, 0xf0, 0xd2, 0x50, 0xa1,
	0x96, 0xe4, 0x3d, 0x68, 0x2d, 0xdd, 0x30, 0xc3, 0xa2, 0x5c, 0x35, 0xb1, 0x72, 0x8f, 0xcf, 0x1b,
	0x9f, 0x39, 0x27, 0x9d, 0xef, 0xda, 0xf9, 0x9f, 0x8e, 0x59, 0x5b, 0xff, 0xa7, 0xf8, 0xe8, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x8f, 0x7b, 0x26, 0x93, 0x0c, 0x00, 0x00,
}
