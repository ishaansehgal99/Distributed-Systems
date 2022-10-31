// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: messages.proto

package ProtoPackage

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MessageType int32

const (
	MessageType_MEMBERSHIP     MessageType = 0
	MessageType_JOIN_REQ       MessageType = 1
	MessageType_JOIN_REP       MessageType = 2
	MessageType_FILE_OP        MessageType = 3 // master request to non-master machine
	MessageType_CLIENT_REQUEST MessageType = 4 // client request to master
	MessageType_RETURN         MessageType = 5 // non-master to master or master to client
	MessageType_REPLICATION    MessageType = 6
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "MEMBERSHIP",
		1: "JOIN_REQ",
		2: "JOIN_REP",
		3: "FILE_OP",
		4: "CLIENT_REQUEST",
		5: "RETURN",
		6: "REPLICATION",
	}
	MessageType_value = map[string]int32{
		"MEMBERSHIP":     0,
		"JOIN_REQ":       1,
		"JOIN_REP":       2,
		"FILE_OP":        3,
		"CLIENT_REQUEST": 4,
		"RETURN":         5,
		"REPLICATION":    6,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type FileOp int32

const (
	FileOp_GET    FileOp = 0
	FileOp_PUT    FileOp = 1
	FileOp_DELETE FileOp = 2
	FileOp_LS     FileOp = 3
)

// Enum value maps for FileOp.
var (
	FileOp_name = map[int32]string{
		0: "GET",
		1: "PUT",
		2: "DELETE",
		3: "LS",
	}
	FileOp_value = map[string]int32{
		"GET":    0,
		"PUT":    1,
		"DELETE": 2,
		"LS":     3,
	}
)

func (x FileOp) Enum() *FileOp {
	p := new(FileOp)
	*p = x
	return p
}

func (x FileOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileOp) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[1].Descriptor()
}

func (FileOp) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[1]
}

func (x FileOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileOp.Descriptor instead.
func (FileOp) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

type MapleJuicePhase int32

const (
	MapleJuicePhase_INIT     MapleJuicePhase = 0
	MapleJuicePhase_ASSIGN   MapleJuicePhase = 1
	MapleJuicePhase_COMPLETE MapleJuicePhase = 2
	MapleJuicePhase_FAIL     MapleJuicePhase = 3
)

// Enum value maps for MapleJuicePhase.
var (
	MapleJuicePhase_name = map[int32]string{
		0: "INIT",
		1: "ASSIGN",
		2: "COMPLETE",
		3: "FAIL",
	}
	MapleJuicePhase_value = map[string]int32{
		"INIT":     0,
		"ASSIGN":   1,
		"COMPLETE": 2,
		"FAIL":     3,
	}
)

func (x MapleJuicePhase) Enum() *MapleJuicePhase {
	p := new(MapleJuicePhase)
	*p = x
	return p
}

func (x MapleJuicePhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapleJuicePhase) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[2].Descriptor()
}

func (MapleJuicePhase) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[2]
}

