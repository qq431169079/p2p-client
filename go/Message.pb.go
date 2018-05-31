// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Message.proto

package botnet_p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_MessageType int32

const (
	Message_UNDEFINED   Message_MessageType = 0
	Message_COMMAND     Message_MessageType = 1
	Message_RESPONSE    Message_MessageType = 2
	Message_FILE_CHUNK  Message_MessageType = 3
	Message_NAT_REQUEST Message_MessageType = 4
	Message_NAT_CHECK   Message_MessageType = 5
	Message_PING        Message_MessageType = 6
	Message_LEAVE       Message_MessageType = 7
	Message_FIND_NODE   Message_MessageType = 8
	Message_FOUND_NODES Message_MessageType = 9
)

var Message_MessageType_name = map[int32]string{
	0: "UNDEFINED",
	1: "COMMAND",
	2: "RESPONSE",
	3: "FILE_CHUNK",
	4: "NAT_REQUEST",
	5: "NAT_CHECK",
	6: "PING",
	7: "LEAVE",
	8: "FIND_NODE",
	9: "FOUND_NODES",
}
var Message_MessageType_value = map[string]int32{
	"UNDEFINED":   0,
	"COMMAND":     1,
	"RESPONSE":    2,
	"FILE_CHUNK":  3,
	"NAT_REQUEST": 4,
	"NAT_CHECK":   5,
	"PING":        6,
	"LEAVE":       7,
	"FIND_NODE":   8,
	"FOUND_NODES": 9,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 0}
}

type Message_Status int32

const (
	Message_FAIL Message_Status = 0
	Message_OK   Message_Status = 1
)

var Message_Status_name = map[int32]string{
	0: "FAIL",
	1: "OK",
}
var Message_Status_value = map[string]int32{
	"FAIL": 0,
	"OK":   1,
}

func (x Message_Status) String() string {
	return proto.EnumName(Message_Status_name, int32(x))
}
func (Message_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 1}
}

type Message struct {
	Uuid        string              `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Type        Message_MessageType `protobuf:"varint,2,opt,name=type,enum=botnet_p2p.Message_MessageType" json:"type,omitempty"`
	Sender      string              `protobuf:"bytes,3,opt,name=sender" json:"sender,omitempty"`
	Receiver    string              `protobuf:"bytes,4,opt,name=receiver" json:"receiver,omitempty"`
	Propagation bool                `protobuf:"varint,5,opt,name=propagation" json:"propagation,omitempty"`
	Signature   []byte              `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Message_Command
	//	*Message_Response
	//	*Message_FileChunk
	//	*Message_NATRequest
	//	*Message_NATCheck
	//	*Message_Leave
	//	*Message_FindNode
	//	*Message_FoundNodes
	//	*Message_Ping
	Payload              isMessage_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Command struct {
	Command *Message_CommandMsg `protobuf:"bytes,7,opt,name=command,oneof"`
}
type Message_Response struct {
	Response *Message_ResponseMsg `protobuf:"bytes,8,opt,name=response,oneof"`
}
type Message_FileChunk struct {
	FileChunk *Message_FileChunkMsg `protobuf:"bytes,9,opt,name=fileChunk,oneof"`
}
type Message_NATRequest struct {
	NATRequest *Message_NATRequestMsg `protobuf:"bytes,10,opt,name=NATRequest,oneof"`
}
type Message_NATCheck struct {
	NATCheck *Message_NATCheckMsg `protobuf:"bytes,11,opt,name=NATCheck,oneof"`
}
type Message_Leave struct {
	Leave *Message_LeaveMsg `protobuf:"bytes,12,opt,name=leave,oneof"`
}
type Message_FindNode struct {
	FindNode *Message_FindNodeMsg `protobuf:"bytes,13,opt,name=findNode,oneof"`
}
type Message_FoundNodes struct {
	FoundNodes *Message_FoundNodesMsg `protobuf:"bytes,14,opt,name=foundNodes,oneof"`
}
type Message_Ping struct {
	Ping *Message_PingMsg `protobuf:"bytes,15,opt,name=ping,oneof"`
}

func (*Message_Command) isMessage_Payload()    {}
func (*Message_Response) isMessage_Payload()   {}
func (*Message_FileChunk) isMessage_Payload()  {}
func (*Message_NATRequest) isMessage_Payload() {}
func (*Message_NATCheck) isMessage_Payload()   {}
func (*Message_Leave) isMessage_Payload()      {}
func (*Message_FindNode) isMessage_Payload()   {}
func (*Message_FoundNodes) isMessage_Payload() {}
func (*Message_Ping) isMessage_Payload()       {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Message) GetType() Message_MessageType {
	if m != nil {
		return m.Type
	}
	return Message_UNDEFINED
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *Message) GetPropagation() bool {
	if m != nil {
		return m.Propagation
	}
	return false
}

func (m *Message) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Message) GetCommand() *Message_CommandMsg {
	if x, ok := m.GetPayload().(*Message_Command); ok {
		return x.Command
	}
	return nil
}

