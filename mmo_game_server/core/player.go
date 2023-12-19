package core

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/ikun666/mmo_game_server/pb"
	"github.com/ikun666/v10/iface"
	"google.golang.org/protobuf/proto"
)

var PidGen int32 = 0
var PidGenLock sync.Mutex

type Player struct {
	Pid  int32
	X    float32 //unity坐标 x
	Y    float32 //z
	Z    float32 //y
	V    float32 //deg 0-360
	Conn iface.IConnection
}

func NewPlayer(conn iface.IConnection) *Player {
	player := &Player{
		X:    float32(100 + rand.Intn(20)),
		Y:    0,
		Z:    float32(100 + rand.Intn(20)),
		V:    0,
		Conn: conn,
	}
	PidGenLock.Lock()
	defer PidGenLock.Unlock()
	player.Pid = PidGen
	PidGen++
	return player
}
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = p.Conn.Write(msgId, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *Player) SyncPid() {
	pmsg := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, pmsg)
}
func (p *Player) BroadCastStartPos() {
	pmsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, pmsg)
}

// func (p *Player) BroadCastMsg(msg string) {
// 	pmsg := &pb.BroadCast{
// 		Pid: p.Pid,
// 		Tp:  1,
// 		Data: &pb.BroadCast_Content{
// 			Content: msg,
// 		},
// 	}
// 	p.SendMsg(200, pmsg)
// }

func (p *Player) SyncSurrounding() {
	//1 获取周围玩家
	pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(int(p.X), int(p.Z))
	//2 向周围玩家发送自己位置
	players := make([]*Player, 0, len(pids))
	for _, v := range pids {
		players = append(players, WorldtMgr.GetPlayer(v))
	}
	pmsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	for _, player := range players {
		player.SendMsg(200, pmsg)
	}
	//3 周围玩家向自己发送位置
	pmsgs := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		msg := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		pmsgs = append(pmsgs, msg)
	}
	syncmsg := &pb.SyncPlayers{
		Ps: pmsgs,
	}
	p.SendMsg(202, syncmsg)
}

func (p *Player) UpdatePlayerPos(x, y, z, v float32) {
	// //得到移动后的网格gid
	// newGid := WorldtMgr.AOIMgr.GetGIDByPos(int(x), int(z))

	// //1 获取原周围玩家
	// pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(int(p.X), int(p.Z))
	// //2 向周围玩家发送自己位置
	// players := make([]*Player, 0, len(pids))
	// for _, v := range pids {
	// 	players = append(players, WorldtMgr.GetPlayer(v))
	// }
	// pmsg := &pb.BroadCast{
	// 	Pid: p.Pid,
	// 	Tp:  4,
	// 	Data: &pb.BroadCast_P{
	// 		P: &pb.Position{
	// 			X: x,
	// 			Y: y,
	// 			Z: z,
	// 			V: v,
	// 		},
	// 	},
	// }
	// for _, player := range players {
	// 	player.SendMsg(200, pmsg)
	// }

	// if newGid != WorldtMgr.AOIMgr.GetGIDByPos(int(p.X), int(p.Z)) {
	// 	//删除原gid
	// 	WorldtMgr.AOIMgr.DeletePIDFromGridByPos(p.Pid, int(p.X), int(p.Z))
	// 	//添加新gid
	// 	WorldtMgr.AOIMgr.AddPID2Grid(p.Pid, newGid)
	// 	//1 获取新周围玩家
	// 	pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(int(x), int(z))
	// 	//2 向周围玩家发送自己位置
	// 	players := make([]*Player, 0, len(pids))
	// 	for _, v := range pids {
	// 		players = append(players, WorldtMgr.GetPlayer(v))
	// 	}
	// 	for _, player := range players {
	// 		player.SendMsg(200, pmsg)
	// 	}
	// }
	// //更新位置
	// p.X = x
	// p.Y = y
	// p.Z = z
	// p.V = v

	//1 获取周围玩家
	pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(int(x), int(z))
	//2 向周围玩家发送自己位置
	players := make([]*Player, 0, len(pids))
	for _, v := range pids {
		players = append(players, WorldtMgr.GetPlayer(v))
	}
	pmsg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: x,
				Y: y,
				Z: z,
				V: v,
			},
		},
	}
	for _, player := range players {
		player.SendMsg(200, pmsg)
	}
	//更新位置
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v
}
func (p *Player) GetSurroundPlayers() []*Player {
	//1 获取周围玩家
	pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(int(p.X), int(p.Z))
	players := make([]*Player, 0, len(pids))
	for _, v := range pids {
		players = append(players, WorldtMgr.GetPlayer(v))
	}
	return players
}
func (p *Player) Offline() {
	pmsg := &pb.SyncPid{
		Pid: p.Pid,
	}
	players := p.GetSurroundPlayers()
	for i, player := range players {
		fmt.Println("-------", i)
		fmt.Println(player)
		fmt.Println("-------", i)
		player.SendMsg(201, pmsg)
	}
	fmt.Println("aaaa")
	WorldtMgr.DeletePlayer(p)
	fmt.Println("bbbb")
	WorldtMgr.AOIMgr.DeletePIDFromGridByPos(p.Pid, int(p.X), int(p.Z))
	fmt.Println("cccc")
}