func (x MapleJuicePhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapleJuicePhase.Descriptor instead.
func (MapleJuicePhase) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

type MapleJuiceType int32

const (
	MapleJuiceType_MAPLE MapleJuiceType = 0
	MapleJuiceType_JUICE MapleJuiceType = 1
)

// Enum value maps for MapleJuiceType.
var (
	MapleJuiceType_name = map[int32]string{
		0: "MAPLE",
		1: "JUICE",
	}
	MapleJuiceType_value = map[string]int32{
		"MAPLE": 0,
		"JUICE": 1,
	}
)

func (x MapleJuiceType) Enum() *MapleJuiceType {
	p := new(MapleJuiceType)
	*p = x
	return p
}

func (x MapleJuiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MapleJuiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[3].Descriptor()
}

func (MapleJuiceType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[3]
}

func (x MapleJuiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MapleJuiceType.Descriptor instead.
func (MapleJuiceType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

type Partition int32

const (
	Partition_HASH  Partition = 0
	Partition_RANGE Partition = 1
)

// Enum value maps for Partition.
var (
	Partition_name = map[int32]string{
		0: "HASH",
		1: "RANGE",
	}
	Partition_value = map[string]int32{
		"HASH":  0,
		"RANGE": 1,
	}
)

func (x Partition) Enum() *Partition {
	p := new(Partition)
	*p = x
	return p
}

func (x Partition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Partition) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[4].Descriptor()
}

func (Partition) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[4]
}

func (x Partition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Partition.Descriptor instead.
func (Partition) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeartbeatCounter int32                `protobuf:"varint,1,opt,name=HeartbeatCounter,proto3" json:"HeartbeatCounter,omitempty"`
	LastSeen         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
	IsLeaving        bool                 `protobuf:"varint,3,opt,name=IsLeaving,proto3" json:"IsLeaving,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetHeartbeatCounter() int32 {
	if x != nil {
		return x.HeartbeatCounter
	}
	return 0
}

func (x *Member) GetLastSeen() *timestamp.Timestamp {
	if x != nil {
		return x.LastSeen
	}
	return nil
}

func (x *Member) GetIsLeaving() bool {
	if x != nil {
		return x.IsLeaving
	}
	return false
}

type MembershipMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          MessageType          `protobuf:"varint,1,opt,name=Type,proto3,enum=tutorial.MessageType" json:"Type,omitempty"`
	MemberList    map[string]*Member   `protobuf:"bytes,2,rep,name=MemberList,proto3" json:"MemberList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MasterID      string               `protobuf:"bytes,3,opt,name=MasterID,proto3" json:"MasterID,omitempty"`
	MasterCounter int32                `protobuf:"varint,4,opt,name=MasterCounter,proto3" json:"MasterCounter,omitempty"`
	FileMap       map[string]*FileInfo `protobuf:"bytes,5,rep,name=FileMap,proto3" json:"FileMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MembershipMessage) Reset() {
	*x = MembershipMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipMessage) ProtoMessage() {}

func (x *MembershipMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipMessage.ProtoReflect.Descriptor instead.
func (*MembershipMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *MembershipMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_MEMBERSHIP
}

func (x *MembershipMessage) GetMemberList() map[string]*Member {
	if x != nil {
		return x.MemberList
	}
	return nil
}

func (x *MembershipMessage) GetMasterID() string {
	if x != nil {
		return x.MasterID
	}
	return ""
}

func (x *MembershipMessage) GetMasterCounter() int32 {
	if x != nil {
		return x.MasterCounter
	}
	return 0
}

func (x *MembershipMessage) GetFileMap() map[string]*FileInfo {
	if x != nil {
		return x.FileMap
	}
	return nil
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineFileInfos map[string]*MachineFileInfo `protobuf:"bytes,1,rep,name=MachineFileInfos,proto3" json:"MachineFileInfos,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *FileInfo) GetMachineFileInfos() map[string]*MachineFileInfo {
	if x != nil {
		return x.MachineFileInfos
	}
	return nil
}

type MachineFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter  int32                `protobuf:"varint,1,opt,name=Counter,proto3" json:"Counter,omitempty"`
	LastSeen *timestamp.Timestamp `protobuf:"bytes,2,opt,name=LastSeen,proto3" json:"LastSeen,omitempty"`
}

func (x *MachineFileInfo) Reset() {
	*x = MachineFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineFileInfo) ProtoMessage() {}

func (x *MachineFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineFileInfo.ProtoReflect.Descriptor instead.
func (*MachineFileInfo) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *MachineFileInfo) GetCounter() int32 {
	if x != nil {
		return x.Counter
	}
	return 0
}

func (x *MachineFileInfo) GetLastSeen() *timestamp.Timestamp {
	if x != nil {
		return x.LastSeen
	}
	return nil
}

type SdfsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type          MessageType `protobuf:"varint,1,opt,name=Type,proto3,enum=tutorial.MessageType" json:"Type,omitempty"`
	FileOp        FileOp      `protobuf:"varint,2,opt,name=FileOp,proto3,enum=tutorial.FileOp" json:"FileOp,omitempty"` // get, put, delete, or ls
	LocalFilename string      `protobuf:"bytes,3,opt,name=LocalFilename,proto3" json:"LocalFilename,omitempty"`         // arg to file op
	SdfsFilename  string      `protobuf:"bytes,4,opt,name=SdfsFilename,proto3" json:"SdfsFilename,omitempty"`           // arg to file op
	ClientIP      string      `protobuf:"bytes,5,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`                   // used to keep track of which client sent the request (for when there are multiple reads at once)
	ReturnString  string      `protobuf:"bytes,6,opt,name=ReturnString,proto3" json:"ReturnString,omitempty"`           // holds return value for put, delete, or ls (or error message)
	FileSeqNum    int32       `protobuf:"varint,7,opt,name=FileSeqNum,proto3" json:"FileSeqNum,omitempty"`
	ShouldAppend  bool        `protobuf:"varint,8,opt,name=ShouldAppend,proto3" json:"ShouldAppend,omitempty"` // used for MapleJuice to always append to existing SDFS file
	File          []byte      `protobuf:"bytes,9,opt,name=File,proto3" json:"File,omitempty"`                  // arg to file op
}

func (x *SdfsMessage) Reset() {
	*x = SdfsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SdfsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SdfsMessage) ProtoMessage() {}

func (x *SdfsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SdfsMessage.ProtoReflect.Descriptor instead.
func (*SdfsMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SdfsMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_MEMBERSHIP
}

func (x *SdfsMessage) GetFileOp() FileOp {
	if x != nil {
		return x.FileOp
	}
	return FileOp_GET
}

func (x *SdfsMessage) GetLocalFilename() string {
	if x != nil {
		return x.LocalFilename
	}
	return ""
}

func (x *SdfsMessage) GetSdfsFilename() string {
	if x != nil {
		return x.SdfsFilename
	}
	return ""
}

func (x *SdfsMessage) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *SdfsMessage) GetReturnString() string {
	if x != nil {
		return x.ReturnString
	}
	return ""
}

func (x *SdfsMessage) GetFileSeqNum() int32 {
	if x != nil {
		return x.FileSeqNum
	}
	return 0
}

func (x *SdfsMessage) GetShouldAppend() bool {
	if x != nil {
		return x.ShouldAppend
	}
	return false
}

func (x *SdfsMessage) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

type MapleJuiceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phase              MapleJuicePhase `protobuf:"varint,1,opt,name=Phase,proto3,enum=tutorial.MapleJuicePhase" json:"Phase,omitempty"`
	TaskType           MapleJuiceType  `protobuf:"varint,2,opt,name=TaskType,proto3,enum=tutorial.MapleJuiceType" json:"TaskType,omitempty"`
	NumWorkers         int32           `protobuf:"varint,3,opt,name=NumWorkers,proto3" json:"NumWorkers,omitempty"`
	ExeFilename        string          `protobuf:"bytes,4,opt,name=ExeFilename,proto3" json:"ExeFilename,omitempty"`
	IntermediatePrefix string          `protobuf:"bytes,5,opt,name=IntermediatePrefix,proto3" json:"IntermediatePrefix,omitempty"`
	SrcDirectory       string          `protobuf:"bytes,6,opt,name=SrcDirectory,proto3" json:"SrcDirectory,omitempty"`                            // maple only
	DestFilename       string          `protobuf:"bytes,7,opt,name=DestFilename,proto3" json:"DestFilename,omitempty"`                            // juice only
	DeleteInput        bool            `protobuf:"varint,8,opt,name=DeleteInput,proto3" json:"DeleteInput,omitempty"`                             // juice only
	PartitionType      Partition       `protobuf:"varint,9,opt,name=PartitionType,proto3,enum=tutorial.Partition" json:"PartitionType,omitempty"` // juice only
	MapleInputFile     string          `protobuf:"bytes,10,opt,name=MapleInputFile,proto3" json:"MapleInputFile,omitempty"`                       // maple only
	JuiceInputKeys     []string        `protobuf:"bytes,11,rep,name=JuiceInputKeys,proto3" json:"JuiceInputKeys,omitempty"`                       // juice only
	MjGUID             string          `protobuf:"bytes,12,opt,name=MjGUID,proto3" json:"MjGUID,omitempty"`
	KeysToValues       []byte          `protobuf:"bytes,13,opt,name=KeysToValues,proto3" json:"KeysToValues,omitempty"`
}

