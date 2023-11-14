package main

import (
	"github.com/ikun666/v1/impl"
)

func main() {
	sever := impl.NewServer("v1")
	sever.Serve()

}
