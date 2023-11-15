package impl

import (
	"fmt"
	"net"

	"github.com/ikun666/v3/iface"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iface.IRouter
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
	var id uint32 = 0
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//v1
		// go func(conn *net.TCPConn) {
		// 	defer conn.Close()
		// 	buf := make([]byte, 512)
		// 	for {
		// 		n, err := conn.Read(buf)
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			return
		// 		}
		// 		fmt.Println(string(buf[:n]))
		// 		if _, err := conn.Write(buf[:n]); err != nil {
		// 			fmt.Println(err)
		// 			return
		// 		}
		// 	}
		// }(conn)

		//v2
		dealConn := NewConnetion(conn, id, s.Router)
		id++
		go dealConn.Start()

	}
}
func (s *Server) Stop() {
	fmt.Println("stop server")
}
func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
}
func NewServer(name string) iface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp",
		IP:        "127.0.0.1",
		Port:      8080,
		Router:    nil,
	}
}

// // 自定义回调函数
// func Echo(conn *net.TCPConn, buf []byte, n int) error {
// 	_, err := conn.Write(buf[:n])
// 	if err != nil {
// 		fmt.Printf("echo err:%v", err)
// 		return fmt.Errorf("echo err:%v", err)
// 	}
// 	return nil
// }
