package core

import (
	"fmt"
	"testing"
)

//	func TestAOIManager(t *testing.T) {
//		aoi := NewAOIManager(100, 320, 3, 350, 550, 6)
//		fmt.Print(aoi)
//	}
func TestGetSuroundGridsByID(t *testing.T) {
	aoi := NewAOIManager(0, 300, 6, 0, 300, 6)
	fmt.Print(aoi)
	test0 := aoi.GetSurroundGridsByID(0)
	for _, v := range test0 {
		fmt.Print(v.GID, " ")
	}
	fmt.Println()
	test1 := aoi.GetSurroundGridsByID(1)
	for _, v := range test1 {
		fmt.Print(v.GID, " ")
	}
	fmt.Println()
	test2 := aoi.GetSurroundGridsByID(11)
	for _, v := range test2 {
		fmt.Print(v.GID, " ")
	}
	fmt.Println()
}
