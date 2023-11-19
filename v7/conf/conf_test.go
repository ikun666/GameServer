package conf

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init("./conf.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", GConfig)
}
