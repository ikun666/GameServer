package core

import (
	"fmt"
	"testing"
)

func TestAOI(t *testing.T) {

	WorldtMgr.AOIMgr.AddPID2Grid(66, 0)
	WorldtMgr.AOIMgr.AddPID2Grid(77, 1)
	WorldtMgr.AOIMgr.AddPID2Grid(88, 2)
	WorldtMgr.AOIMgr.AddPID2Grid(99, 3)
	pids := WorldtMgr.AOIMgr.GetSurroundGridsPID(20, 0)
	fmt.Println(pids)
	players := make([]*Player, 0, len(pids))
	for _, v := range pids {
		fmt.Println(v)
		players = append(players, WorldtMgr.GetPlayer(v))
	}
	fmt.Println(players)
}
