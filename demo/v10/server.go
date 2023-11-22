package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikun666/v10/conf"
	"github.com/ikun666/v10/iface"
	"github.com/ikun666/v10/impl"
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

type IkunRouter struct {
	impl.BaseRouter
}

func (p *IkunRouter) Handle(req iface.IRequest) {
	// fmt.Println("handle hello")
	// _, err := req.GetConnetion().GetConn().Write([]byte("handle ping"))
	req.GetConnetion().Write(1, []byte(fmt.Sprintf("ikun:%v", req.GetConnetion().GetID())))
}

func OnConnCreate(conn iface.IConnection) {
	fmt.Println("----------OnConnCreate----------")
	conn.SetProperty("GitHub", "github.com/ikun666/GameServer")
	conn.Write(123, []byte("----------OnConnCreate----------"))
}
func OnConnDestroy(conn iface.IConnection) {
	fmt.Println("----------OnConnDestroy----------")
	if v, err := conn.GetProperty("GitHub"); err == nil {
		fmt.Println(v)
	}
	conn.Write(456, []byte("----------OnConnDestroy----------"))
}
func main() {
	err := conf.Init("../../v10/conf/conf.json")
	if err != nil {
		fmt.Printf("load conf err:%v", err)
		return
	}
	fmt.Println(conf.GConfig)
	server := impl.NewServer()

	server.SetOnConnCreate(OnConnCreate)
	server.SetOnConnDestroy(OnConnDestroy)

	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})
	server.AddRouter(2, &IkunRouter{})

	go server.Serve()

	// 创建一个接收shutdown信号的channel
	quit := make(chan os.Signal, 1)

	// signal.Notify函数用于将输入信号转发到quit
	// 如果收到指定的信号，将会推送到quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞主线程，直到接收到quit的信号
	<-quit

	server.Stop()
	time.Sleep(15 * time.Second)
	fmt.Println("Server exiting")

}
