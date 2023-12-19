package apis

import (
	"fmt"

	"github.com/ikun666/mmo_game_server/core"
	"github.com/ikun666/mmo_game_server/pb"
	"github.com/ikun666/v10/iface"
	"github.com/ikun666/v10/impl"
	"google.golang.org/protobuf/proto"
)

type MoveApi struct {
	impl.BaseRouter
}

func (m *MoveApi) Handle(req iface.IRequest) {
	pid, err := req.GetConnetion().GetProperty("pid")
	if err != nil {
		fmt.Println(err)
		return
	}
	player := core.WorldtMgr.GetPlayer(pid.(int32))
	msg := &pb.Position{}
	err = proto.Unmarshal(req.GetMessage().GetMsgData(), msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	player.UpdatePlayerPos(msg.X, msg.Y, msg.Z, msg.V)

}
