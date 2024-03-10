package ws

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_tgWServiceBase_StartWS(t *testing.T) {
	socketUrl := "ws://localhost:8080/socket"
	tg := NewTGWServiceBase(socketUrl, dealWsMessage)
	tg.StartWS()

	time.Sleep(time.Second * 600)
	for i := 0; i < 10; i++ {
		if err := tg.PostMessage(websocket.TextMessage, []byte("xxxx")); nil != err {
			log.Errorf("err:[%s]", err)
		}
	}
}

func dealWsMessage(msgType int, msg []byte) {
	log.Infof("recv msg:[%s]", msg)
}
