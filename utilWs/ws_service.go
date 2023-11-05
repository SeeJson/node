package ws

type IWebSocketClient interface {
	StartWS()
	StopWS()
	PostMessage(int, []byte) error
}
