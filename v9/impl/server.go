package impl

import (
	"fmt"
	"net"
	"time"

	"github.com/ikun666/v9/conf"
	"github.com/ikun666/v9/iface"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	MsgHandle   iface.IMsgHandle
	ConnManager iface.IConnManager
}

func (s *Server) Serve() {

	defer s.Stop()
	s.Start()
	//TODO 处理

	//阻塞
	// select {}
}
func (s *Server) Start() {
	fmt.Println("start server")
	s.MsgHandle.StartWorkerPool()
	// addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%v:%v", s.IP, s.Port))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// listener, err := net.ListenTCP(s.IPVersion, addr)
	listener, err := net.Listen(s.IPVersion, fmt.Sprintf("%v:%v", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	var id uint32 = 0
	for {
		// conn, err := listener.AcceptTCP()
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("conn user:", s.ConnManager.Len())
		if s.ConnManager.Len() >= conf.GConfig.MaxConn {
			go DeferCloseConn(conn)
			continue
		}
		dealConn := NewConnetion(s, conn, id, s.MsgHandle)
		//add in there  remove in connection.stop()
		s.ConnManager.Add(dealConn)
		id++
		// fmt.Println("conn:", dealConn.GetConn())
		go dealConn.Start()

	}
}
func (s *Server) Stop() {
	fmt.Println("stop server")
	s.ConnManager.ClearConn()
}
func (s *Server) AddRouter(msgID uint32, router iface.IRouter) {
	s.MsgHandle.AddRouter(msgID, router)
}
func NewServer() iface.IServer {
	return &Server{
		Name:        conf.GConfig.ServerName,
		IPVersion:   conf.GConfig.IPVersion,
		IP:          conf.GConfig.IP,
		Port:        conf.GConfig.Port,
		MsgHandle:   NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
}
func (s *Server) GetConnManager() iface.IConnManager {
	return s.ConnManager
}
func DeferCloseConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("=======users too much,try to connect later======")
	data := []byte("users too much,try to connect later")
	msg := &Message{
		Len:  uint32(len(data)),
		ID:   401,
		Data: data,
	}
	pack := DataPack{}
	sendMsg, _ := pack.Pack(msg)
	conn.Write(sendMsg)
	time.Sleep(5 * time.Second)
}
