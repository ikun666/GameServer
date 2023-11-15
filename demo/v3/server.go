package main

import (
	"fmt"

	"github.com/ikun666/v3/iface"
	"github.com/ikun666/v3/impl"
)

type PingRouter struct {
	impl.BaseRouter
}

func (p *PingRouter) PreHandle(req iface.IRequest) {
	fmt.Println("pre handle ping")
}
func (p *PingRouter) Handle(req iface.IRequest) {
	fmt.Println("handle ping")
	_, err := req.GetConnetion().GetConn().Write([]byte("handle ping"))
	if err != nil {
		fmt.Printf("handle ping err:=", err)
		return
	}
}
func (p *PingRouter) PostHandle(req iface.IRequest) {
	fmt.Println("post handle ping")
}
func main() {
	sever := impl.NewServer("v3")
	sever.AddRouter(&PingRouter{})
	sever.Serve()

}
