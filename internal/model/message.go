package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Message struct {
	Uuid           string         `json:"uuid" gorm:"primaryKey;type:varchar(255)"`
	ContentType    string         `json:"content_type" gorm:"type:varchar(255)"`
	MessageContent MessageContent `gorm:"embedded;embeddedPrefix:message_content_"`
}

type MessageContent struct {
	SendAt   *timestamp.Timestamp `json:"send_at" gorm:"type:timestamp"`
	Provider string               `json:"provider" gorm:"type:varchar(255)"`
	Consumer string               `json:"consumer" gorm:"type:varchar(255)"`
	Title    string               `json:"title" gorm:"type:varchar(255)"`
	Content  string               `json:"content" gorm:"type:text"`
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
