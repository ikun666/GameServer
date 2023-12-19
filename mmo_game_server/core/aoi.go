package core

import (
	"fmt"
	"sync"
)

type AOIManager struct {
	MinX  int
	MaxX  int
	CntX  int
	MinY  int
	MaxY  int
	CntY  int
	Grids map[int]*Grid
	Lock  sync.RWMutex
}

func NewAOIManager(minX, maxX, cntX, minY, maxY, cntY int) *AOIManager {
	aoi := AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntX:  cntX,
		MinY:  minY,
		MaxY:  maxY,
		CntY:  cntY,
		Grids: make(map[int]*Grid),
	}
	width := aoi.GetGridWidth()
	heigth := aoi.GetGridHeigth()
	for y := 0; y < cntY; y++ {
		for x := 0; x < cntX; x++ {
			gid := y*cntX + x
			aoi.Grids[gid] = NewGrid(gid,
				minX+x*width,
				minX+(x+1)*width,
				minY+y*heigth,
				minY+(y+1)*heigth)
		}

	}
	return &aoi
}
func (a *AOIManager) GetGridWidth() int {
	return (a.MaxX - a.MinX) / a.CntX
}
func (a *AOIManager) GetGridHeigth() int {
	return (a.MaxY - a.MinY) / a.CntY
}
func (a *AOIManager) String() string {
	s := fmt.Sprintf("MinX:%d MaxX:%d CntX:%d MinY:%d MaxY:%d CntY:%d\n", a.MinX, a.MaxX, a.CntX, a.MinY, a.MaxY, a.CntY)
	for _, v := range a.Grids {
		s += v.String()
	}
	return s
}
func (a *AOIManager) GetSurroundGridsByID(gid int) []*Grid {
	grids := make([]*Grid, 0, 9)
	if _, ok := a.Grids[gid]; !ok {
		return grids
	}
	//判断gid左右是否有grid
	grids = append(grids, a.Grids[gid])
	idx := gid % a.CntX
	if idx > 0 {
		grids = append(grids, a.Grids[gid-1])
	}
	if idx < a.CntX-1 {
		grids = append(grids, a.Grids[gid+1])
	}
	gx := make([]*Grid, 0, len(grids))
	gx = append(gx, grids...)
	// for _, v := range gx {
	// 	fmt.Print(v.GID, " ")
	// }
	//判断gx上下是否有grid
	for _, v := range gx {
		idy := v.GID / a.CntX
		if idy > 0 {
			grids = append(grids, a.Grids[v.GID-a.CntX])
		}
		if idy < a.CntY-1 {
			grids = append(grids, a.Grids[v.GID+a.CntX])
		}
	}
	return grids
}
func (a *AOIManager) GetGIDByPos(x, y int) int {
	idx := (x - a.MinX) / a.GetGridWidth()
	idy := (y - a.MinY) / a.GetGridHeigth()
	gid := idy*a.CntX + idx
	return gid
}
func (a *AOIManager) GetSurroundGridsPID(x, y int) (pids []int32) {
	gid := a.GetGIDByPos(x, y)
	// fmt.Println("gid:", gid)
	grids := a.GetSurroundGridsByID(gid)
	// fmt.Println("grids:", grids)
	for _, v := range grids {
		pids = append(pids, v.GetAllPID()...)
	}
	return
}
func (a *AOIManager) AddPID2Grid(pid int32, gid int) {
	a.Grids[gid].AddPlayerID(pid)
}
func (a *AOIManager) DeletePIDFromGrid(pid int32, gid int) {
	a.Grids[gid].DeletePlayerID(pid)
}
func (a *AOIManager) GetAllPIDByGrid(gid int) []int32 {
	return a.Grids[gid].GetAllPID()
}
func (a *AOIManager) AddPID2GridByPos(pid int32, x, y int) {
	gid := a.GetGIDByPos(x, y)
	a.Grids[gid].AddPlayerID(pid)
}
func (a *AOIManager) DeletePIDFromGridByPos(pid int32, x, y int) {
	gid := a.GetGIDByPos(x, y)
	a.Grids[gid].DeletePlayerID(pid)
}