func (m *Message) GetResponse() *Message_ResponseMsg {
	if x, ok := m.GetPayload().(*Message_Response); ok {
		return x.Response
	}
	return nil
}

func (m *Message) GetFileChunk() *Message_FileChunkMsg {
	if x, ok := m.GetPayload().(*Message_FileChunk); ok {
		return x.FileChunk
	}
	return nil
}

func (m *Message) GetNATRequest() *Message_NATRequestMsg {
	if x, ok := m.GetPayload().(*Message_NATRequest); ok {
		return x.NATRequest
	}
	return nil
}

func (m *Message) GetNATCheck() *Message_NATCheckMsg {
	if x, ok := m.GetPayload().(*Message_NATCheck); ok {
		return x.NATCheck
	}
	return nil
}

func (m *Message) GetLeave() *Message_LeaveMsg {
	if x, ok := m.GetPayload().(*Message_Leave); ok {
		return x.Leave
	}
	return nil
}

func (m *Message) GetFindNode() *Message_FindNodeMsg {
	if x, ok := m.GetPayload().(*Message_FindNode); ok {
		return x.FindNode
	}
	return nil
}

func (m *Message) GetFoundNodes() *Message_FoundNodesMsg {
	if x, ok := m.GetPayload().(*Message_FoundNodes); ok {
		return x.FoundNodes
	}
	return nil
}

