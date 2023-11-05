/*
 * @Author: hqtest
 * @Date: 2023-08-07 16:45:13
 * @Last Modified by: hqtest
 * @Last Modified time: 2023-08-24 14:06:38
 * @desc:
 */

package ws

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"github.com/gorilla/websocket"
	"net/http"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

type tgWServiceBase struct {
	fExit          context.CancelFunc ///
	pWSConn        *websocket.Conn    /// 连接句柄
	wsURL          string             /// url地址
	connStat       uint32             /// 0-未连接 1-连接
	fReciveMessage func(int, []byte)  /// 接收消息
	chanSendMsg    chan *TGWSMessage  /// 发送消息缓存
}

func NewTGWServiceBase(url string, f func(int, []byte)) *tgWServiceBase {
	return &tgWServiceBase{
		wsURL:          url,
		fReciveMessage: f,
		connStat:       0,
		chanSendMsg:    make(chan *TGWSMessage, 256),
	}
}

func (c *tgWServiceBase) StartWS() {
	ctx, fExit := context.WithCancel(context.Background())
	go c.serve(ctx)
	c.fExit = fExit
}

func (c *tgWServiceBase) StopWS() {
	if c.fExit != nil {
		c.fExit()
		c.fExit = nil
	}
	c.stopInner()
}

func (c *tgWServiceBase) PostMessage(msgType int, pBUFF []byte) error {
	wsConn := c.pWSConn
	if atomic.LoadUint32(&c.connStat) != 1 || wsConn == nil {
		return errors.New("stat is disconnect")
	}

	select {
	case c.chanSendMsg <- NewTGWSMessage(msgType, bytes.NewBuffer(pBUFF)):
	default:
		return errors.New("send cache is full")
	}
	return nil
}

func (c *tgWServiceBase) receiveMessage(wsConn *websocket.Conn) {
	for {
		msgType, pBUFF, e := wsConn.ReadMessage()
		if e != nil {
			log.Errorf("[WSBase] websocket readMessage failure, URL(%s), err(%v)", c.wsURL, e)
			c.stopInner()
			return
		}

		log.Debugf("[WSBase] recv message, type(%d), len(%d)", msgType, len(pBUFF))

		if c.fReciveMessage != nil {
			c.fReciveMessage(msgType, pBUFF)
		}
	}
}

func (c *tgWServiceBase) serve(ctx context.Context) {
	/// 定时检测连接是否正常
	fCheckWSConn := func() error {
		if atomic.LoadUint32(&c.connStat) == 1 {
			return nil
		}

		pDialer := &websocket.Dialer{
			Proxy:            http.ProxyFromEnvironment,
			HandshakeTimeout: 45 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		header := make(http.Header)

		wsConn, _, e := pDialer.DialContext(ctx, c.wsURL, header)
		if e != nil {
			return e
		}

		wsConn.SetCloseHandler(c.handleClose)
		wsConn.SetPingHandler(c.handlePing)
		wsConn.SetPongHandler(c.handlePong)

		go c.receiveMessage(wsConn)

		c.pWSConn = wsConn
		atomic.StoreUint32(&c.connStat, 1)

		log.Infof("[ConfigSVC] WS connect successful, url(%s)", c.wsURL)
		return nil
	}

	fPingTimer := func() {
		c.PostMessage(websocket.PingMessage, []byte(time.Now().String()))
	}

	tkTimer := time.NewTicker(time.Second * 5)
	defer tkTimer.Stop()

	tkPing := time.NewTicker(time.Second * 10)
	defer tkPing.Stop()

	for {
		select {
		case <-ctx.Done(): /// 退出信号
			return

		case <-tkTimer.C:
			e := fCheckWSConn()
			if e != nil {
				log.Warnf("[ConfigSVC] WS connect failure, url(%s), err(%v)", c.wsURL, e)
			}

		case <-tkPing.C:
			fPingTimer()

		case pMsg, ok := <-c.chanSendMsg:
			if !ok {
				return
			}
			c.postMessageInner(pMsg)
		}
	}
}

func (c *tgWServiceBase) stopInner() {
	atomic.StoreUint32(&c.connStat, 0)

	if c.pWSConn == nil {
		return
	}

	c.pWSConn.Close()
	c.pWSConn = nil
}

func (c *tgWServiceBase) postMessageInner(pMsg *TGWSMessage) error {
	wsConn := c.pWSConn
	if atomic.LoadUint32(&c.connStat) != 1 || wsConn == nil {
		return errors.New("stat is disconnect")
	}
	return wsConn.WriteMessage(pMsg.msgType, pMsg.GetMessage().Bytes())
}

func (c *tgWServiceBase) handleClose(code int, text string) error {
	c.stopInner()
	return nil
}

func (c *tgWServiceBase) handlePing(text string) error {
	return c.PostMessage(websocket.PongMessage, []byte(time.Now().String()))
}

func (c *tgWServiceBase) handlePong(text string) error {
	return nil
}
