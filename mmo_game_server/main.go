package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ikun666/mmo_game_server/apis"
	"github.com/ikun666/mmo_game_server/core"
	"github.com/ikun666/v10/conf"
	"github.com/ikun666/v10/iface"
	"github.com/ikun666/v10/impl"
)

func OnConnCreate(conn iface.IConnection) {
	player := core.NewPlayer(conn)
	//msgid=1
	player.SyncPid()

	//msgid=200
	player.BroadCastStartPos()

	core.WorldtMgr.AddPlayer(player)
	conn.SetProperty("pid", player.Pid)
	//msgid=202
	player.SyncSurrounding()
	fmt.Println("pid=", player.Pid)
}
func OnConnDestroy(conn iface.IConnection) {
	pid, err := conn.GetProperty("pid")
	if err != nil {
		fmt.Println(err)
		return
	}
	player := core.WorldtMgr.GetPlayer(pid.(int32))
	player.Offline()
}
func main() {
	err := conf.Init("./conf/conf.json")
	if err != nil {
		fmt.Printf("load conf err:%v", err)
		return
	}
	fmt.Println(conf.GConfig)
	server := impl.NewServer()

	server.SetOnConnCreate(OnConnCreate)
	server.SetOnConnDestroy(OnConnDestroy)

	server.AddRouter(2, &apis.ChatApi{})
	server.AddRouter(3, &apis.MoveApi{})
	go server.Serve()

	// 创建一个接收shutdown信号的channel
	quit := make(chan os.Signal, 1)

	// signal.Notify函数用于将输入信号转发到quit
	// 如果收到指定的信号，将会推送到quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞主线程，直到接收到quit的信号
	<-quit

	server.Stop()
	time.Sleep(5 * time.Second)
	fmt.Println("Server exiting")

}