func (m *Message) GetPing() *Message_PingMsg {
	if x, ok := m.GetPayload().(*Message_Ping); ok {
		return x.Ping
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Command)(nil),
		(*Message_Response)(nil),
		(*Message_FileChunk)(nil),
		(*Message_NATRequest)(nil),
		(*Message_NATCheck)(nil),
		(*Message_Leave)(nil),
		(*Message_FindNode)(nil),
		(*Message_FoundNodes)(nil),
		(*Message_Ping)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_Command:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Command); err != nil {
			return err
		}
	case *Message_Response:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Response); err != nil {
			return err
		}
	case *Message_FileChunk:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FileChunk); err != nil {
			return err
		}
	case *Message_NATRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NATRequest); err != nil {
			return err
		}
	case *Message_NATCheck:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NATCheck); err != nil {
			return err
		}
	case *Message_Leave:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Leave); err != nil {
			return err
		}
	case *Message_FindNode:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FindNode); err != nil {
			return err
		}
	case *Message_FoundNodes:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FoundNodes); err != nil {
			return err
		}
	case *Message_Ping:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ping); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 7: // payload.command
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_CommandMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Command{msg}
		return true, err
	case 8: // payload.response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_ResponseMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Response{msg}
		return true, err
	case 9: // payload.fileChunk
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_FileChunkMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_FileChunk{msg}
		return true, err
	case 10: // payload.NATRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_NATRequestMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_NATRequest{msg}
		return true, err
	case 11: // payload.NATCheck
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_NATCheckMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_NATCheck{msg}
		return true, err
	case 12: // payload.leave
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_LeaveMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Leave{msg}
		return true, err
	case 13: // payload.findNode
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_FindNodeMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_FindNode{msg}
		return true, err
	case 14: // payload.foundNodes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_FoundNodesMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_FoundNodes{msg}
		return true, err
	case 15: // payload.ping
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Message_PingMsg)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Ping{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_Command:
		s := proto.Size(x.Command)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Response:
		s := proto.Size(x.Response)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_FileChunk:
		s := proto.Size(x.FileChunk)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_NATRequest:
		s := proto.Size(x.NATRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_NATCheck:
		s := proto.Size(x.NATCheck)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Leave:
		s := proto.Size(x.Leave)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_FindNode:
		s := proto.Size(x.FindNode)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_FoundNodes:
		s := proto.Size(x.FoundNodes)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Ping:
		s := proto.Size(x.Ping)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Message_Contact struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP" json:"IP,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	IsNAT                bool     `protobuf:"varint,4,opt,name=isNAT" json:"isNAT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Contact) Reset()         { *m = Message_Contact{} }
func (m *Message_Contact) String() string { return proto.CompactTextString(m) }
func (*Message_Contact) ProtoMessage()    {}
func (*Message_Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 0}
}
func (m *Message_Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Contact.Unmarshal(m, b)
}
func (m *Message_Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Contact.Marshal(b, m, deterministic)
}
func (dst *Message_Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Contact.Merge(dst, src)
}
func (m *Message_Contact) XXX_Size() int {
	return xxx_messageInfo_Message_Contact.Size(m)
}
func (m *Message_Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Contact proto.InternalMessageInfo

func (m *Message_Contact) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Message_Contact) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *Message_Contact) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Message_Contact) GetIsNAT() bool {
	if m != nil {
		return m.IsNAT
	}
	return false
}

type Message_CommandMsg struct {
	CommandString        string   `protobuf:"bytes,1,opt,name=commandString" json:"commandString,omitempty"`
	SendResponse         bool     `protobuf:"varint,2,opt,name=sendResponse" json:"sendResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_CommandMsg) Reset()         { *m = Message_CommandMsg{} }
func (m *Message_CommandMsg) String() string { return proto.CompactTextString(m) }
func (*Message_CommandMsg) ProtoMessage()    {}
func (*Message_CommandMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 1}
}
func (m *Message_CommandMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_CommandMsg.Unmarshal(m, b)
}
func (m *Message_CommandMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_CommandMsg.Marshal(b, m, deterministic)
}
func (dst *Message_CommandMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_CommandMsg.Merge(dst, src)
}
func (m *Message_CommandMsg) XXX_Size() int {
	return xxx_messageInfo_Message_CommandMsg.Size(m)
}
func (m *Message_CommandMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_CommandMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_CommandMsg proto.InternalMessageInfo

func (m *Message_CommandMsg) GetCommandString() string {
	if m != nil {
		return m.CommandString
	}
	return ""
}

func (m *Message_CommandMsg) GetSendResponse() bool {
	if m != nil {
		return m.SendResponse
	}
	return false
}

type Message_ResponseMsg struct {
	Value                string         `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Status               Message_Status `protobuf:"varint,2,opt,name=status,enum=botnet_p2p.Message_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message_ResponseMsg) Reset()         { *m = Message_ResponseMsg{} }
func (m *Message_ResponseMsg) String() string { return proto.CompactTextString(m) }
func (*Message_ResponseMsg) ProtoMessage()    {}
func (*Message_ResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 2}
}
func (m *Message_ResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_ResponseMsg.Unmarshal(m, b)
}
func (m *Message_ResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_ResponseMsg.Marshal(b, m, deterministic)
}
func (dst *Message_ResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_ResponseMsg.Merge(dst, src)
}
func (m *Message_ResponseMsg) XXX_Size() int {
	return xxx_messageInfo_Message_ResponseMsg.Size(m)
}
func (m *Message_ResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_ResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_ResponseMsg proto.InternalMessageInfo

func (m *Message_ResponseMsg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Message_ResponseMsg) GetStatus() Message_Status {
	if m != nil {
		return m.Status
	}
	return Message_FAIL
}

type Message_FileChunkMsg struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Name                 bool     `protobuf:"varint,2,opt,name=name" json:"name,omitempty"`
	ChunkNumber          uint32   `protobuf:"varint,3,opt,name=chunkNumber" json:"chunkNumber,omitempty"`
	AllChunks            uint32   `protobuf:"varint,4,opt,name=allChunks" json:"allChunks,omitempty"`
	ChunkSize            uint32   `protobuf:"varint,5,opt,name=chunkSize" json:"chunkSize,omitempty"`
	Data                 []byte   `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_FileChunkMsg) Reset()         { *m = Message_FileChunkMsg{} }
func (m *Message_FileChunkMsg) String() string { return proto.CompactTextString(m) }
func (*Message_FileChunkMsg) ProtoMessage()    {}
func (*Message_FileChunkMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 3}
}
func (m *Message_FileChunkMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_FileChunkMsg.Unmarshal(m, b)
}
func (m *Message_FileChunkMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_FileChunkMsg.Marshal(b, m, deterministic)
}
func (dst *Message_FileChunkMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_FileChunkMsg.Merge(dst, src)
}
func (m *Message_FileChunkMsg) XXX_Size() int {
	return xxx_messageInfo_Message_FileChunkMsg.Size(m)
}
func (m *Message_FileChunkMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_FileChunkMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_FileChunkMsg proto.InternalMessageInfo

func (m *Message_FileChunkMsg) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Message_FileChunkMsg) GetName() bool {
	if m != nil {
		return m.Name
	}
	return false
}

func (m *Message_FileChunkMsg) GetChunkNumber() uint32 {
	if m != nil {
		return m.ChunkNumber
	}
	return 0
}

func (m *Message_FileChunkMsg) GetAllChunks() uint32 {
	if m != nil {
		return m.AllChunks
	}
	return 0
}

func (m *Message_FileChunkMsg) GetChunkSize() uint32 {
	if m != nil {
		return m.ChunkSize
	}
	return 0
}

func (m *Message_FileChunkMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message_NATRequestMsg struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_NATRequestMsg) Reset()         { *m = Message_NATRequestMsg{} }
func (m *Message_NATRequestMsg) String() string { return proto.CompactTextString(m) }
func (*Message_NATRequestMsg) ProtoMessage()    {}
func (*Message_NATRequestMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 4}
}
func (m *Message_NATRequestMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_NATRequestMsg.Unmarshal(m, b)
}
func (m *Message_NATRequestMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_NATRequestMsg.Marshal(b, m, deterministic)
}
func (dst *Message_NATRequestMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_NATRequestMsg.Merge(dst, src)
}
func (m *Message_NATRequestMsg) XXX_Size() int {
	return xxx_messageInfo_Message_NATRequestMsg.Size(m)
}
func (m *Message_NATRequestMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_NATRequestMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_NATRequestMsg proto.InternalMessageInfo

func (m *Message_NATRequestMsg) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type Message_NATCheckMsg struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_NATCheckMsg) Reset()         { *m = Message_NATCheckMsg{} }
func (m *Message_NATCheckMsg) String() string { return proto.CompactTextString(m) }
func (*Message_NATCheckMsg) ProtoMessage()    {}
func (*Message_NATCheckMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 5}
}
func (m *Message_NATCheckMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_NATCheckMsg.Unmarshal(m, b)
}
func (m *Message_NATCheckMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_NATCheckMsg.Marshal(b, m, deterministic)
}
func (dst *Message_NATCheckMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_NATCheckMsg.Merge(dst, src)
}
func (m *Message_NATCheckMsg) XXX_Size() int {
	return xxx_messageInfo_Message_NATCheckMsg.Size(m)
}
func (m *Message_NATCheckMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_NATCheckMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_NATCheckMsg proto.InternalMessageInfo

func (m *Message_NATCheckMsg) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type Message_LeaveMsg struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_LeaveMsg) Reset()         { *m = Message_LeaveMsg{} }
func (m *Message_LeaveMsg) String() string { return proto.CompactTextString(m) }
func (*Message_LeaveMsg) ProtoMessage()    {}
func (*Message_LeaveMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 6}
}
func (m *Message_LeaveMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_LeaveMsg.Unmarshal(m, b)
}
func (m *Message_LeaveMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_LeaveMsg.Marshal(b, m, deterministic)
}
func (dst *Message_LeaveMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_LeaveMsg.Merge(dst, src)
}
func (m *Message_LeaveMsg) XXX_Size() int {
	return xxx_messageInfo_Message_LeaveMsg.Size(m)
}
func (m *Message_LeaveMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_LeaveMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_LeaveMsg proto.InternalMessageInfo

func (m *Message_LeaveMsg) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type Message_FindNodeMsg struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_FindNodeMsg) Reset()         { *m = Message_FindNodeMsg{} }
func (m *Message_FindNodeMsg) String() string { return proto.CompactTextString(m) }
func (*Message_FindNodeMsg) ProtoMessage()    {}
func (*Message_FindNodeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 7}
}
func (m *Message_FindNodeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_FindNodeMsg.Unmarshal(m, b)
}
func (m *Message_FindNodeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_FindNodeMsg.Marshal(b, m, deterministic)
}
func (dst *Message_FindNodeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_FindNodeMsg.Merge(dst, src)
}
func (m *Message_FindNodeMsg) XXX_Size() int {
	return xxx_messageInfo_Message_FindNodeMsg.Size(m)
}
func (m *Message_FindNodeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_FindNodeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_FindNodeMsg proto.InternalMessageInfo

func (m *Message_FindNodeMsg) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type Message_PingMsg struct {
	Guid                 string   `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_PingMsg) Reset()         { *m = Message_PingMsg{} }
func (m *Message_PingMsg) String() string { return proto.CompactTextString(m) }
func (*Message_PingMsg) ProtoMessage()    {}
func (*Message_PingMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 8}
}
func (m *Message_PingMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_PingMsg.Unmarshal(m, b)
}
func (m *Message_PingMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_PingMsg.Marshal(b, m, deterministic)
}
func (dst *Message_PingMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_PingMsg.Merge(dst, src)
}
func (m *Message_PingMsg) XXX_Size() int {
	return xxx_messageInfo_Message_PingMsg.Size(m)
}
func (m *Message_PingMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_PingMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_PingMsg proto.InternalMessageInfo

func (m *Message_PingMsg) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type Message_FoundNodesMsg struct {
	Nodes                []*Message_Contact `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Message_FoundNodesMsg) Reset()         { *m = Message_FoundNodesMsg{} }
func (m *Message_FoundNodesMsg) String() string { return proto.CompactTextString(m) }
func (*Message_FoundNodesMsg) ProtoMessage()    {}
func (*Message_FoundNodesMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_Message_5d4a47f3ca122b15, []int{0, 9}
}
func (m *Message_FoundNodesMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_FoundNodesMsg.Unmarshal(m, b)
}
func (m *Message_FoundNodesMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_FoundNodesMsg.Marshal(b, m, deterministic)
}
func (dst *Message_FoundNodesMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_FoundNodesMsg.Merge(dst, src)
}
func (m *Message_FoundNodesMsg) XXX_Size() int {
	return xxx_messageInfo_Message_FoundNodesMsg.Size(m)
}
func (m *Message_FoundNodesMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_FoundNodesMsg.DiscardUnknown(m)
}

var xxx_messageInfo_Message_FoundNodesMsg proto.InternalMessageInfo

func (m *Message_FoundNodesMsg) GetNodes() []*Message_Contact {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "botnet_p2p.Message")
	proto.RegisterType((*Message_Contact)(nil), "botnet_p2p.Message.Contact")
	proto.RegisterType((*Message_CommandMsg)(nil), "botnet_p2p.Message.CommandMsg")
	proto.RegisterType((*Message_ResponseMsg)(nil), "botnet_p2p.Message.ResponseMsg")
	proto.RegisterType((*Message_FileChunkMsg)(nil), "botnet_p2p.Message.FileChunkMsg")
	proto.RegisterType((*Message_NATRequestMsg)(nil), "botnet_p2p.Message.NATRequestMsg")
	proto.RegisterType((*Message_NATCheckMsg)(nil), "botnet_p2p.Message.NATCheckMsg")
	proto.RegisterType((*Message_LeaveMsg)(nil), "botnet_p2p.Message.LeaveMsg")
	proto.RegisterType((*Message_FindNodeMsg)(nil), "botnet_p2p.Message.FindNodeMsg")
	proto.RegisterType((*Message_PingMsg)(nil), "botnet_p2p.Message.PingMsg")
	proto.RegisterType((*Message_FoundNodesMsg)(nil), "botnet_p2p.Message.FoundNodesMsg")
	proto.RegisterEnum("botnet_p2p.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("botnet_p2p.Message_Status", Message_Status_name, Message_Status_value)
}

func init() { proto.RegisterFile("Message.proto", fileDescriptor_Message_5d4a47f3ca122b15) }

var fileDescriptor_Message_5d4a47f3ca122b15 = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6e, 0xea, 0x46,
	0x10, 0x8e, 0x09, 0x60, 0x7b, 0xc0, 0x1c, 0x6b, 0x55, 0x55, 0x96, 0x7b, 0x7a, 0xea, 0x43, 0x7b,
	0xc1, 0x55, 0xa4, 0x90, 0x5e, 0x55, 0xaa, 0x54, 0x62, 0xec, 0x06, 0x05, 0x0c, 0x5d, 0x20, 0xb9,
	0x44, 0x1b, 0xd8, 0x10, 0x2b, 0x60, 0xbb, 0xfe, 0x89, 0x94, 0x3e, 0x4b, 0x1f, 0xa0, 0x2f, 0xd5,
	0x77, 0x39, 0xda, 0xf5, 0x3a, 0x36, 0x92, 0xb9, 0x62, 0x66, 0xf6, 0xfb, 0xbe, 0xd9, 0x59, 0xbe,
	0x01, 0xd0, 0x66, 0x34, 0x49, 0xc8, 0x9e, 0x5e, 0x45, 0x71, 0x98, 0x86, 0x08, 0x9e, 0xc2, 0x34,
	0xa0, 0xe9, 0x26, 0x1a, 0x46, 0xfd, 0xff, 0xbb, 0x20, 0x8b, 0x53, 0x84, 0xa0, 0x99, 0x65, 0xfe,
	0xce, 0x90, 0x2c, 0x69, 0xa0, 0x62, 0x1e, 0xa3, 0x1b, 0x68, 0xa6, 0xef, 0x11, 0x35, 0x1a, 0x96,
	0x34, 0xe8, 0x0d, 0x7f, 0xba, 0x2a, 0xa9, 0x57, 0x85, 0xa8, 0xf8, 0x5c, 0xbd, 0x47, 0x14, 0x73,
	0x30, 0xfa, 0x1e, 0xda, 0x09, 0x0d, 0x76, 0x34, 0x36, 0x2e, 0xb9, 0x94, 0xc8, 0x90, 0x09, 0x4a,
	0x4c, 0xb7, 0xd4, 0x7f, 0xa3, 0xb1, 0xd1, 0xe4, 0x27, 0x1f, 0x39, 0xb2, 0xa0, 0x13, 0xc5, 0x61,
	0x44, 0xf6, 0x24, 0xf5, 0xc3, 0xc0, 0x68, 0x59, 0xd2, 0x40, 0xc1, 0xd5, 0x12, 0xfa, 0x0c, 0x6a,
	0xe2, 0xef, 0x03, 0x92, 0x66, 0x31, 0x35, 0xda, 0x96, 0x34, 0xe8, 0xe2, 0xb2, 0x80, 0x7e, 0x03,
	0x79, 0x1b, 0x1e, 0x8f, 0x24, 0xd8, 0x19, 0xb2, 0x25, 0x0d, 0x3a, 0xc3, 0x2f, 0x75, 0x77, 0xb5,
	0x73, 0xc8, 0x2c, 0xd9, 0xdf, 0x5d, 0xe0, 0x82, 0x80, 0x7e, 0x67, 0xf7, 0x4a, 0xa2, 0x30, 0x48,
	0xa8, 0xa1, 0x70, 0x72, 0xed, 0xa0, 0x58, 0x60, 0x72, 0xf6, 0x07, 0x05, 0xfd, 0x01, 0xea, 0xb3,
	0x7f, 0xa0, 0xf6, 0x4b, 0x16, 0xbc, 0x1a, 0x2a, 0xe7, 0x5b, 0x75, 0x7c, 0xb7, 0x00, 0xe5, 0x02,
	0x25, 0x09, 0xd9, 0x00, 0xde, 0x68, 0x85, 0xe9, 0xdf, 0x19, 0x4d, 0x52, 0x03, 0xb8, 0xc4, 0xd7,
	0x3a, 0x89, 0x12, 0x95, 0x6b, 0x54, 0x68, 0x6c, 0x0a, 0x6f, 0xb4, 0xb2, 0x5f, 0xe8, 0xf6, 0xd5,
	0xe8, 0x9c, 0x9f, 0xa2, 0xc0, 0x88, 0x29, 0x8a, 0x14, 0xfd, 0x0a, 0xad, 0x03, 0x25, 0x6f, 0xd4,
	0xe8, 0x72, 0xee, 0xe7, 0x3a, 0xee, 0x94, 0x01, 0x72, 0x62, 0x0e, 0x66, 0x4d, 0x9f, 0xfd, 0x60,
	0xe7, 0x85, 0x3b, 0x6a, 0x68, 0xe7, 0x9b, 0xba, 0x02, 0x23, 0x9a, 0x16, 0x14, 0x36, 0xf8, 0x73,
	0x98, 0xe5, 0x49, 0x62, 0xf4, 0xce, 0x0f, 0xee, 0x7e, 0xa0, 0xc4, 0xe0, 0x25, 0x0d, 0x5d, 0x43,
	0x33, 0xf2, 0x83, 0xbd, 0xf1, 0x89, 0xd3, 0x7f, 0xa8, 0xa3, 0x2f, 0xfc, 0x60, 0x9f, 0x13, 0x39,
	0xd4, 0x7c, 0x04, 0xd9, 0x0e, 0x83, 0x94, 0x6c, 0x53, 0xe6, 0xfa, 0x7d, 0xc5, 0xf5, 0x2c, 0x46,
	0x3d, 0x68, 0x4c, 0x16, 0xdc, 0xf3, 0x2a, 0x6e, 0x4c, 0x16, 0x0c, 0x13, 0x85, 0x71, 0xca, 0xed,
	0xac, 0x61, 0x1e, 0xa3, 0xef, 0xa0, 0xe5, 0x27, 0xde, 0x68, 0xc5, 0x9d, 0xac, 0xe0, 0x3c, 0x31,
	0x1f, 0x00, 0x4a, 0x8f, 0xa1, 0x5f, 0x40, 0x13, 0x1e, 0x5b, 0xa6, 0x31, 0xbb, 0x62, 0xde, 0xe4,
	0xb4, 0x88, 0xfa, 0xd0, 0x65, 0x0b, 0x52, 0xd8, 0x8b, 0xf7, 0x55, 0xf0, 0x49, 0xcd, 0x7c, 0x84,
	0x4e, 0xc5, 0x7e, 0xac, 0xf9, 0x1b, 0x39, 0x64, 0x54, 0x08, 0xe6, 0x09, 0x1a, 0x42, 0x3b, 0x49,
	0x49, 0x9a, 0x25, 0x62, 0x5d, 0xcd, 0xba, 0xa7, 0x58, 0x72, 0x04, 0x16, 0x48, 0xf3, 0x3f, 0x09,
	0xba, 0x55, 0x63, 0xf2, 0x59, 0x49, 0xfa, 0x52, 0xbc, 0x07, 0x8b, 0x59, 0x2d, 0x20, 0xc7, 0xe2,
	0x66, 0x3c, 0x66, 0x0b, 0xbb, 0x65, 0x1c, 0x2f, 0x3b, 0x3e, 0x89, 0x4d, 0xd7, 0x70, 0xb5, 0xc4,
	0x16, 0x96, 0x1c, 0x0e, 0x5c, 0x38, 0xe1, 0xaf, 0xa4, 0xe1, 0xb2, 0xc0, 0x4e, 0x39, 0x78, 0xe9,
	0xff, 0x43, 0xf9, 0xba, 0x6b, 0xb8, 0x2c, 0xb0, 0x8e, 0x3b, 0x92, 0x12, 0xb1, 0xe7, 0x3c, 0x36,
	0x7f, 0x06, 0xed, 0xc4, 0xff, 0x75, 0x5f, 0x9d, 0xf9, 0x15, 0x3a, 0x15, 0x87, 0xd7, 0x42, 0xbe,
	0x80, 0x52, 0x18, 0xf9, 0x9c, 0x44, 0xc5, 0xaf, 0xb5, 0x90, 0x1f, 0x41, 0x16, 0x96, 0xaa, 0x3d,
	0xbe, 0x05, 0xed, 0xc4, 0xb0, 0xe8, 0x1a, 0x5a, 0x01, 0xb7, 0xb8, 0x64, 0x5d, 0x9e, 0xf3, 0xa8,
	0x30, 0x24, 0xce, 0x91, 0xfd, 0x7f, 0x25, 0xe8, 0x54, 0x7e, 0x5a, 0x91, 0x06, 0xea, 0xda, 0x1b,
	0x3b, 0xee, 0xc4, 0x73, 0xc6, 0xfa, 0x05, 0xea, 0x80, 0x6c, 0xcf, 0x67, 0xb3, 0x91, 0x37, 0xd6,
	0x25, 0xd4, 0x05, 0x05, 0x3b, 0xcb, 0xc5, 0xdc, 0x5b, 0x3a, 0x7a, 0x03, 0xf5, 0x00, 0xdc, 0xc9,
	0xd4, 0xd9, 0xd8, 0x77, 0x6b, 0xef, 0x5e, 0xbf, 0x44, 0x9f, 0xf8, 0x93, 0x6c, 0xb0, 0xf3, 0xd7,
	0xda, 0x59, 0xae, 0xf4, 0x26, 0x93, 0x62, 0x05, 0xfb, 0xce, 0xb1, 0xef, 0xf5, 0x16, 0x52, 0xa0,
	0xb9, 0x98, 0x78, 0x7f, 0xea, 0x6d, 0xa4, 0x42, 0x6b, 0xea, 0x8c, 0x1e, 0x1c, 0x5d, 0x66, 0x18,
	0x77, 0xe2, 0x8d, 0x37, 0xde, 0x7c, 0xec, 0xe8, 0x0a, 0xd3, 0x70, 0xe7, 0x6b, 0x91, 0x2f, 0x75,
	0xb5, 0x6f, 0x42, 0x3b, 0x77, 0x12, 0xa3, 0xbb, 0xa3, 0xc9, 0x54, 0xbf, 0x40, 0x6d, 0x68, 0xcc,
	0xef, 0x75, 0xe9, 0x56, 0x05, 0x39, 0x22, 0xef, 0x87, 0x90, 0xec, 0x9e, 0xda, 0xfc, 0x2f, 0xe7,
	0xe6, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x9a, 0xfb, 0x8e, 0x83, 0x06, 0x00, 0x00,
}
