package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	PlayerIDs map[int32]struct{}
	Lock      sync.RWMutex
}

func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIDs: map[int32]struct{}{},
	}
}
func (g *Grid) AddPlayerID(pid int32) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	g.PlayerIDs[pid] = struct{}{}
}
func (g *Grid) DeletePlayerID(pid int32) {
	g.Lock.Lock()
	defer g.Lock.Unlock()
	delete(g.PlayerIDs, pid)
}
func (g *Grid) GetAllPID() []int32 {
	g.Lock.RLock()
	defer g.Lock.RUnlock()
	ids := make([]int32, 0, len(g.PlayerIDs))
	for k := range g.PlayerIDs {
		ids = append(ids, k)
	}
	return ids
}
func (g *Grid) String() string {

	return fmt.Sprintf("GID:%d MinX:%d MaxX:%d MinY:%d MaxY:%d\n", g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY)

}
