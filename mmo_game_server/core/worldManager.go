package core

import "sync"

type WorldManager struct {
	Players    map[int32]*Player
	PlayerLock sync.RWMutex
	AOIMgr     *AOIManager
}

var WorldtMgr *WorldManager

func init() {
	WorldtMgr = &WorldManager{
		Players: make(map[int32]*Player),
		AOIMgr:  NewAOIManager(0, 300, 30, 0, 300, 30),
	}
}
func (c *WorldManager) AddPlayer(player *Player) {
	c.AOIMgr.AddPID2GridByPos(player.Pid, int(player.X), int(player.Z))
	c.PlayerLock.Lock()
	defer c.PlayerLock.Unlock()
	c.Players[player.Pid] = player
}
func (c *WorldManager) DeletePlayer(player *Player) {
	c.AOIMgr.DeletePIDFromGridByPos(player.Pid, int(player.X), int(player.Z))
	c.PlayerLock.Lock()
	defer c.PlayerLock.Unlock()
	delete(c.Players, player.Pid)
}
func (c *WorldManager) GetPlayer(pid int32) *Player {
	c.PlayerLock.RLock()
	defer c.PlayerLock.RUnlock()
	return c.Players[pid]
}
func (c *WorldManager) GetAllPlayer() []*Player {
	c.PlayerLock.RLock()
	defer c.PlayerLock.RUnlock()
	players := make([]*Player, 0, len(c.Players))
	for _, v := range c.Players {
		players = append(players, v)
	}
	return players
}
