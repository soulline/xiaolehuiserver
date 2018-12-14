package xlhsocket

import (
	"net"
	"sync"
	"time"
)

type Session struct {
	SessionId string
	Con       net.Conn
	Time      int64
	Lock      sync.Mutex
}

func NewSession(sessionId string, con net.Conn) *Session {
	return &Session{
		SessionId: sessionId,
		Con:       con,
		Time:      time.Now().Unix(),
	}
}

func (ses *Session) Write(msg string) error {
	ses.Lock.Lock()
	defer ses.Lock.Unlock()
	_, errs := ses.Con.Write([]byte(msg))
	return errs
}

func (ses *Session) Close() {
	ses.Con.Close()
}

func (ses *Session) UpdateTime() {
	ses.Time = time.Now().Unix()
}
