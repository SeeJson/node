package ws

import "bytes"

type TGWSMessage struct {
	msgType  int           /// 消息类型
	pMessage *bytes.Buffer /// 消息数据
}

func NewTGWSMessage(msgType int, pBUFF *bytes.Buffer) *TGWSMessage {
	return &TGWSMessage{
		msgType:  msgType,
		pMessage: pBUFF,
	}
}

func (c *TGWSMessage) GetMessageType() int {
	return c.msgType
}

func (c *TGWSMessage) GetMessage() *bytes.Buffer {
	return c.pMessage
}
