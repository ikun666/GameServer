package main

import (
	"fmt"

	"github.com/ikun666/v4/conf"
	"github.com/ikun666/v4/iface"
	"github.com/ikun666/v4/impl"
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
	err := conf.Init("../../v4/conf/conf.json")
	if err != nil {
		fmt.Printf("load conf err:%v", err)
		return
	}
	sever := impl.NewServer()
	sever.AddRouter(&PingRouter{})
	sever.Serve()

}
