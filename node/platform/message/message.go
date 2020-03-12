package message

import "encoding/json"

type Message struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func FromBytes(b []byte) *Message {
	msg := &Message{}
	json.Unmarshal(b, msg)
	return msg
}

func Create(t string, v interface{}) *Message {
	return &Message{
		Type:  t,
		Value: v,
	}
}

func (m *Message) Get() *Message {
	return m
}

func (m *Message) ToBytes() []byte {
	v, _ := json.Marshal(m)
	return v
}

func (m *Message) ToString() string {
	v, _ := json.Marshal(m)
	return string(v)
}
