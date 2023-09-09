package webscockt

import (
	"github.com/olahol/melody"
	"net/http"
)

type handleMessageFunc func(*melody.Session, []byte)
type handleErrorFunc func(*melody.Session, error)
type handleCloseFunc func(*melody.Session, int, string) error
type handleSessionFunc func(*melody.Session)

func NewWebsocket(config *melody.Config) *WebsocketManage {

	instance := melody.New()
	if config != nil {
		instance.Config = config
	}
	wm := WebsocketManage{Instance: instance}
	return &wm
}

type WebsocketManage struct {
	Instance *melody.Melody
	Debug    bool
}

func (w *WebsocketManage) HandleRequest(write http.ResponseWriter, r *http.Request, keys map[string]any) error {

	if keys != nil {
		return w.Instance.HandleRequestWithKeys(write, r, keys)
	} else {
		return w.Instance.HandleRequest(write, r)
	}

}

func (w *WebsocketManage) HandleConnect(fn handleSessionFunc) {
	f := func(s *melody.Session) {
		if fn != nil {
			fn(s)
		}
	}
	w.Instance.HandleConnect(f)
}

// HandleMessage 注册 WebSocket 消息处理函数
func (w *WebsocketManage) HandleMessage(fn handleMessageFunc) {
	f := func(s *melody.Session, data []byte) {
		if fn != nil {
			fn(s, data)
		}
	}
	w.Instance.HandleMessage(f)
}

// 注册 WebSocket 连接关闭处理函数
func (w *WebsocketManage) handleDisconnect(fn handleSessionFunc) {

	f := func(s *melody.Session) {
		if fn != nil {
			fn(s)
		}
	}
	w.Instance.HandleDisconnect(f)
}

func (w *WebsocketManage) handlePong(fn handleSessionFunc) {
	f := func(s *melody.Session) {
		if fn != nil {
			fn(s)
		}
	}
	w.Instance.HandlePong(f)

}

func (w *WebsocketManage) HandleCloseFunc(fn handleCloseFunc) {
	f := func(s *melody.Session, code int, text string) error {
		if fn != nil {
			return fn(s, code, text)
		}
		return nil
	}
	w.Instance.HandleClose(f)

}

func (w *WebsocketManage) HandleError(fn handleErrorFunc) {
	f := func(s *melody.Session, err error) {
		if fn != nil {
			fn(s, err)
		}
	}
	w.Instance.HandleError(f)

}