func (x *MapleJuiceMessage) Reset() {
	*x = MapleJuiceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapleJuiceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapleJuiceMessage) ProtoMessage() {}

func (x *MapleJuiceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapleJuiceMessage.ProtoReflect.Descriptor instead.
func (*MapleJuiceMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{5}
}

func (x *MapleJuiceMessage) GetPhase() MapleJuicePhase {
	if x != nil {
		return x.Phase
	}
	return MapleJuicePhase_INIT
}

func (x *MapleJuiceMessage) GetTaskType() MapleJuiceType {
	if x != nil {
		return x.TaskType
	}
	return MapleJuiceType_MAPLE
}

func (x *MapleJuiceMessage) GetNumWorkers() int32 {
	if x != nil {
		return x.NumWorkers
	}
	return 0
}

func (x *MapleJuiceMessage) GetExeFilename() string {
	if x != nil {
		return x.ExeFilename
	}
	return ""
}

func (x *MapleJuiceMessage) GetIntermediatePrefix() string {
	if x != nil {
		return x.IntermediatePrefix
	}
	return ""
}

func (x *MapleJuiceMessage) GetSrcDirectory() string {
	if x != nil {
		return x.SrcDirectory
	}
	return ""
}

func (x *MapleJuiceMessage) GetDestFilename() string {
	if x != nil {
		return x.DestFilename
	}
	return ""
}

func (x *MapleJuiceMessage) GetDeleteInput() bool {
	if x != nil {
		return x.DeleteInput
	}
	return false
}

func (x *MapleJuiceMessage) GetPartitionType() Partition {
	if x != nil {
		return x.PartitionType
	}
	return Partition_HASH
}

func (x *MapleJuiceMessage) GetMapleInputFile() string {
	if x != nil {
		return x.MapleInputFile
	}
	return ""
}

func (x *MapleJuiceMessage) GetJuiceInputKeys() []string {
	if x != nil {
		return x.JuiceInputKeys
	}
	return nil
}

func (x *MapleJuiceMessage) GetMjGUID() string {
	if x != nil {
		return x.MjGUID
	}
	return ""
}

func (x *MapleJuiceMessage) GetKeysToValues() []byte {
	if x != nil {
		return x.KeysToValues
	}
	return nil
}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x06,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73,
	0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49,
	0x73, 0x4c, 0x65, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x22, 0xb2, 0x03, 0x0a, 0x11, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x1a, 0x4f, 0x0a, 0x0f,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a,
	0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc0, 0x01,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x10, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x1a, 0x5e, 0x0a, 0x15, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x63, 0x0a, 0x0f, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x4c, 0x61, 0x73,
	0x74, 0x53, 0x65, 0x65, 0x6e, 0x22, 0xc4, 0x02, 0x0a, 0x0b, 0x53, 0x64, 0x66, 0x73, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x28, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x64, 0x66, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x64, 0x66, 0x73, 0x46, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x71, 0x4e,
	0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65,
	0x71, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x41, 0x70,
	0x70, 0x65, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x53, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x9d, 0x04, 0x0a,
	0x11, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x4a, 0x75, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x50, 0x68, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x4d, 0x61, 0x70,
	0x6c, 0x65, 0x4a, 0x75, 0x69, 0x63, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x52, 0x05, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x4a, 0x75, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x75, 0x6d,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4e,
	0x75, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x78, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x45, 0x78, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x72, 0x63, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x72, 0x63, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x22, 0x0a, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x39, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4a, 0x75, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x4a, 0x75, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6a, 0x47, 0x55, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x6a, 0x47, 0x55, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x73,
	0x54, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x4b, 0x65, 0x79, 0x73, 0x54, 0x6f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0x77, 0x0a, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x53, 0x48, 0x49, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4a,
	0x4f, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4a, 0x4f, 0x49,
	0x4e, 0x5f, 0x52, 0x45, 0x50, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x4f, 0x50, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x54, 0x55,
	0x52, 0x4e, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x06, 0x2a, 0x2e, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x06, 0x0a,
	0x02, 0x4c, 0x53, 0x10, 0x03, 0x2a, 0x3f, 0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x4a, 0x75,
	0x69, 0x63, 0x65, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x53, 0x53, 0x49, 0x47, 0x4e, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x2a, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x6c, 0x65, 0x4a,
	0x75, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x50, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4a, 0x55, 0x49, 0x43, 0x45, 0x10, 0x01, 0x2a, 0x20,
	0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x48,
	0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_messages_proto_goTypes = []interface{}{
	(MessageType)(0),            // 0: tutorial.MessageType
	(FileOp)(0),                 // 1: tutorial.FileOp
	(MapleJuicePhase)(0),        // 2: tutorial.MapleJuicePhase
	(MapleJuiceType)(0),         // 3: tutorial.MapleJuiceType
	(Partition)(0),              // 4: tutorial.Partition
	(*Member)(nil),              // 5: tutorial.Member
	(*MembershipMessage)(nil),   // 6: tutorial.MembershipMessage
	(*FileInfo)(nil),            // 7: tutorial.FileInfo
	(*MachineFileInfo)(nil),     // 8: tutorial.MachineFileInfo
	(*SdfsMessage)(nil),         // 9: tutorial.SdfsMessage
	(*MapleJuiceMessage)(nil),   // 10: tutorial.MapleJuiceMessage
	nil,                         // 11: tutorial.MembershipMessage.MemberListEntry
	nil,                         // 12: tutorial.MembershipMessage.FileMapEntry
	nil,                         // 13: tutorial.FileInfo.MachineFileInfosEntry
	(*timestamp.Timestamp)(nil), // 14: google.protobuf.Timestamp
}
var file_messages_proto_depIdxs = []int32{
	14, // 0: tutorial.Member.LastSeen:type_name -> google.protobuf.Timestamp
	0,  // 1: tutorial.MembershipMessage.Type:type_name -> tutorial.MessageType
	11, // 2: tutorial.MembershipMessage.MemberList:type_name -> tutorial.MembershipMessage.MemberListEntry
	12, // 3: tutorial.MembershipMessage.FileMap:type_name -> tutorial.MembershipMessage.FileMapEntry
	13, // 4: tutorial.FileInfo.MachineFileInfos:type_name -> tutorial.FileInfo.MachineFileInfosEntry
	14, // 5: tutorial.MachineFileInfo.LastSeen:type_name -> google.protobuf.Timestamp
	0,  // 6: tutorial.SdfsMessage.Type:type_name -> tutorial.MessageType
	1,  // 7: tutorial.SdfsMessage.FileOp:type_name -> tutorial.FileOp
	2,  // 8: tutorial.MapleJuiceMessage.Phase:type_name -> tutorial.MapleJuicePhase
	3,  // 9: tutorial.MapleJuiceMessage.TaskType:type_name -> tutorial.MapleJuiceType
	4,  // 10: tutorial.MapleJuiceMessage.PartitionType:type_name -> tutorial.Partition
	5,  // 11: tutorial.MembershipMessage.MemberListEntry.value:type_name -> tutorial.Member
	7,  // 12: tutorial.MembershipMessage.FileMapEntry.value:type_name -> tutorial.FileInfo
	8,  // 13: tutorial.FileInfo.MachineFileInfosEntry.value:type_name -> tutorial.MachineFileInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembershipMessage); i {
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
		file_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineFileInfo); i {
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
		file_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SdfsMessage); i {
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
		file_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapleJuiceMessage); i {
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
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
