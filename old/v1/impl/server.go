package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/old/v1/iface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(conn *net.TCPConn) {
			defer conn.Close()
			buf := make([]byte, 512)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(buf[:n]))
				if _, err := conn.Write(buf[:n]); err != nil {
					fmt.Println(err)
					return
				}
			}
		}(conn)
	}
}
func (s *Server) Stop() {
	fmt.Println("stop server")
}
func NewServer(name string) iface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp",
		IP:        "127.0.0.1",
		Port:      8080,
	}
}
