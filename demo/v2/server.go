package main

import (
	"github.com/ikun666/v2/impl"
)

func main() {
	sever := impl.NewServer("v2")
	sever.Serve()

}
