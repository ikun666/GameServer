package main

import (
	"fmt"

	"github.com/ikun666/v5/conf"
	"github.com/ikun666/v5/iface"
	"github.com/ikun666/v5/impl"
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
	req.GetConnetion().Write(1, []byte(fmt.Sprintf("hello:%v", req.GetConnetion().GetID())))
}
func (p *PingRouter) PostHandle(req iface.IRequest) {
	// fmt.Println("post handle ping")
}
func main() {
	err := conf.Init("../../v5/conf/conf.json")
	if err != nil {
		fmt.Printf("load conf err:%v", err)
		return
	}
	fmt.Println(conf.GConfig)
	sever := impl.NewServer()
	sever.AddRouter(&PingRouter{})
	sever.Serve()

}
