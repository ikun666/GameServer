package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	go func() {
		buf := make([]byte, 512)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("recieve msg:%v\n", string(buf[:n]))
		}
	}()
	s := "hello ikun666"
	for {

		_, err = conn.Write([]byte(s))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}

}
