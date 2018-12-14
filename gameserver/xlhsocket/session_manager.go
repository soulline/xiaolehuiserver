package xlhsocket

import (
	"sync"
	"time"
)

type SessionM struct {
	sessions map[string]*Session
	num      uint32
	lock     sync.RWMutex
	ser      *Msf
}

func NewSessonM(msf *Msf) *SessionM {
	return &SessionM{
		sessions: make(map[string]*Session),
		num:      0,
		ser:      msf,
	}
}

func (this *SessionM) GetSessionById(sessionId string) *Session {
	if v, exit := this.sessions[sessionId]; exit {
		return v
	}
	return nil
}

func (this *SessionM) SetSession(sessionId string, sess *Session) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.sessions[sessionId] = sess
}

//关闭连接并删除
func (this *SessionM) DelSessionById(sessionId string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if v, exit := this.sessions[sessionId]; exit {
		v.Close()
	}
	delete(this.sessions, sessionId)
}

//向所有客户端发送消息
func (this *SessionM) WriteToAll(msg []byte) {
	for i, _ := range this.sessions {
		this.WriteByid(i, msg)
	}
}

//向单个客户端发送信息
func (this *SessionM) WriteByid(sessionId string, msg []byte) bool {

	if v, exit := this.sessions[sessionId]; exit {
		if err := v.Write(string(msg)); err != nil {
			this.DelSessionById(sessionId)
			return false
		} else {
			return true
		}
	}
	return false
}

//心跳检测   每秒遍历一次 查看所有sess 上次接收消息时间  如果超过 num 就删除该 sess
func (this *SessionM) HeartBeat(num int64) {
	for {
		time.Sleep(time.Second)
		for i, v := range this.sessions {
			//logrus.Printf("sessions : %v, value : %v", i, v)
			//diff := time.Now().Unix() - v.Time
			//logrus.Println("time : ", diff)
			if time.Now().Unix()-v.Time > num {
				this.DelSessionById(i)
			}
		}
	}
}
