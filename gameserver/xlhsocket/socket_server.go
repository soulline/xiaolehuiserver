package xlhsocket

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"xiaolehuigo/gameserver/util"
)

type MsfEventer interface {
	OnHandel(sessionId string, conn net.Conn) bool
	OnClose(sessionId string)
	OnMessage(sessionId string, msg string) bool
}

type Msf struct {
	SessionMaster *SessionM
	MsfEvent      MsfEventer
}

func NewMsf(msfEvent MsfEventer) *Msf {
	msf := &Msf{
		MsfEvent: msfEvent,
	}
	msf.SessionMaster = NewSessonM(msf)
	return msf
}

func (msf *Msf) Listen(address string) {
	netListen, err := net.Listen("tcp", address)
	logrus.Info("listen : " + address)
	CheckError(err)

	logrus.Info("Waiting for clients")
	go msf.SessionMaster.HeartBeat(60)
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		sessionId := "session_" + util.GetUUID()
		sessionP := NewSession(sessionId, conn)
		msf.SessionMaster.SetSession(sessionId, sessionP)
		logrus.Info(conn.RemoteAddr().String(), " tcp connect success")
		go HandleConnection(msf, sessionP)
	}
}

/**
处理连接
*/
func HandleConnection(msf *Msf, session *Session) {
	defer func() {
		msf.SessionMaster.DelSessionById(session.SessionId)
		//调用断开链接事件
		msf.MsfEvent.OnClose(session.SessionId)
	}()
	buffer := make([]byte, 1024)
	for {
		n, err := session.Con.Read(buffer)
		if err != nil {
			logrus.Info(session.Con.RemoteAddr().String(), " connection error: ", err)
			return
		}
		session.UpdateTime()
		msg := string(buffer[:n])
		msf.MsfEvent.OnMessage(session.SessionId, msg)
		logrus.Info(session.Con.RemoteAddr().String(), "receive data  ", msg)
	}

}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
