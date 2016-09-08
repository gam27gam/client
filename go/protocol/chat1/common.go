// Auto-generated by avdl-compiler v1.3.7 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/common.avdl

package chat1

import (
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type ThreadID []byte
type MessageID uint
type TopicID []byte
type ConversationID uint64
type TLFID []byte
type MessageType int

const (
	MessageType_NONE       MessageType = 0
	MessageType_TEXT       MessageType = 1
	MessageType_ATTACHMENT MessageType = 2
	MessageType_EDIT       MessageType = 3
	MessageType_DELETE     MessageType = 4
	MessageType_METADATA   MessageType = 5
	MessageType_TLFNAME    MessageType = 6
)

var MessageTypeMap = map[string]MessageType{
	"NONE":       0,
	"TEXT":       1,
	"ATTACHMENT": 2,
	"EDIT":       3,
	"DELETE":     4,
	"METADATA":   5,
	"TLFNAME":    6,
}

var MessageTypeRevMap = map[MessageType]string{
	0: "NONE",
	1: "TEXT",
	2: "ATTACHMENT",
	3: "EDIT",
	4: "DELETE",
	5: "METADATA",
	6: "TLFNAME",
}

type TopicType int

const (
	TopicType_NONE TopicType = 0
	TopicType_CHAT TopicType = 1
	TopicType_DEV  TopicType = 2
)

var TopicTypeMap = map[string]TopicType{
	"NONE": 0,
	"CHAT": 1,
	"DEV":  2,
}

var TopicTypeRevMap = map[TopicType]string{
	0: "NONE",
	1: "CHAT",
	2: "DEV",
}

type Pagination struct {
	Next     []byte `codec:"next" json:"next"`
	Previous []byte `codec:"previous" json:"previous"`
	Num      int    `codec:"num" json:"num"`
	Last     bool   `codec:"last" json:"last"`
}

type RateLimit struct {
	CallsRemaining int `codec:"callsRemaining" json:"callsRemaining"`
	WindowReset    int `codec:"windowReset" json:"windowReset"`
	MaxCalls       int `codec:"maxCalls" json:"maxCalls"`
}

type ConversationIDTriple struct {
	Tlfid     TLFID     `codec:"tlfid" json:"tlfid"`
	TopicType TopicType `codec:"topicType" json:"topicType"`
	TopicID   TopicID   `codec:"topicID" json:"topicID"`
}

type ConversationMetadata struct {
	IdTriple       ConversationIDTriple `codec:"idTriple" json:"idTriple"`
	ConversationID ConversationID       `codec:"conversationID" json:"conversationID"`
}

type ConversationReaderInfo struct {
	Mtime     gregor1.Time `codec:"mtime" json:"mtime"`
	ReadMsgid MessageID    `codec:"readMsgid" json:"readMsgid"`
	MaxMsgid  MessageID    `codec:"maxMsgid" json:"maxMsgid"`
}

type Conversation struct {
	Metadata   ConversationMetadata    `codec:"metadata" json:"metadata"`
	ReaderInfo *ConversationReaderInfo `codec:"readerInfo,omitempty" json:"readerInfo,omitempty"`
	MaxHeaders []MessageServerHeader   `codec:"maxHeaders" json:"maxHeaders"`
}

type MessageServerHeader struct {
	MessageType  MessageType      `codec:"messageType" json:"messageType"`
	MessageID    MessageID        `codec:"messageID" json:"messageID"`
	Sender       gregor1.UID      `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID `codec:"senderDevice" json:"senderDevice"`
	SupersededBy MessageID        `codec:"supersededBy" json:"supersededBy"`
	Ctime        gregor1.Time     `codec:"ctime" json:"ctime"`
}

type MessagePreviousPointer struct {
	Id   MessageID `codec:"id" json:"id"`
	Hash []byte    `codec:"hash" json:"hash"`
}

type MessageClientHeader struct {
	Conv         ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName      string                   `codec:"tlfName" json:"tlfName"`
	MessageType  MessageType              `codec:"messageType" json:"messageType"`
	Prev         []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender       gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
}

type EncryptedData struct {
	V int    `codec:"v" json:"v"`
	E []byte `codec:"e" json:"e"`
	N []byte `codec:"n" json:"n"`
}

type SignatureInfo struct {
	V int    `codec:"v" json:"v"`
	S []byte `codec:"s" json:"s"`
	K []byte `codec:"k" json:"k"`
}

type InboxView struct {
	Conversations []Conversation `codec:"conversations" json:"conversations"`
	Pagination    *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type CommonInterface interface {
}

func CommonProtocol(i CommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type CommonClient struct {
	Cli rpc.GenericClient
}
