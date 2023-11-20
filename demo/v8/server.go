package main

import (
	"fmt"

	"github.com/ikun666/v8/conf"
	"github.com/ikun666/v8/iface"
	"github.com/ikun666/v8/impl"
)

type PingRouter struct {
	impl.BaseRouter
}

func (p *PingRouter) PreHandle(req iface.IRequest) {
	// fmt.Println("pre handle ping")
}
func (p *PingRouter) Handle(req iface.IRequest) {
	// fmt.Println("handle ping")
	// _, err := req.GetConnetion().GetConn().Write([]byte("handle ping"))
	req.GetConnetion().Write(1, []byte(fmt.Sprintf("ping ping:%v", req.GetConnetion().GetID())))
}
func (p *PingRouter) PostHandle(req iface.IRequest) {
	// fmt.Println("post handle ping")
}

type HelloRouter struct {
	impl.BaseRouter
}

func (p *HelloRouter) Handle(req iface.IRequest) {
	// fmt.Println("handle hello")
	// _, err := req.GetConnetion().GetConn().Write([]byte("handle ping"))
	req.GetConnetion().Write(1, []byte(fmt.Sprintf("hello:%v", req.GetConnetion().GetID())))
}

func main() {
	err := conf.Init("../../v8/conf/conf.json")
	if err != nil {
		fmt.Printf("load conf err:%v", err)
		return
	}
	fmt.Println(conf.GConfig)
	sever := impl.NewServer()
	sever.AddRouter(0, &PingRouter{})
	sever.AddRouter(1, &HelloRouter{})
	sever.Serve()

}
