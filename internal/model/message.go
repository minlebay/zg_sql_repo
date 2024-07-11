package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Message struct {
	Uuid           string `json:"uuid"`
	ContentType    string `json:"content_type"`
	MessageContent MessageContent
}

type MessageContent struct {
	SendAt   *timestamp.Timestamp `json:"send_at"`
	Provider string               `json:"provider"`
	Consumer string               `json:"consumer"`
	Title    string               `json:"title"`
	Content  string               `json:"content"`
}

func (m *Message) Marshal() ([]byte, error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(m)
	return result.Bytes(), err
}

func (m *Message) Unmarshal(data []byte) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	return dec.Decode(m)
}

func (m Message) String() string {
	return fmt.Sprintf(
		"UUID: %s, ContentType: %s, MessageContent: %s",
		m.Uuid, m.ContentType, m.MessageContent.String(),
	)
}

func (mc *MessageContent) Marshal() ([]byte, error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	err := enc.Encode(mc)
	return result.Bytes(), err
}

func (mc *MessageContent) Unmarshal(data []byte) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	return dec.Decode(mc)
}

func (mc MessageContent) String() string {
	return fmt.Sprintf(
		"SendAt: %v, Provider: %s, Consumer: %s, Title: %s, Content: %s",
		mc.SendAt, mc.Provider, mc.Consumer, mc.Title, mc.Content,
	)
}
