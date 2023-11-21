package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/old/v8/conf"
	"github.com/ikun666/old/v8/iface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	MsgHandle iface.IMsgHandle
}

func (s *Server) Serve() {
	go s.Start()
	defer s.Stop()
	//TODO 处理

	//阻塞
	select {}
}
func (s *Server) Start() {
	fmt.Println("start server")
	s.MsgHandle.StartWorkerPool()
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%v:%v", s.IP, s.Port))
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	var id uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		dealConn := NewConnetion(conn, id, s.MsgHandle)
		id++
		// fmt.Println("conn:", dealConn.GetConn())
		go dealConn.Start()

	}
}
func (s *Server) Stop() {
	fmt.Println("stop server")
}
func (s *Server) AddRouter(msgID uint32, router iface.IRouter) {
	s.MsgHandle.AddRouter(msgID, router)
}
func NewServer() iface.IServer {
	return &Server{
		Name:      conf.GConfig.ServerName,
		IPVersion: "tcp",
		IP:        conf.GConfig.IP,
		Port:      conf.GConfig.Port,
		MsgHandle: NewMsgHandle(),
	}
}
