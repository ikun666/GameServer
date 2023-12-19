package apis

import (
	"fmt"

	"github.com/ikun666/mmo_game_server/core"
	"github.com/ikun666/mmo_game_server/pb"
	"github.com/ikun666/v10/iface"
	"github.com/ikun666/v10/impl"

	"google.golang.org/protobuf/proto"
)

type ChatApi struct {
	impl.BaseRouter
}

func (c *ChatApi) Handle(req iface.IRequest) {
	pid, err := req.GetConnetion().GetProperty("pid")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := &pb.Talk{}
	err = proto.Unmarshal(req.GetMessage().GetMsgData(), msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	players := core.WorldtMgr.GetAllPlayer()
	for _, player := range players {
		pmsg := &pb.BroadCast{
			Pid: pid.(int32),
			Tp:  1,
			Data: &pb.BroadCast_Content{
				Content: msg.Content,
			},
		}
		player.SendMsg(200, pmsg)
	}

}
